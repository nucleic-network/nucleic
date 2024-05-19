package keeper

import (
	"context"

	"github.com/eve-network/eve/x/rollapp/types"
)

type msgServer struct {
	Keeper
}

// CreateRollapp implements types.MsgServer.
func (m msgServer) CreateRollapp(context.Context, *types.MsgCreateRollapp) (*types.MsgCreateRollappResponse, error) {
	panic("unimplemented")
}

// TriggerGenesisEvent implements types.MsgServer.
func (m msgServer) TriggerGenesisEvent(context.Context, *types.MsgRollappGenesisEvent) (*types.MsgRollappGenesisEventResponse, error) {
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
