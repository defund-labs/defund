package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/defund-labs/defund/x/etf/types"
)

func (k msgServer) Uninvest(goCtx context.Context, msg *types.MsgUninvest) (*types.MsgUninvestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUninvestResponse{}, nil
}
