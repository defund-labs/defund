package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	proto "github.com/gogo/protobuf/proto"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	connectiontypes "github.com/cosmos/ibc-go/v4/modules/core/03-connection/types"
	channeltypes "github.com/cosmos/ibc-go/v4/modules/core/04-channel/types"

	transfertypes "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
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
func (k Keeper) OnAcknowledgementPacketSuccess(ctx sdk.Context, packet channeltypes.Packet, ack channeltypes.Acknowledgement, txMsgData *sdk.TxMsgData) error {
	// loop through each ICA msg in the tx (one ack respresents one tx)
	for _, msgData := range txMsgData.Data {
		switch msgData.MsgType {
		case sdk.MsgTypeURL(&banktypes.MsgSend{}):
			// get the redeem from the store. If not found return nil and do not run logic
			redeem, found := k.brokerKeeper.GetRedeem(ctx, fmt.Sprintf("%s-%d", packet.SourceChannel, packet.Sequence))
			if !found {
				k.Logger(ctx).Error(fmt.Sprintf("Redeem %s not found. Skipping redeem success logic.", fmt.Sprintf("%s-%d", packet.SourceChannel, packet.Sequence)))
				return nil
			}
			msgResponse := &transfertypes.MsgTransferResponse{}
			if err := proto.Unmarshal(msgData.Data, msgResponse); err != nil {
				err = sdkerrors.Wrapf(sdkerrors.ErrJSONUnmarshal, "cannot unmarshal ica transfer response message: %s", err.Error())
				k.Logger(ctx).Error("Failed running redeem success logic during msgResponse unmarshalling. ===>>> Error: %s", err.Error())
			}
			k.Logger(ctx).Info("Redeem shares ICA transfer msg ran successfully. Running redeem success logic.", "response", msgResponse.String())
			// Run redeem success logic
			k.OnRedeemSuccess(ctx, packet, redeem)

			return nil
		case sdk.MsgTypeURL(&osmosisgammtypes.MsgSwapExactAmountIn{}):
			// get the rebalance from the store. If not found return nil and do not run logic
			rebalance, found := k.brokerKeeper.GetRebalance(ctx, fmt.Sprintf("%s-%d", packet.SourceChannel, packet.Sequence))
			if !found {
				k.Logger(ctx).Error(fmt.Sprintf("Fund rebalance %s not found. Skipping rebalance success logic.", fmt.Sprintf("%s-%d", packet.SourceChannel, packet.Sequence)))
				return nil
			}
			msgResponse := &osmosisgammtypes.MsgSwapExactAmountInResponse{}
			if err := proto.Unmarshal(msgData.Data, msgResponse); err != nil {
				err = sdkerrors.Wrapf(sdkerrors.ErrJSONUnmarshal, "cannot unmarshal Osmosis swap in response message: %s", err.Error())
				k.Logger(ctx).Error("Failed running rebalance success logic during msgResponse unmarshalling. ===>>> Error: %s", err.Error())
			}
			k.Logger(ctx).Info("Fund rebalance ICA msg ran successfully. Running rebalance success logic.", "response", msgResponse.String())
			// Run rebalance success logic
			k.OnRebalanceSuccess(ctx, rebalance, rebalance.Fund)

			return nil
		default:
			return nil
		}
	}
	return nil
}

func (k Keeper) OnAcknowledgementPacketFailure(ctx sdk.Context, packet channeltypes.Packet, ack channeltypes.Acknowledgement, txMsgData *sdk.TxMsgData) error {
	// loop through each ICA msg in the tx (one ack respresents one tx)
	for _, msgData := range txMsgData.Data {
		switch msgData.MsgType {
		case sdk.MsgTypeURL(&banktypes.MsgSend{}):
			// get the redeem from the store. If not found return nil and do not run logic
			redeem, found := k.brokerKeeper.GetRedeem(ctx, fmt.Sprintf("%s-%d", packet.SourceChannel, packet.Sequence))
			if !found {
				k.Logger(ctx).Error(fmt.Sprintf("Redeem %s not found. Skipping redeem failure logic.", fmt.Sprintf("%s-%d", packet.SourceChannel, packet.Sequence)))
				return nil
			}
			msgResponse := &transfertypes.MsgTransferResponse{}
			if err := proto.Unmarshal(msgData.Data, msgResponse); err != nil {
				err = sdkerrors.Wrapf(sdkerrors.ErrJSONUnmarshal, "cannot unmarshal ica transfer response message: %s", err.Error())
				k.Logger(ctx).Error("Failed running redeem failure logic during msgResponse unmarshalling. ===>>> Error: %s", err.Error())
			}
			k.Logger(ctx).Error("Redeem shares ICA transfer msg ran unsuccessfully. Running redeem failure logic.", "response: ", msgResponse.String())

			// Run redeem failure logic
			k.OnRedeemFailure(ctx, packet, redeem)

			return nil
		case sdk.MsgTypeURL(&osmosisgammtypes.MsgSwapExactAmountIn{}):
			// get the rebalance from the store. If not found return nil and do not run logic
			rebalance, found := k.brokerKeeper.GetRebalance(ctx, fmt.Sprintf("%s-%d", packet.SourceChannel, packet.Sequence))
			if !found {
				k.Logger(ctx).Error(fmt.Sprintf("Fund rebalance %s not found. Skipping rebalance failure logic.", fmt.Sprintf("%s-%d", packet.SourceChannel, packet.Sequence)))
				return nil
			}
			fund := rebalance.Fund
			switch rebalance.Broker {
			case "osmosis":
				msgResponse := &osmosisgammtypes.MsgSwapExactAmountInResponse{}
				if err := proto.Unmarshal(msgData.Data, msgResponse); err != nil {
					err = sdkerrors.Wrapf(sdkerrors.ErrJSONUnmarshal, "cannot unmarshal Osmosis swap in response message: %s", err.Error())
					k.Logger(ctx).Error("Failed running rebalance failure logic during msgResponse unmarshalling. ===>>> Error: %s", err.Error())
				}
				k.Logger(ctx).Error("Fund rebalance ICA msg ran unsuccessfully. Running rebalance failure logic.", "response: ", msgResponse.String())
			}

			// Run rebalance failure logic
			k.OnRebalanceFailure(ctx, rebalance, fund)

			return nil
		default:
			return nil
		}
	}
	return nil
}

func (k Keeper) OnAcknowledgementPacketICA(ctx sdk.Context, packet channeltypes.Packet, ack channeltypes.Acknowledgement, txMsgData *sdk.TxMsgData) error {
	switch ack.Response.(type) {
	// on successful ack
	case *channeltypes.Acknowledgement_Result:
		return k.OnAcknowledgementPacketSuccess(ctx, packet, ack, txMsgData)
	// on failure ack
	case *channeltypes.Acknowledgement_Error:
		return k.OnAcknowledgementPacketFailure(ctx, packet, ack, txMsgData)
	default:
		return nil
	}
}
