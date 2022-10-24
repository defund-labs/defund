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

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.InterqueryKeyPrefix)
	interqueryKey := types.GetKeyPrefixInterquery(interquery.GetStoreid())
	store.Set(interqueryKey, bz)

	return nil
}

// GetInterquery returns a interquery from its storeid
func (k Keeper) GetInterquery(
	ctx sdk.Context,
	storeid string,

) (val types.Interquery, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.InterqueryKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Interquery
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.Storeid == storeid {
			return val, true
		}
	}

	return val, false
}

// RemoveInterquery removes an interquery from the store
func (k Keeper) RemoveInterquery(
	ctx sdk.Context,
	storeid string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.InterqueryKeyPrefix)
	interqueryKey := types.GetKeyPrefixInterquery(storeid)

	store.Delete(interqueryKey)
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
	storeid string,

) (val types.InterqueryResult, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(string(types.InterqueryResultKeyPrefix)))

	iterator := store.Iterator(nil, nil)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.InterqueryResult
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.Storeid == storeid {
			return val, true
		}
	}

	return val, false
}

// RemoveInterqueryResult removes an interquery result from the store
func (k Keeper) RemoveInterqueryResult(ctx sdk.Context, storeid string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.InterqueryResultKeyPrefix)
	store.Delete(types.InterqueryResultKey(
		storeid,
	))
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

// RemoveInterqueryTimeoutResult removes an interquery timeout result from the store
func (k Keeper) RemoveInterqueryTimeoutResult(
	ctx sdk.Context,
	storeid string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.InterqueryTimeoutResultKeyPrefix)
	store.Delete(types.InterqueryTimeoutResultKey(
		storeid,
	))
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
