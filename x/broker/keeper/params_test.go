package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/defund-labs/defund/testutil/keeper"
	"github.com/defund-labs/defund/x/broker/types"
	dbm "github.com/tendermint/tm-db"
)

func TestGetParams(t *testing.T) {
	db := dbm.NewMemDB()
	k, ctx := testkeeper.BrokerKeeper(db, t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
