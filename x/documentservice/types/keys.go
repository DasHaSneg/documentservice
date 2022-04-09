package types

const (
	// ModuleName defines the module name
	ModuleName = "documentservice"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_documentservice"

	ContractKey      = "Contract-value-"
	ContractCountKey = "Contract-count-"

	SupplementKey      = "Supplement-value-"
	SupplementCountKey = "Supplement-count-"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
