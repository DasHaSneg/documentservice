package keeper

import (
	"context"

	"github.com/cosmonaut/documentservice/x/documentservice/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SignAnnex(goCtx context.Context, msg *types.MsgSignAnnex) (*types.MsgSignAnnexResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	annex, found := k.GetAnnex(ctx, msg.AnnexId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrID, "Annex Id %d does not exist", msg.AnnexId)
	}

	if annex.Buyer != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrID, "Creator and Buyer address do not match")
	}

	annex.State = types.Signed

	contract, found := k.GetContract(ctx, annex.ContractId)

	if !found {
		return nil, sdkerrors.Wrapf(types.ErrID, "Contract Id %d does not exist", annex.ContractId)
	}

	if contract.State != types.WaitingForAnnexSignature(1) {
		return nil, sdkerrors.Wrapf(types.ErrState, "Contract state is not equal %s", types.WaitingForAnnexSignature(1))
	}

	contract.State = types.WaitingForSignature(1)

	k.SetAnnex(ctx, annex)

	k.SetContract(ctx, contract)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, "documentservice"),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.SignAnnexEventKey),
			sdk.NewAttribute(types.SignAnnexState, annex.State),
		),
	)

	return &types.MsgSignAnnexResponse{}, nil
}
