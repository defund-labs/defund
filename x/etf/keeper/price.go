package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

// CreateFundPrice creates a current fund price for a fund symbol in the funds base denom
func (k Keeper) CreateFundPrice(ctx sdk.Context, symbol string) (price sdk.Coin, err error) {
	comp := []sdk.Dec{}
	fund, found := k.GetFund(ctx, symbol)
	if !found {
		return price, sdkerrors.Wrapf(types.ErrFundNotFound, "fund %s not found", symbol)
	}
	for _, holding := range fund.Holdings {
		// check that a pool with the broker exists
		_, found := k.brokerKeeper.GetPoolFromBroker(ctx, holding.BrokerId, holding.PoolId)
		if !found {
			return price, sdkerrors.Wrapf(types.ErrInvalidPool, "pool %d not found on broker %s", holding.PoolId, holding.BrokerId)
		}
		priceUnweighted, err := k.brokerKeeper.CalculateOsmosisSpotPrice(ctx, holding.PoolId, holding.Token, fund.BaseDenom)
		if err != nil {
			return price, err
		}
		priceWeighted := priceUnweighted.Mul(sdk.NewDec(holding.Percent / 100))
		comp = append(comp, priceWeighted)
	}
	// If the fund is brand new, the price starts at price specifed in BaseDenom (5,000,000 uusdc for example)
	if fund.Shares.Amount.Uint64() == 0 {
		price = fund.StartingPrice
	}
	if fund.Shares.Amount.Uint64() > 0 {
		total := sum(comp)
		price = sdk.NewCoin(fund.BaseDenom, sdk.NewInt(total.RoundInt64()))
	}
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
		accBalance, err := k.brokerKeeper.GetOsmosisBalance(ctx, accAddress)
		if err != nil {
			return ownership, err
		}

		// take holding and find per etf share of holding from fund balance then multiply it by
		// the amount of fundShares
		amt := accBalance.Coins.Sort().AmountOf(holding.Token).Quo(fund.Shares.Amount).Mul(fundShares.Amount)
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
	amt := tokenIn.Amount.Quo(fundPrice.Amount)

	// Create the coin fund shares price
	etfShares = sdk.NewCoin(fund.Shares.Denom, amt)

	return etfShares, nil
}
