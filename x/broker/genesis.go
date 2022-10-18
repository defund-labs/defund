package broker

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/defund-labs/defund/x/broker/keeper"
	"github.com/defund-labs/defund/x/broker/types"
)

// Create a list of all pool ids supported on osmosis at init/genesis
// Pool 678 is USCD/OSMO so base denom pool option.
var poolsOsmosis = []uint64{
	1, 497, 674, 604, 9, 498, 584, 3, 10, 601, 2, 611, 585, 13, 4, 482, 481, 6, 577, 5, 463,
	629, 641, 15, 461, 560, 586, 587, 42, 600, 627, 608, 571, 631, 548, 7, 605, 572, 648,
	606, 643, 8, 597, 619, 553, 625, 602, 618, 574, 578, 651, 626, 573, 22, 555, 637, 464,
	645, 644, 596, 547, 616, 558, 621, 613, 197, 617, 670, 612, 638, 561, 567, 649, 653,
	633, 557, 662, 615, 565, 562, 592, 151, 183, 673, 549, 624, 642, 678,
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
	// Set all the brokers Defund supports at genesis
	AddOsmosisToBrokers(ctx, k)
	// set the broker params
	k.SetParams(ctx, genState.Params)
	// this line is used by starport scaffolding # genesis/module/init
}

// ExportGenesis returns the brokers module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx, types.ParamsKey)

	genesis.Brokers = k.GetAllBrokers(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
