package broker

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/defund-labs/defund/x/broker/keeper"
	"github.com/defund-labs/defund/x/broker/types"
)

// Create a list of all pool ids supported on osmosis at init/genesis
var poolsOsmosis = []uint64{
	1, 678, 704, 712, 497, 674, 604, 9, 498, 584, 3, 10, 601, 2, 722, 611, 719, 585, 738, 13,
	4, 482, 481, 6, 577, 5, 463, 629, 641, 690, 15, 461, 560, 586, 587, 42, 600, 627, 608, 571,
	631, 548, 7, 605, 572, 648, 606, 643, 8, 597, 619, 553, 625, 602, 618, 574, 578, 651, 626, 573,
	22, 555, 637, 681, 464, 645, 644, 596, 547, 616, 558, 621, 613, 197, 679, 617, 670, 612, 638, 561,
	567, 649, 732, 653, 633, 557, 706, 662, 615, 701, 565, 669, 562, 592, 693, 151, 183, 695, 726, 673,
	549, 716, 624, 731, 718, 642, 721, 640, 734, 713, 725, 710, 737, 729, 700, 707, 717, 676,
	579, 682, 580, 730,
}

// AddOsmosisToBrokers adds Osmosis as a broker to state manually
func AddOsmosisToBrokers(ctx sdk.Context, k keeper.Keeper) error {
	var pools []*types.Pool

	for _, pool := range poolsOsmosis {
		addPool := types.Pool{
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
