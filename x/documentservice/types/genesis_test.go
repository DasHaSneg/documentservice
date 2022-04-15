package types_test

import (
	"testing"

	"github.com/cosmonaut/documentservice/x/documentservice/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				ContractList: []types.Contract{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				ContractCount: 2,
				AnnexList: []types.Annex{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				AnnexCount: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated contract",
			genState: &types.GenesisState{
				ContractList: []types.Contract{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid contract count",
			genState: &types.GenesisState{
				ContractList: []types.Contract{
					{
						Id: 1,
					},
				},
				ContractCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated annex",
			genState: &types.GenesisState{
				AnnexList: []types.Annex{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid annex count",
			genState: &types.GenesisState{
				AnnexList: []types.Annex{
					{
						Id: 1,
					},
				},
				AnnexCount: 0,
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
