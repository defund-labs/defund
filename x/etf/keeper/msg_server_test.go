package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/defund-labs/defund/testutil/keeper"
	"github.com/defund-labs/defund/x/etf/keeper"
	"github.com/defund-labs/defund/x/etf/types"
	"github.com/stretchr/testify/require"
	dbm "github.com/tendermint/tm-db"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	db := dbm.NewMemDB()
	k, ctx := keepertest.EtfKeeper(db, t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func TestFundMsgServerCreate(t *testing.T) {
	db := dbm.NewMemDB()
	k, ctx := keepertest.EtfKeeper(db, t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateFund{Creator: creator}
		_, err := srv.CreateFund(wctx, expected)
		require.NoError(t, err)
	}
}

func TestSharesMsgServerCreate(t *testing.T) {
	db := dbm.NewMemDB()
	k, ctx := keepertest.EtfKeeper(db, t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateFund{Creator: creator}
		_, err := srv.CreateFund(wctx, expected)
		require.NoError(t, err)
	}
}

func TestSharesMsgServerRedeem(t *testing.T) {
	db := dbm.NewMemDB()
	k, ctx := keepertest.EtfKeeper(db, t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateFund{Creator: creator}
		_, err := srv.CreateFund(wctx, expected)
		require.NoError(t, err)
	}
}
