package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/defundhub/defund/x/query/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) InterqueryAll(c context.Context, req *types.QueryAllInterqueryRequest) (*types.QueryAllInterqueryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var interquerys []types.Interquery
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	interqueryStore := prefix.NewStore(store, types.KeyPrefix(types.InterqueryKeyPrefix))

	pageRes, err := query.Paginate(interqueryStore, req.Pagination, func(key []byte, value []byte) error {
		var interquery types.Interquery
		if err := k.cdc.Unmarshal(value, &interquery); err != nil {
			return err
		}

		interquerys = append(interquerys, interquery)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllInterqueryResponse{Interquery: interquerys, Pagination: pageRes}, nil
}

func (k Keeper) Interquery(c context.Context, req *types.QueryGetInterqueryRequest) (*types.QueryGetInterqueryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetInterquery(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetInterqueryResponse{Interquery: val}, nil
}
