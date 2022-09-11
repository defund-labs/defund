package keeper

// All Osmosis Logic Lives Here

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/defund-labs/defund/x/broker/types"

	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v3/modules/core/24-host"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	osmosisbalancertypes "github.com/osmosis-labs/osmosis/v8/x/gamm/pool-models/balancer"
	osmosisgammtypes "github.com/osmosis-labs/osmosis/v8/x/gamm/types"
)

type PoolsKey struct {
}

type BalanceKey struct {
	Address string `json:"address"`
}

// calcSpotPrice returns the spot price of the pool
// This is the weight-adjusted balance of the tokens in the pool.
// so spot_price = (B_in / W_in) / (B_out / W_out)
func calcSpotPrice(
	tokenBalanceIn,
	tokenWeightIn,
	tokenBalanceOut,
	tokenWeightOut sdk.Dec,
) sdk.Dec {
	number := tokenBalanceIn.Quo(tokenWeightIn)
	denom := tokenBalanceOut.Quo(tokenWeightOut)
	ratio := number.Quo(denom)

	return ratio
}

// calcSpotPriceWithSwapFee returns the spot price of the pool accounting for
// the input taken by the swap fee.
// This is the weight-adjusted balance of the tokens in the pool.
// so spot_price = (B_in / W_in) / (B_out / W_out)
// and spot_price_with_fee = spot_price / (1 - swapfee)
func calcSpotPriceWithSwapFee(
	tokenBalanceIn,
	tokenWeightIn,
	tokenBalanceOut,
	tokenWeightOut,
	swapFee sdk.Dec,
) sdk.Dec {
	spotPrice := calcSpotPrice(tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut)
	// Q: Why is this not just (1 - swapfee)
	// A: Its because its being applied to the other asset.
	// TODO: write this up more coherently
	// 1 / (1 - swapfee)
	scale := sdk.OneDec().Quo(sdk.OneDec().Sub(swapFee))

	return spotPrice.Mul(scale)
}

// QueryOsmosisPool sets an interquery request in store for a Osmosis pool to be run by relayers
func (k Keeper) CreateQueryOsmosisPool(ctx sdk.Context, poolId uint64) error {
	path := "/store/gamm/key"
	connectionid := "connection-0"
	key := osmosisgammtypes.GetKeyPrefixPools(poolId)
	timeoutHeight := uint64(ctx.BlockHeight() + 50)
	storeid := fmt.Sprintf("osmosis-%d", poolId)
	chainid := "osmosis-1"

	err := k.queryKeeper.CreateInterqueryRequest(ctx, chainid, storeid, path, key, timeoutHeight, connectionid)
	if err != nil {
		return err
	}
	return nil
}

// CreateQueryOsmosisBalance sets an interquery request in store for a Osmosis account balance to be run by relayers
func (k Keeper) CreateQueryOsmosisBalance(ctx sdk.Context, account string) error {
	path := "/store/bank/key"
	connectionid := "connection-0"
	accAddr, err := sdk.AccAddressFromBech32(account)
	if err != nil {
		return err
	}
	key := banktypes.CreateAccountBalancesPrefix(accAddr.Bytes())
	timeoutHeight := uint64(ctx.BlockHeight() + 50)
	storeid := fmt.Sprintf("account-%s", account)
	chainid := "osmosis-1"

	err = k.queryKeeper.CreateInterqueryRequest(ctx, chainid, storeid, path, key, timeoutHeight, connectionid)
	if err != nil {
		return err
	}
	return nil
}

// ChangeBrokerPoolStatus finds the pool via poolid for broker specifed and changes the status
// of the pool to the status provided
func (k Keeper) ChangeBrokerPoolStatus(ctx sdk.Context, broker types.Broker, poolId uint64, status string) error {
	for i, item := range broker.Pools {
		if item.PoolId == poolId {
			broker.Pools[i].Status = status
			k.SetBroker(ctx, broker)
			return nil
		}
	}
	return sdkerrors.Wrapf(types.ErrInvalidPool, "pool (%d) not found", poolId)
}

