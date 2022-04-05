package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/proto/tendermint/crypto"
)

var _ sdk.Msg = &MsgCreateInterquery{}

func NewMsgCreateInterquery(
	creator string,
	id string,
	name string,
	key []byte,
	path string,
	timeoutHeight uint64,
	clientId string,

) *MsgCreateInterquery {
	return &MsgCreateInterquery{
		Creator:       creator,
		Id:            id,
		Name:          name,
		Key:           key,
		Path:          path,
		TimeoutHeight: timeoutHeight,
		ClientId:      clientId,
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

var _ sdk.Msg = &MsgCreateInterqueryResult{}

func NewMsgCreateInterqueryResult(
	creator string,
	storeid string,
	key string,
	data []byte,
	height uint64,
	clientid string,
	success bool,
	proof *crypto.ProofOps,

) *MsgCreateInterqueryResult {
	return &MsgCreateInterqueryResult{
		Creator:  creator,
		Storeid:  storeid,
		Data:     data,
		Height:   height,
		ClientId: clientid,
		Success:  success,
		Proof:    proof,
	}
}

func (msg *MsgCreateInterqueryResult) Route() string {
	return RouterKey
}

func (msg *MsgCreateInterqueryResult) Type() string {
	return "CreateInterqueryResult"
}

func (msg *MsgCreateInterqueryResult) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateInterqueryResult) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateInterqueryResult) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgCreateInterqueryTimeout{}

func NewMsgCreateInterqueryTimeout(
	creator string,
	storeid string,
	key string,
	timeoutheight uint64,
	clientid string,
	proof *crypto.ProofOps,

) *MsgCreateInterqueryTimeout {
	return &MsgCreateInterqueryTimeout{
		Creator:       creator,
		Storeid:       storeid,
		TimeoutHeight: timeoutheight,
		ClientId:      clientid,
	}
}
func (msg *MsgCreateInterqueryTimeout) Route() string {
	return RouterKey
}

func (msg *MsgCreateInterqueryTimeout) Type() string {
	return "CreateInterqueryTimeout"
}

func (msg *MsgCreateInterqueryTimeout) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateInterqueryTimeout) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateInterqueryTimeout) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
