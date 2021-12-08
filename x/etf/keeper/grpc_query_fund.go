package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/defundhub/defund/x/etf/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) FundAll(c context.Context, req *types.QueryAllFundRequest) (*types.QueryAllFundResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var funds []types.Fund
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	fundStore := prefix.NewStore(store, types.KeyPrefix(types.FundKeyPrefix))

	pageRes, err := query.Paginate(fundStore, req.Pagination, func(key []byte, value []byte) error {
		var fund types.Fund
		if err := k.cdc.Unmarshal(value, &fund); err != nil {
			return err
		}

		funds = append(funds, fund)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllFundResponse{Fund: funds, Pagination: pageRes}, nil
}

func (k Keeper) Fund(c context.Context, req *types.QueryGetFundRequest) (*types.QueryGetFundResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetFund(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetFundResponse{Fund: val}, nil
}
