package keeper

import (
	"encoding/json"
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
	querykeeper "github.com/defund-labs/defund/x/query/keeper"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey

		accountKeeper types.AccountKeeper
		bankKeeper    types.BankKeeper
		brokerKeeper  types.BrokerKeeper
		queryKeeper   querykeeper.Keeper
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
) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,

		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
		channelKeeper: channelKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

type PoolsKey struct {
}

type BalanceKey struct {
	Address string `json:"address"`
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

// Helper function to be implemented in end blocker to interchain query pools on gravity dex (Cosmos Hub)
func (k Keeper) QueryGravityDex(ctx sdk.Context) error {
	path := "custom/liquidity/liquidityPools/"
	clientid := "07-tendermint-0"
	keyRaw := PoolsKey{}
	key, err := json.Marshal(keyRaw)
	heightStr := strconv.FormatInt(ctx.BlockHeight(), 10)
	timeoutHeight := uint64(ctx.BlockHeight() + 10)
	storeid := fmt.Sprintf("CosmosPools-%s", heightStr)

	err = k.queryKeeper.CreateInterqueryRequest(ctx, storeid, path, key, timeoutHeight, clientid)
	if err != nil {
		return err
	}
	return nil
}

// Helper function that creates an interquery for an account balance on Cosmos with the accountType as port of the store id
func (k Keeper) QueryFundAccounts(ctx sdk.Context, address string, accountType string) error {
	path := "custom/bank/all_balances/"
	clientid := "07-tendermint-0"
	keyRaw := BalanceKey{address}
	key, err := json.Marshal(keyRaw)
	if err != nil {
		return err
	}
	timeoutHeight := uint64(ctx.BlockHeight() + 10)
	storeid := fmt.Sprintf("%s-Balance-%s", accountType, address)

	err = k.queryKeeper.CreateInterqueryRequest(ctx, storeid, path, key, timeoutHeight, clientid)
	if err != nil {
		return err
	}
	return nil
}

// Helper function that gets all funds from store and creates an interquery for all ICA account balances
// associated with the fund on the broker chains fund ICA account.
func (k Keeper) QueryAllAccounts(ctx sdk.Context) error {
	funds := k.GetAllFund(ctx)
	for _, fund := range funds {
		portID, err := icatypes.NewControllerPortID(fund.Address)
		if err != nil {
			return status.Errorf(codes.InvalidArgument, "could not find account: %s", err)
		}

		addr, found := k.brokerKeeper.GetBrokerAccount(ctx, fund.ConnectionId, portID)
		if !found {
			return status.Errorf(codes.NotFound, "no account found for portID %s", portID)
		}
		err = k.QueryFundAccounts(ctx, addr, "Fund")
		if err != nil {
			return status.Errorf(codes.InvalidArgument, "could not find account: %s", err)
		}
	}

	return nil
}

func (k Keeper) EndBlockerRun(ctx sdk.Context) error {
	// Run every 10th block
	if ctx.BlockHeight()%10 == 0 {
		// Add gravity dex interquery
		err := k.QueryGravityDex(ctx)
		if err != nil {
			ctx.Logger().Error(fmt.Sprintf("Error Creating Cosmos GDex Pool Interquery: %s", err))
		}
		// Add gravity dex account balances interquery for all funds
		err = k.QueryAllAccounts(ctx)
		if err != nil {
			ctx.Logger().Error(fmt.Sprintf("Error Creating Cosmos GDex Account Balance Interquery: %s", err))
		}
	}
	return nil
}
