package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	icatypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/types"
	brokertypes "github.com/defund-labs/defund/x/broker/types"
	"github.com/defund-labs/defund/x/etf/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (k Keeper) GetBalanceForFundByAddress(ctx sdk.Context, symbol string, address string) (banktypes.Balance, error) {
	var coins []sdk.Coin

	fund, found := k.GetFund(ctx, symbol)
	if !found {
		return banktypes.Balance{}, sdkerrors.Wrapf(types.ErrFundNotFound, "fund %s not found", symbol)
	}

	for i := range fund.Balances[address].Balances {
		coins = append(coins, *fund.Balances[address].Balances[i])
	}

	balance := banktypes.Balance{
		Address: address,
		Coins:   sdk.NewCoins(coins...),
	}

	return balance, nil
}

// CreateFundPrice creates a current fund price for a fund symbol in the funds base denom
func (k Keeper) CreateFundPrice(ctx sdk.Context, symbol string) (price sdk.Coin, err error) {
	comp := []sdk.Dec{}
	fund, found := k.GetFund(ctx, symbol)
	if !found {
		return price, sdkerrors.Wrapf(types.ErrFundNotFound, "fund %s not found", symbol)
	}
	// If the fund is brand new, the price starts at price specifed in BaseDenom (5,000,000 uusdc for example)
	if fund.Shares.Amount.IsZero() {
		price = *fund.StartingPrice
		return price, nil
	}
	for _, holding := range fund.Holdings {
		broker, found := k.brokerKeeper.GetBroker(ctx, holding.BrokerId)
		if !found {
			return price, sdkerrors.Wrap(types.ErrWrongBroker, fmt.Sprintf("broker %s not found", holding.BrokerId))
		}
		// get the ica account for the fund on the broker chain
		portID, err := icatypes.NewControllerPortID(fund.Address)
		if err != nil {
			return price, err
		}
		fundBrokerAddress, found := k.brokerKeeper.GetBrokerAccount(ctx, broker.ConnectionId, portID)
		if !found {
			return price, sdkerrors.Wrapf(brokertypes.ErrIBCAccountNotExist, "failed to find ica account for owner %s on connection %s and port %s", fund.Address, broker.ConnectionId, portID)
		}
		// check that a pool with the broker exists
		_, found = k.brokerKeeper.GetPoolFromBroker(ctx, holding.BrokerId, holding.PoolId)
		if !found {
			return price, sdkerrors.Wrapf(types.ErrInvalidPool, "pool %d not found on broker %s", holding.PoolId, holding.BrokerId)
		}
		var balances banktypes.Balance
		var priceInBaseDenom sdk.Dec
		switch holding.BrokerId {
		case "osmosis":
			// get the account balances for the fund account on the broker chain
			balances, err = k.GetBalanceForFundByAddress(ctx, fund.Symbol, fundBrokerAddress)
			if err != nil {
				return price, err
			}
			// Calculate spot price for 1 holding token in base denom
			priceInBaseDenom, err = k.brokerKeeper.CalculateOsmosisSpotPrice(ctx, holding.PoolId, fund.BaseDenom, holding.Token)
			if err != nil {
				return price, err
			}
		}
		// get the holding denom amount from balances
		holdingBalance := balances.Coins.Sort().AmountOf(holding.Token).ToDec()
		// compute the weighted price by taking the holding balance of token, multiplying by price in base denom
		// to obtain the balance in base denom.
		priceWeighted := holdingBalance.Mul(priceInBaseDenom)
		comp = append(comp, priceWeighted)
	}
	total := sumDecs(comp).Mul(sdk.NewDec(1000000)).RoundInt()
	price = sdk.NewCoin(fund.BaseDenom, total.Quo(fund.Shares.Amount))
	return price, nil
}

// GetOwnershipSharesInFund computes and returns the tokens the fund coins provided
// represent/owns within the fund.
func (k Keeper) GetOwnershipSharesInFund(ctx sdk.Context, fund types.Fund, fundShares sdk.Coin) (ownership []sdk.Coin, err error) {
	// check to make sure the shares provided are for the correct fund
	if fund.Shares.Denom != fundShares.Denom {
		return ownership, sdkerrors.Wrapf(types.ErrInvalidDenom, "invalid etf denom. looking for %s, received %s", fund.Shares.Denom, fundShares.Denom)
	}
	for _, holding := range fund.Holdings {
		// get the broker
		broker, found := k.brokerKeeper.GetBroker(ctx, holding.BrokerId)
		if !found {
			return ownership, sdkerrors.Wrap(brokertypes.ErrBrokerNotFound, fmt.Sprintf("broker %s not found for holding %s", holding.BrokerId, holding.Token))
		}

		// get the ica account address port
		portID, err := icatypes.NewControllerPortID(fund.Address)
		if err != nil {
			return ownership, err
		}
		// get the ica account address
		accAddress, found := k.icaControllerKeeper.GetInterchainAccountAddress(ctx, broker.ConnectionId, portID)
		if !found {
			return ownership, status.Errorf(codes.NotFound, "no account found for portID %s", portID)
		}
		// get the ica accounts token balances
		accBalance, err := k.GetBalanceForFundByAddress(ctx, fund.Symbol, accAddress)
		if err != nil {
			return ownership, err
		}

		// take holding and find per etf share of holding from fund balance then multiply it by
		// the amount of fundShares
		amt := accBalance.Coins.Sort().AmountOf(holding.Token).ToDec().Quo(fund.Shares.Amount.ToDec()).Mul(fundShares.Amount.ToDec()).RoundInt()
		amtCoin := sdk.NewCoin(holding.Token, amt)
		ownership = append(ownership, sdk.NewCoin(holding.Token, amtCoin.Amount))
	}

	return ownership, nil
}

// GetAmountETFSharesForTokens computes and returns the amount of etf shares the tokenIn will create.
// The base denom must be used for the tokenIn or it will error.
func (k Keeper) GetAmountETFSharesForToken(ctx sdk.Context, fund types.Fund, tokenIn sdk.Coin) (etfShares sdk.Coin, err error) {
	// Make sure the tokenIn is the correct base denom for the fund
	if fund.BaseDenom != tokenIn.Denom {
		return etfShares, sdkerrors.Wrapf(types.ErrWrongBaseDenom, "the base denom for the fund is %s not %s", fund.BaseDenom, tokenIn.Denom)
	}

	fundPrice, err := k.CreateFundPrice(ctx, fund.Symbol)
	if err != nil {
		return etfShares, err
	}

	// Divide the amount of tokens provided by the price of the fund
	amt := tokenIn.Amount.ToDec().Quo(fundPrice.Amount.ToDec()).Mul(sdk.NewDec(1000000)).RoundInt()

	// Create the coin fund shares price
	etfShares = sdk.NewCoin(fund.Shares.Denom, amt)

	return etfShares, nil
}
