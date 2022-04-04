package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/defund-labs/defund/v1/testutil/keeper"
	"github.com/defund-labs/defund/v1/x/query/keeper"
	"github.com/defund-labs/defund/v1/x/query/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.QueryKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
