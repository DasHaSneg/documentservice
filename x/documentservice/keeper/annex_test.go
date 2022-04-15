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

func createNAnnex(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Annex {
	items := make([]types.Annex, n)
	for i := range items {
		items[i].Id = keeper.AppendAnnex(ctx, items[i])
	}
	return items
}

func TestAnnexGet(t *testing.T) {
	keeper, ctx := keepertest.DocumentserviceKeeper(t)
	items := createNAnnex(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetAnnex(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestAnnexRemove(t *testing.T) {
	keeper, ctx := keepertest.DocumentserviceKeeper(t)
	items := createNAnnex(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAnnex(ctx, item.Id)
		_, found := keeper.GetAnnex(ctx, item.Id)
		require.False(t, found)
	}
}

func TestAnnexGetAll(t *testing.T) {
	keeper, ctx := keepertest.DocumentserviceKeeper(t)
	items := createNAnnex(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllAnnex(ctx)),
	)
}

func TestAnnexCount(t *testing.T) {
	keeper, ctx := keepertest.DocumentserviceKeeper(t)
	items := createNAnnex(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetAnnexCount(ctx))
}
