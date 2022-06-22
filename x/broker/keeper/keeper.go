package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	"github.com/tendermint/tendermint/libs/log"

	icacontrollerkeeper "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/controller/keeper"
	clientkeeper "github.com/cosmos/ibc-go/v3/modules/core/02-client/keeper"
	connectionkeeper "github.com/cosmos/ibc-go/v3/modules/core/03-connection/keeper"
	connectiontypes "github.com/cosmos/ibc-go/v3/modules/core/03-connection/types"
	channelkeeper "github.com/cosmos/ibc-go/v3/modules/core/04-channel/keeper"
	host "github.com/cosmos/ibc-go/v3/modules/core/24-host"
	"github.com/defund-labs/defund/x/broker/types"

	transferkeeper "github.com/cosmos/ibc-go/v3/modules/apps/transfer/keeper"
)

type Keeper struct {
	cdc codec.Codec

	storeKey sdk.StoreKey

	scopedKeeper        capabilitykeeper.ScopedKeeper
	icaControllerKeeper icacontrollerkeeper.Keeper
	transferKeeper      transferkeeper.Keeper
	channelKeeper       channelkeeper.Keeper
	connectionKeeper    connectionkeeper.Keeper
	clientKeeper        clientkeeper.Keeper
}

func NewKeeper(cdc codec.Codec, storeKey sdk.StoreKey, iaKeeper icacontrollerkeeper.Keeper, scopedKeeper capabilitykeeper.ScopedKeeper, transferKeeper transferkeeper.Keeper, channelKeeper channelkeeper.Keeper, connectionkeeper connectionkeeper.Keeper, clientkeeper clientkeeper.Keeper) Keeper {
	return Keeper{
		cdc:      cdc,
		storeKey: storeKey,

		scopedKeeper:        scopedKeeper,
		icaControllerKeeper: iaKeeper,
		transferKeeper:      transferKeeper,
		channelKeeper:       channelKeeper,
		connectionKeeper:    connectionkeeper,
		clientKeeper:        clientkeeper,
	}
}

// Logger returns the application logger, scoped to the associated module
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s-%s", host.ModuleName, types.ModuleName))
}

// ClaimCapability claims the channel capability passed via the OnOpenChanInit callback
func (k *Keeper) ClaimCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) error {
	return k.scopedKeeper.ClaimCapability(ctx, cap, name)
}

// Get a ICA account address on a host chain
func (k Keeper) GetBrokerAccount(ctx sdk.Context, ConnectionId string, portId string) (string, bool) {
	acc, found := k.icaControllerKeeper.GetInterchainAccountAddress(ctx, ConnectionId, portId)
	if found != true {
		return "", false
	}
	return acc, true
}

// Registers an ICA account on a host chain
func (k Keeper) RegisterBrokerAccount(ctx sdk.Context, connectionID, owner string) error {
	if err := k.icaControllerKeeper.RegisterInterchainAccount(ctx, connectionID, owner); err != nil {
		return err
	}
	return nil
}

// GetIBCConnection is a wrapper to get a connection from id
func (k Keeper) GetIBCConnection(ctx sdk.Context, connectionID string) (connectiontypes.ConnectionEnd, bool) {
	connection, found := k.connectionKeeper.GetConnection(ctx, connectionID)
	return connection, found
}
