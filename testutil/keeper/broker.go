package keeper

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/defund-labs/defund/app"
	"github.com/defund-labs/defund/x/broker/keeper"
	"github.com/defund-labs/defund/x/broker/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"
)

func BrokerKeeper(db *dbm.MemDB, t testing.TB) (*keeper.Keeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, sdk.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	encoding := app.MakeEncodingConfig(app.ModuleBasics)

	a := app.New(log.NewNopLogger(), db, nil, true, map[int64]bool{}, app.DefaultNodeHome, 0, encoding,
		simapp.EmptyAppOptions{})

	k := keeper.NewKeeper(
		codec.NewProtoCodec(registry),
		storeKey,
		a.GetSubspace(types.ModuleName),
		a.ICAControllerKeeper,
		a.TransferKeeper,
		a.IBCKeeper.ChannelKeeper,
		a.IBCKeeper.ConnectionKeeper,
		a.IBCKeeper.ClientKeeper,
		a.QueryKeeper,
		a.EtfKeeper,
		a.BankKeeper,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())
	return &k, ctx
}
