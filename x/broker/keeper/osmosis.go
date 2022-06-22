package keeper

// All Osmosis Logic Lives Here

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v3/modules/core/24-host"

	osmosisgammtypes "github.com/osmosis-labs/osmosis/x/gamm/types"
)

// Helper function that creates and returns a MsgSwapExactAmountIn msg type to be run on Osmosis via ICA
func (k Keeper) CreateOsmosisTrade(ctx sdk.Context, trader string, routes []osmosisgammtypes.SwapAmountInRoute, tokenin sdk.Coin, tokenoutminamount sdk.Int) (*osmosisgammtypes.MsgSwapExactAmountIn, error) {
	trade := osmosisgammtypes.MsgSwapExactAmountIn{
		Sender:            trader,
		Routes:            routes,
		TokenIn:           tokenin,
		TokenOutMinAmount: tokenoutminamount,
	}
	trade.ValidateBasic()
	return &trade, nil
}

// This keeper function creates and sends a list of trades via ICA to Osmosis
func (k Keeper) SendOsmosisTrades(ctx sdk.Context, msgs []*osmosisgammtypes.MsgSwapExactAmountIn, owner string, connectionID string) (sequence uint64, err error) {

	seralizeMsgs := []sdk.Msg{}
	for _, msg := range msgs {
		msg.ValidateBasic()
		seralizeMsgs = append(seralizeMsgs, msg)
	}

	portID, err := icatypes.NewControllerPortID(owner)
	if err != nil {
		return 0, err
	}

	channelID, found := k.icaControllerKeeper.GetActiveChannelID(ctx, connectionID, portID)
	if !found {
		return 0, sdkerrors.Wrapf(icatypes.ErrActiveChannelNotFound, "failed to retrieve active channel for port %s", portID)
	}

	chanCap, found := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(portID, channelID))
	if !found {
		return 0, sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	data, err := icatypes.SerializeCosmosTx(k.cdc, seralizeMsgs)
	if err != nil {
		return sequence, err
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
		return sequence, err
	}

	return sequence, nil
}
