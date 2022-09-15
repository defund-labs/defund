package keeper_test

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	"github.com/cosmos/cosmos-sdk/types/query"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	icatypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	"github.com/defund-labs/defund/app"
	ibctesting "github.com/defund-labs/defund/testing"
	brokertypes "github.com/defund-labs/defund/x/broker/types"
	"github.com/defund-labs/defund/x/etf/types"
	querytypes "github.com/defund-labs/defund/x/query/types"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"
)

type IntegrationTestSuite struct {
	suite.Suite

	coordinator *ibctesting.Coordinator

	// testing chains used for convenience and readability
	chainA *ibctesting.TestChain
	chainB *ibctesting.TestChain

	chainActx sdk.Context
	chainBctx sdk.Context

	queryClient ibctransfertypes.QueryClient
}

type GenesisState map[string]json.RawMessage

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

var (
	setupAccountCounter = sdk.ZeroInt()
	testFundSymbol      = "test"
	testFundDesc        = "test"
	testFundName        = "test"
	baseDenom           = "uosmo"
	testChannelId       = "channel-0"

	poolsOsmosis = []uint64{
		1, 497, 674, 604, 9, 498, 584, 3, 10, 601, 2, 611, 585, 13, 4, 482, 481, 6, 577, 5, 463,
		629, 641, 15, 461, 560, 586, 587, 42, 600, 627, 608, 571, 631, 548, 7, 605, 572, 648,
		606, 643, 8, 597, 619, 553, 625, 602, 618, 574, 578, 651, 626, 573, 22, 555, 637, 464,
		645, 644, 596, 547, 616, 558, 621, 613, 197, 617, 670, 612, 638, 561, 567, 649, 653,
		633, 557, 662, 615, 565, 562, 592, 151, 183, 673, 549, 624, 642,
	}
)

func NewFundAddress(fundId string) sdk.AccAddress {
	key := append([]byte("etf"), []byte(fundId)...)
	return address.Module("etf", key)
}

func GetFundDenom(symbol string) string {
	return fmt.Sprintf("etf/pool/%s", symbol)
}

func NewDefaultGenesisState(cdc codec.JSONCodec) GenesisState {
	return app.ModuleBasics.DefaultGenesis(cdc)
}

func SetDefundTestingApp() (ibctesting.TestingApp, map[string]json.RawMessage) {
	db := dbm.NewMemDB()
	encCdc := app.MakeEncodingConfig(app.ModuleBasics)
	app := app.New(log.NewNopLogger(), db, nil, true, map[int64]bool{}, app.DefaultNodeHome, 5, encCdc, app.EmptyAppOptions{})
	return app, NewDefaultGenesisState(encCdc.Marshaler)
}

func (s *IntegrationTestSuite) SetupTest() {
	s.coordinator = ibctesting.NewCoordinator(s.T(), 0)

	chains := make(map[string]*ibctesting.TestChain)
	for i := 0; i < 2; i++ {
		ibctesting.DefaultTestingAppInit = SetDefundTestingApp

		// create a chain with the temporary coordinator that we'll later override
		chainID := ibctesting.GetChainID(i)
		chain := ibctesting.NewTestChain(s.T(), ibctesting.NewCoordinator(s.T(), 0), chainID)

		balance := banktypes.Balance{
			Address: chain.SenderAccount.GetAddress().String(),
			Coins:   sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100000000000000))),
		}

		// create application and override files in the IBC test chain
		app := ibctesting.SetupWithGenesisValSet(
			s.T(),
			chain.Vals,
			[]authtypes.GenesisAccount{
				chain.SenderAccount.(authtypes.GenesisAccount),
			},
			chainID,
			sdk.DefaultPowerReduction,
			balance,
		)

		chain.App = app
		chain.QueryServer = app.GetIBCKeeper()
		chain.TxConfig = app.GetTxConfig()
		chain.Codec = app.AppCodec()
		chain.CurrentHeader = tmproto.Header{
			ChainID: chainID,
			Height:  1,
			Time:    s.coordinator.CurrentTime.UTC(),
		}

		chain.Coordinator = s.coordinator
		s.coordinator.CommitBlock(chain)

		chains[chainID] = chain

		s.GetDefundApp(chain).TransferKeeper.SetParams(chain.GetContext(), ibctransfertypes.DefaultParams())
	}

	s.coordinator.Chains = chains
	s.chainA = s.coordinator.GetChain(ibctesting.GetChainID(0))
	s.chainB = s.coordinator.GetChain(ibctesting.GetChainID(1))

	defundApp := s.GetDefundApp(s.chainA)

	queryHelper := baseapp.NewQueryServerTestHelper(s.chainA.GetContext(), defundApp.InterfaceRegistry())
	ibctransfertypes.RegisterQueryServer(queryHelper, defundApp.TransferKeeper)
	s.queryClient = ibctransfertypes.NewQueryClient(queryHelper)

	s.chainActx = s.chainA.GetContext()
	s.chainBctx = s.chainB.GetContext()
}

