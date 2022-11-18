package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	"github.com/tendermint/tendermint/proto/tendermint/crypto"
)

var _ sdk.Msg = &MsgCreateInterqueryResult{}

func NewMsgCreateInterqueryResult(
	creator string,
	storeid string,
	data string,
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
