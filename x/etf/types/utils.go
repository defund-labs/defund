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
	PriceInBase    sdk.Dec
	PriceInHolding sdk.Dec
	CurrentComp    *sdk.Dec
}

type PricedHoldings []PricedHolding

func (p PricedHoldings) GetPercentComposition(denom string) (price PricedHolding, err error) {
	total := sdk.NewDec(0)
	var found bool
	for _, priced := range p {
		total = total.Add(priced.PriceInBase)
		if priced.Holding.Token == denom {
			// set the price to the current priced since its what we are looking for
			price = priced
			// set to found
			found = true
		}
	}
	comp := price.PriceInBase.Quo(total)
	price.CurrentComp = &comp
	// if we never found the denom specified error out
	if !found {
		return price, sdkerrors.Wrapf(ErrInvalidHolding, "denom %s not found in priced holdings list", denom)
	}

	return price, nil
}

func (p PricedHoldings) GetAmountOf(denom string, basedenom bool) (sdk.Dec, error) {
	for _, priced := range p {
		if priced.Holding.Token == denom {
			if basedenom {
				return priced.PriceInBase, nil
			} else {
				return priced.PriceInHolding, nil
			}
		}
	}

	return sdk.Dec{}, sdkerrors.Wrapf(ErrInvalidHolding, "denom %s not found in priced holdings list", denom)
}

func (b *FundBalances) GetBalancesByAddress(address string) sdk.Coins {
	var coins []sdk.Coin
	switch address {
	case b.GetOsmosis().Address:
		for i := range b.Osmosis.Balances {
			coins = append(coins, *b.Osmosis.Balances[i])
		}
		return sdk.NewCoins(coins...)
	default:
		return sdk.NewCoins()
	}
}

func (b *FundBalances) GetBalancesByBroker(brokerid string) (bals sdk.Coins, found bool) {
	var coins []sdk.Coin
	switch brokerid {
	case "osmosis":
		for i := range b.Osmosis.Balances {
			coins = append(coins, *b.Osmosis.Balances[i])
		}
		return sdk.NewCoins(coins...), true
	default:
		return sdk.Coins{}, false
	}
}

// SetBalance sets a new balance for the fund + broker specified
func (fund *Fund) SetBalances(brokerid string, address string, coin sdk.Coin) error {
	switch brokerid {
	case "osmosis":
		var updated bool = false
		fund.Balances.Osmosis.Address = address
		for i := range fund.Balances.Osmosis.Balances {
			// if we already have this denom update it and break the loop
			if fund.Balances.Osmosis.Balances[i].Denom == coin.Denom {
				fund.Balances.Osmosis.Balances[i] = &coin
				updated = true
				break
			}
		}
		if !updated {
			// if we did not update by end of last loop, we do not have denom so just set as new denom
			fund.Balances.Osmosis.Balances = append(fund.Balances.Osmosis.Balances, &coin)
		}
		return nil
	default:
		return sdkerrors.Wrapf(ErrInvalidBalance, "the fund %s does not have a fund balance for the broker %s", fund.Symbol, brokerid)
	}
}

// HasBalance checks if the fund supplied has any balances. Returns boolean
func (b *FundBalances) HasBalance(symbol string) (hasbalance bool) {
	/////// Check Osmosis //////////
	if len(b.Osmosis.Balances) > 0 {
		return true
	} else {
		return false
	}
}
