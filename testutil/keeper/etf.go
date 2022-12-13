package keeper

import (
	"testing"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	icacontrollerkeeper "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/controller/keeper"
	icacontrollertypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/controller/types"
	"github.com/defund-labs/defund/app"
	"github.com/defund-labs/defund/x/etf/keeper"
	"github.com/defund-labs/defund/x/etf/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"
)

func GetSubspace(keeper paramskeeper.Keeper, moduleName string) paramstypes.Subspace {
	subspace, _ := paramskeeper.Keeper.GetSubspace(keeper, moduleName)
	return subspace
}

func EtfKeeper(db *dbm.MemDB, t testing.TB) (keeper.Keeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)
	storeKeyAcc := sdk.NewKVStoreKey(authtypes.StoreKey)
	storeKeyIcaC := sdk.NewKVStoreKey(icacontrollertypes.StoreKey)

	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, sdk.StoreTypeMemory, nil)
	stateStore.MountStoreWithDB(storeKeyAcc, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(storeKeyIcaC, sdk.StoreTypeIAVL, db)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	encoding := app.MakeEncodingConfig(app.ModuleBasics)

	maccPerms := map[string][]string{
		authtypes.FeeCollectorName: nil,
	}

	a := app.New(log.NewNopLogger(), db, nil, true, map[int64]bool{}, app.DefaultNodeHome, 0, encoding, app.GetEnabledProposals(),
		simapp.EmptyAppOptions{}, []wasm.Option{})

	a.AccountKeeper = authkeeper.NewAccountKeeper(
		codec.NewProtoCodec(registry), storeKeyAcc, a.GetSubspace(authtypes.ModuleName), authtypes.ProtoBaseAccount, maccPerms,
	)
	a.ICAControllerKeeper = icacontrollerkeeper.NewKeeper(
		codec.NewProtoCodec(registry), storeKeyIcaC, a.GetSubspace(icacontrollertypes.SubModuleName),
		a.IBCKeeper.ChannelKeeper, // may be replaced with middleware such as ics29 fee
		a.IBCKeeper.ChannelKeeper, &a.IBCKeeper.PortKeeper,
		a.ScopedICAControllerKeeper, a.MsgServiceRouter(),
	)

	a.WasmInternalKeeper = wasmkeeper.NewDefaultPermissionKeeper(a.WasmKeeper)

	capKeeper := *capabilitykeeper.NewKeeper(codec.NewProtoCodec(registry), storeKey, memStoreKey)

	scopedEtfKeeper := capKeeper.ScopeToModule("etf")

	k := keeper.NewKeeper(
		codec.NewProtoCodec(registry),
		storeKey,
		memStoreKey,
		scopedEtfKeeper,
		a.AccountKeeper,
		a.BankKeeper,
		a.IBCKeeper.ChannelKeeper,
		a.QueryKeeper,
		a.BrokerKeeper,
		a.IBCKeeper.ConnectionKeeper,
		a.IBCKeeper.ClientKeeper,
		a.ICAControllerKeeper,
		a.TransferKeeper,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())
	return k, ctx
}
