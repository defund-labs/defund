package broker

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/defund-labs/defund/x/broker/keeper"
	"github.com/defund-labs/defund/x/broker/types"
)

// Create a list of all pool ids supported on osmosis at init/genesis
var poolsOsmosis = []uint64{
	1, 2, 560, 604, 561, 497, 584, 3, 9, 601, 605, 627, 481, 5, 629, 15, 619, 608, 7, 463, 571, 42, 586,
	577, 621, 573, 553, 602, 197, 597, 651, 645, 557, 626, 631, 613, 647, 624, 637, 11, 641, 822,
}

// AddOsmosisToBrokers adds Osmosis as a broker to state manually
func AddOsmosisToBrokers(ctx sdk.Context, k keeper.Keeper) error {
	var pools []*types.Source

	for i := range poolsOsmosis {
		addPool := types.Source{
			PoolId:       poolsOsmosis[i],
			InterqueryId: fmt.Sprintf("%s-%d", "osmosis", poolsOsmosis[i]),
			Status:       "active",
		}
		pools = append(pools, &addPool)
	}

	broker := types.Broker{
		Id:     "osmosis",
		Pools:  pools,
		Status: "inactive",
	}

	k.SetBroker(ctx, broker)
	return nil
}

// InitGenesis initializes the brokers module's state from a provided genesis
// state via a genesis file.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// set the broker params
	k.SetParams(ctx, genState.Params)
	// Set all the brokers Defund supports at genesis
	AddOsmosisToBrokers(ctx, k)
	// this line is used by starport scaffolding # genesis/module/init
}

// ExportGenesis returns the brokers module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.Brokers = k.GetAllBrokers(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
