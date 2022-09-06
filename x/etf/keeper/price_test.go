package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/defund-labs/defund/testutil/keeper"
	dbm "github.com/tendermint/tm-db"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestPrice(t *testing.T) {
	db := dbm.NewMemDB()
	keeper, ctx := keepertest.EtfKeeper(db, t)
	items := createNFund(keeper, ctx, 10)
	for _, item := range items {
		keeper.CreateFundPrice(ctx, item.Symbol)
	}
}
