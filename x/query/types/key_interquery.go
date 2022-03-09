package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// InterqueryKeyPrefix is the prefix to retrieve all Interquery
	InterqueryKeyPrefix = "Interquery/value/"
	// InterqueryResultKeyPrefix is the prefix to retrieve all Interquery Results
	InterqueryResultKeyPrefix = "InterqueryResult/value/"
	// InterqueryKeyPrefix is the prefix to retrieve all Interquery Timeout Results
	InterqueryTimeoutResultKeyPrefix = "InterqueryTimeout/value/"
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

// InterqueryKey returns the store key to retrieve a Interquery Result from the index fields
func InterqueryResultKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}

// InterqueryKey returns the store key to retrieve a Interquery Timeout Result from the index fields
func InterqueryTimeoutResultKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
