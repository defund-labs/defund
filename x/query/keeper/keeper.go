package keeper

import (
	"fmt"
	"strconv"

	"github.com/tendermint/tendermint/libs/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	etftypes "github.com/defund-labs/defund/x/etf/types"
	"github.com/defund-labs/defund/x/query/types"

	json "github.com/tendermint/tendermint/libs/json"

	liquiditytypes "github.com/tendermint/liquidity/x/liquidity/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey

		accountKeeper types.AccountKeeper
		etfKeeper     types.EtfKeeper
		brokerKeeper  types.BrokerKeeper
	}
)

type PoolsKey struct {
}

type BalanceKey struct {
	Address string `json:"address"`
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,

	accountkeeper types.AccountKeeper,
	etfkeeper types.EtfKeeper,
	brokerkeeper types.BrokerKeeper,

) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,

		accountKeeper: accountkeeper,
		etfKeeper:     etfkeeper,
		brokerKeeper:  brokerkeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) NewQueryAddress(id uint64) sdk.AccAddress {
	key := append([]byte("query"), sdk.Uint64ToBigEndian(id)...)
	return address.Module("query", key)
}

func (k Keeper) CreateInterqueryRequest(ctx sdk.Context, storeid string, path string, key []byte, timeoutheight uint64, clientid string) error {
	var queryModuleAddress authtypes.ModuleAccountI
	if k.accountKeeper.GetModuleAccount(ctx, "query") == nil {
		queryAddress := k.NewQueryAddress(1)
		queryModuleAddress = authtypes.NewModuleAccount(
			authtypes.NewBaseAccountWithAddress(
				queryAddress,
			),
			"query",
		)
		k.accountKeeper.SetModuleAccount(ctx, queryModuleAddress)
	} else {
		queryModuleAddress = k.accountKeeper.GetModuleAccount(ctx, "query")
	}
	interquery := types.Interquery{
		Creator:       queryModuleAddress.GetAddress().String(),
		Storeid:       storeid,
		Path:          path,
		Key:           key,
		TimeoutHeight: timeoutheight,
		ClientId:      clientid,
	}
	k.SetInterquery(ctx, interquery)

	k.Logger(ctx).Info(fmt.Sprintf("Interquery request for path %s on clientid of %s has been initiated", path, clientid))

	return nil
}

// GetInterqueryResult gets an interquery based on storeid from store
func (k Keeper) GetInterqueryResultFromStore(ctx sdk.Context, storeid string) ([]byte, error) {
	results := k.GetAllInterqueryResult(ctx)
	for _, result := range results {
		if result.Storeid == storeid {
			return result.Data, nil
		}
	}
	return nil, status.Errorf(codes.InvalidArgument, "could not find interquery with storeid: %s", storeid)
}

// Helper function to be implemented in end blocker to interchain query pools on gravity dex (Cosmos Hub)
func (k Keeper) QueryGravityDex(ctx sdk.Context) error {
	path := "custom/liquidity/liquidityPools/"
	clientid := "07-tendermint-0"
	keyRaw := PoolsKey{}
	key, err := json.Marshal(keyRaw)
	timeoutHeight := uint64(ctx.BlockHeight() + 10)
	storeid := fmt.Sprintf("gdex-pools")

	err = k.CreateInterqueryRequest(ctx, storeid, path, key, timeoutHeight, clientid)
	if err != nil {
		return err
	}
	return nil
}

// QueryPoolAccount creates and interchain query for the balance of an address for a relayer to submit as a interquery result
func (k Keeper) QueryPoolAccount(ctx sdk.Context, address string) error {
	path := "custom/bank/all_balances/"
	clientid := "07-tendermint-0"
	keyRaw := BalanceKey{address}
	key, err := json.Marshal(keyRaw)
	if err != nil {
		return err
	}
	timeoutHeight := uint64(ctx.BlockHeight() + 10)
	storeid := fmt.Sprintf("balance-%d", address)

	err = k.CreateInterqueryRequest(ctx, storeid, path, key, timeoutHeight, clientid)
	if err != nil {
		return err
	}
	return nil
}

