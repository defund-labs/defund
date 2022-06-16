package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/defund-labs/defund/x/query/types"
)

/////////////////////////////// Interquery Store Helpers ///////////////////////////////////////////

func (k Keeper) MarshalInterquery(interquery types.Interquery) ([]byte, error) {
	return interquery.Marshal()
}

func (k Keeper) UnmarshalInterquery(bz []byte) (types.Interquery, error) {
	var query types.Interquery
	return query, k.cdc.UnmarshalInterface(bz, &query)
}

func (k Keeper) MarshalInterqueryResult(interquery types.InterqueryResult) ([]byte, error) {
	return interquery.Marshal()
}

func (k Keeper) UnmarshalInterqueryResult(bz []byte) (types.InterqueryResult, error) {
	var query types.InterqueryResult
	return query, k.cdc.UnmarshalInterface(bz, &query)
}

func (k Keeper) MarshalInterqueryTimeoutResult(interquery types.InterqueryTimeoutResult) ([]byte, error) {
	return interquery.Marshal()
}

func (k Keeper) UnmarshalInterqueryTimeoutResult(bz []byte) (types.InterqueryTimeoutResult, error) {
	var query types.InterqueryTimeoutResult
	return query, k.cdc.UnmarshalInterface(bz, &query)
}

// SetInterquery set a specific interquery in the store from its index
func (k Keeper) SetInterquery(ctx sdk.Context, interquery types.Interquery) error {
	bz, err := k.MarshalInterquery(interquery)
	if err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	interqueryKey := types.GetKeyPrefixInterquery(interquery.GetStoreid())
	store.Set(interqueryKey, bz)

	return nil
}

// GetInterquery returns a interquery from its index
func (k Keeper) GetInterquery(
	ctx sdk.Context,
	index string,

) (val types.Interquery, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.InterqueryKeyPrefix)

	b := store.Get(types.InterqueryKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllInterquery returns all interquery
func (k Keeper) GetAllInterquery(ctx sdk.Context) (list []types.Interquery) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.InterqueryKeyPrefix)
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
func (k Keeper) SetInterqueryResult(ctx sdk.Context, interquery types.InterqueryResult) error {
	bz, err := k.MarshalInterqueryResult(interquery)
	if err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	interqueryKey := types.GetKeyPrefixInterqueryResult(interquery.GetStoreid())
	store.Set(interqueryKey, bz)

	return nil
}

// GetInterquery returns a interquery from its index
func (k Keeper) GetInterqueryResult(
	ctx sdk.Context,
	index string,

) (val types.InterqueryResult, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.InterqueryResultKeyPrefix)

	b := store.Get(types.InterqueryResultKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllInterqueryResult returns all interquery results from the store
func (k Keeper) GetAllInterqueryResult(ctx sdk.Context) (list []types.InterqueryResult) {
	store := ctx.KVStore(k.storeKey)
	interqueryResultStore := prefix.NewStore(store, types.InterqueryResultKeyPrefix)

	iterator := interqueryResultStore.Iterator(nil, nil)

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
func (k Keeper) SetInterqueryTimeoutResult(ctx sdk.Context, interquery types.InterqueryTimeoutResult) error {
	bz, err := k.MarshalInterqueryTimeoutResult(interquery)
	if err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	interqueryKey := types.GetKeyPrefixInterqueryTimeoutResult(interquery.GetStoreid())
	store.Set(interqueryKey, bz)

	return nil
}

// GetInterquery returns a interquery from its index
func (k Keeper) GetInterqueryTimeoutResult(
	ctx sdk.Context,
	storeid string,

) (val types.InterqueryTimeoutResult, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.InterqueryTimeoutResultKeyPrefix)

	b := store.Get(types.InterqueryTimeoutResultKey(
		storeid,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllInterquery returns all interquery
func (k Keeper) GetAllInterqueryTimeoutResult(ctx sdk.Context) (list []types.InterqueryTimeoutResult) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.InterqueryTimeoutResultKeyPrefix)
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

//
func (k Keeper) TimeoutInterqueries(ctx sdk.Context) (list []types.Interquery) {
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

func (k Keeper) ModuleEndBlocker(ctx sdk.Context) {
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
