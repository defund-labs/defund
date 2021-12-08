package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// FundKeyPrefix is the prefix to retrieve all Fund
	FundKeyPrefix = "Fund/value/"
)

// FundKey returns the store key to retrieve a Fund from the index fields
func FundKey(
	id string,
) []byte {
	var key []byte

	idBytes := []byte(id)
	key = append(key, idBytes...)
	key = append(key, []byte("/")...)

	return key
}
