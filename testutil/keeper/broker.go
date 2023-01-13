package keeper

import (
	"testing"

	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
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
	paramsStoreKey := sdk.NewKVStoreKey(paramstypes.StoreKey)
	paramsStoreKeyT := sdk.NewTransientStoreKey(paramstypes.TStoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(paramsStoreKeyT, sdk.StoreTypeTransient, db)
	stateStore.MountStoreWithDB(paramsStoreKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, sdk.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	encoding := app.MakeEncodingConfig(app.ModuleBasics)

	a := app.New(log.NewNopLogger(), db, nil, true, map[int64]bool{}, app.DefaultNodeHome, 0, encoding, app.GetEnabledProposals(),
		simapp.EmptyAppOptions{}, []wasm.Option{})

	cdc := codec.NewProtoCodec(app.MakeEncodingConfig(app.ModuleBasics).InterfaceRegistry)

	paramsSubspace := paramstypes.NewSubspace(cdc,
		codec.NewLegacyAmino(),
		storeKey,
		memStoreKey,
		"DexParams",
	)
	_ = paramskeeper.NewKeeper(cdc, codec.NewLegacyAmino(), paramsStoreKey, paramsStoreKeyT)

	k := keeper.NewKeeper(
		cdc,
		storeKey,
		paramsSubspace,
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
