package keeper

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

// GetAllFundPrice returns all funds prices in store
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
func (k Keeper) CreateFundPrice(ctx sdk.Context, symbol string) (price sdk.Coin, err error) {
	comp := []sdk.Dec{}
	fund, found := k.GetFund(ctx, symbol)
	if !found {
		return price, sdkerrors.Wrapf(types.ErrFundNotFound, "fund %s not found", symbol)
	}
	for _, holding := range fund.Holdings {
		_, found := k.brokerKeeper.GetPoolFromBroker(ctx, fund.Broker.Id, holding.PoolId)
		if !found {
			return price, sdkerrors.Wrapf(types.ErrInvalidPool, "pool %s not found on broker %s", holding.PoolId, fund.Broker.Id)
		}
		priceUnweighted, err := k.brokerKeeper.CalculateOsmosisSpotPrice(ctx, holding.PoolId, holding.Token, fund.Broker.BaseDenom)
		if err != nil {
			return price, err
		}
		priceWeighted := priceUnweighted.Mul(sdk.NewDec(holding.Percent))
		comp = append(comp, priceWeighted)
	}
	// If the fund is brand new, the price starts at price specifed BaseDenom (5,000,000 uatom for example)
	if len(comp) == 0 {
		price = sdk.NewCoin(fund.Broker.BaseDenom, sdk.NewInt(5000000))
	}
	if len(comp) > 0 {
		total := sum(comp)
		price = sdk.NewCoin(fund.Broker.BaseDenom, sdk.NewInt(total.RoundInt64()))
	}
	return price, nil
}

// CreateAllFundPriceEndBlock is a function that runs at each end block that logs the fund price for each fund
// and purges unneeded fund prices from the store
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
			Time:   ctx.BlockTime(),
			Symbol: fund.Symbol,
			Id:     fmt.Sprintf("%s-%s", fund.Symbol, strconv.FormatInt(ctx.BlockHeight(), 10)),
		}
		k.SetFundPrice(ctx, fundPrice)
	}
	return nil
}
