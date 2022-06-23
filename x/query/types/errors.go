package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/query module sentinel errors
var (
	ErrMarshallingError = sdkerrors.Register(ModuleName, 1101, "Marshalling error")
)
