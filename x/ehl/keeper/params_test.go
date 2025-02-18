package keeper_test

import (
	"testing"

	testkeeper "bitmail/testutil/keeper"
	"bitmail/x/ehl/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.EhlKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
