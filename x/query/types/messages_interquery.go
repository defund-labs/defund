package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateInterquery{}

func NewMsgCreateInterquery(
	creator string,
	index string,
	height string,
	path string,
	chainId string,
	typeName string,

) *MsgCreateInterquery {
	return &MsgCreateInterquery{
		Creator:  creator,
		Index:    index,
		Height:   height,
		Path:     path,
		ChainId:  chainId,
		TypeName: typeName,
	}
}

func (msg *MsgCreateInterquery) Route() string {
	return RouterKey
}

func (msg *MsgCreateInterquery) Type() string {
	return "CreateInterquery"
}

func (msg *MsgCreateInterquery) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateInterquery) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateInterquery) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateInterquery{}

func NewMsgUpdateInterquery(
	creator string,
	index string,
	height string,
	path string,
	chainId string,
	typeName string,

) *MsgUpdateInterquery {
	return &MsgUpdateInterquery{
		Creator:  creator,
		Index:    index,
		Height:   height,
		Path:     path,
		ChainId:  chainId,
		TypeName: typeName,
	}
}

func (msg *MsgUpdateInterquery) Route() string {
	return RouterKey
}

func (msg *MsgUpdateInterquery) Type() string {
	return "UpdateInterquery"
}

func (msg *MsgUpdateInterquery) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateInterquery) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateInterquery) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteInterquery{}

func NewMsgDeleteInterquery(
	creator string,
	index string,

) *MsgDeleteInterquery {
	return &MsgDeleteInterquery{
		Creator: creator,
		Index:   index,
	}
}
func (msg *MsgDeleteInterquery) Route() string {
	return RouterKey
}

func (msg *MsgDeleteInterquery) Type() string {
	return "DeleteInterquery"
}

func (msg *MsgDeleteInterquery) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteInterquery) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteInterquery) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
