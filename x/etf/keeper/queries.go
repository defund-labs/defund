package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	icatypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/types"
	brokertypes "github.com/defund-labs/defund/x/broker/types"
	"github.com/defund-labs/defund/x/etf/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateBalances creates interqueries all balances for each broker of a fund
func (k Keeper) CreateBalances(ctx sdk.Context, fund types.Fund) error {
	for _, holding := range fund.Holdings {
		switch holding.BrokerId {
		case "osmosis":
			// get the broker
			broker, found := k.brokerKeeper.GetBroker(ctx, holding.BrokerId)
			if !found {
				return sdkerrors.Wrap(brokertypes.ErrBrokerNotFound, fmt.Sprintf("broker %s not found for holding %s", holding.BrokerId, holding.Token))
			}

			portID, err := icatypes.NewControllerPortID(fund.Address)
			if err != nil {
				return err
			}
			addr, found := k.icaControllerKeeper.GetInterchainAccountAddress(ctx, broker.ConnectionId, portID)
			if !found {
				return status.Errorf(codes.NotFound, "no account found for portID %s on connection %s", portID, broker.ConnectionId)
			}

			err = k.brokerKeeper.CreateQueryOsmosisBalance(ctx, fund.Symbol, addr, holding.Token)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
