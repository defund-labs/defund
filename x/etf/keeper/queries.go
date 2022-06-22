package keeper

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/defund-labs/defund/x/etf/types"
)

type PoolsKey struct {
}

type BalanceKey struct {
	Address string `json:"address"`
}

// Helper function to be implemented in end blocker to interchain query pools on gravity dex (Cosmos Hub)
func (k Keeper) QueryGravityDex(ctx sdk.Context) error {
	path := "custom/liquidity/liquidityPools/"
	clientid := "07-tendermint-0"
	keyRaw := PoolsKey{}
	key, _ := json.Marshal(keyRaw)
	timeoutHeight := uint64(ctx.BlockHeight() + 10)
	storeid := "gdex-pools"

	err := k.queryKeeper.CreateInterqueryRequest(ctx, storeid, path, key, timeoutHeight, clientid)
	if err != nil {
		return err
	}
	return nil
}

// Helper function that creates an interquery for an account balance on Cosmos with the accountType as part of the store id
func (k Keeper) QueryPoolAccount(ctx sdk.Context, pool uint64, address string) error {
	path := "custom/bank/all_balances/"
	clientid := "07-tendermint-0"
	keyRaw := BalanceKey{address}
	key, err := json.Marshal(keyRaw)
	if err != nil {
		return err
	}
	timeoutHeight := uint64(ctx.BlockHeight() + 10)
	storeid := fmt.Sprintf("poolbalance-%d", pool)

	err = k.queryKeeper.CreateInterqueryRequest(ctx, storeid, path, key, timeoutHeight, clientid)
	if err != nil {
		return err
	}
	return nil
}

// GetHighestHeightPools gets the most recent (highest height) of all pools in interqueryresult store
func (k Keeper) GetHighestHeightPools(ctx sdk.Context) ([]liquiditytypes.Pool, error) {
	queries := k.GetAllInterqueryResult(ctx)
	poolQueries := []types.InterqueryResult{}
	pools := []liquiditytypes.Pool{}
	for _, query := range queries {
		idSplit := strings.Split(query.Storeid, "-")
		if idSplit[0] == "gdex" && idSplit[1] == "pools" {
			poolQueries = append(poolQueries, query)
		}
	}
	// Sort tje poolQueries from largest to smallest
	sort.SliceStable(poolQueries, func(i, j int) bool {
		return poolQueries[i].Height > poolQueries[j].Height
	})
	// Take the query with the most recent height (first in sorted slice)
	if len(poolQueries) > 0 {
		query := poolQueries[0]
		json.Unmarshal(query.Data, &pools)
	}

	if len(poolQueries) == 0 {
		return nil, sdkerrors.Wrapf(types.ErrInvalidPools, "No pools interqueried. Need pools interqueried to proceed")
	}
	return pools, nil
}

// QueryAllPools queries all pool accounts from gdex from most recent pools store
func (k Keeper) QueryAllPools(ctx sdk.Context) error {
	recentPools, err := k.GetHighestHeightPools(ctx)
	// Log error if error returns from query. Do not want to panic application. Just log
	if err != nil {
		ctx.Logger().Debug(err.Error())
	}
	// Take the query with the most recent height (first in sorted slice)
	if len(recentPools) > 0 {
		for _, pool := range recentPools {
			err := k.QueryPoolAccount(ctx, pool.Id, pool.ReserveAccountAddress)
			if err != nil {
				return err
			}
		}
	} else {
		ctx.Logger().Debug("no pools in store to interquery")
	}

	return nil
}

// CreateDefundQueries creates all the repeated interqueries for defund
func (k Keeper) CreateDefundQueries(ctx sdk.Context) error {
	// Run every 10th block
	if ctx.BlockHeight()%10 == 0 {
		// Add gravity dex interquery
		err := k.QueryGravityDex(ctx)
		if err != nil {
			ctx.Logger().Debug(fmt.Sprintf("Error Creating Cosmos GDex Pool Interquery: %s", err))
		}
		// Add interquery for all pool account balances on gdex
		err = k.QueryAllPools(ctx)
		if err != nil {
			ctx.Logger().Debug(fmt.Sprintf("Error Creating Interquery For All Cosmos GDex Pool Accounts: %s", err))
		}
	}
	return nil
}
