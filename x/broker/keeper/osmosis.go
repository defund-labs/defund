package keeper

// All Osmosis Logic Lives Here

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/defund-labs/defund/x/etf/types"

	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v3/modules/core/24-host"

	osmosisgammtypes "github.com/osmosis-labs/osmosis/x/gamm/types"
)

type PoolsKey struct {
}

type BalanceKey struct {
	Address string `json:"address"`
}

// QueryOsmosisPool sets an interquery request in store for a Osmosis pool to be run by relayers
func (k Keeper) QueryOsmosisPool(ctx sdk.Context, poolId uint64) error {
	path := "/store/gamm/key/"
	clientid := "07-tendermint-0"
	key := osmosisgammtypes.GetKeyPrefixPools(poolId)
	timeoutHeight := uint64(ctx.BlockHeight() + 10)
	storeid := fmt.Sprintf("osmosis-%d", poolId)

	err := k.queryKeeper.CreateInterqueryRequest(ctx, storeid, path, key, timeoutHeight, clientid)
	if err != nil {
		return err
	}
	return nil
}

// QueryOsmosisPools queries all pools specified in the Osmosis broker
func (k Keeper) QueryOsmosisPools(ctx sdk.Context) {
	broker, found := k.GetBroker(ctx, "osmosis")
	if !found {
		return
	}
	for _, pool := range broker.Pools {
		err := k.QueryOsmosisPool(ctx, pool.PoolId)
		if err != nil {
			ctx.Logger().Debug(fmt.Sprintf("error creating osmosis pool query (%d): %s", pool.PoolId, err.Error()))
			continue
		}
	}
	return
}

// GetPriceOfAssetFromQuery gets a pool from an interquery result and computes the price of that pool pair
func (k Keeper) GetPriceOfAssetFromQuery(ctx sdk.Context, poolId string, tokenIn string) (sdk.Int, error) {
	query, found := k.queryKeeper.GetInterqueryResult(ctx, fmt.Sprintf("osmosis-%s", poolId))
	if !found {
		return sdk.Int{}, sdkerrors.Wrapf(types.ErrInvalidPool, "could not find pool query for %s", fmt.Sprintf("osmosis-%s", poolId))
	}
	var pool = osmosisgammtypes.Pool{}
	err := json.Unmarshal(query.Data, &pool)
	if err != nil {
		return sdk.Int{}, sdkerrors.Wrapf(types.ErrMarshallingError, "cannot decode osmosis pool query (%s)", strings.Split(query.Storeid, "-")[1])
	}
	assets := pool.PoolAssets
	firstAsset := assets[0]
	secondAsset := assets[1]
	if firstAsset.Token.Denom == tokenIn {
		return firstAsset.Token.Amount.Quo(secondAsset.Token.Amount), nil
	}
	if secondAsset.Token.Denom == tokenIn {
		return secondAsset.Token.Amount.Quo(firstAsset.Token.Amount), nil
	}
	return sdk.Int{}, sdkerrors.Wrap(types.ErrInvalidDenom, "denom could not be found in pool query")
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
