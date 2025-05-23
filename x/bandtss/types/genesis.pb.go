// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: band/bandtss/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the bandtss module's genesis state.
type GenesisState struct {
	// params defines all the parameters of the module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// members is an array containing members information.
	Members []Member `protobuf:"bytes,2,rep,name=members,proto3" json:"members"`
	// current_group is the current group information.
	CurrentGroup CurrentGroup `protobuf:"bytes,3,opt,name=current_group,json=currentGroup,proto3" json:"current_group"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fe6a4345855d3c9, []int{0}
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

func (m *GenesisState) GetMembers() []Member {
	if m != nil {
		return m.Members
	}
	return nil
}

func (m *GenesisState) GetCurrentGroup() CurrentGroup {
	if m != nil {
		return m.CurrentGroup
	}
	return CurrentGroup{}
}

// Params defines the set of module parameters.
type Params struct {
	// reward_percentage is the percentage of block rewards allocated to active TSS members.
	// The reward proportion is calculated after being allocated to oracle rewards.
	RewardPercentage uint64 `protobuf:"varint,1,opt,name=reward_percentage,json=rewardPercentage,proto3" json:"reward_percentage,omitempty"`
	// inactive_penalty_duration is the duration where a member cannot activate back after being set to inactive.
	InactivePenaltyDuration time.Duration `protobuf:"bytes,2,opt,name=inactive_penalty_duration,json=inactivePenaltyDuration,proto3,stdduration" json:"inactive_penalty_duration"`
	// min_transition_duration is the minimum duration that the transition process waits before execution.
	MinTransitionDuration time.Duration `protobuf:"bytes,3,opt,name=min_transition_duration,json=minTransitionDuration,proto3,stdduration" json:"min_transition_duration"`
	// max_transition_duration is the maximum duration that the transition process waits before execution.
	MaxTransitionDuration time.Duration `protobuf:"bytes,4,opt,name=max_transition_duration,json=maxTransitionDuration,proto3,stdduration" json:"max_transition_duration"`
	// fee_per_signer is the tokens that will be paid per signer.
	FeePerSigner github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,5,rep,name=fee_per_signer,json=feePerSigner,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"fee_per_signer"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fe6a4345855d3c9, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetRewardPercentage() uint64 {
	if m != nil {
		return m.RewardPercentage
	}
	return 0
}

func (m *Params) GetInactivePenaltyDuration() time.Duration {
	if m != nil {
		return m.InactivePenaltyDuration
	}
	return 0
}

func (m *Params) GetMinTransitionDuration() time.Duration {
	if m != nil {
		return m.MinTransitionDuration
	}
	return 0
}

func (m *Params) GetMaxTransitionDuration() time.Duration {
	if m != nil {
		return m.MaxTransitionDuration
	}
	return 0
}

func (m *Params) GetFeePerSigner() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.FeePerSigner
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "band.bandtss.v1beta1.GenesisState")
	proto.RegisterType((*Params)(nil), "band.bandtss.v1beta1.Params")
}

func init() {
	proto.RegisterFile("band/bandtss/v1beta1/genesis.proto", fileDescriptor_3fe6a4345855d3c9)
}

var fileDescriptor_3fe6a4345855d3c9 = []byte{
	// 505 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xe3, 0x26, 0x04, 0xb4, 0x04, 0x54, 0xac, 0xa0, 0x26, 0x15, 0xb2, 0xab, 0x9c, 0x7a,
	0x61, 0xb7, 0xa5, 0xb7, 0x8a, 0x0b, 0x06, 0xa9, 0x12, 0x52, 0xa5, 0xc8, 0xe5, 0x04, 0x07, 0x6b,
	0xed, 0x4c, 0xdc, 0x15, 0xf1, 0xae, 0xd9, 0x5d, 0x87, 0xf4, 0x2d, 0x38, 0x72, 0xe4, 0xcc, 0x93,
	0xf4, 0xd8, 0x1b, 0x9c, 0x1a, 0x94, 0x5c, 0x78, 0x0c, 0xb4, 0x6b, 0x3b, 0x2d, 0xc8, 0x48, 0x70,
	0xf1, 0x9f, 0x99, 0x6f, 0x7e, 0xf3, 0xad, 0x67, 0x8c, 0x46, 0x31, 0xe5, 0x13, 0x62, 0x2e, 0x5a,
	0x29, 0x32, 0x3f, 0x8c, 0x41, 0xd3, 0x43, 0x92, 0x02, 0x07, 0xc5, 0x14, 0xce, 0xa5, 0xd0, 0xc2,
	0xed, 0x9b, 0x34, 0xae, 0x34, 0xb8, 0xd2, 0xec, 0xf6, 0x53, 0x91, 0x0a, 0x2b, 0x20, 0xe6, 0xa9,
	0xd4, 0xee, 0x7a, 0xa9, 0x10, 0xe9, 0x0c, 0x88, 0x7d, 0x8b, 0x8b, 0x29, 0x99, 0x14, 0x92, 0x6a,
	0x26, 0x78, 0x9d, 0x4f, 0x84, 0xca, 0x84, 0x22, 0x31, 0x55, 0xb0, 0x69, 0x97, 0x08, 0x56, 0xe7,
	0x9b, 0xfd, 0xd4, 0xbd, 0xad, 0x66, 0xf4, 0xcd, 0x41, 0xbd, 0x93, 0xd2, 0xe1, 0x99, 0xa6, 0x1a,
	0xdc, 0x63, 0xd4, 0xcd, 0xa9, 0xa4, 0x99, 0x1a, 0x38, 0x7b, 0xce, 0xfe, 0xfd, 0x67, 0x4f, 0x70,
	0x93, 0x63, 0x3c, 0xb6, 0x9a, 0xa0, 0x73, 0x79, 0xed, 0xb7, 0xc2, 0xaa, 0xc2, 0x7d, 0x8e, 0xee,
	0x66, 0x90, 0xc5, 0x20, 0xd5, 0x60, 0x6b, 0xaf, 0xfd, 0xf7, 0xe2, 0x53, 0x2b, 0xaa, 0x8a, 0xeb,
	0x12, 0xf7, 0x14, 0x3d, 0x48, 0x0a, 0x29, 0x81, 0xeb, 0x28, 0x95, 0xa2, 0xc8, 0x07, 0x6d, 0x6b,
	0x60, 0xd4, 0xcc, 0x78, 0x59, 0x4a, 0x4f, 0x8c, 0xb2, 0x22, 0xf5, 0x92, 0x5b, 0xb1, 0xd1, 0xb2,
	0x8d, 0xba, 0xa5, 0x4b, 0xf7, 0x05, 0x7a, 0x24, 0xe1, 0x23, 0x95, 0x93, 0x28, 0x07, 0x99, 0x00,
	0xd7, 0x34, 0x05, 0x7b, 0xbc, 0x4e, 0xd0, 0x5f, 0x5d, 0xfb, 0xdb, 0xa1, 0x4d, 0x8e, 0x37, 0xb9,
	0x70, 0x5b, 0xfe, 0x11, 0x71, 0x23, 0x34, 0x64, 0x9c, 0x26, 0x9a, 0xcd, 0x21, 0xca, 0x81, 0xd3,
	0x99, 0xbe, 0x88, 0xea, 0x71, 0x0c, 0xb6, 0xac, 0xd1, 0x21, 0x2e, 0xe7, 0x85, 0xeb, 0x79, 0xe1,
	0x57, 0x95, 0x20, 0xb8, 0x67, 0xfc, 0x7d, 0x5e, 0xfa, 0x4e, 0xb8, 0x53, 0x53, 0xc6, 0x25, 0xa4,
	0x96, 0xb8, 0xef, 0xd0, 0x4e, 0xc6, 0x78, 0xa4, 0x25, 0xe5, 0x8a, 0x99, 0xc8, 0x0d, 0xbe, 0xfd,
	0xef, 0xf8, 0xc7, 0x19, 0xe3, 0x6f, 0x36, 0x88, 0xdf, 0xe0, 0x74, 0xd1, 0x08, 0xef, 0xfc, 0x0f,
	0x9c, 0x2e, 0x1a, 0xe0, 0x1f, 0xd0, 0xc3, 0x29, 0x98, 0xaf, 0x22, 0x23, 0xc5, 0x52, 0x0e, 0x72,
	0x70, 0xc7, 0x0e, 0x7f, 0x88, 0xcb, 0xfd, 0xc4, 0x66, 0x3f, 0x6f, 0xe6, 0x26, 0x18, 0x0f, 0x0e,
	0x0c, 0xf3, 0xeb, 0xd2, 0xdf, 0x4f, 0x99, 0x3e, 0x2f, 0x62, 0x9c, 0x88, 0x8c, 0x54, 0xcb, 0x5c,
	0xde, 0x9e, 0xaa, 0xc9, 0x7b, 0xa2, 0x2f, 0x72, 0x50, 0xb6, 0x40, 0x85, 0xbd, 0x29, 0xc0, 0x18,
	0xe4, 0x99, 0x6d, 0x70, 0xdc, 0xf9, 0xf9, 0xc5, 0x77, 0x82, 0xd7, 0x97, 0x2b, 0xcf, 0xb9, 0x5a,
	0x79, 0xce, 0x8f, 0x95, 0xe7, 0x7c, 0x5a, 0x7b, 0xad, 0xab, 0xb5, 0xd7, 0xfa, 0xbe, 0xf6, 0x5a,
	0x6f, 0x0f, 0x6e, 0x71, 0xcd, 0xe2, 0xd8, 0x53, 0x25, 0x62, 0x46, 0x92, 0x73, 0xca, 0x38, 0x99,
	0x1f, 0x91, 0xc5, 0xe6, 0xbf, 0xb0, 0x5d, 0xe2, 0xae, 0x95, 0x1c, 0xfd, 0x0a, 0x00, 0x00, 0xff,
	0xff, 0xe9, 0x4e, 0xb2, 0x50, 0xc4, 0x03, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.RewardPercentage != that1.RewardPercentage {
		return false
	}
	if this.InactivePenaltyDuration != that1.InactivePenaltyDuration {
		return false
	}
	if this.MinTransitionDuration != that1.MinTransitionDuration {
		return false
	}
	if this.MaxTransitionDuration != that1.MaxTransitionDuration {
		return false
	}
	if len(this.FeePerSigner) != len(that1.FeePerSigner) {
		return false
	}
	for i := range this.FeePerSigner {
		if !this.FeePerSigner[i].Equal(&that1.FeePerSigner[i]) {
			return false
		}
	}
	return true
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
	{
		size, err := m.CurrentGroup.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Members) > 0 {
		for iNdEx := len(m.Members) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Members[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FeePerSigner) > 0 {
		for iNdEx := len(m.FeePerSigner) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FeePerSigner[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	n3, err3 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.MaxTransitionDuration, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.MaxTransitionDuration):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintGenesis(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x22
	n4, err4 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.MinTransitionDuration, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.MinTransitionDuration):])
	if err4 != nil {
		return 0, err4
	}
	i -= n4
	i = encodeVarintGenesis(dAtA, i, uint64(n4))
	i--
	dAtA[i] = 0x1a
	n5, err5 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.InactivePenaltyDuration, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.InactivePenaltyDuration):])
	if err5 != nil {
		return 0, err5
	}
	i -= n5
	i = encodeVarintGenesis(dAtA, i, uint64(n5))
	i--
	dAtA[i] = 0x12
	if m.RewardPercentage != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.RewardPercentage))
		i--
		dAtA[i] = 0x8
	}
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
	if len(m.Members) > 0 {
		for _, e := range m.Members {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.CurrentGroup.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RewardPercentage != 0 {
		n += 1 + sovGenesis(uint64(m.RewardPercentage))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.InactivePenaltyDuration)
	n += 1 + l + sovGenesis(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.MinTransitionDuration)
	n += 1 + l + sovGenesis(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.MaxTransitionDuration)
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.FeePerSigner) > 0 {
		for _, e := range m.FeePerSigner {
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
				return fmt.Errorf("proto: wrong wireType = %d for field Members", wireType)
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
			m.Members = append(m.Members, Member{})
			if err := m.Members[len(m.Members)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentGroup", wireType)
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
			if err := m.CurrentGroup.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *Params) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardPercentage", wireType)
			}
			m.RewardPercentage = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RewardPercentage |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InactivePenaltyDuration", wireType)
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
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.InactivePenaltyDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinTransitionDuration", wireType)
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
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.MinTransitionDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxTransitionDuration", wireType)
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
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.MaxTransitionDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeePerSigner", wireType)
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
			m.FeePerSigner = append(m.FeePerSigner, types.Coin{})
			if err := m.FeePerSigner[len(m.FeePerSigner)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
