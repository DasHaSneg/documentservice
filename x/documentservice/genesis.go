package documentservice

import (
	"github.com/cosmonaut/documentservice/x/documentservice/keeper"
	"github.com/cosmonaut/documentservice/x/documentservice/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the contract
	for _, elem := range genState.ContractList {
		k.SetContract(ctx, elem)
	}

	// Set contract count
	k.SetContractCount(ctx, genState.ContractCount)
	// Set all the annex
	for _, elem := range genState.AnnexList {
		k.SetAnnex(ctx, elem)
	}

	// Set annex count
	k.SetAnnexCount(ctx, genState.AnnexCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.ContractList = k.GetAllContract(ctx)
	genesis.ContractCount = k.GetContractCount(ctx)
	genesis.AnnexList = k.GetAllAnnex(ctx)
	genesis.AnnexCount = k.GetAnnexCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
