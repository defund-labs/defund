package types

import (
	"encoding/binary"
)

var _ binary.ByteOrder

var (
	// InterqueryKeyPrefix is the prefix to retrieve all Interquery
	InterqueryKeyPrefix = []byte{0x01}
	// InterqueryResultKeyPrefix is the prefix to retrieve all Interquery Results
	InterqueryResultKeyPrefix = []byte{0x02}
	// InterqueryTimeoutResultKeyPrefix is the prefix to retrieve all Interquery Timeout Results
	InterqueryTimeoutResultKeyPrefix = []byte{0x03}
)

// InterqueryKey returns the store key to retrieve a Interquery from the index fields
func InterqueryKey(
	index string,
) []byte {
	var key []byte

	key = append(key, InterqueryKeyPrefix...)
	key = append(key, []byte("/")...)

	return key
}

func GetKeyPrefixInterquery(storeid string) []byte {
	return append(InterqueryKeyPrefix, []byte(storeid)...)
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

func GetKeyPrefixInterqueryResult(storeid string) []byte {
	return append(InterqueryResultKeyPrefix, []byte(storeid)...)
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

func GetKeyPrefixInterqueryTimeoutResult(storeid string) []byte {
	return append(InterqueryTimeoutResultKeyPrefix, []byte(storeid)...)
}
