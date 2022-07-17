package keeper

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/defund-labs/defund/x/etf/types"
	"github.com/tendermint/tendermint/libs/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	porttypes "github.com/cosmos/ibc-go/v3/modules/core/05-port/types"
	brokertypes "github.com/defund-labs/defund/x/broker/types"
	querytypes "github.com/defund-labs/defund/x/query/types"
	osmosisbalancertypes "github.com/osmosis-labs/osmosis/v7/x/gamm/pool-models/balancer"
	osmosisgammtypes "github.com/osmosis-labs/osmosis/v7/x/gamm/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey

		accountKeeper       types.AccountKeeper
		bankKeeper          types.BankKeeper
		brokerKeeper        types.BrokerKeeper
		queryKeeper         types.InterqueryKeeper
		channelKeeper       types.ChannelKeeper
		ics4Wrapper         porttypes.ICS4Wrapper
		connectionKeeper    types.ConnectionKeeper
		clientKeeper        types.ClientKeeper
		icaControllerKeeper types.ICAControllerKeeper
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
	connectionKeeper types.ConnectionKeeper,
	clientKeeper types.ClientKeeper,
	iaKeeper types.ICAControllerKeeper,
) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,

		accountKeeper:       accountKeeper,
		bankKeeper:          bankKeeper,
		channelKeeper:       channelKeeper,
		queryKeeper:         interqueryKeeper,
		brokerKeeper:        brokerKeeper,
		connectionKeeper:    connectionKeeper,
		clientKeeper:        clientKeeper,
		icaControllerKeeper: iaKeeper,
	}
}

// SetICS4Wrapper sets the ICS4 wrapper to the keeper.
// It panics if already set
func (k *Keeper) SetICS4Wrapper(ics4Wrapper porttypes.ICS4Wrapper) {
	if k.ics4Wrapper != nil {
		panic("ICS4 wrapper already set")
	}

	k.ics4Wrapper = ics4Wrapper
}

// Logger returns the module logger
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// helper function to check if a osmosis pool contains denom specified
func containsAssets(assets []osmosisgammtypes.PoolAsset, denom string) bool {
	for _, pool := range assets {
		if pool.Token.Denom == denom {
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

// CreateShares sends a multi-send of assets to create ETF shares from creator to the module account
// which then sends an IBC transfer to the fund account on the broker chain and creates a pending transfer store.
// Initializes the create shares process which continues in Broker module in OnAckRec.
func (k Keeper) CreateShares(ctx sdk.Context, fund types.Fund, channel string, tokens []*sdk.Coin, creator string, timeoutHeight clienttypes.Height, timeoutTimestamp uint64) error {
	// Need to convert the coins to plain coins for multi send
	coins := sdk.Coins{}
	for _, token := range tokens {
		coins = append(coins, *token)
	}
	creatorAcc, err := sdk.AccAddressFromBech32(creator)
	if err != nil {
		return err
	}
	fundAcc, err := sdk.AccAddressFromBech32(fund.Address)
	if err != nil {
		return err
	}

	// get the ica account for the fund on the broker chain
	portID, err := icatypes.NewControllerPortID(fund.Address)
	if err != nil {
		return err
	}
	fundBrokerAddress, found := k.brokerKeeper.GetBrokerAccount(ctx, fund.ConnectionId, portID)
	if !found {
		return sdkerrors.Wrapf(brokertypes.ErrIBCAccountNotExist, "failed to find ica account for owner %s on connection %s and port %s", fund.Address, fund.ConnectionId, portID)
	}

	// send the tokens to the Defund fund account to ensure that we receive the
	// tokens correctly.
	err = k.bankKeeper.SendCoins(ctx, creatorAcc, fundAcc, sdk.NewCoins(coins...))
	if err != nil {
		return err
	}

	// for each token send IBC transfer to move funds to broker chain. logic continues in ibc callbacks
	for _, token := range tokens {
		sequence, err := k.brokerKeeper.SendTransfer(ctx, channel, *token, fund.Address, fundBrokerAddress, timeoutHeight, timeoutTimestamp)
		if err != nil {
			return err
		}
		transfer := brokertypes.Transfer{
			Id:       fmt.Sprintf("%s-%d", channel, sequence),
			Channel:  channel,
			Sequence: sequence,
			Status:   "tranferring",
			Token:    token,
			Sender:   fund.Address,
			Receiver: fundBrokerAddress,
		}
		k.brokerKeeper.SetTransfer(ctx, transfer)
	}

	// compute the amount of etf shares this creator is given
	numETFShares, err := k.GetAmountETFSharesForTokens(ctx, fund, tokens)
	if err != nil {
		return err
	}
	newETFCoins := sdk.NewCoins(numETFShares)

	// finally mint coins (to module account) and then send them to the creator of the create
	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, newETFCoins)
	if err != nil {
		return err
	}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creatorAcc, newETFCoins)
	if err != nil {
		return err
	}

	return nil
}

// RedeemShares sends an ICA MultiSend message to the broker chain to be run on that chain.
// Initializes the redemption of shares process which continues in Broker module in OnAckRec.
func (k Keeper) RedeemShares(ctx sdk.Context, id string, fund types.Fund, channel string, amount sdk.Coin, fundAccount string, receiver string) error {
	receiverAcc, err := sdk.AccAddressFromBech32(receiver)
	if err != nil {
		return err
	}
	portid, err := icatypes.NewControllerPortID(fund.Address)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "could not find account: %s", err)
	}
	sequence, found := k.channelKeeper.GetNextSequenceSend(ctx, portid, channel)
	if !found {
		return sdkerrors.Wrapf(types.ErrNextSequenceNotFound, "failed to retrieve the next sequence for channel %s and port %s", channel, portid)
	}
	redeem := types.Redeem{
		Id:       id,
		Creator:  receiver,
		Fund:     &fund,
		Amount:   &amount,
		Channel:  channel,
		Sequence: strconv.FormatUint(sequence, 10),
		Status:   "pending",
	}

	// get the amount of tokens that these shares represent
	ownership, err := k.GetOwnershipSharesInFund(ctx, fund, amount)
	if err != nil {
		return err
	}

	msg, err := k.brokerKeeper.CreateMultiSendMsg(ctx, fundAccount, receiver, sdk.NewCoins(ownership...))
	if err != nil {
		return err
	}
	// take the fund etf shares and escrow them in the module account. in the ack callback, on success
	// we will burn these shares. If unsuccessful we will send them back to the user (same on timeout).
	k.bankKeeper.SendCoinsFromAccountToModule(ctx, receiverAcc, types.ModuleName, sdk.NewCoins(amount))
	// create the ica multi send message
	k.brokerKeeper.SendIBCSend(ctx, []*banktypes.MsgSend{msg}, fund.Address, fund.ConnectionId)

	k.SetRedeem(ctx, redeem)

	return nil
}

