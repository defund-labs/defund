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
	ibctransfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
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
	poolsOsmosisAppend = []string{
		"CtIC", "Cs0C", "Cs0C", "CsIC", "CtUC", "CokD", "CtEC", "Cs8C", "CpQD", "Cs8C", "Co4C",
		"CpED", "Co8D", "CosD", "Co4D", "CosD", "Cs4C", "CpID", "Cs4C", "CtQC", "Cs8C", "CtoC",
		"CsIC", "CtEC", "Co0D", "CtQC", "Cs8C", "CowD", "Cs0C", "CvwC", "Cs8C", "Cr0C", "Cs8C",
		"CssC", "CosD", "CtEC", "CtcC", "CowD", "Co0D", "CpUD", "CsAC", "Co8D", "Ct0C", "CtAC",
		"CtwC", "CsIC", "CtAC", "Co4D", "Co0D", "CowD", "CssC", "CssC", "Cs4C", "Co0D", "CpYD",
		"CswC", "CooD", "CrwC", "CvoC", "CpsD", "CosD", "CvoC", "Co0D", "CsIC", "CssC", "Cs8C",
		"Co0D", "Cs0C", "CpAD", "CooD", "CtEC", "CpAD", "CsUC", "CooD", "CswC", "CswC", "CooC",
		"CvwC", "CooD", "CpID", "Co4D", "CsMC", "Co0D", "Cs4C", "CoYD", "CsgC", "CsEC", "CscC",
		"CocD", "CooD",
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

func (s *IntegrationTestSuite) initOsmosisBroker() {
	var pools []*brokertypes.Source

	for i := range poolsOsmosis {
		app, err := base64.StdEncoding.DecodeString(poolsOsmosisAppend[i])
		s.Assert().NoError(err)
		addPool := brokertypes.Source{
			PoolId:       poolsOsmosis[i],
			InterqueryId: fmt.Sprintf("%s-%d", "osmosis", poolsOsmosis[i]),
			Status:       "active",
			Append:       app,
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

func (s *IntegrationTestSuite) TestGetOsmosisPool_Valid() {
	s.initOsmosisBroker()
	data, err := base64.StdEncoding.DecodeString("Chovb3Ntb3Npcy5nYW1tLnYxYmV0YTEuUG9vbBKzAgo/b3NtbzFtdzBhYzZyd2xwNXI4d2Fwd2szenM2ZzI5aDhmY3NjeHFha2R6dzllbWtuZTZjOHdqcDlxMHQzdjh0EAEaFQoQMjAwMDAwMDAwMDAwMDAwMBIBMCIDMjRoKioKC2dhbW0vcG9vbC8xEhszODQ4OTk3OTEzNDAyMzI0NjU1NzkzNjk0NjIyaApVCkRpYmMvMjczOTRGQjA5MkQyRUNDRDU2MTIzQzc0RjM2RTRDMUY5MjYwMDFDRUFEQTlDQTk3RUE2MjJCMjVGNDFFNUVCMhINNzY5MDA2NjIzODY0ORIPNTM2ODcwOTEyMDAwMDAwMioKFwoFdW9zbW8SDjI1MjM4OTEwOTUxNjQ1Eg81MzY4NzA5MTIwMDAwMDA6EDEwNzM3NDE4MjQwMDAwMDA=")
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
	_, err = s.GetDefundApp(s.chainA).BrokerKeeper.GetOsmosisPool(s.chainActx, 1)
	s.Assert().NoError(err)
}