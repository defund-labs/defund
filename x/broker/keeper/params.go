package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/defund-labs/defund/x/broker/types"
)

// Gets the base denom param directly from params
func (k Keeper) GetParam(ctx sdk.Context, key []byte) (bd *types.BaseDenoms) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.ParamsKey)
	if bz == nil {
		return bd
	}

	params := types.Params{}
	k.cdc.MustUnmarshal(bz, &params)

	bd = params.BaseDenoms

	return bd
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) error {
	err := params.Validate()
	if err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	bz, err := k.cdc.Marshal(&params)
	if err != nil {
		return err
	}
	store.Set(types.ParamsKey, bz)

	return nil
}

func (k *Keeper) GetParams(ctx sdk.Context, key []byte) (params types.Params) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.ParamsKey)
	if bz == nil {
		return params
	}

	k.cdc.MustUnmarshal(bz, &params)
	return params
}
