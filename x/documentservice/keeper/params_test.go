package keeper_test

import (
	"testing"

	testkeeper "github.com/cosmonaut/documentservice/testutil/keeper"
	"github.com/cosmonaut/documentservice/x/documentservice/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.DocumentserviceKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
