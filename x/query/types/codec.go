package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateInterquery{}, "query/CreateInterquery", nil)
	cdc.RegisterConcrete(&MsgCreateInterqueryResult{}, "query/CreateInterqueryResult", nil)
	cdc.RegisterConcrete(&MsgCreateInterqueryTimeout{}, "query/CreateInterqueryTimeout", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateInterquery{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateInterqueryResult{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateInterqueryTimeout{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
