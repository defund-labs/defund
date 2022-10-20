package keeper

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	icatypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/types"
	"github.com/defund-labs/defund/x/etf/types"
	querytypes "github.com/defund-labs/defund/x/query/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func removeCoin(s []*sdk.Coin, i int) []*sdk.Coin {
	coins := []*sdk.Coin{}
	for index := range s {
		if index != i {
			coins = append(coins, s[i])
		}
	}
	return coins
}

// OnFundBalanceSubmissionCallback is called on the submission of an interquery result. It runs if the
// interquery is a /store/bank/key. It takes the raw balance query for each denom from the remote
// broker chains and adjusts the balances store for the fund on Defund to reflect the updated balance
func (k Keeper) OnFundBalanceSubmissionCallback(ctx sdk.Context, result *querytypes.InterqueryResult) error {
	// store id for balance should look like balance:{fundSymbol}:{brokerId}:{address}:{denom}
	var denom string
	var account string
	var brokerId string
	var fundSymbol string
	sep := strings.Split(result.Storeid, ":")
	if sep[0] == "balance" && result.Proved && result.Success {
		fundSymbol = sep[1]
		brokerId = sep[2]
		account = sep[3]
		denom = sep[4]

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

		bals := types.Balances{
			Balances: []*sdk.Coin{},
		}

		if _, ok := fund.Balances[account]; ok {
			// remove this balance from the list of balances
			for i := range bals.Balances {
				if bals.Balances[i].Denom == denom {
					bals.Balances = removeCoin(bals.Balances, i)
					break
				}
			}
		} else {
			fund.Balances[account] = &bals
		}

		// lets add the balance to list of balances for broker
		var coin sdk.Coin
		err = coin.Unmarshal(result.Data)
		if err != nil {
			return err
		}
		bals.Balances = append(bals.Balances, &coin)

		// set the new balances for this broker
		fund.Balances[account].Balances = bals.Balances
		// reset fund with updated balances
		k.SetFund(ctx, fund)
	}

	return nil
}