// QueryAllPools queries all pool accounts from most recent pools store
func (k Keeper) QueryAllPools(ctx sdk.Context, broker string) error {
	recentPools, err := k.GetInterqueryPoolAll(ctx, broker)
	// Log error if error returns from query. Do not want to panic application. Just log
	if err != nil {
		ctx.Logger().Error(err.Error())
	}
	// Take the query with the most recent height (first in sorted slice)
	if len(recentPools) > 0 {
		for _, pool := range recentPools {
			err := k.QueryPoolAccount(ctx, pool.ReserveAccountAddress)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// GetInterqueryPoolAll gets all the most recent pools from the broker specified in the interquery result store
func (k Keeper) GetInterqueryPoolAll(ctx sdk.Context, broker string) ([]liquiditytypes.Pool, error) {
	storeid := fmt.Sprintf("%s-pools", broker)
	rawpools, found := k.GetInterqueryResult(ctx, storeid)
	if !found {
		return nil, fmt.Errorf("pools interquery not found: (%s)", storeid)
	}
	pools := []liquiditytypes.Pool{}
	json.Unmarshal(rawpools.Data, &pools)
	return pools, nil
}

// GetInterqueryPool gets the most recent pool details of a specific pool in an interquery result store by looping through all pools
func (k Keeper) GetInterqueryPool(ctx sdk.Context, broker string, poolid string) (liquiditytypes.Pool, error) {
	pools, err := k.GetInterqueryPoolAll(ctx, broker)
	if err != nil {
		ctx.Logger().Error(err.Error())
	}
	for _, pool := range pools {
		if strconv.FormatUint(pool.Id, 10) == poolid {
			return pool, nil
		}
	}
	return liquiditytypes.Pool{}, sdkerrors.Wrapf(types.ErrInvalidPool, "Pool not found (%s)", poolid)
}

// GetBalance gets the most recent balance/holdings of an account in an interquery result store
func (k Keeper) GetBalance(ctx sdk.Context, account string) ([]sdk.Coin, error) {
	storeid := fmt.Sprintf("balance-%s", account)
	rawbalance, found := k.GetInterqueryResult(ctx, storeid)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrInterqueryNotFound, "Interquery not found (%s)", storeid)
	}
	balances := []sdk.Coin{}
	err := json.Unmarshal(rawbalance.Data, &balances)
	if err != nil {
		return []sdk.Coin{}, sdkerrors.Wrapf(types.ErrMarshallingError, "Marshalling error for interquery: (%s)", storeid)
	}
	return balances, nil
}

// CheckHoldings checks to make sure the specified holdings and the pool for each holding are valid
// by checking the interchain queried pools for the broker specified
func (k Keeper) CheckHoldings(ctx sdk.Context, broker string, holdings []etftypes.Holding) error {
	percentCheck := uint64(0)
	for _, holding := range holdings {
		percentCheck = percentCheck + uint64(holding.Percent)
		pool, err := k.GetInterqueryPool(ctx, broker, holding.PoolId)
		if err != nil {
			return err
		}
		// Checks to see if the holding pool contains the holding token specified and if not returns error
		if !contains(pool.ReserveCoinDenoms, holding.Token) {
			return sdkerrors.Wrapf(types.ErrInvalidDenom, "invalid denom (%s)", holding.Token)
		}
	}
	// Make sure all fund holdings add up to 100%
	if percentCheck != uint64(100) {
		return sdkerrors.Wrapf(types.ErrPercentComp, "percent composition must add up to 100%")
	}
	return nil
}

// EndBlockerRun creates all the repeated interqueries
func (k Keeper) EndBlockerRun(ctx sdk.Context) error {
	// Run every 10th block
	if ctx.BlockHeight()%50 == 0 {
		// Add gravity dex interquery
		err := k.QueryGravityDex(ctx)
		if err != nil {
			ctx.Logger().Debug(fmt.Sprintf("Error Creating Cosmos GDex Pool Interquery: %s", err))
		}
		// Add interquery for all pool account balances on gdex
		err = k.QueryAllPools(ctx, "gdex")
		if err != nil {
			ctx.Logger().Debug(fmt.Sprintf("Error Creating Interquery For All Pool Accounts: %s", err))
		}
	}
	return nil
}
