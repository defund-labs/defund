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

	b := store.Get(types.BrokerKey(
		id,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllBrokers returns all brokers in store
func (k Keeper) GetAllBrokers(ctx sdk.Context) (list []types.Broker) {
	store := ctx.KVStore(k.storeKey)
	brokerResultStore := prefix.NewStore(store, types.KeyPrefix(types.BrokerKeyPrefix))

	iterator := brokerResultStore.Iterator(nil, nil)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Broker
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetOsmosisPoolFromBroker gets an osmosis pool from a broker store
func (k Keeper) GetPoolFromBroker(ctx sdk.Context, brokerId string, poolId uint64) (val types.Pool, found bool) {
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
			return val, true
		}
	}

	return val, false
}
