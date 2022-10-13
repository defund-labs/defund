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
	return types.Fund{}, sdkerrors.Wrapf(types.ErrFundNotFound, "fund with the symbol %s does not exist", symbol)
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

// GetRedeem returns a redeem if the redeem includes the transferId
func (k Keeper) GetRedeem(
	ctx sdk.Context,
	id string,

) (val types.Redeem, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RedeemKeyPrefix))

	b := store.Get(types.RedeemKey(
		id,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveRedeem removes an redeem from the store
func (k Keeper) RemoveRedeem(
	ctx sdk.Context,
	id string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RedeemKeyPrefix))
	store.Delete(types.RedeemKey(
		id,
	))
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

// RemoveRebalance removes an rebalance from the store
func (k Keeper) RemoveRebalance(
	ctx sdk.Context,
	id string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RebalanceKeyPrefix))
	store.Delete(types.RebalanceKey(
		id,
	))
}
