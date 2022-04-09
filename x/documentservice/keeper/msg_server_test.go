package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/cosmonaut/documentservice/testutil/keeper"
	"github.com/cosmonaut/documentservice/x/documentservice/keeper"
	"github.com/cosmonaut/documentservice/x/documentservice/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.DocumentserviceKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
