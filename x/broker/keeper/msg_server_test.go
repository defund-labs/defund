package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/defund-labs/defund/testutil/keeper"
	"github.com/defund-labs/defund/x/etf/keeper"
	"github.com/defund-labs/defund/x/etf/types"
	dbm "github.com/tendermint/tm-db"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	db := dbm.NewMemDB()
	k, ctx := keepertest.EtfKeeper(db, t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
