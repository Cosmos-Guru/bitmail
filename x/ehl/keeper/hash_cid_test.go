package keeper_test

import (
	"testing"

	keepertest "bitmail/testutil/keeper"
	"bitmail/testutil/nullify"
	"bitmail/x/ehl/keeper"
	"bitmail/x/ehl/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNHashCid(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.HashCid {
	items := make([]types.HashCid, n)
	for i := range items {
		items[i].Id = keeper.AppendHashCid(ctx, items[i])
	}
	return items
}

func TestHashCidGet(t *testing.T) {
	keeper, ctx := keepertest.EhlKeeper(t)
	items := createNHashCid(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetHashCid(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestHashCidRemove(t *testing.T) {
	keeper, ctx := keepertest.EhlKeeper(t)
	items := createNHashCid(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveHashCid(ctx, item.Id)
		_, found := keeper.GetHashCid(ctx, item.Id)
		require.False(t, found)
	}
}

func TestHashCidGetAll(t *testing.T) {
	keeper, ctx := keepertest.EhlKeeper(t)
	items := createNHashCid(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllHashCid(ctx)),
	)
}

func TestHashCidCount(t *testing.T) {
	keeper, ctx := keepertest.EhlKeeper(t)
	items := createNHashCid(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetHashCidCount(ctx))
}
