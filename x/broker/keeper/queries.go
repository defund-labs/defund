package keeper

// All Osmosis Logic Lives Here

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// CreateDefundQueries creates all the repeated interqueries for broker chains
func (k Keeper) CreateDefundQueries(ctx sdk.Context) error {
	// Run every 10th block (1 minute)
	if ctx.BlockHeight()%10 == 0 {
		// Add Osmosis broker interquery for all pools
		k.CreateQueryOsmosisPools(ctx)
	}
	return nil
}
