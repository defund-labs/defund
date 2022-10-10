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
	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
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
	return fmt.Sprintf("etf/%s", symbol)
}

func containsString(strings []string, value string) bool {
	for _, string := range strings {
		if string == value {
			return true
		}
	}
	return false
}

// RegisterBrokerAccounts checks to make sure if all broker accounts are created for holdings within
// a fund. If no broker account exists, one is created and then stored in the Broker store
func (k msgServer) RegisterBrokerAccounts(ctx sdk.Context, holdings []types.Holding, acc authtypes.AccountI) error {
	// we must keep track of broker accounts registered so we can make sure we create only one
	// account per broker.
	var registeredBrokers []string
	for _, holding := range holdings {
		// make sure we do not already have account for broker for this fund
		if containsString(registeredBrokers, holding.BrokerId) {
			continue
		}
		broker, found := k.brokerKeeper.GetBroker(ctx, holding.BrokerId)
		if !found {
			return sdkerrors.Wrap(types.ErrWrongBroker, fmt.Sprintf("broker %s not found for holding %s", holding.BrokerId, holding.Token))
		}

		// ensure the broker is active and has connection id assigned to it
		if broker.Status != "active" {
			return sdkerrors.Wrap(types.ErrWrongBroker, fmt.Sprintf("broker %s status is not active (status: %s) for holding %s", holding.BrokerId, broker.Status, holding.Token))
		}

		// Create and save the broker fund ICA account on the broker chain
		err := k.brokerKeeper.RegisterBrokerAccount(ctx, broker.ConnectionId, acc.GetAddress().String())
		if err != nil {
			return err
		}
		registeredBrokers = append(registeredBrokers, holding.BrokerId)
	}

	return nil
}

// Helper function that parses a string of holdings in the format "ATOM:50:osmosis:1,OSMO:50:osmosis:2" (DENOM:PERCENT:BROKER:POOL,...) into a slice of type holding
// and checks to make sure that the holdings are all supported denoms from the specified broker and pool
func (k msgServer) ParseStringHoldings(ctx sdk.Context, holdings string, basedenom string) ([]types.Holding, error) {
	rawHoldings := strings.Split(holdings, ",")
	holdingsList := []types.Holding{}
	var foundBaseDenom bool
	for _, holding := range rawHoldings {
		sepHoldings := strings.Split(holding, ":")
		perc, err := strconv.ParseInt(sepHoldings[1], 10, 64)
		if err != nil {
			return nil, err
		}
		poolid, err := strconv.ParseUint(sepHoldings[3], 10, 64)
		if err != nil {
			return nil, err
		}
		// if this is base denom mark we have a holding for it
		if sepHoldings[0] == basedenom {
			foundBaseDenom = true
		}
		holdingsList = append(holdingsList, types.Holding{
			Token:    sepHoldings[0],
			Percent:  perc,
			PoolId:   poolid,
			BrokerId: sepHoldings[2],
			Type:     sepHoldings[4],
		})
	}
	// if no base denom holding error
	if !foundBaseDenom {
		return nil, sdkerrors.Wrapf(types.ErrWrongBaseDenom, "the base denom %s must be included as a holding", basedenom)
	}
	// Run keeper that checks to make sure all holdings specified are valid and supported in the pool provided for the broker provided
	err := k.CheckHoldings(ctx, holdingsList)
	if err != nil {
		return nil, err
	}
	return holdingsList, nil
}

// CreateFund is the Msg handler that creates a new fund in store and initializes everything
// for the creation of that fund
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

	holdings, err := k.ParseStringHoldings(ctx, msg.Holdings, msg.BaseDenom)
	if err != nil {
		return nil, err
	}

	// Check and create all broker accounts for fund
	err = k.RegisterBrokerAccounts(ctx, holdings, acc)
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
		Holdings:      holdings,
		BaseDenom:     msg.BaseDenom,
		Rebalance:     msg.Rebalance,
		StartingPrice: startPrice,
	}

	k.SetFund(
		ctx,
		fund,
	)
	return &types.MsgCreateFundResponse{}, nil
}

// Create is the Msg handler that creates new fund shares
func (k msgServer) Create(goCtx context.Context, msg *types.MsgCreate) (*types.MsgCreateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	//basic validation of the message
	msg.ValidateBasic()

	fund, found := k.GetFund(ctx, msg.Fund)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrFundNotFound, "failed to find fund with id of %s", fund.Symbol)
	}

	timeoutHeight, err := clienttypes.ParseHeight(msg.TimeoutHeight)
	if err != nil {
		return nil, err
	}

	_, err = k.Keeper.CreateShares(ctx, fund, msg.Channel, *msg.TokenIn, msg.Creator, timeoutHeight, msg.TimeoutTimestamp)
	if err != nil {
		return nil, err
	}

	return &types.MsgCreateResponse{}, nil
}

// Redeem is the Msg handler that redeems new fund shares
func (k msgServer) Redeem(goCtx context.Context, msg *types.MsgRedeem) (*types.MsgRedeemResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	//basic validation of the message
	msg.ValidateBasic()

	// get the fund and check if it exists
	fund, found := k.GetFund(ctx, msg.Fund)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrFundNotFound, "failed to find fund with id of %s", fund.Symbol)
	}

	err := k.Keeper.RedeemShares(ctx, msg.Creator, fund, *msg.Amount, *msg.Addresses)
	if err != nil {
		return nil, err
	}

	return &types.MsgRedeemResponse{}, nil
}
