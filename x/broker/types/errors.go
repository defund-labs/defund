package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrIBCAccountAlreadyExist = sdkerrors.Register(ModuleName, 1, "interchain account already registered")
	ErrIBCAccountNotExist     = sdkerrors.Register(ModuleName, 2, "interchain account does not exist")
	ErrBrokerNotFound         = sdkerrors.Register(ModuleName, 3, "broker could not be found")
	ErrConnectionNotFound     = sdkerrors.Register(ModuleName, 4, "connection does not exist")
	ErrBrokerActive           = sdkerrors.Register(ModuleName, 5, "broker is active")
	ErrInvalidPool            = sdkerrors.Register(ModuleName, 6, "pool query does not exist")
	ErrMarshallingError       = sdkerrors.Register(ModuleName, 7, "Marshalling error")
	ErrNextSequenceNotFound   = sdkerrors.Register(ModuleName, 8, "Next Sequence not found")
	ErrNotPositiveWeight      = sdkerrors.Register(ModuleName, 9, "token weight should be greater than 0")
	ErrWeightTooLarge         = sdkerrors.Register(ModuleName, 10, "user specified token weight should be less than 2^20")
	ErrDenomNotFoundInPool    = sdkerrors.Register(ModuleName, 11, "denom does not exist in pool")
	ErrInvalidAccount         = sdkerrors.Register(ModuleName, 12, "error with remote account balance query")
)
