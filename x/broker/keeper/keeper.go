package keeper

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	"github.com/tendermint/tendermint/libs/log"

	icacontrollerkeeper "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/controller/keeper"
	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	channelkeeper "github.com/cosmos/ibc-go/v3/modules/core/04-channel/keeper"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v3/modules/core/24-host"
	"github.com/defund-labs/defund/x/broker/types"

	transferkeeper "github.com/cosmos/ibc-go/v3/modules/apps/transfer/keeper"
	liquiditytypes "github.com/tendermint/liquidity/x/liquidity/types"
)

type Keeper struct {
	cdc codec.Codec

	storeKey sdk.StoreKey

	scopedKeeper        capabilitykeeper.ScopedKeeper
	icaControllerKeeper icacontrollerkeeper.Keeper
	transferKeeper      transferkeeper.Keeper
	channelKeeper       channelkeeper.Keeper
}

func NewKeeper(cdc codec.Codec, storeKey sdk.StoreKey, iaKeeper icacontrollerkeeper.Keeper, scopedKeeper capabilitykeeper.ScopedKeeper, transferKeeper transferkeeper.Keeper, channelKeeper channelkeeper.Keeper) Keeper {
	return Keeper{
		cdc:      cdc,
		storeKey: storeKey,

		scopedKeeper:        scopedKeeper,
		icaControllerKeeper: iaKeeper,
		transferKeeper:      transferKeeper,
		channelKeeper:       channelKeeper,
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
	if found != true {
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

// Creates an ICA Transfer msg on a host ICA chain
func (k Keeper) IBCTransfer(ctx sdk.Context) {

}

// Sends an IBC transfer
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

// Helper function that creates and returns a MsgSwapWithinBatch msg to be run on Cosmos Hub via ICA
func (k Keeper) CreateCosmosTrade(ctx sdk.Context, trader string, poolid uint64, offercoin sdk.Coin, demandcoin string, swapfeerate sdk.Dec, limitprice sdk.Dec) (*liquiditytypes.MsgSwapWithinBatch, error) {
	trade := liquiditytypes.MsgSwapWithinBatch{
		SwapRequesterAddress: trader,
		PoolId:               poolid,
		SwapTypeId:           1,
		OfferCoin:            offercoin,
		DemandCoinDenom:      demandcoin,
		OfferCoinFee:         liquiditytypes.GetOfferCoinFee(offercoin, swapfeerate),
		OrderPrice:           limitprice,
	}
	trade.ValidateBasic()
	return &trade, nil
}

// This keeper function creates and sends a list of trades via ICA to the Gravity Dex (on Cosmos Hub)
func (k Keeper) SendCosmosTrades(ctx sdk.Context, msgs []*liquiditytypes.MsgSwapWithinBatch, owner string, connectionID string) (sequence uint64, err error) {

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
