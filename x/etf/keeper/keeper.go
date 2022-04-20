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

func sum(items []sdk.Dec) sdk.Dec {
	sum := sdk.NewDec(0)
	for _, item := range items {
		sum = sum.Add(item)
	}
	return sum
}

// CreateFundPrice creates a current fund price for a fund symbol
func (k Keeper) CreateFundPrice(ctx sdk.Context, symbol string) (sdk.Dec, error) {
	fund, found := k.GetFund(ctx, symbol)
	if !found {
		return sdk.Dec{}, sdkerrors.Wrapf(types.ErrFundNotFound, "Could not find fund (%s)", symbol)
	}
	comp := []sdk.Dec{}
	for _, holding := range fund.Holdings {
		balances, err := k.queryKeeper.GetHighestHeightPoolBalance(ctx, holding.PoolId)
		if err != nil {
			return sdk.Dec{}, err
		}
		if balances[0].Denom == holding.Token && fund.BaseDenom != holding.Token {
			baseAmount := balances[0].Amount.ToDec()
			tokenAmount := balances[1].Amount.ToDec()
			priceInBaseDenom := tokenAmount.Quo(baseAmount)
			percentDec := sdk.NewDec(holding.Percent).Quo(sdk.NewDec(100))
			comp = append(comp, priceInBaseDenom.Mul(percentDec))
		}
		if balances[1].Denom == holding.Token && fund.BaseDenom != holding.Token {
			baseAmount := balances[1].Amount.ToDec()
			tokenAmount := balances[0].Amount.ToDec()
			priceInBaseDenom := tokenAmount.Quo(baseAmount)
			percentDec := sdk.NewDec(holding.Percent).Quo(sdk.NewDec(100))
			comp = append(comp, priceInBaseDenom.Mul(percentDec))
		}
		// If the holding token is the baseDenom, just multiply it by the % it represents since we already know its price relative
		// to itself. Aka -> 1/1
		if fund.BaseDenom == holding.Token {
			percentDec := sdk.NewDec(holding.Percent).Quo(sdk.NewDec(100))
			comp = append(comp, sdk.NewDec(1).Mul(percentDec))
		}
		if len(comp) == 0 {
			return sdk.Dec{}, sdkerrors.Wrapf(types.ErrFundNotFound, "No price details found for symbol (%s)", symbol)
		}
	}

	price := sum(comp)

	return price, nil
}

// CreateAllFundPriceEndBlock is a function that runs at each end block that logs the fund price for each fund at current height
func (k Keeper) CreateAllFundPriceEndBlock(ctx sdk.Context) error {
	funds := k.GetAllFund(ctx)
	for _, fund := range funds {
		price, err := k.CreateFundPrice(ctx, fund.Symbol)
		if err != nil {
			return err
		}

		fundPrice := types.FundPrice{
			Height: uint64(ctx.BlockHeight()),
			Price:  price,
			Symbol: fund.Symbol,
			Id:     fmt.Sprintf("%s-%s", fund.Symbol, strconv.FormatInt(ctx.BlockHeight(), 10)),
		}
		k.SetFundPrice(ctx, fundPrice)
	}
	return nil
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

func (k Keeper) EndBlocker(ctx sdk.Context) error {
	err := k.CreateAllFundPriceEndBlock(ctx)
	if err != nil {
		ctx.Logger().Error("Error Creating Fund Price Log:", err.Error())
	}
	return nil
}
