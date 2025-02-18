package ehl

import (
	"bitmail/x/ehl/keeper"
	"bitmail/x/ehl/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the hashCid
	for _, elem := range genState.HashCidList {
		k.SetHashCid(ctx, elem)
	}

	// Set hashCid count
	k.SetHashCidCount(ctx, genState.HashCidCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.HashCidList = k.GetAllHashCid(ctx)
	genesis.HashCidCount = k.GetHashCidCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
