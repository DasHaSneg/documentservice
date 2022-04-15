package keeper

import (
	"context"

	"github.com/cosmonaut/documentservice/x/documentservice/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AnnexAll(c context.Context, req *types.QueryAllAnnexRequest) (*types.QueryAllAnnexResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var annexs []types.Annex
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	annexStore := prefix.NewStore(store, types.KeyPrefix(types.AnnexKey))

	pageRes, err := query.Paginate(annexStore, req.Pagination, func(key []byte, value []byte) error {
		var annex types.Annex
		if err := k.cdc.Unmarshal(value, &annex); err != nil {
			return err
		}

		annexs = append(annexs, annex)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAnnexResponse{Annex: annexs, Pagination: pageRes}, nil
}

func (k Keeper) Annex(c context.Context, req *types.QueryGetAnnexRequest) (*types.QueryGetAnnexResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	annex, found := k.GetAnnex(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetAnnexResponse{Annex: annex}, nil
}
