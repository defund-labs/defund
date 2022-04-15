package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	"github.com/defund-labs/defund/x/etf/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
)

func (k msgServer) Invest(goCtx context.Context, msg *types.MsgInvest) (*types.MsgInvestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	fund, found := k.GetFund(ctx, msg.Fund)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrFundNotFound, "failed to find fund with id of %s", fund.Symbol)
	}

	portID, err := icatypes.NewControllerPortID(fund.Address)
	if err != nil {
		return nil, err
	}

	acc, found := k.brokerKeeper.GetBrokerAccount(ctx, fund.ConnectionId, portID)
	if !found {
		return nil, sdkerrors.Wrapf(icatypes.ErrInterchainAccountNotFound, "failed to retrieve interchain account for owner %s", fund.Address)
	}

	id := k.GetNextID(ctx)

	timeoutHeight, err := clienttypes.ParseHeight(msg.TimeoutHeight)
	if err != nil {
		return nil, err
	}

	err = k.Keeper.Invest(ctx, id, acc, fund, msg.Channel, *msg.Amount, msg.Creator, fund.Address, timeoutHeight, msg.TimeoutTimestamp)
	if err != nil {
		return nil, err
	}

	return &types.MsgInvestResponse{}, nil
}
