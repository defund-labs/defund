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
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	"github.com/defund-labs/defund/x/etf/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey

		paramSpace paramtypes.Subspace

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

	paramSubspace paramtypes.Subspace,

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

		paramSpace: paramSubspace,

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

// helper function to check if a list of coins contains a token denom
func containsdenom(list []sdk.Coin, denom string) bool {
	for _, value := range list {
		if value.Denom == denom {
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
func (k Keeper) CreateFundPrice(ctx sdk.Context, symbol string) (sdk.Coin, error) {
	fund, found := k.GetFund(ctx, symbol)
	if !found {
		return sdk.Coin{}, sdkerrors.Wrapf(types.ErrFundNotFound, "Could not find fund (%s)", symbol)
	}
	// If there are no shares thus no holdings, funds price is 1 $BASEDENOM
	if fund.Shares.Amount.Uint64() == uint64(0) {
		price := sdk.NewCoin(fund.BaseDenom, sdk.NewInt(1000000))
		return price, nil
	}
	comp := []sdk.Dec{}
	// Get the current balances from interquery of the broker account on the broker chain
	portID, err := icatypes.NewControllerPortID(fund.Address)
	if err != nil {
		return sdk.Coin{}, err
	}
	// Get the ica broker account
	brokeraccount, found := k.brokerKeeper.GetBrokerAccount(ctx, fund.ConnectionId, portID)
	if !found {
		return sdk.Coin{}, sdkerrors.Wrapf(icatypes.ErrInterchainAccountNotFound, "failed to retrieve interchain account for owner %s", fund.Address)
	}
	// use ica broker account to get the interquery balance of that account
	balances, err := k.queryKeeper.GetBalance(ctx, brokeraccount)
	if err != nil {
		return sdk.Coin{}, err
	}
	for _, holding := range fund.Holdings {
		// Check to ensure we have a query for each holding currently or error out
		check := containsdenom(balances, holding.Token)
		if !check {
			return sdk.Coin{}, sdkerrors.Wrapf(types.ErrNoBalanceForDenom, "failed to retrieve a balance for holding %s. Cannot produce fund price", holding.Token)
		}
		for _, balance := range balances {
			if holding.Token == balance.Denom && fund.BaseDenom != holding.Token {
				baseAmount := balances[0].Amount.ToDec()
				tokenAmount := balances[1].Amount.ToDec()
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
		}
	}

	total := sum(comp)
	price := sdk.NewCoin(fund.BaseDenom, sdk.NewInt(total.RoundInt64()))

	return price, nil
}

// CreateAllFundPriceEndBlock is a function that runs at each end block that logs the fund price for each fund at the current height
func (k Keeper) CreateAllFundPriceEndBlock(ctx sdk.Context) error {
	funds := k.GetAllFund(ctx)
	for _, fund := range funds {
		price, err := k.CreateFundPrice(ctx, fund.Symbol)
		if err != nil {
			return err
		}

		fundPrice := types.FundPrice{
			Height: uint64(ctx.BlockHeight()),
			Amount: &price,
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
		Sequence: sequence,
		Status:   "pending",
	}
	err = k.brokerKeeper.SendTransfer(ctx, fund.Address, channel, amount, sender, receiver, timeoutHeight, timeoutTimestamp)
	if err != nil {
		return err
	}

	k.SetInvest(ctx, invest)

	return nil
}

func (k Keeper) EndBlocker(ctx sdk.Context) error {
	err := k.CreateAllFundPriceEndBlock(ctx)
	if err != nil {
		ctx.Logger().Error("Error Creating Fund Price Log:", err.Error())
	}
	return nil
}
