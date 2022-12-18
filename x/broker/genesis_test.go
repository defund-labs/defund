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
	genesisState := types.GenesisState{
		Brokers: []types.Broker{},
		Params:  types.DefaultParams(),
		// this line is used by starport scaffolding # genesis/test/state
	}

	db := dbm.NewMemDB()
	k, ctx := keepertest.BrokerKeeper(db, t)
	broker.InitGenesis(ctx, *k, genesisState)
	got := broker.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	require.Len(t, got.Brokers, len(genesisState.Brokers))
	require.Equal(t, got.Params.AtomIBCPath, types.DefaultAtomPath)
	require.Equal(t, got.Params.OsmoIBCPath, types.DefaultOsmoPath)
	t.Log(genesisState.Brokers)
	t.Log(got.Brokers)
	require.Subset(t, genesisState.Brokers, got.Brokers)
	// this line is used by starport scaffolding # genesis/test/assert
}
