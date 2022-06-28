package keeper

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/defund-labs/defund/x/etf/types"
)

// SetFund set a specific fund in the store from its index
func (k Keeper) SetFundPrice(ctx sdk.Context, fundprice types.FundPrice) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FundPriceKeyPrefix))
	b := k.cdc.MustMarshal(&fundprice)
	store.Set(types.FundPriceKey(
		fundprice.Id,
	), b)
}

// GetFund returns a fund from its index
func (k Keeper) GetFundPrice(
	ctx sdk.Context,
	index string,

) (val types.FundPrice, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FundPriceKeyPrefix))

	b := store.Get(types.FundPriceKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllFund returns all funds in store
func (k Keeper) GetAllFundPrice(ctx sdk.Context) (list []types.FundPrice) {
	store := ctx.KVStore(k.storeKey)
	fundPriceResultStore := prefix.NewStore(store, []byte(types.FundPriceKeyPrefix))

	iterator := fundPriceResultStore.Iterator(nil, nil)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.FundPrice
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// CreateFundPrice creates a current fund price for a fund symbol
func (k Keeper) CreateFundPrice(ctx sdk.Context, symbol string) (sdk.Coin, error) {
	fund, found := k.GetFund(ctx, symbol)
	return price, nil
}

// CreateAllFundPriceEndBlock is a function that runs at each end block that logs the fund price for each fund at current height
func (k Keeper) CreatePriceEndBlock(ctx sdk.Context) error {
	funds := k.GetAllFund(ctx)
	for _, fund := range funds {
		price, err := k.CreateFundPrice(ctx, fund.Symbol)
		if err != nil {
			return err
		}

		fundPrice := types.FundPrice{
			Height: uint64(ctx.BlockHeight()),
			Amount: &price,
			Symbol: fund.Symbol,
			Id:     fmt.Sprintf("%s-%s", fund.Symbol, strconv.FormatInt(ctx.BlockHeight(), 10)),
		}
		k.SetFundPrice(ctx, fundPrice)
	}
	return nil
}
