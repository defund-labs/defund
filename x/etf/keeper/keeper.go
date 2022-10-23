package keeper

import (
	"fmt"

	"github.com/defund-labs/defund/x/etf/types"
	"github.com/tendermint/tendermint/libs/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	icacontrollerkeeper "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/controller/keeper"
	icatypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/types"
	transferkeeper "github.com/cosmos/ibc-go/v4/modules/apps/transfer/keeper"
	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	porttypes "github.com/cosmos/ibc-go/v4/modules/core/05-port/types"
	brokertypes "github.com/defund-labs/defund/x/broker/types"
	osmosisgammtypes "github.com/osmosis-labs/osmosis/v8/x/gamm/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey

		scopedKeeper        capabilitykeeper.ScopedKeeper
		accountKeeper       types.AccountKeeper
		bankKeeper          types.BankKeeper
		brokerKeeper        types.BrokerKeeper
		queryKeeper         types.InterqueryKeeper
		channelKeeper       types.ChannelKeeper
		ics4Wrapper         porttypes.ICS4Wrapper
		connectionKeeper    types.ConnectionKeeper
		clientKeeper        types.ClientKeeper
		icaControllerKeeper icacontrollerkeeper.Keeper
		transferKeeper      transferkeeper.Keeper
	}

	Surplus struct {
		BaseDenom      sdk.Coin
		HoldingDenom   sdk.Coin
		Holding        types.Holding
		SurplusPercent sdk.Int
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,

	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	channelKeeper types.ChannelKeeper,
	interqueryKeeper types.InterqueryKeeper,
	brokerKeeper types.BrokerKeeper,
	connectionKeeper types.ConnectionKeeper,
	clientKeeper types.ClientKeeper,
	iaKeeper icacontrollerkeeper.Keeper,
	transferKeeper transferkeeper.Keeper,
) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,

		accountKeeper:       accountKeeper,
		bankKeeper:          bankKeeper,
		channelKeeper:       channelKeeper,
		queryKeeper:         interqueryKeeper,
		brokerKeeper:        brokerKeeper,
		connectionKeeper:    connectionKeeper,
		clientKeeper:        clientKeeper,
		icaControllerKeeper: iaKeeper,
		transferKeeper:      transferKeeper,
	}
}

// ClaimCapability claims the channel capability passed via the OnOpenChanInit callback
func (k *Keeper) ClaimCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) error {
	return k.scopedKeeper.ClaimCapability(ctx, cap, name)
}

// SetICS4Wrapper sets the ICS4 wrapper to the keeper.
// It panics if already set
func (k *Keeper) SetICS4Wrapper(ics4Wrapper porttypes.ICS4Wrapper) {
	if k.ics4Wrapper != nil {
		panic("ICS4 wrapper already set")
	}

	k.ics4Wrapper = ics4Wrapper
}

// Logger returns the module logger
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// helper function to check if a osmosis pool contains denom specified
func containsAssets(assets []osmosisgammtypes.PoolAsset, denom string) bool {
	for _, asset := range assets {
		if asset.Token.Denom == denom {
			return true
		}
	}

	return false
}

func sumDecs(items []sdk.Dec) sdk.Dec {
	sum := sdk.NewDec(0)
	for _, item := range items {
		sum = sum.Add(item)
	}
	return sum
}

