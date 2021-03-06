package etf_test

import (
	"testing"

	keepertest "github.com/defund-labs/defund/testutil/keeper"
	"github.com/defund-labs/defund/x/etf"
	"github.com/defund-labs/defund/x/etf/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		FundList: []types.Fund{
			{
				Symbol: "0",
			},
			{
				Symbol: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.EtfKeeper(t)
	etf.InitGenesis(ctx, *k, genesisState)
	got := etf.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	require.Len(t, got.FundList, len(genesisState.FundList))
	require.Subset(t, genesisState.FundList, got.FundList)
	// this line is used by starport scaffolding # genesis/test/assert
}
