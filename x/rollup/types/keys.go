package types

const (
	// ModuleName defines the module name
	ModuleName = "rollup"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_rollup"
)

var ParamsKey = []byte("p_rollup")

func KeyPrefix(p string) []byte {
	return []byte(p)
}
