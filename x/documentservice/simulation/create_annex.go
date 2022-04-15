package simulation

import (
	"math/rand"

	"github.com/cosmonaut/documentservice/x/documentservice/keeper"
	"github.com/cosmonaut/documentservice/x/documentservice/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgCreateAnnex(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCreateAnnex{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the CreateAnnex simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CreateAnnex simulation not implemented"), nil, nil
	}
}
