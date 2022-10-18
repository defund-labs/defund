package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/defund-labs/defund/x/broker/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramstore.GetParamSet(ctx, &params)
	return params
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

func (k *Keeper) GetParam(ctx sdk.Context, key []byte) *types.BaseDenoms {
	var bd types.BaseDenoms
	k.paramstore.Get(ctx, key, &bd)
	return &bd
}
