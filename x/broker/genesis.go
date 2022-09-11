package broker

import (
	"encoding/base64"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/defund-labs/defund/x/broker/keeper"
	"github.com/defund-labs/defund/x/broker/types"
)

// Create a list of all pool ids supported on osmosis at init/genesis
var poolsOsmosis = []uint64{
	1, 497, 674, 604, 9, 498, 584, 3, 10, 601, 2, 611, 585, 13, 4, 482, 481, 6, 577, 5, 463,
	629, 641, 15, 461, 560, 586, 587, 42, 600, 627, 608, 571, 631, 548, 7, 605, 572, 648,
	606, 643, 8, 597, 619, 553, 625, 602, 618, 574, 578, 651, 626, 573, 22, 555, 637, 464,
	645, 644, 596, 547, 616, 558, 621, 613, 197, 617, 670, 612, 638, 561, 567, 649, 653,
	633, 557, 662, 615, 565, 562, 592, 151, 183, 673, 549, 624, 642,
}

// add the list of bytes we must append to properly unmarshal
var poolsOsmosisAppend = []string{
	"CtIC", "Cs0C", "Cs0C", "CsIC", "CtUC", "CokD", "CtEC", "Cs8C", "CpQD", "Cs8C", "Co4C",
	"CpED", "Co8D", "CosD", "Co4D", "CosD", "Cs4C", "CpID", "Cs4C", "CtQC", "Cs8C", "CtoC",
	"CsIC", "CtEC", "Co0D", "CtQC", "Cs8C", "CowD", "Cs0C", "CvwC", "Cs8C", "Cr0C", "Cs8C",
	"CssC", "CosD", "CtEC", "CtcC", "CowD", "Co0D", "CpUD", "CsAC", "Co8D", "Ct0C", "CtAC",
	"CtwC", "CsIC", "CtAC", "Co4D", "Co0D", "CowD", "CssC", "CssC", "Cs4C", "Co0D", "CpYD",
	"CswC", "CooD", "CrwC", "CvoC", "CpsD", "CosD", "CvoC", "Co0D", "CsIC", "CssC", "Cs8C",
	"Co0D", "Cs0C", "CpAD", "CooD", "CtEC", "CpAD", "CsUC", "CooD", "CswC", "CswC", "CooC",
	"CvwC", "CooD", "CpID", "Co4D", "CsMC", "Co0D", "Cs4C", "CoYD", "CsgC", "CsEC", "CscC",
	"CocD", "CooD",
}

// AddOsmosisToBrokers adds Osmosis as a broker to state manually
func AddOsmosisToBrokers(ctx sdk.Context, k keeper.Keeper) error {
	var pools []*types.Source

	for i := range poolsOsmosis {
		app, err := base64.StdEncoding.DecodeString(poolsOsmosisAppend[i])
		if err != nil {
			return err
		}
		addPool := types.Source{
			PoolId:       poolsOsmosis[i],
			InterqueryId: fmt.Sprintf("%s-%d", "osmosis", poolsOsmosis[i]),
			Status:       "active",
			Append:       app,
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
