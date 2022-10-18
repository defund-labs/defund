package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	channeltypes "github.com/cosmos/ibc-go/v4/modules/core/04-channel/types"

	etftypes "github.com/defund-labs/defund/x/etf/types"
)

// OnRedeemSuccess runs the redeem etf shares logic which takes escrowed etf shares and
// proportionally burns them.
func (k Keeper) OnRedeemSuccess(ctx sdk.Context, packet channeltypes.Packet, redeem etftypes.Redeem) error {
	// lets burn the etf shares in escrow
	err := k.bankKeeper.BurnCoins(ctx, "etf", sdk.NewCoins(*redeem.Amount))
	if err != nil {
		return err
	}

	// reset the proper amount of etf shares in fund by subtracted newly burnt shares
	fund, found := k.GetFund(ctx, redeem.Fund.Symbol)
	if !found {
		return sdkerrors.Wrapf(etftypes.ErrFundNotFound, "fund %s not found", redeem.Fund.Symbol)
	}
	fund.Shares.Amount = fund.Shares.Amount.Sub(redeem.Amount.Amount)
	k.SetFund(ctx, fund)

	// lets clear up the redeem from the store
	k.RemoveRedeem(ctx, redeem.Id)

	return nil
}

// OnRedeemFailure runs the redeem etf shares failure logic which takes escrowed etf shares
// and proportionally sends them back to the redeemer. This is used in Timeout as well
func (k Keeper) OnRedeemFailure(ctx sdk.Context, packet channeltypes.Packet, redeem etftypes.Redeem) error {
	// lets send the etf shares in escrow back to the redeemer
	redeemer, err := sdk.AccAddressFromBech32(redeem.Creator)
	if err != nil {
		return err
	}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, "etf", redeemer, sdk.NewCoins(*redeem.Amount))
	if err != nil {
		return err
	}

	// lets clear up the redeem from the store
	k.RemoveRedeem(ctx, redeem.Id)

	return nil
}

// OnRebalanceSuccess runs the rebalance etf logic which just deletes the rebalance
// in the store and updates the funds last rebalance height.
func (k Keeper) OnRebalanceSuccess(ctx sdk.Context, rebalance etftypes.Rebalance, fund *etftypes.Fund) error {
	fund.LastRebalanceHeight = int64(ctx.BlockHeight())
	k.SetFund(ctx, *fund)
	// Remove the rebalance from store. Clean up store
	k.RemoveRebalance(ctx, rebalance.Id)
	return nil
}

// OnRebalanceFailure runs the rebalance etf failure logic which just deletes the rebalance
// from store. Used for Timeout as well.
//
// NOTE: Potentially add a timeout/retry for failed rebalances?
func (k Keeper) OnRebalanceFailure(ctx sdk.Context, rebalance etftypes.Rebalance, fund *etftypes.Fund) error {
	// Remove the rebalance from store. Clean up store
	k.RemoveRebalance(ctx, rebalance.Id)
	return nil
}
