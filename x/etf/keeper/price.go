package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
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

// CreateFundPrice creates a current fund price for a fund symbol
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
	// If the fund is brand new, the price starts at price specifed in BaseDenom (5,000,000 uosmo for example)
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
		amt := accBalance.Coins.AmountOf(holding.Token).Quo(fund.Shares.Amount).Mul(fundShares.Amount)
		amtCoin := sdk.NewCoin(holding.Token, amt)
		ownership = append(ownership, sdk.NewCoin(holding.Token, amtCoin.Amount))
	}

	return ownership, nil
}

// GetAmountETFSharesForTokens computes and returns the amount of etf shares a list of tokens would create for
// an etf. This function errors out if each token in tokens do not represent the same amount of etf shares
// in an ETF. Also errors out if all holdings in fund are not supplied.
func (k Keeper) GetAmountETFSharesForTokens(ctx sdk.Context, fund types.Fund, tokens []*sdk.Coin) (etfShares sdk.Coin, err error) {
	// get what one share of etf represents in ownership of underlying funds
	oneETFShareOwnershipRaw, err := k.GetOwnershipSharesInFund(ctx, fund, sdk.NewCoin(fund.BaseDenom, sdk.NewInt(1000000)))
	if err != nil {
		return etfShares, err
	}
	// turn list of coin into coins type
	oneETFShareOwnership := sdk.NewCoins(oneETFShareOwnershipRaw...)

	// compute the amount of etf tokens the first amount of tokens supplied represents in list.
	// this is next used to check if the rest of the tokens represent the same amount of ownership or it errors out
	etfSharesRaw := tokens[0].Amount.Quo(oneETFShareOwnership.AmountOf(tokens[0].Denom))
	etfShares = sdk.NewCoin(fund.Shares.Denom, etfSharesRaw)

	// check what the rest of tokens represent in etf
	for _, token := range tokens {
		etfSharesCheck := token.Amount.Quo(oneETFShareOwnership.AmountOf(token.Denom))
		if etfSharesCheck != etfSharesRaw {
			return etfShares, sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "all tokens do not represent the same ownership amount in this fund (denom: %s with amount: %d represents %d shares while it should represent %d shares)", token.Denom, token.Amount, etfSharesCheck, etfSharesRaw)
		}
	}

	return etfShares, nil
}
