package keeper

// All Custom Transfer Logic Lives Here

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	icatypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v4/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v4/modules/core/24-host"
)

// Helper function that creates and returns a MsgTransfer msg type to be run via ICA
func (k Keeper) CreateMultiSendMsg(ctx sdk.Context, fromAddress string, toAddress string, amount sdk.Coins) (*banktypes.MsgSend, error) {
	send := banktypes.MsgSend{
		FromAddress: fromAddress,
		ToAddress:   toAddress,
		Amount:      amount,
	}
	send.ValidateBasic()
	return &send, nil
}

// Creates an ICA Bank Send msg on a host/broker ICA chain to
// send funds from an account on the host chain to someone else on the host chain
func (k Keeper) SendIBCSend(ctx sdk.Context, msgs []*banktypes.MsgSend, owner string, connectionID string) (sequence uint64, channel string, err error) {
	seralizeMsgs := []sdk.Msg{}
	for _, msg := range msgs {
		msg.ValidateBasic()
		seralizeMsgs = append(seralizeMsgs, msg)
	}

	portID, err := icatypes.NewControllerPortID(owner)
	if err != nil {
		return sequence, channel, err
	}

	channel, found := k.icaControllerKeeper.GetActiveChannelID(ctx, connectionID, portID)
	if !found {
		return sequence, channel, sdkerrors.Wrapf(icatypes.ErrActiveChannelNotFound, "failed to retrieve active channel for port %s", portID)
	}

	chanCap, found := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(portID, channel))
	if !found {
		return sequence, channel, sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	data, err := icatypes.SerializeCosmosTx(k.cdc, seralizeMsgs)
	if err != nil {
		return sequence, channel, err
	}

	packetData := icatypes.InterchainAccountPacketData{
		Type: icatypes.EXECUTE_TX,
		Data: data,
	}

	// timeoutTimestamp set to max value with the unsigned bit shifted to sastisfy hermes timestamp conversion
	// it is the responsibility of the auth module developer to ensure an appropriate timeout timestamp
	timeoutTimestamp := uint64(time.Now().Add(time.Minute).UnixNano())
	sequence, err = k.icaControllerKeeper.SendTx(ctx, chanCap, connectionID, portID, packetData, uint64(timeoutTimestamp))
	if err != nil {
		return sequence, channel, err
	}

	return sequence, channel, nil
}

// Creates an ICA Transfer msg on a host/broker ICA chain to send funds through IBC to Defund
func (k Keeper) SendIBCTransferICA(ctx sdk.Context, msgs []*ibctransfertypes.MsgTransfer, owner string, connectionID string) (sequence uint64, channel string, err error) {
	seralizeMsgs := []sdk.Msg{}
	for _, msg := range msgs {
		msg.ValidateBasic()
		seralizeMsgs = append(seralizeMsgs, msg)
	}

	portID, err := icatypes.NewControllerPortID(owner)
	if err != nil {
		return 0, "", err
	}

	channel, found := k.icaControllerKeeper.GetActiveChannelID(ctx, connectionID, portID)
	if !found {
		return 0, "", sdkerrors.Wrapf(icatypes.ErrActiveChannelNotFound, "failed to retrieve active channel for port %s", portID)
	}

	chanCap, found := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(portID, channel))
	if !found {
		return 0, "", sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	data, err := icatypes.SerializeCosmosTx(k.cdc, seralizeMsgs)
	if err != nil {
		return sequence, "", err
	}

	packetData := icatypes.InterchainAccountPacketData{
		Type: icatypes.EXECUTE_TX,
		Data: data,
	}

	// timeoutTimestamp set to max value with the unsigned bit shifted to sastisfy hermes timestamp conversion
	// it is the responsibility of the auth module developer to ensure an appropriate timeout timestamp
	timeoutTimestamp := uint64(time.Now().Add(time.Minute).UnixNano())
	sequence, err = k.icaControllerKeeper.SendTx(ctx, chanCap, connectionID, portID, packetData, uint64(timeoutTimestamp))
	if err != nil {
		return sequence, "", err
	}

	return sequence, channel, nil
}
