package query_test

import (
	"testing"

	keepertest "github.com/defund-labs/defund/testutil/keeper"
	"github.com/defund-labs/defund/x/query"
	"github.com/defund-labs/defund/x/query/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		InterqueryList: []types.Interquery{
			{
				Storeid: "Key-0",
			},
			{
				Storeid: "Key-1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.QueryKeeper(t)
	query.InitGenesis(ctx, *k, genesisState)
	got := query.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	require.Len(t, got.InterqueryList, len(genesisState.InterqueryList))
	require.Subset(t, genesisState.InterqueryList, got.InterqueryList)
	// this line is used by starport scaffolding # genesis/test/assert
}
