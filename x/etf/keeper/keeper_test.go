package keeper_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	"github.com/defund-labs/defund/app"
	ibctesting "github.com/defund-labs/defund/testing"
	brokertypes "github.com/defund-labs/defund/x/broker/types"
	"github.com/defund-labs/defund/x/etf/types"
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
		1, 678, 704, 712, 497, 674, 604, 9, 498, 584, 3, 10, 601, 2, 722, 611, 719, 585, 738, 13,
		4, 482, 481, 6, 577, 5, 463, 629, 641, 690, 15, 461, 560, 586, 587, 42, 600, 627, 608, 571,
		631, 548, 7, 605, 572, 648, 606, 643, 8, 597, 619, 553, 625, 602, 618, 574, 578, 651, 626, 573,
		22, 555, 637, 681, 464, 645, 644, 596, 547, 616, 558, 621, 613, 197, 679, 617, 670, 612, 638, 561,
		567, 649, 732, 653, 633, 557, 706, 662, 615, 701, 565, 669, 562, 592, 693, 151, 183, 695, 726, 673,
		549, 716, 624, 731, 718, 642, 721, 640, 734, 713, 725, 710, 737, 729, 700, 707, 717, 676,
		579, 682, 580, 730,
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
	app := app.New(log.NewNopLogger(), db, nil, true, map[int64]bool{}, app.DefaultNodeHome, 0, encCdc, app.EmptyAppOptions{})
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
	atomCoin = sdk.NewCoin(denomAtom.GetFullDenomPath(), sdk.NewInt(100000000000))

	// create the denom for osmo that lives on osmosis broker
	denomOsmo := ibctransfertypes.DenomTrace{
		Path:      "transfer/channel-0",
		BaseDenom: "uosmo",
	}
	// set the new denom trace in store
	s.GetDefundApp(s.chainA).TransferKeeper.SetDenomTrace(s.chainActx, denomOsmo)
	osmoCoin = sdk.NewCoin(denomOsmo.GetFullDenomPath(), sdk.NewInt(100000000000))

	// create the ibc akt that lives on osmosis broker
	denomAkt := ibctransfertypes.DenomTrace{
		Path:      "transfer/channel-0/transfer/channel-0",
		BaseDenom: "uakt",
	}
	// set the new denom trace in store
	s.GetDefundApp(s.chainA).TransferKeeper.SetDenomTrace(s.chainActx, denomAkt)
	aktCoin = sdk.NewCoin(denomAkt.GetFullDenomPath(), sdk.NewInt(100000000000))

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

	for _, pool := range poolsOsmosis {
		addPool := brokertypes.Source{
			PoolId:       pool,
			InterqueryId: fmt.Sprintf("%s-%d", "osmosis", pool),
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

func (s *IntegrationTestSuite) TestCreateShares_Valid() {
	// setup transfer channels
	path := ibctesting.NewPath(s.chainA, s.chainB)
	path.EndpointA.ChannelConfig.PortID = ibctesting.TransferPort
	path.EndpointB.ChannelConfig.PortID = ibctesting.TransferPort
	path.EndpointA.ChannelConfig.Version = "ics20-1"
	path.EndpointB.ChannelConfig.Version = "ics20-1"
	s.coordinator.Setup(path)

	s.SetupTest()
	s.initTestTokens()
	s.initOsmosisBroker()
	s.initTestFund()

	// create a unique address
	setupAccountCounter = setupAccountCounter.Add(sdk.OneInt())
	addr := sdk.AccAddress([]byte("addr_______________" + setupAccountCounter.String()))

	// register the account in AccountKeeper
	acct := s.GetDefundApp(s.chainA).AccountKeeper.NewAccountWithAddress(s.chainActx, addr)
	s.GetDefundApp(s.chainA).AccountKeeper.SetAccount(s.chainActx, acct)

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

	// check each token is send enabled
	for _, token := range tokens {
		s.Assert().True(s.GetDefundApp(s.chainA).BankKeeper.IsSendEnabledCoin(s.chainActx, *token))
	}

	// try to create etf shares with keeper function
	err := s.GetDefundApp(s.chainA).EtfKeeper.CreateShares(s.chainActx, fund, "channel-0", tokens, addr.String(), clienttypes.NewHeight(uint64(0), uint64(s.chainActx.BlockHeight()+100)), 0)
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) TestRedeemShares_Valid() {}

func (s *IntegrationTestSuite) TestRebalanceShares_Valid() {}
