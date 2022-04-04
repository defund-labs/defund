package keeper

import (
	"context"
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/defund-labs/defund/v1/x/etf/types"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func NewFundAddress(fundId uint64) sdk.AccAddress {
	key := append([]byte("etf"), sdk.Uint64ToBigEndian(fundId)...)
	return address.Module("etf", key)
}

func GetFundDenom(symbol string) string {
	return fmt.Sprintf("etf/pool/%s", symbol)
}

func (k msgServer) CreateFund(goCtx context.Context, msg *types.MsgCreateFund) (*types.MsgCreateFundResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get next id for new fund
	id := k.GetNextId(
		ctx,
	)

	// Check if the value already exists
	_, isFound := k.GetFund(
		ctx,
		id,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "id already exists")
	}

	idInt, _ := strconv.ParseUint(id, 10, 64)

	// Generate and get a new fund address
	fundAddress := NewFundAddress(idInt)

	// Create and save corresponding module account to the account keeper
	acc := k.accountKeeper.NewAccount(ctx, authtypes.NewModuleAccount(
		authtypes.NewBaseAccountWithAddress(
			fundAddress,
		),
		fundAddress.String(),
		"mint",
		"burn",
	))
	k.accountKeeper.SetAccount(ctx, acc)

	var fund = types.Fund{
		Creator:     msg.Creator,
		Id:          id,
		Address:     fundAddress.String(),
		Symbol:      msg.Symbol,
		Name:        msg.Name,
		Description: msg.Description,
		Shares:      sdk.NewCoin(GetFundDenom(msg.Symbol), sdk.ZeroInt()),
	}

	k.SetFund(
		ctx,
		fund,
	)
	return &types.MsgCreateFundResponse{}, nil
}

func (k msgServer) UpdateFund(goCtx context.Context, msg *types.MsgUpdateFund) (*types.MsgUpdateFundResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetFund(
		ctx,
		msg.Id,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "id of fund not found")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var fund = types.Fund{
		Creator:     msg.Creator,
		Id:          msg.Id,
		Address:     valFound.Address,
		Symbol:      valFound.Symbol,
		Name:        msg.Name,
		Description: msg.Description,
		Shares:      valFound.Shares,
	}

	k.SetFund(ctx, fund)

	return &types.MsgUpdateFundResponse{}, nil
}
