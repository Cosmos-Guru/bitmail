package keeper_test

import (
	"testing"

	testkeeper "bitmail/testutil/keeper"
	"bitmail/x/bitmail/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.BitmailKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
