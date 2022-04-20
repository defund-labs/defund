package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
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
	}

	return &types.QueryFundPriceResponse{Price: fundPrice}, nil
}
