package keeper

import (
	"fmt"

	"github.com/eve-network/eve/x/rollapp/types"

	"cosmossdk.io/store/prefix"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// FinalizeQueue is called every block to finalize states when their dispute period over.
func (k Keeper) FinalizeQueue(ctx sdk.Context) error {
	if uint64(ctx.BlockHeight()) < k.DisputePeriodInBlocks(ctx) {
		return nil
	}
	// TODO: need implement this
	return nil
}

func (k Keeper) finalizePending(ctx sdk.Context, pendingFinalizationQueue []types.BlockHeightToFinalizationQueue) error {
	// Iterate over all the pending finalization queue
	for _, blockHeightToFinalizationQueue := range pendingFinalizationQueue {
		// finalize pending states
		for _, stateInfoIndex := range blockHeightToFinalizationQueue.FinalizationQueue {
			stateInfo := k.MustGetStateInfo(ctx, stateInfoIndex.RollappId, stateInfoIndex.Index)
			if stateInfo.Status != types.Status_PENDING {
				panic(fmt.Sprintf("invariant broken: stateInfo is not in pending state: rollapp: %s: status: %s", stateInfoIndex.RollappId, stateInfo.Status))
			}
			stateInfo.Finalize()
			// update the status of the stateInfo
			k.SetStateInfo(ctx, stateInfo)
			// update the LatestStateInfoIndex of the rollapp
			k.SetLatestFinalizedStateIndex(ctx, stateInfoIndex)
			// call the after-update-state hook
			keeperHooks := k.GetHooks()
			err := keeperHooks.AfterStateFinalized(ctx, stateInfoIndex.RollappId, &stateInfo)
			if err != nil {
				return fmt.Errorf("after state finalized: %w", err)
			}

			// emit event
			ctx.EventManager().EmitEvent(
				sdk.NewEvent(types.EventTypeStateUpdate,
					stateInfo.GetEvents()...,
				),
			)
		}
		k.RemoveBlockHeightToFinalizationQueue(ctx, blockHeightToFinalizationQueue.CreationHeight)
	}
	return nil
}

// SetBlockHeightToFinalizationQueue set a specific blockHeightToFinalizationQueue in the store from its index
func (k Keeper) SetBlockHeightToFinalizationQueue(ctx sdk.Context, blockHeightToFinalizationQueue types.BlockHeightToFinalizationQueue) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BlockHeightToFinalizationQueueKeyPrefix))
	b := k.cdc.MustMarshal(&blockHeightToFinalizationQueue)
	store.Set(types.BlockHeightToFinalizationQueueKey(
		blockHeightToFinalizationQueue.CreationHeight,
	), b)
}

// GetBlockHeightToFinalizationQueue returns a blockHeightToFinalizationQueue from its index
func (k Keeper) GetBlockHeightToFinalizationQueue(
	ctx sdk.Context,
	creationHeight uint64,
) (val types.BlockHeightToFinalizationQueue, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BlockHeightToFinalizationQueueKeyPrefix))

	b := store.Get(types.BlockHeightToFinalizationQueueKey(
		creationHeight,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveBlockHeightToFinalizationQueue removes a blockHeightToFinalizationQueue from the store
func (k Keeper) RemoveBlockHeightToFinalizationQueue(
	ctx sdk.Context,
	creationHeight uint64,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BlockHeightToFinalizationQueueKeyPrefix))
	store.Delete(types.BlockHeightToFinalizationQueueKey(
		creationHeight,
	))
}

// GetAllFinalizationQueueUntilHeight returns all the blockHeightToFinalizationQueues with creation height equal or less to the input height
func (k Keeper) GetAllFinalizationQueueUntilHeight(ctx sdk.Context, height uint64) (list []types.BlockHeightToFinalizationQueue) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BlockHeightToFinalizationQueueKeyPrefix))
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.BlockHeightToFinalizationQueue
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if height < val.CreationHeight {
			break
		}
		list = append(list, val)
	}

	return
}

// GetAllBlockHeightToFinalizationQueue returns all blockHeightToFinalizationQueue
func (k Keeper) GetAllBlockHeightToFinalizationQueue(ctx sdk.Context) (list []types.BlockHeightToFinalizationQueue) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BlockHeightToFinalizationQueueKeyPrefix))
	iterator := store.Iterator(nil, nil)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.BlockHeightToFinalizationQueue
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