func (s *IntegrationTestSuite) GetDefundApp(chain *ibctesting.TestChain) *app.App {
	app, ok := chain.App.(*app.App)
	if !ok {
		panic("not defund app")
	}

	return app
}

func (s *IntegrationTestSuite) initTestTokens() (atomCoin sdk.Coin, osmoCoin sdk.Coin, aktCoin sdk.Coin) {
	// create the ibc atom that lives on osmosis broker
	denomAtom := ibctransfertypes.DenomTrace{
		Path:      "transfer/channel-0/transfer/channel-0",
		BaseDenom: "uatom",
	}

	// set the new denom trace in store
	s.GetDefundApp(s.chainA).TransferKeeper.SetDenomTrace(s.chainActx, denomAtom)
	s.GetDefundApp(s.chainA).TransferKeeper.SetParams(s.chainActx, ibctransfertypes.DefaultParams())
	atomCoin = sdk.NewCoin(denomAtom.IBCDenom(), sdk.NewInt(100000000000))

	// create the denom for osmo that lives on osmosis broker
	denomOsmo := ibctransfertypes.DenomTrace{
		Path:      "transfer/channel-0",
		BaseDenom: "uosmo",
	}
	// set the new denom trace in store
	s.GetDefundApp(s.chainA).TransferKeeper.SetDenomTrace(s.chainActx, denomOsmo)
	s.GetDefundApp(s.chainA).TransferKeeper.SetParams(s.chainActx, ibctransfertypes.DefaultParams())
	osmoCoin = sdk.NewCoin(denomOsmo.IBCDenom(), sdk.NewInt(100000000000))

	// create the ibc akt that lives on osmosis broker
	denomAkt := ibctransfertypes.DenomTrace{
		Path:      "transfer/channel-1/transfer/channel-0",
		BaseDenom: "uakt",
	}
	// set the new denom trace in store
	s.GetDefundApp(s.chainA).TransferKeeper.SetDenomTrace(s.chainActx, denomAkt)
	s.GetDefundApp(s.chainA).TransferKeeper.SetParams(s.chainActx, ibctransfertypes.DefaultParams())
	aktCoin = sdk.NewCoin(denomAkt.IBCDenom(), sdk.NewInt(100000000000))

	// create test tokens, atom, osmo, akt
	s.GetDefundApp(s.chainA).BankKeeper.MintCoins(s.chainActx, types.ModuleName, sdk.NewCoins(atomCoin))
	s.GetDefundApp(s.chainA).BankKeeper.MintCoins(s.chainActx, types.ModuleName, sdk.NewCoins(osmoCoin))
	s.GetDefundApp(s.chainA).BankKeeper.MintCoins(s.chainActx, types.ModuleName, sdk.NewCoins(aktCoin))

	return atomCoin, osmoCoin, aktCoin
}

