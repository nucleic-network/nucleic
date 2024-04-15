package keeper

import (
	"github.com/eve-network/eve/x/rollup/types"

	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetRollup(ctx sdk.Context, rollup types.Rollup) error {
	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	bz, err := k.cdc.Marshal(&rollup)
	if err != nil {
		return err
	}
	store.Set(types.RollupKey(
		rollup.RollupId,
	), bz)
	return nil
}
