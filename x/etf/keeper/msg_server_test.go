package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/defund-labs/defund/testutil/keeper"
	"github.com/defund-labs/defund/x/etf/keeper"
	"github.com/defund-labs/defund/x/etf/types"
	"github.com/stretchr/testify/require"
	dbm "github.com/tendermint/tm-db"
)

func (s *KeeperTestSuite) setup(ctx sdk.Context) (outctx sdk.Context, fund types.Fund, connectionId string, portId string) {
	path := s.NewTransferPath()
	s.Require().Equal(path.EndpointA.ChannelID, "channel-0")

	fund = s.CreateTestFund()
	// Commit new block to store info
	s.coordinator.CommitBlock(s.chainA, s.chainB)
	// We must create an ICA channel here on the broker chain with the test fund address as the owner
	connectionId, portId = s.CreateChannelICA(fund.Address, path)
	accAddress, found := s.GetDefundApp(s.chainA).ICAControllerKeeper.GetInterchainAccountAddress(ctx, connectionId, portId)
	s.Assert().True(found)
	atomCoin, osmoCoin, aktCoin := s.CreateTestTokens()
	// add them to an account balance
	err := s.GetDefundApp(s.chainA).BankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, s.chainA.SenderAccounts[1].SenderAccount.GetAddress(), sdk.NewCoins(atomCoin))
	s.Assert().NoError(err)
	err = s.GetDefundApp(s.chainA).BankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, s.chainA.SenderAccounts[1].SenderAccount.GetAddress(), sdk.NewCoins(osmoCoin))
	s.Assert().NoError(err)
	err = s.GetDefundApp(s.chainA).BankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, s.chainA.SenderAccounts[1].SenderAccount.GetAddress(), sdk.NewCoins(aktCoin))
	s.Assert().NoError(err)
	// create the fake balance query for fund
	s.CreateFundBalanceQuery(accAddress, []sdk.Coin{atomCoin, osmoCoin, aktCoin}, 1)
	s.CreatePoolQueries(fund)

	outctx = ctx

	return outctx, fund, connectionId, portId
}

func (s *KeeperTestSuite) TestFundMsgServerCreate() {
	db := dbm.NewMemDB()
	k, ctx := keepertest.EtfKeeper(db, s.T())
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	s.setup(ctx)
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateFund{
			Creator:   creator,
			Holdings:  "ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2:33:osmosis:1:spot,uosmo:34:osmosis:1:spot,ibc/1480B8FD20AD5FCAE81EA87584D269547DD4D436843C1D20F15E00EB64743EF4:33:osmosis:4:spot",
			BaseDenom: "uosmo",
		}
		_, err := srv.CreateFund(wctx, expected)
		require.NoError(s.T(), err)
	}
}

func (s *KeeperTestSuite) TestSharesMsgServerCreate() {
	db := dbm.NewMemDB()
	k, ctx := keepertest.EtfKeeper(db, s.T())
	srv := keeper.NewMsgServerImpl(*k)
	ctx, fund, _, _ := s.setup(ctx)
	wctx := sdk.WrapSDKContext(ctx)
	creator := s.chainA.SenderAccount.GetAddress().String()
	token := sdk.NewCoin(fund.BaseDenom, sdk.NewInt(10000000))

	s.coordinator.CommitBlock(s.chainA)

	expected := &types.MsgCreate{
		Creator:          creator,
		Fund:             fund.Symbol,
		TokenIn:          &token,
		Channel:          testChannelId,
		TimeoutHeight:    "50",
		TimeoutTimestamp: 0,
	}
	_, err := srv.Create(wctx, expected)
	require.NoError(s.T(), err)
}

func (s *KeeperTestSuite) TestSharesMsgServerRedeem() {
	db := dbm.NewMemDB()
	k, ctx := keepertest.EtfKeeper(db, s.T())
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgRedeem{Creator: creator}
		_, err := srv.Redeem(wctx, expected)
		require.NoError(s.T(), err)
	}
}
