package types

import (
	fmt "fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	transfertypes "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
)

var (
	DefaultBrokerParams = BaseDenoms{
		AtomTrace: &transfertypes.DenomTrace{},
		OsmoTrace: &transfertypes.DenomTrace{},
	}

	KeyBaseDenoms = []byte("BaseDenoms")
)

func NewParams(
	atompath string,
	atombasedenom string,
	osmopath string,
	osmobasedenom string,
) Params {
	bd := BaseDenoms{
		AtomTrace: &transfertypes.DenomTrace{
			Path:      atompath,
			BaseDenom: atombasedenom,
		},
		OsmoTrace: &transfertypes.DenomTrace{
			Path:      osmopath,
			BaseDenom: osmobasedenom,
		},
	}
	return Params{
		BaseDenoms: &bd,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		"", "uatom", "", "uosmo",
	)
}

func validateBaseDenom(i interface{}) error {
	denoms, ok := i.(BaseDenoms)
	if !ok {
		return fmt.Errorf("param is not accepted: %T", i)
	}
	if denoms.AtomTrace != nil {
		return fmt.Errorf("atom denom trace cannot be empty")
	}
	if denoms.OsmoTrace != nil {
		return fmt.Errorf("osmo denom trace cannot be empty")
	}
	return nil
}

func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyBaseDenoms, &p.BaseDenoms, validateBaseDenom),
	}
}
