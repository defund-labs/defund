package keeper

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	commitmenttypes "github.com/cosmos/ibc-go/v3/modules/core/23-commitment/types"
	tmclienttypes "github.com/cosmos/ibc-go/v3/modules/light-clients/07-tendermint/types"
	"github.com/defund-labs/defund/x/query/types"
)

func (k msgServer) CreateInterquery(goCtx context.Context, msg *types.MsgCreateInterquery) (*types.MsgCreateInterqueryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists in pending interquery
	_, isFound := k.GetInterquery(
		ctx,
		msg.Storeid,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("storeid %s is already set. all store id's must be unique.", msg.Storeid))
	}

	// Check if the value already exists in submitted interquery
	_, isFound = k.GetInterqueryResult(
		ctx,
		msg.Storeid,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("storeid %s is already set. all store id's must be unique.", msg.Storeid))
	}

	// Check if the value already exists in timedout interquery
	_, isFound = k.GetInterqueryTimeoutResult(
		ctx,
		msg.Storeid,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("storeid %s is already set. all store id's must be unique.", msg.Storeid))
	}

	var interquery = types.Interquery{
		Storeid:       msg.Storeid,
		Chainid:       msg.Chainid,
		Path:          msg.Path,
		Key:           msg.Key,
		TimeoutHeight: msg.TimeoutHeight,
		ConnectionId:  msg.ConnectionId,
	}

	k.SetInterquery(
		ctx,
		interquery,
	)
	return &types.MsgCreateInterqueryResponse{}, nil
}

func (k msgServer) CreateInterqueryResult(goCtx context.Context, msg *types.MsgCreateInterqueryResult) (*types.MsgCreateInterqueryResultResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)
	var interqueryresult = types.InterqueryResult{}

	// Get the interquery from store
	interquery, isFound := k.GetInterquery(ctx, msg.Storeid)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("Interquery with StoreId %s could not be found.", msg.Storeid))
	}

	pathList := strings.Split(interquery.Path, "/")

	///////////////////////// Verify Proof Logic //////////////////////////////////

	if msg.Proof == nil {
		return nil, sdkerrors.Wrapf(types.ErInvalidProof, "no proof provided")
	}
	connection, found := k.connectionKeeper.GetConnection(ctx, interquery.ConnectionId)
	if !found {
		return nil, fmt.Errorf("unable to find connection %s", interquery.ConnectionId)
	}

	clientState, found := k.clientKeeper.GetClientState(ctx, connection.ClientId)
	if !found {
		return nil, fmt.Errorf("unable to fetch client state")
	}

	height := clienttypes.NewHeight(clienttypes.ParseChainID(interquery.Chainid), uint64(msg.Height))
	consensusState, found := k.clientKeeper.GetClientConsensusState(ctx, connection.ClientId, height)

	if !found {
		return nil, fmt.Errorf("unable to fetch consensus state for (client id: %s) (msg height: %d) (revision height: %d + revision number: %d)", connection.ClientId, msg.Height, height.RevisionHeight, height.RevisionNumber)
	}

	path := commitmenttypes.NewMerklePath([]string{pathList[1], url.PathEscape(string(interquery.Key))}...)

	merkleProof, err := commitmenttypes.ConvertProofs(msg.Proof)
	if err != nil {
		k.Logger(ctx).Error("error converting proofs")
	}

	tmclientstate, ok := clientState.(*tmclienttypes.ClientState)
	if !ok {
		k.Logger(ctx).Error("error unmarshaling client state", "cs", clientState)
	}

	///////////////////////////////////////////////////////////////////////////////

	if len(msg.Data) != 0 {
		// if we got a non-nil response, verify inclusion proof.
		if err := merkleProof.VerifyMembership(tmclientstate.ProofSpecs, consensusState.GetRoot(), path, msg.Data); err != nil {
			return nil, fmt.Errorf("unable to verify proof: %s", err)
		}
		interqueryresult = types.InterqueryResult{
			Creator: msg.Creator,
			Storeid: msg.Storeid,
			Data:    msg.Data,
			Height:  msg.Height,
			Success: true,
			Proved:  true,
		}
		k.Logger(ctx).Debug("interquery result proof validated", "module", types.ModuleName, "queryId", msg.Storeid)

	} else {
		// if we got a nil response, verify non inclusion proof.
		if err := merkleProof.VerifyNonMembership(tmclientstate.ProofSpecs, consensusState.GetRoot(), path); err != nil {
			return nil, fmt.Errorf("unable to verify proof: %s", err)
		}
		interqueryresult = types.InterqueryResult{
			Creator: msg.Creator,
			Storeid: msg.Storeid,
			Data:    msg.Data,
			Height:  msg.Height,
			Success: false,
			Proved:  true,
		}
		k.Logger(ctx).Debug("interquery result non-inclusion proof has been validated!", "module", types.ModuleName, "queryId", msg.Storeid)
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
		Storeid:       msg.Storeid,
		TimeoutHeight: msg.TimeoutHeight,
	}

	k.SetInterqueryTimeoutResult(ctx, interquerytimeoutresult)

	// Remove/cleanup the pending interquery from the store
	k.RemoveInterquery(ctx, interquerytimeoutresult.Storeid)

	return &types.MsgCreateInterqueryTimeoutResponse{}, nil
}
