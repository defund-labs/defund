package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrIBCAccountAlreadyExist = sdkerrors.Register(ModuleName, 1, "interchain account already registered")
	ErrIBCAccountNotExist     = sdkerrors.Register(ModuleName, 2, "interchain account not exist")
	ErrHandlingICAMsg         = sdkerrors.Register(ModuleName, 3, "error handling interchain account message")
)