// initTestFund creates a test fund in store and initializes all the requirements
// for this fund
func (s *IntegrationTestSuite) initTestFund() {

	// create test account address
	addr := NewFundAddress(fmt.Sprintf("defund1njx8c8yjfsj5g4xnzej9lfl2ugmhldfh4x8c5%d", setupAccountCounter))
	// add one to the account counter
	setupAccountCounter = setupAccountCounter.Add(sdk.OneInt())
	// create a new module account for the fund account
	acct := s.GetDefundApp(s.chainA).AccountKeeper.NewAccount(s.chainActx, authtypes.NewModuleAccount(
		authtypes.NewBaseAccountWithAddress(
			addr,
		),
		addr.String(),
		"mint",
		"burn",
	))
	// set the new fund account in the store
	s.GetDefundApp(s.chainA).AccountKeeper.SetAccount(s.chainActx, acct)

	// get the osmosis broker
	broker, found := s.GetDefundApp(s.chainA).BrokerKeeper.GetBroker(s.chainActx, "osmosis")
	s.Assert().True(found)

	// create the ica account
	err := s.GetDefundApp(s.chainA).BrokerKeeper.RegisterBrokerAccount(s.chainActx, broker.ConnectionId, acct.GetAddress().String())
	s.Assert().NoError(err)

	// generate new portId for ica account
	portID, err := icatypes.NewControllerPortID(acct.GetAddress().String())
	s.Assert().NoError(err)

	// set the interchain accounts in store since IBC callback will not
	s.GetDefundApp(s.chainA).ICAControllerKeeper.SetActiveChannelID(s.chainActx, broker.BaseDenom, portID, testChannelId)
	s.GetDefundApp(s.chainA).ICAControllerKeeper.SetInterchainAccountAddress(s.chainActx, broker.ConnectionId, portID, acct.GetAddress().String())

	// init all the tokens. returns all the initialized coins that were sent to module
	testAtomIBC, testOsmoIBC, testAktIBC := s.initTestTokens()

	// create the holdings to add to fund
	holdingOne := types.Holding{
		Token:    testOsmoIBC.Denom,
		Percent:  34,
		PoolId:   1,
		BrokerId: "osmosis",
	}
	holdingTwo := types.Holding{
		Token:    testAtomIBC.Denom,
		Percent:  33,
		PoolId:   1,
		BrokerId: "osmosis",
	}
	holdingThree := types.Holding{
		Token:    testAktIBC.Denom,
		Percent:  33,
		PoolId:   1,
		BrokerId: "osmosis",
	}
	// add the holdings as slice of holdings
	holdings := []types.Holding{holdingOne, holdingTwo, holdingThree}

	// create the test fund
	TestFund := types.Fund{
		Symbol:        testFundSymbol,
		Address:       acct.GetAddress().String(),
		Name:          testFundName,
		Description:   testFundDesc,
		Shares:        sdk.NewCoin(GetFundDenom(testFundSymbol), sdk.ZeroInt()),
		Holdings:      holdings,
		BaseDenom:     baseDenom,
		Rebalance:     10,
		StartingPrice: sdk.NewCoin(baseDenom, sdk.NewInt(5000000)),
	}
	// set the test fund in store
	s.GetDefundApp(s.chainA).EtfKeeper.SetFund(s.chainActx, TestFund)
}

func (s *IntegrationTestSuite) initOsmosisBroker() {
	var pools []*brokertypes.Source

	for i := range poolsOsmosis {
		addPool := brokertypes.Source{
			PoolId:       poolsOsmosis[i],
			InterqueryId: fmt.Sprintf("%s-%d", "osmosis", poolsOsmosis[i]),
			Status:       "active",
		}
		pools = append(pools, &addPool)
	}

	broker := brokertypes.Broker{
		Id:           "osmosis",
		ConnectionId: "connection-0",
		Pools:        pools,
		BaseDenom:    "uosmo",
		Status:       "inactive",
	}

	s.GetDefundApp(s.chainA).BrokerKeeper.SetBroker(s.chainActx, broker)
}

