package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const TypeMsgUpdateState = "update_state"

var _ sdk.Msg = &MsgUpdateState{}

func NewMsgUpdateState(creator string, rollappId string, startHeight uint64, numBlocks uint64, daPath string, version uint64, blockDescs *BlockDescriptors) *MsgUpdateState {
	return &MsgUpdateState{
		Creator:     creator,
		RollappId:   rollappId,
		StartHeight: startHeight,
		NumBlocks:   numBlocks,
		DaPath:      daPath,
		Version:     version,
		BlockDescs:  *blockDescs,
	}
}

func (msg *MsgUpdateState) Route() string {
	return RouterKey
}

func (msg *MsgUpdateState) Type() string {
	return TypeMsgUpdateState
}

func (msg *MsgUpdateState) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateState) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateState) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	// an update cann't be with no block descriptors
	if msg.NumBlocks == uint64(0) {
		return errorsmod.Wrap(ErrInvalidNumBlocks, "number of blocks can not be zero")
	}

	// check to see that update contains all BDs
	if len(msg.BlockDescs.BlockDescs) != int(msg.NumBlocks) {
		return errorsmod.Wrapf(ErrInvalidNumBlocks, "number of blocks (%d) != number of block descriptors(%d)", msg.NumBlocks, len(msg.BlockDescs.BlockDescs))
	}

	// check to see that startHeight is not zaro
	if msg.StartHeight == 0 {
		return errorsmod.Wrapf(ErrWrongBlockHeight, "StartHeight must be greater than zero")
	}

	// check that the blocks are sequential by height
	for bdIndex := uint64(0); bdIndex < msg.NumBlocks; bdIndex += 1 {
		if msg.BlockDescs.BlockDescs[bdIndex].Height != msg.StartHeight+bdIndex {
			return ErrInvalidBlockSequence
		}
		// check to see stateRoot is a 32 byte array
		if len(msg.BlockDescs.BlockDescs[bdIndex].StateRoot) != 32 {
			return errorsmod.Wrapf(ErrInvalidStateRoot, "StateRoot of block high (%d) must be 32 byte array. But received (%d) bytes",
				msg.BlockDescs.BlockDescs[bdIndex].Height, len(msg.BlockDescs.BlockDescs[bdIndex].StateRoot))
		}
	}

	return nil
}
