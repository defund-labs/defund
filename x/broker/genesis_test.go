package broker_test

import (
	"testing"

	keepertest "github.com/defund-labs/defund/testutil/keeper"
	"github.com/defund-labs/defund/x/broker"
	"github.com/defund-labs/defund/x/broker/types"
	"github.com/stretchr/testify/require"
	dbm "github.com/tendermint/tm-db"
)

func TestGenesis(t *testing.T) {
	genesisState := types.DefaultGenesis()

	db := dbm.NewMemDB()
	k, ctx := keepertest.BrokerKeeper(db, t)
	broker.InitGenesis(ctx, *k, *genesisState)
	got := broker.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	require.Equal(t, got.Brokers[0].Id, "osmosis")
	require.Equal(t, got.Params.AtomIBCPath, types.DefaultAtomPath)
	require.Equal(t, got.Params.OsmoIBCPath, types.DefaultOsmoPath)
	// this line is used by starport scaffolding # genesis/test/assert
}
