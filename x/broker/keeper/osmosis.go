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

	osmosisbalancertypes "github.com/osmosis-labs/osmosis/v7/x/gamm/pool-models/balancer"
	osmosisgammtypes "github.com/osmosis-labs/osmosis/v7/x/gamm/types"
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
	path := "/store/gamm/key/"
	connectionid := "connection-0"
	key := osmosisgammtypes.GetKeyPrefixPools(poolId)
	timeoutHeight := uint64(ctx.BlockHeight() + 10)
	storeid := fmt.Sprintf("osmosis-%d", poolId)

	err := k.queryKeeper.CreateInterqueryRequest(ctx, storeid, path, key, timeoutHeight, connectionid)
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
	return sdkerrors.Wrapf(types.ErrInvalidPool, "pool (%s) not found", poolId)
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
func (k Keeper) GetOsmosisPool(ctx sdk.Context, poolId string) (osmosisbalancertypes.Pool, error) {
	query, found := k.queryKeeper.GetInterqueryResult(ctx, fmt.Sprintf("osmosis-%s", poolId))
	if !found {
		return osmosisbalancertypes.Pool{}, sdkerrors.Wrapf(types.ErrInvalidPool, "could not find pool query for %s", fmt.Sprintf("osmosis-%s", poolId))
	}
	var pool = osmosisbalancertypes.Pool{}
	err := json.Unmarshal(query.Data, &pool)
	if err != nil {
		return osmosisbalancertypes.Pool{}, sdkerrors.Wrapf(types.ErrMarshallingError, "cannot decode osmosis pool query (%s)", strings.Split(query.Storeid, "-")[1])
	}
	return pool, nil
}

// GetPriceOfAssetFromQuery gets a pool from an interquery result and computes the price of that pool pair
func (k Keeper) CalculateSpotPrice(ctx sdk.Context, poolId string, tokenInDenom string, tokenOutDenom string) (sdk.Dec, error) {
	query, found := k.queryKeeper.GetInterqueryResult(ctx, fmt.Sprintf("osmosis-%s", poolId))
	if !found {
		return sdk.Dec{}, sdkerrors.Wrapf(types.ErrInvalidPool, "could not find pool query for %s", fmt.Sprintf("osmosis-%s", poolId))
	}
	var pool = osmosisbalancertypes.Pool{}
	err := json.Unmarshal(query.Data, &pool)
	if err != nil {
		return sdk.Dec{}, sdkerrors.Wrapf(types.ErrMarshallingError, "cannot decode osmosis pool query (%s)", strings.Split(query.Storeid, "-")[1])
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
		sdk.ZeroDec(),
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
