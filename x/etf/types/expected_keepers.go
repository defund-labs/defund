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
	osmosisgammtypes "github.com/osmosis-labs/osmosis/v8/x/gamm/types"
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
	GetBrokerAccount(ctx sdk.Context, ConnectionId string, portIDstring string) (string, bool)
	RegisterBrokerAccount(ctx sdk.Context, connectionID, owner string) error
	SetTransfer(ctx sdk.Context, transfer brokertypes.Transfer)
	GetAllTransfer(ctx sdk.Context) (list []brokertypes.Transfer)
	GetTransfer(ctx sdk.Context, index string) (val brokertypes.Transfer, found bool)
	RemoveTransfer(ctx sdk.Context, id string)
	GetOsmosisPool(ctx sdk.Context, poolId uint64) (p osmosisgammtypes.PoolI, err error)
	GetOsmosisBalance(ctx sdk.Context, account string) (banktypes.Balance, error)
	CalculateOsmosisSpotPrice(ctx sdk.Context, poolId uint64, tokenInDenom string, tokenOutDenom string) (sdk.Dec, error)
	GetPoolFromBroker(ctx sdk.Context, brokerId string, poolId uint64) (val brokertypes.Source, found bool)
	CreateQueryOsmosisBalance(ctx sdk.Context, account string) error
	CreateOsmosisTrade(ctx sdk.Context, trader string, routes []osmosisgammtypes.SwapAmountInRoute, tokenin sdk.Coin, tokenoutminamount sdk.Int) (*osmosisgammtypes.MsgSwapExactAmountIn, error)
	SendOsmosisTrades(ctx sdk.Context, msgs []*osmosisgammtypes.MsgSwapExactAmountIn, owner string, connectionID string) (sequence uint64, err error)
}

type ChannelKeeper interface {
	GetChannel(ctx sdk.Context, srcPort, srcChan string) (channel channeltypes.Channel, found bool)
	GetNextSequenceSend(ctx sdk.Context, portID, channelID string) (uint64, bool)
}

type InterqueryKeeper interface {
	GetInterqueryResult(ctx sdk.Context, index string) (querytypes.InterqueryResult, bool)
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
