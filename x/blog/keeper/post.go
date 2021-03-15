package keeper

import (
	"github.com/clockworkgr/spvue/x/blog/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"
)

// GetPostCount get the total number of post
func (k Keeper) GetPostCount(ctx sdk.Context) int64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostCountKey))
	byteKey := types.KeyPrefix(types.PostCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseInt(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to int64
		panic("cannot decode count")
	}

	return count
}

// SetPostCount set the total number of post
func (k Keeper) SetPostCount(ctx sdk.Context, count int64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostCountKey))
	byteKey := types.KeyPrefix(types.PostCountKey)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// AppendPost appends a post in the store with a new id and update the count
func (k Keeper) AppendPost(
	ctx sdk.Context,
	creator string,
	title string,
	body string,
	votes int32,
) string {
	// Create the post
	count := k.GetPostCount(ctx)
	id := strconv.FormatInt(count, 10)
	var post = types.Post{
		Creator: creator,
		Id:      id,
		Title:   title,
		Body:    body,
		Votes:   votes,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))
	key := types.KeyPrefix(types.PostKey + post.Id)
	value := k.cdc.MustMarshalBinaryBare(&post)
	store.Set(key, value)

	// Update post count
	k.SetPostCount(ctx, count+1)

	return id
}

// SetPost set a specific post in the store
func (k Keeper) SetPost(ctx sdk.Context, post types.Post) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))
	b := k.cdc.MustMarshalBinaryBare(&post)
	store.Set(types.KeyPrefix(types.PostKey+post.Id), b)
}

// GetPost returns a post from its id
func (k Keeper) GetPost(ctx sdk.Context, key string) types.Post {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))
	var post types.Post
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.PostKey+key)), &post)
	return post
}

// HasPost checks if the post exists in the store
func (k Keeper) HasPost(ctx sdk.Context, id string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))
	return store.Has(types.KeyPrefix(types.PostKey + id))
}

// GetPostOwner returns the creator of the post
func (k Keeper) GetPostOwner(ctx sdk.Context, key string) string {
	return k.GetPost(ctx, key).Creator
}

// DeletePost removes a post from the store
func (k Keeper) RemovePost(ctx sdk.Context, key string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))
	store.Delete(types.KeyPrefix(types.PostKey + key))
}

// GetAllPost returns all post
func (k Keeper) GetAllPost(ctx sdk.Context) (list []types.Post) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.PostKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Post
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
