package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		FundList: []Fund{},
		// this line is used by starport scaffolding # genesis/types/default
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in fund
	fundIndexMap := make(map[string]struct{})

	for _, elem := range gs.FundList {
		id := string(FundKey(elem.Id))
		if _, ok := fundIndexMap[id]; ok {
			return fmt.Errorf("duplicated id for fund")
		}
		fundIndexMap[id] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return nil
}
