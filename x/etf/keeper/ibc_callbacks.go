package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	"github.com/defund-labs/defund/x/etf/types"
)

// OnTransferSuccess runs the transfer success logic which deletes the transfer from the store
func (k Keeper) OnTransferSuccess(ctx sdk.Context, packet channeltypes.Packet, ack channeltypes.Acknowledgement) error {
	transfer, found := k.brokerKeeper.GetTransfer(ctx, fmt.Sprintf("%s-%d", packet.SourceChannel, packet.Sequence))
	if !found {
		return sdkerrors.Wrapf(types.ErrCreateNotFound, "transfer %s not found", fmt.Sprintf("%s-%d", packet.SourceChannel, packet.Sequence))
	}
	ctx.Logger().Debug(fmt.Sprintf("transfer %s was completed successfully", transfer.Id))
	k.brokerKeeper.RemoveTransfer(ctx, transfer.Id)
	return nil
}

// OnTransferFailure runs the on transfer failure logic which logs the failure and nothing else.
// In the future may add a retry limit.
func (k Keeper) OnTransferFailure(ctx sdk.Context, packet channeltypes.Packet, ack channeltypes.Acknowledgement) error {
	transfer, found := k.brokerKeeper.GetTransfer(ctx, fmt.Sprintf("%s-%d", packet.SourceChannel, packet.Sequence))
	if !found {
		return sdkerrors.Wrapf(types.ErrCreateNotFound, "transfer %s not found", fmt.Sprintf("%s-%d", packet.SourceChannel, packet.Sequence))
	}
	ctx.Logger().Debug(fmt.Sprintf("transfer %s failed", transfer.Id))
	return nil
}

// OnTransferTimeout runs the on transfer timeout logic which logs the timeout and nothing else.
func (k Keeper) OnTransferTimeout(ctx sdk.Context, packet channeltypes.Packet) error {
	transfer, found := k.brokerKeeper.GetTransfer(ctx, fmt.Sprintf("%s-%d", packet.SourceChannel, packet.Sequence))
	if !found {
		return sdkerrors.Wrapf(types.ErrCreateNotFound, "transfer %s not found", fmt.Sprintf("%s-%d", packet.SourceChannel, packet.Sequence))
	}
	ctx.Logger().Debug(fmt.Sprintf("transfer %s timed out", transfer.Id))
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
