package keeper

import (
	"context"

	"github.com/cosmonaut/documentservice/x/documentservice/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ContractsByInn(goCtx context.Context, req *types.QueryContractsByInnRequest) (*types.QueryContractsByInnResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	var contracts []types.Contract

	store := ctx.KVStore(k.storeKey)
	contractStore := prefix.NewStore(store, types.KeyPrefix(types.ContractKey))

	pageRes, err := query.Paginate(contractStore, req.Pagination, func(key []byte, value []byte) error {
		var contract types.Contract
		if err := k.cdc.Unmarshal(value, &contract); err != nil {
			return err
		}
		if contract.SellerInn == req.Inn || contract.BuyerInn == req.Inn {
			contracts = append(contracts, contract)
		}

		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &types.QueryContractsByInnResponse{Contract: contracts, Pagination: pageRes}, nil
}
