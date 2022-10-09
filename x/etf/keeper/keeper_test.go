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
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	icatypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v4/modules/core/04-channel/types"
	"github.com/defund-labs/defund/app"
	ibctesting "github.com/defund-labs/defund/testing"
	brokertypes "github.com/defund-labs/defund/x/broker/types"
	"github.com/defund-labs/defund/x/etf/types"
	querytypes "github.com/defund-labs/defund/x/query/types"
	osmosisbalancertypes "github.com/osmosis-labs/osmosis/v8/x/gamm/pool-models/balancer"
	osmosisgammtypes "github.com/osmosis-labs/osmosis/v8/x/gamm/types"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"
)

type KeeperTestSuite struct {
	suite.Suite

	coordinator *ibctesting.Coordinator

	// testing chains used for convenience and readability
	chainA *ibctesting.TestChain
	chainB *ibctesting.TestChain

	queryClient ibctransfertypes.QueryClient
}

type GenesisState map[string]json.RawMessage

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

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func NewFundAddress(fundId string) sdk.AccAddress {
	key := append([]byte("etf"), []byte(fundId)...)
	return address.Module("etf", key)
}

func GetFundDenom(symbol string) string {
	return fmt.Sprintf("etf/%s", symbol)
}

func NewDefaultGenesisState(cdc codec.JSONCodec) GenesisState {
	return app.ModuleBasics.DefaultGenesis(cdc)
}

func SetDefundTestingApp() (ibctesting.TestingApp, map[string]json.RawMessage) {
	db := dbm.NewMemDB()
	encCdc := app.MakeEncodingConfig(app.ModuleBasics)
	appd := app.New(log.NewNopLogger(), db, nil, true, map[int64]bool{}, app.DefaultNodeHome, 5, encCdc, app.EmptyAppOptions{})
	gensisState := app.NewDefaultGenesisState(encCdc.Marshaler)
	return appd, gensisState
}

func (s *KeeperTestSuite) SetupTest() {
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

		s.GetDefundApp(chain).TransferKeeper.SetParams(chain.GetContext(), ibctransfertypes.DefaultParams())

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
}

func (s *KeeperTestSuite) GetDefundApp(chain *ibctesting.TestChain) *app.App {
	app, ok := chain.App.(*app.App)
	if !ok {
		panic("not defund app")
	}

	return app
}

func (s *KeeperTestSuite) NewTransferPath() *ibctesting.Path {
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
	return path
}

func (s *KeeperTestSuite) NewICAPath(TestPortID string, TestVersion string, clientIdA string, clientIdB string) *ibctesting.Path {
	// setup ica channels
	path := ibctesting.NewPath(s.chainA, s.chainB)
	path.EndpointA.ConnectionID = "connection-0"
	path.EndpointB.ConnectionID = "connection-0"
	path.EndpointA.ClientID = clientIdA
	path.EndpointB.ClientID = clientIdB
	path.EndpointA.ChannelConfig.PortID = TestPortID
	path.EndpointB.ChannelConfig.PortID = icatypes.PortID
	path.EndpointA.ChannelConfig.Order = channeltypes.ORDERED
	path.EndpointB.ChannelConfig.Order = channeltypes.ORDERED
	path.EndpointA.ChannelConfig.Version = TestVersion
	path.EndpointB.ChannelConfig.Version = TestVersion
	path.EndpointA.ClientID = "07-tendermint-0"
	path.EndpointB.ClientID = "07-tendermint-0"

	return path
}