// DecodeLiquiditySourceQuery decodes a query based on if/what broker the query is for
// returns error if not supported/cannot unmarshall
func (k Keeper) DecodeLiquiditySourceQuery(ctx sdk.Context, query querytypes.InterqueryResult) (osmosisbalancertypes.Pool, error) {
	switch strings.Split(query.Storeid, "-")[0] {
	case "osmosis":
		var pool = osmosisbalancertypes.Pool{}
		err := json.Unmarshal(query.Data, &pool)
		if err != nil {
			return pool, sdkerrors.Wrapf(types.ErrMarshallingError, "cannot decode osmosis pool query (%s)", strings.Split(query.Storeid, "-")[1])
		}
		return pool, nil
	default:
		var pool = osmosisbalancertypes.Pool{}
		return pool, sdkerrors.Wrapf(types.ErrMarshallingError, "cannot decode liquidity source query. not supported (%s)", strings.Split(query.Storeid, "-")[0])
	}
}

// CheckHoldings checks to make sure the specified holdings and the pool for each holding are valid
// by checking the interchain queried pools for the broker specified
func (k Keeper) CheckHoldings(ctx sdk.Context, brokerId string, holdings []types.Holding) error {
	percentCheck := uint64(0)
	for _, holding := range holdings {
		// Add percent composition to percentCheck to later confirm adds to 100%
		percentCheck = percentCheck + uint64(holding.Percent)
		poolQuery, found := k.queryKeeper.GetInterqueryResult(ctx, fmt.Sprint(holding.PoolId))
		if !found {
			return sdkerrors.Wrapf(types.ErrInvalidPool, "could not find pool details for (broker: %s, pool: %s)", brokerId, holding.PoolId)
		}
		pool, err := k.DecodeLiquiditySourceQuery(ctx, poolQuery)
		if err != nil {
			return err
		}
		// Checks to see if the holding pool contains the holding token specified and if not returns error
		if !containsAssets(pool.PoolAssets, holding.Token) {
			return sdkerrors.Wrapf(types.ErrInvalidDenom, "invalid/unsupported denom (%s) in pool (%s)", holding.Token, holding.PoolId)
		}
	}
	// Make sure all fund holdings add up to 100%
	if percentCheck != uint64(100) {
		return sdkerrors.Wrapf(types.ErrPercentComp, "percent composition must add up to 100%")
	}
	return nil
}

// getOsmosisRoutes is a helper function that looks up the Osmosis broker, takes in the currentDenom, needDenom
// and returns a list of the best routes to go through. It does this by first checking if a currentDenom
// has a direct pool with uosmo, if it does not, it then finds a curentDenom -> usomo with uosmo -> needDenom
// pair to create the routes needed to go from currentDenom -> needDenom.
func (k Keeper) getRoutes(ctx sdk.Context, currentDenom string, needDenom string) {}

// SendRebalanceTx sends one ICA tx for the fund with a list of swap msg's to rebalance
// an ETF. Each swap message will have multiple routes within it to swap to the needed
// rebalanced asset (see getRoutes above).
// For calculation of rebalances needed, each holding is converted to the base denom,
// then each holdings current weight in the base denom is subtracted from the expected composition.
// Then each needed composition that is positive (over owned) is matched with each negative composition (under owned)
// to create a swap message until no negative compositions exist.
func (k Keeper) SendRebalanceTx(ctx sdk.Context, fund types.Fund) {

}
