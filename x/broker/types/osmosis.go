package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	osmosisbalancertypes "github.com/osmosis-labs/osmosis/v8/x/gamm/pool-models/balancer"
	osmosisgammtypes "github.com/osmosis-labs/osmosis/v8/x/gamm/types"
)

type BalancerPoolPretty struct {
	Address            sdk.AccAddress                  `json:"address" yaml:"address"`
	Id                 uint64                          `json:"id" yaml:"id"`
	PoolParams         osmosisbalancertypes.PoolParams `json:"pool_params" yaml:"pool_params"`
	FuturePoolGovernor string                          `json:"future_pool_governor" yaml:"future_pool_governor"`
	TotalWeight        sdk.Dec                         `json:"total_weight" yaml:"total_weight"`
	TotalShares        sdk.Coin                        `json:"total_shares" yaml:"total_shares"`
	PoolAssets         []osmosisgammtypes.PoolAsset    `json:"pool_assets" yaml:"pool_assets"`
}
