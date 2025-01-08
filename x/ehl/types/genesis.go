package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		HashCidList: []HashCid{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in hashCid
	hashCidIdMap := make(map[uint64]bool)
	hashCidCount := gs.GetHashCidCount()
	for _, elem := range gs.HashCidList {
		if _, ok := hashCidIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for hashCid")
		}
		if elem.Id >= hashCidCount {
			return fmt.Errorf("hashCid id should be lower or equal than the last id")
		}
		hashCidIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
