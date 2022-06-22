package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

type AccountKeeper interface {
	// Methods imported from account should be defined here
	NewAccount(sdk.Context, authtypes.AccountI) authtypes.AccountI
	NewAccountWithAddress(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI

	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI
	GetAllAccounts(ctx sdk.Context) []authtypes.AccountI
	SetAccount(ctx sdk.Context, acc authtypes.AccountI)

	IterateAccounts(ctx sdk.Context, process func(authtypes.AccountI) bool)

	ValidatePermissions(macc authtypes.ModuleAccountI) error

	GetModuleAddress(moduleName string) sdk.AccAddress
	GetModuleAddressAndPermissions(moduleName string) (addr sdk.AccAddress, permissions []string)
	GetModuleAccountAndPermissions(ctx sdk.Context, moduleName string) (authtypes.ModuleAccountI, []string)
	GetModuleAccount(ctx sdk.Context, moduleName string) authtypes.ModuleAccountI
	SetModuleAccount(ctx sdk.Context, macc authtypes.ModuleAccountI)

	UnmarshalAccount(bz []byte) (authtypes.AccountI, error)
}

type BrokerKeeper interface {
	GetBrokerAccount(ctx sdk.Context, ConnectionId string, portId string) (string, bool)
}
