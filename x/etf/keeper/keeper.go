package keeper

import (
	"fmt"
	"strconv"

	"github.com/tendermint/tendermint/libs/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	"github.com/defund-labs/defund/x/etf/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey

		accountKeeper types.AccountKeeper
		bankKeeper    types.BankKeeper
		brokerKeeper  types.BrokerKeeper
		queryKeeper   types.InterqueryKeeper
		channelKeeper types.ChannelKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,

	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	channelKeeper types.ChannelKeeper,
	interqueryKeeper types.InterqueryKeeper,
	brokerKeeper types.BrokerKeeper,
) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,

		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
		channelKeeper: channelKeeper,
		queryKeeper:   interqueryKeeper,
		brokerKeeper: brokerKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// Keeper function that send an IBC transfer to the account specified and creates a pending invest store.
// Initializes the investment process.
func (k Keeper) Invest(ctx sdk.Context, id string, sendFrom string, fund types.Fund, channel string, amount sdk.Coin, sender string, receiver string, timeoutHeight clienttypes.Height, timeoutTimestamp uint64) error {
	portid, err := icatypes.NewControllerPortID(fund.Address)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "could not find account: %s", err)
	}
	sequence, found := k.channelKeeper.GetNextSequenceSend(ctx, portid, channel)
	if !found {
		return sdkerrors.Wrapf(types.ErrNextSequenceNotFound, "failed to retrieve the next sequence for channel %s and port %s", channel, portid)
	}
	invest := types.Invest{
		Id:       id,
		Creator:  sender,
		Fund:     &fund,
		Amount:   &amount,
		Channel:  channel,
		Sequence: strconv.FormatUint(sequence, 10),
		Status:   "pending",
	}
	k.SetInvest(ctx, invest)
	k.brokerKeeper.SendTransfer(ctx, fund.Address, channel, amount, sender, receiver, timeoutHeight, timeoutTimestamp)
	return nil
}

// Keeper function to rebalance an ETF using the IBC Broker Module
func (k Keeper) RebalanceETF(ctx sdk.Context) {

}

// Keeper function that goes through all ETFs and rebalances them if the current height is a rebalance period
func (k Keeper) RebalanceEndBlocker(ctx sdk.Context) {

}
