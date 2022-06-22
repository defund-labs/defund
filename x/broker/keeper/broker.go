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
	interqueryResultStore := prefix.NewStore(store, []byte(types.BrokerKeyPrefix))

	iterator := interqueryResultStore.Iterator(nil, nil)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Broker
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
