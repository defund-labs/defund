package wasmbinding

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	brokerwasm "github.com/defund-labs/defund/x/broker/client/wasm"
	brokerbindings "github.com/defund-labs/defund/x/broker/client/wasm/bindings"
	brokertypes "github.com/defund-labs/defund/x/broker/types"
	etfwasm "github.com/defund-labs/defund/x/etf/client/wasm"
	etfbindings "github.com/defund-labs/defund/x/etf/client/wasm/bindings"
	etftypes "github.com/defund-labs/defund/x/etf/types"
)

type QueryPlugin struct {
	brokerHandler brokerwasm.BrokerWasmQueryHandler
	etfHandler    etfwasm.EtfWasmQueryHandler
}

// NewQueryPlugin returns a reference to a new QueryPlugin.
func NewQueryPlugin(brokerquery *brokerwasm.BrokerWasmQueryHandler, etfquery *etfwasm.EtfWasmQueryHandler) *QueryPlugin {
	return &QueryPlugin{
		brokerHandler: *brokerquery,
		etfHandler:    *etfquery,
	}
}

func (qp QueryPlugin) HandleBrokerQuery(ctx sdk.Context, queryData json.RawMessage) ([]byte, error) {
	var query brokerbindings.BrokerQuery
	if err := json.Unmarshal(queryData, &query); err != nil {
		return nil, brokertypes.ErrUnmarshallJson
	}
	switch {
	case query.GetBroker != nil:
		res, err := qp.brokerHandler.GetBroker(ctx, query.GetBroker)
		if err != nil {
			return nil, err
		}
		bz, err := json.Marshal(res)
		if err != nil {
			return nil, brokertypes.ErrEncodingBroker
		}

		return bz, nil
	case query.GetBrokers != nil:
		res, err := qp.brokerHandler.GetBrokers(ctx, query.GetBrokers)
		if err != nil {
			return nil, err
		}
		bz, err := json.Marshal(res)
		if err != nil {
			return nil, brokertypes.ErrEncodingBrokers
		}

		return bz, nil
	case query.GetBrokerAddress != nil:
		res, err := qp.brokerHandler.GetBrokerAddress(ctx, query.GetBrokerAddress)
		if err != nil {
			return nil, err
		}
		bz, err := json.Marshal(res)
		if err != nil {
			return nil, brokertypes.ErrEncodingBrokerAddress
		}

		return bz, nil
	default:
		return nil, brokertypes.ErrUnknownQuery
	}
}

func (qp QueryPlugin) HandleEtfQuery(ctx sdk.Context, queryData json.RawMessage) ([]byte, error) {
	var query etfbindings.EtfQuery
	if err := json.Unmarshal(queryData, &query); err != nil {
		return nil, etftypes.ErrUnmarshallJson
	}
	switch {
	case query.GetFund != nil:
		res, err := qp.etfHandler.GetFund(ctx, query.GetFund)
		if err != nil {
			return nil, err
		}
		bz, err := json.Marshal(res)
		if err != nil {
			return nil, etftypes.ErrMarshallGetFund
		}

		return bz, nil
	case query.GetFunds != nil:
		res, err := qp.etfHandler.GetFunds(ctx, query.GetFunds)
		if err != nil {
			return nil, err
		}
		bz, err := json.Marshal(res)
		if err != nil {
			return nil, etftypes.ErrMarshallGetFunds
		}

		return bz, nil
	case query.GetFundPrice != nil:
		res, err := qp.etfHandler.GetFundPrice(ctx, query.GetFundPrice)
		if err != nil {
			return nil, err
		}
		bz, err := json.Marshal(res)
		if err != nil {
			return nil, etftypes.ErrMarshallGetFundPrice
		}

		return bz, nil
	default:
		return nil, etftypes.ErrUnknownEtfQuery
	}
}
