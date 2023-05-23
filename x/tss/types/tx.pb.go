// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tss/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_bandprotocol_chain_v2_pkg_tss "github.com/bandprotocol/chain/v2/pkg/tss"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// MsgCreateGroup is the Msg/CreateGroup request type.
type MsgCreateGroup struct {
	// members is a list of members in this group.
	Members []string `protobuf:"bytes,1,rep,name=members,proto3" json:"members,omitempty"`
	// threshold is a minimum number of members required to produce a
	// signature.
	Threshold uint64 `protobuf:"varint,2,opt,name=threshold,proto3" json:"threshold,omitempty"`
	// sender is the signer of this message.
	Sender string `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty"`
}

func (m *MsgCreateGroup) Reset()         { *m = MsgCreateGroup{} }
func (m *MsgCreateGroup) String() string { return proto.CompactTextString(m) }
func (*MsgCreateGroup) ProtoMessage()    {}
func (*MsgCreateGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_58d13e1023e3ffaf, []int{0}
}
func (m *MsgCreateGroup) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateGroup.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateGroup.Merge(m, src)
}
func (m *MsgCreateGroup) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateGroup.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateGroup proto.InternalMessageInfo

func (m *MsgCreateGroup) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

func (m *MsgCreateGroup) GetThreshold() uint64 {
	if m != nil {
		return m.Threshold
	}
	return 0
}

func (m *MsgCreateGroup) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

// MsgCreateGroupResponse is the Msg/CreateGroup response type.
type MsgCreateGroupResponse struct {
}

func (m *MsgCreateGroupResponse) Reset()         { *m = MsgCreateGroupResponse{} }
func (m *MsgCreateGroupResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateGroupResponse) ProtoMessage()    {}
func (*MsgCreateGroupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_58d13e1023e3ffaf, []int{1}
}
func (m *MsgCreateGroupResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateGroupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateGroupResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateGroupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateGroupResponse.Merge(m, src)
}
func (m *MsgCreateGroupResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateGroupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateGroupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateGroupResponse proto.InternalMessageInfo

// MsgSubmitDKGRound1 is the Msg/SubmitDKGRound1 request type.
type MsgSubmitDKGRound1 struct {
	// group_id is ID of the group.
	GroupID github_com_bandprotocol_chain_v2_pkg_tss.GroupID `protobuf:"varint,1,opt,name=group_id,json=groupId,proto3,casttype=github.com/bandprotocol/chain/v2/pkg/tss.GroupID" json:"group_id,omitempty"`
	// round1_data is all data that require to handle round1.
	Round1Data Round1Data `protobuf:"bytes,2,opt,name=round1_data,json=round1Data,proto3" json:"round1_data"`
	// member is the signer of this message. Must be the member of this group.
	Member string `protobuf:"bytes,3,opt,name=member,proto3" json:"member,omitempty"`
}

func (m *MsgSubmitDKGRound1) Reset()         { *m = MsgSubmitDKGRound1{} }
func (m *MsgSubmitDKGRound1) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitDKGRound1) ProtoMessage()    {}
func (*MsgSubmitDKGRound1) Descriptor() ([]byte, []int) {
	return fileDescriptor_58d13e1023e3ffaf, []int{2}
}
func (m *MsgSubmitDKGRound1) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitDKGRound1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitDKGRound1.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitDKGRound1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitDKGRound1.Merge(m, src)
}
func (m *MsgSubmitDKGRound1) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitDKGRound1) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitDKGRound1.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitDKGRound1 proto.InternalMessageInfo

func (m *MsgSubmitDKGRound1) GetGroupID() github_com_bandprotocol_chain_v2_pkg_tss.GroupID {
	if m != nil {
		return m.GroupID
	}
	return 0
}

func (m *MsgSubmitDKGRound1) GetRound1Data() Round1Data {
	if m != nil {
		return m.Round1Data
	}
	return Round1Data{}
}

func (m *MsgSubmitDKGRound1) GetMember() string {
	if m != nil {
		return m.Member
	}
	return ""
}

// MsgSubmitDKGRound1Response is the Msg/SubmitDKGRound1 response type.
type MsgSubmitDKGRound1Response struct {
}