func (s *KeeperTestSuite) CreateTestTokens() (atomCoin sdk.Coin, osmoCoin sdk.Coin, aktCoin sdk.Coin) {
	// create the ibc atom that lives on osmosis broker
	denomAtom := ibctransfertypes.DenomTrace{
		Path:      "transfer/channel-0",
		BaseDenom: "uatom",
	}
	// set the new denom trace in store
	s.GetDefundApp(s.chainA).TransferKeeper.SetDenomTrace(s.chainA.GetContext(), denomAtom)
	atomCoin = sdk.NewCoin(denomAtom.IBCDenom(), sdk.NewInt(50000000))

	// set the new denom trace in store
	osmoCoin = sdk.NewCoin("uosmo", sdk.NewInt(50000000))

	// create the ibc akt that lives on osmosis broker
	denomAkt := ibctransfertypes.DenomTrace{
		Path:      "transfer/channel-1",
		BaseDenom: "uakt",
	}
	// set the new denom trace in store
	s.GetDefundApp(s.chainA).TransferKeeper.SetDenomTrace(s.chainA.GetContext(), denomAkt)
	aktCoin = sdk.NewCoin(denomAkt.IBCDenom(), sdk.NewInt(50000000))

	// create test tokens, atom, osmo, akt
	s.GetDefundApp(s.chainA).BankKeeper.MintCoins(s.chainA.GetContext(), types.ModuleName, sdk.NewCoins(atomCoin))
	s.GetDefundApp(s.chainA).BankKeeper.MintCoins(s.chainA.GetContext(), types.ModuleName, sdk.NewCoins(osmoCoin))
	s.GetDefundApp(s.chainA).BankKeeper.MintCoins(s.chainA.GetContext(), types.ModuleName, sdk.NewCoins(aktCoin))

	return atomCoin, osmoCoin, aktCoin
}

func (s *KeeperTestSuite) CreateFundBalanceQuery(fund types.Fund, tokens []sdk.Coin, multiplier int64) {
	for i := range tokens {
		tokens[i].Amount.Mul(sdk.NewInt(multiplier))
	}
	balances := banktypes.Balance{
		Address: fund.Address,
		Coins:   sdk.Coins(tokens),
	}
	data, err := balances.Marshal()
	s.Assert().NoError(err)
	height := clienttypes.NewHeight(0, 0)
	query := querytypes.InterqueryResult{
		Creator:     s.chainA.SenderAccount.GetAddress().String(),
		Storeid:     fmt.Sprintf("balance-%s", fund.Address),
		Chainid:     "osmosis-1",
		Data:        data,
		Height:      &height,
		LocalHeight: uint64(0),
		Success:     true,
		Proved:      true,
	}
	err = s.GetDefundApp(s.chainA).QueryKeeper.SetInterqueryResult(s.chainA.GetContext(), query)
	s.Assert().NoError(err)
}

func (s *KeeperTestSuite) CreatePoolQueries(fund types.Fund) {
	// Create the Osmo Atom Pool Pair Interquery
	OsmoAtomPool := osmosisbalancertypes.Pool{
		Address: fund.Address,
		Id:      1,
		PoolParams: osmosisbalancertypes.PoolParams{
			SwapFee: sdk.NewDecWithPrec(20, 4),
			ExitFee: sdk.NewDecWithPrec(0, 4),
		},
		FuturePoolGovernor: "24h",
		TotalShares: sdk.Coin{
			Denom:  "gamm/pool/1",
			Amount: sdk.NewInt(int64(3849029151596493552)),
		},
		PoolAssets: []osmosisgammtypes.PoolAsset{
			{
				Token:  sdk.NewCoin("ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2", sdk.NewInt(7689908).Mul(sdk.NewInt(1000000))),
				Weight: sdk.NewInt(536870912000000),
			},
			{
				Token:  sdk.NewCoin("uosmo", sdk.NewInt(25239880).Mul(sdk.NewInt(1000000))),
				Weight: sdk.NewInt(536870912000000),
			},
		},
	}
	data, err := OsmoAtomPool.Marshal()
	s.Assert().NoError(err)
	height := clienttypes.NewHeight(0, 0)
	query := querytypes.InterqueryResult{
		Creator:     s.chainA.SenderAccount.GetAddress().String(),
		Storeid:     "osmosis-1",
		Chainid:     "osmosis-1",
		Data:        data,
		Height:      &height,
		LocalHeight: uint64(0),
		Success:     true,
		Proved:      true,
	}
	err = s.GetDefundApp(s.chainA).QueryKeeper.SetInterqueryResult(s.chainA.GetContext(), query)
	s.Assert().NoError(err)

	// Create the Osmo Akt Pool Pair Interquery
	OsmoAktPool := osmosisbalancertypes.Pool{
		Address: fund.Address,
		Id:      3,
		PoolParams: osmosisbalancertypes.PoolParams{
			SwapFee: sdk.NewDecWithPrec(20, 4),
			ExitFee: sdk.NewDecWithPrec(0, 4),
		},
		FuturePoolGovernor: "24h",
		TotalShares: sdk.Coin{
			Denom:  "gamm/pool/3",
			Amount: sdk.NewInt(int64(3849029151596493552)),
		},
		PoolAssets: []osmosisgammtypes.PoolAsset{
			{
				Token:  sdk.NewCoin("ibc/1480B8FD20AD5FCAE81EA87584D269547DD4D436843C1D20F15E00EB64743EF4", sdk.NewInt(4261789).Mul(sdk.NewInt(1000000))),
				Weight: sdk.NewInt(536870912000000),
			},
			{
				Token:  sdk.NewCoin("uosmo", sdk.NewInt(743575).Mul(sdk.NewInt(1000000))),
				Weight: sdk.NewInt(536870912000000),
			},
		},
	}
	data, err = OsmoAktPool.Marshal()
	s.Assert().NoError(err)
	height = clienttypes.NewHeight(0, 0)
	query = querytypes.InterqueryResult{
		Creator:     s.chainA.SenderAccount.GetAddress().String(),
		Storeid:     "osmosis-3",
		Chainid:     "osmosis-1",
		Data:        data,
		Height:      &height,
		LocalHeight: uint64(0),
		Success:     true,
		Proved:      true,
	}
	err = s.GetDefundApp(s.chainA).QueryKeeper.SetInterqueryResult(s.chainA.GetContext(), query)
	s.Assert().NoError(err)
}

