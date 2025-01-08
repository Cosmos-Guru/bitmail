package bitmail_test

import (
	"testing"

	keepertest "bitmail/testutil/keeper"
	"bitmail/testutil/nullify"
	"bitmail/x/bitmail"
	"bitmail/x/bitmail/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BitmailKeeper(t)
	bitmail.InitGenesis(ctx, *k, genesisState)
	got := bitmail.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
