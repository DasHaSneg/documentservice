package keeper

import (
	"context"

	"github.com/cosmonaut/documentservice/x/documentservice/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SignContract(goCtx context.Context, msg *types.MsgSignContract) (*types.MsgSignContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	contract, found := k.GetContract(ctx, msg.ContractId)

	if !found {
		return nil, sdkerrors.Wrapf(types.ErrID, "Contract Id %d does not exist", msg.ContractId)
	}

	if contract.Buyer != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrID, "Creator and Buyer address do not match")
	}

	if contract.State != types.WaitingForSignature(1) {
		return nil, sdkerrors.Wrapf(types.ErrState, "Contract state is not equal %s", types.WaitingForSignature(1))
	}

	contract.State = types.Signed

	k.SetContract(ctx, contract)

	return &types.MsgSignContractResponse{}, nil
}
