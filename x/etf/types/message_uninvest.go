package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUninvest{}

func NewMsgUninvest(creator string, fund string, amount *sdk.Coin) *MsgUninvest {
	return &MsgUninvest{
		Creator: creator,
		Fund:    fund,
		Amount:  amount,
	}
}

func (msg *MsgUninvest) Route() string {
	return RouterKey
}

func (msg *MsgUninvest) Type() string {
	return "Uninvest"
}

func (msg *MsgUninvest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUninvest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUninvest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
