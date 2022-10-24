package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/query module sentinel errors
var (
	ErrMarshallingError = sdkerrors.Register(ModuleName, 1101, "marshalling error")
	ErInvalidProof      = sdkerrors.Register(ModuleName, 1102, "invalid proof")
	ErrInvalidQuery     = sdkerrors.Register(ModuleName, 1103, "invalid query")
	ErrFailedCallback   = sdkerrors.Register(ModuleName, 1104, "interquery callback failed")
)
