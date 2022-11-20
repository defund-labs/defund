package bindings

import (
	"github.com/defund-labs/defund/x/broker/types"
)

type BrokerQuery struct {
	GetBroker        *types.QueryBrokerRequest                       `json:"get_broker,omitempty"`
	GetBrokers       *types.QueryBrokersRequest                      `json:"get_brokers,omitempty"`
	GetBrokerAddress *types.QueryInterchainAccountFromAddressRequest `json:"get_broker_address,omitempty"`
}
