// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nucleic/rollapp/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the rollapp module's genesis state.
type GenesisState struct {
	Params                             Params                           `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	RollappList                        []Rollapp                        `protobuf:"bytes,2,rep,name=rollapp_list,json=rollappList,proto3" json:"rollapp_list"`
	StateInfoList                      []StateInfo                      `protobuf:"bytes,3,rep,name=state_info_list,json=stateInfoList,proto3" json:"state_info_list"`
	LatestStateInfoIndexList           []StateInfoIndex                 `protobuf:"bytes,4,rep,name=latest_state_info_index_list,json=latestStateInfoIndexList,proto3" json:"latest_state_info_index_list"`
	LatestFinalizedStateIndexList      []StateInfoIndex                 `protobuf:"bytes,5,rep,name=latest_finalized_state_index_list,json=latestFinalizedStateIndexList,proto3" json:"latest_finalized_state_index_list"`
	BlockHeightToFinalizationQueueList []BlockHeightToFinalizationQueue `protobuf:"bytes,6,rep,name=block_height_to_finalization_queue_list,json=blockHeightToFinalizationQueueList,proto3" json:"block_height_to_finalization_queue_list"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_165a9e3e67ef0cbd, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetRollappList() []Rollapp {
	if m != nil {
		return m.RollappList
	}
	return nil
}

func (m *GenesisState) GetStateInfoList() []StateInfo {
	if m != nil {
		return m.StateInfoList
	}
	return nil
}

func (m *GenesisState) GetLatestStateInfoIndexList() []StateInfoIndex {
	if m != nil {
		return m.LatestStateInfoIndexList
	}
	return nil
}

func (m *GenesisState) GetLatestFinalizedStateIndexList() []StateInfoIndex {
	if m != nil {
		return m.LatestFinalizedStateIndexList
	}
	return nil
}

func (m *GenesisState) GetBlockHeightToFinalizationQueueList() []BlockHeightToFinalizationQueue {
	if m != nil {
		return m.BlockHeightToFinalizationQueueList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "nucleic.rollapp.GenesisState")
}

func init() { proto.RegisterFile("nucleic/rollapp/genesis.proto", fileDescriptor_165a9e3e67ef0cbd) }

var fileDescriptor_165a9e3e67ef0cbd = []byte{
	// 400 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xb1, 0x6e, 0xda, 0x40,
	0x18, 0xc7, 0xed, 0x42, 0x19, 0x0e, 0x2a, 0x24, 0xab, 0x52, 0x2d, 0x0b, 0x0c, 0x65, 0x29, 0x4b,
	0x6d, 0x89, 0xaa, 0x0f, 0x50, 0x06, 0x0a, 0x12, 0x43, 0x4b, 0x3b, 0x75, 0xb1, 0x6c, 0x73, 0x98,
	0x13, 0xe6, 0xce, 0xf5, 0x9d, 0x15, 0x92, 0x47, 0xc8, 0x94, 0xb7, 0xc9, 0x2b, 0x30, 0x32, 0x66,
	0x8a, 0x22, 0x78, 0x91, 0xc8, 0x77, 0x9f, 0x09, 0xc1, 0x4a, 0xa4, 0x4c, 0x67, 0xdf, 0xff, 0xfb,
	0x7e, 0xbf, 0xff, 0x70, 0xa8, 0x4d, 0xb3, 0x30, 0xc6, 0x24, 0x74, 0x53, 0x16, 0xc7, 0x7e, 0x92,
	0xb8, 0x11, 0xa6, 0x98, 0x13, 0xee, 0x24, 0x29, 0x13, 0xcc, 0x68, 0x42, 0xec, 0x40, 0x6c, 0x7d,
	0x8c, 0x58, 0xc4, 0x64, 0xe6, 0xe6, 0x5f, 0x6a, 0xcc, 0x6a, 0x9d, 0x53, 0x12, 0x3f, 0xf5, 0xd7,
	0x00, 0xb1, 0x4a, 0x0e, 0x38, 0x21, 0xee, 0x9e, 0xc7, 0x5c, 0xf8, 0x02, 0x7b, 0x84, 0x2e, 0x00,
	0xdf, 0xbb, 0xad, 0xa2, 0xc6, 0x4f, 0xd5, 0xeb, 0x4f, 0x9e, 0x19, 0xdf, 0x51, 0x4d, 0x19, 0x4c,
	0xbd, 0xab, 0xf7, 0xeb, 0x83, 0x4f, 0xce, 0x59, 0x4f, 0xe7, 0x97, 0x8c, 0x87, 0xd5, 0xed, 0x7d,
	0x47, 0x9b, 0xc1, 0xb0, 0xf1, 0x03, 0x35, 0x20, 0xf7, 0x62, 0xc2, 0x85, 0xf9, 0xae, 0x5b, 0xe9,
	0xd7, 0x07, 0x66, 0x69, 0x79, 0xa6, 0x4e, 0xd8, 0xae, 0xc3, 0xf5, 0x94, 0x70, 0x61, 0x8c, 0x51,
	0xf3, 0xa9, 0x9e, 0xa2, 0x54, 0x24, 0xc5, 0x2a, 0x51, 0x64, 0xd5, 0x09, 0x5d, 0x30, 0xe0, 0x7c,
	0xe0, 0xc5, 0x85, 0x24, 0x61, 0xd4, 0x8a, 0x7d, 0x81, 0xb9, 0xf0, 0x4e, 0x80, 0x84, 0xce, 0xf1,
	0x46, 0x61, 0xab, 0x12, 0xdb, 0x79, 0x19, 0x3b, 0xc9, 0x67, 0x81, 0x6d, 0x2a, 0xd4, 0xf3, 0x4c,
	0x6a, 0x18, 0xfa, 0x0c, 0x9a, 0x05, 0xa1, 0x7e, 0x4c, 0xae, 0xf0, 0xfc, 0x28, 0x3c, 0xba, 0xde,
	0xbf, 0xc5, 0xd5, 0x56, 0xbc, 0x51, 0x81, 0x83, 0xa1, 0x42, 0x78, 0xad, 0xa3, 0x2f, 0x41, 0xcc,
	0xc2, 0x95, 0xb7, 0xc4, 0x24, 0x5a, 0x0a, 0x4f, 0xb0, 0x42, 0xed, 0x0b, 0xc2, 0xa8, 0xf7, 0x3f,
	0xc3, 0x19, 0x56, 0xde, 0x9a, 0xf4, 0xba, 0x25, 0xef, 0x30, 0xdf, 0x1f, 0xcb, 0xf5, 0xbf, 0x6c,
	0x74, 0xb2, 0xfc, 0x3b, 0xdf, 0x85, 0x1e, 0xbd, 0xe0, 0xd5, 0xa9, 0xbc, 0xcc, 0x70, 0xba, 0xdd,
	0xdb, 0xfa, 0x6e, 0x6f, 0xeb, 0x0f, 0x7b, 0x5b, 0xbf, 0x39, 0xd8, 0xda, 0xee, 0x60, 0x6b, 0x77,
	0x07, 0x5b, 0xfb, 0x37, 0x88, 0x88, 0x58, 0x66, 0x81, 0x13, 0xb2, 0xb5, 0x0b, 0xfa, 0xaf, 0x14,
	0x8b, 0x0b, 0x96, 0xae, 0x8a, 0x7f, 0x77, 0x73, 0x7c, 0x92, 0xe2, 0x32, 0xc1, 0x3c, 0xa8, 0xc9,
	0xe7, 0xf8, 0xed, 0x31, 0x00, 0x00, 0xff, 0xff, 0x68, 0x84, 0x55, 0x48, 0x35, 0x03, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BlockHeightToFinalizationQueueList) > 0 {
		for iNdEx := len(m.BlockHeightToFinalizationQueueList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BlockHeightToFinalizationQueueList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.LatestFinalizedStateIndexList) > 0 {
		for iNdEx := len(m.LatestFinalizedStateIndexList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LatestFinalizedStateIndexList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.LatestStateInfoIndexList) > 0 {
		for iNdEx := len(m.LatestStateInfoIndexList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LatestStateInfoIndexList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.StateInfoList) > 0 {
		for iNdEx := len(m.StateInfoList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.StateInfoList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.RollappList) > 0 {
		for iNdEx := len(m.RollappList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RollappList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.RollappList) > 0 {
		for _, e := range m.RollappList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.StateInfoList) > 0 {
		for _, e := range m.StateInfoList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.LatestStateInfoIndexList) > 0 {
		for _, e := range m.LatestStateInfoIndexList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.LatestFinalizedStateIndexList) > 0 {
		for _, e := range m.LatestFinalizedStateIndexList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.BlockHeightToFinalizationQueueList) > 0 {
		for _, e := range m.BlockHeightToFinalizationQueueList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RollappList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RollappList = append(m.RollappList, Rollapp{})
			if err := m.RollappList[len(m.RollappList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StateInfoList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StateInfoList = append(m.StateInfoList, StateInfo{})
			if err := m.StateInfoList[len(m.StateInfoList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestStateInfoIndexList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LatestStateInfoIndexList = append(m.LatestStateInfoIndexList, StateInfoIndex{})
			if err := m.LatestStateInfoIndexList[len(m.LatestStateInfoIndexList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestFinalizedStateIndexList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LatestFinalizedStateIndexList = append(m.LatestFinalizedStateIndexList, StateInfoIndex{})
			if err := m.LatestFinalizedStateIndexList[len(m.LatestFinalizedStateIndexList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeightToFinalizationQueueList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlockHeightToFinalizationQueueList = append(m.BlockHeightToFinalizationQueueList, BlockHeightToFinalizationQueue{})
			if err := m.BlockHeightToFinalizationQueueList[len(m.BlockHeightToFinalizationQueueList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
