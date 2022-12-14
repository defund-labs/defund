package keeper

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	icatypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/types"
	"github.com/defund-labs/defund/x/etf/types"
	querytypes "github.com/defund-labs/defund/x/query/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// OnFundBalanceSubmissionCallback is called on the submission of an interquery result. It runs if the
// interquery is a /store/bank/key. It takes the raw balance query for each denom from the remote
// broker chains and adjusts the balances store for the fund on Defund to reflect the updated balance
func (k Keeper) OnFundBalanceSubmissionCallback(ctx sdk.Context, result *querytypes.InterqueryResult) error {
	// store id for balance should look like balance:{fundSymbol}:{brokerId}:{address}:{denom}
	var account string
	var brokerId string
	var fundSymbol string
	sep := strings.Split(result.Storeid, ":")
	if sep[0] == "balance" && result.Proved && result.Success {
		fundSymbol = sep[1]
		brokerId = sep[2]
		account = sep[3]
		// denom placeholder
		_ = sep[4]

		fund, err := k.GetFundBySymbol(ctx, fundSymbol)
		if err != nil {
			return err
		}
		broker, found := k.brokerKeeper.GetBroker(ctx, brokerId)
		if !found {
			return status.Errorf(codes.NotFound, "broker %s not found", broker)
		}
		portID, err := icatypes.NewControllerPortID(fund.Address)
		if err != nil {
			return err
		}

		// check to ensure the query balance address matches the ICA for this fund + broker
		icaAddr, found := k.icaControllerKeeper.GetInterchainAccountAddress(ctx, broker.ConnectionId, portID)
		if !found {
			return status.Errorf(codes.NotFound, "no ICA account found for portID %s", portID)
		}
		if icaAddr != account {
			return status.Errorf(codes.InvalidArgument, "account mismatch. ICA fund account for fund %s on broker %s is %s (received %s)", fundSymbol, brokerId, icaAddr, account)
		}

		var coin sdk.Coin
		err = coin.Unmarshal(result.Data)
		if err != nil {
			return err
		}

		// reset fund with updated balances
		fund.SetBalances(brokerId, account, coin)
		k.SetFund(ctx, fund)
	}

	return nil
}

func (k Keeper) EtfQueryCleanerBeginBlocker(ctx sdk.Context) {
	results := k.queryKeeper.GetAllInterqueryResult(ctx)
	for _, result := range results {
		err := k.OnFundBalanceSubmissionCallback(ctx, &result)
		if err != nil {
			retErr := sdkerrors.Wrapf(types.ErrBeginBlocker, "EtfQueryCleanerBeginBlocker: error running OnFundBalanceSubmissionCallback (%s)", err.Error())
			k.Logger(ctx).Error(retErr.Error())
			continue
		}
		// clean the store result
		k.queryKeeper.RemoveInterqueryResult(ctx, result.Storeid)
	}
}
