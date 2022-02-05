package query

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/defundhub/defund/x/query/keeper"
	"github.com/defundhub/defund/x/query/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the interquery
	for _, elem := range genState.InterqueryList {
		k.SetInterquery(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.InterqueryList = k.GetAllInterquery(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
