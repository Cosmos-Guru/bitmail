package ehl_test

import (
	"testing"

	keepertest "bitmail/testutil/keeper"
	"bitmail/testutil/nullify"
	"bitmail/x/ehl"
	"bitmail/x/ehl/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		HashCidList: []types.HashCid{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		HashCidCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.EhlKeeper(t)
	ehl.InitGenesis(ctx, *k, genesisState)
	got := ehl.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.HashCidList, got.HashCidList)
	require.Equal(t, genesisState.HashCidCount, got.HashCidCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