func (m *MsgSubmitDKGRound1Response) Reset()         { *m = MsgSubmitDKGRound1Response{} }
func (m *MsgSubmitDKGRound1Response) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitDKGRound1Response) ProtoMessage()    {}
func (*MsgSubmitDKGRound1Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_58d13e1023e3ffaf, []int{3}
}
func (m *MsgSubmitDKGRound1Response) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitDKGRound1Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitDKGRound1Response.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitDKGRound1Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitDKGRound1Response.Merge(m, src)
}
func (m *MsgSubmitDKGRound1Response) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitDKGRound1Response) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitDKGRound1Response.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitDKGRound1Response proto.InternalMessageInfo

// MsgSubmitDKGRound2 is the Msg/SubmitDKGRound2 request type.
type MsgSubmitDKGRound2 struct {
	// group_id is ID of the group.
	GroupID github_com_bandprotocol_chain_v2_pkg_tss.GroupID `protobuf:"varint,1,opt,name=group_id,json=groupId,proto3,casttype=github.com/bandprotocol/chain/v2/pkg/tss.GroupID" json:"group_id,omitempty"`
	// round2_data is is all data that require to handle round2.
	Round2Data Round2Data `protobuf:"bytes,2,opt,name=round2_data,json=round2Data,proto3" json:"round2_data"`
	// member is the signer of this message. Must be the member of this group.
	Member string `protobuf:"bytes,3,opt,name=member,proto3" json:"member,omitempty"`
}

func (m *MsgSubmitDKGRound2) Reset()         { *m = MsgSubmitDKGRound2{} }
func (m *MsgSubmitDKGRound2) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitDKGRound2) ProtoMessage()    {}
func (*MsgSubmitDKGRound2) Descriptor() ([]byte, []int) {
	return fileDescriptor_58d13e1023e3ffaf, []int{4}
}
func (m *MsgSubmitDKGRound2) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitDKGRound2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitDKGRound2.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitDKGRound2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitDKGRound2.Merge(m, src)
}
func (m *MsgSubmitDKGRound2) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitDKGRound2) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitDKGRound2.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitDKGRound2 proto.InternalMessageInfo

func (m *MsgSubmitDKGRound2) GetGroupID() github_com_bandprotocol_chain_v2_pkg_tss.GroupID {
	if m != nil {
		return m.GroupID
	}
	return 0
}

func (m *MsgSubmitDKGRound2) GetRound2Data() Round2Data {
	if m != nil {
		return m.Round2Data
	}
	return Round2Data{}
}

func (m *MsgSubmitDKGRound2) GetMember() string {
	if m != nil {
		return m.Member
	}
	return ""
}

type MsgSubmitDKGRound2Response struct {
}

func (m *MsgSubmitDKGRound2Response) Reset()         { *m = MsgSubmitDKGRound2Response{} }
func (m *MsgSubmitDKGRound2Response) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitDKGRound2Response) ProtoMessage()    {}
func (*MsgSubmitDKGRound2Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_58d13e1023e3ffaf, []int{5}
}
func (m *MsgSubmitDKGRound2Response) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitDKGRound2Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitDKGRound2Response.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitDKGRound2Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitDKGRound2Response.Merge(m, src)
}
func (m *MsgSubmitDKGRound2Response) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitDKGRound2Response) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitDKGRound2Response.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitDKGRound2Response proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateGroup)(nil), "tss.v1beta1.MsgCreateGroup")
	proto.RegisterType((*MsgCreateGroupResponse)(nil), "tss.v1beta1.MsgCreateGroupResponse")
	proto.RegisterType((*MsgSubmitDKGRound1)(nil), "tss.v1beta1.MsgSubmitDKGRound1")
	proto.RegisterType((*MsgSubmitDKGRound1Response)(nil), "tss.v1beta1.MsgSubmitDKGRound1Response")
	proto.RegisterType((*MsgSubmitDKGRound2)(nil), "tss.v1beta1.MsgSubmitDKGRound2")
	proto.RegisterType((*MsgSubmitDKGRound2Response)(nil), "tss.v1beta1.MsgSubmitDKGRound2Response")
}

