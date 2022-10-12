package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	osmosisgammtypes "github.com/osmosis-labs/osmosis/v8/x/gamm/types"
)

type RebalanceMsgs struct {
	Osmosis []*osmosisgammtypes.MsgSwapExactAmountIn
}

type PricedHolding struct {
	Holding        Holding
	PriceInBase    sdk.Coin
	PriceInHolding sdk.Coin
	CurrentComp    *sdk.Dec
}

type PricedHoldings []PricedHolding

func (p PricedHoldings) GetPercentComposition(denom string) (price PricedHolding, err error) {
	total := sdk.NewDec(0)
	var found bool
	for _, priced := range p {
		total = total.Add(priced.PriceInBase.Amount.ToDec())
		if priced.Holding.Token == denom {
			// set the price to the current priced since its what we are looking for
			price = priced
			// set to found
			found = true
		}
	}
	comp := price.PriceInBase.Amount.ToDec().Quo(total)
	price.CurrentComp = &comp
	// if we never found the denom specified error out
	if !found {
		return price, sdkerrors.Wrapf(ErrInvalidHolding, "denom %s not found in priced holdings list", denom)
	}

	return price, nil
}

func (p PricedHoldings) GetAmountOf(denom string, basedenom bool) (sdk.Coin, error) {
	for _, priced := range p {
		if priced.Holding.Token == denom {
			if basedenom {
				return priced.PriceInBase, nil
			} else {
				return priced.PriceInHolding, nil
			}
		}
	}

	return sdk.Coin{}, sdkerrors.Wrapf(ErrInvalidHolding, "denom %s not found in priced holdings list", denom)
}
