package keeper

import (
	"defund/app"
	"defund/x/dex/keeper"
	"defund/x/dex/types"
	"testing"

	"cosmossdk.io/log"
	dbm "github.com/cosmos/cosmos-db"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestDexKeeper(t testing.TB) (keeper.Keeper, sdk.Context, *app.App) {
	db := dbm.NewMemDB()
	app := app.New(
		log.NewNopLogger(),
		db,
		nil,
		true,
		simtestutil.EmptyAppOptions{},
	)

	k := app.DexKeeper

	ctx := app.BaseApp.NewContext(true)

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return k, ctx, app
}
