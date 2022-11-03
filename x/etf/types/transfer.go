package types

import (
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

type RedeemMsgs struct {
	Osmosis []*banktypes.MsgSend
}

const (
	TransferStateTransferring = "tranferring"
	TransferStateFailed       = "failed"
	TransferStateCompleted    = "complete"
	TransferStateBurned       = "burned"
	RedeemState               = "redeeming"
)
