package keeper

import (
	"github.com/defund-labs/defund/v1/x/query/types"
)

var _ types.QueryServer = Keeper{}
