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
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	icacontrollerkeeper "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/controller/keeper"
	icatypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/types"
	transferkeeper "github.com/cosmos/ibc-go/v4/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
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

// helper function to check if a list of coins contains a token denom
func containsdenom(list []sdk.Coin, denom string) bool {
	for _, value := range list {
		if value.Denom == denom {
			return true
		}
	}

	return false
}

func sum(items []sdk.Dec) sdk.Dec {
	sum := sdk.NewDec(0)
	for _, item := range items {
		sum = sum.Add(item)
	}
	return sum
}

func sumInts(items []sdk.Int) sdk.Int {
	sum := sdk.NewInt(0)
	for _, item := range items {
		sum = sum.Add(item)
	}
	return sum
}

// helper that removes the index item from the slice of Surplus and then returns the modified slice
func remove(index int, s []Surplus) []Surplus {
	s[index] = s[len(s)-1]
	s[len(s)-1] = Surplus{}
	s = s[:len(s)-1]
	return s
}

// CreateShares sends a multi-send of assets to create ETF shares from creator to the module account
// which then sends an IBC transfer to the fund account on the broker chain and creates a pending transfer store.
// Initializes the create shares process which continues in Broker module in OnAckRec.
func (k Keeper) CreateShares(ctx sdk.Context, fund types.Fund, channel string, tokens []*sdk.Coin, creator string, timeoutHeight clienttypes.Height, timeoutTimestamp uint64) error {
	// Need to convert the coins to plain coins for multi send
	coins := sdk.Coins{}
	// Placeholder for current coin to be set below
	currentCoin := sdk.Coin{}
	for _, token := range tokens {
		coins = append(coins, *token)
	}
	creatorAcc, err := sdk.AccAddressFromBech32(creator)
	if err != nil {
		return err
	}
	fundAcc, err := sdk.AccAddressFromBech32(fund.Address)
	if err != nil {
		return err
	}

	// send the tokens to the Defund fund account to ensure that we receive the
	// tokens correctly.
	err = k.bankKeeper.SendCoins(ctx, creatorAcc, fundAcc, sdk.NewCoins(coins...))
	if err != nil {
		return err
	}

	// for each token send IBC transfer to move funds to broker chain. logic continues in ibc callbacks
	for _, holding := range fund.Holdings {
		broker, found := k.brokerKeeper.GetBroker(ctx, holding.BrokerId)
		if !found {
			return sdkerrors.Wrap(types.ErrWrongBroker, fmt.Sprintf("broker %s not found", holding.BrokerId))
		}

		// get the ica account for the fund on the broker chain
		portID, err := icatypes.NewControllerPortID(fund.Address)
		if err != nil {
			return err
		}
		fundBrokerAddress, found := k.brokerKeeper.GetBrokerAccount(ctx, broker.ConnectionId, portID)
		if !found {
			return sdkerrors.Wrapf(brokertypes.ErrIBCAccountNotExist, "failed to find ica account for owner %s on connection %s and port %s", fund.Address, broker.ConnectionId, portID)
		}

		// get the token supplied by user for the current holding in loop
		for i, coin := range coins {
			if coin.Denom == holding.Token {
				currentCoin = coin
				break
			}
			// if we are at the end of the list and it has not broke, we are missing the token. Return error
			if (len(coins) - 1) == i {
				return sdkerrors.Wrapf(types.ErrInvalidDenom, "could not find supplied token representing holding denom: %s", holding.Token)
			}
		}

		sequence, err := k.SendTransfer(ctx, channel, currentCoin, fund.Address, fundBrokerAddress, timeoutHeight, timeoutTimestamp)
		if err != nil {
			return err
		}
		transfer := brokertypes.Transfer{
			Id:       fmt.Sprintf("%s-%d", channel, sequence),
			Channel:  channel,
			Sequence: sequence,
			Status:   "tranferring",
			Token:    &currentCoin,
			Sender:   fund.Address,
			Receiver: fundBrokerAddress,
		}
		k.brokerKeeper.SetTransfer(ctx, transfer)
	}

	// compute the amount of etf shares this creator is given
	numETFShares, err := k.GetAmountETFSharesForTokens(ctx, fund, tokens)
	if err != nil {
		return err
	}
	newETFCoins := sdk.NewCoins(numETFShares)

	// finally mint coins (to module account) and then send them to the creator of the create
	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, newETFCoins)
	if err != nil {
		return err
	}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creatorAcc, newETFCoins)
	if err != nil {
		return err
	}

	// finally reflect the new shares in the fund store for shares
	fund.Shares = fund.Shares.Add(numETFShares)
	k.SetFund(ctx, fund)

	return nil
}

