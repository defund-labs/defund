package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	connectiontypes "github.com/cosmos/ibc-go/v4/modules/core/03-connection/types"
	channeltypes "github.com/cosmos/ibc-go/v4/modules/core/04-channel/types"

	osmosisgammtypes "github.com/osmosis-labs/osmosis/v8/x/gamm/types"
)

// Get a ICA account address on a host chain
func (k Keeper) GetBrokerAccount(ctx sdk.Context, ConnectionId string, portId string) (string, bool) {
	acc, found := k.icaControllerKeeper.GetInterchainAccountAddress(ctx, ConnectionId, portId)
	if !found {
		return "", false
	}
	return acc, true
}

// Registers an ICA account on a host chain
func (k Keeper) RegisterBrokerAccount(ctx sdk.Context, connectionID, owner string) error {
	if err := k.icaControllerKeeper.RegisterInterchainAccount(ctx, connectionID, owner, ""); err != nil {
		return err
	}
	return nil
}

// GetIBCConnection is a wrapper to get a connection from id
func (k Keeper) GetIBCConnection(ctx sdk.Context, connectionID string) (connectiontypes.ConnectionEnd, bool) {
	connection, found := k.connectionKeeper.GetConnection(ctx, connectionID)
	return connection, found
}

// GetChannel is a wrapper to get a channel from id
func (k Keeper) GetChannel(ctx sdk.Context, portID string, channelID string) (channeltypes.Channel, bool) {
	channel, found := k.channelKeeper.GetChannel(ctx, portID, channelID)
	return channel, found
}

// OnAcknowledgementPacketSuccess is the logic called on the IBC OnAcknowledgementPacket callback.
// In this function we check the incoming packet as an ICS-27 packet. We then take that ICS-27
// packet and run through each ICA message for the ack.
//
// If the ICA message is an ICA Send message then we know it is a Redeem message for redeeming ETF shares.
// We check to see the sequence corresponds with a redeem store, if it does, we then proceed to check
// if it was a successful msg. If so we then burn the fund shares held by the module account from the initial Redeem flow.
// If it failed, we take the escrowed etf shares and proportionally send them back to the redeemer.
//
// If the ICA message is an ICA Swap Message, we know it is a rebalance workflow, and we mark the rebalance
// from pending to complete.
func (k Keeper) OnAcknowledgementPacketSuccess(ctx sdk.Context, packet channeltypes.Packet, ack channeltypes.Acknowledgement, txMsgData *sdk.TxMsgData) {
	// loop through each ICA msg in the tx (one ack respresents one tx)
	for _, msgData := range txMsgData.Data {
		switch msgData.MsgType {
		case sdk.MsgTypeURL(&banktypes.MsgSend{}):
			// get the redeem from the store. If not found return nil and do not run logic
			redeem, found := k.brokerKeeper.GetRedeem(ctx, fmt.Sprintf("%s-%d", packet.SourceChannel, packet.Sequence))
			if !found {
				k.Logger(ctx).Error(fmt.Sprintf("Redeem %s not found. Skipping redeem success logic.", fmt.Sprintf("%s-%d", packet.SourceChannel, packet.Sequence)))
				return
			}
			k.Logger(ctx).Info("Redeem shares ICA transfer msg ran successfully. Running redeem success logic.")
			// Run redeem success logic
			err := k.OnRedeemSuccess(ctx, packet, redeem)
			if err != nil {
				k.Logger(ctx).Error("Error occured during run of ICA callback.", "callback", "OnRedeemSuccess", "error", err.Error())
			}

			return
		case sdk.MsgTypeURL(&osmosisgammtypes.MsgSwapExactAmountIn{}):
			// get the rebalance from the store. If not found return nil and do not run logic
			rebalance, found := k.brokerKeeper.GetRebalance(ctx, fmt.Sprintf("%s-%d", packet.SourceChannel, packet.Sequence))
			if !found {
				k.Logger(ctx).Error(fmt.Sprintf("Fund rebalance %s not found. Skipping rebalance success logic.", fmt.Sprintf("%s-%d", packet.SourceChannel, packet.Sequence)))
				return
			}
			k.Logger(ctx).Info("Fund rebalance ICA msg ran successfully. Running rebalance success logic.")
			// Run rebalance success logic
			err := k.OnRebalanceSuccess(ctx, rebalance, rebalance.Fund)
			if err != nil {
				k.Logger(ctx).Error("Error occured during run of ICA callback.", "callback", "OnRebalanceSuccess", "error", err.Error())
			}

			return
		default:
			k.Logger(ctx).Error("Received ICA failure msg type we do not recognize.", "typeUrl", msgData.MsgType)

			return
		}
	}
}

func (k Keeper) OnAcknowledgementPacketFailure(ctx sdk.Context, packet channeltypes.Packet, ack channeltypes.Acknowledgement, txMsgData *sdk.TxMsgData) {
	// get the redeem from the store. If not found skip redeem logic
	redeem, found := k.brokerKeeper.GetRedeem(ctx, fmt.Sprintf("%s-%d", packet.SourceChannel, packet.Sequence))
	if !found {
		k.Logger(ctx).Debug(fmt.Sprintf("Redeem %s not found. Skipping redeem failure logic.", fmt.Sprintf("%s-%d", packet.SourceChannel, packet.Sequence)))
	} else {
		k.Logger(ctx).Error("Redeem shares ICA transfer msg ran unsuccessfully. Running redeem failure logic.")
		// Run redeem failure logic
		err := k.OnRedeemFailure(ctx, packet, redeem)
		if err != nil {
			k.Logger(ctx).Error("Error occured during run of ICA callback.", "callback", "OnRedeemFailure", "error", err.Error())
		}
	}
	// get the rebalance from the store. If not found skip rebalance logic
	rebalance, found := k.brokerKeeper.GetRebalance(ctx, fmt.Sprintf("%s-%d", packet.SourceChannel, packet.Sequence))
	if !found {
		k.Logger(ctx).Debug(fmt.Sprintf("Fund rebalance %s not found. Skipping rebalance failure logic.", fmt.Sprintf("%s-%d", packet.SourceChannel, packet.Sequence)))
	} else {
		k.Logger(ctx).Error("Fund rebalance ICA msg ran unsuccessfully. Running rebalance failure logic.")
		// Run rebalance failure logic
		err := k.OnRebalanceFailure(ctx, rebalance, rebalance.Fund)
		if err != nil {
			k.Logger(ctx).Error("Error occured during run of ICA callback.", "callback", "OnRebalanceFailure", "error", err.Error())
		}
	}
}

func (k Keeper) OnAcknowledgementPacketICA(ctx sdk.Context, packet channeltypes.Packet, ack channeltypes.Acknowledgement, txMsgData *sdk.TxMsgData) {
	switch ack.Response.(type) {
	// on successful ack
	case *channeltypes.Acknowledgement_Result:
		ctx.Logger().Info("received successful ICA acknowledgement. running ICA successful acknowledgement logic.", "channel", packet.SourceChannel, "sequence", packet.Sequence, "ack", ack.GetResponse())
		k.OnAcknowledgementPacketSuccess(ctx, packet, ack, txMsgData)
	// on failure ack. defaults to failure in switch statement
	default:
		ctx.Logger().Info("received error ICA acknowledgement. running ICA failure acknowledgement logic.", "channel", packet.SourceChannel, "sequence", packet.Sequence, "ack", ack.GetError())
		k.OnAcknowledgementPacketFailure(ctx, packet, ack, txMsgData)
	}
}