func (s *IntegrationTestSuite) TestCheckHoldings_Valid() {
	s.initOsmosisBroker()

	data, err := base64.StdEncoding.DecodeString("Chovb3Ntb3Npcy5nYW1tLnYxYmV0YTEuUG9vbBKzAgo/b3NtbzFtdzBhYzZyd2xwNXI4d2Fwd2szenM2ZzI5aDhmY3NjeHFha2R6dzllbWtuZTZjOHdqcDlxMHQzdjh0EAEaFQoQMjAwMDAwMDAwMDAwMDAwMBIBMCIDMjRoKioKC2dhbW0vcG9vbC8xEhszODQ4OTk3OTEzNDAyMzI0NjU1NzkzNjk0NjIyaApVCkRpYmMvMjczOTRGQjA5MkQyRUNDRDU2MTIzQzc0RjM2RTRDMUY5MjYwMDFDRUFEQTlDQTk3RUE2MjJCMjVGNDFFNUVCMhINNzY5MDAzNjMwNTU0MhIPNTM2ODcwOTEyMDAwMDAwMioKFwoFdW9zbW8SDjI1MjM5MDA5NTkzODk3Eg81MzY4NzA5MTIwMDAwMDA6EDEwNzM3NDE4MjQwMDAwMDA=")
	s.Assert().NoError(err)
	height := clienttypes.NewHeight(0, 0)
	interquery := querytypes.InterqueryResult{
		Creator:     "defund1y295kyv2upsy6swhj0dulghf208ngec5k7zpjq",
		Storeid:     "osmosis-1",
		Chainid:     "osmosis-1",
		Data:        data,
		Height:      &height,
		LocalHeight: 0,
		Success:     true,
		Proved:      true,
	}
	err = s.GetDefundApp(s.chainA).QueryKeeper.SetInterqueryResult(s.chainActx, interquery)
	s.Assert().NoError(err)

	data, err = base64.StdEncoding.DecodeString("Chovb3Ntb3Npcy5nYW1tLnYxYmV0YTEuUG9vbBLvAgo/b3NtbzFsend2MGdsY2hmY3cwZnB3emR3ZmRzZXBtdmx1djZ6NmVoNHF1bnhkbWwzM3NqMDZxM3lxN3h3dGRlEAQaFQoQMzAwMDAwMDAwMDAwMDAwMBIBMCIDMjRoKikKC2dhbW0vcG9vbC80Eho1MjE3NzY4NjUwNzYzMzUwMzYzNjczMzUzOTJoClUKRGliYy8xNDgwQjhGRDIwQUQ1RkNBRTgxRUE4NzU4NEQyNjk1NDdERDRENDM2ODQzQzFEMjBGMTVFMDBFQjY0NzQzRUY0Eg00NjI2NTYyOTEwMzg5Eg83MDg2Njk2MDM4NDAwMDAyZwpUCkRpYmMvMjczOTRGQjA5MkQyRUNDRDU2MTIzQzc0RjM2RTRDMUY5MjYwMDFDRUFEQTlDQTk3RUE2MjJCMjVGNDFFNUVCMhIMMTI2NDE3MzA2NDY0Eg8zNjUwNzIyMjAxNjAwMDA6EDEwNzM3NDE4MjQwMDAwMDA=")
	s.Assert().NoError(err)
	interquery = querytypes.InterqueryResult{
		Creator:     "defund1y295kyv2upsy6swhj0dulghf208ngec5k7zpjq",
		Storeid:     "osmosis-4",
		Chainid:     "osmosis-4",
		Data:        data,
		Height:      &height,
		LocalHeight: 0,
		Success:     true,
		Proved:      true,
	}
	err = s.GetDefundApp(s.chainA).QueryKeeper.SetInterqueryResult(s.chainActx, interquery)
	s.Assert().NoError(err)

	holdings := []types.Holding{
		{
			Token:    "ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2",
			Percent:  33,
			PoolId:   1,
			BrokerId: "osmosis",
		},
		{
			Token:    "uosmo",
			Percent:  34,
			PoolId:   1,
			BrokerId: "osmosis",
		},
		{
			Token:    "ibc/1480B8FD20AD5FCAE81EA87584D269547DD4D436843C1D20F15E00EB64743EF4",
			Percent:  33,
			PoolId:   4,
			BrokerId: "osmosis",
		},
	}
	err = s.GetDefundApp(s.chainA).EtfKeeper.CheckHoldings(s.chainActx, holdings)
	s.Assert().NoError(err)
}

