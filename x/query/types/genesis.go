package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		InterqueryList: []Interquery{},
		// this line is used by starport scaffolding # genesis/types/default
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in interquery
	interqueryIndexMap := make(map[string]struct{})
	interqueryResultIndexMap := make(map[string]struct{})
	interqueryTimeoutIndexMap := make(map[string]struct{})

	for _, elem := range gs.InterqueryList {
		index := string(InterqueryKey(elem.Storeid))
		if _, ok := interqueryIndexMap[index]; ok {
			fmt.Print("ok: ", ok, "\n")
			return fmt.Errorf("duplicated index for interquery")
		}
		interqueryIndexMap[index] = struct{}{}
	}
	for _, elem := range gs.InterqueryResultList {
		index := string(InterqueryKey(elem.Storeid))
		if _, ok := interqueryResultIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for interquery result")
		}
		interqueryResultIndexMap[index] = struct{}{}
	}
	for _, elem := range gs.InterqueryTimeoutResultList {
		index := string(InterqueryKey(elem.Storeid))
		if _, ok := interqueryTimeoutIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for interquery timeout result")
		}
		interqueryTimeoutIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return nil
}
