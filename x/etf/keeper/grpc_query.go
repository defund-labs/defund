package keeper

import (
	"github.com/defundhub/defund/x/etf/types"
)

var _ types.QueryServer = Keeper{}
