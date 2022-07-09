package keeper

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	"github.com/defund-labs/defund/x/etf/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

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
		poolid, err := strconv.ParseUint(sepHoldings[2], 10, 64)
		if err != nil {
			return nil, err
		}
		holdingsList = append(holdingsList, types.Holding{
			Token:   sepHoldings[0],
			Percent: perc,
			PoolId:  poolid,
		})
	}
	// Run keeper that checks to make sure all holdings specified are valid and supported in the pool provided for the broker provided
	err := k.CheckHoldings(ctx, broker, holdingsList)
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

	broker, found := k.brokerKeeper.GetBroker(ctx, msg.Broker)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrWrongBroker, fmt.Sprintf("broker %s not found", msg.Broker))
	}

	// Create and save the broker fund ICA account on the broker chain
	err := k.brokerKeeper.RegisterBrokerAccount(ctx, broker.ConnectionId, acc.GetAddress().String())
	if err != nil {
		return nil, err
	}

	holdings, err := k.ParseStringHoldings(ctx, msg.Broker, msg.Holdings)
	if err != nil {
		return nil, err
	}

	// Convert starting price to coin format
	rawIntStartingPrice, err := strconv.ParseInt(msg.StartingPrice, 10, 64)
	if err != nil {
		return nil, err
	}
	startPrice := sdk.NewCoin(msg.BaseDenom, sdk.NewInt(rawIntStartingPrice))

	var fund = types.Fund{
		Creator:       msg.Creator,
		Symbol:        msg.Symbol,
		Address:       acc.GetAddress().String(),
		Name:          msg.Name,
		Description:   msg.Description,
		Shares:        sdk.NewCoin(GetFundDenom(msg.Symbol), sdk.ZeroInt()),
		Broker:        &broker,
		Holdings:      holdings,
		BaseDenom:     msg.BaseDenom,
		Rebalance:     msg.Rebalance,
		ConnectionId:  broker.ConnectionId,
		StartingPrice: startPrice,
	}

	k.SetFund(
		ctx,
		fund,
	)
	return &types.MsgCreateFundResponse{}, nil
}

func (k msgServer) Create(goCtx context.Context, msg *types.MsgCreate) (*types.MsgCreateResponse, error) {
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

	err = k.Keeper.CreateShares(ctx, id, acc, fund, msg.Channel, *msg.Amount, msg.Creator, fund.Address, timeoutHeight, msg.TimeoutTimestamp)
	if err != nil {
		return nil, err
	}

	return &types.MsgCreateResponse{}, nil
}

func (k msgServer) Redeem(goCtx context.Context, msg *types.MsgRedeem) (*types.MsgRedeemResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// get the fund and check if it exists
	fund, found := k.GetFund(ctx, msg.Fund)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrFundNotFound, "failed to find fund with id of %s", fund.Symbol)
	}

	id := k.GetNextID(ctx)

	err := k.Keeper.RedeemShares(ctx, id, fund, msg.Channel, *msg.Amount, fund.Address, msg.Creator)
	if err != nil {
		return nil, err
	}

	return &types.MsgRedeemResponse{}, nil
}
