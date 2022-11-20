package wasm

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/defund-labs/defund/x/broker/keeper"
	"github.com/defund-labs/defund/x/broker/types"
)

type BrokerWasmQueryHandler struct {
	brokerkeeper keeper.Keeper
}

func NewEtfWasmQueryHandler(keeper *keeper.Keeper) *BrokerWasmQueryHandler {
	return &BrokerWasmQueryHandler{
		brokerkeeper: *keeper,
	}
}

func (handler BrokerWasmQueryHandler) GetBroker(ctx sdk.Context, req *types.QueryBrokerRequest) (*types.QueryBrokerResponse, error) {
	c := sdk.WrapSDKContext(ctx)
	return handler.brokerkeeper.Broker(c, req)
}

func (handler BrokerWasmQueryHandler) GetBrokers(ctx sdk.Context, req *types.QueryBrokersRequest) (*types.QueryBrokersResponse, error) {
	c := sdk.WrapSDKContext(ctx)
	return handler.brokerkeeper.Brokers(c, req)
}

func (handler BrokerWasmQueryHandler) GetBrokerAddress(ctx sdk.Context, req *types.QueryInterchainAccountFromAddressRequest) (*types.QueryInterchainAccountFromAddressResponse, error) {
	c := sdk.WrapSDKContext(ctx)
	return handler.brokerkeeper.InterchainAccountFromAddress(c, req)
}
