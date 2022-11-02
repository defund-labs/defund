package types

const (
	ModuleName = "broker"

	StoreKey = ModuleName

	RouterKey = ModuleName

	QuerierRoute = ModuleName

	MemStoreKey = "mem_etf"

	// BrokerKeyPrefix is the prefix to retrieve all Broker stores
	BrokerKeyPrefix = "Broker/value/"

	// RedeemKeyPrefix is the prefix to retrieve all Redeem stores
	RedeemKeyPrefix = "Redeem/value/"

	// RebalanceKeyPrefix is the prefix to retrieve all Redeem stores
	RebalanceKeyPrefix = "Rebalance/value/"

	// TransferKeyPrefix is the prefix to retrieve all Transfer stores
	TransferKeyPrefix = "Transfer/value/"

	StatusComplete = "complete"
	StatusError    = "error"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

// BrokerKey returns the store key to retrieve a Broker from the index fields
func BrokerKey(
	id string,
) []byte {
	var key []byte

	idBytes := []byte(id)
	key = append(key, idBytes...)
	key = append(key, []byte("/")...)

	return key
}

// TransferKey returns the store key to retrieve a Transfer from the index fields
func TransferKey(
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

// FundKey returns the store key to retrieve a Fund from the index fields
func RebalanceKey(
	id string,
) []byte {
	var key []byte

	idBytes := []byte(id)
	key = append(key, idBytes...)
	key = append(key, []byte("/")...)

	return key
}
