package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	"github.com/tendermint/tendermint/libs/log"

	icacontrollerkeeper "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/controller/keeper"
	clientkeeper "github.com/cosmos/ibc-go/v4/modules/core/02-client/keeper"
	connectionkeeper "github.com/cosmos/ibc-go/v4/modules/core/03-connection/keeper"
	channelkeeper "github.com/cosmos/ibc-go/v4/modules/core/04-channel/keeper"
	host "github.com/cosmos/ibc-go/v4/modules/core/24-host"
	"github.com/defund-labs/defund/x/broker/types"

	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	transferkeeper "github.com/cosmos/ibc-go/v4/modules/apps/transfer/keeper"
	etfkeeper "github.com/defund-labs/defund/x/etf/keeper"
	querykeeper "github.com/defund-labs/defund/x/query/keeper"
	osmosisgammtypes "github.com/osmosis-labs/osmosis/v8/x/gamm/types"
)

type Keeper struct {
	cdc codec.Codec

	storeKey sdk.StoreKey

	paramstore paramtypes.Subspace

	scopedKeeper        capabilitykeeper.ScopedKeeper
	icaControllerKeeper icacontrollerkeeper.Keeper
	transferKeeper      transferkeeper.Keeper
	channelKeeper       channelkeeper.Keeper
	connectionKeeper    connectionkeeper.Keeper
	clientKeeper        clientkeeper.Keeper
	queryKeeper         querykeeper.Keeper
	etfKeeper           etfkeeper.Keeper
	bankKeeper          bankkeeper.Keeper
}

func NewKeeper(cdc codec.Codec, storeKey sdk.StoreKey, paramstore paramtypes.Subspace, iaKeeper icacontrollerkeeper.Keeper, transferKeeper transferkeeper.Keeper, channelKeeper channelkeeper.Keeper, connectionkeeper connectionkeeper.Keeper, clientkeeper clientkeeper.Keeper, querykeeper querykeeper.Keeper, etfkeeper etfkeeper.Keeper, bankkeeper bankkeeper.Keeper) Keeper {
	if !paramstore.HasKeyTable() {
		paramstore = paramstore.WithKeyTable(types.ParamKeyTable())
	}
	return Keeper{
		cdc:      cdc,
		storeKey: storeKey,

		paramstore: paramstore,

		icaControllerKeeper: iaKeeper,
		transferKeeper:      transferKeeper,
		channelKeeper:       channelKeeper,
		connectionKeeper:    connectionkeeper,
		clientKeeper:        clientkeeper,
		queryKeeper:         querykeeper,
		etfKeeper:           etfkeeper,
		bankKeeper:          bankkeeper,
	}
}

func (k *Keeper) SetEtfKeeper(keeper etfkeeper.Keeper) {
	k.etfKeeper = keeper
}

// Logger returns the application logger, scoped to the associated module
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s-%s", host.ModuleName, types.ModuleName))
}

// CreateDefundQueries creates all the repeated interqueries for broker chains
func (k Keeper) CreateDefundQueries(ctx sdk.Context) error {
	// Run every 10th block (1 minute)
	if ctx.BlockHeight()%10 == 0 {
		// Add Osmosis broker interquery for all pools
		k.CreateQueryOsmosisPools(ctx)
	}
	return nil
}

// QueryOsmosisPool sets an interquery request in store for a Osmosis pool to be run by relayers
func (k Keeper) CreateQueryOsmosisPool(ctx sdk.Context, poolId uint64) error {
	broker, found := k.GetBroker(ctx, "osmosis")
	if !found {
		return sdkerrors.Wrap(types.ErrBrokerNotFound, "broker not found: osmosis")
	}
	path := "/store/gamm/key"
	connectionid := broker.ConnectionId
	key := osmosisgammtypes.GetKeyPrefixPools(poolId)
	timeoutHeight := uint64(ctx.BlockHeight() + 50)
	storeid := fmt.Sprintf("osmosis-%d", poolId)
	chainid := "osmo-test-4"

	err := k.queryKeeper.CreateInterqueryRequest(ctx, chainid, storeid, path, key, timeoutHeight, connectionid)
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
