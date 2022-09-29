package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRedeem{}

func NewMsgRedeem(creator string, fund string, amount *sdk.Coin, channel string, timeoutheight string, timeouttimestamp uint64, addresses AddressMap) *MsgRedeem {
	return &MsgRedeem{
		Creator:   creator,
		Fund:      fund,
		Amount:    amount,
		Channel:   channel,
		Addresses: &addresses,
	}
}

func (msg *MsgRedeem) Route() string {
	return RouterKey
}

func (msg *MsgRedeem) Type() string {
	return "Redeem"
}

func (msg *MsgRedeem) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRedeem) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRedeem) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
