package types

const (
	// ModuleName defines the module name
	ModuleName = "etf"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_etf"

	// FundKeyPrefix is the prefix to retrieve all Fund stores
	FundKeyPrefix = "Fund/value/"

	// InvestKeyPrefix is the prefix to retrieve all Invest stores
	InvestKeyPrefix = "Invest/value/"

	// UninvestKeyPrefix is the prefix to retrieve all Uninvest stores
	UninvestKeyPrefix = "Uninvest/value/"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

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

// InvestKey returns the store key to retrieve a Invest from the index fields
func InvestKey(
	id string,
) []byte {
	var key []byte

	idBytes := []byte(id)
	key = append(key, idBytes...)
	key = append(key, []byte("/")...)

	return key
}

// UninvestKey returns the store key to retrieve a Uninvest from the index fields
func UninvestKey(
	id string,
) []byte {
	var key []byte

	idBytes := []byte(id)
	key = append(key, idBytes...)
	key = append(key, []byte("/")...)

	return key
}
