package wasmbinding

import (
	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	brokerwasm "github.com/defund-labs/defund/x/broker/client/wasm"
	brokerkeeper "github.com/defund-labs/defund/x/broker/keeper"
	etfwasm "github.com/defund-labs/defund/x/etf/client/wasm"
	etfkeeper "github.com/defund-labs/defund/x/etf/keeper"
)

func RegisterPlugins(
	etfkeeper *etfkeeper.Keeper,
	brokerkeeper *brokerkeeper.Keeper,
	accountKeeper *authkeeper.AccountKeeper,
	router wasmkeeper.MessageRouter,
) []wasmkeeper.Option {
	etfHandler := etfwasm.NewEtfWasmQueryHandler(etfkeeper)
	brokerHandler := brokerwasm.NewEtfWasmQueryHandler(brokerkeeper)
	wasmQueryPlugin := NewQueryPlugin(brokerHandler, etfHandler)

	queryOpt := wasmkeeper.WithQueryPlugins(&wasmkeeper.QueryPlugins{
		Custom: CustomQuerier(wasmQueryPlugin),
	})
	messengerHandlerOpt := wasmkeeper.WithMessageHandlerDecorator(
		CustomMessageDecorator(router, accountKeeper),
	)

	return []wasm.Option{
		queryOpt,
		messengerHandlerOpt,
	}
}
