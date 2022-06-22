package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddLiquiditySource = "add_liquidity_source"
const TypeMsgAddConnectionBroker = "add_connection_broker"

var _ sdk.Msg = &MsgAddLiquiditySource{}
var _ sdk.Msg = &MsgAddConnectionBroker{}

// helper function to check if a string is within a slice
func contains(list []string, str string) bool {
	for _, value := range list {
		if value == str {
			return true
		}
	}

	return false
}

func NewMsgAddLiquiditySource(creator string, brokerId string, poolId string) *MsgAddLiquiditySource {
	return &MsgAddLiquiditySource{
		Creator:  creator,
		BrokerId: brokerId,
		PoolId:   poolId,
	}
}

func (msg *MsgAddLiquiditySource) Route() string {
	return RouterKey
}

func (msg *MsgAddLiquiditySource) Type() string {
	return TypeMsgAddLiquiditySource
}

func (msg *MsgAddLiquiditySource) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddLiquiditySource) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddLiquiditySource) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func NewMsgAddConnectionBroker(creator string, brokerId string, connectionId string) *MsgAddConnectionBroker {
	return &MsgAddConnectionBroker{
		Creator:      creator,
		BrokerId:     brokerId,
		ConnectionId: connectionId,
	}
}

func (msg *MsgAddConnectionBroker) Route() string {
	return RouterKey
}

func (msg *MsgAddConnectionBroker) Type() string {
	return TypeMsgAddConnectionBroker
}

func (msg *MsgAddConnectionBroker) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddConnectionBroker) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddConnectionBroker) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
