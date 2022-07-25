package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateFund{}, "etf/CreateFund", nil)
	cdc.RegisterConcrete(&MsgCreate{}, "etf/Create", nil)
	cdc.RegisterConcrete(&MsgRedeem{}, "etf/Redeem", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateFund{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreate{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRedeem{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
