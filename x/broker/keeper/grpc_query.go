package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	icatypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/types"

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

// Broker implements the Query/Broker gRPC method
func (k Keeper) Brokers(goCtx context.Context, req *types.QueryBrokersRequest) (*types.QueryBrokersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var brokers []types.Broker
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	brokerStore := prefix.NewStore(store, []byte(types.BrokerKeyPrefix))

	pageRes, err := query.Paginate(brokerStore, req.Pagination, func(key []byte, value []byte) error {
		var broker types.Broker
		if err := k.cdc.Unmarshal(value, &broker); err != nil {
			return err
		}

		brokers = append(brokers, broker)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryBrokersResponse{Brokers: brokers, Pagination: pageRes}, nil
}

// Create implements the Query/Broker gRPC method
func (k Keeper) Create(goCtx context.Context, req *types.QueryCreateRequest) (*types.QueryCreateResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var creates []types.Create
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	createStore := prefix.NewStore(store, []byte(types.TransferKeyPrefix))

	pageRes, err := query.Paginate(createStore, req.Pagination, func(key []byte, value []byte) error {
		var create types.Create
		if err := k.cdc.Unmarshal(value, &create); err != nil {
			return err
		}

		creates = append(creates, create)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryCreateResponse{Creates: creates, Pagination: pageRes}, nil
}

// Redeem implements the Query/Broker gRPC method
func (k Keeper) Redeem(goCtx context.Context, req *types.QueryRedeemRequest) (*types.QueryRedeemResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var redeems []types.Redeem
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	redeemStore := prefix.NewStore(store, []byte(types.RedeemKeyPrefix))

	pageRes, err := query.Paginate(redeemStore, req.Pagination, func(key []byte, value []byte) error {
		var redeem types.Redeem
		if err := k.cdc.Unmarshal(value, &redeem); err != nil {
			return err
		}

		redeems = append(redeems, redeem)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryRedeemResponse{Redeems: redeems, Pagination: pageRes}, nil
}
