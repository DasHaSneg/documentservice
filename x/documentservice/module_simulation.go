package documentservice

import (
	"math/rand"

	"github.com/cosmonaut/documentservice/testutil/sample"
	documentservicesimulation "github.com/cosmonaut/documentservice/x/documentservice/simulation"
	"github.com/cosmonaut/documentservice/x/documentservice/types"
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
	_ = documentservicesimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateContract = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateContract int = 100

	opWeightMsgCreateAnnex = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateAnnex int = 100

	opWeightMsgSignAnnex = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSignAnnex int = 100

	opWeightMsgSignContract = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSignContract int = 100

	opWeightMsgCompleteContract = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCompleteContract int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	documentserviceGenesis := types.GenesisState{
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&documentserviceGenesis)
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

	var weightMsgCreateContract int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateContract, &weightMsgCreateContract, nil,
		func(_ *rand.Rand) {
			weightMsgCreateContract = defaultWeightMsgCreateContract
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateContract,
		documentservicesimulation.SimulateMsgCreateContract(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateAnnex int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateAnnex, &weightMsgCreateAnnex, nil,
		func(_ *rand.Rand) {
			weightMsgCreateAnnex = defaultWeightMsgCreateAnnex
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateAnnex,
		documentservicesimulation.SimulateMsgCreateAnnex(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSignAnnex int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSignAnnex, &weightMsgSignAnnex, nil,
		func(_ *rand.Rand) {
			weightMsgSignAnnex = defaultWeightMsgSignAnnex
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSignAnnex,
		documentservicesimulation.SimulateMsgSignAnnex(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSignContract int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSignContract, &weightMsgSignContract, nil,
		func(_ *rand.Rand) {
			weightMsgSignContract = defaultWeightMsgSignContract
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSignContract,
		documentservicesimulation.SimulateMsgSignContract(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCompleteContract int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCompleteContract, &weightMsgCompleteContract, nil,
		func(_ *rand.Rand) {
			weightMsgCompleteContract = defaultWeightMsgCompleteContract
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCompleteContract,
		documentservicesimulation.SimulateMsgCompleteContract(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
