package keeper

import (
	"github.com/defund-labs/defund/x/etf/types"
)

var _ types.QueryServer = Keeper{}