// CreateShares send an IBC transfer to all the brokers for each holding with the proportion of tokenIn
// represented in baseDenom that the broker will then rebalance on the next rebalance.
func (k Keeper) CreateShares(ctx sdk.Context, fund types.Fund, channel string, tokenIn sdk.Coin, creator string, timeoutHeight clienttypes.Height, timeoutTimestamp uint64) (numETFShares sdk.Coin, err error) {
	creatorAcc, err := sdk.AccAddressFromBech32(creator)
	if err != nil {
		return numETFShares, err
	}
	fundAcc, err := sdk.AccAddressFromBech32(fund.Address)
	if err != nil {
		return numETFShares, err
	}

	// send the tokenIn to the Defund fund account to ensure that we receive the
	// tokens correctly and instantly to proceed.
	err = k.bankKeeper.SendCoins(ctx, creatorAcc, fundAcc, sdk.NewCoins(tokenIn))
	if err != nil {
		return numETFShares, err
	}

	// for each holding send proportional tokenIn to the holdings broker chain. logic continues in
	// ibc callbacks
	for _, holding := range fund.Holdings {
		broker, found := k.brokerKeeper.GetBroker(ctx, holding.BrokerId)
		if !found {
			return numETFShares, sdkerrors.Wrap(types.ErrWrongBroker, fmt.Sprintf("broker %s not found", holding.BrokerId))
		}

		// get the ica account for the fund on the broker chain
		portID, err := icatypes.NewControllerPortID(fund.Address)
		if err != nil {
			return numETFShares, err
		}
		fundBrokerAddress, found := k.brokerKeeper.GetBrokerAccount(ctx, broker.ConnectionId, portID)
		if !found {
			return numETFShares, sdkerrors.Wrapf(brokertypes.ErrIBCAccountNotExist, "failed to find ica account for owner %s on connection %s and port %s", fund.Address, broker.ConnectionId, portID)
		}

		// Multiply the tokenIn by the % this holding should represent
		sendAmt := tokenIn.Amount.Mul(sdk.NewInt(holding.Percent)).Quo(sdk.NewInt(100))
		sendCoin := sdk.NewCoin(tokenIn.Denom, sendAmt)

		sequence, err := k.SendTransfer(ctx, channel, sendCoin, fund.Address, fundBrokerAddress, timeoutHeight, timeoutTimestamp)
		if err != nil {
			return numETFShares, err
		}
		transfer := brokertypes.Transfer{
			Id:       fmt.Sprintf("%s-%d", channel, sequence),
			Channel:  channel,
			Sequence: sequence,
			Status:   types.TransferStateTransferring,
			Token:    &sendCoin,
			Sender:   fund.Address,
			Receiver: fundBrokerAddress,
		}
		k.brokerKeeper.SetTransfer(ctx, transfer)
	}

	// compute the amount of etf shares this creator is given
	numETFShares, err = k.GetAmountETFSharesForToken(ctx, fund, tokenIn)
	if err != nil {
		return numETFShares, err
	}
	newETFCoins := sdk.NewCoins(numETFShares)

	// finally mint coins (to module account) and then send them to the creator of the create
	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, newETFCoins)
	if err != nil {
		return numETFShares, err
	}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creatorAcc, newETFCoins)
	if err != nil {
		return numETFShares, err
	}

	// finally reflect the new shares in the fund store for shares
	newShares := fund.Shares.Add(numETFShares)
	fund.Shares = &newShares
	k.SetFund(ctx, fund)

	return numETFShares, nil
}

