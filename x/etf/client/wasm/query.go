package wasm

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/defund-labs/defund/x/etf/keeper"
	"github.com/defund-labs/defund/x/etf/types"
)

type EtfWasmQueryHandler struct {
	etfkeeper keeper.Keeper
}

func NewEtfWasmQueryHandler(keeper *keeper.Keeper) *EtfWasmQueryHandler {
	return &EtfWasmQueryHandler{
		etfkeeper: *keeper,
	}
}

func (handler EtfWasmQueryHandler) GetFund(ctx sdk.Context, req *types.QueryGetFundRequest) (*types.QueryGetFundResponse, error) {
	c := sdk.WrapSDKContext(ctx)
	return handler.etfkeeper.Fund(c, req)
}

func (handler EtfWasmQueryHandler) GetFunds(ctx sdk.Context, req *types.QueryAllFundRequest) (*types.QueryAllFundResponse, error) {
	c := sdk.WrapSDKContext(ctx)
	return handler.etfkeeper.FundAll(c, req)
}

func (handler EtfWasmQueryHandler) GetFundPrice(ctx sdk.Context, req *types.QueryFundPriceRequest) (*types.QueryFundPriceResponse, error) {
	c := sdk.WrapSDKContext(ctx)
	return handler.etfkeeper.FundPrice(c, req)
}
