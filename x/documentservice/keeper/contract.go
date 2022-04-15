package keeper

import (
	"encoding/binary"

	"github.com/cosmonaut/documentservice/x/documentservice/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetContractCount get the total number of contract
func (k Keeper) GetContractCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ContractCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetContractCount set the total number of contract
func (k Keeper) SetContractCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ContractCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendContract appends a contract in the store with a new id and update the count
func (k Keeper) AppendContract(
	ctx sdk.Context,
	contract types.Contract,
) uint64 {
	// Create the contract
	count := k.GetContractCount(ctx)

	// Set the ID of the appended value
	contract.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractKey))
	appendedValue := k.cdc.MustMarshal(&contract)
	store.Set(GetContractIDBytes(contract.Id), appendedValue)

	// Update contract count
	k.SetContractCount(ctx, count+1)

	return count
}

// SetContract set a specific contract in the store
func (k Keeper) SetContract(ctx sdk.Context, contract types.Contract) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractKey))
	b := k.cdc.MustMarshal(&contract)
	store.Set(GetContractIDBytes(contract.Id), b)
}

// GetContract returns a contract from its id
func (k Keeper) GetContract(ctx sdk.Context, id uint64) (val types.Contract, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractKey))
	b := store.Get(GetContractIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveContract removes a contract from the store
func (k Keeper) RemoveContract(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractKey))
	store.Delete(GetContractIDBytes(id))
}

// GetAllContract returns all contract
func (k Keeper) GetAllContract(ctx sdk.Context) (list []types.Contract) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Contract
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetContractIDBytes returns the byte representation of the ID
func GetContractIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetContractIDFromBytes returns ID in uint64 format from a byte array
func GetContractIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
