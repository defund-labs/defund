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
	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	querytypes "github.com/defund-labs/defund/x/query/types"
	osmosisbalancertypes "github.com/osmosis-labs/osmosis/v7/x/gamm/pool-models/balancer"
	osmosisgammtypes "github.com/osmosis-labs/osmosis/v7/x/gamm/types"
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

// Create sends an IBC transfer to the account specified and creates a pending invest store.
// Initializes the investment process which continues in Broker module in OnAckRec.
func (k Keeper) Create(ctx sdk.Context, id string, sendFrom string, fund types.Fund, channel string, amount sdk.Coin, sender string, receiver string, timeoutHeight clienttypes.Height, timeoutTimestamp uint64) error {
	portid, err := icatypes.NewControllerPortID(fund.Address)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "could not find account: %s", err)
	}
	sequence, found := k.channelKeeper.GetNextSequenceSend(ctx, portid, channel)
	if !found {
		return sdkerrors.Wrapf(types.ErrNextSequenceNotFound, "failed to retrieve the next sequence for channel %s and port %s", channel, portid)
	}
	invest := types.Create{
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

func (k Keeper) EndBlocker(ctx sdk.Context) error {
	err := k.CreatePriceEndBlock(ctx)
	if err != nil {
		ctx.Logger().Debug("Error Creating Fund Price Log:", err.Error())
	}
	return nil
}
