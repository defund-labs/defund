package bindings

import (
	"github.com/defund-labs/defund/x/etf/types"
)

type EtfQuery struct {
	GetFund      *types.QueryGetFundRequest   `json:"get_fund,omitempty"`
	GetFunds     *types.QueryAllFundRequest   `json:"get_all_funds,omitempty"`
	GetFundPrice *types.QueryFundPriceRequest `json:"get_fund_price,omitempty"`
}
