package keeper

import (
	"github.com/defundhub/defund/x/query/types"
)

var _ types.QueryServer = Keeper{}
