package testutil

import (
	"defund/app"
	"testing"

	defundtypes "defund/types"

	"cosmossdk.io/log"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
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

	hdr := cmtproto.Header{
		Height: 1,
		Time:   defundtypes.ParseTime("2022-01-01T00:00:00Z"),
	}

	ctx := app.BaseApp.NewContext(true).WithBlockHeader(hdr)

	return app, ctx
}
