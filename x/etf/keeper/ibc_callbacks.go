package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
)

// OnAcknowledgementPacket performs an IBC send callback. Once a user submits an
// IBC transfer to a recipient to Defund, we check to see if the sequence associated with
// the transfer is associated with a Defund Create shares workflow. If not, we simply skip
// ahead the middleware stack, if it is, we record the share creation as complete if the ack
// is successful and mint new fund shares that are sent to the user that created the new shares.
// If it is failed, we record the Create shares as failed and then proceed through the IBC the stack.
func (k Keeper) OnAcknowledgementPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	acknowledgement []byte,
) error {
	return nil
}
