package keeper

import (
	"github.com/defund-labs/defund/v1/x/etf/types"
)

var _ types.QueryServer = Keeper{}
