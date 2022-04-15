package keeper

import (
	"context"

	"github.com/cosmonaut/documentservice/x/documentservice/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CompleteContract(goCtx context.Context, msg *types.MsgCompleteContract) (*types.MsgCompleteContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	contract, found := k.GetContract(ctx, msg.ContractId)

	if !found {
		return nil, sdkerrors.Wrapf(types.ErrID, "Contract Id %d does not exist", msg.ContractId)
	}

	if contract.Seller != msg.Creator || contract.Buyer != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrID, "Creator and Seller address do not match")
	}

	if contract.State != types.Signed {
		return nil, sdkerrors.Wrapf(types.ErrState, "Contract state is not equal %s", types.Signed)
	}

	contract.State = types.Completed

	k.SetContract(ctx, contract)

	return &types.MsgCompleteContractResponse{}, nil
}
