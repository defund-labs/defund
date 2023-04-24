package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/defund-labs/defund/x/broker/types"
)

// SetBroker sets a specific broker in the store from its index
func (k Keeper) SetBroker(ctx sdk.Context, broker types.Broker) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BrokerKeyPrefix))
	b := k.cdc.MustMarshal(&broker)
	store.Set(types.BrokerKey(
		broker.Id,
	), b)
}

// GetBroker returns a broker from its index by id
func (k Keeper) GetBroker(
	ctx sdk.Context,
	id string,

) (val types.Broker, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BrokerKeyPrefix))

	iterator := store.Iterator(nil, nil)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Broker
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.Id == id {
			return val, true
		}
	}

	return val, false
}

// GetAllBrokers returns all brokers in store
func (k Keeper) GetAllBrokers(ctx sdk.Context) (list []types.Broker) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BrokerKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Broker
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetOsmosisPoolFromBroker gets an osmosis pool from a broker store
func (k Keeper) GetPoolFromBroker(ctx sdk.Context, brokerId string, poolId uint64) (val types.Source, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BrokerKeyPrefix))

	var broker types.Broker

	b := store.Get(types.BrokerKey(
		brokerId,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &broker)

	for _, pool := range broker.Pools {
		if pool.PoolId == poolId {
			val = *pool
			return val, true
		}
	}

	return val, false
}

// SetCreate set a specific create in the store from its index
func (k Keeper) SetCreate(ctx sdk.Context, transfer types.Create) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TransferKeyPrefix))
	b := k.cdc.MustMarshal(&transfer)
	store.Set(types.TransferKey(
		transfer.Id,
	), b)
}

// GetCreate returns a create from its index
func (k Keeper) GetCreate(
	ctx sdk.Context,
	index string,

) (val types.Create, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TransferKeyPrefix))

	b := store.Get(types.TransferKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveCreate removes an create from the store
func (k Keeper) RemoveCreate(
	ctx sdk.Context,
	id string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TransferKeyPrefix))
	store.Delete(types.TransferKey(
		id,
	))
}

// Create returns all creates from store
func (k Keeper) GetAllCreate(ctx sdk.Context) (list []types.Create) {
	store := ctx.KVStore(k.storeKey)
	transferStore := prefix.NewStore(store, []byte(types.TransferKeyPrefix))

	iterator := transferStore.Iterator(nil, nil)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Create
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAllCreateWithDep returns all creates from store based on the query it depends on
func (k Keeper) GetAllCreateWithDep(ctx sdk.Context, queryId string, status string) (list []types.Create) {
	store := ctx.KVStore(k.storeKey)
	transferStore := prefix.NewStore(store, []byte(types.TransferKeyPrefix))

	iterator := transferStore.Iterator(nil, nil)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Create
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.Query == queryId && val.Status == status {
			list = append(list, val)
		}
	}

	return
}