func init() { proto.RegisterFile("tss/v1beta1/tx.proto", fileDescriptor_58d13e1023e3ffaf) }

var fileDescriptor_58d13e1023e3ffaf = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0xeb, 0x75, 0x5a, 0xa9, 0x2b, 0x81, 0x64, 0x8d, 0x11, 0xc2, 0x94, 0x56, 0xe5, 0x40,
	0x0f, 0x28, 0xa6, 0xe6, 0x0e, 0x52, 0xa9, 0x98, 0x26, 0x54, 0x21, 0x85, 0xdb, 0x38, 0x0c, 0xa7,
	0x31, 0x6e, 0x44, 0x13, 0x47, 0xb6, 0x33, 0x6d, 0xdf, 0x62, 0x1f, 0xab, 0xc7, 0x1d, 0x39, 0x55,
	0x28, 0xfd, 0x16, 0x9c, 0x50, 0x9c, 0x64, 0x2c, 0x9d, 0x68, 0xc4, 0x65, 0xb7, 0xbc, 0xff, 0x7b,
	0xef, 0xff, 0xf2, 0x7b, 0x7e, 0xf0, 0x50, 0x2b, 0x85, 0x2f, 0xc6, 0x3e, 0xd3, 0x74, 0x8c, 0xf5,
	0xa5, 0x9b, 0x48, 0xa1, 0x05, 0xea, 0x69, 0xa5, 0xdc, 0x52, 0xb5, 0x0f, 0xb9, 0xe0, 0xc2, 0xe8,
	0x38, 0xff, 0x2a, 0x4a, 0xec, 0xe7, 0x5c, 0x08, 0xbe, 0x64, 0xd8, 0x44, 0x7e, 0xfa, 0x1d, 0xd3,
	0xf8, 0xaa, 0x4c, 0x3d, 0xad, 0x79, 0x2a, 0x55, 0xc8, 0xc3, 0x6f, 0xf0, 0xf1, 0x4c, 0xf1, 0x0f,
	0x92, 0x51, 0xcd, 0x4e, 0xa4, 0x48, 0x13, 0x64, 0xc1, 0x4e, 0xc4, 0x22, 0x9f, 0x49, 0x65, 0x81,
	0x41, 0x7b, 0xd4, 0xf5, 0xaa, 0x10, 0x1d, 0xc3, 0xae, 0x5e, 0x48, 0xa6, 0x16, 0x62, 0x19, 0x58,
	0x7b, 0x03, 0x30, 0xda, 0xf7, 0xfe, 0x0a, 0xe8, 0x08, 0x1e, 0x28, 0x16, 0x07, 0x4c, 0x5a, 0xed,
	0x01, 0x18, 0x75, 0xbd, 0x32, 0x1a, 0x5a, 0xf0, 0xa8, 0x3e, 0xc1, 0x63, 0x2a, 0x11, 0xb1, 0x62,
	0xc3, 0x15, 0x80, 0x68, 0xa6, 0xf8, 0x97, 0xd4, 0x8f, 0x42, 0x3d, 0xfd, 0x74, 0xe2, 0x89, 0x34,
	0x0e, 0xc6, 0xe8, 0x0c, 0x3e, 0xe2, 0x79, 0xdd, 0x79, 0x18, 0x58, 0x20, 0x9f, 0x32, 0x79, 0x9f,
	0xad, 0xfb, 0x1d, 0xd3, 0x7b, 0x3a, 0xfd, 0xbd, 0xee, 0xbf, 0xe1, 0xa1, 0x5e, 0xa4, 0xbe, 0x3b,
	0x17, 0x11, 0xf6, 0x69, 0x1c, 0x18, 0x92, 0xb9, 0x58, 0xe2, 0xf9, 0x82, 0x86, 0x31, 0xbe, 0x20,
	0x38, 0xf9, 0xc1, 0x0d, 0x63, 0xd9, 0xe3, 0x75, 0x8c, 0xe1, 0x69, 0x80, 0xde, 0xc1, 0x9e, 0x34,
	0x53, 0xce, 0x03, 0xaa, 0xa9, 0x81, 0xe8, 0x91, 0x67, 0xee, 0x9d, 0xcd, 0xba, 0xc5, 0x5f, 0x4c,
	0xa9, 0xa6, 0x93, 0xfd, 0xd5, 0xba, 0xdf, 0xf2, 0xa0, 0xbc, 0x55, 0x72, 0xc8, 0x62, 0x1b, 0x15,
	0x64, 0x11, 0x0d, 0x8f, 0xa1, 0x7d, 0x9f, 0x64, 0x37, 0x28, 0x79, 0x10, 0x50, 0xd2, 0x00, 0x4a,
	0xee, 0x81, 0x92, 0xff, 0x06, 0x25, 0x15, 0x28, 0xb9, 0xde, 0x83, 0xed, 0x99, 0xe2, 0xe8, 0x33,
	0xec, 0xdd, 0x3d, 0xa9, 0x17, 0xb5, 0xb9, 0xf5, 0x6b, 0xb0, 0x5f, 0xee, 0x48, 0x56, 0xc6, 0xe8,
	0x2b, 0x7c, 0xb2, 0x7d, 0x26, 0xfd, 0xed, 0xbe, 0xad, 0x02, 0xfb, 0x55, 0x43, 0xc1, 0xbf, 0xcd,
	0x49, 0x93, 0x39, 0x69, 0x32, 0xbf, 0x5d, 0xc9, 0xe4, 0xe3, 0x2a, 0x73, 0xc0, 0x4d, 0xe6, 0x80,
	0x5f, 0x99, 0x03, 0xae, 0x37, 0x4e, 0xeb, 0x66, 0xe3, 0xb4, 0x7e, 0x6e, 0x9c, 0xd6, 0xd9, 0xeb,
	0xc6, 0xd7, 0xbd, 0xcc, 0xdf, 0x16, 0xeb, 0xab, 0x84, 0x29, 0xff, 0xc0, 0xa4, 0xdf, 0xfe, 0x09,
	0x00, 0x00, 0xff, 0xff, 0x4c, 0x85, 0x15, 0x8f, 0x1c, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// CreateGroup creates a new group with a list of members.
	CreateGroup(ctx context.Context, in *MsgCreateGroup, opts ...grpc.CallOption) (*MsgCreateGroupResponse, error)
	// SubmmitDKGRound1 submit dkg for compute round1.
	SubmitDKGRound1(ctx context.Context, in *MsgSubmitDKGRound1, opts ...grpc.CallOption) (*MsgSubmitDKGRound1Response, error)
	// SubmitDKGRound2 submmit dkg for compute round 2.
	SubmitDKGRound2(ctx context.Context, in *MsgSubmitDKGRound2, opts ...grpc.CallOption) (*MsgSubmitDKGRound2Response, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateGroup(ctx context.Context, in *MsgCreateGroup, opts ...grpc.CallOption) (*MsgCreateGroupResponse, error) {
	out := new(MsgCreateGroupResponse)
	err := c.cc.Invoke(ctx, "/tss.v1beta1.Msg/CreateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SubmitDKGRound1(ctx context.Context, in *MsgSubmitDKGRound1, opts ...grpc.CallOption) (*MsgSubmitDKGRound1Response, error) {
	out := new(MsgSubmitDKGRound1Response)
	err := c.cc.Invoke(ctx, "/tss.v1beta1.Msg/SubmitDKGRound1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SubmitDKGRound2(ctx context.Context, in *MsgSubmitDKGRound2, opts ...grpc.CallOption) (*MsgSubmitDKGRound2Response, error) {
	out := new(MsgSubmitDKGRound2Response)
	err := c.cc.Invoke(ctx, "/tss.v1beta1.Msg/SubmitDKGRound2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// CreateGroup creates a new group with a list of members.
	CreateGroup(context.Context, *MsgCreateGroup) (*MsgCreateGroupResponse, error)
	// SubmmitDKGRound1 submit dkg for compute round1.
	SubmitDKGRound1(context.Context, *MsgSubmitDKGRound1) (*MsgSubmitDKGRound1Response, error)
	// SubmitDKGRound2 submmit dkg for compute round 2.
	SubmitDKGRound2(context.Context, *MsgSubmitDKGRound2) (*MsgSubmitDKGRound2Response, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateGroup(ctx context.Context, req *MsgCreateGroup) (*MsgCreateGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroup not implemented")
}
func (*UnimplementedMsgServer) SubmitDKGRound1(ctx context.Context, req *MsgSubmitDKGRound1) (*MsgSubmitDKGRound1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitDKGRound1 not implemented")
}
func (*UnimplementedMsgServer) SubmitDKGRound2(ctx context.Context, req *MsgSubmitDKGRound2) (*MsgSubmitDKGRound2Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitDKGRound2 not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tss.v1beta1.Msg/CreateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateGroup(ctx, req.(*MsgCreateGroup))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SubmitDKGRound1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSubmitDKGRound1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitDKGRound1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tss.v1beta1.Msg/SubmitDKGRound1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitDKGRound1(ctx, req.(*MsgSubmitDKGRound1))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SubmitDKGRound2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSubmitDKGRound2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitDKGRound2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tss.v1beta1.Msg/SubmitDKGRound2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitDKGRound2(ctx, req.(*MsgSubmitDKGRound2))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tss.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGroup",
			Handler:    _Msg_CreateGroup_Handler,
		},
		{
			MethodName: "SubmitDKGRound1",
			Handler:    _Msg_SubmitDKGRound1_Handler,
		},
		{
			MethodName: "SubmitDKGRound2",
			Handler:    _Msg_SubmitDKGRound2_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tss/v1beta1/tx.proto",
}

func (m *MsgCreateGroup) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateGroup) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateGroup) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Threshold != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Threshold))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Members) > 0 {
		for iNdEx := len(m.Members) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Members[iNdEx])
			copy(dAtA[i:], m.Members[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.Members[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateGroupResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateGroupResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateGroupResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgSubmitDKGRound1) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitDKGRound1) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitDKGRound1) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Member) > 0 {
		i -= len(m.Member)
		copy(dAtA[i:], m.Member)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Member)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size, err := m.Round1Data.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.GroupID != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.GroupID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgSubmitDKGRound1Response) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitDKGRound1Response) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitDKGRound1Response) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgSubmitDKGRound2) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitDKGRound2) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitDKGRound2) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Member) > 0 {
		i -= len(m.Member)
		copy(dAtA[i:], m.Member)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Member)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size, err := m.Round2Data.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.GroupID != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.GroupID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgSubmitDKGRound2Response) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitDKGRound2Response) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitDKGRound2Response) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateGroup) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Members) > 0 {
		for _, s := range m.Members {
			l = len(s)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if m.Threshold != 0 {
		n += 1 + sovTx(uint64(m.Threshold))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateGroupResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgSubmitDKGRound1) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GroupID != 0 {
		n += 1 + sovTx(uint64(m.GroupID))
	}
	l = m.Round1Data.Size()
	n += 1 + l + sovTx(uint64(l))
	l = len(m.Member)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSubmitDKGRound1Response) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgSubmitDKGRound2) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GroupID != 0 {
		n += 1 + sovTx(uint64(m.GroupID))
	}
	l = m.Round2Data.Size()
	n += 1 + l + sovTx(uint64(l))
	l = len(m.Member)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSubmitDKGRound2Response) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateGroup) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateGroup: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateGroup: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Members", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Members = append(m.Members, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Threshold", wireType)
			}
			m.Threshold = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Threshold |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCreateGroupResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateGroupResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateGroupResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSubmitDKGRound1) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSubmitDKGRound1: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitDKGRound1: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupID", wireType)
			}
			m.GroupID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GroupID |= github_com_bandprotocol_chain_v2_pkg_tss.GroupID(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Round1Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Round1Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Member", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Member = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSubmitDKGRound1Response) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSubmitDKGRound1Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitDKGRound1Response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSubmitDKGRound2) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSubmitDKGRound2: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitDKGRound2: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupID", wireType)
			}
			m.GroupID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GroupID |= github_com_bandprotocol_chain_v2_pkg_tss.GroupID(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Round2Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Round2Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Member", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Member = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSubmitDKGRound2Response) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSubmitDKGRound2Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitDKGRound2Response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
