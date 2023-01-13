package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
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

func (k Keeper) GetBaseDenomParam(ctx sdk.Context) types.BaseDenoms {
	params := k.GetParams(ctx)

	return types.BaseDenoms{
		AtomTrace: &ibctransfertypes.DenomTrace{
			Path:      params.AtomIBCPath,
			BaseDenom: "uatom",
		},
		OsmoTrace: &ibctransfertypes.DenomTrace{
			Path:      params.OsmoIBCPath,
			BaseDenom: "uosmo",
		},
	}
}
