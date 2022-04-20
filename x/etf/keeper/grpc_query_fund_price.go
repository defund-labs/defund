package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/defund-labs/defund/x/etf/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) FundPrice(goCtx context.Context, req *types.QueryFundPriceRequest) (*types.QueryFundPriceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	price, err := k.CreateFundPrice(ctx, req.Symbol)
	if err != nil {
		return &types.QueryFundPriceResponse{}, err
	}

	fundPrice := types.FundPrice{
		Height: uint64(ctx.BlockHeight()),
		Price:  price,
		Symbol: req.Symbol,
	}

	return &types.QueryFundPriceResponse{Price: fundPrice}, nil
}

func (k Keeper) FundPriceAll(c context.Context, req *types.QueryAllFundPriceRequest) (*types.QueryAllFundPriceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var fundprices []types.FundPrice
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	fundStore := prefix.NewStore(store, types.KeyPrefix(types.FundPriceKeyPrefix))

	pageRes, err := query.Paginate(fundStore, req.Pagination, func(key []byte, value []byte) error {
		var fundprice types.FundPrice
		if err := k.cdc.Unmarshal(value, &fundprice); err != nil {
			return err
		}

		fundprices = append(fundprices, fundprice)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllFundPriceResponse{Price: fundprices, Pagination: pageRes}, nil
}
