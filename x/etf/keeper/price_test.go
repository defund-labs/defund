package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/defund-labs/defund/testutil/keeper"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestPrice(t *testing.T) {
	keeper, ctx := keepertest.EtfKeeper(t)
	items := createNFund(keeper, ctx, 10)
	for _, item := range items {
		keeper.CreateFundPrice(ctx, item.Symbol)
	}
}
