package types

// NewQueryInterchainAccountRequest creates and returns a new QueryInterchainAccountFromAddressRequest
func NewQueryInterchainAccountRequest(connectionID, owner string) *QueryInterchainAccountFromAddressRequest {
	return &QueryInterchainAccountFromAddressRequest{
		ConnectionId: connectionID,
		Owner:        owner,
	}
}

// NewQueryInterchainAccountResponse creates and returns a new QueryInterchainAccountFromAddressResponse
func NewQueryInterchainAccountResponse(interchainAccAddr string) *QueryInterchainAccountFromAddressResponse {
	return &QueryInterchainAccountFromAddressResponse{
		InterchainAccountAddress: interchainAccAddr,
	}
}

// NewQueryBrokerRequest creates and returns a new QueryBrokerRequest
func NewQueryBrokerRequest(broker string) *QueryBrokerRequest {
	return &QueryBrokerRequest{
		Broker: broker,
	}
}

// NewQueryBrokerResponse creates and returns a new QueryBrokerResponse
func NewQueryBrokerResponse(broker Broker) *QueryBrokerResponse {
	return &QueryBrokerResponse{
		Broker: &broker,
	}
}
