package slashing

import (
	"time"

	"emoney/x/slashing/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MaxEvidenceAge - max age for evidence
func (k Keeper) MaxEvidenceAge(ctx sdk.Context) (res time.Duration) {
	k.paramspace.Get(ctx, types.KeyMaxEvidenceAge, &res)
	return
}

// SignedBlocksWindowDuration - sliding window for downtime slashing
func (k Keeper) SignedBlocksWindowDuration(ctx sdk.Context) (res time.Duration) {
	k.paramspace.Get(ctx, types.KeySignedBlocksWindowDuration, &res)
	return
}

// Downtime slashing threshold
func (k Keeper) MinSignedPerWindow(ctx sdk.Context) (res sdk.Dec) {
	k.paramspace.Get(ctx, types.KeyMinSignedPerWindow, &res)
	return
}

// Downtime unbond duration
func (k Keeper) DowntimeJailDuration(ctx sdk.Context) (res time.Duration) {
	k.paramspace.Get(ctx, types.KeyDowntimeJailDuration, &res)
	return
}

// SlashFractionDoubleSign
func (k Keeper) SlashFractionDoubleSign(ctx sdk.Context) (res sdk.Dec) {
	k.paramspace.Get(ctx, types.KeySlashFractionDoubleSign, &res)
	return
}

// SlashFractionDowntime
func (k Keeper) SlashFractionDowntime(ctx sdk.Context) (res sdk.Dec) {
	k.paramspace.Get(ctx, types.KeySlashFractionDowntime, &res)
	return
}

// GetParams returns the total set of slashing parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramspace.GetParamSet(ctx, &params)
	return params
}
