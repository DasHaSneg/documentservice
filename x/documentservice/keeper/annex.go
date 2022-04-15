package keeper

import (
	"encoding/binary"

	"github.com/cosmonaut/documentservice/x/documentservice/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetAnnexCount get the total number of annex
func (k Keeper) GetAnnexCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.AnnexCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetAnnexCount set the total number of annex
func (k Keeper) SetAnnexCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.AnnexCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendAnnex appends a annex in the store with a new id and update the count
func (k Keeper) AppendAnnex(
	ctx sdk.Context,
	annex types.Annex,
) uint64 {
	// Create the annex
	count := k.GetAnnexCount(ctx)

	// Set the ID of the appended value
	annex.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AnnexKey))
	appendedValue := k.cdc.MustMarshal(&annex)
	store.Set(GetAnnexIDBytes(annex.Id), appendedValue)

	// Update annex count
	k.SetAnnexCount(ctx, count+1)

	return count
}

// SetAnnex set a specific annex in the store
func (k Keeper) SetAnnex(ctx sdk.Context, annex types.Annex) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AnnexKey))
	b := k.cdc.MustMarshal(&annex)
	store.Set(GetAnnexIDBytes(annex.Id), b)
}

// GetAnnex returns a annex from its id
func (k Keeper) GetAnnex(ctx sdk.Context, id uint64) (val types.Annex, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AnnexKey))
	b := store.Get(GetAnnexIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAnnex removes a annex from the store
func (k Keeper) RemoveAnnex(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AnnexKey))
	store.Delete(GetAnnexIDBytes(id))
}

// GetAllAnnex returns all annex
func (k Keeper) GetAllAnnex(ctx sdk.Context) (list []types.Annex) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AnnexKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Annex
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAnnexIDBytes returns the byte representation of the ID
func GetAnnexIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetAnnexIDFromBytes returns ID in uint64 format from a byte array
func GetAnnexIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
