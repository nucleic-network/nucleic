package keeper

import (
	"context"

	"github.com/eve-network/eve/x/rollup/types"
)

type msgServer struct {
	Keeper
}

// CreateRollup implements types.MsgServer.
func (m msgServer) CreateRollup(context.Context, *types.MsgCreateRollup) (*types.MsgCreateRollupResponse, error) {
	panic("unimplemented")
}

// UpdateState implements types.MsgServer.
func (m msgServer) UpdateState(context.Context, *types.MsgUpdateState) (*types.MsgUpdateStateResponse, error) {
	panic("unimplemented")
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}
