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

func NewFundAddress(fundId string) sdk.AccAddress {
	key := append([]byte("etf"), []byte(fundId)...)
	return address.Module("etf", key)
}

func GetFundDenom(symbol string) string {
	return fmt.Sprintf("etf/pool/%s", symbol)
}

// Helper function that parses a string of holdings in the format "ATOM:50:1,OSMO:50:2" (DENOM:PERCENT:POOL,...) into a slice of type holding
// and checks to make sure that the holdings are all supported denoms from the specified broker and pool
func (k msgServer) ParseStringHoldings(ctx sdk.Context, broker string, holdings string) ([]types.Holding, error) {
	rawHoldings := strings.Split(holdings, ",")
	holdingsList := []types.Holding{}
	for _, holding := range rawHoldings {
		sepHoldings := strings.Split(holding, ":")
		perc, err := strconv.ParseInt(sepHoldings[1], 10, 64)
		if err != nil {
			return nil, err
		}
		holdingsList = append(holdingsList, types.Holding{
			Token:   sepHoldings[0],
			Percent: perc,
			PoolId:  sepHoldings[2],
		})
	}
	// Run keeper that checks to make sure all holdings specified are valid and supported in the pool provided for the broker provided
	err := k.queryKeeper.CheckHoldings(ctx, broker, holdingsList)
	if err != nil {
		return nil, err
	}
	return holdingsList, nil
}

func (k msgServer) CreateFund(goCtx context.Context, msg *types.MsgCreateFund) (*types.MsgCreateFundResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Basic CreateFund validation
	msg.ValidateBasic()

	// Check if the value already exists
	_, isFound := k.GetFund(
		ctx,
		msg.Symbol,
	)
	if isFound {
		return nil, sdkerrors.Wrap(types.ErrSymbolExists, fmt.Sprintf("symbol %s already exists", msg.Symbol))
	}

	// Generate and get a new fund address
	fundAddress := NewFundAddress(msg.Symbol)

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
	err := k.brokerKeeper.RegisterBrokerAccount(ctx, msg.ConnectionId, acc.GetAddress().String())
	if err != nil {
		return nil, err
	}

	holdings, err := k.ParseStringHoldings(ctx, msg.Broker, msg.Holdings)
	if err != nil {
		return nil, err
	}

	var fund = types.Fund{
		Creator:      msg.Creator,
		Symbol:       msg.Symbol,
		Address:      acc.GetAddress().String(),
		Name:         msg.Name,
		Description:  msg.Description,
		Shares:       sdk.NewCoin(GetFundDenom(msg.Symbol), sdk.ZeroInt()),
		Broker:       msg.Broker,
		Holdings:     holdings,
		BaseDenom:    msg.BaseDenom,
		Rebalance:    msg.Rebalance,
		ConnectionId: msg.ConnectionId,
	}

	k.SetFund(
		ctx,
		fund,
	)
	return &types.MsgCreateFundResponse{}, nil
}
