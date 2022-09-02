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
	broker string,
	holdings string,
	rebalance int64,
	basedenom string,
	startingprice string,

) *MsgCreateFund {
	return &MsgCreateFund{
		Creator:       creator,
		Symbol:        symbol,
		Name:          name,
		Description:   description,
		Holdings:      holdings,
		Rebalance:     rebalance,
		BaseDenom:     basedenom,
		StartingPrice: startingprice,
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

	// Ensure that an allowed basedenom is used. Currently we only support uosmo
	if msg.BaseDenom != "uosmo" {
		return sdkerrors.Wrapf(ErrWrongBaseDenom, "invalid base denom (%s)", msg.BaseDenom)
	}

	return nil
}
