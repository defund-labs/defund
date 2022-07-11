package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	proto "github.com/gogo/protobuf/proto"
	"github.com/tendermint/tendermint/libs/log"

	icacontrollerkeeper "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/controller/keeper"
	clientkeeper "github.com/cosmos/ibc-go/v3/modules/core/02-client/keeper"
	connectionkeeper "github.com/cosmos/ibc-go/v3/modules/core/03-connection/keeper"
	connectiontypes "github.com/cosmos/ibc-go/v3/modules/core/03-connection/types"
	channelkeeper "github.com/cosmos/ibc-go/v3/modules/core/04-channel/keeper"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v3/modules/core/24-host"
	"github.com/defund-labs/defund/x/broker/types"

	transferkeeper "github.com/cosmos/ibc-go/v3/modules/apps/transfer/keeper"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	querykeeper "github.com/defund-labs/defund/x/query/keeper"
)

type Keeper struct {
	cdc codec.Codec

	storeKey sdk.StoreKey

	scopedKeeper        capabilitykeeper.ScopedKeeper
	icaControllerKeeper icacontrollerkeeper.Keeper
	transferKeeper      transferkeeper.Keeper
	channelKeeper       channelkeeper.Keeper
	connectionKeeper    connectionkeeper.Keeper
	clientKeeper        clientkeeper.Keeper
	queryKeeper         querykeeper.Keeper
}

func NewKeeper(cdc codec.Codec, storeKey sdk.StoreKey, iaKeeper icacontrollerkeeper.Keeper, scopedKeeper capabilitykeeper.ScopedKeeper, transferKeeper transferkeeper.Keeper, channelKeeper channelkeeper.Keeper, connectionkeeper connectionkeeper.Keeper, clientkeeper clientkeeper.Keeper, querykeeper querykeeper.Keeper) Keeper {
	return Keeper{
		cdc:      cdc,
		storeKey: storeKey,

		scopedKeeper:        scopedKeeper,
		icaControllerKeeper: iaKeeper,
		transferKeeper:      transferKeeper,
		channelKeeper:       channelKeeper,
		connectionKeeper:    connectionkeeper,
		clientKeeper:        clientkeeper,
		queryKeeper:         querykeeper,
	}
}

// Logger returns the application logger, scoped to the associated module
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s-%s", host.ModuleName, types.ModuleName))
}

// ClaimCapability claims the channel capability passed via the OnOpenChanInit callback
func (k *Keeper) ClaimCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) error {
	return k.scopedKeeper.ClaimCapability(ctx, cap, name)
}

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
	if err := k.icaControllerKeeper.RegisterInterchainAccount(ctx, connectionID, owner); err != nil {
		return err
	}
	return nil
}

// GetIBCConnection is a wrapper to get a connection from id
func (k Keeper) GetIBCConnection(ctx sdk.Context, connectionID string) (connectiontypes.ConnectionEnd, bool) {
	connection, found := k.connectionKeeper.GetConnection(ctx, connectionID)
	return connection, found
}

func (k Keeper) OnRedeemSuccess(ctx sdk.Context) error {
	return nil
}

func (k Keeper) OnRedeemFailure(ctx sdk.Context) error {
	return nil
}

func (k Keeper) OnRebalanceSuccess(ctx sdk.Context) error {
	return nil
}

func (k Keeper) OnRebalanceFailure(ctx sdk.Context) error {
	return nil
}

// OnAcknowledgementPacketSuccess is the logic called on the IBC OnAcknowledgementPacket callback.
// In this function we check the incoming packet as an ICS-27 packet. We then take that ICS-27
// packet and run through each ICA message for the ack.
//
// If the ICA message is an ICA IBC transfer message then we know it is a Redeem message for redeeming ETF shares.
// We check to see the sequence corresponds with a redeem store, if it does, we then proceed to check
// if it was a successful msg. If so we then burn the fund shares held by the module account from the initial Redeem flow.
// If it failed, we mark the redeem in store as failed and then proceed as usual.
//
// If the ICA message is an ICA Swap Message, we know it is a rebalance workflow, and we mark the rebalance
// from pending to complete.
func (k Keeper) OnAcknowledgementPacketSuccess(ctx sdk.Context, packet channeltypes.Packet, ack channeltypes.Acknowledgement, txMsgData *sdk.TxMsgData) error {
	// loop through each ICA msg in the tx (one ack respresents one tx)
	for _, msgData := range txMsgData.Data {
		switch msgData.MsgType {
		case sdk.MsgTypeURL(&transfertypes.MsgTransfer{}):
			msgResponse := &transfertypes.MsgTransferResponse{}
			if err := proto.Unmarshal(msgData.Data, msgResponse); err != nil {
				return sdkerrors.Wrapf(sdkerrors.ErrJSONUnmarshal, "cannot unmarshal send response message: %s", err.Error())
			}
			k.Logger(ctx).Info("message response in ICS-27 packet response", "response", msgResponse.String())

			// Run redeem success logic
			k.OnRedeemSuccess(ctx)

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
		case sdk.MsgTypeURL(&transfertypes.MsgTransfer{}):
			msgResponse := &transfertypes.MsgTransferResponse{}
			if err := proto.Unmarshal(msgData.Data, msgResponse); err != nil {
				return sdkerrors.Wrapf(sdkerrors.ErrJSONUnmarshal, "cannot unmarshal send response message: %s", err.Error())
			}
			k.Logger(ctx).Info("message response in ICS-27 packet response", "response", msgResponse.String())

			// Run redeem failure logic
			k.OnRedeemFailure(ctx)

			return nil
		default:
			return nil
		}
	}
	return nil
}

func (k Keeper) OnAcknowledgementPacket(ctx sdk.Context, packet channeltypes.Packet, ack channeltypes.Acknowledgement, txMsgData *sdk.TxMsgData) error {
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
