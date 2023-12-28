package tracking

import (
	"time"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/MonikaCat/archway/v5/x/tracking/keeper"
	"github.com/MonikaCat/archway/v5/x/tracking/types"
)

// EndBlocker modifies tracked transactions using tracked contract operations.
func EndBlocker(ctx sdk.Context, k keeper.Keeper) []abci.ValidatorUpdate {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)

	k.FinalizeBlockTxTracking(ctx)

	return []abci.ValidatorUpdate{}
}
