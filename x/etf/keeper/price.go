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
	invests := k.GetAllInvestbySymbol(ctx, symbol)
	if !found {
		return sdk.Coin{}, sdkerrors.Wrapf(types.ErrFundNotFound, "Could not find fund (%s)", symbol)
	}
	comp := []sdk.Dec{}
	for _, holding := range fund.Holdings {
		balances, err := k.GetHighestHeightPoolBalance(ctx, holding.PoolId)
		if err != nil {
			return sdk.Coin{}, err
		}
		if balances[0].Denom == holding.Token && fund.BaseDenom != holding.Token {
			baseAmount := balances[0].Amount.ToDec()
			tokenAmount := balances[1].Amount.ToDec()
			priceInBaseDenom := tokenAmount.Quo(baseAmount)
			percentDec := sdk.NewDec(holding.Percent).Quo(sdk.NewDec(100))
			comp = append(comp, priceInBaseDenom.Mul(percentDec))
		}
		if balances[1].Denom == holding.Token && fund.BaseDenom != holding.Token {
			baseAmount := balances[1].Amount.ToDec()
			tokenAmount := balances[0].Amount.ToDec()
			priceInBaseDenom := tokenAmount.Quo(baseAmount)
			percentDec := sdk.NewDec(holding.Percent).Quo(sdk.NewDec(100))
			comp = append(comp, priceInBaseDenom.Mul(percentDec))
		}
		// If the holding token is the baseDenom, just multiply it by the % it represents since we already know its price relative
		// to itself. Aka -> 1/1
		if fund.BaseDenom == holding.Token {
			percentDec := sdk.NewDec(holding.Percent).Quo(sdk.NewDec(100))
			comp = append(comp, sdk.NewDec(1).Mul(percentDec))
		}
		if len(comp) == 0 {
			return sdk.Coin{}, sdkerrors.Wrapf(types.ErrFundNotFound, "No price details found for symbol (%s)", symbol)
		}
	}

	price := sdk.Coin{}

	// If the fund is brand new, the price starts at 1,000,000 BaseDenom (1,000,000 uatom for example)
	if len(invests) == 0 {
		price = sdk.NewCoin(fund.BaseDenom, sdk.NewInt(1000000))
	}

	if len(invests) > 0 {
		total := sum(comp)
		price = sdk.NewCoin(fund.BaseDenom, sdk.NewInt(total.RoundInt64()))
	}

	return price, nil
}

// CreateAllFundPriceEndBlock is a function that runs at each end block that logs the fund price for each fund at current height
func (k Keeper) CreateAllFundPriceEndBlock(ctx sdk.Context) error {
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
