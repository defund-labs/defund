package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateFund{}, "etf/CreateFund", nil)
	cdc.RegisterConcrete(&MsgInvest{}, "etf/Invest", nil)
	cdc.RegisterConcrete(&MsgUninvest{}, "etf/Uninvest", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateFund{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgInvest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUninvest{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
