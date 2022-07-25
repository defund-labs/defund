package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateQueryOsmosisBalances queries all balances for each fund on Osmosis
func (k Keeper) CreateQueryOsmosisBalances(ctx sdk.Context) {
	funds := k.GetAllFund(ctx)
	for _, fund := range funds {
		portID, err := icatypes.NewControllerPortID(fund.Address)
		if err != nil {
			ctx.Logger().Debug(err.Error())
		}
		addr, found := k.icaControllerKeeper.GetInterchainAccountAddress(ctx, fund.ConnectionId, portID)
		if !found {
			err := status.Errorf(codes.NotFound, "no account found for portID %s", portID)
			ctx.Logger().Debug(err.Error())
		}
		k.brokerKeeper.CreateQueryOsmosisBalance(ctx, addr)
	}
}

// CreateDefundQueries creates all the queries for the ETF module. Run in end blocker
func (k Keeper) CreateETFQueries(ctx sdk.Context) {
	// Add Osmosis broker interquery for all pools
	k.CreateQueryOsmosisBalances(ctx)
}