// RedeemShares sends an ICA MultiSend message to the broker chain to be run on that chain.
// Initializes the redemption of shares process which continues in Broker module in OnAckRec.
func (k Keeper) RedeemShares(ctx sdk.Context, id string, fund types.Fund, channel string, amount sdk.Coin, fundAccount string, receiver string) error {
	// Placeholder for current coin to be set below
	currentCoin := sdk.Coin{}
	// Map for holding all the messages for each broker to send later
	msgs := make(map[string][]*ibctransfertypes.MsgTransfer)

	receiverAcc, err := sdk.AccAddressFromBech32(receiver)
	if err != nil {
		return err
	}

	// get the amount of tokens that these shares represent
	ownership, err := k.GetOwnershipSharesInFund(ctx, fund, amount)
	if err != nil {
		return err
	}

	redeem := types.Redeem{
		Id:        id,
		Creator:   receiver,
		Fund:      &fund,
		Amount:    &amount,
		Channel:   channel,
		Status:    "pending",
		Transfers: []brokertypes.Transfer{},
	}

	for _, holding := range fund.Holdings {
		// get the token to redeem for the current holding in loop
		for i, coin := range ownership {
			if coin.Denom == holding.Token {
				currentCoin = coin
				break
			}
			// if we are at the end of the list and it has not broke, we are missing the token. Return error
			if (len(ownership) - 1) == i {
				return sdkerrors.Wrapf(types.ErrInvalidDenom, "could not find supplied token representing holding denom: %s", holding.Token)
			}
		}

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
			ctx.Logger().Debug(err.Error())
		}

		msg := ibctransfertypes.NewMsgTransfer("transfer", channel, currentCoin, addr, fund.Address, clienttypes.NewHeight(0, 0), 0)
		if err != nil {
			return err
		}

		msgs[holding.BrokerId] = append(msgs[holding.BrokerId], msg)

		// take the fund etf shares and escrow them in the module account. in the ack callback, on success
		// of all sequences we will burn these shares. If unsuccessful we will send them back to the user (same on timeout).
		err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, receiverAcc, types.ModuleName, sdk.NewCoins(amount))
		if err != nil {
			return err
		}
	}

	// send each ICA message and add it to the redeem which will be used in end blocker
	// to check status of ICA message
	for brokerId, msg := range msgs {
		// get the broker
		broker, found := k.brokerKeeper.GetBroker(ctx, brokerId)
		if !found {
			return sdkerrors.Wrap(types.ErrWrongBroker, fmt.Sprintf("broker %s not found", broker.Id))
		}
		// create the ica multi send message
		sequence, channelICA, err := k.SendIBCTransferICA(ctx, msg, fund.Address, broker.ConnectionId)
		if err != nil {
			return err
		}
		// get the ica account address port
		portID, err := icatypes.NewControllerPortID(fund.Address)
		if err != nil {
			return err
		}
		// get the ica account address
		accAddress, found := k.icaControllerKeeper.GetInterchainAccountAddress(ctx, broker.ConnectionId, portID)
		if !found {
			return status.Errorf(codes.NotFound, "no account found for portID %s", portID)
		}
		transfer := brokertypes.Transfer{
			Id:       fmt.Sprintf("%s-%d", channelICA, sequence),
			Channel:  channelICA,
			Sequence: sequence,
			Status:   "tranferring",
			Token:    &currentCoin,
			Sender:   accAddress,
			Receiver: fund.Address,
		}
		redeem.Transfers = append(redeem.Transfers, transfer)
	}

	k.SetRedeem(ctx, redeem)

	return nil
}

// CheckHoldings checks to make sure the specified holdings and the pool for each holding are valid
// by checking the interchain queried pools for the broker specified
func (k Keeper) CheckHoldings(ctx sdk.Context, holdings []types.Holding) error {
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
			return sdkerrors.Wrapf(types.ErrInvalidDenom, "invalid/unsupported denom (%s) in pool (%d)", holding.Token, holding.PoolId)
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
	// for loop to create a multi pool route. will run if no direct pool can be found. should only
	// contain no more than two routes
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
		if len(routes) >= 2 {
			break
		}
	}
	return routes, nil
}

