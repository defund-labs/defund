package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	channeltypes "github.com/cosmos/ibc-go/v4/modules/core/04-channel/types"
	"github.com/defund-labs/defund/x/etf/types"
)

// OnTransferSuccess runs the transfer success logic which creates shares for the fund and sends them to creator.
// The transfer is then deleted from the store.
func (k Keeper) OnTransferSuccess(ctx sdk.Context, packet channeltypes.Packet, ack channeltypes.Acknowledgement) error {
	transfer, found := k.brokerKeeper.GetTransfer(ctx, fmt.Sprintf("%s-%d", packet.SourceChannel, packet.Sequence))
	// if not found continue with logic as normal as its not a create shares transfer
	if !found {
		return nil
	}
	fund, found := k.GetFund(ctx, transfer.Fund)
	if !found {
		return sdkerrors.Wrapf(types.ErrFundNotFound, "OnTransferSuccess: fund %s not found", transfer.Fund)
	}

	// otherwise the base denom in has been successfully sent via ibc so create new etf shares and send to creator of shares
	// compute the amount of etf shares this creator is given
	numETFShares, err := k.GetAmountETFSharesForToken(ctx, fund, *transfer.Token)
	if err != nil {
		return err
	}
	newETFCoins := sdk.NewCoins(numETFShares)

	// finally mint coins (to module account) and then send them to the creator of the create
	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, newETFCoins)
	if err != nil {
		return err
	}
	creatorAcc, err := sdk.AccAddressFromBech32(transfer.Sender)
	if err != nil {
		return err
	}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creatorAcc, newETFCoins)
	if err != nil {
		return err
	}

	// finally reflect the new shares in the fund store for shares
	newShares := fund.Shares.Add(numETFShares)
	fund.Shares = &newShares
	k.SetFund(ctx, fund)

	// log success and remove transfer from store
	ctx.Logger().Debug(fmt.Sprintf("create shares transfer %s was completed successfully", transfer.Id))
	k.brokerKeeper.RemoveTransfer(ctx, transfer.Id)
	return nil
}

// OnTransferFailure runs the on transfer failure logic which logs the failure. No need to release escrow tokenIn since this
// is handled on Transfer module at the end of Transfer middleware stack automatically.
func (k Keeper) OnTransferFailure(ctx sdk.Context, packet channeltypes.Packet, ack channeltypes.Acknowledgement) error {
	transfer, found := k.brokerKeeper.GetTransfer(ctx, fmt.Sprintf("%s-%d", packet.SourceChannel, packet.Sequence))
	// if not found continue with logic
	if !found {
		return nil
	}
	ctx.Logger().Debug(fmt.Sprintf("transfer %s failed. sending escrowed assets back to user", transfer.Id))
	k.brokerKeeper.RemoveTransfer(ctx, transfer.Id)
	return nil
}

// OnTransferTimeout runs the on transfer timeout logic which logs the timeout. No need to release escrow tokenIn since this
// is handled on Transfer module at the end of Transfer middleware stack automatically.
func (k Keeper) OnTransferTimeout(ctx sdk.Context, packet channeltypes.Packet) error {
	transfer, found := k.brokerKeeper.GetTransfer(ctx, fmt.Sprintf("%s-%d", packet.SourceChannel, packet.Sequence))
	// if not found continue with logic
	if !found {
		return nil
	}
	ctx.Logger().Debug(fmt.Sprintf("transfer %s timed out. sending escrowed assets back to user", transfer.Id))
	k.brokerKeeper.RemoveTransfer(ctx, transfer.Id)
	return nil
}

// OnAcknowledgementPacket performs an IBC ack callback. Once a user submits an
// IBC transfer to a recipient to Defund, we check to see if the sequence associated with
// the transfer is associated with a Defund Transfer workflow. If not, we simply skip
// ahead the middleware stack, if it is, we record the share creation as complete if the ack
// is successful and mint new fund shares that are sent to the user that created the new shares.
// If it is failed, we record the Create shares as failed and then proceed through the IBC middleware stack.
func (k Keeper) OnAcknowledgementPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	acknowledgement []byte,
) error {
	// unmarshal the ack to be used later
	var ack channeltypes.Acknowledgement
	if err := channeltypes.SubModuleCdc.UnmarshalJSON(acknowledgement, &ack); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal Broker packet acknowledgement: %v", err)
	}
	switch ack.Response.(type) {
	// on successful ack
	case *channeltypes.Acknowledgement_Result:
		return k.OnTransferSuccess(ctx, packet, ack)
	// on failure ack
	case *channeltypes.Acknowledgement_Error:
		return k.OnTransferFailure(ctx, packet, ack)
	default:
		return nil
	}
}

// OnTimeoutPacket implements the transfer middleware for handling Create logic on timeout which
// removes the create from store. No need to handle escrow as the base application does all of that for us
func (k Keeper) OnTimeoutPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) error {
	err := k.OnTransferTimeout(ctx, packet)
	if err != nil {
		return err
	}
	return nil
}
