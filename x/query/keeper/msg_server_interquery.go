package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/defund-labs/defund/x/query/types"
)

func (k msgServer) CreateInterquery(goCtx context.Context, msg *types.MsgCreateInterquery) (*types.MsgCreateInterqueryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Create the store Id by using the Key-Id combination
	storeId := fmt.Sprintf("%s-%s", msg.Name, msg.Id)

	// Check if the value already exists
	_, isFound := k.GetInterquery(
		ctx,
		storeId,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("%s Key to Id is already set. All Key to Id values must be unique.", storeId))
	}

	var interquery = types.Interquery{
		Creator:       msg.Creator,
		Storeid:       storeId,
		Path:          msg.Path,
		Key:           msg.Key,
		TimeoutHeight: msg.TimeoutHeight,
		ClientId:      msg.ClientId,
	}

	k.SetInterquery(
		ctx,
		interquery,
	)
	return &types.MsgCreateInterqueryResponse{}, nil
}

func (k msgServer) CreateInterqueryResult(goCtx context.Context, msg *types.MsgCreateInterqueryResult) (*types.MsgCreateInterqueryResultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetInterqueryResult(
		ctx,
		msg.Storeid,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("%s Key to Id is already set. All Key to Id values must be unique.", msg.Storeid))
	}

	var interqueryresult = types.InterqueryResult{
		Creator:  msg.Creator,
		Storeid:  msg.Storeid,
		Data:     msg.Data,
		Height:   msg.Height,
		ClientId: msg.ClientId,
		Success:  msg.Success,
		Proof:    msg.Proof,
	}

	// Create the interquery result in the store
	k.SetInterqueryResult(ctx, interqueryresult)

	// Remove/cleanup the pending interquery from the store
	k.RemoveInterquery(ctx, interqueryresult.Storeid)

	return &types.MsgCreateInterqueryResultResponse{}, nil
}

func (k msgServer) CreateInterqueryTimeout(goCtx context.Context, msg *types.MsgCreateInterqueryTimeout) (*types.MsgCreateInterqueryTimeoutResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetInterqueryTimeoutResult(
		ctx,
		msg.Storeid,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("%s Key to Id is already set. All Key to Id values must be unique.", msg.Storeid))
	}

	var interquerytimeoutresult = types.InterqueryTimeoutResult{
		Creator:       msg.Creator,
		Storeid:       msg.Storeid,
		TimeoutHeight: msg.TimeoutHeight,
		ClientId:      msg.ClientId,
	}

	k.SetInterqueryTimeoutResult(ctx, interquerytimeoutresult)

	return &types.MsgCreateInterqueryTimeoutResponse{}, nil
}