// HandleSurplus attempts to match a surplus with an opposite sign surplus if there are any. Any leftovers
// will be added to the surplusList slice to be used later
func (k Keeper) HandleSurplus(ctx sdk.Context, fund types.Fund, holding types.Holding, msgs []*osmosisgammtypes.MsgSwapExactAmountIn, surplusList []Surplus, currentSurplus Surplus) ([]*osmosisgammtypes.MsgSwapExactAmountIn, []Surplus, error) {
	broker, found := k.brokerKeeper.GetBroker(ctx, holding.BrokerId)
	if !found {
		return msgs, surplusList, status.Errorf(codes.NotFound, "broker %s not found", holding.BrokerId)
	}
	// get the ica account for the fund on the broker chain
	portID, err := icatypes.NewControllerPortID(fund.Address)
	if err != nil {
		return msgs, surplusList, err
	}
	fundBrokerAddress, found := k.brokerKeeper.GetBrokerAccount(ctx, broker.ConnectionId, portID)
	if !found {
		return msgs, surplusList, sdkerrors.Wrapf(brokertypes.ErrIBCAccountNotExist, "failed to find ica account for owner %s on connection %s and port %s", fund.Address, broker.ConnectionId, portID)
	}
	if len(surplusList) > 0 {
		// if the holding is over owned
		if currentSurplus.SurplusPercent.GT(sdk.NewInt(0)) {
			// search through each surplusList looking for a negative match
			for i := range surplusList {
				// if the surplusList item is a negative surplus (we can match with positive)
				if surplusList[i].SurplusPercent.LT(sdk.NewInt(0)) {
					// get the routes needed to swap for from this currentSurplus to surplusList item
					routes, err := k.getOsmosisRoutes(ctx, currentSurplus.Holding.Token, surplusList[i].Holding.Token)
					if err != nil {
						return msgs, surplusList, err
					}
					// if the current surplus is greater then this surplusList item, add full item as swap
					// in msg list and then delete the surplusList item continuing to the next iteration in the loop
					if currentSurplus.BaseDenom.Amount.GT(surplusList[i].BaseDenom.Amount) {
						// compute the % needed to swap of the total currentSurplus by dividing the surplusList item in base denom by
						// by the currentSurplus in base denom
						needToSwapPortion := surplusList[i].BaseDenom.Amount.Quo(currentSurplus.BaseDenom.Amount)
						// create the tokenin coin by multiplying the portion by currentSurlus we have to much of
						tokenIn := sdk.NewCoin(currentSurplus.HoldingDenom.Denom, currentSurplus.HoldingDenom.Amount.Mul(needToSwapPortion).Quo(sdk.NewInt(100)))
						// create the max amount out by using the full surplusList amount (since we are going to use all of it)
						// and then creating a 2% slippage on it (potentially add this as fund param?)
						tokenOut := surplusList[i].HoldingDenom.Amount.Mul(sdk.NewInt(98)).Quo(sdk.NewInt(100))
						msg, err := k.brokerKeeper.CreateOsmosisTrade(ctx, fundBrokerAddress, routes, tokenIn, tokenOut)
						if err != nil {
							return msgs, surplusList, err
						}
						// add the msg to the msgs list
						msgs = append(msgs, msg)
						// update the HoldingDenom of currentSurplus by subtracting what we are swapping in the HoldingDenom
						currentSurplus.HoldingDenom = currentSurplus.HoldingDenom.Sub(tokenIn)
						// update the BaseDenom of currentSurplus by subtracting the total BaseDenom amount of surplusList item
						currentSurplus.BaseDenom = currentSurplus.BaseDenom.Sub(surplusList[i].BaseDenom)
						// remove the fully used surplusList item from list
						surplusList = remove(i, surplusList)
					}
					// if the current surplus is less then this surplusList item, make out the full swap for currentSurplus,
					// remove the amount swapped from in surplusList item and break the loop as this currentSurplus is done
					if currentSurplus.BaseDenom.Amount.LT(surplusList[i].BaseDenom.Amount) {
						// compute the % needed to swap of the total surplusList item by dividing the currentSurplus in base denom by
						// by the surplusList item in base denom
						needToSwapPortion := currentSurplus.BaseDenom.Amount.Quo(surplusList[i].BaseDenom.Amount)
						// create the tokenin coin by using the full currentSurplus
						tokenIn := currentSurplus.HoldingDenom
						// create the max amount out by using the full surplusList amount and multiplying it by the portion
						// (since we are not going to use all of it) and then creating a 2% slippage on it (potentially add this as fund param?)
						tokenOut := surplusList[i].HoldingDenom.Amount.Mul(needToSwapPortion).Mul(sdk.NewInt(98)).Quo(sdk.NewInt(100))
						msg, err := k.brokerKeeper.CreateOsmosisTrade(ctx, fundBrokerAddress, routes, tokenIn, tokenOut)
						if err != nil {
							return msgs, surplusList, err
						}
						// add the msg to the msgs list
						msgs = append(msgs, msg)
						// update the HoldingDenom of surplusList item by subtracting what we are swapping out the HoldingDenom
						surplusList[i].HoldingDenom = surplusList[i].HoldingDenom.Sub(sdk.NewCoin(surplusList[i].HoldingDenom.Denom, tokenOut))
						// update the BaseDenom of surplusList item by subtracting the total BaseDenom amount of currentSurplus item
						surplusList[i].BaseDenom = surplusList[i].BaseDenom.Sub(currentSurplus.BaseDenom)
						break
					}
					// if the current surplus equals the surplusList item, swap the current surplus over in full,
					// remove the surplusList item and then break the loop as this currentSurplus is done
					if currentSurplus.BaseDenom.Amount.Equal(surplusList[i].BaseDenom.Amount) {
						// create the tokenin coin by using the full currentSurplus
						tokenIn := currentSurplus.HoldingDenom
						// create the max amount out by using the full surplusList amount
						// (since we are going to use all of it) and then creating a 2% slippage on it (potentially add this as fund param?)
						tokenOut := surplusList[i].HoldingDenom.Amount.Mul(sdk.NewInt(98)).Quo(sdk.NewInt(100))
						// create the osmosis swap message with 2% slippage
						msg, err := k.brokerKeeper.CreateOsmosisTrade(ctx, fundBrokerAddress, routes, tokenIn, tokenOut)
						if err != nil {
							return msgs, surplusList, err
						}
						// add the msg to the msgs list
						msgs = append(msgs, msg)
						// remove the fully used surplusList item from list
						surplusList = remove(i, surplusList)
						break
					}
				}
			}
		}
		// if the holding is under owned
		if currentSurplus.SurplusPercent.LT(sdk.NewInt(0)) {
			// search through each surplusList looking for a positive match
			for i := range surplusList {
				// if the surplusList item is a positive surplus (we can match with negative)
				if surplusList[i].SurplusPercent.GT(sdk.NewInt(0)) {
					// get the routes needed to swap from surplus item to currentSurplus
					routes, err := k.getOsmosisRoutes(ctx, surplusList[i].Holding.Token, currentSurplus.Holding.Token)
					if err != nil {
						return msgs, surplusList, err
					}
					// if the current surplus is greater then this surplusList item, add full surplusList item as swap
					// in msg list and then delete the surplusList item continuing to the next iteration in the loop
					if currentSurplus.BaseDenom.Amount.GT(surplusList[i].BaseDenom.Amount) {
						// compute the % needed to swap of the total currentSurplus by dividing the surplusList item in base denom by
						// the currentSurplus in base denom
						needToSwapPortion := surplusList[i].BaseDenom.Amount.Quo(currentSurplus.BaseDenom.Amount)
						// create the tokenin coin by using the full surplusList amount (since we are going to use all of it)
						tokenIn := surplusList[i].HoldingDenom
						// create the tokenin coin by multiplying the portion by currentSurlus we have to little of
						tokenOut := currentSurplus.HoldingDenom.Amount.Mul(needToSwapPortion)
						// create the swap msg with 2% slippage (potentially add this as fund param)
						msg, err := k.brokerKeeper.CreateOsmosisTrade(ctx, fundBrokerAddress, routes, tokenIn, tokenOut.Mul(sdk.NewInt(98)).Quo(sdk.NewInt(100)))
						if err != nil {
							return msgs, surplusList, err
						}
						// add the msg to the msgs list
						msgs = append(msgs, msg)
						// update the HoldingDenom of currentSurplus by subtracting what we are swapping in the HoldingDenom
						currentSurplus.HoldingDenom = currentSurplus.HoldingDenom.Sub(sdk.NewCoin(currentSurplus.HoldingDenom.Denom, tokenOut))
						// update the BaseDenom of currentSurplus by subtracting the total BaseDenom amount of surplusList item
						currentSurplus.BaseDenom = currentSurplus.BaseDenom.Sub(surplusList[i].BaseDenom)
						// remove the fully used surplusList item from list
						surplusList = remove(i, surplusList)
					}
					// if the current surplus is less then this surplusList item, make out the full swap for currentSurplus,
					// remove the amount swapped from in surplusList item and break the loop as this currentSurplus is done
					if currentSurplus.BaseDenom.Amount.LT(surplusList[i].BaseDenom.Amount) {
						// compute the % needed to swap of the total surplusList item by dividing the currentSurplus in base denom by
						// by the surplusList item in base denom
						needToSwapPortion := currentSurplus.BaseDenom.Amount.Quo(surplusList[i].BaseDenom.Amount)
						// create the tokein by using the full surplusList amount and multiplying it by the portion
						// (since we are not going to use all of it) and then creating a 2% slippage on it (potentially add this as fund param?)
						tokenIn := sdk.NewCoin(surplusList[i].HoldingDenom.Denom, surplusList[i].HoldingDenom.Amount.Mul(needToSwapPortion))
						// create the max amount out coin by using the full currentSurplus
						tokenOut := currentSurplus.HoldingDenom.Amount.Mul(sdk.NewInt(98)).Quo(sdk.NewInt(100))
						msg, err := k.brokerKeeper.CreateOsmosisTrade(ctx, fundBrokerAddress, routes, tokenIn, tokenOut)
						if err != nil {
							return msgs, surplusList, err
						}
						// add the msg to the msgs list
						msgs = append(msgs, msg)
						// update the HoldingDenom of surplusList item by subtracting what we are swapping out the HoldingDenom
						surplusList[i].HoldingDenom = surplusList[i].HoldingDenom.Sub(tokenIn)
						// update the BaseDenom of surplusList item by subtracting the total BaseDenom amount of currentSurplus item
						surplusList[i].BaseDenom = surplusList[i].BaseDenom.Sub(currentSurplus.BaseDenom)
						break
					}
					// if the current surplus equals the surplusList item, swap the current surplus over in full,
					// remove the surplusList item and then break the loop as this currentSurplus is done
					if currentSurplus.BaseDenom.Amount.Equal(surplusList[i].BaseDenom.Amount) {
						// create the tokenin coin by using the full currentSurplus
						tokenIn := currentSurplus.HoldingDenom
						// create the max amount out by using the full surplusList amount
						// (since we are going to use all of it) and then creating a 2% slippage on it (potentially add this as fund param?)
						tokenOut := surplusList[i].HoldingDenom.Amount.Mul(sdk.NewInt(98)).Quo(sdk.NewInt(100))
						// create the osmosis swap message with 2% slippage
						msg, err := k.brokerKeeper.CreateOsmosisTrade(ctx, fundBrokerAddress, routes, tokenIn, tokenOut)
						if err != nil {
							return msgs, surplusList, err
						}
						// add the msg to the msgs list
						msgs = append(msgs, msg)
						// remove the fully used surplusList item from list
						surplusList = remove(i, surplusList)
						break
					}
				}
			}
		}
	}
	// if leftovers in the currentSurplus, add to the left over slice
	if currentSurplus.BaseDenom.Amount.GT(sdk.NewInt(0)) {
		surplusList = append(surplusList, currentSurplus)
	}
	return msgs, surplusList, nil
}