// initTestFund creates a test fund in store and initializes all the requirements
// for this fund
func (s *KeeperTestSuite) CreateTestFund() types.Fund {

	// create test account address
	addr := NewFundAddress(fmt.Sprintf("defund1njx8c8yjfsj5g4xnzej9lfl2ugmhldfh4x8c5%d", setupAccountCounter))
	// add one to the account counter
	setupAccountCounter = setupAccountCounter.Add(sdk.OneInt())
	// create a new module account for the fund account
	acct := s.GetDefundApp(s.chainA).AccountKeeper.NewAccount(s.chainA.GetContext(), authtypes.NewModuleAccount(
		authtypes.NewBaseAccountWithAddress(
			addr,
		),
		addr.String(),
		"mint",
		"burn",
	))
	// set the new fund account in the store
	s.GetDefundApp(s.chainA).AccountKeeper.SetAccount(s.chainA.GetContext(), acct)

	// initialize and get the broker
	broker := s.CreateOsmosisBroker()

	// create the ica account
	err := s.GetDefundApp(s.chainA).BrokerKeeper.RegisterBrokerAccount(s.chainA.GetContext(), broker.ConnectionId, acct.GetAddress().String())
	s.Assert().NoError(err)

	// generate new portId for ica account
	portID, err := icatypes.NewControllerPortID(acct.GetAddress().String())
	s.Assert().NoError(err)

	// set the interchain accounts in store since IBC callback will not
	s.GetDefundApp(s.chainA).ICAControllerKeeper.SetActiveChannelID(s.chainA.GetContext(), broker.ConnectionId, portID, testChannelId)
	s.GetDefundApp(s.chainA).ICAControllerKeeper.SetInterchainAccountAddress(s.chainA.GetContext(), broker.ConnectionId, portID, acct.GetAddress().String())

	// init all the tokens. returns all the initialized coins that were sent to module
	testAtomIBC, testOsmoIBC, testAktIBC := s.CreateTestTokens()

	holdingOne := types.Holding{
		Token:    testOsmoIBC.Denom,
		Percent:  34,
		PoolId:   1,
		BrokerId: "osmosis",
		Type:     "spot",
	}
	holdingTwo := types.Holding{
		Token:    testAtomIBC.Denom,
		Percent:  33,
		PoolId:   1,
		BrokerId: "osmosis",
		Type:     "spot",
	}
	holdingThree := types.Holding{
		Token:    testAktIBC.Denom,
		Percent:  33,
		PoolId:   3,
		BrokerId: "osmosis",
		Type:     "spot",
	}
	// add the holdings as slice of holdings
	holdings := []types.Holding{holdingOne, holdingTwo, holdingThree}

	// create the test fund
	TestFund := types.Fund{
		Symbol:        testFundSymbol,
		Address:       acct.GetAddress().String(),
		Name:          testFundName,
		Description:   testFundDesc,
		Shares:        sdk.NewCoin(GetFundDenom(testFundSymbol), sdk.NewInt(5000000)),
		Holdings:      holdings,
		BaseDenom:     baseDenom,
		Rebalance:     10,
		StartingPrice: sdk.NewCoin(baseDenom, sdk.NewInt(5000000)),
	}
	// set the test fund in store
	s.GetDefundApp(s.chainA).EtfKeeper.SetFund(s.chainA.GetContext(), TestFund)

	return TestFund
}

