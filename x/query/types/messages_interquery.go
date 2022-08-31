package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	"github.com/tendermint/tendermint/proto/tendermint/crypto"
)

var _ sdk.Msg = &MsgCreateInterquery{}

func NewMsgCreateInterquery(
	creator string,
	storeid string,
	Chainid string,
	key []byte,
	path string,
	timeoutHeight uint64,
	connectionid string,

) *MsgCreateInterquery {
	return &MsgCreateInterquery{
		Creator:       creator,
		Storeid:       storeid,
		Chainid:       Chainid,
		Key:           key,
		Path:          path,
		TimeoutHeight: timeoutHeight,
		ConnectionId:  connectionid,
	}
}

func (msg *MsgCreateInterquery) Route() string {
	return RouterKey
}

func (msg *MsgCreateInterquery) Type() string {
	return "create_interquery"
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
	data []byte,
	height *clienttypes.Height,
	proof *crypto.ProofOps,

) *MsgCreateInterqueryResult {
	return &MsgCreateInterqueryResult{
		Creator: creator,
		Storeid: storeid,
		Data:    data,
		Height:  height,
		Proof:   proof,
	}
}

func (msg *MsgCreateInterqueryResult) Route() string {
	return RouterKey
}

func (msg *MsgCreateInterqueryResult) Type() string {
	return "create_interquery_result"
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
	proof *crypto.ProofOps,

) *MsgCreateInterqueryTimeout {
	return &MsgCreateInterqueryTimeout{
		Creator:       creator,
		Storeid:       storeid,
		TimeoutHeight: timeoutheight,
	}
}
func (msg *MsgCreateInterqueryTimeout) Route() string {
	return RouterKey
}

func (msg *MsgCreateInterqueryTimeout) Type() string {
	return "create_interquery_timeout"
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