// SendRebalanceTx sends one ICA tx for the fund with a list of swap msg's to rebalance
// the ETF. Each swap message will have multiple routes within it to swap to the needed
// rebalanced asset (see getRoutes above).
// For calculation of rebalances needed, each holding is converted to the base denom,
// then each holdings current weight in the base denom is subtracted from the expected composition.
// Then each needed composition that is positive (over owned) is matched with each negative composition (under owned)
// to create a swap message until no negative compositions exist.
//
// NOTE: currently very computationally expensive, look at ways to improve
func (k Keeper) SendRebalanceTx(ctx sdk.Context, fund types.Fund) error {
	// create slice to hold the trade messages
	//msgs := []*osmosisgammtypes.MsgSwapExactAmountIn{}
	// slice to store the extra surpluses
	surplusList := []Surplus{}
	// slice to store the amount of each holding in the base denom as a coin type
	allHoldingsInBaseDenom := []sdk.Coin{}
	// slice to store all the holdings in base denom amounts so we can add them to get total for etf
	allHoldingsInBaseDenomAmount := []sdk.Int{}
	// Map for holding all the messages for each broker to send later
	msgs := make(map[string][]*osmosisgammtypes.MsgSwapExactAmountIn)

	for _, holding := range fund.Holdings {
		broker, found := k.brokerKeeper.GetBroker(ctx, holding.BrokerId)
		if !found {
			return sdkerrors.Wrap(types.ErrWrongBroker, fmt.Sprintf("broker %s not found", holding.BrokerId))
		}
		// get the ica account for the fund on the broker chain
		portID, err := icatypes.NewControllerPortID(fund.Address)
		if err != nil {
			return err
		}
		fundBrokerAddress, found := k.brokerKeeper.GetBrokerAccount(ctx, broker.ConnectionId, portID)
		if !found {
			return sdkerrors.Wrapf(brokertypes.ErrIBCAccountNotExist, "failed to find ica account for owner %s on connection %s and port %s", fund.Address, broker.ConnectionId, portID)
		}
		// get the account balances for the fund account on the broker chain
		balances, err := k.brokerKeeper.GetOsmosisBalance(ctx, fundBrokerAddress)
		if err != nil {
			return err
		}
		// get amount of holding token from balances
		amount := balances.GetCoins().AmountOf(holding.Token)
		// get the price of the asset in base denom
		priceInBaseDenom, err := k.brokerKeeper.CalculateOsmosisSpotPrice(ctx, holding.PoolId, holding.Token, broker.BaseDenom)
		if err != nil {
			return err
		}
		// calculate the amount held of holding in base denom by taking amount and multiplying by price in base denom
		amountInBaseDenom := amount.Quo(priceInBaseDenom.RoundInt())

		// add the pricing/amounts to the slices needed
		allHoldingsInBaseDenomAmount = append(allHoldingsInBaseDenomAmount, amountInBaseDenom)
		allHoldingsInBaseDenom = append(allHoldingsInBaseDenom, sdk.NewCoin(broker.BaseDenom, amountInBaseDenom))
	}

	allHoldingsInBaseDenomCoins := sdk.NewCoins(allHoldingsInBaseDenom...)
	totalInBaseDenom := sumInts(allHoldingsInBaseDenomAmount)

	// for loop that determines the surpluses (both positive and negative) from balances and holdings in etf
	for _, holding := range fund.Holdings {
		broker, found := k.brokerKeeper.GetBroker(ctx, holding.BrokerId)
		if !found {
			return sdkerrors.Wrap(types.ErrWrongBroker, fmt.Sprintf("broker %s not found", holding.BrokerId))
		}
		// get the price of the asset in base denom
		priceInBaseDenom, err := k.brokerKeeper.CalculateOsmosisSpotPrice(ctx, holding.PoolId, holding.Token, broker.BaseDenom)
		if err != nil {
			return err
		}
		// use some math to get the current composition % for this holding in the fund
		// (holding in base denom / total in base denom)
		currentComposition := allHoldingsInBaseDenomCoins.AmountOf(holding.Token).Quo(totalInBaseDenom)
		// get the surplus composition % by subtracting the current composition % from what its supposed to be
		overUnderCompPerc := currentComposition.Sub(sdk.NewInt(holding.Percent / 100))
		// get the surplus in the base denom by dividing the % surplus comp by the current comp % to
		// yield the % of base denom in surplus and then muliplying it by the current balance
		// of this base denom in the fund account
		surplus := Surplus{
			BaseDenom:      sdk.NewCoin(fund.BaseDenom, overUnderCompPerc.Abs().Quo(currentComposition).Mul(allHoldingsInBaseDenomCoins.AmountOf(holding.Token))),
			HoldingDenom:   sdk.NewCoin(holding.Token, overUnderCompPerc.Abs().Quo(currentComposition).Mul(allHoldingsInBaseDenomCoins.AmountOf(holding.Token)).Mul(priceInBaseDenom.RoundInt())),
			Holding:        holding,
			SurplusPercent: overUnderCompPerc,
		}

		// get the new surplus msgs
		newMsgs, _, err := k.HandleSurplus(ctx, fund, holding, msgs[holding.BrokerId], surplusList, surplus)
		if err != nil {
			return err
		}
		// append new msgs for broker into broker map
		msgs[holding.BrokerId] = append(msgs[holding.BrokerId], newMsgs...)
	}

	// send msgs for each broker
	for brokerId, msg := range msgs {
		broker, found := k.brokerKeeper.GetBroker(ctx, brokerId)
		if !found {
			return sdkerrors.Wrap(types.ErrWrongBroker, fmt.Sprintf("broker %s not found", brokerId))
		}
		// send the trade through ICA
		_, err := k.brokerKeeper.SendOsmosisTrades(ctx, msg, fund.Address, broker.ConnectionId)
		if err != nil {
			return err
		}
	}

	return nil
}
