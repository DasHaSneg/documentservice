package keeper_test

import (
	"testing"

	keepertest "github.com/cosmonaut/documentservice/testutil/keeper"
	"github.com/cosmonaut/documentservice/testutil/nullify"
	"github.com/cosmonaut/documentservice/x/documentservice/keeper"
	"github.com/cosmonaut/documentservice/x/documentservice/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNContract(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Contract {
	items := make([]types.Contract, n)
	for i := range items {
		items[i].Id = keeper.AppendContract(ctx, items[i])
	}
	return items
}

func TestContractGet(t *testing.T) {
	keeper, ctx := keepertest.DocumentserviceKeeper(t)
	items := createNContract(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetContract(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestContractRemove(t *testing.T) {
	keeper, ctx := keepertest.DocumentserviceKeeper(t)
	items := createNContract(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveContract(ctx, item.Id)
		_, found := keeper.GetContract(ctx, item.Id)
		require.False(t, found)
	}
}

func TestContractGetAll(t *testing.T) {
	keeper, ctx := keepertest.DocumentserviceKeeper(t)
	items := createNContract(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllContract(ctx)),
	)
}

func TestContractCount(t *testing.T) {
	keeper, ctx := keepertest.DocumentserviceKeeper(t)
	items := createNContract(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetContractCount(ctx))
}
