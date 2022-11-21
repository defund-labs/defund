package wasmbinding_test

import (
	"encoding/json"
	"testing"

	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
	"github.com/defund-labs/defund/app"
	ibctesting "github.com/defund-labs/defund/testing"
	brokerwasm "github.com/defund-labs/defund/x/broker/client/wasm"
	etfwasm "github.com/defund-labs/defund/x/etf/client/wasm"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"
)

type WasmTestSuite struct {
	suite.Suite

	coordinator *ibctesting.Coordinator

	chainA *ibctesting.TestChain

	queryClient ibctransfertypes.QueryClient

	etfwasmq    *etfwasm.EtfWasmQueryHandler
	brokerwasmq *brokerwasm.BrokerWasmQueryHandler
}

func TestWasmTestSuite(t *testing.T) {
	suite.Run(t, new(WasmTestSuite))
}

func SetDefundTestingApp() (ibctesting.TestingApp, map[string]json.RawMessage) {
	db := dbm.NewMemDB()
	encCdc := app.MakeEncodingConfig(app.ModuleBasics)
	appd := app.New(log.NewNopLogger(), db, nil, true, map[int64]bool{}, app.DefaultNodeHome, 5, encCdc, app.GetEnabledProposals(), app.EmptyAppOptions{}, []wasm.Option{})
	gensisState := app.NewDefaultGenesisState(encCdc.Marshaler)
	return appd, gensisState
}

func (s *WasmTestSuite) GetDefundApp(chain *ibctesting.TestChain) *app.App {
	app, ok := chain.App.(*app.App)
	if !ok {
		panic("not defund app")
	}

	return app
}

func (s *WasmTestSuite) SetupTest() {
	s.coordinator = ibctesting.NewCoordinator(s.T(), 1)

	chains := make(map[string]*ibctesting.TestChain)

	ibctesting.DefaultTestingAppInit = SetDefundTestingApp

	// create a chain with the temporary coordinator that we'll later override
	chainID := "testchain0"
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

	s.coordinator.Chains = chains
	s.chainA = s.coordinator.GetChain(ibctesting.GetChainID(0))

	defundApp := s.GetDefundApp(s.chainA)

	queryHelper := baseapp.NewQueryServerTestHelper(s.chainA.GetContext(), defundApp.InterfaceRegistry())
	ibctransfertypes.RegisterQueryServer(queryHelper, defundApp.TransferKeeper)
	s.queryClient = ibctransfertypes.NewQueryClient(queryHelper)

	s.etfwasmq = etfwasm.NewEtfWasmQueryHandler(&s.GetDefundApp(s.chainA).EtfKeeper)
	s.brokerwasmq = brokerwasm.NewEtfWasmQueryHandler(&s.GetDefundApp(s.chainA).BrokerKeeper)
}

func (s *WasmTestSuite) TestWasmContracts() {}
