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
		brokerKeeper:  brokerKeeper,
	}
}

// Logger returns the module logger
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// helper function to check if a string is within a slice
func contains(list []string, str string) bool {
	for _, value := range list {
		if value == str {
			return true
		}
	}

	return false
}

// CreateFundPrice creates a current fund price for a fund symbol
func (k Keeper) CreateFundPrice(ctx sdk.Context, symbol string) (sdk.Int, error) {
	funds := k.GetAllFund(ctx)
	for _, fund := range funds {
		if fund.Symbol == symbol {
			for _, holding := range fund.Holdings {
				balances, err := k.queryKeeper.GetHighestHeightPoolBalance(ctx, holding.PoolId)
				if err != nil {
					return sdk.Int{}, err
				}
				if balances.Coins[0].Denom == holding.Token {
					percentDec := sdk.NewInt(holding.Percent/100)
					priceComp := sdk.NewInt(balances.Coins[0].Amount.Int64()).Mul(percentDec)
					return priceComp, nil
				}
			}
		}
	}
	return sdk.Int{}, sdkerrors.Wrapf(types.ErrFundNotFound, "No fund found (%s)", symbol)
}

// Invest sends an IBC transfer to the account specified and creates a pending invest store.
// Initializes the investment process which continues in Broker module in OnAckRec.
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

// RebalanceETF rebalances an ETF using the IBC Broker Module
func (k Keeper) RebalanceETF(ctx sdk.Context) {

}

// RebalanceEndBlocker goes through all ETFs and rebalances them if the current height is a rebalance period
func (k Keeper) RebalanceEndBlocker(ctx sdk.Context) {

}