func (s *KeeperTestSuite) CreateOsmosisBroker() brokertypes.Broker {
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
		Status:       "inactive",
	}

	s.GetDefundApp(s.chainA).BrokerKeeper.SetBroker(s.chainA.GetContext(), broker)

	return broker
}

// RegisterInterchainAccount is a helper function for starting the channel handshake
func (s *KeeperTestSuite) RegisterInterchainAccount(endpoint *ibctesting.Endpoint, owner string, TestVersion string) error {
	portID, err := icatypes.NewControllerPortID(owner)
	if err != nil {
		return err
	}

	channelSequence := endpoint.Chain.App.GetIBCKeeper().ChannelKeeper.GetNextChannelSequence(endpoint.Chain.GetContext())

	if err := s.GetDefundApp(s.chainA).ICAControllerKeeper.RegisterInterchainAccount(endpoint.Chain.GetContext(), endpoint.ConnectionID, owner, TestVersion); err != nil {
		return err
	}

	// commit state changes for proof verification
	endpoint.Chain.NextBlock()

	// update port/channel ids
	endpoint.ChannelID = channeltypes.FormatChannelIdentifier(channelSequence)
	endpoint.ChannelConfig.PortID = portID

	return nil
}

func (s *KeeperTestSuite) SetupICAPath(path *ibctesting.Path, owner string, TestVersion string) error {
	if err := s.RegisterInterchainAccount(path.EndpointA, owner, TestVersion); err != nil {
		return err
	}

	if err := path.EndpointB.ChanOpenTry(); err != nil {
		return err
	}

	if err := path.EndpointA.ChanOpenAck(); err != nil {
		return err
	}

	if err := path.EndpointB.ChanOpenConfirm(); err != nil {
		return err
	}

	return nil
}

