package keeper

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/url"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	commitmenttypes "github.com/cosmos/ibc-go/v4/modules/core/23-commitment/types"
	tmclienttypes "github.com/cosmos/ibc-go/v4/modules/light-clients/07-tendermint/types"
	"github.com/defund-labs/defund/x/query/types"
)

func (k msgServer) CreateInterqueryResult(goCtx context.Context, msg *types.MsgCreateInterqueryResult) (*types.MsgCreateInterqueryResultResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)
	var interqueryresult = types.InterqueryResult{}

	// replace the base64 encoded data with bytes
	data, err := base64.StdEncoding.DecodeString(msg.Data)
	if err != nil {
		return nil, err
	}

	// Get the interquery from store
	interquery, isFound := k.GetInterquery(ctx, msg.Storeid)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("Interquery with StoreId %s could not be found.", msg.Storeid))
	}
	// get past interquery result from store if it exists. If it doesnts proceed
	interqueryResult, isFound := k.GetInterqueryResult(ctx, msg.Storeid)

	// if a past interquery exists, if the height of this data is less then the current height, abort
	// as we only accept updated data.
	if isFound {
		if interqueryResult.Height.RevisionHeight > msg.Height.RevisionHeight {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("Interquery results must be a higher block then current result. current interquery height: %d. submitted interquery height: %d", interqueryResult.Height.RevisionHeight, msg.Height.RevisionHeight))
		}
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

	height := msg.Height
	consensusState, found := k.clientKeeper.GetClientConsensusState(ctx, connection.ClientId, height)

	if !found {
		return nil, fmt.Errorf("unable to fetch consensus state for (client id: %s) (msg height: %d) (revision height: %d + revision number: %d)", connection.ClientId, msg.Height, height.GetRevisionHeight(), height.GetRevisionNumber())
	}

	path := commitmenttypes.NewMerklePath([]string{pathList[2], url.PathEscape(string(interquery.Key))}...)

	merkleProof, err := commitmenttypes.ConvertProofs(msg.Proof)
	if err != nil {
		k.Logger(ctx).Error("error converting proofs")
	}

	tmclientstate, ok := clientState.(*tmclienttypes.ClientState)
	if !ok {
		k.Logger(ctx).Error("error unmarshaling client state", "cs", clientState)
	}

	///////////////////////////////////////////////////////////////////////////////

	if len(data) != 0 {
		// if we got a non-nil response, verify inclusion proof.
		if err := merkleProof.VerifyMembership(tmclientstate.ProofSpecs, consensusState.GetRoot(), path, data); err != nil {
			return nil, fmt.Errorf("unable to verify proof: %s", err)
		}
		interqueryresult = types.InterqueryResult{
			Creator:     msg.Creator,
			Storeid:     msg.Storeid,
			Chainid:     interquery.Chainid,
			Data:        data,
			Height:      msg.Height,
			LocalHeight: uint64(ctx.BlockHeight()),
			Success:     true,
			Proved:      true,
		}
		k.Logger(ctx).Debug("interquery result proof validated. Updating query state", "module", types.ModuleName, "queryId", msg.Storeid)

		// Create the interquery result in the store
		k.SetInterqueryResult(ctx, interqueryresult)

		// Run callbacks
		for _, callback := range k.callBacks {
			run := *callback
			err := run(ctx, interqueryResult)
			if err != nil {
				return nil, err
			}
		}
	} else {
		// if we got a nil response, verify non inclusion proof.
		if err := merkleProof.VerifyNonMembership(tmclientstate.ProofSpecs, consensusState.GetRoot(), path); err != nil {
			return nil, fmt.Errorf("unable to verify proof: %s", err)
		}
		interqueryresult = types.InterqueryResult{
			Creator: msg.Creator,
			Storeid: msg.Storeid,
			Data:    data,
			Height:  msg.Height,
			Success: false,
			Proved:  true,
		}
		k.Logger(ctx).Debug("interquery result non-inclusion proof has been validated! Not updating query state", "module", types.ModuleName, "queryId", msg.Storeid)
	}

	// Remove/cleanup the pending interquery from the store
	k.RemoveInterquery(ctx, interqueryresult.Storeid)

	return &types.MsgCreateInterqueryResultResponse{}, nil
}
