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

	// FundKeyPrefix is the prefix to retrieve all Fund stores
	FundPriceKeyPrefix = "FundPrice/value/"

	// CreateKeyPrefix is the prefix to retrieve all Create stores
	CreateKeyPrefix = "Create/value/"

	// RedeemKeyPrefix is the prefix to retrieve all Redeem stores
	RedeemKeyPrefix = "Redeem/value/"
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

// FundPriceKey returns the store key to retrieve a FundPrice from the index fields
func FundPriceKey(
	id string,
) []byte {
	var key []byte

	idBytes := []byte(id)
	key = append(key, idBytes...)
	key = append(key, []byte("/")...)

	return key
}

// CreateKey returns the store key to retrieve a Create from the index fields
func CreateKey(
	id string,
) []byte {
	var key []byte

	idBytes := []byte(id)
	key = append(key, idBytes...)
	key = append(key, []byte("/")...)

	return key
}

// RedeemKey returns the store key to retrieve a Redeem from the index fields
func RedeemKey(
	id string,
) []byte {
	var key []byte

	idBytes := []byte(id)
	key = append(key, idBytes...)
	key = append(key, []byte("/")...)

	return key
}
