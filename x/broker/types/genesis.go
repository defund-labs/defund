package types

import (
	"fmt"
)

// DefaultIndex is the default broker global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Broker genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Brokers: []Broker{},
		Params:  DefaultParams(),
		// this line is used by starport scaffolding # genesis/types/default
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	brokerIndexMap := make(map[string]struct{})

	for _, elem := range gs.Brokers {
		id := string(BrokerKey(elem.Id))
		// Check for duplicated index in broker
		if _, ok := brokerIndexMap[id]; ok {
			return fmt.Errorf("duplicated id for broker")
		}
		brokerIndexMap[id] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return nil
}