// RedeemShares sends an ICA Send message to each broker chain for each holding to be run on that chain.
// Initializes the redemption of shares process which continues in Broker module in OnAckRec.
func (k Keeper) RedeemShares(ctx sdk.Context, creator string, fund types.Fund, amount sdk.Coin, addressMap types.AddressMap) error {
	// Map for holding all the messages for each broker to send later
	msgs := make(map[string][]*banktypes.MsgSend)
	// Map for keeping track of the total of etf shares redeeming for each broker tx so we can
	// perform logic on acks based on success or failure
	redeem := make(map[string]sdk.Int)

	// Add Osmosis broker to the mappings
	msgs["osmosis"] = []*banktypes.MsgSend{}
	redeem["osmosis"] = sdk.NewInt(0)

	creatorAcc, err := sdk.AccAddressFromBech32(creator)
	if err != nil {
		return err
	}

	// get the amount of tokens that these shares represent
	ownershipRaw, err := k.GetOwnershipSharesInFund(ctx, fund, amount)
	if err != nil {
		return err
	}

	ownership := sdk.NewCoins(ownershipRaw...)

	for _, holding := range fund.Holdings {

		// get the current holding amount and coin from ownership coins
		currentCoinAmt := ownership.Sort().AmountOf(holding.Token)
		currentCoin := sdk.NewCoin(holding.Token, currentCoinAmt)

		// get the current price of this fund (in base denom)
		fundPrice, err := k.CreateFundPrice(ctx, fund.Symbol)
		if err != nil {
			return err
		}
		// multiply the fund price by the current amount of etf tokens to get the basedenom total amount
		totalWorthOfFundTokensInBaseDenom := fundPrice.Amount.Mul(amount.Amount)

		spotPrice, err := k.brokerKeeper.CalculateOsmosisSpotPrice(ctx, holding.PoolId, holding.Token, fund.BaseDenom.OnBroker)
		if err != nil {
			return err
		}
		// calculate the ownership amount in basedenom
		totalOwnershipInBaseDenom := spotPrice.Mul(currentCoinAmt.ToDec())
		// calculate the percent amount of fund token this ownership represents
		percentOwnership := totalOwnershipInBaseDenom.Quo(totalWorthOfFundTokensInBaseDenom.ToDec())
		// multiply total etf shares amount by ownership percent
		etfSharesThisHoldingRepresents := amount.Amount.ToDec().Mul(percentOwnership).RoundInt()
		// Add the amount of the etf shares this holding redeem represents in basedenom
		redeem["osmosis"] = redeem["osmosis"].Add(etfSharesThisHoldingRepresents)

		broker, found := k.brokerKeeper.GetBroker(ctx, holding.BrokerId)
		if !found {
			return sdkerrors.Wrap(types.ErrWrongBroker, fmt.Sprintf("broker %s not found", holding.BrokerId))
		}

		portID, err := icatypes.NewControllerPortID(fund.Address)
		if err != nil {
			ctx.Logger().Debug(err.Error())
		}
		addr, found := k.icaControllerKeeper.GetInterchainAccountAddress(ctx, broker.ConnectionId, portID)
		if !found {
			err := status.Errorf(codes.NotFound, "no account found for portID %s on connection %s", portID, broker.ConnectionId)
			return err
		}
		fundICAAddress, err := sdk.AccAddressFromBech32(addr)
		if err != nil {
			return err
		}

		receiveAddress, err := sdk.AccAddressFromBech32(addressMap.OsmosisAddress)
		if err != nil {
			return err
		}

		msg := banktypes.NewMsgSend(fundICAAddress, receiveAddress, sdk.NewCoins(currentCoin))
		if err != nil {
			return err
		}

		msgs[holding.BrokerId] = append(msgs[holding.BrokerId], msg)
	}

	// take the fund etf shares and escrow them in the module account. in the ack callback, on success
	// of sequence we will burn proportionally. If unsuccessful the transfer is reattempted until successful.
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, creatorAcc, types.ModuleName, sdk.NewCoins(amount))
	if err != nil {
		return err
	}

	// send each ICA message and add it to the redeem which will be used in end blocker
	// to check status of ICA message
	for brokerId, msg := range msgs {
		// get the broker
		broker, found := k.brokerKeeper.GetBroker(ctx, brokerId)
		if !found {
			return sdkerrors.Wrap(types.ErrWrongBroker, fmt.Sprintf("broker %s not found", broker.Id))
		}
		// send the send message to broker chain
		sequence, channel, err := k.brokerKeeper.SendIBCSend(ctx, msg, fund.Address, broker.ConnectionId)
		if err != nil {
			return err
		}

		etfAmount := sdk.NewCoin(fund.Shares.Denom, redeem[brokerId])

		// Create the redeem store
		redeem := types.Redeem{
			Id:      fmt.Sprintf("%s-%d", channel, sequence),
			Creator: creator,
			Fund:    &fund,
			Amount:  &etfAmount,
			Status:  types.RedeemState,
		}

		k.SetRedeem(ctx, redeem)
	}

	return nil
}

// CheckHoldings checks to make sure the specified holdings and the pool for each holding are valid
// by checking the interchain queried pools for the broker specified
func (k Keeper) CheckHoldings(ctx sdk.Context, holdings []*types.Holding) error {
	percentCheck := uint64(0)
	for _, holding := range holdings {
		// Add percent composition to percentCheck to later confirm adds to 100%
		percentCheck = percentCheck + uint64(holding.Percent)
		pool, err := k.brokerKeeper.GetOsmosisPool(ctx, holding.PoolId)
		if err != nil {
			return err
		}
		// Checks to see if the holding pool contains the holding token specified and if not returns error
		if !containsAssets(pool.GetAllPoolAssets(), holding.Token) {
			return sdkerrors.Wrapf(types.ErrInvalidDenom, "invalid denom (%s) for pool (%d) on broker (%s)", holding.Token, holding.PoolId, holding.BrokerId)
		}
		// checks to ensure we have the right holding types current: spot, add: staked
		if holding.Type != "spot" {
			return sdkerrors.Wrapf(types.ErrInvalidHolding, "unsupported holding type. received %s. supported types are 'spot'", holding.Type)
		}
	}
	// Make sure all fund holdings add up to 100%
	if percentCheck != uint64(100) {
		return sdkerrors.Wrapf(types.ErrPercentComp, "percent composition must add up to 100 percent")
	}
	return nil
}

