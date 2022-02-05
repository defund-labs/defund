package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/defundhub/defund/x/query/types"
)

func (k msgServer) CreateInterquery(goCtx context.Context, msg *types.MsgCreateInterquery) (*types.MsgCreateInterqueryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetInterquery(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var interquery = types.Interquery{
		Creator:  msg.Creator,
		Index:    msg.Index,
		Height:   msg.Height,
		Path:     msg.Path,
		ChainId:  msg.ChainId,
		TypeName: msg.TypeName,
	}

	k.SetInterquery(
		ctx,
		interquery,
	)
	return &types.MsgCreateInterqueryResponse{}, nil
}

func (k msgServer) UpdateInterquery(goCtx context.Context, msg *types.MsgUpdateInterquery) (*types.MsgUpdateInterqueryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetInterquery(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var interquery = types.Interquery{
		Creator:  msg.Creator,
		Index:    msg.Index,
		Height:   msg.Height,
		Path:     msg.Path,
		ChainId:  msg.ChainId,
		TypeName: msg.TypeName,
	}

	k.SetInterquery(ctx, interquery)

	return &types.MsgUpdateInterqueryResponse{}, nil
}

func (k msgServer) DeleteInterquery(goCtx context.Context, msg *types.MsgDeleteInterquery) (*types.MsgDeleteInterqueryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetInterquery(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveInterquery(
		ctx,
		msg.Index,
	)

	return &types.MsgDeleteInterqueryResponse{}, nil
}
