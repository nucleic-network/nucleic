package keeper

import (
	"context"

	"github.com/eve-network/eve/x/rollapp/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/store/prefix"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
)

var _ types.QueryServer = Keeper{}

// LatestHeight implements types.QueryServer.
func (k Keeper) LatestHeight(c context.Context, req *types.QueryGetLatestHeightRequest) (*types.QueryGetLatestHeightResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	var val types.StateInfoIndex
	var found bool
	if req.Finalized {
		val, found = k.GetLatestFinalizedStateIndex(
			ctx,
			req.RollappId,
		)
	} else {
		val, found = k.GetLatestStateInfoIndex(
			ctx,
			req.RollappId,
		)
	}
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	state := k.MustGetStateInfo(ctx, req.RollappId, val.Index)

	return &types.QueryGetLatestHeightResponse{
		Height: state.GetLatestHeight(),
	}, nil
}

// LatestStateIndex implements types.QueryServer.
func (k Keeper) LatestStateIndex(c context.Context, req *types.QueryGetLatestStateIndexRequest) (*types.QueryGetLatestStateIndexResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	var val types.StateInfoIndex
	var found bool
	if req.Finalized {
		val, found = k.GetLatestFinalizedStateIndex(
			ctx,
			req.RollappId,
		)
	} else {
		val, found = k.GetLatestStateInfoIndex(
			ctx,
			req.RollappId,
		)
	}

	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetLatestStateIndexResponse{StateIndex: val}, nil
}

// Params implements types.QueryServer.
func (k Keeper) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryParamsResponse{Params: k.GetParams(ctx)}, nil
}

