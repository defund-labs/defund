package keeper

import (
	"github.com/defund-labs/defund/x/query/types"
)

var _ types.QueryServer = Keeper{}
