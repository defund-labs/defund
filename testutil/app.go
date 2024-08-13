package testutil

import (
	"defund/app"
	"testing"

	"cosmossdk.io/log"
	dbm "github.com/cosmos/cosmos-db"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestApp(t testing.TB) (*app.App, sdk.Context) {
	db := dbm.NewMemDB()
	app := app.New(
		log.NewNopLogger(),
		db,
		nil,
		true,
		simtestutil.EmptyAppOptions{},
	)

	ctx := app.BaseApp.NewContext(true)

	return app, ctx
}
