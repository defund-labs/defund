package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/defund-labs/defund/x/broker/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	Keeper
}

// NewMsgServerImpl creates and returns a new types.MsgServer, fulfilling the broker Msg service interface
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

func (k msgServer) AddLiquiditySource(goCtx context.Context, msg *types.MsgAddLiquiditySource) (*types.MsgAddLiquiditySourceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the broker from the store
	broker, found := k.GetBroker(ctx, msg.BrokerId)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrBrokerNotFound, fmt.Sprintf("broker not found: %s", msg.BrokerId))
	}

	// Create pool to be added to brokers list of pools
	addPool := types.Pool{
		PoolId:       msg.PoolId,
		InterqueryId: fmt.Sprintf("%s-%d", msg.BrokerId, msg.PoolId),
	}

	// Append new pool to brokers
	broker.Pools = append(broker.Pools, &addPool)

	k.SetBroker(ctx, broker)

	return &types.MsgAddLiquiditySourceResponse{}, nil
}

func (k msgServer) AddConnectionBroker(goCtx context.Context, msg *types.MsgAddConnectionBroker) (*types.MsgAddConnectionBrokerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	broker, found := k.GetBroker(ctx, msg.BrokerId)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrBrokerNotFound, fmt.Sprintf("broker not found: %s", msg.BrokerId))
	}

	if broker.Status == "active" {
		return nil, sdkerrors.Wrap(types.ErrBrokerActive, fmt.Sprintf("broker %s has connection set", msg.BrokerId))
	}

	_, found = k.connectionKeeper.GetConnection(ctx, msg.ConnectionId)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrConnectionNotFound, fmt.Sprintf("(%s)", msg.ConnectionId))
	}

	broker.ConnectionId = msg.ConnectionId
	broker.Status = "active"

	k.SetBroker(ctx, broker)

	return &types.MsgAddConnectionBrokerResponse{}, nil
}
