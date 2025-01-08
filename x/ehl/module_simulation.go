package ehl

import (
	"math/rand"

	"bitmail/testutil/sample"
	ehlsimulation "bitmail/x/ehl/simulation"
	"bitmail/x/ehl/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = ehlsimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgCreateHashCid = "op_weight_msg_hash_cid"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateHashCid int = 100

	opWeightMsgUpdateHashCid = "op_weight_msg_hash_cid"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateHashCid int = 100

	opWeightMsgDeleteHashCid = "op_weight_msg_hash_cid"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteHashCid int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	ehlGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		HashCidList: []types.HashCid{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		HashCidCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&ehlGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateHashCid int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateHashCid, &weightMsgCreateHashCid, nil,
		func(_ *rand.Rand) {
			weightMsgCreateHashCid = defaultWeightMsgCreateHashCid
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateHashCid,
		ehlsimulation.SimulateMsgCreateHashCid(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateHashCid int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateHashCid, &weightMsgUpdateHashCid, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateHashCid = defaultWeightMsgUpdateHashCid
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateHashCid,
		ehlsimulation.SimulateMsgUpdateHashCid(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteHashCid int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteHashCid, &weightMsgDeleteHashCid, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteHashCid = defaultWeightMsgDeleteHashCid
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteHashCid,
		ehlsimulation.SimulateMsgDeleteHashCid(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateHashCid,
			defaultWeightMsgCreateHashCid,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				ehlsimulation.SimulateMsgCreateHashCid(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateHashCid,
			defaultWeightMsgUpdateHashCid,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				ehlsimulation.SimulateMsgUpdateHashCid(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteHashCid,
			defaultWeightMsgDeleteHashCid,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				ehlsimulation.SimulateMsgDeleteHashCid(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
