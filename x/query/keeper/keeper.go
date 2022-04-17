package keeper

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/tendermint/tendermint/libs/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/defund-labs/defund/x/query/types"

	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
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

// Helper function to be implemented in end blocker to interchain query pools on gravity dex (Cosmos Hub)
func (k Keeper) QueryGravityDex(ctx sdk.Context) error {
	path := "custom/liquidity/liquidityPools/"
	clientid := "07-tendermint-0"
	keyRaw := PoolsKey{}
	key, err := json.Marshal(keyRaw)
	heightStr := strconv.FormatInt(ctx.BlockHeight(), 10)
	timeoutHeight := uint64(ctx.BlockHeight() + 10)
	storeid := fmt.Sprintf("CosmosPools-%s", heightStr)

	err = k.CreateInterqueryRequest(ctx, storeid, path, key, timeoutHeight, clientid)
	if err != nil {
		return err
	}
	return nil
}

// Helper function that creates an interquery for an account balance on Cosmos with the accountType as part of the store id
func (k Keeper) QueryFundAccounts(ctx sdk.Context, address string, accountType string) error {
	path := "custom/bank/all_balances/"
	clientid := "07-tendermint-0"
	keyRaw := BalanceKey{address}
	key, err := json.Marshal(keyRaw)
	if err != nil {
		return err
	}
	timeoutHeight := uint64(ctx.BlockHeight() + 10)
	storeid := fmt.Sprintf("%s-Balance-%s", accountType, address)

	err = k.CreateInterqueryRequest(ctx, storeid, path, key, timeoutHeight, clientid)
	if err != nil {
		return err
	}
	return nil
}

// Helper function that gets all funds from store and creates an interquery for all ICA account balances
// associated with the fund on the broker chains fund ICA account.
func (k Keeper) QueryAllAccounts(ctx sdk.Context) error {
	funds := k.etfKeeper.GetAllFund(ctx)
	for _, fund := range funds {
		portID, err := icatypes.NewControllerPortID(fund.Address)
		if err != nil {
			return status.Errorf(codes.InvalidArgument, "could not find account: %s", err)
		}

		addr, found := k.brokerKeeper.GetBrokerAccount(ctx, fund.ConnectionId, portID)
		if !found {
			return status.Errorf(codes.NotFound, "no account found for portID %s", portID)
		}
		err = k.QueryFundAccounts(ctx, addr, "FundAccount")
		if err != nil {
			return status.Errorf(codes.InvalidArgument, "could not find account: %s", err)
		}
	}

	return nil
}

func (k Keeper) EndBlockerRun(ctx sdk.Context) error {
	// Run every 10th block
	if ctx.BlockHeight()%10 == 0 {
		// Add gravity dex interquery
		err := k.QueryGravityDex(ctx)
		if err != nil {
			ctx.Logger().Error(fmt.Sprintf("Error Creating Cosmos GDex Pool Interquery: %s", err))
		}
		// Add gravity dex account balances interquery for all funds
		err = k.QueryAllAccounts(ctx)
		if err != nil {
			ctx.Logger().Error(fmt.Sprintf("Error Creating Cosmos GDex Account Balance Interquery: %s", err))
		}
	}
	return nil
}
