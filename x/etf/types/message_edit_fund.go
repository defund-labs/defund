package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgEditFund = "edit_fund"

var _ sdk.Msg = &MsgEditFund{}

func NewMsgEditFund(creator string, symbol string, holdings string) *MsgEditFund {
	return &MsgEditFund{
		Creator:  creator,
		Symbol:   symbol,
		Holdings: holdings,
	}
}

func (msg *MsgEditFund) Route() string {
	return RouterKey
}

func (msg *MsgEditFund) Type() string {
	return TypeMsgEditFund
}

func (msg *MsgEditFund) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgEditFund) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEditFund) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
