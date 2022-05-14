package keeper

import (
	"context"
	"strconv"
	"time"

	"github.com/cosmonaut/documentservice/x/documentservice/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateAnnex(goCtx context.Context, msg *types.MsgCreateAnnex) (*types.MsgCreateAnnexResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	contract, found := k.GetContract(ctx, msg.ContractId)

	if !found {
		return nil, sdkerrors.Wrapf(types.ErrID, "Contract Id %d does not exist", msg.ContractId)
	}

	var annex = types.Annex{
		Creator:    msg.Creator,
		AnnexHash:  msg.AnnexHash,
		ContractId: msg.ContractId,
		State:      types.WaitingForSignature(1),
		Seller:     msg.Creator,
		Buyer:      msg.Buyer,
		CreateDate: strconv.FormatInt(time.Now().Unix(), 10),
	}

	id := k.AppendAnnex(ctx, annex)

	contract.State = types.WaitingForAnnexSignature(1)

	k.SetContract(ctx, contract)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, "documentservice"),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.AnnexEventKey),
			sdk.NewAttribute(types.AnnexEventId, strconv.FormatUint(id, 10)),
			sdk.NewAttribute(types.AnnexEventAnnexState, annex.State),
			sdk.NewAttribute(types.AnnexEventContractState, contract.State),
		),
	)

	return &types.MsgCreateAnnexResponse{Id: id, CreateDate: annex.CreateDate}, nil
}
