package broker

import (
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	"github.com/defund-labs/defund/x/broker/keeper"
	proto "github.com/gogo/protobuf/proto"

	channeltypes "github.com/cosmos/ibc-go/v4/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v4/modules/core/05-port/types"
	host "github.com/cosmos/ibc-go/v4/modules/core/24-host"
	ibcexported "github.com/cosmos/ibc-go/v4/modules/core/exported"
	etfkeeper "github.com/defund-labs/defund/x/etf/keeper"
)

var _ porttypes.IBCModule = IBCModule{}

// IBCModule implements the ICS26 interface for interchain accounts controller chains
type IBCModule struct {
	keeper    keeper.Keeper
	etfKeeper etfkeeper.Keeper
}

// NewIBCModule creates a new IBCModule given the keeper
func NewIBCModule(k keeper.Keeper, etfkeeper etfkeeper.Keeper) IBCModule {
	return IBCModule{
		keeper:    k,
		etfKeeper: etfkeeper,
	}
}

// OnChanOpenInit implements the IBCModule interface
func (im IBCModule) OnChanOpenInit(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID string,
	channelID string,
	chanCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	version string,
) (string, error) {
	// Claim channel capability passed back by IBC module
	if err := im.keeper.ClaimCapability(ctx, chanCap, host.ChannelCapabilityPath(portID, channelID)); err != nil {
		return "", err
	}

	return "", nil
}

// OnChanOpenTry implements the IBCModule interface
func (im IBCModule) OnChanOpenTry(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID,
	channelID string,
	chanCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	counterpartyVersion string,
) (string, error) {
	return "", nil
}

// OnChanOpenAck implements the IBCModule interface
func (im IBCModule) OnChanOpenAck(
	ctx sdk.Context,
	portID,
	channelID string,
	counterpartyChannelID string,
	counterpartyVersion string,
) error {
	return nil
}

// OnChanOpenConfirm implements the IBCModule interface
func (im IBCModule) OnChanOpenConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	return nil
}

// OnChanCloseInit implements the IBCModule interface
func (im IBCModule) OnChanCloseInit(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	// Disallow user-initiated channel closing for broker channels
	return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "user cannot close channel")
}

// OnChanCloseConfirm implements the IBCModule interface
func (im IBCModule) OnChanCloseConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	// Disallow user-initiated channel closing for broker channels
	return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "user cannot close channel")
}

// OnRecvPacket implements the IBCModule interface. A successful acknowledgement
// is returned if the packet data is succesfully decoded and the receive application
// logic returns without error.
func (im IBCModule) OnRecvPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) ibcexported.Acknowledgement {
	return channeltypes.NewErrorAcknowledgement(errors.New("cannot receive packet via broker module"))
}

// OnAcknowledgementPacket implements the IBCModule interface
func (im IBCModule) OnAcknowledgementPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	acknowledgement []byte,
	relayer sdk.AccAddress,
) error {
	// unmarshal the ack to be used later
	var ack channeltypes.Acknowledgement
	if err := channeltypes.SubModuleCdc.UnmarshalJSON(acknowledgement, &ack); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal Broker packet acknowledgement: %v", err)
	}
	// unmarshal the msg data from the tx to be used later
	txMsgData := &sdk.TxMsgData{}
	if err := proto.Unmarshal(ack.GetResult(), txMsgData); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal Broker tx message data: %v", err)
	}
	// if the length of the msg data is 0 skip/return, otherwise run through logic
	switch len(txMsgData.Data) {
	case 0:
		return nil
	default:
		err := im.keeper.OnAcknowledgementPacket(ctx, packet, ack, txMsgData)
		if err != nil {
			return err
		}
	}
	return nil
}

// OnTimeoutPacket implements the IBCModule interface.
func (im IBCModule) OnTimeoutPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) error {
	id := fmt.Sprintf("%s-%d", packet.SourceChannel, packet.Sequence)
	// get the redeem from the store. If not found return nil and do not run logic if found run the redeem timeout logic
	redeem, redeemExists := im.etfKeeper.GetRedeem(ctx, id)
	if redeemExists {
		im.keeper.Logger(ctx).Error("redeem %s timed out. Running the redeem timeout logic.", id)
		im.etfKeeper.OnRedeemFailure(ctx, packet, redeem)
	}
	// get the rebalance from the store. If not found return nil and do not run logic if found run the redeem timeout logic
	rebalance, rebalanceExists := im.etfKeeper.GetRebalance(ctx, id)
	if rebalanceExists {
		im.keeper.Logger(ctx).Error("rebalance %s timed out. Running the rebalance timeout logic.", id)
		im.etfKeeper.OnRebalanceFailure(ctx, rebalance, rebalance.Fund)
	}

	return nil
}

// NegotiateAppVersion implements the IBCModule interface
func (im IBCModule) NegotiateAppVersion(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionID string,
	portID string,
	counterparty channeltypes.Counterparty,
	proposedVersion string,
) (string, error) {
	return "", nil
}
