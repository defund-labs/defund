package keeper

// All Custom Transfer Logic Lives Here

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
)

// Creates an ICA Transfer msg on a host/broker ICA chain to
// send funds from that chain back to Defund address
func (k Keeper) IBCTransfer(ctx sdk.Context) {

}

// Sends an IBC transfer to a broker chain
func (k Keeper) SendTransfer(ctx sdk.Context, owner string, channel string, token sdk.Coin, sender string, receiver string, timeoutHeight clienttypes.Height, timeoutTimestamp uint64) error {
	portID, err := icatypes.NewControllerPortID(owner)
	if err != nil {
		return err
	}

	senderAddr, err := sdk.AccAddressFromBech32(sender)
	if err != nil {
		return err
	}

	err = k.transferKeeper.SendTransfer(ctx, portID, channel, token, senderAddr, receiver, timeoutHeight, uint64(timeoutTimestamp))
	if err != nil {
		return err
	}

	return nil
}
