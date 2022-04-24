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

// SetInvest set a specific invest in the store from its index
func (k Keeper) SetInvest(ctx sdk.Context, invest types.Invest) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InvestKeyPrefix))
	b := k.cdc.MustMarshal(&invest)
	store.Set(types.InvestKey(
		invest.Id,
	), b)
}

// GetInvest returns a invest from its index
func (k Keeper) GetInvest(
	ctx sdk.Context,
	index string,

) (val types.Invest, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InvestKeyPrefix))

	b := store.Get(types.InvestKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllInvest returns all invests from store
func (k Keeper) GetAllInvest(ctx sdk.Context) (list []types.Invest) {
	store := ctx.KVStore(k.storeKey)
	investStore := prefix.NewStore(store, []byte(types.InvestKeyPrefix))

	iterator := investStore.Iterator(nil, nil)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Invest
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAllInvestbySymbol returns all invests from store based on symbol
func (k Keeper) GetAllInvestbySymbol(ctx sdk.Context, symbol string) (list []types.Invest) {
	store := ctx.KVStore(k.storeKey)
	investStore := prefix.NewStore(store, []byte(types.InvestKeyPrefix))

	iterator := investStore.Iterator(nil, nil)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Invest
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.Fund.Symbol == symbol {
			list = append(list, val)
		}
	}

	return
}

// GetInvestBySequence returns an invest store based on its sequence and channel
func (k Keeper) GetInvestBySequence(ctx sdk.Context, sequence uint64, channel string) (types.Invest, error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InvestKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Invest
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.Channel == channel && val.Sequence == sequence {
			return val, nil
		}
	}
	return types.Invest{}, sdkerrors.Wrapf(types.ErrInvestNotFound, "invest not found for sequence %s on channel %s", sequence, channel)
}

// GetNextIDInvest gets the count of all invest and then adds 1 for the next invest id
func (k Keeper) GetNextIDInvest(ctx sdk.Context) (id string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InvestKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	count := 0

	for ; iterator.Valid(); iterator.Next() {
		count = count + 1
	}

	return strconv.Itoa(count)
}

// SetUninvest set a specific uninvest in the store from its index
func (k Keeper) SetUninvest(ctx sdk.Context, uninvest types.Uninvest) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InvestKeyPrefix))
	b := k.cdc.MustMarshal(&uninvest)
	store.Set(types.InvestKey(
		uninvest.Id,
	), b)
}

// GetUninvest returns a invest from its index
func (k Keeper) GetUninvest(
	ctx sdk.Context,
	index string,

) (val types.Uninvest, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UninvestKeyPrefix))

	b := store.Get(types.UninvestKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllUninvest returns all invests from store
func (k Keeper) GetAllUninvest(ctx sdk.Context) (list []types.Uninvest) {
	store := ctx.KVStore(k.storeKey)
	uninvestStore := prefix.NewStore(store, []byte(types.UninvestKeyPrefix))

	iterator := uninvestStore.Iterator(nil, nil)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Uninvest
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAllUninvestbySymbol returns all uninvests from store based on symbol
func (k Keeper) GetAllUninvestbySymbol(ctx sdk.Context, symbol string) (list []types.Uninvest) {
	store := ctx.KVStore(k.storeKey)
	uninvestStore := prefix.NewStore(store, []byte(types.UninvestKeyPrefix))

	iterator := uninvestStore.Iterator(nil, nil)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Uninvest
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.Fund.Symbol == symbol {
			list = append(list, val)
		}
	}

	return
}

// GetUninvestBySequence returns an Uninvest store based on its sequence and channel
func (k Keeper) GetUninvestBySequence(ctx sdk.Context, sequence uint64, channel string) (types.Uninvest, error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UninvestKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Uninvest
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.Channel == channel && val.Sequence == sequence {
			return val, nil
		}
	}
	return types.Uninvest{}, sdkerrors.Wrapf(types.ErrUninvestNotFound, "Uninvest not found for sequence %s on channel %s", sequence, channel)
}

// GetNextIDUninvest gets the count of all uninvests and then adds 1 for the next uninvest id
func (k Keeper) GetNextIDUninvest(ctx sdk.Context) (id string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UninvestKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	count := 0

	for ; iterator.Valid(); iterator.Next() {
		count = count + 1
	}

	return strconv.Itoa(count)
}
