package dkg

import (
	"math/rand"

	"dkg/testutil/sample"
	dkgsimulation "dkg/x/dkg/simulation"
	"dkg/x/dkg/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = dkgsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgRefundMsgRequest = "op_weight_msg_refund_msg_request"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRefundMsgRequest int = 100

	opWeightMsgFileDispute = "op_weight_msg_file_dispute"
	// TODO: Determine the simulation weight value
	defaultWeightMsgFileDispute int = 100

	opWeightMsgStartKeygen = "op_weight_msg_start_keygen"
	// TODO: Determine the simulation weight value
	defaultWeightMsgStartKeygen int = 100

	opWeightMsgKeygenResult = "op_weight_msg_keygen_result"
	// TODO: Determine the simulation weight value
	defaultWeightMsgKeygenResult int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	dkgGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&dkgGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgRefundMsgRequest int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRefundMsgRequest, &weightMsgRefundMsgRequest, nil,
		func(_ *rand.Rand) {
			weightMsgRefundMsgRequest = defaultWeightMsgRefundMsgRequest
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRefundMsgRequest,
		dkgsimulation.SimulateMsgRefundMsgRequest(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgFileDispute int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgFileDispute, &weightMsgFileDispute, nil,
		func(_ *rand.Rand) {
			weightMsgFileDispute = defaultWeightMsgFileDispute
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgFileDispute,
		dkgsimulation.SimulateMsgFileDispute(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgStartKeygen int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgStartKeygen, &weightMsgStartKeygen, nil,
		func(_ *rand.Rand) {
			weightMsgStartKeygen = defaultWeightMsgStartKeygen
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgStartKeygen,
		dkgsimulation.SimulateMsgStartKeygen(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgKeygenResult int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgKeygenResult, &weightMsgKeygenResult, nil,
		func(_ *rand.Rand) {
			weightMsgKeygenResult = defaultWeightMsgKeygenResult
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgKeygenResult,
		dkgsimulation.SimulateMsgKeygenResult(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