// getOsmosisRoutes is a helper function that looks up the Osmosis broker, takes in the currentDenom, needDenom
// and returns a list of the best routes to go through. It does this by first checking if a currentDenom
// has a direct pool with uosmo, if it does not, it then finds a curentDenom -> usomo with uosmo -> needDenom
// pair to create the routes needed to go from currentDenom -> needDenom.
func (k Keeper) getOsmosisRoutes(ctx sdk.Context, currentDenom string, needDenom string) (routes []osmosisgammtypes.SwapAmountInRoute, err error) {
	broker, found := k.brokerKeeper.GetBroker(ctx, "osmosis")
	if !found {
		return nil, sdkerrors.Wrapf(brokertypes.ErrBrokerNotFound, "broker %s not found", "osmosis")
	}
	// for loop to check if there is a direct pool between currentDenom and needDenom
	for _, pool := range broker.Pools {
		osmoPool, err := k.brokerKeeper.GetOsmosisPool(ctx, pool.PoolId)
		if err != nil {
			return routes, err
		}
		poolAssets := osmoPool.GetAllPoolAssets()

		currentDenomCheck := containsAssets(poolAssets, currentDenom)
		wantDenomCheck := containsAssets(poolAssets, needDenom)

		if currentDenomCheck && wantDenomCheck {
			route := osmosisgammtypes.SwapAmountInRoute{
				PoolId:        osmoPool.GetId(),
				TokenOutDenom: needDenom,
			}
			routes = append(routes, route)
			return routes, nil
		}
	}
	// for loop to create a multi pool route. will run if no direct pool can be found above
	for _, pool := range broker.Pools {
		osmoPool, err := k.brokerKeeper.GetOsmosisPool(ctx, pool.PoolId)
		if err != nil {
			return routes, err
		}
		poolAssets := osmoPool.GetAllPoolAssets()

		currentDenomCheck := containsAssets(poolAssets, currentDenom)
		wantDenomCheck := containsAssets(poolAssets, needDenom)

		if currentDenomCheck {
			route := osmosisgammtypes.SwapAmountInRoute{
				PoolId:        osmoPool.GetId(),
				TokenOutDenom: currentDenom,
			}
			routes = append(routes, route)
		}
		if wantDenomCheck {
			route := osmosisgammtypes.SwapAmountInRoute{
				PoolId:        osmoPool.GetId(),
				TokenOutDenom: needDenom,
			}
			routes = append(routes, route)
		}
	}
	return routes, nil
}

