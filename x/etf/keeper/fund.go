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

// GetFund returns a fund from its symbol
func (k Keeper) GetFund(
	ctx sdk.Context,
	symbol string,

) (val types.Fund, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FundKeyPrefix))

	b := store.Get(types.FundKey(
		symbol,
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
