package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/etf module sentinel errors
var (
	ErrWrongBroker          = sdkerrors.Register(ModuleName, 1100, "Invalid broker")
	ErrFundNotFound         = sdkerrors.Register(ModuleName, 1101, "Fund not found")
	ErrNextSequenceNotFound = sdkerrors.Register(ModuleName, 1102, "Next Sequence not found")
	ErrInvalidDenom         = sdkerrors.Register(ModuleName, 1103, "Denom is currenly not supported")
	ErrInvalidPool          = sdkerrors.Register(ModuleName, 1104, "Pool is currenly not supported")
	ErrSymbolExists         = sdkerrors.Register(ModuleName, 1105, "Symbol already exists")
	ErrWrongBaseDenom       = sdkerrors.Register(ModuleName, 1106, "Invalid base denom")
)
