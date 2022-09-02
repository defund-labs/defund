package keeper_test

import (
	"fmt"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	connectiontypes "github.com/cosmos/ibc-go/v3/modules/core/03-connection/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	"github.com/defund-labs/defund/app"
	brokertypes "github.com/defund-labs/defund/x/broker/types"
	"github.com/defund-labs/defund/x/etf/types"
	"github.com/stretchr/testify/suite"
	abci "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtypes "github.com/tendermint/tendermint/types"
)

type IntegrationTestSuite struct {
	suite.Suite

	ctx         sdk.Context
	app         *app.App
	queryClient types.QueryClient
}

var defaultConsensusParams = &abci.ConsensusParams{
	Block: &abci.BlockParams{
		MaxBytes: 200000,
		MaxGas:   2000000,
	},
	Evidence: &tmproto.EvidenceParams{
		MaxAgeNumBlocks: 302400,
		MaxAgeDuration:  504 * time.Hour, // 3 weeks is the max duration
		MaxBytes:        10000,
	},
	Validator: &tmproto.ValidatorParams{
		PubKeyTypes: []string{
			tmtypes.ABCIPubKeyTypeEd25519,
		},
	},
}

var (
	setupAccountCounter = sdk.ZeroInt()
	testFundSymbol      = "test"
	testFundDesc        = "test"
	testFundName        = "test"
	baseDenom           = "uosmo"
	testConnectionId    = "connection-0"
	testChannelId       = "channel-0"
	testPortId          = "icacontroller"

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

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

func (s *IntegrationTestSuite) SetupTest() {
}

func (s *IntegrationTestSuite) initTestTokens() (atomCoin sdk.Coin, osmoCoin sdk.Coin, aktCoin sdk.Coin) {
	// create the ibc atom that lives on osmosis broker
	denomAtom := ibctransfertypes.DenomTrace{
		Path:      "transfer/channel-0/transfer/channel-0",
		BaseDenom: "uatom",
	}
	// set the new denom trace in store
	s.app.TransferKeeper.SetDenomTrace(s.ctx, denomAtom)
	atomCoin = sdk.NewCoin(denomAtom.GetFullDenomPath(), sdk.NewInt(100000000000))

	// create the denom for osmo that lives on osmosis broker
	denomOsmo := ibctransfertypes.DenomTrace{
		Path:      "transfer/channel-0",
		BaseDenom: "uosmo",
	}
	// set the new denom trace in store
	s.app.TransferKeeper.SetDenomTrace(s.ctx, denomOsmo)
	osmoCoin = sdk.NewCoin(denomOsmo.GetFullDenomPath(), sdk.NewInt(100000000000))

	// create the ibc akt that lives on osmosis broker
	denomAkt := ibctransfertypes.DenomTrace{
		Path:      "transfer/channel-0/transfer/channel-1",
		BaseDenom: "uakt",
	}
	// set the new denom trace in store
	s.app.TransferKeeper.SetDenomTrace(s.ctx, denomAkt)
	aktCoin = sdk.NewCoin(denomAkt.GetFullDenomPath(), sdk.NewInt(100000000000))

	// create test tokens, atom, osmo, akt
	s.app.BankKeeper.MintCoins(s.ctx, types.ModuleName, sdk.NewCoins(atomCoin))
	s.app.BankKeeper.MintCoins(s.ctx, types.ModuleName, sdk.NewCoins(osmoCoin))
	s.app.BankKeeper.MintCoins(s.ctx, types.ModuleName, sdk.NewCoins(aktCoin))

	return atomCoin, osmoCoin, aktCoin
}

// initTestFund creates a test fund in store and initializes all the requirements
// for this fund
func (s *IntegrationTestSuite) initTestFund() {
	// create the test connection
	connection := connectiontypes.ConnectionEnd{}
	// set the connection in store
	s.app.IBCKeeper.ConnectionKeeper.SetConnection(s.ctx, testConnectionId, connection)

	// create the test channel
	TestChannel := channeltypes.Channel{
		ConnectionHops: []string{testConnectionId},
	}
	// set the test channel in store
	s.app.IBCKeeper.ChannelKeeper.SetChannel(s.ctx, testPortId, testChannelId, TestChannel)

	// create test account address
	addr := NewFundAddress(fmt.Sprintf("defund1njx8c8yjfsj5g4xnzej9lfl2ugmhldfh4x8c5%d", setupAccountCounter))
	// add one to the account counter
	setupAccountCounter = setupAccountCounter.Add(sdk.OneInt())
	// create a new module account for the fund account
	acct := s.app.AccountKeeper.NewAccount(s.ctx, authtypes.NewModuleAccount(
		authtypes.NewBaseAccountWithAddress(
			addr,
		),
		addr.String(),
		"mint",
		"burn",
	))
	// set the new fund account in the store
	s.app.AccountKeeper.SetAccount(s.ctx, acct)

	// get the osmosis broker
	broker, found := s.app.BrokerKeeper.GetBroker(s.ctx, "osmosis")
	s.Assert().True(found)

	// create the ica account
	err := s.app.BrokerKeeper.RegisterBrokerAccount(s.ctx, broker.ConnectionId, acct.GetAddress().String())
	s.Assert().NoError(err)

	// generate new portId for ica account
	portID, err := icatypes.NewControllerPortID(acct.GetAddress().String())
	s.Assert().NoError(err)

	// set the interchain accounts in store since IBC callback will not
	s.app.ICAControllerKeeper.SetActiveChannelID(s.ctx, broker.BaseDenom, portID, "channel-7")
	s.app.ICAControllerKeeper.SetInterchainAccountAddress(s.ctx, broker.ConnectionId, portID, acct.GetAddress().String())

	// init all the tokens. returns all the initialized coins that were sent to module
	testAtomIBC, testOsmoIBC, testAktIBC := s.initTestTokens()

	// create the holdings to add to fund
	holdingOne := types.Holding{
		Token:   testOsmoIBC.Denom,
		Percent: 34,
		PoolId:  1,
	}
	holdingTwo := types.Holding{
		Token:   testAtomIBC.Denom,
		Percent: 33,
		PoolId:  1,
	}
	holdingThree := types.Holding{
		Token:   testAktIBC.Denom,
		Percent: 33,
		PoolId:  1,
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
	s.app.EtfKeeper.SetFund(s.ctx, TestFund)
}

func (s *IntegrationTestSuite) initOsmosisBroker() {
	var pools []*brokertypes.Pool

	for _, pool := range poolsOsmosis {
		addPool := brokertypes.Pool{
			PoolId:       pool,
			InterqueryId: fmt.Sprintf("%s-%d", "osmosis", pool),
			Status:       "active",
		}
		pools = append(pools, &addPool)
	}

	broker := brokertypes.Broker{
		Id:        "osmosis",
		Pools:     pools,
		BaseDenom: "uosmo",
		Status:    "inactive",
	}

	s.app.BrokerKeeper.SetBroker(s.ctx, broker)
}

func (s *IntegrationTestSuite) TestCreateShares_Valid() {
	// create a unique address
	setupAccountCounter = setupAccountCounter.Add(sdk.OneInt())
	addr := sdk.AccAddress([]byte("addr_______________" + setupAccountCounter.String()))

	// register the account in AccountKeeper
	acct := s.app.AccountKeeper.NewAccountWithAddress(s.ctx, addr)
	s.app.AccountKeeper.SetAccount(s.ctx, acct)

	// init all the tokens. returns all the initialized coins that were sent to module
	testAtomIBC, testOsmoIBC, testAktIBC := s.initTestTokens()

	// add them to an account balance
	s.app.BankKeeper.SendCoinsFromModuleToAccount(s.ctx, types.ModuleName, addr, sdk.NewCoins(testAtomIBC))
	s.app.BankKeeper.SendCoinsFromModuleToAccount(s.ctx, types.ModuleName, addr, sdk.NewCoins(testOsmoIBC))
	s.app.BankKeeper.SendCoinsFromModuleToAccount(s.ctx, types.ModuleName, addr, sdk.NewCoins(testAktIBC))

	// get test fund from store
	fund, found := s.app.EtfKeeper.GetFund(s.ctx, testFundName)
	s.Assert().True(found)

	tokens := []*sdk.Coin{&testAtomIBC, &testOsmoIBC, &testAktIBC}

	// try to create etf shares with keeper function
	err := s.app.EtfKeeper.CreateShares(s.ctx, fund, testChannelId, tokens, addr.String(), clienttypes.NewHeight(uint64(0), uint64(s.ctx.BlockHeight()+100)), 0)
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) TestRedeemShares_Valid() {}

func (s *IntegrationTestSuite) TestRebalanceShares_Valid() {}
