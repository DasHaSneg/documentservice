package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ContractList: []Contract{},
		AnnexList:    []Annex{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in contract
	contractIdMap := make(map[uint64]bool)
	contractCount := gs.GetContractCount()
	for _, elem := range gs.ContractList {
		if _, ok := contractIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for contract")
		}
		if elem.Id >= contractCount {
			return fmt.Errorf("contract id should be lower or equal than the last id")
		}
		contractIdMap[elem.Id] = true
	}
	// Check for duplicated ID in annex
	annexIdMap := make(map[uint64]bool)
	annexCount := gs.GetAnnexCount()
	for _, elem := range gs.AnnexList {
		if _, ok := annexIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for annex")
		}
		if elem.Id >= annexCount {
			return fmt.Errorf("annex id should be lower or equal than the last id")
		}
		annexIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
