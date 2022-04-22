package keeper

import (
	"context"
	"strconv"
	"time"

	"github.com/cosmonaut/documentservice/x/documentservice/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateContract(goCtx context.Context, msg *types.MsgCreateContract) (*types.MsgCreateContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var contract = types.Contract{
		Creator:      msg.Creator,
		ContractHash: msg.ContractHash,
		State:        types.PendingAnnexCreation,
		Seller:       msg.Creator,
		Buyer:        msg.Buyer,
		CreateDate:   strconv.FormatInt(time.Now().Unix(), 10),
	}

	id := k.AppendContract(ctx, contract)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, "udrservice"),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.ContractEventKey),
			sdk.NewAttribute(types.ContractEventId, strconv.FormatUint(id, 10)),
		),
	)

	return &types.MsgCreateContractResponse{Id: id, CreateDate: contract.CreateDate}, nil
}
