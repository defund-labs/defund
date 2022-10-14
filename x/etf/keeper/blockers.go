package keeper

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	connectiontypes "github.com/cosmos/ibc-go/v4/modules/core/03-connection/types"
	channeltypes "github.com/cosmos/ibc-go/v4/modules/core/04-channel/types"
)

// SendPendingTransfers takes all pending transfers from the store
// and sends the IBC transfers for each transfer. These transfers represent
// the unsuccessful transfers from creates and redeems. If an error occurs we just log and continue to next
// iteration as we do not want to stop all transfers for one transfer error.
func (k Keeper) SendPendingTransfers(ctx sdk.Context) {
	transfers := k.brokerKeeper.GetAllTransfer(ctx)
	for _, transfer := range transfers {
		// get client and then get current height of the counterparty chain
		channel, found := k.channelKeeper.GetChannel(ctx, "transfer", transfer.Channel)
		if !found {
			err := sdkerrors.Wrapf(channeltypes.ErrChannelNotFound, "channel %s not found", transfer.Channel)
			ctx.Logger().Debug(err.Error())
			continue
		}
		connectionEnd, found := k.connectionKeeper.GetConnection(ctx, channel.ConnectionHops[0])
		if !found {
			err := sdkerrors.Wrap(connectiontypes.ErrConnectionNotFound, channel.ConnectionHops[0])
			ctx.Logger().Debug(err.Error())
			continue
		}
		clientState, found := k.clientKeeper.GetClientState(ctx, connectionEnd.GetClientID())
		if !found {
			err := sdkerrors.Wrapf(clienttypes.ErrConsensusStateNotFound, "consensus state for %s not found", connectionEnd.GetClientID())
			ctx.Logger().Debug(err.Error())
			continue
		}
		// create timeout info for transfer packet
		timeoutHeight := clientState.GetLatestHeight().GetRevisionHeight() + 50
		timeoutTimestamp := uint64(time.Now().Add(time.Minute).UnixNano())

		k.SendTransfer(ctx, transfer.Channel, *transfer.Token, transfer.Sender, transfer.Receiver, clienttypes.NewHeight(clientState.GetLatestHeight().GetRevisionNumber(), timeoutHeight), timeoutTimestamp)
	}
}

// SendRebalancesEndBlocker is the end blocker function that sends rebalance ICA's to all broker
// chains for each fund. If there is an error we just log it and continue
func (k Keeper) SendRebalancesEndBlocker(ctx sdk.Context) {
	funds := k.GetAllFund(ctx)

	for _, fund := range funds {
		err := k.SendRebalanceTx(ctx, fund)
		if err != nil {
			ctx.Logger().Error(fmt.Sprintf("rebalance failed for fund %s with error: %s", fund.Symbol, err.Error()))
		}
	}
}