func (s *KeeperTestSuite) TestETFFundActions() {
	TestOwnerAddress := s.chainA.SenderAccount.GetAddress()
	TestPortID, _ := icatypes.NewControllerPortID(TestOwnerAddress.String())
	// TestVersion defines a reusable interchainaccounts version string for testing purposes
	TestVersion := string(icatypes.ModuleCdc.MustMarshalJSON(&icatypes.Metadata{
		Version:                icatypes.Version,
		ControllerConnectionId: "connection-0",
		HostConnectionId:       "connection-0",
		Encoding:               icatypes.EncodingProtobuf,
		TxType:                 icatypes.TxTypeSDKMultiMsg,
	}))

	path := s.NewTransferPath()
	s.coordinator.SetupClients(path)
	s.coordinator.SetupConnections(path)
	s.coordinator.CreateTransferChannels(path)

	pathIca := s.NewICAPath(TestPortID, TestVersion, path.EndpointA.ClientID, path.EndpointB.ClientID)

	err := path.EndpointA.UpdateClient()
	s.Assert().NoError(err)
	err = path.EndpointA.Counterparty.UpdateClient()
	s.Assert().NoError(err)
	err = path.EndpointB.UpdateClient()
	s.Assert().NoError(err)
	err = path.EndpointB.Counterparty.UpdateClient()
	s.Assert().NoError(err)

	err = s.SetupICAPath(pathIca, TestOwnerAddress.String(), TestVersion)
	s.Assert().NoError(err)

	s.coordinator.CommitBlock(s.chainA)

	s.Run("CheckHoldings", func() {
		s.CreateOsmosisBroker()

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
		err = s.GetDefundApp(s.chainA).QueryKeeper.SetInterqueryResult(s.chainA.GetContext(), interquery)
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
		err = s.GetDefundApp(s.chainA).QueryKeeper.SetInterqueryResult(s.chainA.GetContext(), interquery)
		s.Assert().NoError(err)

		holdings := []types.Holding{
			{
				Token:    "ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2",
				Percent:  33,
				PoolId:   1,
				BrokerId: "osmosis",
				Type:     "spot",
			},
			{
				Token:    "uosmo",
				Percent:  34,
				PoolId:   1,
				BrokerId: "osmosis",
				Type:     "spot",
			},
			{
				Token:    "ibc/1480B8FD20AD5FCAE81EA87584D269547DD4D436843C1D20F15E00EB64743EF4",
				Percent:  33,
				PoolId:   4,
				BrokerId: "osmosis",
				Type:     "spot",
			},
		}
		err = s.GetDefundApp(s.chainA).EtfKeeper.CheckHoldings(s.chainA.GetContext(), holdings)
		s.Assert().NoError(err)
	})

	s.Run("Create", func() {
		fund := s.CreateTestFund()
		atomCoin, osmoCoin, aktCoin := s.CreateTestTokens()
		// add them to an account balance
		err := s.GetDefundApp(s.chainA).BankKeeper.SendCoinsFromModuleToAccount(s.chainA.GetContext(), types.ModuleName, s.chainA.SenderAccounts[1].SenderAccount.GetAddress(), sdk.NewCoins(atomCoin))
		s.Assert().NoError(err)
		err = s.GetDefundApp(s.chainA).BankKeeper.SendCoinsFromModuleToAccount(s.chainA.GetContext(), types.ModuleName, s.chainA.SenderAccounts[1].SenderAccount.GetAddress(), sdk.NewCoins(osmoCoin))
		s.Assert().NoError(err)
		err = s.GetDefundApp(s.chainA).BankKeeper.SendCoinsFromModuleToAccount(s.chainA.GetContext(), types.ModuleName, s.chainA.SenderAccounts[1].SenderAccount.GetAddress(), sdk.NewCoins(aktCoin))
		s.Assert().NoError(err)
		// create the fake balance query for fund
		s.CreateFundBalanceQuery(fund, []sdk.Coin{atomCoin, osmoCoin, aktCoin}, 1)
		s.CreatePoolQueries(fund)
		tokenIn := sdk.NewCoin(fund.BaseDenom, sdk.NewInt(20000000))
		_, err = s.GetDefundApp(s.chainA).EtfKeeper.CreateShares(s.chainA.GetContext(), fund, "channel-0", tokenIn, s.chainA.SenderAccounts[1].SenderAccount.GetAddress().String(), clienttypes.NewHeight(0, 100), 0)
		s.Assert().NoError(err)
	})

	s.Run("Redeem", func() {
		fund, found := s.GetDefundApp(s.chainA).EtfKeeper.GetFund(s.chainA.GetContext(), "test")
		s.Assert().True(found)
		tokenInRedeem := sdk.NewCoin(fund.Shares.Denom, sdk.NewInt(89753))
		addrMap := types.AddressMap{
			OsmosisAddress: s.chainA.SenderAccount.GetAddress().String(),
		}

		pathIca.EndpointA.UpdateClient()
		pathIca.EndpointB.UpdateClient()
		s.coordinator.CommitBlock(s.chainA, s.chainB)

		// redeem the created shares from above
		err := s.GetDefundApp(s.chainA).EtfKeeper.RedeemShares(s.chainA.GetContext(), s.chainA.SenderAccounts[1].SenderAccount.GetAddress().String(), fund, path.EndpointA.ChannelID, tokenInRedeem, addrMap)
		s.Assert().NoError(err)
	})

	s.Run("CreateRebalanceMsgs", func() {})
}
