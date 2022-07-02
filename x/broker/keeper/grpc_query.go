package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"
	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"

	"github.com/defund-labs/defund/x/broker/types"
)

// InterchainAccountFromAddress implements the Query/InterchainAccountFromAddress gRPC method
func (k Keeper) InterchainAccountFromAddress(goCtx context.Context, req *types.QueryInterchainAccountFromAddressRequest) (*types.QueryInterchainAccountFromAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	portID, err := icatypes.NewControllerPortID(req.Owner)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "could not find account: %s", err)
	}

	addr, found := k.icaControllerKeeper.GetInterchainAccountAddress(ctx, req.ConnectionId, portID)
	if !found {
		return nil, status.Errorf(codes.NotFound, "no account found for portID %s", portID)
	}

	return types.NewQueryInterchainAccountResponse(addr), nil
}

// Broker implements the Query/Broker gRPC method
func (k Keeper) Broker(goCtx context.Context, req *types.QueryBrokerRequest) (*types.QueryBrokerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	broker, found := k.GetBroker(ctx, req.Broker)
	if !found {
		return nil, status.Errorf(codes.NotFound, "broker %s not found", req.Broker)
	}

	return types.NewQueryBrokerResponse(broker), nil
}
