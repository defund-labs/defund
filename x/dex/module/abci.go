package dex

import (
	"context"
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"defund/x/dex/keeper"
	"defund/x/dex/types"
)

func BeginBlocker(ctx context.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	k.DeleteOutdatedRequests(ctx.(sdk.Context))
}

func EndBlocker(ctx context.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)

	params := k.GetParams(ctx.(sdk.Context))
	if ctx.(sdk.Context).BlockHeight()%int64(params.BatchSize) == 0 {
		k.ExecuteRequests(ctx.(sdk.Context))
	}
}
