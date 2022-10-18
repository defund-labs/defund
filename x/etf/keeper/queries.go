package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	icatypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/types"
	brokertypes "github.com/defund-labs/defund/x/broker/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateBalances queries all balances for each broker for each funds holdings
func (k Keeper) CreateBalances(ctx sdk.Context) {
	funds := k.GetAllFund(ctx)
	for _, fund := range funds {
		for _, holding := range fund.Holdings {
			switch holding.BrokerId {
			case "osmosis":
				// get the broker
				broker, found := k.brokerKeeper.GetBroker(ctx, holding.BrokerId)
				if !found {
					err := sdkerrors.Wrap(brokertypes.ErrBrokerNotFound, fmt.Sprintf("broker %s not found for holding %s", holding.BrokerId, holding.Token))
					ctx.Logger().Debug(err.Error())
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

				k.brokerKeeper.CreateQueryOsmosisBalance(ctx, addr)
			}
		}
	}
}

// CreateDefundQueries creates all the queries for the ETF module. Run in end blocker
func (k Keeper) CreateETFQueries(ctx sdk.Context) {
	// Add broker interquery for all fund balance accounts
	k.CreateBalances(ctx)
}
