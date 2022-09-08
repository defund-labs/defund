package broker

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/defund-labs/defund/x/broker/keeper"
	"github.com/defund-labs/defund/x/broker/types"
)

// Create a list of all pool ids supported on osmosis at init/genesis
var poolsOsmosis = []uint64{
	1, 560, 497, 604, 2, 561, 584, 481, 9, 605, 11, 666, 577, 557, 621, 3,
	619, 631, 613, 463, 602, 197, 586, 601, 597, 573, 7, 553, 571, 5, 626,
	627, 608, 15, 629, 637, 42, 624, 634, 633,
}

// AddOsmosisToBrokers adds Osmosis as a broker to state manually
func AddOsmosisToBrokers(ctx sdk.Context, k keeper.Keeper) error {
	var pools []*types.Source

	for _, pool := range poolsOsmosis {
		addPool := types.Source{
			PoolId:       pool,
			InterqueryId: fmt.Sprintf("%s-%d", "osmosis", pool),
			Status:       "active",
		}
		pools = append(pools, &addPool)
	}

	broker := types.Broker{
		Id:        "osmosis",
		Pools:     pools,
		BaseDenom: "uosmo",
		Status:    "inactive",
	}

	k.SetBroker(ctx, broker)
	return nil
}

// InitGenesis initializes the brokers module's state from a provided genesis
// state via a genesis file.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the brokers Defund supports at genesis
	AddOsmosisToBrokers(ctx, k)
	// this line is used by starport scaffolding # genesis/module/init
}

// ExportGenesis returns the brokers module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.Brokers = k.GetAllBrokers(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
