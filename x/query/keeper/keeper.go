package keeper

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/defund-labs/defund/x/query/types"

	json "github.com/tendermint/tendermint/libs/json"

	liquiditytypes "github.com/tendermint/liquidity/x/liquidity/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey

		accountKeeper types.AccountKeeper
		brokerKeeper  types.BrokerKeeper
	}
)

type PoolsKey struct {
}

type BalanceKey struct {
	Address string `json:"address"`
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,

	accountkeeper types.AccountKeeper,
	brokerkeeper types.BrokerKeeper,

) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,

		accountKeeper: accountkeeper,
		brokerKeeper:  brokerkeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) NewQueryAddress(id uint64) sdk.AccAddress {
	key := append([]byte("query"), sdk.Uint64ToBigEndian(id)...)
	return address.Module("query", key)
}

// Helper Function to check if a string is within a slice
func contains(list []string, str string) bool {
	for _, value := range list {
		if value == str {
			return true
		}
	}

	return false
}

func (k Keeper) CreateInterqueryRequest(ctx sdk.Context, storeid string, path string, key []byte, timeoutheight uint64, clientid string) error {
	var queryModuleAddress authtypes.ModuleAccountI
	if k.accountKeeper.GetModuleAccount(ctx, "query") == nil {
		queryAddress := k.NewQueryAddress(1)
		queryModuleAddress = authtypes.NewModuleAccount(
			authtypes.NewBaseAccountWithAddress(
				queryAddress,
			),
			"query",
		)
		k.accountKeeper.SetModuleAccount(ctx, queryModuleAddress)
	} else {
		queryModuleAddress = k.accountKeeper.GetModuleAccount(ctx, "query")
	}
	interquery := types.Interquery{
		Creator:       queryModuleAddress.GetAddress().String(),
		Storeid:       storeid,
		Path:          path,
		Key:           key,
		TimeoutHeight: timeoutheight,
		ClientId:      clientid,
	}
	k.SetInterquery(ctx, interquery)

	k.Logger(ctx).Info(fmt.Sprintf("Interquery request for path %s on clientid of %s has been initiated", path, clientid))

	return nil
}

// Helper function to be implemented in end blocker to interchain query pools on gravity dex (Cosmos Hub)
func (k Keeper) QueryGravityDex(ctx sdk.Context) error {
	path := "custom/liquidity/liquidityPools/"
	clientid := "07-tendermint-0"
	keyRaw := PoolsKey{}
	key, _ := json.Marshal(keyRaw)
	timeoutHeight := uint64(ctx.BlockHeight() + 10)
	storeid := "gdex-pools"

	err := k.CreateInterqueryRequest(ctx, storeid, path, key, timeoutHeight, clientid)
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

	err = k.CreateInterqueryRequest(ctx, storeid, path, key, timeoutHeight, clientid)
	if err != nil {
		return err
	}
	return nil
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

// GetHighestHeightPoolDetails gets the most recent (highest height) pool details of a specific pool in interqueryresult store
func (k Keeper) GetHighestHeightPoolDetails(ctx sdk.Context, poolid string) (liquiditytypes.Pool, error) {
	pools, err := k.GetHighestHeightPools(ctx)
	if err != nil {
		ctx.Logger().Debug(err.Error())
	}
	for _, pool := range pools {
		if strconv.FormatUint(pool.Id, 10) == poolid {
			return pool, nil
		}
	}
	return liquiditytypes.Pool{}, sdkerrors.Wrapf(types.ErrInvalidPool, "Pool not found (%s)", poolid)
}

// GetHighestHeightPoolBalance gets the most recent (highest height) balance/holdings of a pool in interqueryresult store
func (k Keeper) GetHighestHeightPoolBalance(ctx sdk.Context, poolid string) ([]sdk.Coin, error) {
	queries := k.GetAllInterqueryResult(ctx)
	poolQueries := []types.InterqueryResult{}
	balances := []sdk.Coin{}
	for _, query := range queries {
		idSplit := strings.Split(query.Storeid, "-")
		if idSplit[0] == "poolbalance" {
			if idSplit[1] == poolid {
				poolQueries = append(poolQueries, query)
			}
		}
	}
	// Sort tje poolQueries from largest to smallest
	sort.SliceStable(poolQueries, func(i, j int) bool {
		return poolQueries[i].Height > poolQueries[j].Height
	})
	// Take the query with the most recent height (first in sorted slice)
	if len(poolQueries) > 0 {
		rawquery := poolQueries[0]
		err := json.Unmarshal(rawquery.Data, &balances)
		if err != nil {
			return []sdk.Coin{}, sdkerrors.Wrapf(types.ErrMarshallingError, "Marshalling error for pool balances (PoolId: %s)", poolid)
		}
	}

	if len(poolQueries) == 0 {
		return []sdk.Coin{}, sdkerrors.Wrapf(types.ErrInvalidPools, "No pools interqueried. Need pools interqueried to proceed")
	}
	return balances, nil
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

func (k Keeper) TimeoutInterqueries(ctx sdk.Context) {
	// Get all interqueries from store
	interqueries := k.GetAllInterquery(ctx)

	// Loop through each query and timeout if not valid
	for _, query := range interqueries {
		if uint64(ctx.BlockHeight()) > query.TimeoutHeight {
			queryTimeout := types.InterqueryTimeoutResult{
				Creator:       query.Creator,
				Storeid:       query.Storeid,
				TimeoutHeight: query.TimeoutHeight,
				ClientId:      query.ClientId,
			}
			// Set the query as a timed out interquery in store
			k.SetInterqueryTimeoutResult(ctx, queryTimeout)
			// Remove/cleanup the interquery from pending interqueries
			k.RemoveInterquery(ctx, query.Storeid)
		}
	}
}

func (k Keeper) ModuleEndBlocker(ctx sdk.Context) {

	//Timeout all timedout/invalid interqueries at the beginning of the end block
	k.TimeoutInterqueries(ctx)

	//Get all interqueries that have not been submitted yet
	pendingqueries := k.GetAllInterquery(ctx)

	// Create holder for all events
	events := sdk.Events{}

	for _, query := range pendingqueries {

		event := sdk.NewEvent(
			types.EventTypeQuery,
			sdk.NewAttribute(types.AttributeKeyQueryClientId, query.ClientId),
			sdk.NewAttribute(types.AttributeKeyQueryPath, query.Path),
			sdk.NewAttribute(types.AttributeKeyQueryStoreid, query.Storeid),
		)

		events = append(events, event)

	}

	//Emit the query event
	ctx.EventManager().EmitEvents(events)
}
