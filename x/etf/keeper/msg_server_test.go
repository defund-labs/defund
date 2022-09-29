package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/defund-labs/defund/testutil/keeper"
	"github.com/defund-labs/defund/x/etf/keeper"
	"github.com/defund-labs/defund/x/etf/types"
	"github.com/stretchr/testify/require"
	dbm "github.com/tendermint/tm-db"
)

func (s *KeeperTestSuite) TestFundMsgServerCreate() {
	db := dbm.NewMemDB()
	k, ctx := keepertest.EtfKeeper(db, s.T())
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateFund{Creator: creator}
		_, err := srv.CreateFund(wctx, expected)
		require.NoError(s.T(), err)
	}
}

func (s *KeeperTestSuite) TestSharesMsgServerCreate() {
	db := dbm.NewMemDB()
	k, ctx := keepertest.EtfKeeper(db, s.T())
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)

	path := s.NewTransferPath()
	s.coordinator.SetupClients(path)
	s.coordinator.SetupConnections(path)
	s.coordinator.CreateChannels(path)

	creator := s.chainA.SenderAccount.GetAddress().String()
	fund := s.CreateTestFund()
	token := sdk.NewCoin(fund.BaseDenom, sdk.NewInt(10000000))

	s.coordinator.CommitBlock(s.chainA)

	expected := &types.MsgCreate{
		Creator:          creator,
		Fund:             fund.Symbol,
		TokenIn:          &token,
		Channel:          path.EndpointA.ChannelID,
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
