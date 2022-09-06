package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/defund-labs/defund/testutil/keeper"
	"github.com/defund-labs/defund/x/etf/keeper"
	"github.com/defund-labs/defund/x/etf/types"
	"github.com/stretchr/testify/require"
	dbm "github.com/tendermint/tm-db"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNFund(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Fund {
	items := make([]types.Fund, n)
	for i := range items {
		items[i].Symbol = strconv.Itoa(i)

		keeper.SetFund(ctx, items[i])
	}
	return items
}

func TestFundGet(t *testing.T) {
	db := dbm.NewMemDB()
	keeper, ctx := keepertest.EtfKeeper(db, t)
	items := createNFund(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetFund(ctx,
			item.Symbol,
		)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}

func TestFundGetAll(t *testing.T) {
	db := dbm.NewMemDB()
	keeper, ctx := keepertest.EtfKeeper(db, t)
	items := createNFund(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllFund(ctx))
}
