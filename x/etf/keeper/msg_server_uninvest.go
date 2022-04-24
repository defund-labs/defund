package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	"github.com/defund-labs/defund/x/etf/types"
)

func (k msgServer) Uninvest(goCtx context.Context, msg *types.MsgUninvest) (*types.MsgUninvestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Grab the fund
	fund, err := k.Keeper.GetFundBySymbol(ctx, msg.Fund)
	if err != nil {
		return nil, err
	}

	portID, err := icatypes.NewControllerPortID(fund.Address)
	if err != nil {
		return nil, err
	}

	// Get the next uninvest id
	id := k.Keeper.GetNextIDInvest(ctx)
	// Get the next sequence
	sequence, found := k.channelKeeper.GetNextSequenceSend(ctx, portID, msg.Channel)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrNextSequenceNotFound, "failed to retrieve the next sequence for channel %s and port %s", msg.Channel, portID)
	}

	uninvest := types.Uninvest{
		Id:       id,
		Creator:  msg.Creator,
		Fund:     &fund,
		Amount:   msg.Amount,
		Channel:  msg.Channel,
		Sequence: sequence,
		Status:   "pending",
		Type:     uint64(0),
		Error:    "",
	}

	// Set uninvest store. This store triggers uninvest logic via the UninvestBeginBlocker
	k.Keeper.SetUninvest(ctx, uninvest)

	return &types.MsgUninvestResponse{}, nil
}
