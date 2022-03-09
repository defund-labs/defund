package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/defundhub/defund/x/query/types"
)

/////////////////////////////// Interquery Store Helpers ///////////////////////////////////////////

// SetInterquery set a specific interquery in the store from its index
func (k Keeper) SetInterquery(ctx sdk.Context, interquery types.Interquery) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InterqueryKeyPrefix))
	b := k.cdc.MustMarshal(&interquery)
	store.Set(types.InterqueryKey(
		interquery.Storeid,
	), b)
}

// GetInterquery returns a interquery from its index
func (k Keeper) GetInterquery(
	ctx sdk.Context,
	index string,

) (val types.Interquery, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InterqueryKeyPrefix))

	b := store.Get(types.InterqueryKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveInterquery removes a interquery from the store
func (k Keeper) RemoveInterquery(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InterqueryKeyPrefix))
	store.Delete(types.InterqueryKey(
		index,
	))
}

// GetAllInterquery returns all interquery
func (k Keeper) GetAllInterquery(ctx sdk.Context) (list []types.Interquery) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InterqueryKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Interquery
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

/////////////////////////////// InterqueryResult Store Helpers ///////////////////////////////////////////

// SetInterquery set a specific interquery in the store from its index
func (k Keeper) SetInterqueryResult(ctx sdk.Context, interqueryresult types.InterqueryResult) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InterqueryResultKeyPrefix))
	b := k.cdc.MustMarshal(&interqueryresult)
	store.Set(types.InterqueryResultKey(
		interqueryresult.Storeid,
	), b)
}

// GetInterquery returns a interquery from its index
func (k Keeper) GetInterqueryResult(
	ctx sdk.Context,
	index string,

) (val types.InterqueryResult, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InterqueryResultKeyPrefix))

	b := store.Get(types.InterqueryResultKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveInterquery removes a interquery from the store
func (k Keeper) RemoveInterqueryResult(
	ctx sdk.Context,
	storeid string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InterqueryResultKeyPrefix))
	store.Delete(types.InterqueryResultKey(
		storeid,
	))
}

// GetAllInterquery returns all interquery
func (k Keeper) GetAllInterqueryResult(ctx sdk.Context) (list []types.InterqueryResult) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InterqueryResultKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.InterqueryResult
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

/////////////////////////////// InterqueryTimeoutResult Store Helpers ///////////////////////////////////////////

// SetInterquery set a specific interquery in the store from its index
func (k Keeper) SetInterqueryTimeoutResult(ctx sdk.Context, interquerytimeout types.InterqueryTimeoutResult) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InterqueryTimeoutResultKeyPrefix))
	b := k.cdc.MustMarshal(&interquerytimeout)
	store.Set(types.InterqueryTimeoutResultKey(
		interquerytimeout.Storeid,
	), b)
}

// GetInterquery returns a interquery from its index
func (k Keeper) GetInterqueryTimeoutResult(
	ctx sdk.Context,
	storeid string,

) (val types.InterqueryTimeoutResult, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InterqueryTimeoutResultKeyPrefix))

	b := store.Get(types.InterqueryTimeoutResultKey(
		storeid,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveInterquery removes a interquery from the store
func (k Keeper) RemoveInterqueryTimeoutResult(
	ctx sdk.Context,
	storeid string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InterqueryTimeoutResultKeyPrefix))
	store.Delete(types.InterqueryTimeoutResultKey(
		storeid,
	))
}

// GetAllInterquery returns all interquery
func (k Keeper) GetAllInterqueryTimeoutResult(ctx sdk.Context) (list []types.InterqueryTimeoutResult) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InterqueryTimeoutResultKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.InterqueryTimeoutResult
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

/////////////////////////////// Interquery Endblock Helpers ///////////////////////////////////////////

// Function to check if a string is within a slice
func contains(list []string, str string) bool {
	for _, value := range list {
		if value == str {
			return true
		}
	}

	return false
}

// Helper function to get all interqueries and then compare to
// submitted & timed out queries based on key and id (key-id should always be unique)
func (k Keeper) GetAllPendingInterqueriesFiltered(ctx sdk.Context) (list []types.Interquery) {
	// Get all interqueries from store
	interqueries := k.GetAllInterquery(ctx)
	// Get all interquery results from store
	interqueriesresults := k.GetAllInterqueryResult(ctx)
	// Get all interquery timeouts from store
	interqueriestimeouts := k.GetAllInterqueryTimeoutResult(ctx)

	// Set up blank list to store all the completed ids below
	completeids := []string{}
	// Set up blank list to store all the pending queries that need to be emitted
	pendingqueries := []types.Interquery{}

	// Add to list of complete ids for interquery results
	for _, queryresult := range interqueriesresults {
		if !contains(completeids, queryresult.Storeid) {
			completeids = append(completeids, queryresult.Storeid)
		}
	}
	// Add to list of timedout ids for interquery timeout results
	for _, querytimeout := range interqueriestimeouts {
		if !contains(completeids, querytimeout.Storeid) {
			completeids = append(completeids, querytimeout.Storeid)
		}
	}

	for _, query := range interqueries {
		if !contains(completeids, query.Storeid) {
			pendingqueries = append(pendingqueries, query)
		}
	}

	return pendingqueries
}

func (k Keeper) EmitInterqueryEvents(ctx sdk.Context) {
	//Get all interqueries that have not been submitted yet
	pendingqueries := k.GetAllPendingInterqueriesFiltered(ctx)

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
