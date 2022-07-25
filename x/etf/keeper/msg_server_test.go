package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/defund-labs/defund/testutil/keeper"
	"github.com/defund-labs/defund/x/etf/keeper"
	"github.com/defund-labs/defund/x/etf/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.EtfKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func TestFundMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.EtfKeeper(t)
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
	k, ctx := keepertest.EtfKeeper(t)
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
	k, ctx := keepertest.EtfKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateFund{Creator: creator}
		_, err := srv.CreateFund(wctx, expected)
		require.NoError(t, err)
	}
}
