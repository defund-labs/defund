package types

import (
	fmt "fmt"

	transfertypes "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
)

var (
	DefaultBrokerParams = BaseDenoms{
		AtomTrace: &transfertypes.DenomTrace{},
		OsmoTrace: &transfertypes.DenomTrace{},
	}

	ParamsKey = []byte{0x51}
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
		"transfer/channel-0/transfer/channel-0", "uatom", "transfer/channel-0", "uosmo",
	)
}

func (p Params) Validate() error {
	if p.BaseDenoms.AtomTrace == nil {
		return fmt.Errorf("atom denom trace cannot be empty")
	}
	if p.BaseDenoms.OsmoTrace == nil {
		return fmt.Errorf("osmo denom trace cannot be empty")
	}
	return nil
}
