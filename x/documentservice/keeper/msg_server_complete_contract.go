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

	if msg.Creator != contract.Seller || msg.Creator != contract.Buyer {
		return nil, sdkerrors.Wrapf(types.ErrAccess, "Transaction Creator %s is not the contracting party", msg.Creator)
	}

	if contract.State == types.WaitingForCompletionConfirmationBuyer && msg.Creator == contract.Seller {
		return nil, sdkerrors.Wrapf(types.ErrState, "Seller %s already confirmed completion", msg.Creator)
	}

	if contract.State == types.WaitingForCompletionConfirmationSeller && msg.Creator == contract.Buyer {
		return nil, sdkerrors.Wrapf(types.ErrState, "Buyer %s already confirmed completion", msg.Creator)
	}

	if !found {
		return nil, sdkerrors.Wrapf(types.ErrID, "Contract Id %d does not exist", msg.ContractId)
	}

	if contract.State != types.Signed {
		return nil, sdkerrors.Wrapf(types.ErrState, "Contract state %s is not equal %s", contract.State, types.Signed)
	}

	if msg.Creator == contract.Seller {
		if contract.State == types.WaitingForCompletionConfirmationSeller {
			contract.State = types.Completed
		} else {
			contract.State = types.WaitingForCompletionConfirmationBuyer
		}
	}

	if msg.Creator == contract.Buyer {
		if contract.State == types.WaitingForCompletionConfirmationBuyer {
			contract.State = types.Completed
		} else {
			contract.State = types.WaitingForCompletionConfirmationSeller
		}
	}

	if contract.State == types.Completed {
		annexes := k.GetAllAnnexByContractId(ctx, contract.Id)
		if len(annexes) != 0 {
			for _, annex := range annexes {
				annex.State = types.Completed
				k.SetAnnex(ctx, annex)
			}
		}
	}

	k.SetContract(ctx, contract)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, "documentservice"),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.CompleteContractEventKey),
		),
	)

	return &types.MsgCompleteContractResponse{}, nil
}
