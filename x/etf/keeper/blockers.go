package keeper

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	connectiontypes "github.com/cosmos/ibc-go/v4/modules/core/03-connection/types"
	channeltypes "github.com/cosmos/ibc-go/v4/modules/core/04-channel/types"
	brokertypes "github.com/defund-labs/defund/x/broker/types"
	"github.com/defund-labs/defund/x/etf/types"
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

// CheckRedeemsAndFinishIfDone checks the redeem for all transfers and if all are complete
// it finishes the redeem process by burning the escrowed shares and removing the redeem from store
func (k Keeper) CheckRedeemsAndFinishIfDone(ctx sdk.Context) {
	// get all redeems to loop through
	redeems := k.GetAllRedeem(ctx)
	for _, redeem := range redeems {
		// create holder to keep track of counts of completes
		countDone := 0
		countError := 0
		for _, transfer := range redeem.Transfers {
			if transfer.Status == brokertypes.StatusComplete {
				// if the transfer is complete add to the complete count
				countDone = countDone + 1
			}
			if transfer.Status == brokertypes.StatusError {
				// if any transfer errored must add to count to reverse later
				countError = countError + 1
			}
		}
		// if the complete count is equal to the length of transfer we are done
		if countDone == len(redeem.Transfers) {
			// Burn escrowed shares in the etf module
			err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(*redeem.Amount))
			if err != nil {
				ctx.Logger().Error(fmt.Sprintf("error occured during redeem check. could not burn escrowed shares: %s", err.Error()))
				break
			}
			fundAddress, err := sdk.AccAddressFromBech32(redeem.Fund.Address)
			if err != nil {
				ctx.Logger().Error(fmt.Sprintf("error occured during redeem check. could not convert to address: %s", err.Error()))
				break
			}
			redeemerAddress, err := sdk.AccAddressFromBech32(redeem.Creator)
			if err != nil {
				ctx.Logger().Error(fmt.Sprintf("error occured during redeem check. could not convert to address: %s", err.Error()))
				break
			}
			// send the shares to the defund account for the redeemer
			err = k.bankKeeper.SendCoins(ctx, fundAddress, redeemerAddress, sdk.NewCoins(*redeem.Amount))
			if err != nil {
				ctx.Logger().Error(fmt.Sprintf("error occured during redeem check. could not send transfers: %s", err.Error()))
				break
			}
			// finally reflect removal of shares in the fund store for shares
			redeem.Fund.Shares = redeem.Fund.Shares.Sub(*redeem.Amount)
			k.SetFund(ctx, *redeem.Fund)
			// Remove the redeem from store. Clean up store
			k.RemoveRedeem(ctx, redeem.Id)
		}
		// if the count error has any errors then we need to revert
		if countError > 0 {
			// Send escrowed shares back to user
			acc, err := sdk.AccAddressFromBech32(redeem.Creator)
			if err != nil {
				ctx.Logger().Error(fmt.Sprintf("error occured during redeem check. could not convert to address: %s", err.Error()))
				break
			}
			err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, acc, sdk.NewCoins(*redeem.Amount))
			if err != nil {
				ctx.Logger().Error(fmt.Sprintf("error occured during redeem check. could not send coins from module: %s", err.Error()))
				break
			}
			// loop through each transfer and send back to broker chain
			for _, transfer := range redeem.Transfers {
				// get the transfer channel
				channel, found := k.channelKeeper.GetChannel(ctx, "transfer", transfer.Channel)
				if !found {
					err := sdkerrors.Wrapf(channeltypes.ErrChannelNotFound, "port: transfer, channel: %s", transfer.Channel)
					ctx.Logger().Error(err.Error())
					continue
				}
				_, err = k.SendTransfer(ctx, channel.String(), *transfer.Token, transfer.Receiver, transfer.Sender, clienttypes.NewHeight(0, 0), 0)
				if err != nil {
					ctx.Logger().Error(fmt.Sprintf("error occured during redeem check. could not send transfers: %s", err.Error()))
					break
				}
			}
			// Remove the redeem from store. Clean up store
			k.RemoveRedeem(ctx, redeem.Id)
		}
	}
}
