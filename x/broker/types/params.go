package types

import (
	fmt "fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	transfertypes "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
)

var (
	DefaultAtomPath = "transfer/channel-0/transfer/channel-0"
	DefaultOsmoPath = "transfer/channel-0"

	AtomPathParamsKey = []byte("AtomIBCPath")
	OsmoPathParamsKey = []byte("OsmoIBCPath")
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(
	atomIbcPath string,
	osmoIbcPath string,
) Params {
	return Params{
		AtomIBCPath: atomIbcPath,
		OsmoIBCPath: osmoIbcPath,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultAtomPath, DefaultOsmoPath,
	)
}

func validate(i interface{}) error {
	_, ok := i.(string)
	if !ok {
		return fmt.Errorf("parameter not accepted: %T", i)
	}
	return nil
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(AtomPathParamsKey, &p.AtomIBCPath, validate),
		paramtypes.NewParamSetPair(OsmoPathParamsKey, &p.OsmoIBCPath, validate),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	return nil
}

// Validate validates the set of params
func (p Params) GetBaseDenomParam() BaseDenoms {
	return BaseDenoms{
		AtomTrace: &transfertypes.DenomTrace{
			Path:      p.AtomIBCPath,
			BaseDenom: "uatom",
		},
		OsmoTrace: &transfertypes.DenomTrace{
			Path:      p.OsmoIBCPath,
			BaseDenom: "uosmo",
		},
	}
}
