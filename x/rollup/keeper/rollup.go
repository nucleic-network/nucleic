package keeper

import (
	"github.com/eve-network/eve/x/rollup/types"

	"cosmossdk.io/store/prefix"

	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetRollup(ctx sdk.Context, rollup types.Rollup) error {
	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	bz, err := k.cdc.Marshal(&rollup)
	if err != nil {
		return err
	}
	store.Set(types.RollupKey(rollup.RollupId), bz)
	return nil
}

func (k Keeper) GetRollup(ctx sdk.Context, rollupId string) (types.Rollup, bool) {
	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	b := store.Get(types.RollupKey(rollupId))
	if b == nil {
		return types.Rollup{}, false
	}

	var val types.Rollup
	err := k.cdc.Unmarshal(b, &val)
	if err != nil {
		return types.Rollup{}, false
	}
	return val, true
}

func (k Keeper) RemoveRollup(ctx sdk.Context, rollupId string) {
	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store.Delete(types.RollupKey(rollupId))
}

func (k Keeper) GetAllRollups(ctx sdk.Context) []types.Rollup {
	prefixStore := prefix.NewStore(runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx)), []byte(types.RollupKeyPrefix))
	iter := prefixStore.Iterator(nil, nil)
	defer iter.Close()

	var list []types.Rollup

	for ; iter.Valid(); iter.Next() {
		var val types.Rollup
		k.cdc.MustUnmarshal(iter.Value(), &val)
		list = append(list, val)
	}

	return list
}