// CreateRebalanceMsgs creates the rebalance messages and returns them for the fund in standard interface
// format.
func (k Keeper) CreateRebalanceMsgs(ctx sdk.Context, fund types.Fund) (types.RebalanceMsgs, error) {

	msgs := types.RebalanceMsgs{}

	// slice to store the holdings with price info
	holdings := types.PricedHoldings{}
	// used to store the total amount of surplus in basedenom
	totalOverInBaseDenom := sdk.NewDec(0)

	for _, holding := range fund.Holdings {
		broker, found := k.brokerKeeper.GetBroker(ctx, holding.BrokerId)
		if !found {
			return msgs, sdkerrors.Wrap(types.ErrWrongBroker, fmt.Sprintf("broker %s not found", holding.BrokerId))
		}
		// get the ica account for the fund on the broker chain
		portID, err := icatypes.NewControllerPortID(fund.Address)
		if err != nil {
			return msgs, err
		}
		fundBrokerAddress, found := k.icaControllerKeeper.GetInterchainAccountAddress(ctx, broker.ConnectionId, portID)
		if !found {
			return msgs, sdkerrors.Wrapf(brokertypes.ErrIBCAccountNotExist, "failed to find ica account for owner %s on connection %s and port %s", fund.Address, broker.ConnectionId, portID)
		}
		var balances banktypes.Balance
		// create the balances based on broker
		switch broker.Id {
		case "osmosis":
			// get the account balances for the fund account on the broker chain
			balances, err = k.GetBalanceForFundByAddress(ctx, fund.Symbol, fundBrokerAddress)
			if err != nil {
				return msgs, err
			}
		}
		// get amount of holding token from balances
		amount := balances.GetCoins().Sort().AmountOf(holding.Token)
		// get the price of the asset in base denom
		priceInBaseDenom, err := k.brokerKeeper.CalculateOsmosisSpotPrice(ctx, holding.PoolId, fund.BaseDenom.OnBroker, holding.Token)
		if err != nil {
			return msgs, err
		}
		// calculate the amount held of holding in base denom by taking amount and multiplying by price in base denom
		amountInBaseDenom := amount.ToDec().Mul(priceInBaseDenom)

		// add to the totalOverInBaseDenom
		totalOverInBaseDenom = totalOverInBaseDenom.Add(amountInBaseDenom)

		holding := types.PricedHolding{
			Holding:        *holding,
			PriceInBase:    amountInBaseDenom,
			PriceInHolding: amount.ToDec(),
		}

		holdings = append(holdings, holding)

	}

	// for loop that creates the ICA messages to swap all proportioned surplus holdings to base denom
	for _, holding := range fund.Holdings {
		broker, found := k.brokerKeeper.GetBroker(ctx, holding.BrokerId)
		if !found {
			return msgs, status.Errorf(codes.NotFound, "broker %s not found", holding.BrokerId)
		}
		// get the ica account for the fund on the broker chain
		portID, err := icatypes.NewControllerPortID(fund.Address)
		if err != nil {
			return msgs, err
		}
		fundBrokerAddress, found := k.brokerKeeper.GetBrokerAccount(ctx, broker.ConnectionId, portID)
		if !found {
			return msgs, sdkerrors.Wrapf(brokertypes.ErrIBCAccountNotExist, "failed to find ica account for owner %s on connection %s and port %s", fund.Address, broker.ConnectionId, portID)
		}
		// use some math to get the current composition % for this holding in the fund
		// (holding in base denom / total in base denom)
		currentComposition, err := holdings.GetPercentComposition(holding.Token)
		if err != nil {
			return msgs, err
		}
		// get the surplus composition % by subtracting the current composition % from what its supposed to be
		overUnderCompPerc := currentComposition.CurrentComp.Sub(sdk.NewDecWithPrec(holding.Percent, 2))

		// if we over own the asset
		if overUnderCompPerc.IsPositive() && !overUnderCompPerc.IsZero() {
			// convert overUnderCompPerc to absolute value (if negative)
			overUnderCompPerc = overUnderCompPerc.Abs()
			// compute the % needed to swap into by multiplying % overUnderCompPerc by the balance of
			// this holding
			amtInHoldingDenom, err := holdings.GetAmountOf(holding.Token, true)
			if err != nil {
				return msgs, err
			}
			// lets calculate the amount we should have
			amountWeShouldHaveInBaseDenom := totalOverInBaseDenom.Mul(sdk.NewDecWithPrec(holding.Percent, 2))
			// get the price of the asset in holding denom
			priceInBaseDenom, err := k.brokerKeeper.CalculateOsmosisSpotPrice(ctx, holding.PoolId, holding.Token, fund.BaseDenom.OnBroker)
			if err != nil {
				return msgs, err
			}
			needToSwapTokenInHoldingDenom := amtInHoldingDenom.Sub(amountWeShouldHaveInBaseDenom).Mul(priceInBaseDenom).RoundInt()
			needToSwapTokenInBaseDenomDenom := overUnderCompPerc.Mul(totalOverInBaseDenom).RoundInt()
			// create the tokenIn coin
			tokenIn := sdk.NewCoin(holding.Token, needToSwapTokenInHoldingDenom)
			// create the min amount out by using the current holding amount in base denom and then
			// NOTE: creating a 2% slippage on it (potentially add this as fund param?)
			tokenOut := needToSwapTokenInBaseDenomDenom.Mul(sdk.NewInt(98)).Quo(sdk.NewInt(100))
			// create holder for msg in switch statement
			var msg *osmosisgammtypes.MsgSwapExactAmountIn
			switch holding.BrokerId {
			case "osmosis":
				// get the routes needed to swap for from this current denom to base denom
				routes, err := k.getOsmosisRoutes(ctx, holding.Token, fund.BaseDenom.OnBroker)
				if err != nil {
					return msgs, err
				}
				msg, err = k.brokerKeeper.CreateOsmosisTrade(ctx, fundBrokerAddress, routes, tokenIn, tokenOut)
				if err != nil {
					return msgs, err
				}
			}
			// append the new message for the broker
			msgs.Osmosis = append(msgs.Osmosis, msg)
		}
	}

	// for loop that creates the ICA messages to swap from base denom to needed/under owned proportioned holding.
	// We must do this after the positive for loop (above) that creates swaps to base denom so we can
	// run these swaps after the swap to base denom.
	for _, holding := range fund.Holdings {
		broker, found := k.brokerKeeper.GetBroker(ctx, holding.BrokerId)
		if !found {
			return msgs, status.Errorf(codes.NotFound, "broker %s not found", holding.BrokerId)
		}
		// get the ica account for the fund on the broker chain
		portID, err := icatypes.NewControllerPortID(fund.Address)
		if err != nil {
			return msgs, err
		}
		fundBrokerAddress, found := k.brokerKeeper.GetBrokerAccount(ctx, broker.ConnectionId, portID)
		if !found {
			return msgs, sdkerrors.Wrapf(brokertypes.ErrIBCAccountNotExist, "failed to find ica account for owner %s on connection %s and port %s", fund.Address, broker.ConnectionId, portID)
		}
		// use some math to get the current composition % for this holding in the fund
		// (holding in base denom / total in base denom)
		currentComposition, err := holdings.GetPercentComposition(holding.Token)
		if err != nil {
			return msgs, err
		}
		// get the surplus composition % by subtracting the current composition % from what its supposed to be
		overUnderCompPerc := currentComposition.CurrentComp.Sub(sdk.NewDec(holding.Percent).Quo(sdk.NewDec(100)))

		if overUnderCompPerc.IsNegative() && !overUnderCompPerc.IsZero() {
			// if the current denom is the base denom, we can just skip for base denom as the holding
			// already is in base denom from the above
			if fund.BaseDenom.OnBroker == holding.Token {
				continue
			}
			// convert overUnderCompPerc to absolute value (if negative)
			overUnderCompPerc = overUnderCompPerc.Abs()
			needToSwapTokenInBaseDenom := overUnderCompPerc.Mul(totalOverInBaseDenom).RoundInt()
			// create the tokenIn coin
			tokenIn := sdk.NewCoin(fund.BaseDenom.OnBroker, needToSwapTokenInBaseDenom)
			tokenPriceInHoldingDenom, err := k.brokerKeeper.CalculateOsmosisSpotPrice(ctx, holding.PoolId, holding.Token, fund.BaseDenom.OnBroker)
			if err != nil {
				return msgs, err
			}
			// create the min amount out by using the current holding amount in base denom and then
			// NOTE: creating a 2% slippage on it (potentially add this as fund param?)
			tokenOut := needToSwapTokenInBaseDenom.ToDec().Mul(tokenPriceInHoldingDenom).Mul(sdk.NewDec(98)).Quo(sdk.NewDec(100)).RoundInt()
			// create holder for msg in switch statement
			var msg *osmosisgammtypes.MsgSwapExactAmountIn
			switch holding.BrokerId {
			case "osmosis":
				// get the routes needed to swap for from this current denom to base denom
				routes, err := k.getOsmosisRoutes(ctx, fund.BaseDenom.OnBroker, holding.Token)
				if err != nil {
					return msgs, err
				}
				msg, err = k.brokerKeeper.CreateOsmosisTrade(ctx, fundBrokerAddress, routes, tokenIn, tokenOut)
				if err != nil {
					return msgs, err
				}
			}
			// append the new message for the broker
			msgs.Osmosis = append(msgs.Osmosis, msg)
		}
	}

	return msgs, nil
}

// SendRebalanceTx sends an ICA tx to each broker chain with msgs that will rebalance the underlying
// holdings for the fund
func (k Keeper) SendRebalanceTx(ctx sdk.Context, fund types.Fund) error {
	// Create the rebalance messages
	msgs, err := k.CreateRebalanceMsgs(ctx, fund)
	if err != nil {
		return err
	}

	// send trades for each supported brokers as long as we have more then one rebalance message for it

	// Osmosis Broker Send
	// TODO: add a SendStakingTxs function after SendOsmosisTrades for staking ETFs for each holding thats staked
	if len(msgs.Osmosis) > 0 {
		broker, found := k.brokerKeeper.GetBroker(ctx, "osmosis")
		if !found {
			return status.Errorf(codes.NotFound, "broker %s not found", "osmosis")
		}
		_, err = k.brokerKeeper.SendOsmosisTrades(ctx, msgs.Osmosis, fund.Address, broker.ConnectionId)
		if err != nil {
			return err
		}
	}

	return nil
}
