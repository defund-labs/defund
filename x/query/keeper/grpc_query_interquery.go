package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/defund-labs/defund/x/query/types"
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
	interqueryStore := prefix.NewStore(store, types.InterqueryKeyPrefix)

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
		req.Storeid,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetInterqueryResponse{Interquery: val}, nil
}

func (k Keeper) InterqueryResultAll(c context.Context, req *types.QueryAllInterqueryResultRequest) (*types.QueryAllInterqueryResultResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var interqueryresults []types.InterqueryResult
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	interqueryResultStore := prefix.NewStore(store, types.InterqueryResultKeyPrefix)

	pageRes, err := query.Paginate(interqueryResultStore, req.Pagination, func(key []byte, value []byte) error {
		var interqueryresult types.InterqueryResult
		if err := k.cdc.Unmarshal(value, &interqueryresult); err != nil {
			return err
		}

		interqueryresults = append(interqueryresults, interqueryresult)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllInterqueryResultResponse{Interqueryresult: interqueryresults, Pagination: pageRes}, nil
}

func (k Keeper) InterqueryResult(c context.Context, req *types.QueryGetInterqueryResultRequest) (*types.QueryGetInterqueryResultResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetInterqueryResult(
		ctx,
		req.Storeid,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetInterqueryResultResponse{Interqueryresult: val}, nil
}

func (k Keeper) InterqueryTimeoutResultAll(c context.Context, req *types.QueryAllInterqueryTimeoutResultRequest) (*types.QueryAllInterqueryTimeoutResultResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var interquerytimeoutresults []types.InterqueryTimeoutResult
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	interqueryTimeoutResultStore := prefix.NewStore(store, types.InterqueryTimeoutResultKeyPrefix)

	pageRes, err := query.Paginate(interqueryTimeoutResultStore, req.Pagination, func(key []byte, value []byte) error {
		var interquerytimeoutresult types.InterqueryTimeoutResult
		if err := k.cdc.Unmarshal(value, &interquerytimeoutresult); err != nil {
			return err
		}

		interquerytimeoutresults = append(interquerytimeoutresults, interquerytimeoutresult)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllInterqueryTimeoutResultResponse{Interquerytimeoutresult: interquerytimeoutresults, Pagination: pageRes}, nil
}

func (k Keeper) InterqueryTimeoutResult(c context.Context, req *types.QueryGetInterqueryTimeoutResultRequest) (*types.QueryGetInterqueryTimeoutResultResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetInterqueryTimeoutResult(
		ctx,
		req.Storeid,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetInterqueryTimeoutResultResponse{Interquerytimeoutresult: val}, nil
}
