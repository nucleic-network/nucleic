package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// RollupKeyPrefix is the prefix to retrieve all Rollup
	RollupKeyPrefix         = "Rollup/value/"
	RollupByEIP155KeyPrefix = "RollupByEIP155/value/"
)

// RollupKey returns the store key to retrieve a Rollup from the index fields
func RollupKey(
	rollupId string,
) []byte {
	var key []byte

	rollupIdBytes := []byte(rollupId)
	key = append(key, rollupIdBytes...)
	key = append(key, []byte("/")...)

	return key
}

// RollupByEIP155Key returns the store key to retrieve a Rollup from the index fields
func RollupByEIP155Key(
	eip155 uint64,
) []byte {
	var key []byte

	eip155Bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(eip155Bytes, eip155)
	key = append(key, eip155Bytes...)
	key = append(key, []byte("/")...)

	return key
}
