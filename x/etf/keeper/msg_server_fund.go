package keeper

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/defund-labs/defund/x/etf/types"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func NewFundAddress(fundId uint64) sdk.AccAddress {
	key := append([]byte("etf"), sdk.Uint64ToBigEndian(fundId)...)
	return address.Module("etf", key)
}

func GetFundDenom(symbol string) string {
	return fmt.Sprintf("etf/pool/%s", symbol)
}

// Helper function that parses a string of holdings in the format "ATOM:50:1,OSMO:50:2" (DENOM:PERCENT:POOL,...) into a list of holdings
func ParseStringHoldings(holdings string) ([]types.Holding, error) {
	rawHoldings := strings.Split(holdings, ",")
	holdingsList := []types.Holding{}
	for _, holding := range rawHoldings {
		sepHoldings := strings.Split(holding, ":")
		perc, err := strconv.ParseInt(sepHoldings[1], 10, 64)
		if err != nil {
			return nil, err
		}
		holdingsList = append(holdingsList, types.Holding{
			Denom:   sepHoldings[0],
			Percent: perc,
			PoolId:  sepHoldings[2],
		})
	}
	return holdingsList, nil
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

	// Create and save the broker fund ICA account on the broker chain
	err := k.icaKeeper.RegisterInterchainAccount(ctx, msg.ConnectionId, acc.GetAddress().String())
	if err != nil {
		return nil, err
	}

	holdings, err := ParseStringHoldings(msg.Holdings)
	if err != nil {
		return nil, err
	}

	var fund = types.Fund{
		Creator:     msg.Creator,
		Id:          id,
		Address:     acc.GetAddress().String(),
		Symbol:      msg.Symbol,
		Name:        msg.Name,
		Description: msg.Description,
		Shares:      sdk.NewCoin(GetFundDenom(msg.Symbol), sdk.ZeroInt()),
		Broker:      msg.Broker,
		Holdings:    holdings,
		BaseDenom:   msg.BaseDenom,
		Rebalance:   msg.Rebalance,
		ConnectionId: msg.ConnectionId,
	}

	k.SetFund(
		ctx,
		fund,
	)
	return &types.MsgCreateFundResponse{}, nil
}
