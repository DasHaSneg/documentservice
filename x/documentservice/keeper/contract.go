package keeper

import (
	"encoding/binary"

	"github.com/cosmonaut/documentservice/x/documentservice/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AppendContract(ctx sdk.Context, contract types.Contract) uint64 {
	count := k.GetContractCount(ctx)

	contract.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.ContractKey))

	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, contract.Id)

	byteValue := k.cdc.MustMarshal(&contract)

	store.Set(byteKey, byteValue)

	k.SetContractCount(ctx, count+1)
	return count
}

func (k Keeper) GetContractCount(ctx sdk.Context) uint64 {
	byteKey := []byte(types.ContractCountKey)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), byteKey)

	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}

	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetContractCount(ctx sdk.Context, count uint64) {
	byteKey := []byte(types.ContractCountKey)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), byteKey)

	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)

	store.Set(byteKey, bz)
}

func (k Keeper) GetContract(ctx sdk.Context, id uint64) (contract types.Contract) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.ContractKey))

	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, id)

	byteResult := store.Get(byteKey)

	//TODO: check if value exist

	k.cdc.MustUnmarshal(byteResult, &contract)
	return contract
}
