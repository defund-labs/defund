package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/defund-labs/defund/x/etf/types"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper Keeper
}

// NewMigrator returns a new Migrator.
func NewMigrator(keeper Keeper) Migrator {
	return Migrator{keeper: keeper}
}

// Migrate1to2 migrates from version 1 to 2. Symbolic upgrade since version was set to 2 initially
func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	return nil
}

// Migrate2to3 migrates x/bank storage from version 2 to 3.
func (m Migrator) Migrate2to3(ctx sdk.Context) error {
	store := ctx.KVStore(m.keeper.storeKey)
	interqueryResultStore := prefix.NewStore(store, []byte(types.FundKeyPrefix))

	iterator := interqueryResultStore.Iterator(nil, nil)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Fund
		m.keeper.cdc.MustUnmarshal(iterator.Value(), &val)
		newFund := types.Fund{
			Symbol:              val.Symbol,
			Address:             val.Address,
			Name:                val.Name,
			Description:         val.Description,
			Shares:              val.Shares,
			Holdings:            val.Holdings,
			Rebalance:           val.Rebalance,
			BaseDenom:           val.BaseDenom,
			StartingPrice:       val.StartingPrice,
			Creator:             val.Creator,
			Rebalancing:         val.Rebalancing,
			LastRebalanceHeight: val.LastRebalanceHeight,
			Balances: &types.FundBalances{
				Osmosis: types.Balances{
					Address:  "",
					Balances: []*sdk.Coin{},
				},
			},
			FundType: types.FundType_PASSIVE,
		}
		m.keeper.SetFund(ctx, newFund)
	}
	return nil
}
