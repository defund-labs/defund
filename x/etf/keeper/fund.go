package keeper

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/defund-labs/defund/x/etf/types"
)

// SetFund set a specific fund in the store from its index
func (k Keeper) SetFund(ctx sdk.Context, fund types.Fund) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FundKeyPrefix))
	b := k.cdc.MustMarshal(&fund)
	store.Set(types.FundKey(
		fund.Symbol,
	), b)
}

// GetFund returns a fund from its index
func (k Keeper) GetFund(
	ctx sdk.Context,
	index string,

) (val types.Fund, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FundKeyPrefix))

	b := store.Get(types.FundKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllFund returns all funds in store
func (k Keeper) GetAllFund(ctx sdk.Context) (list []types.Fund) {
	store := ctx.KVStore(k.storeKey)
	interqueryResultStore := prefix.NewStore(store, []byte(types.FundKeyPrefix))

	iterator := interqueryResultStore.Iterator(nil, nil)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Fund
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetFundBySymbol returns a fund by the funds symbol
func (k Keeper) GetFundBySymbol(ctx sdk.Context, symbol string) (types.Fund, error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FundKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Fund
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.Symbol == symbol {
			return val, nil
		}
	}
	return types.Fund{}, sdkerrors.Wrapf(types.ErrFundNotFound, "fund with the sumbol %s does not exist", symbol)
}

// GetNextID gets the count of all funds and then adds 1 for the next fund id
func (k Keeper) GetNextID(ctx sdk.Context) (id string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FundKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	count := 0

	for ; iterator.Valid(); iterator.Next() {
		count = count + 1
	}

	return strconv.Itoa(count)
}

// GetNextCreateID gets the count of all creates and then adds 1 for the next create id
func (k Keeper) GetNextCreateID(ctx sdk.Context) (id string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CreateKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	count := 0

	for ; iterator.Valid(); iterator.Next() {
		count = count + 1
	}

	return strconv.Itoa(count)
}

// SetCreate set a specific create in the store from its index
func (k Keeper) SetCreate(ctx sdk.Context, create types.Create) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CreateKeyPrefix))
	b := k.cdc.MustMarshal(&create)
	store.Set(types.CreateKey(
		create.Id,
	), b)
}

// GetCreate returns a increatevest from its index
func (k Keeper) GetCreate(
	ctx sdk.Context,
	index string,

) (val types.Create, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CreateKeyPrefix))

	b := store.Get(types.CreateKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllCreate returns all creates from store
func (k Keeper) GetAllCreate(ctx sdk.Context) (list []types.Create) {
	store := ctx.KVStore(k.storeKey)
	createStore := prefix.NewStore(store, []byte(types.CreateKeyPrefix))

	iterator := createStore.Iterator(nil, nil)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Create
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAllCreatebySymbol returns all creates from store based on symbol
func (k Keeper) GetAllCreatebySymbol(ctx sdk.Context, symbol string) (list []types.Create) {
	store := ctx.KVStore(k.storeKey)
	createStore := prefix.NewStore(store, []byte(types.CreateKeyPrefix))

	iterator := createStore.Iterator(nil, nil)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Create
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.Fund.Symbol == symbol {
			list = append(list, val)
		}
	}

	return
}

// GetNextRedeemID gets the count of all redeems and then adds 1 for the next redeem id
func (k Keeper) GetNextRedeemID(ctx sdk.Context) (id string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RedeemKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	count := 0

	for ; iterator.Valid(); iterator.Next() {
		count = count + 1
	}

	return strconv.Itoa(count)
}

// SetRedeem set a specific redeem in the store from its index
func (k Keeper) SetRedeem(ctx sdk.Context, redeem types.Redeem) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RedeemKeyPrefix))
	b := k.cdc.MustMarshal(&redeem)
	store.Set(types.RedeemKey(
		redeem.Id,
	), b)
}

// GetRedeem returns a redeem from its index
func (k Keeper) GetRedeem(
	ctx sdk.Context,
	index string,

) (val types.Redeem, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RedeemKeyPrefix))

	b := store.Get(types.RedeemKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllRedeem returns all redeems from store
func (k Keeper) GetAllRedeem(ctx sdk.Context) (list []types.Redeem) {
	store := ctx.KVStore(k.storeKey)
	redeemStore := prefix.NewStore(store, []byte(types.RedeemKeyPrefix))

	iterator := redeemStore.Iterator(nil, nil)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Redeem
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAllRedeembySymbol returns all redeems from store based on symbol
func (k Keeper) GetAllRedeembySymbol(ctx sdk.Context, symbol string) (list []types.Redeem) {
	store := ctx.KVStore(k.storeKey)
	redeemStore := prefix.NewStore(store, []byte(types.RedeemKeyPrefix))

	iterator := redeemStore.Iterator(nil, nil)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Redeem
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.Fund.Symbol == symbol {
			list = append(list, val)
		}
	}

	return
}

// GetRebalance returns a rebalance from its index
func (k Keeper) GetRebalance(
	ctx sdk.Context,
	index string,

) (val types.Rebalance, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RebalanceKeyPrefix))

	b := store.Get(types.RebalanceKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}
