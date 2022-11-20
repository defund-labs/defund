package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/etf module sentinel errors
var (
	ErrWrongBroker          = sdkerrors.Register(ModuleName, 0, "Invalid broker")
	ErrFundNotFound         = sdkerrors.Register(ModuleName, 1, "Fund not found")
	ErrNextSequenceNotFound = sdkerrors.Register(ModuleName, 2, "Next Sequence not found")
	ErrInvalidDenom         = sdkerrors.Register(ModuleName, 3, "Denom is currenly not supported")
	ErrInvalidPool          = sdkerrors.Register(ModuleName, 4, "Pool is currenly not supported")
	ErrSymbolExists         = sdkerrors.Register(ModuleName, 5, "Symbol already exists")
	ErrWrongBaseDenom       = sdkerrors.Register(ModuleName, 6, "Invalid base denom")
	ErrPercentComp          = sdkerrors.Register(ModuleName, 7, "Invalid Percent Composition")
	ErrInvalidPools         = sdkerrors.Register(ModuleName, 8, "No pools found in store")
	ErrMarshallingError     = sdkerrors.Register(ModuleName, 9, "Marshalling error")
	ErrCreateNotFound       = sdkerrors.Register(ModuleName, 10, "Create not found")
	ErrInvalidHolding       = sdkerrors.Register(ModuleName, 11, "Holding is invalid")
	ErrInvalidRebalance     = sdkerrors.Register(ModuleName, 12, "Rebalance failed")
	ErrBeginBlocker         = sdkerrors.Register(ModuleName, 13, "Begin blocker failed")
	ErrInvalidBalance       = sdkerrors.Register(ModuleName, 14, "Error on getting remote account balance")
	ErrUnmarshallJson       = sdkerrors.Register(ModuleName, 15, "JSON unmarshalling error")
	ErrMarshallGetFund      = sdkerrors.Register(ModuleName, 16, "Error marshalling GetFund")
	ErrMarshallGetFunds     = sdkerrors.Register(ModuleName, 17, "Error marshalling GetFunds")
	ErrMarshallGetFundPrice = sdkerrors.Register(ModuleName, 18, "Error marshalling GetFundPrice")
	ErrUnknownEtfQuery      = sdkerrors.Register(ModuleName, 19, "Unknown ETF query")
	ErrParsingWasmMsg       = sdkerrors.Register(ModuleName, 20, "Error parsing wasm msg")
)
