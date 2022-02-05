package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/defundhub/defund/x/query/types"
)

// SetInterquery set a specific interquery in the store from its index
func (k Keeper) SetInterquery(ctx sdk.Context, interquery types.Interquery) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InterqueryKeyPrefix))
	b := k.cdc.MustMarshal(&interquery)
	store.Set(types.InterqueryKey(
		interquery.Index,
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
