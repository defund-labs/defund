package etf

import (
	"errors"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	proto "github.com/gogo/protobuf/proto"

	channeltypes "github.com/cosmos/ibc-go/v4/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v4/modules/core/05-port/types"
	host "github.com/cosmos/ibc-go/v4/modules/core/24-host"
	ibcexported "github.com/cosmos/ibc-go/v4/modules/core/exported"
	brokerkeeper "github.com/defund-labs/defund/x/broker/keeper"
	etfkeeper "github.com/defund-labs/defund/x/etf/keeper"
	"github.com/defund-labs/defund/x/etf/types"
)

var _ porttypes.IBCModule = IBCModule{}

// IBCModule implements the ICS26 interface for interchain accounts controller chains
type IBCModule struct {
	keeper       etfkeeper.Keeper
	brokerKeeper brokerkeeper.Keeper
}

// NewIBCModule creates a new IBCModule given the keeper
func NewIBCModule(k etfkeeper.Keeper, brokerkeeper brokerkeeper.Keeper) IBCModule {
	return IBCModule{
		keeper:       k,
		brokerKeeper: brokerkeeper,
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

	return version, nil
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
	// on a ICA channel close we must ensure the fund rebalance is set to false
	addr := strings.Split(portID, "icacontroller-")[1]
	if len(addr) > 0 {
		fund, err := im.keeper.GetFundByDefundAddr(ctx, addr)
		if err != nil {
			return err
		}
		fund.Rebalancing = false
		im.keeper.SetFund(ctx, fund)
	}
	return nil
}

// OnChanCloseConfirm implements the IBCModule interface
func (im IBCModule) OnChanCloseConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	return nil
}

// OnRecvPacket implements the IBCModule interface. A successful acknowledgement
// is returned if the packet data is succesfully decoded and the receive application
// logic returns without error.
func (im IBCModule) OnRecvPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) ibcexported.Acknowledgement {
	return channeltypes.NewErrorAcknowledgement(errors.New("cannot receive packet via etf module"))
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
	im.keeper.OnAcknowledgementPacketICA(ctx, packet, ack, txMsgData)

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
	redeem, redeemExists := im.brokerKeeper.GetRedeem(ctx, id)
	if redeemExists {
		// on a timeout, ICA channel closes so automatically set the fund as not autorebalancing
		fund, found := im.keeper.GetFund(ctx, redeem.Fund)
		if !found {
			return sdkerrors.Wrapf(types.ErrFundNotFound, "fund %s not found", redeem.Fund)
		}
		fund.Rebalancing = false
		im.keeper.SetFund(ctx, fund)

		im.keeper.Logger(ctx).Error("redeem %s timed out. Running the redeem timeout logic.", id)
		err := im.keeper.OnRedeemFailure(ctx, packet, redeem)
		if err != nil {
			im.keeper.Logger(ctx).Error("Error occured during run of ICA callback.", "callback", "OnRedeemFailure", "error", err.Error())
		}
	} else {
		im.keeper.Logger(ctx).Debug(fmt.Sprintf("Redeem %s not found. Skipping redeem timeout logic.", id))
	}
	// get the rebalance from the store. If not found return nil and do not run logic if found run the redeem timeout logic
	rebalance, rebalanceExists := im.brokerKeeper.GetRebalance(ctx, id)
	if rebalanceExists {
		// on a timeout, ICA channel closes so automatically set the fund as not autorebalancing
		fund, found := im.keeper.GetFund(ctx, rebalance.Fund)
		if !found {
			return sdkerrors.Wrapf(types.ErrFundNotFound, "fund %s not found", rebalance.Fund)
		}
		fund.Rebalancing = false
		im.keeper.SetFund(ctx, fund)

		im.keeper.Logger(ctx).Error(fmt.Sprintf("rebalance %s timed out. Running the rebalance timeout logic.", id))
		err := im.keeper.OnRebalanceFailure(ctx, rebalance, fund.Symbol)
		if err != nil {
			im.keeper.Logger(ctx).Error("Error occured during run of ICA callback.", "callback", "OnRebalanceFailure", "error", err.Error())
		}
	} else {
		im.keeper.Logger(ctx).Debug(fmt.Sprintf("Rebalance %s not found. Skipping rebalance timeout logic.", id))
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
