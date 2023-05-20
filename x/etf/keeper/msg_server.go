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
	ibctransfertypes "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
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

func removeLastChannel(path string) string {
	var newPath string = ""
	splits := strings.Split(path, "/")
	if len(splits) > 2 {
		for i := range splits[0 : len(splits)-2] {
			newPath = newPath + splits[i]
			if i != len(splits)-3 {
				newPath = newPath + "/"
			}
		}
	}
	if len(splits) == 2 {
		newPath = ""
	}
	return newPath
}

// RegisterBrokerAccounts opens one broker account per broker Defund supports
func (k msgServer) RegisterBrokerAccounts(ctx sdk.Context, holdings []*types.Holding, acc string) error {
	// get all brokers
	brokers := k.brokerKeeper.GetAllBrokers(ctx)
	for _, broker := range brokers {
		// ensure the broker is active and has connection id assigned to it
		if broker.Status != "active" {
			return sdkerrors.Wrap(types.ErrWrongBroker, fmt.Sprintf("broker %s status is not active (status: %s)", broker.Id, broker.Status))
		}

		// Create and save the broker fund ICA account on the broker chain
		err := k.RegisterBrokerAccount(ctx, broker.ConnectionId, acc)
		if err != nil {
			return err
		}
	}

	return nil
}

// Helper function that parses a string of holdings in the format "uatom:50:osmosis:1:spot,uosmo:50:osmosis:2:spot" (DENOM:PERCENT:BROKER:POOL:TYPE,...) into a slice of type holding
// and checks to make sure that the holdings are all supported denoms from the specified broker and pool
func (k msgServer) ParseStringHoldings(ctx sdk.Context, holdings string, basedenom string) ([]*types.Holding, error) {
	rawHoldings := strings.Split(holdings, ",")
	holdingsList := []*types.Holding{}
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
		holding := types.Holding{
			Token:    sepHoldings[0],
			Percent:  perc,
			PoolId:   poolid,
			BrokerId: sepHoldings[2],
			Type:     sepHoldings[4],
		}
		holdingsList = append(holdingsList, &holding)
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

	// Check if the value already exists
	_, isFound := k.GetFund(
		ctx,
		msg.Symbol,
	)
	if isFound {
		return nil, sdkerrors.Wrap(types.ErrSymbolExists, fmt.Sprintf("symbol %s already exists", msg.Symbol))
	}

	// check to make sure proper base denom is used
	if msg.BaseDenom != "atom" && msg.BaseDenom != "osmo" {
		return nil, sdkerrors.Wrap(types.ErrWrongBaseDenom, fmt.Sprintf("denom %s is not a valid base denom. must be atom or osmo", msg.BaseDenom))
	}

	var basedenom types.BaseDenom
	basedenoms := k.brokerKeeper.GetBaseDenomParam(ctx)
	if msg.BaseDenom == "atom" {
		basedenom.OnDefund = basedenoms.AtomTrace.IBCDenom()
		// create the broker chain denom
		newPath := removeLastChannel(basedenoms.AtomTrace.Path)
		newTrace := ibctransfertypes.DenomTrace{
			Path:      newPath,
			BaseDenom: basedenoms.AtomTrace.BaseDenom,
		}
		basedenom.OnBroker = newTrace.IBCDenom()
	}
	if msg.BaseDenom == "osmo" {
		basedenom.OnDefund = basedenoms.OsmoTrace.IBCDenom()
		// create the broker chain denom
		newPath := removeLastChannel(basedenoms.OsmoTrace.Path)
		newTrace := ibctransfertypes.DenomTrace{
			Path:      newPath,
			BaseDenom: basedenoms.OsmoTrace.BaseDenom,
		}
		basedenom.OnBroker = newTrace.IBCDenom()
	}

	// Generate and get a new fund address
	fundAddress := NewFundAddress(msg.Symbol)

	var t types.FundType = types.FundType_PASSIVE
	var contractAddress string
	if msg.Active {
		// set the fund type to active
		t = types.FundType_ACTIVE
		// instantiate a wasm contract from code id provided
		contract, _, err := k.wasmInternalKeeper.Instantiate(ctx, msg.WasmCodeId, fundAddress, fundAddress, []byte(fmt.Sprintf(`{"fund": "%s"}`, msg.Symbol)), msg.Symbol, sdk.NewCoins())
		if err != nil {
			return nil, err
		}
		contractAddress = contract.String()
	}

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

	holdings, err := k.ParseStringHoldings(ctx, msg.Holdings, basedenom.OnBroker)
	if err != nil {
		return nil, err
	}

	// Check and create all broker accounts for fund
	err = k.RegisterBrokerAccounts(ctx, holdings, fundAddress.String())
	if err != nil {
		return nil, err
	}

	// Convert starting price to coin format
	rawIntStartingPrice, err := strconv.ParseInt(msg.StartingPrice, 10, 64)
	if err != nil {
		return nil, err
	}
	startPrice := sdk.NewCoin(basedenom.OnBroker, sdk.NewInt(rawIntStartingPrice))
	shares := sdk.NewCoin(msg.Symbol, sdk.ZeroInt())

	balances := types.FundBalances{
		Osmosis: types.Balances{
			Address:  "",
			Balances: []*sdk.Coin{},
		},
	}

	var fund = types.Fund{
		Creator:             msg.Creator,
		Symbol:              msg.Symbol,
		Address:             acc.GetAddress().String(),
		Name:                msg.Name,
		Description:         msg.Description,
		Shares:              &shares,
		Holdings:            holdings,
		BaseDenom:           &basedenom,
		Rebalance:           msg.Rebalance,
		Rebalancing:         false,
		LastRebalanceHeight: 0,
		StartingPrice:       &startPrice,
		Balances:            &balances,
		FundType:            t,
		Contract:            contractAddress,
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

	// basic validation of the message
	msg.ValidateBasic()

	fund, found := k.GetFund(ctx, msg.Fund)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrFundNotFound, "failed to find fund with symbol of %s", msg.Fund)
	}

	// base denom tokenIn check
	if msg.TokenIn.Denom != fund.BaseDenom.OnDefund {
		return nil, sdkerrors.Wrap(types.ErrWrongBaseDenom, fmt.Sprintf("denom %s is not a valid base denom for token in. the base denom for this fund is %s", msg.TokenIn.Denom, fund.BaseDenom.OnDefund))
	}

	timeoutHeight, err := clienttypes.ParseHeight(msg.TimeoutHeight)
	if err != nil {
		return nil, err
	}

	err = k.Keeper.CreateShares(ctx, fund, msg.Channel, *msg.TokenIn, msg.Creator, timeoutHeight, msg.TimeoutTimestamp)
	if err != nil {
		return nil, err
	}

	return &types.MsgCreateResponse{}, nil
}

