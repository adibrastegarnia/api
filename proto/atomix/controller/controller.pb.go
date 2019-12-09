// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: atomix/controller/controller.proto

package controller

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

// Gets a list of partition groups in the given namespace
type GetPartitionGroupsRequest struct {
	ID                   *PartitionGroupId `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetPartitionGroupsRequest) Reset()         { *m = GetPartitionGroupsRequest{} }
func (m *GetPartitionGroupsRequest) String() string { return proto.CompactTextString(m) }
func (*GetPartitionGroupsRequest) ProtoMessage()    {}
func (*GetPartitionGroupsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6541e67aaed2c172, []int{0}
}
func (m *GetPartitionGroupsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPartitionGroupsRequest.Unmarshal(m, b)
}
func (m *GetPartitionGroupsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPartitionGroupsRequest.Marshal(b, m, deterministic)
}
func (m *GetPartitionGroupsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPartitionGroupsRequest.Merge(m, src)
}
func (m *GetPartitionGroupsRequest) XXX_Size() int {
	return xxx_messageInfo_GetPartitionGroupsRequest.Size(m)
}
func (m *GetPartitionGroupsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPartitionGroupsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPartitionGroupsRequest proto.InternalMessageInfo

func (m *GetPartitionGroupsRequest) GetID() *PartitionGroupId {
	if m != nil {
		return m.ID
	}
	return nil
}

// Returns a list of partition groups
type GetPartitionGroupsResponse struct {
	Groups               []*PartitionGroup `protobuf:"bytes,1,rep,name=groups,proto3" json:"groups,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetPartitionGroupsResponse) Reset()         { *m = GetPartitionGroupsResponse{} }
func (m *GetPartitionGroupsResponse) String() string { return proto.CompactTextString(m) }
func (*GetPartitionGroupsResponse) ProtoMessage()    {}
func (*GetPartitionGroupsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6541e67aaed2c172, []int{1}
}
func (m *GetPartitionGroupsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPartitionGroupsResponse.Unmarshal(m, b)
}
func (m *GetPartitionGroupsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPartitionGroupsResponse.Marshal(b, m, deterministic)
}
func (m *GetPartitionGroupsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPartitionGroupsResponse.Merge(m, src)
}
func (m *GetPartitionGroupsResponse) XXX_Size() int {
	return xxx_messageInfo_GetPartitionGroupsResponse.Size(m)
}
func (m *GetPartitionGroupsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPartitionGroupsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPartitionGroupsResponse proto.InternalMessageInfo

func (m *GetPartitionGroupsResponse) GetGroups() []*PartitionGroup {
	if m != nil {
		return m.Groups
	}
	return nil
}

// Creates a new partition group
type CreatePartitionGroupRequest struct {
	ID                   *PartitionGroupId   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Spec                 *PartitionGroupSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *CreatePartitionGroupRequest) Reset()         { *m = CreatePartitionGroupRequest{} }
func (m *CreatePartitionGroupRequest) String() string { return proto.CompactTextString(m) }
func (*CreatePartitionGroupRequest) ProtoMessage()    {}
func (*CreatePartitionGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6541e67aaed2c172, []int{2}
}
func (m *CreatePartitionGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePartitionGroupRequest.Unmarshal(m, b)
}
func (m *CreatePartitionGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePartitionGroupRequest.Marshal(b, m, deterministic)
}
func (m *CreatePartitionGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePartitionGroupRequest.Merge(m, src)
}
func (m *CreatePartitionGroupRequest) XXX_Size() int {
	return xxx_messageInfo_CreatePartitionGroupRequest.Size(m)
}
func (m *CreatePartitionGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePartitionGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePartitionGroupRequest proto.InternalMessageInfo

func (m *CreatePartitionGroupRequest) GetID() *PartitionGroupId {
	if m != nil {
		return m.ID
	}
	return nil
}

func (m *CreatePartitionGroupRequest) GetSpec() *PartitionGroupSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

// Returns the status of a partition group
type CreatePartitionGroupResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePartitionGroupResponse) Reset()         { *m = CreatePartitionGroupResponse{} }
func (m *CreatePartitionGroupResponse) String() string { return proto.CompactTextString(m) }
func (*CreatePartitionGroupResponse) ProtoMessage()    {}
func (*CreatePartitionGroupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6541e67aaed2c172, []int{3}
}
func (m *CreatePartitionGroupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePartitionGroupResponse.Unmarshal(m, b)
}
func (m *CreatePartitionGroupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePartitionGroupResponse.Marshal(b, m, deterministic)
}
func (m *CreatePartitionGroupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePartitionGroupResponse.Merge(m, src)
}
func (m *CreatePartitionGroupResponse) XXX_Size() int {
	return xxx_messageInfo_CreatePartitionGroupResponse.Size(m)
}
func (m *CreatePartitionGroupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePartitionGroupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePartitionGroupResponse proto.InternalMessageInfo

// Deletes a partition group
type DeletePartitionGroupRequest struct {
	ID                   *PartitionGroupId `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *DeletePartitionGroupRequest) Reset()         { *m = DeletePartitionGroupRequest{} }
func (m *DeletePartitionGroupRequest) String() string { return proto.CompactTextString(m) }
func (*DeletePartitionGroupRequest) ProtoMessage()    {}
func (*DeletePartitionGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6541e67aaed2c172, []int{4}
}
func (m *DeletePartitionGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePartitionGroupRequest.Unmarshal(m, b)
}
func (m *DeletePartitionGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePartitionGroupRequest.Marshal(b, m, deterministic)
}
func (m *DeletePartitionGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePartitionGroupRequest.Merge(m, src)
}
func (m *DeletePartitionGroupRequest) XXX_Size() int {
	return xxx_messageInfo_DeletePartitionGroupRequest.Size(m)
}
func (m *DeletePartitionGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePartitionGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePartitionGroupRequest proto.InternalMessageInfo

func (m *DeletePartitionGroupRequest) GetID() *PartitionGroupId {
	if m != nil {
		return m.ID
	}
	return nil
}

// Returns the result of deleting a partition group
type DeletePartitionGroupResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePartitionGroupResponse) Reset()         { *m = DeletePartitionGroupResponse{} }
func (m *DeletePartitionGroupResponse) String() string { return proto.CompactTextString(m) }
func (*DeletePartitionGroupResponse) ProtoMessage()    {}
func (*DeletePartitionGroupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6541e67aaed2c172, []int{5}
}
func (m *DeletePartitionGroupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePartitionGroupResponse.Unmarshal(m, b)
}
func (m *DeletePartitionGroupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePartitionGroupResponse.Marshal(b, m, deterministic)
}
func (m *DeletePartitionGroupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePartitionGroupResponse.Merge(m, src)
}
func (m *DeletePartitionGroupResponse) XXX_Size() int {
	return xxx_messageInfo_DeletePartitionGroupResponse.Size(m)
}
func (m *DeletePartitionGroupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePartitionGroupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePartitionGroupResponse proto.InternalMessageInfo

// Enters a primary election
type PartitionElectionRequest struct {
	PartitionID          *PartitionId `protobuf:"bytes,1,opt,name=partition_id,json=partitionId,proto3" json:"partition_id,omitempty"`
	Member               string       `protobuf:"bytes,2,opt,name=member,proto3" json:"member,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PartitionElectionRequest) Reset()         { *m = PartitionElectionRequest{} }
func (m *PartitionElectionRequest) String() string { return proto.CompactTextString(m) }
func (*PartitionElectionRequest) ProtoMessage()    {}
func (*PartitionElectionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6541e67aaed2c172, []int{6}
}
func (m *PartitionElectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PartitionElectionRequest.Unmarshal(m, b)
}
func (m *PartitionElectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PartitionElectionRequest.Marshal(b, m, deterministic)
}
func (m *PartitionElectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartitionElectionRequest.Merge(m, src)
}
func (m *PartitionElectionRequest) XXX_Size() int {
	return xxx_messageInfo_PartitionElectionRequest.Size(m)
}
func (m *PartitionElectionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PartitionElectionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PartitionElectionRequest proto.InternalMessageInfo

func (m *PartitionElectionRequest) GetPartitionID() *PartitionId {
	if m != nil {
		return m.PartitionID
	}
	return nil
}

func (m *PartitionElectionRequest) GetMember() string {
	if m != nil {
		return m.Member
	}
	return ""
}

// Enter response
type PartitionElectionResponse struct {
	Term                 *PrimaryTerm `protobuf:"bytes,1,opt,name=term,proto3" json:"term,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PartitionElectionResponse) Reset()         { *m = PartitionElectionResponse{} }
func (m *PartitionElectionResponse) String() string { return proto.CompactTextString(m) }
func (*PartitionElectionResponse) ProtoMessage()    {}
func (*PartitionElectionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6541e67aaed2c172, []int{7}
}
func (m *PartitionElectionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PartitionElectionResponse.Unmarshal(m, b)
}
func (m *PartitionElectionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PartitionElectionResponse.Marshal(b, m, deterministic)
}
func (m *PartitionElectionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartitionElectionResponse.Merge(m, src)
}
func (m *PartitionElectionResponse) XXX_Size() int {
	return xxx_messageInfo_PartitionElectionResponse.Size(m)
}
func (m *PartitionElectionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PartitionElectionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PartitionElectionResponse proto.InternalMessageInfo

func (m *PartitionElectionResponse) GetTerm() *PrimaryTerm {
	if m != nil {
		return m.Term
	}
	return nil
}

// Primary term
type PrimaryTerm struct {
	Term                 int64    `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	Primary              string   `protobuf:"bytes,2,opt,name=primary,proto3" json:"primary,omitempty"`
	Candidates           []string `protobuf:"bytes,3,rep,name=candidates,proto3" json:"candidates,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimaryTerm) Reset()         { *m = PrimaryTerm{} }
func (m *PrimaryTerm) String() string { return proto.CompactTextString(m) }
func (*PrimaryTerm) ProtoMessage()    {}
func (*PrimaryTerm) Descriptor() ([]byte, []int) {
	return fileDescriptor_6541e67aaed2c172, []int{8}
}
func (m *PrimaryTerm) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimaryTerm.Unmarshal(m, b)
}
func (m *PrimaryTerm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimaryTerm.Marshal(b, m, deterministic)
}
func (m *PrimaryTerm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimaryTerm.Merge(m, src)
}
func (m *PrimaryTerm) XXX_Size() int {
	return xxx_messageInfo_PrimaryTerm.Size(m)
}
func (m *PrimaryTerm) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimaryTerm.DiscardUnknown(m)
}

var xxx_messageInfo_PrimaryTerm proto.InternalMessageInfo

func (m *PrimaryTerm) GetTerm() int64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *PrimaryTerm) GetPrimary() string {
	if m != nil {
		return m.Primary
	}
	return ""
}

func (m *PrimaryTerm) GetCandidates() []string {
	if m != nil {
		return m.Candidates
	}
	return nil
}

func init() {
	proto.RegisterType((*GetPartitionGroupsRequest)(nil), "atomix.controller.GetPartitionGroupsRequest")
	proto.RegisterType((*GetPartitionGroupsResponse)(nil), "atomix.controller.GetPartitionGroupsResponse")
	proto.RegisterType((*CreatePartitionGroupRequest)(nil), "atomix.controller.CreatePartitionGroupRequest")
	proto.RegisterType((*CreatePartitionGroupResponse)(nil), "atomix.controller.CreatePartitionGroupResponse")
	proto.RegisterType((*DeletePartitionGroupRequest)(nil), "atomix.controller.DeletePartitionGroupRequest")
	proto.RegisterType((*DeletePartitionGroupResponse)(nil), "atomix.controller.DeletePartitionGroupResponse")
	proto.RegisterType((*PartitionElectionRequest)(nil), "atomix.controller.PartitionElectionRequest")
	proto.RegisterType((*PartitionElectionResponse)(nil), "atomix.controller.PartitionElectionResponse")
	proto.RegisterType((*PrimaryTerm)(nil), "atomix.controller.PrimaryTerm")
}

func init() { proto.RegisterFile("atomix/controller/controller.proto", fileDescriptor_6541e67aaed2c172) }

var fileDescriptor_6541e67aaed2c172 = []byte{
	// 465 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xdd, 0x6a, 0xd4, 0x40,
	0x14, 0x6e, 0x36, 0x65, 0xa5, 0x27, 0x8a, 0x74, 0x28, 0x92, 0xa6, 0xb2, 0xdd, 0x8e, 0x08, 0x0b,
	0xd6, 0xac, 0xc4, 0xab, 0xe2, 0x5d, 0xbb, 0xa5, 0xec, 0x95, 0x32, 0x15, 0x14, 0xbd, 0x90, 0x34,
	0x39, 0x2c, 0x03, 0x49, 0x66, 0x9c, 0x4c, 0xc5, 0xbe, 0x80, 0x6f, 0xe0, 0xeb, 0xf5, 0xa2, 0x0f,
	0xe0, 0x33, 0x48, 0x7e, 0x37, 0x76, 0xc7, 0xec, 0x5e, 0xe8, 0x55, 0x66, 0xce, 0x7c, 0xdf, 0xf9,
	0xbe, 0xe4, 0x7c, 0x19, 0xa0, 0xa1, 0x16, 0x29, 0xff, 0x3e, 0x8d, 0x44, 0xa6, 0x95, 0x48, 0x12,
	0x54, 0x9d, 0xa5, 0x2f, 0x95, 0xd0, 0x82, 0xec, 0x56, 0x18, 0x7f, 0x79, 0xe0, 0x1d, 0xad, 0xd2,
	0x64, 0xa8, 0x34, 0xd7, 0x5c, 0x64, 0x15, 0xcb, 0xdb, 0x5b, 0x88, 0x85, 0x28, 0x97, 0xd3, 0x62,
	0x55, 0x55, 0xe9, 0x47, 0xd8, 0xbf, 0x40, 0xfd, 0xae, 0xc1, 0x5e, 0x28, 0x71, 0x2d, 0x73, 0x86,
	0x5f, 0xaf, 0x31, 0xd7, 0xe4, 0x0d, 0x0c, 0x78, 0xec, 0x5a, 0x63, 0x6b, 0xe2, 0x04, 0xcf, 0xfc,
	0x15, 0x55, 0xff, 0x4f, 0xda, 0x3c, 0x3e, 0x1d, 0xde, 0xdd, 0x1e, 0x0e, 0xe6, 0x33, 0x36, 0xe0,
	0x31, 0xfd, 0x00, 0x9e, 0xa9, 0x73, 0x2e, 0x45, 0x96, 0x23, 0x39, 0x81, 0xe1, 0xa2, 0xac, 0xb8,
	0xd6, 0xd8, 0x9e, 0x38, 0xc1, 0xd1, 0xda, 0xf6, 0xac, 0x26, 0xd0, 0x9f, 0x16, 0x1c, 0x9c, 0x29,
	0x0c, 0x35, 0xde, 0x03, 0xfc, 0x03, 0xd7, 0xe4, 0x04, 0xb6, 0x73, 0x89, 0x91, 0x3b, 0x28, 0xe9,
	0xcf, 0xd7, 0xd2, 0x2f, 0x25, 0x46, 0xac, 0xa4, 0xd0, 0x11, 0x3c, 0x35, 0xdb, 0xaa, 0x5e, 0x99,
	0x7e, 0x82, 0x83, 0x19, 0x26, 0xf8, 0x3f, 0x6c, 0x17, 0xda, 0xe6, 0xde, 0xb5, 0xf6, 0x0f, 0x0b,
	0xdc, 0xf6, 0xe8, 0x3c, 0xc1, 0xa8, 0x78, 0x36, 0xca, 0x0c, 0x1e, 0xb6, 0x61, 0xf9, 0xd2, 0x7a,
	0x18, 0xf5, 0x79, 0x98, 0xc7, 0xa7, 0x8f, 0xef, 0x6e, 0x0f, 0x9d, 0x65, 0x61, 0xc6, 0x1c, 0xb9,
	0x3c, 0x25, 0x4f, 0x60, 0x98, 0x62, 0x7a, 0x85, 0xaa, 0xfc, 0x92, 0x3b, 0xac, 0xde, 0xd1, 0xb7,
	0xb0, 0x6f, 0xf0, 0x51, 0x87, 0x22, 0x80, 0x6d, 0x8d, 0x2a, 0xed, 0x33, 0xa0, 0x78, 0x1a, 0xaa,
	0x9b, 0xf7, 0xa8, 0x52, 0x56, 0x62, 0xe9, 0x67, 0x70, 0x3a, 0x45, 0x42, 0x3a, 0x2d, 0xec, 0x0a,
	0x42, 0x5c, 0x78, 0x20, 0x2b, 0x48, 0x6d, 0xa6, 0xd9, 0x92, 0x11, 0x40, 0x14, 0x66, 0x31, 0x8f,
	0x43, 0x8d, 0xb9, 0x6b, 0x8f, 0xed, 0xc9, 0x0e, 0xeb, 0x54, 0x82, 0x5f, 0x36, 0xec, 0x9e, 0xb5,
	0xea, 0x97, 0xa8, 0xbe, 0xf1, 0x08, 0xc9, 0x0d, 0xec, 0x99, 0x06, 0x4d, 0x7c, 0x83, 0xe1, 0x9e,
	0xa0, 0x7a, 0xd3, 0x8d, 0xf1, 0xf5, 0x14, 0xb7, 0x0a, 0x69, 0xd3, 0x9c, 0x8d, 0xd2, 0x3d, 0x61,
	0x33, 0x4a, 0xf7, 0x06, 0x68, 0x8b, 0xe4, 0x40, 0x56, 0xff, 0x67, 0x72, 0x6c, 0x68, 0xf4, 0xd7,
	0x0b, 0xc5, 0x7b, 0xb9, 0x21, 0xba, 0x15, 0xcd, 0xe0, 0xd1, 0x79, 0xa6, 0x51, 0x35, 0x51, 0x21,
	0x2f, 0xfa, 0x52, 0x79, 0x2f, 0xd8, 0xde, 0xf1, 0x66, 0xe0, 0x46, 0xed, 0x95, 0x75, 0x35, 0x2c,
	0x6f, 0xc5, 0xd7, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x7e, 0xac, 0xfa, 0xa7, 0x87, 0x05, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ControllerServiceClient is the client API for ControllerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ControllerServiceClient interface {
	// Creates a partition group
	CreatePartitionGroup(ctx context.Context, in *CreatePartitionGroupRequest, opts ...grpc.CallOption) (*CreatePartitionGroupResponse, error)
	// Deletes a partition group
	DeletePartitionGroup(ctx context.Context, in *DeletePartitionGroupRequest, opts ...grpc.CallOption) (*DeletePartitionGroupResponse, error)
	// Gets a list of active partition groups
	GetPartitionGroups(ctx context.Context, in *GetPartitionGroupsRequest, opts ...grpc.CallOption) (*GetPartitionGroupsResponse, error)
	// Enters a primary election for a specific partition
	EnterElection(ctx context.Context, in *PartitionElectionRequest, opts ...grpc.CallOption) (ControllerService_EnterElectionClient, error)
}

type controllerServiceClient struct {
	cc *grpc.ClientConn
}

func NewControllerServiceClient(cc *grpc.ClientConn) ControllerServiceClient {
	return &controllerServiceClient{cc}
}

func (c *controllerServiceClient) CreatePartitionGroup(ctx context.Context, in *CreatePartitionGroupRequest, opts ...grpc.CallOption) (*CreatePartitionGroupResponse, error) {
	out := new(CreatePartitionGroupResponse)
	err := c.cc.Invoke(ctx, "/atomix.controller.ControllerService/CreatePartitionGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controllerServiceClient) DeletePartitionGroup(ctx context.Context, in *DeletePartitionGroupRequest, opts ...grpc.CallOption) (*DeletePartitionGroupResponse, error) {
	out := new(DeletePartitionGroupResponse)
	err := c.cc.Invoke(ctx, "/atomix.controller.ControllerService/DeletePartitionGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controllerServiceClient) GetPartitionGroups(ctx context.Context, in *GetPartitionGroupsRequest, opts ...grpc.CallOption) (*GetPartitionGroupsResponse, error) {
	out := new(GetPartitionGroupsResponse)
	err := c.cc.Invoke(ctx, "/atomix.controller.ControllerService/GetPartitionGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controllerServiceClient) EnterElection(ctx context.Context, in *PartitionElectionRequest, opts ...grpc.CallOption) (ControllerService_EnterElectionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ControllerService_serviceDesc.Streams[0], "/atomix.controller.ControllerService/EnterElection", opts...)
	if err != nil {
		return nil, err
	}
	x := &controllerServiceEnterElectionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ControllerService_EnterElectionClient interface {
	Recv() (*PartitionElectionResponse, error)
	grpc.ClientStream
}

type controllerServiceEnterElectionClient struct {
	grpc.ClientStream
}

func (x *controllerServiceEnterElectionClient) Recv() (*PartitionElectionResponse, error) {
	m := new(PartitionElectionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ControllerServiceServer is the server API for ControllerService service.
type ControllerServiceServer interface {
	// Creates a partition group
	CreatePartitionGroup(context.Context, *CreatePartitionGroupRequest) (*CreatePartitionGroupResponse, error)
	// Deletes a partition group
	DeletePartitionGroup(context.Context, *DeletePartitionGroupRequest) (*DeletePartitionGroupResponse, error)
	// Gets a list of active partition groups
	GetPartitionGroups(context.Context, *GetPartitionGroupsRequest) (*GetPartitionGroupsResponse, error)
	// Enters a primary election for a specific partition
	EnterElection(*PartitionElectionRequest, ControllerService_EnterElectionServer) error
}

// UnimplementedControllerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedControllerServiceServer struct {
}

func (*UnimplementedControllerServiceServer) CreatePartitionGroup(ctx context.Context, req *CreatePartitionGroupRequest) (*CreatePartitionGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePartitionGroup not implemented")
}
func (*UnimplementedControllerServiceServer) DeletePartitionGroup(ctx context.Context, req *DeletePartitionGroupRequest) (*DeletePartitionGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePartitionGroup not implemented")
}
func (*UnimplementedControllerServiceServer) GetPartitionGroups(ctx context.Context, req *GetPartitionGroupsRequest) (*GetPartitionGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPartitionGroups not implemented")
}
func (*UnimplementedControllerServiceServer) EnterElection(req *PartitionElectionRequest, srv ControllerService_EnterElectionServer) error {
	return status.Errorf(codes.Unimplemented, "method EnterElection not implemented")
}

func RegisterControllerServiceServer(s *grpc.Server, srv ControllerServiceServer) {
	s.RegisterService(&_ControllerService_serviceDesc, srv)
}

func _ControllerService_CreatePartitionGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePartitionGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerServiceServer).CreatePartitionGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/atomix.controller.ControllerService/CreatePartitionGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerServiceServer).CreatePartitionGroup(ctx, req.(*CreatePartitionGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControllerService_DeletePartitionGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePartitionGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerServiceServer).DeletePartitionGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/atomix.controller.ControllerService/DeletePartitionGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerServiceServer).DeletePartitionGroup(ctx, req.(*DeletePartitionGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControllerService_GetPartitionGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPartitionGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerServiceServer).GetPartitionGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/atomix.controller.ControllerService/GetPartitionGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerServiceServer).GetPartitionGroups(ctx, req.(*GetPartitionGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControllerService_EnterElection_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PartitionElectionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ControllerServiceServer).EnterElection(m, &controllerServiceEnterElectionServer{stream})
}

type ControllerService_EnterElectionServer interface {
	Send(*PartitionElectionResponse) error
	grpc.ServerStream
}

type controllerServiceEnterElectionServer struct {
	grpc.ServerStream
}

func (x *controllerServiceEnterElectionServer) Send(m *PartitionElectionResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _ControllerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "atomix.controller.ControllerService",
	HandlerType: (*ControllerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePartitionGroup",
			Handler:    _ControllerService_CreatePartitionGroup_Handler,
		},
		{
			MethodName: "DeletePartitionGroup",
			Handler:    _ControllerService_DeletePartitionGroup_Handler,
		},
		{
			MethodName: "GetPartitionGroups",
			Handler:    _ControllerService_GetPartitionGroups_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "EnterElection",
			Handler:       _ControllerService_EnterElection_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "atomix/controller/controller.proto",
}
