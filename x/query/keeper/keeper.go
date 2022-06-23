package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/defund-labs/defund/x/query/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey

		accountKeeper types.AccountKeeper
		brokerKeeper  types.BrokerKeeper
	}
)

type PoolsKey struct {
}

type BalanceKey struct {
	Address string `json:"address"`
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,

	accountkeeper types.AccountKeeper,
	brokerkeeper types.BrokerKeeper,

) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,

		accountKeeper: accountkeeper,
		brokerKeeper:  brokerkeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) NewQueryAddress(id uint64) sdk.AccAddress {
	key := append([]byte("query"), sdk.Uint64ToBigEndian(id)...)
	return address.Module("query", key)
}

func (k Keeper) CreateInterqueryRequest(ctx sdk.Context, storeid string, path string, key []byte, timeoutheight uint64, clientid string) error {
	var queryModuleAddress authtypes.ModuleAccountI
	if k.accountKeeper.GetModuleAccount(ctx, "query") == nil {
		queryAddress := k.NewQueryAddress(1)
		queryModuleAddress = authtypes.NewModuleAccount(
			authtypes.NewBaseAccountWithAddress(
				queryAddress,
			),
			"query",
		)
		k.accountKeeper.SetModuleAccount(ctx, queryModuleAddress)
	} else {
		queryModuleAddress = k.accountKeeper.GetModuleAccount(ctx, "query")
	}
	interquery := types.Interquery{
		Creator:       queryModuleAddress.GetAddress().String(),
		Storeid:       storeid,
		Path:          path,
		Key:           key,
		TimeoutHeight: timeoutheight,
		ClientId:      clientid,
	}
	k.SetInterquery(ctx, interquery)

	k.Logger(ctx).Info(fmt.Sprintf("Interquery request for path %s on clientid of %s has been initiated", path, clientid))

	return nil
}

func (k Keeper) TimeoutInterqueries(ctx sdk.Context) {
	// Get all interqueries from store
	interqueries := k.GetAllInterquery(ctx)

	// Loop through each query and timeout if not valid
	for _, query := range interqueries {
		if uint64(ctx.BlockHeight()) > query.TimeoutHeight {
			queryTimeout := types.InterqueryTimeoutResult{
				Creator:       query.Creator,
				Storeid:       query.Storeid,
				TimeoutHeight: query.TimeoutHeight,
				ClientId:      query.ClientId,
			}
			// Set the query as a timed out interquery in store
			k.SetInterqueryTimeoutResult(ctx, queryTimeout)
			// Remove/cleanup the interquery from pending interqueries
			k.RemoveInterquery(ctx, query.Storeid)
		}
	}
}

func (k Keeper) ModuleEndBlocker(ctx sdk.Context) {

	//Timeout all timedout/invalid interqueries at the beginning of the end block
	k.TimeoutInterqueries(ctx)

	//Get all interqueries that have not been submitted yet
	pendingqueries := k.GetAllInterquery(ctx)

	// Create holder for all events
	events := sdk.Events{}

	for _, query := range pendingqueries {

		event := sdk.NewEvent(
			types.EventTypeQuery,
			sdk.NewAttribute(types.AttributeKeyQueryClientId, query.ClientId),
			sdk.NewAttribute(types.AttributeKeyQueryPath, query.Path),
			sdk.NewAttribute(types.AttributeKeyQueryStoreid, query.Storeid),
		)

		events = append(events, event)

	}

	//Emit the query event
	ctx.EventManager().EmitEvents(events)
}