// Redeem is the Msg handler that redeems new fund shares
func (k msgServer) Redeem(goCtx context.Context, msg *types.MsgRedeem) (*types.MsgRedeemResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// basic validation of the message
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

func (k msgServer) EditFund(goCtx context.Context, msg *types.MsgEditFund) (*types.MsgEditFundResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	msg.ValidateBasic()

	// get the fund
	fund, err := k.GetFundBySymbol(ctx, msg.Symbol)
	if err != nil {
		return &types.MsgEditFundResponse{}, err
	}

	// update the fund for the holdings if provided
	if fund.Holdings != nil {
		// ensure the fund is active type
		if fund.FundType != types.FundType_ACTIVE {
			return &types.MsgEditFundResponse{}, sdkerrors.Wrapf(types.ErrUnauthorized, "invalid fund type only active funds can edit holdings")
		}

		// convert the raw string address to addr bytes
		contractAddr, err := sdk.AccAddressFromBech32(fund.Contract)
		if err != nil {
			return &types.MsgEditFundResponse{}, err
		}
		// ensure the contract specified in fund is a wasm contract
		isContract := k.wasmKeeper.HasContractInfo(ctx, contractAddr)
		if !isContract {
			return &types.MsgEditFundResponse{}, sdkerrors.Wrapf(types.ErrUnauthorized, "contract specified for fund (%s) is not a cosmwasm contract", fund.Contract)
		}
		// make sure the signer is the wasm contract
		if msg.Creator != fund.Contract {
			return &types.MsgEditFundResponse{}, sdkerrors.Wrapf(types.ErrUnauthorized, "msg signer (%s) is not the funds contract (%s)", msg.Creator, fund.Contract)
		}

		fund.Holdings = msg.GetHoldings()
	}

	k.SetFund(ctx, fund)

	return &types.MsgEditFundResponse{}, nil
}