// Rollapp implements types.QueryServer.
func (k Keeper) Rollapp(c context.Context, req *types.QueryGetRollappRequest) (*types.QueryGetRollappResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetRollapp(
		ctx,
		req.RollappId,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	rollappResponse := &types.QueryGetRollappResponse{Rollapp: val}
	latestStateInfoIndex, found := k.GetLatestStateInfoIndex(ctx, val.RollappId)
	if found {
		rollappResponse.LatestStateIndex = &latestStateInfoIndex
	}
	latestFinalizedStateInfoIndex, found := k.GetLatestFinalizedStateIndex(ctx, val.RollappId)
	if found {
		rollappResponse.LatestFinalizedStateIndex = &latestFinalizedStateInfoIndex
	}

	return rollappResponse, nil
}

// RollappAll implements types.QueryServer.
func (k Keeper) RollappAll(c context.Context, req *types.QueryAllRollappRequest) (*types.QueryAllRollappResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var rollapps []types.RollappSummary
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	rollappStore := prefix.NewStore(store, types.KeyPrefix(types.RollappKeyPrefix))

	pageRes, err := query.Paginate(rollappStore, req.Pagination, func(key, value []byte) error {
		var rollapp types.Rollapp
		if err := k.cdc.Unmarshal(value, &rollapp); err != nil {
			return err
		}
		rollappSummary := types.RollappSummary{
			RollappId: rollapp.RollappId,
		}
		latestStateInfoIndex, found := k.GetLatestStateInfoIndex(ctx, rollapp.RollappId)
		if found {
			rollappSummary.LatestStateIndex = &latestStateInfoIndex
		}
		latestFinalizedStateInfoIndex, found := k.GetLatestFinalizedStateIndex(ctx, rollapp.RollappId)
		if found {
			rollappSummary.LatestFinalizedStateIndex = &latestFinalizedStateInfoIndex
		}

		rollapps = append(rollapps, rollappSummary)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRollappResponse{Rollapp: rollapps, Pagination: pageRes}, nil
}

// RollappByEIP155 implements types.QueryServer.
func (k Keeper) RollappByEIP155(c context.Context, req *types.QueryGetRollappByEIP155Request) (*types.QueryGetRollappResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetRollappByEIP155(
		ctx,
		req.Eip155,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	rollappResponse := &types.QueryGetRollappResponse{Rollapp: val}
	latestStateInfoIndex, found := k.GetLatestStateInfoIndex(ctx, val.RollappId)
	if found {
		rollappResponse.LatestStateIndex = &latestStateInfoIndex
	}
	latestFinalizedStateInfoIndex, found := k.GetLatestFinalizedStateIndex(ctx, val.RollappId)
	if found {
		rollappResponse.LatestFinalizedStateIndex = &latestFinalizedStateInfoIndex
	}

	return rollappResponse, nil
}

// StateInfo implements types.QueryServer.
func (k Keeper) StateInfo(c context.Context, req *types.QueryGetStateInfoRequest) (*types.QueryGetStateInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	if req.Height == 0 && req.Index == 0 {
		if req.Finalized {
			latestFinalizedStateIndex, found := k.GetLatestFinalizedStateIndex(ctx, req.RollappId)
			if !found {
				return nil, errorsmod.Wrapf(types.ErrNoFinalizedStateYetForRollapp,
					"LatestFinalizedStateIndex wasn't found for rollappId=%s", req.RollappId)
			}
			req.Index = latestFinalizedStateIndex.Index
		} else {
			latestStateIndex, found := k.GetLatestStateInfoIndex(ctx, req.RollappId)
			if !found {
				if _, exists := k.GetRollapp(ctx, req.RollappId); !exists {
					return nil, types.ErrRollappNotRegistered
				}
				return nil, status.Error(codes.NotFound, "not found")
			}
			req.Index = latestStateIndex.Index
		}
	}

	var stateInfo types.StateInfo
	if req.Index != 0 {
		val, found := k.GetStateInfo(ctx, req.RollappId, req.Index)
		if !found {
			return nil, status.Error(codes.NotFound, "not found")
		}
		stateInfo = val
	} else if req.Height != 0 {
		val, err := k.FindStateInfoByHeight(ctx, req.RollappId, req.Height)
		if err != nil {
			return nil, err
		}
		stateInfo = *val
	}

	return &types.QueryGetStateInfoResponse{StateInfo: stateInfo}, nil
}

func (k Keeper) FindStateInfoByHeight(ctx sdk.Context, rollappId string, height uint64) (*types.StateInfo, error) {
	// check that height not zero
	if height == 0 {
		return nil, types.ErrInvalidHeight
	}

	_, found := k.GetRollapp(ctx, rollappId)
	if !found {
		return nil, types.ErrUnknownRollappID
	}

	stateInfoIndex, found := k.GetLatestStateInfoIndex(ctx, rollappId)
	if !found {
		return nil, errorsmod.Wrapf(types.ErrNotFound,
			"LatestStateInfoIndex wasn't found for rollappId=%s",
			rollappId)
	}
	// initial interval to search in
	startInfoIndex := uint64(1) // see TODO bellow
	endInfoIndex := stateInfoIndex.Index

	// get state info
	LatestStateInfo, found := k.GetStateInfo(ctx, rollappId, endInfoIndex)
	if !found {
		return nil, errorsmod.Wrapf(types.ErrNotFound,
			"StateInfo wasn't found for rollappId=%s, index=%d",
			rollappId, endInfoIndex)
	}

	// check that height exists
	if height >= LatestStateInfo.StartHeight+LatestStateInfo.NumBlocks {
		return nil, errorsmod.Wrapf(types.ErrStateNotExists,
			"rollappId=%s, height=%d",
			rollappId, height)
	}

	// check if the height belongs to this batch
	if height >= LatestStateInfo.StartHeight {
		return &LatestStateInfo, nil
	}

	maxNumberOfSteps := endInfoIndex - startInfoIndex + 1
	stepNum := uint64(0)
	for ; stepNum < maxNumberOfSteps; stepNum += 1 {
		// we know that endInfoIndex > startInfoIndex
		// otherwise the height should have been found
		if endInfoIndex <= startInfoIndex {
			return nil, errorsmod.Wrapf(types.ErrLogic,
				"endInfoIndex should be != than startInfoIndex rollappId=%s, startInfoIndex=%d, endInfoIndex=%d",
				rollappId, startInfoIndex, endInfoIndex)
		}
		// 1. get state info
		startStateInfo, found := k.GetStateInfo(ctx, rollappId, startInfoIndex)
		if !found {
			// TODO:
			// if stateInfo is missing it won't be logic error if history deletion be implemented
			// for that we will have to check the oldest we have
			return nil, errorsmod.Wrapf(types.ErrNotFound,
				"StateInfo wasn't found for rollappId=%s, index=%d",
				rollappId, startInfoIndex)
		}
		endStateInfo, found := k.GetStateInfo(ctx, rollappId, endInfoIndex)
		if !found {
			return nil, errorsmod.Wrapf(types.ErrNotFound,
				"StateInfo wasn't found for rollappId=%s, index=%d",
				rollappId, endInfoIndex)
		}
		startHeight := startStateInfo.StartHeight
		endHeight := endStateInfo.StartHeight + endStateInfo.NumBlocks - 1

		// 2. check startStateInfo
		if height >= startStateInfo.StartHeight &&
			(startStateInfo.StartHeight+startStateInfo.NumBlocks) > height {
			return &startStateInfo, nil
		}

		// 3. check endStateInfo
		if height >= endStateInfo.StartHeight &&
			(endStateInfo.StartHeight+endStateInfo.NumBlocks) > height {
			return &endStateInfo, nil
		}

		// 4. calculate the average blocks per batch
		avgBlocksPerBatch := (endHeight - startHeight + 1) / (endInfoIndex - startInfoIndex + 1)
		if avgBlocksPerBatch == 0 {
			return nil, errorsmod.Wrapf(types.ErrLogic,
				"avgBlocksPerBatch is zero!!! rollappId=%s, endHeight=%d, startHeight=%d, endInfoIndex=%d, startInfoIndex=%d",
				rollappId, endHeight, startHeight, endInfoIndex, startInfoIndex)
		}

		// 5. load the candidate block batch
		infoIndexStep := (height - startHeight) / avgBlocksPerBatch
		if infoIndexStep == 0 {
			infoIndexStep = 1
		}
		candidateInfoIndex := startInfoIndex + infoIndexStep
		if candidateInfoIndex > endInfoIndex {
			// skip to the last, probably the steps to big
			candidateInfoIndex = endInfoIndex
		}
		if candidateInfoIndex == endInfoIndex {
			candidateInfoIndex = endInfoIndex - 1
		}
		candidateStateInfo, found := k.GetStateInfo(ctx, rollappId, candidateInfoIndex)
		if !found {
			return nil, errorsmod.Wrapf(types.ErrNotFound,
				"StateInfo wasn't found for rollappId=%s, index=%d",
				rollappId, candidateInfoIndex)
		}

		// 6. check the candidate
		if candidateStateInfo.StartHeight > height {
			endInfoIndex = candidateInfoIndex - 1
		} else {
			if candidateStateInfo.StartHeight+candidateStateInfo.NumBlocks-1 < height {
				startInfoIndex = candidateInfoIndex + 1
			} else {
				return &candidateStateInfo, nil
			}
		}
	}

	return nil, errorsmod.Wrapf(types.ErrLogic,
		"More searching steps than indexes! rollappId=%s, stepNum=%d, maxNumberOfSteps=%d",
		rollappId, stepNum, maxNumberOfSteps)
}
