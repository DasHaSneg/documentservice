package documentservice_test

import (
	"testing"

	keepertest "github.com/cosmonaut/documentservice/testutil/keeper"
	"github.com/cosmonaut/documentservice/testutil/nullify"
	"github.com/cosmonaut/documentservice/x/documentservice"
	"github.com/cosmonaut/documentservice/x/documentservice/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ContractList: []types.Contract{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		ContractCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DocumentserviceKeeper(t)
	documentservice.InitGenesis(ctx, *k, genesisState)
	got := documentservice.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ContractList, got.ContractList)
	require.Equal(t, genesisState.ContractCount, got.ContractCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