// QueryOsmosisPools queries all pools specified in the Osmosis broker
func (k Keeper) CreateQueryOsmosisPools(ctx sdk.Context) {
	broker, found := k.GetBroker(ctx, "osmosis")
	if !found {
		return
	}
	for _, pool := range broker.Pools {
		err := k.CreateQueryOsmosisPool(ctx, pool.PoolId)
		if err != nil {
			ctx.Logger().Debug(fmt.Sprintf("error creating osmosis pool query (%d): %s. setting pool as inactive", pool.PoolId, err.Error()))
			k.ChangeBrokerPoolStatus(ctx, broker, pool.PoolId, "inactive")
			continue
		}
	}
}

// GetOsmosisPool gets an osmosis pool from the interquery store and returns the unmarshalled pool
func (k Keeper) GetOsmosisPool(ctx sdk.Context, poolId uint64) (osmosisbalancertypes.Pool, error) {
	pool := &osmosisbalancertypes.Pool{}
	// get the pool source from broker
	brokerPool, found := k.GetPoolFromBroker(ctx, "osmosis", poolId)
	if !found {
		return *pool, sdkerrors.Wrapf(types.ErrInvalidPool, "pool %d not found on broker osmosis", poolId)
	}
	query, found := k.queryKeeper.GetInterqueryResult(ctx, fmt.Sprintf("osmosis-%d", poolId))
	if !found {
		return *pool, sdkerrors.Wrapf(types.ErrInvalidPool, "could not find pool query for %s", fmt.Sprintf("osmosis-%d", poolId))
	}
	// we must append an append to the data if it is osmosis pool due to store error
	// whichr results in failed unmarshalling for osmosis pools
	if len(brokerPool.Append) > 0 {
		query.Data = append(brokerPool.Append, query.Data...)
	}
	err := pool.Unmarshal(query.Data)
	if err != nil {
		return *pool, sdkerrors.Wrapf(types.ErrMarshallingError, "cannot decode osmosis pool query (%s). %e", query.Storeid, err)
	}
	return *pool, nil
}

// GetOsmosisBalance gets an osmosis bank balance from the interquery store and returns the unmarshalled balance
func (k Keeper) GetOsmosisBalance(ctx sdk.Context, account string) (banktypes.Balance, error) {
	query, found := k.queryKeeper.GetInterqueryResult(ctx, fmt.Sprintf("account-%s", account))
	if !found {
		return banktypes.Balance{}, sdkerrors.Wrapf(types.ErrInvalidPool, "could not find account query for %s", fmt.Sprintf("account-%s", account))
	}
	var balance = banktypes.Balance{}
	err := json.Unmarshal(query.Data, &balance)
	if err != nil {
		return banktypes.Balance{}, sdkerrors.Wrapf(types.ErrMarshallingError, "cannot decode osmosis account query (%s)", strings.Split(query.Storeid, "-")[1])
	}
	return balance, nil
}

// CalculateOsmosisSpotPrice gets a pool from an interquery result and computes the price of that pool pair
func (k Keeper) CalculateOsmosisSpotPrice(ctx sdk.Context, poolId uint64, tokenInDenom string, tokenOutDenom string) (sdk.Dec, error) {
	query, found := k.queryKeeper.GetInterqueryResult(ctx, fmt.Sprintf("osmosis-%d", poolId))
	if !found {
		return sdk.Dec{}, sdkerrors.Wrapf(types.ErrInvalidPool, "could not find pool query for %s", fmt.Sprintf("osmosis-%d", poolId))
	}
	var pool = osmosisbalancertypes.Pool{}
	err := pool.Unmarshal(query.Data)
	if err != nil {
		return sdk.Dec{}, sdkerrors.Wrapf(types.ErrMarshallingError, "cannot decode osmosis pool query (%s)", query.Storeid)
	}
	inPoolAsset, err := pool.GetPoolAsset(tokenInDenom)
	if err != nil {
		return sdk.Dec{}, err
	}
	outPoolAsset, err := pool.GetPoolAsset(tokenOutDenom)
	if err != nil {
		return sdk.Dec{}, err
	}
	// calcSpotPriceWithSwapFee, but with fee = 0
	return calcSpotPriceWithSwapFee(
		inPoolAsset.Token.Amount.ToDec(),
		inPoolAsset.Weight.ToDec(),
		outPoolAsset.Token.Amount.ToDec(),
		outPoolAsset.Weight.ToDec(),
		pool.PoolParams.SwapFee,
	), nil
}

