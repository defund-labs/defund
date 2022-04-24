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
	ibctransfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	channelkeeper "github.com/cosmos/ibc-go/v3/modules/core/04-channel/keeper"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v3/modules/core/24-host"
	"github.com/defund-labs/defund/x/broker/types"

	transferkeeper "github.com/cosmos/ibc-go/v3/modules/apps/transfer/keeper"
	etfkeeper "github.com/defund-labs/defund/x/etf/keeper"
	liquiditytypes "github.com/tendermint/liquidity/x/liquidity/types"
)

type Keeper struct {
	cdc codec.Codec

	storeKey sdk.StoreKey

	scopedKeeper        capabilitykeeper.ScopedKeeper
	icaControllerKeeper icacontrollerkeeper.Keeper
	transferKeeper      transferkeeper.Keeper
	channelKeeper       channelkeeper.Keeper
	etfkeeper           etfkeeper.Keeper
}

func NewKeeper(cdc codec.Codec, storeKey sdk.StoreKey, iaKeeper icacontrollerkeeper.Keeper, scopedKeeper capabilitykeeper.ScopedKeeper, transferKeeper transferkeeper.Keeper, channelKeeper channelkeeper.Keeper, etfKeeper etfkeeper.Keeper) Keeper {
	return Keeper{
		cdc:      cdc,
		storeKey: storeKey,

		scopedKeeper:        scopedKeeper,
		icaControllerKeeper: iaKeeper,
		transferKeeper:      transferKeeper,
		channelKeeper:       channelKeeper,
		etfkeeper:           etfKeeper,
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
func (k Keeper) ICAIBCTransfer(ctx sdk.Context, msgs []*ibctransfertypes.MsgTransfer, owner string, connectionID string) (sequence uint64, err error) {
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

// SendTransfer sends an IBC transfer
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

// CreateCosmosTrade creates and returns a MsgSwapWithinBatch msg to be run on Cosmos Hub via ICA
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

// SendCosmosTrades creates and sends a list of trades via ICA to the Gravity Dex (on Cosmos Hub)
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

// HandleICASwapInvest handles the logic when a swap pertaining to an invest comes in as an IBC ack, error or timeout
func (k Keeper) HandleICASwapInvest(ctx sdk.Context, msgData *liquiditytypes.MsgSwapWithinBatch, packet channeltypes.Packet, ackErr bool, timeout bool) error {
	invest, err := k.etfkeeper.GetInvestBySequence(ctx, packet.Sequence, packet.SourceChannel)
	if err != nil {
		return err
	}
	if !ackErr && !timeout {
		// Change the invest status from pending to complete once swap goes through
		invest.Status = "complete"
		k.etfkeeper.SetInvest(ctx, invest)
	}
	if ackErr {
		/// Change the invest status from pending to error and log error if ica error occurs
		invest.Status = "error"
		invest.Error = msgData.String()
		k.etfkeeper.SetInvest(ctx, invest)
	}
	if timeout {
		/// Change the invest status from pending to timeout if ica timeout occurs
		invest.Status = "timeout"
		k.etfkeeper.SetInvest(ctx, invest)
	}
	return nil
}

// HandleICATransferInvest handles the logic when a IBC transfer pertaining to an invest comes in as an IBC ack, error or timeout
func (k Keeper) HandleICATransferInvest(ctx sdk.Context, msgData *ibctransfertypes.MsgTransfer, packet channeltypes.Packet, ackErr bool, timeout bool) error {
	invest, err := k.etfkeeper.GetInvestBySequence(ctx, packet.Sequence, packet.SourceChannel)
	if err != nil {
		return err
	}
	// Handle successfull execution logic
	if !ackErr && !timeout {
		// Change the invest status from pending to complete once swap goes through
		invest.Status = "pending-swap"
		k.etfkeeper.SetInvest(ctx, invest)
	}
	// Handle unsuccessfull execution logic
	if ackErr {
		// Change the invest status from pending to error and log error if ica error occurs
		invest.Status = "error"
		invest.Error = msgData.String()
		k.etfkeeper.SetInvest(ctx, invest)
	}
	if timeout {
		/// Change the invest status from pending to timeout if ica timeout occurs
		invest.Status = "timeout"
		k.etfkeeper.SetInvest(ctx, invest)
	}
	return nil
}

// HandleICATransferUninvest handles the logic when a IBC transfer pertaining to an uninvest comes in as an IBC ack or timeout
func (k Keeper) HandleICATransferUninvest(ctx sdk.Context, msgData *ibctransfertypes.MsgTransfer, packet channeltypes.Packet, ackErr bool, timeout bool) error {
	uninvest, err := k.etfkeeper.GetUninvestBySequence(ctx, packet.Sequence, packet.SourceChannel)
	if err != nil {
		return err
	}
	// Handle successfull execution logic
	if !ackErr && !timeout {
		// Change the uninvest status from pending to complete once swap goes through
		uninvest.Status = "complete"
		k.etfkeeper.SetUninvest(ctx, uninvest)
	}
	// Handle unsuccessfull execution logic
	if ackErr {
		// Change the uninvest status from pending to error and log error if ica error occurs
		uninvest.Status = "error"
		uninvest.Error = msgData.String()
		k.etfkeeper.SetUninvest(ctx, uninvest)
	}
	if timeout {
		/// Change the uninvest status from pending to timeout if ica timeout occurs
		uninvest.Status = "timeout"
		k.etfkeeper.SetUninvest(ctx, uninvest)
	}
	return nil
}

// HandleICASwapRebalance handles the logic when a swap pertaining to a fund rebalance comes in as an IBC ack, error or timeout
func (k Keeper) HandleICASwapRebalance(ctx sdk.Context, msgData *liquiditytypes.MsgSwapWithinBatch, packet channeltypes.Packet, ackErr bool, timeout bool) error {
	return nil
}

// HandleICASwap looks up the channel and sequence in all rebalance, invest and uninvest messages to determine which ICA
// response belongs to either invest or uninvest
func (k Keeper) HandleICASwap(ctx sdk.Context, msgData *liquiditytypes.MsgSwapWithinBatch, packet channeltypes.Packet, ackErr bool, timeout bool) error {
	return nil
}

// HandleICASend handles an ICA IBC send from another chain. In Defunds case, we need to handle it if it is an uninvest
func (k Keeper) HandleICASend(ctx sdk.Context, msgData *ibctransfertypes.MsgTransfer, packet channeltypes.Packet, ackErr bool, timeout bool) error {
	return nil
}
