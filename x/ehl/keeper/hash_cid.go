package keeper

import (
	"encoding/binary"

	"bitmail/x/ehl/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetHashCidCount get the total number of hashCid
func (k Keeper) GetHashCidCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.HashCidCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetHashCidCount set the total number of hashCid
func (k Keeper) SetHashCidCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.HashCidCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendHashCid appends a hashCid in the store with a new id and update the count
func (k Keeper) AppendHashCid(
	ctx sdk.Context,
	hashCid types.HashCid,
) uint64 {
	// Create the hashCid
	count := k.GetHashCidCount(ctx)

	// Set the ID of the appended value
	hashCid.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HashCidKey))
	appendedValue := k.cdc.MustMarshal(&hashCid)
	store.Set(GetHashCidIDBytes(hashCid.Id), appendedValue)

	// Update hashCid count
	k.SetHashCidCount(ctx, count+1)

	return count
}

// SetHashCid set a specific hashCid in the store
func (k Keeper) SetHashCid(ctx sdk.Context, hashCid types.HashCid) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HashCidKey))
	b := k.cdc.MustMarshal(&hashCid)
	store.Set(GetHashCidIDBytes(hashCid.Id), b)
}

// GetHashCid returns a hashCid from its id
func (k Keeper) GetHashCid(ctx sdk.Context, id uint64) (val types.HashCid, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HashCidKey))
	b := store.Get(GetHashCidIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveHashCid removes a hashCid from the store
func (k Keeper) RemoveHashCid(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HashCidKey))
	store.Delete(GetHashCidIDBytes(id))
}

// GetAllHashCid returns all hashCid
func (k Keeper) GetAllHashCid(ctx sdk.Context) (list []types.HashCid) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HashCidKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.HashCid
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetHashCidIDBytes returns the byte representation of the ID
func GetHashCidIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetHashCidIDFromBytes returns ID in uint64 format from a byte array
func GetHashCidIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
