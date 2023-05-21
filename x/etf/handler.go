package etf

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/defund-labs/defund/x/etf/keeper"
	"github.com/defund-labs/defund/x/etf/types"
)

// NewHandler returns a new SDK handler for the ETF module.
func NewHandler(k keeper.Keeper) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgCreateFund:
			return handleMsgCreateFund(ctx, msgServer, msg)
		case *types.MsgCreate:
			return handleMsgCreate(ctx, msgServer, msg)
		case *types.MsgRedeem:
			return handleMsgRedeem(ctx, msgServer, msg)
		case *types.MsgEditFund:
			return handleMsgEditFund(ctx, msgServer, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

func handleMsgCreateFund(ctx sdk.Context, msgServer keeper.MsgServer, msg *types.MsgCreateFund) (*sdk.Result, error) {
	res, err := msgServer.CreateFund(sdk.WrapSDKContext(ctx), msg)
	return sdk.WrapServiceResult(ctx, res, err)
}

func handleMsgCreate(ctx sdk.Context, msgServer keeper.MsgServer, msg *types.MsgCreate) (*sdk.Result, error) {
	res, err := msgServer.Create(sdk.WrapSDKContext(ctx), msg)
	return sdk.WrapServiceResult(ctx, res, err)
}

func handleMsgRedeem(ctx sdk.Context, msgServer keeper.MsgServer, msg *types.MsgRedeem) (*sdk.Result, error) {
	res, err := msgServer.Redeem(sdk.WrapSDKContext(ctx), msg)
	return sdk.WrapServiceResult(ctx, res, err)
}

func handleMsgEditFund(ctx sdk.Context, msgServer keeper.MsgServer, msg *types.MsgEditFund) (*sdk.Result, error) {
	res, err := msgServer.EditFund(sdk.WrapSDKContext(ctx), msg)
	return sdk.WrapServiceResult(ctx, res, err)
}
