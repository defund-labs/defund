package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/query module sentinel errors
var (
	ErrInvalidDenom       = sdkerrors.Register(ModuleName, 1100, "Denom is currenly not supported")
	ErrInvalidPool        = sdkerrors.Register(ModuleName, 1101, "Pool is currenly not supported")
	ErrInvalidPools       = sdkerrors.Register(ModuleName, 1102, "No pools found in store")
	ErrMarshallingError   = sdkerrors.Register(ModuleName, 1103, "Marshalling error")
	ErrPercentComp        = sdkerrors.Register(ModuleName, 1104, "Invalid Percent Composition")
	ErrInterqueryNotFound = sdkerrors.Register(ModuleName, 1105, "Interquery not found")
)
