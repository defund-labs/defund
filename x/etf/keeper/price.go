package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/defund-labs/defund/x/etf/types"
)

type PricesSort []*types.FundPrice

func (p PricesSort) Len() int {
	return len(p)
}

func (p PricesSort) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p PricesSort) Less(i, j int) bool {
	return p[i].Height < p[j].Height
}

// SetFundPrices set a specific funds list of prices in the store from its index
func (k Keeper) SetFundPrices(ctx sdk.Context, fundprices types.FundPrices) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FundPriceKeyPrefix))
	b := k.cdc.MustMarshal(&fundprices)
	store.Set(types.FundPriceKey(
		fundprices.Id,
	), b)
}

// GetFundPrices returns a funds prices from its index
func (k Keeper) GetFundPrices(
	ctx sdk.Context,
	index string,

) (val types.FundPrices, found bool) {
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

// GetAllFundPrices returns all funds prices in store
func (k Keeper) GetAllFundPrices(ctx sdk.Context) (list []types.FundPrices) {
	store := ctx.KVStore(k.storeKey)
	fundPriceResultStore := prefix.NewStore(store, []byte(types.FundPriceKeyPrefix))

	iterator := fundPriceResultStore.Iterator(nil, nil)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.FundPrices
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
		price = fund.StartingPrice
	}
	if len(comp) > 0 {
		total := sum(comp)
		price = sdk.NewCoin(fund.Broker.BaseDenom, sdk.NewInt(total.RoundInt64()))
	}
	return price, nil
}
