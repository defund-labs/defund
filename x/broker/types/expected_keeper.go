package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	connectiontypes "github.com/cosmos/ibc-go/v3/modules/core/03-connection/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	"github.com/cosmos/ibc-go/v3/modules/core/exported"
	querytypes "github.com/defund-labs/defund/x/query/types"
)

type TransferKeeper interface {
	SendTransfer(ctx sdk.Context, sourcePort, sourceChannel string, token sdk.Coin, sender sdk.AccAddress, receiver string, timeoutHeight clienttypes.Height, timeoutTimestamp uint64) error
}

type ChannelKeeper interface {
	GetChannel(ctx sdk.Context, srcPort, srcChan string) (channel channeltypes.Channel, found bool)
	GetNextSequenceSend(ctx sdk.Context, portID, channelID string) (uint64, bool)
}

type ConnectionKeeper interface {
	GetConnection(ctx sdk.Context, connectionID string) (connectiontypes.ConnectionEnd, bool)
	ConnOpenInit(ctx sdk.Context, clientID string, counterparty connectiontypes.Counterparty, version *connectiontypes.Version, delayPeriod uint64) (string, error)
}

type ClientKeeper interface {
	CreateClient(ctx sdk.Context, clientState exported.ClientState, consensusState exported.ConsensusState) (string, error)
}

type InterqueryKeeper interface {
	CreateInterqueryRequest(ctx sdk.Context, storeid string, path string, key []byte, timeoutheight uint64, clientid string) error
	GetInterqueryResult(ctx sdk.Context, index string) (querytypes.InterqueryResult, bool)
}