// Helper function that creates and returns a MsgSwapExactAmountIn msg type to be run on Osmosis via ICA
func (k Keeper) CreateOsmosisTrade(ctx sdk.Context, trader string, routes []osmosisgammtypes.SwapAmountInRoute, tokenin sdk.Coin, tokenoutminamount sdk.Int) (*osmosisgammtypes.MsgSwapExactAmountIn, error) {
	trade := osmosisgammtypes.MsgSwapExactAmountIn{
		Sender:            trader,
		Routes:            routes,
		TokenIn:           tokenin,
		TokenOutMinAmount: tokenoutminamount,
	}
	trade.ValidateBasic()
	return &trade, nil
}

// This keeper function creates and sends a list of trades via ICA to Osmosis
func (k Keeper) SendOsmosisTrades(ctx sdk.Context, msgs []*osmosisgammtypes.MsgSwapExactAmountIn, owner string, connectionID string) (sequence uint64, err error) {

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

// SetPoolStatusHook is a hook run at begin blocker to set the status of all pools
// within each broker depending if the last query is recent. We state recent as updated within
// last 30 blocks (5 minutes).
// NOTE: CHANGE RECENT BLOCK PARAM TO MODULE PARAM SET BY GOVERNANCE
func (k Keeper) SetPoolStatusHookOsmosis(ctx sdk.Context) {
	k.Logger(ctx).Debug("SetPoolStatusHookOsmosis: Running Update Osmosis Pool Status Hook")
	osmosisBroker, found := k.GetBroker(ctx, "osmosis")
	if !found {
		err := sdkerrors.Wrapf(types.ErrBrokerNotFound, "broker %s not found", "osmosis")
		k.Logger(ctx).Error(err.Error())
	}
	for _, pool := range osmosisBroker.Pools {
		// lookup interquery for pool
		iq, found := k.queryKeeper.GetInterqueryResult(ctx, pool.InterqueryId)
		// if no interquery result for pool set as inactive
		if !found {
			k.Logger(ctx).Debug(fmt.Sprintf("SetPoolStatusHookOsmosis: Osmosis Pool (%d) Interquery Not Found, Updated To Inactive", pool.PoolId))
			k.ChangeBrokerPoolStatus(ctx, osmosisBroker, pool.PoolId, "inactive")
			continue
		}
		// check if interquery was updated within last 100 blocks
		updated := (uint64(ctx.BlockHeight()) - iq.LocalHeight) < 100
		k.Logger(ctx).Debug(fmt.Sprintf("SetPoolStatusHookOsmosis: Running Update Pool Status (%d)", pool.PoolId))
		// set the status to inactive or active depending on check
		if updated {
			k.Logger(ctx).Debug(fmt.Sprintf("SetPoolStatusHookOsmosis: Osmosis Pool (%d) Updated To Active", pool.PoolId))
			err := k.ChangeBrokerPoolStatus(ctx, osmosisBroker, pool.PoolId, "active")
			if err != nil {
				k.Logger(ctx).Error(err.Error())
			}
			continue
		} else {
			k.Logger(ctx).Debug(fmt.Sprintf("SetPoolStatusHookOsmosis: Osmosis Pool (%d) Updated To Inactive", pool.PoolId))
			err := k.ChangeBrokerPoolStatus(ctx, osmosisBroker, pool.PoolId, "inactive")
			if err != nil {
				k.Logger(ctx).Error(err.Error())
			}
			continue
		}
	}
}
