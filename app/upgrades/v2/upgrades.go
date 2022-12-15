package v2

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	etfkeeper "github.com/defund-labs/defund/x/etf/keeper"
	"github.com/defund-labs/defund/x/etf/types"
)

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	etfkeeper *etfkeeper.Keeper,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		newVM, err := mm.RunMigrations(ctx, configurator, vm)
		if err != nil {
			return newVM, err
		}

		// Change fund states to account for new type, contract and balance fields
		ctx.Logger().Info("Updating ETF stores for in place upgrade v2")
		funds := etfkeeper.GetAllFund(ctx)
		for i := range funds {
			etfkeeper.SetFund(ctx, types.Fund{
				Symbol:      funds[i].Symbol,
				Address:     funds[i].Address,
				Name:        funds[i].Name,
				Description: funds[i].Description,
			})
		}

		return newVM, err
	}
}
