package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrIBCAccountAlreadyExist = sdkerrors.Register(ModuleName, 1, "interchain account already registered")
	ErrIBCAccountNotExist     = sdkerrors.Register(ModuleName, 2, "interchain account not exist")
	ErrBrokerNotFound         = sdkerrors.Register(ModuleName, 3, "broker could not be found")
	ErrConnectionNotFound     = sdkerrors.Register(ModuleName, 4, "connection does not exist")
	ErrBrokerActive           = sdkerrors.Register(ModuleName, 5, "broker is active")
	ErrInvalidPool            = sdkerrors.Register(ModuleName, 6, "pool query does not exist")
	ErrMarshallingError       = sdkerrors.Register(ModuleName, 7, "Marshalling error")
)
