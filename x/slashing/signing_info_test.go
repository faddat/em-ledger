package slashing

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestGetSetValidatorSigningInfo(t *testing.T) {
	ctx, _, _, _, keeper, _ := createTestInput(t, DefaultParams())
	info, found := keeper.getValidatorSigningInfo(ctx, sdk.ConsAddress(addrs[0]))
	require.False(t, found)
	newInfo := NewValidatorSigningInfo(
		sdk.ConsAddress(addrs[0]),
		time.Unix(2, 0),
		false,
	)
	keeper.SetValidatorSigningInfo(ctx, sdk.ConsAddress(addrs[0]), newInfo)
	info, found = keeper.getValidatorSigningInfo(ctx, sdk.ConsAddress(addrs[0]))
	require.True(t, found)
	//require.Equal(t, info.StartHeight, int64(4))
	//require.Equal(t, info.IndexOffset, int64(3))
	require.Equal(t, info.JailedUntil, time.Unix(2, 0).UTC())
	//require.Equal(t, info.MissedBlocksCounter, int64(10))
}

func TestGetSetValidatorMissedBlockBitArray(t *testing.T) {
	ctx, _, _, _, keeper, _ := createTestInput(t, DefaultParams())
	missed := keeper.getValidatorMissedBlockBitArray(ctx, sdk.ConsAddress(addrs[0]), 0)
	require.False(t, missed) // treat empty key as not missed
	keeper.setValidatorMissedBlockBitArray(ctx, sdk.ConsAddress(addrs[0]), 0, true)
	missed = keeper.getValidatorMissedBlockBitArray(ctx, sdk.ConsAddress(addrs[0]), 0)
	require.True(t, missed) // now should be missed
}
