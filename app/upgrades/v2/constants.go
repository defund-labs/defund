package v2

import (
	"github.com/CosmWasm/wasmd/x/wasm"

	"github.com/defund-labs/defund/app/upgrades"

	store "github.com/cosmos/cosmos-sdk/store/types"
)

const UpgradeName = "v2"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added: []string{wasm.ModuleName},
	},
}
