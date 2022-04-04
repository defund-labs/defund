package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/defund-labs/defund/v1/testutil/keeper"
	"github.com/defund-labs/defund/v1/x/query/keeper"
	"github.com/defund-labs/defund/v1/x/query/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNInterquery(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Interquery {
	items := make([]types.Interquery, n)
	for i := range items {
		items[i].Storeid = strconv.Itoa(i)

		keeper.SetInterquery(ctx, items[i])
	}
	return items
}

func TestInterqueryGet(t *testing.T) {
	keeper, ctx := keepertest.QueryKeeper(t)
	items := createNInterquery(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetInterquery(ctx,
			item.Storeid,
		)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}

func TestInterqueryGetAll(t *testing.T) {
	keeper, ctx := keepertest.QueryKeeper(t)
	items := createNInterquery(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllInterquery(ctx))
}
