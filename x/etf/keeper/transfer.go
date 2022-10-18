package keeper

// All Custom Transfer Logic Lives Here

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	"github.com/defund-labs/defund/x/broker/types"
)

// Sends an IBC transfer to another chain
func (k Keeper) SendTransfer(ctx sdk.Context, channel string, token sdk.Coin, sender string, receiver string, timeoutHeight clienttypes.Height, timeoutTimestamp uint64) (sequence uint64, err error) {
	portID := "transfer"

	senderAddr, err := sdk.AccAddressFromBech32(sender)
	if err != nil {
		return sequence, err
	}

	sequence, found := k.channelKeeper.GetNextSequenceSend(ctx, portID, channel)
	if !found {
		return sequence, sdkerrors.Wrapf(types.ErrNextSequenceNotFound, "failed to retrieve the next sequence for channel %s and port %s", channel, portID)
	}

	err = k.transferKeeper.SendTransfer(ctx, portID, channel, token, senderAddr, receiver, timeoutHeight, uint64(timeoutTimestamp))
	if err != nil {
		return sequence, err
	}

	return sequence, nil
}