func (s *IntegrationTestSuite) TestCreateShares_Valid() {
	s.SetupTest()

	// setup transfer channels
	path := ibctesting.NewPath(s.chainA, s.chainB)
	path.EndpointA.ChannelConfig.PortID = "transfer"
	path.EndpointB.ChannelConfig.PortID = "transfer"
	path.EndpointA.ChannelConfig.Version = "ics20-1"
	path.EndpointB.ChannelConfig.Version = "ics20-1"
	path.EndpointA.ConnectionID = "connection-0"
	path.EndpointB.ConnectionID = "connection-0"
	path.EndpointA.ChannelID = "channel-0"
	path.EndpointB.ChannelID = "channel-0"
	path.EndpointA.ClientID = "07-tendermint-0"
	path.EndpointB.ClientID = "07-tendermint-0"
	s.coordinator.SetupConnections(path)
	s.coordinator.CreateChannels(path)

	s.initTestTokens()
	s.initOsmosisBroker()
	s.initTestFund()

	// create a unique address
	addr := s.chainA.SenderAccount.GetAddress()

	// init all the tokens. returns all the initialized coins that were sent to module
	testAtomIBC, testOsmoIBC, testAktIBC := s.initTestTokens()

	// add them to an account balance
	s.GetDefundApp(s.chainA).BankKeeper.SendCoinsFromModuleToAccount(s.chainActx, types.ModuleName, addr, sdk.NewCoins(testAtomIBC))
	s.GetDefundApp(s.chainA).BankKeeper.SendCoinsFromModuleToAccount(s.chainActx, types.ModuleName, addr, sdk.NewCoins(testOsmoIBC))
	s.GetDefundApp(s.chainA).BankKeeper.SendCoinsFromModuleToAccount(s.chainActx, types.ModuleName, addr, sdk.NewCoins(testAktIBC))

	// get test fund from store
	fund, found := s.GetDefundApp(s.chainA).EtfKeeper.GetFund(s.chainActx, testFundName)
	s.Assert().True(found)

	tokens := []*sdk.Coin{&testAtomIBC, &testOsmoIBC, &testAktIBC}

	// ensure token transer sends are enabled
	s.Assert().True(s.GetDefundApp(s.chainA).TransferKeeper.GetSendEnabled(s.chainActx))
	s.Assert().True(s.GetDefundApp(s.chainB).TransferKeeper.GetSendEnabled(s.chainBctx))

	// check each token is send enabled and balance has some tokens
	for i := range tokens {
		s.Assert().True(s.GetDefundApp(s.chainA).TransferKeeper.GetSendEnabled(s.chainActx))
		s.Assert().True(s.GetDefundApp(s.chainA).BankKeeper.IsSendEnabledCoin(s.chainActx, *tokens[i]))
		balanceQuery := banktypes.NewQueryAllBalancesRequest(s.chainA.SenderAccount.GetAddress(), &query.PageRequest{Offset: 0})
		context := sdk.WrapSDKContext(s.chainActx)
		res, err := s.GetDefundApp(s.chainA).BankKeeper.AllBalances(context, balanceQuery)
		s.Assert().NoError(err)
		bal := res.Balances.AmountOf(tokens[i].Denom)
		add := sdk.NewCoin(tokens[i].Denom, bal)
		tokens[i] = &add
	}

	// try to create etf shares with keeper function
	err := s.GetDefundApp(s.chainA).EtfKeeper.CreateShares(s.chainActx, fund, "channel-0", tokens, addr.String(), clienttypes.NewHeight(uint64(0), uint64(s.chainActx.BlockHeight()+100)), 0)
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) TestRedeemShares_Valid() {}

func (s *IntegrationTestSuite) TestRebalanceShares_Valid() {}
