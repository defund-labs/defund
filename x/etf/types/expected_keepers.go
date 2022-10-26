package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	connectiontypes "github.com/cosmos/ibc-go/v4/modules/core/03-connection/types"
	channeltypes "github.com/cosmos/ibc-go/v4/modules/core/04-channel/types"
	"github.com/cosmos/ibc-go/v4/modules/core/exported"
	brokertypes "github.com/defund-labs/defund/x/broker/types"
	querytypes "github.com/defund-labs/defund/x/query/types"
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

type BankKeeper interface {
	// Methods imported from bank should be defined here
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromModuleToModule(ctx sdk.Context, senderPool, recipientPool string, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error

	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error

	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, name string, amt sdk.Coins) error

	SetDenomMetaData(ctx sdk.Context, denomMetaData banktypes.Metadata)

	// Only needed for simulation interface matching
	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	// SetBalances(ctx sdk.Context, addr sdk.AccAddress, balances sdk.Coins) error
	LockedCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	GetSupply(ctx sdk.Context, denom string) sdk.Coin
	UndelegateCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	DelegateCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
}

type BrokerKeeper interface {
	GetBroker(ctx sdk.Context, id string) (val brokertypes.Broker, found bool)
	SetBroker(ctx sdk.Context, broker brokertypes.Broker)
	SetTransfer(ctx sdk.Context, transfer brokertypes.Transfer)
	GetAllTransfer(ctx sdk.Context) (list []brokertypes.Transfer)
	GetTransfer(ctx sdk.Context, index string) (val brokertypes.Transfer, found bool)
	RemoveTransfer(ctx sdk.Context, id string)
	GetPoolFromBroker(ctx sdk.Context, brokerId string, poolId uint64) (val brokertypes.Source, found bool)
	GetParam(ctx sdk.Context, key []byte) *brokertypes.BaseDenoms
	GetRebalance(ctx sdk.Context, index string) (val brokertypes.Rebalance, found bool)
	RemoveRebalance(ctx sdk.Context, id string)
	RemoveRedeem(ctx sdk.Context, id string)
	GetRedeem(ctx sdk.Context, id string) (val brokertypes.Redeem, found bool)
	SetRedeem(ctx sdk.Context, redeem brokertypes.Redeem)
	SetRebalance(ctx sdk.Context, rebalance brokertypes.Rebalance)
}

type ChannelKeeper interface {
	GetChannel(ctx sdk.Context, srcPort, srcChan string) (channel channeltypes.Channel, found bool)
	GetNextSequenceSend(ctx sdk.Context, portID, channelID string) (uint64, bool)
}

type InterqueryKeeper interface {
	GetInterqueryResult(ctx sdk.Context, index string) (querytypes.InterqueryResult, bool)
	GetAllInterqueryResult(ctx sdk.Context) (list []querytypes.InterqueryResult)
	RemoveInterqueryResult(ctx sdk.Context, storeid string)
	CreateInterqueryRequest(ctx sdk.Context, chainid string, storeid string, path string, key []byte, timeoutheight uint64, connectionid string) error
}

type ConnectionKeeper interface {
	GetConnection(ctx sdk.Context, connectionID string) (connectiontypes.ConnectionEnd, bool)
}

type ClientKeeper interface {
	GetClientState(ctx sdk.Context, clientID string) (exported.ClientState, bool)
}

type TransferKeeper interface {
	SendTransfer(ctx sdk.Context, sourcePort, sourceChannel string, token sdk.Coin, sender sdk.AccAddress, receiver string, timeoutHeight clienttypes.Height, timeoutTimestamp uint64) error
}
