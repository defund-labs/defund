package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// InterqueryKeyPrefix is the prefix to retrieve all Interquery
	InterqueryKeyPrefix = "Interquery/value/"
)

// InterqueryKey returns the store key to retrieve a Interquery from the index fields
func InterqueryKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
