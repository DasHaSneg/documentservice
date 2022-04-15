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

func (k Keeper) ContractAll(c context.Context, req *types.QueryAllContractRequest) (*types.QueryAllContractResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var contracts []types.Contract
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	contractStore := prefix.NewStore(store, types.KeyPrefix(types.ContractKey))

	pageRes, err := query.Paginate(contractStore, req.Pagination, func(key []byte, value []byte) error {
		var contract types.Contract
		if err := k.cdc.Unmarshal(value, &contract); err != nil {
			return err
		}

		contracts = append(contracts, contract)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllContractResponse{Contract: contracts, Pagination: pageRes}, nil
}

func (k Keeper) Contract(c context.Context, req *types.QueryGetContractRequest) (*types.QueryGetContractResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	contract, found := k.GetContract(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetContractResponse{Contract: contract}, nil
}
