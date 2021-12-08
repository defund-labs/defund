package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateFund{}

func NewMsgCreateFund(
	creator string,
	symbol string,
	name string,
	description string,

) *MsgCreateFund {
	return &MsgCreateFund{
		Creator:     creator,
		Symbol:      symbol,
		Name:        name,
		Description: description,
	}
}

func (msg *MsgCreateFund) Route() string {
	return RouterKey
}

func (msg *MsgCreateFund) Type() string {
	return "CreateFund"
}

func (msg *MsgCreateFund) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateFund) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateFund) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateFund{}

func NewMsgUpdateFund(
	name string,
	description string,

) *MsgUpdateFund {
	return &MsgUpdateFund{
		Name:        name,
		Description: description,
	}
}

func (msg *MsgUpdateFund) Route() string {
	return RouterKey
}

func (msg *MsgUpdateFund) Type() string {
	return "UpdateFund"
}

func (msg *MsgUpdateFund) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateFund) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateFund) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
