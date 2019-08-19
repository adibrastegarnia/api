// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: atomix/controller/partition.proto

package controller

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Partition identifier
type PartitionId struct {
	Partition            int32             `protobuf:"varint,1,opt,name=partition,proto3" json:"partition,omitempty"`
	Group                *PartitionGroupId `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PartitionId) Reset()         { *m = PartitionId{} }
func (m *PartitionId) String() string { return proto.CompactTextString(m) }
func (*PartitionId) ProtoMessage()    {}
func (*PartitionId) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fbf59d47cbc2727, []int{0}
}
func (m *PartitionId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PartitionId.Unmarshal(m, b)
}
func (m *PartitionId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PartitionId.Marshal(b, m, deterministic)
}
func (m *PartitionId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartitionId.Merge(m, src)
}
func (m *PartitionId) XXX_Size() int {
	return xxx_messageInfo_PartitionId.Size(m)
}
func (m *PartitionId) XXX_DiscardUnknown() {
	xxx_messageInfo_PartitionId.DiscardUnknown(m)
}

var xxx_messageInfo_PartitionId proto.InternalMessageInfo

func (m *PartitionId) GetPartition() int32 {
	if m != nil {
		return m.Partition
	}
	return 0
}

func (m *PartitionId) GetGroup() *PartitionGroupId {
	if m != nil {
		return m.Group
	}
	return nil
}

// Partition group name
type PartitionGroupId struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace            string   `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PartitionGroupId) Reset()         { *m = PartitionGroupId{} }
func (m *PartitionGroupId) String() string { return proto.CompactTextString(m) }
func (*PartitionGroupId) ProtoMessage()    {}
func (*PartitionGroupId) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fbf59d47cbc2727, []int{1}
}
func (m *PartitionGroupId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PartitionGroupId.Unmarshal(m, b)
}
func (m *PartitionGroupId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PartitionGroupId.Marshal(b, m, deterministic)
}
func (m *PartitionGroupId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartitionGroupId.Merge(m, src)
}
func (m *PartitionGroupId) XXX_Size() int {
	return xxx_messageInfo_PartitionGroupId.Size(m)
}
func (m *PartitionGroupId) XXX_DiscardUnknown() {
	xxx_messageInfo_PartitionGroupId.DiscardUnknown(m)
}

var xxx_messageInfo_PartitionGroupId proto.InternalMessageInfo

func (m *PartitionGroupId) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PartitionGroupId) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

// Partition configuration
type PartitionConfig struct {
	Partition            *PartitionId  `protobuf:"bytes,1,opt,name=partition,proto3" json:"partition,omitempty"`
	Controller           *NodeConfig   `protobuf:"bytes,2,opt,name=controller,proto3" json:"controller,omitempty"`
	Members              []*NodeConfig `protobuf:"bytes,3,rep,name=members,proto3" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *PartitionConfig) Reset()         { *m = PartitionConfig{} }
func (m *PartitionConfig) String() string { return proto.CompactTextString(m) }
func (*PartitionConfig) ProtoMessage()    {}
func (*PartitionConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fbf59d47cbc2727, []int{2}
}
func (m *PartitionConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PartitionConfig.Unmarshal(m, b)
}
func (m *PartitionConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PartitionConfig.Marshal(b, m, deterministic)
}
func (m *PartitionConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartitionConfig.Merge(m, src)
}
func (m *PartitionConfig) XXX_Size() int {
	return xxx_messageInfo_PartitionConfig.Size(m)
}
func (m *PartitionConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_PartitionConfig.DiscardUnknown(m)
}

var xxx_messageInfo_PartitionConfig proto.InternalMessageInfo

func (m *PartitionConfig) GetPartition() *PartitionId {
	if m != nil {
		return m.Partition
	}
	return nil
}

func (m *PartitionConfig) GetController() *NodeConfig {
	if m != nil {
		return m.Controller
	}
	return nil
}

func (m *PartitionConfig) GetMembers() []*NodeConfig {
	if m != nil {
		return m.Members
	}
	return nil
}

// Node configuration
type NodeConfig struct {
	ID                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Host                 string   `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	Port                 int32    `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeConfig) Reset()         { *m = NodeConfig{} }
func (m *NodeConfig) String() string { return proto.CompactTextString(m) }
func (*NodeConfig) ProtoMessage()    {}
func (*NodeConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fbf59d47cbc2727, []int{3}
}
func (m *NodeConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeConfig.Unmarshal(m, b)
}
func (m *NodeConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeConfig.Marshal(b, m, deterministic)
}
func (m *NodeConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeConfig.Merge(m, src)
}
func (m *NodeConfig) XXX_Size() int {
	return xxx_messageInfo_NodeConfig.Size(m)
}
func (m *NodeConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeConfig.DiscardUnknown(m)
}

var xxx_messageInfo_NodeConfig proto.InternalMessageInfo

func (m *NodeConfig) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *NodeConfig) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *NodeConfig) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

// Partition group info
type PartitionGroup struct {
	ID                   *PartitionGroupId   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Spec                 *PartitionGroupSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	Partitions           []*Partition        `protobuf:"bytes,3,rep,name=partitions,proto3" json:"partitions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *PartitionGroup) Reset()         { *m = PartitionGroup{} }
func (m *PartitionGroup) String() string { return proto.CompactTextString(m) }
func (*PartitionGroup) ProtoMessage()    {}
func (*PartitionGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fbf59d47cbc2727, []int{4}
}
func (m *PartitionGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PartitionGroup.Unmarshal(m, b)
}
func (m *PartitionGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PartitionGroup.Marshal(b, m, deterministic)
}
func (m *PartitionGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartitionGroup.Merge(m, src)
}
func (m *PartitionGroup) XXX_Size() int {
	return xxx_messageInfo_PartitionGroup.Size(m)
}
func (m *PartitionGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_PartitionGroup.DiscardUnknown(m)
}

var xxx_messageInfo_PartitionGroup proto.InternalMessageInfo

func (m *PartitionGroup) GetID() *PartitionGroupId {
	if m != nil {
		return m.ID
	}
	return nil
}

func (m *PartitionGroup) GetSpec() *PartitionGroupSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *PartitionGroup) GetPartitions() []*Partition {
	if m != nil {
		return m.Partitions
	}
	return nil
}

// Partition info
type Partition struct {
	PartitionID          int32                `protobuf:"varint,1,opt,name=partition_id,json=partitionId,proto3" json:"partition_id,omitempty"`
	Endpoints            []*PartitionEndpoint `protobuf:"bytes,2,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Partition) Reset()         { *m = Partition{} }
func (m *Partition) String() string { return proto.CompactTextString(m) }
func (*Partition) ProtoMessage()    {}
func (*Partition) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fbf59d47cbc2727, []int{5}
}
func (m *Partition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Partition.Unmarshal(m, b)
}
func (m *Partition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Partition.Marshal(b, m, deterministic)
}
func (m *Partition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Partition.Merge(m, src)
}
func (m *Partition) XXX_Size() int {
	return xxx_messageInfo_Partition.Size(m)
}
func (m *Partition) XXX_DiscardUnknown() {
	xxx_messageInfo_Partition.DiscardUnknown(m)
}

var xxx_messageInfo_Partition proto.InternalMessageInfo

func (m *Partition) GetPartitionID() int32 {
	if m != nil {
		return m.PartitionID
	}
	return 0
}

func (m *Partition) GetEndpoints() []*PartitionEndpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

// Partition endpoint
type PartitionEndpoint struct {
	Host                 string   `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port                 int32    `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PartitionEndpoint) Reset()         { *m = PartitionEndpoint{} }
func (m *PartitionEndpoint) String() string { return proto.CompactTextString(m) }
func (*PartitionEndpoint) ProtoMessage()    {}
func (*PartitionEndpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fbf59d47cbc2727, []int{6}
}
func (m *PartitionEndpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PartitionEndpoint.Unmarshal(m, b)
}
func (m *PartitionEndpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PartitionEndpoint.Marshal(b, m, deterministic)
}
func (m *PartitionEndpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartitionEndpoint.Merge(m, src)
}
func (m *PartitionEndpoint) XXX_Size() int {
	return xxx_messageInfo_PartitionEndpoint.Size(m)
}
func (m *PartitionEndpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_PartitionEndpoint.DiscardUnknown(m)
}

var xxx_messageInfo_PartitionEndpoint proto.InternalMessageInfo

func (m *PartitionEndpoint) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *PartitionEndpoint) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

// Partition group specification
type PartitionGroupSpec struct {
	Partitions           uint32     `protobuf:"varint,1,opt,name=partitions,proto3" json:"partitions,omitempty"`
	PartitionSize        uint32     `protobuf:"varint,2,opt,name=partition_size,json=partitionSize,proto3" json:"partition_size,omitempty"`
	Protocol             *types.Any `protobuf:"bytes,3,opt,name=protocol,proto3" json:"protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *PartitionGroupSpec) Reset()         { *m = PartitionGroupSpec{} }
func (m *PartitionGroupSpec) String() string { return proto.CompactTextString(m) }
func (*PartitionGroupSpec) ProtoMessage()    {}
func (*PartitionGroupSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fbf59d47cbc2727, []int{7}
}
func (m *PartitionGroupSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PartitionGroupSpec.Unmarshal(m, b)
}
func (m *PartitionGroupSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PartitionGroupSpec.Marshal(b, m, deterministic)
}
func (m *PartitionGroupSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartitionGroupSpec.Merge(m, src)
}
func (m *PartitionGroupSpec) XXX_Size() int {
	return xxx_messageInfo_PartitionGroupSpec.Size(m)
}
func (m *PartitionGroupSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_PartitionGroupSpec.DiscardUnknown(m)
}

var xxx_messageInfo_PartitionGroupSpec proto.InternalMessageInfo

func (m *PartitionGroupSpec) GetPartitions() uint32 {
	if m != nil {
		return m.Partitions
	}
	return 0
}

func (m *PartitionGroupSpec) GetPartitionSize() uint32 {
	if m != nil {
		return m.PartitionSize
	}
	return 0
}

func (m *PartitionGroupSpec) GetProtocol() *types.Any {
	if m != nil {
		return m.Protocol
	}
	return nil
}

func init() {
	proto.RegisterType((*PartitionId)(nil), "atomix.controller.PartitionId")
	proto.RegisterType((*PartitionGroupId)(nil), "atomix.controller.PartitionGroupId")
	proto.RegisterType((*PartitionConfig)(nil), "atomix.controller.PartitionConfig")
	proto.RegisterType((*NodeConfig)(nil), "atomix.controller.NodeConfig")
	proto.RegisterType((*PartitionGroup)(nil), "atomix.controller.PartitionGroup")
	proto.RegisterType((*Partition)(nil), "atomix.controller.Partition")
	proto.RegisterType((*PartitionEndpoint)(nil), "atomix.controller.PartitionEndpoint")
	proto.RegisterType((*PartitionGroupSpec)(nil), "atomix.controller.PartitionGroupSpec")
}

func init() { proto.RegisterFile("atomix/controller/partition.proto", fileDescriptor_6fbf59d47cbc2727) }

var fileDescriptor_6fbf59d47cbc2727 = []byte{
	// 477 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0x9d, 0xa6, 0xe0, 0x31, 0x69, 0xe9, 0xaa, 0x42, 0xa6, 0x2a, 0x6d, 0x30, 0x54, 0xea,
	0xc9, 0x46, 0xe1, 0x80, 0xaa, 0x96, 0x03, 0x21, 0x08, 0x45, 0x42, 0x08, 0x6d, 0x3f, 0x00, 0x39,
	0xf6, 0xc6, 0xac, 0x94, 0xec, 0xac, 0xec, 0xad, 0x44, 0x7b, 0xe5, 0xcc, 0x6f, 0xc1, 0x5f, 0xf4,
	0xd0, 0x2f, 0x41, 0xde, 0x75, 0xd6, 0x4e, 0x53, 0x25, 0x9c, 0x76, 0x34, 0xf3, 0x66, 0xde, 0xcc,
	0x9b, 0x59, 0x78, 0x99, 0x28, 0x9c, 0xf3, 0x9f, 0x71, 0x8a, 0x42, 0x15, 0x38, 0x9b, 0xb1, 0x22,
	0x96, 0x49, 0xa1, 0xb8, 0xe2, 0x28, 0x22, 0x59, 0xa0, 0x42, 0xb2, 0x67, 0x20, 0x51, 0x03, 0x39,
	0x78, 0x9e, 0x23, 0xe6, 0x33, 0x16, 0x6b, 0xc0, 0xe4, 0x6a, 0x1a, 0x27, 0xe2, 0xda, 0xa0, 0x0f,
	0xf6, 0x73, 0xcc, 0x51, 0x9b, 0x71, 0x65, 0x19, 0x6f, 0x38, 0x05, 0xff, 0xdb, 0xa2, 0xec, 0x38,
	0x23, 0x87, 0xe0, 0x59, 0x96, 0xc0, 0xe9, 0x3b, 0xa7, 0x5d, 0xda, 0x38, 0xc8, 0x19, 0x74, 0xf3,
	0x02, 0xaf, 0x64, 0xe0, 0xf6, 0x9d, 0x53, 0x7f, 0xf0, 0x2a, 0x5a, 0x69, 0x20, 0xb2, 0xc5, 0x3e,
	0x57, 0xc0, 0x71, 0x46, 0x4d, 0x46, 0x38, 0x82, 0xa7, 0xf7, 0x43, 0x84, 0xc0, 0x96, 0x48, 0xe6,
	0x4c, 0xf3, 0x78, 0x54, 0xdb, 0x55, 0x03, 0xd5, 0x5b, 0xca, 0x24, 0x65, 0x9a, 0xc6, 0xa3, 0x8d,
	0x23, 0xfc, 0xeb, 0xc0, 0xae, 0x2d, 0xf3, 0x11, 0xc5, 0x94, 0xe7, 0xe4, 0xe2, 0x7e, 0xcb, 0xfe,
	0xe0, 0x68, 0x5d, 0x63, 0xe3, 0xac, 0x3d, 0xd2, 0x7b, 0x80, 0x06, 0x54, 0xcf, 0xf5, 0xe2, 0x81,
	0xf4, 0xaf, 0x98, 0x31, 0x43, 0x48, 0x5b, 0x09, 0xe4, 0x1d, 0x3c, 0x9a, 0xb3, 0xf9, 0x84, 0x15,
	0x65, 0xd0, 0xe9, 0x77, 0x36, 0xe7, 0x2e, 0xd0, 0xe1, 0x17, 0x80, 0xc6, 0x4d, 0x9e, 0x81, 0xcb,
	0x33, 0xa3, 0xc3, 0x70, 0xfb, 0xee, 0xf6, 0xd8, 0x1d, 0x8f, 0xa8, 0xcb, 0xb5, 0x42, 0x3f, 0xb0,
	0x54, 0xb5, 0x10, 0xda, 0xae, 0x7c, 0x12, 0x0b, 0x15, 0x74, 0xf4, 0x76, 0xb4, 0x1d, 0xfe, 0x71,
	0x60, 0x67, 0x59, 0x5e, 0x72, 0x6e, 0x4b, 0xfe, 0xdf, 0xa2, 0x96, 0x78, 0xcf, 0x60, 0xab, 0x94,
	0x2c, 0xad, 0xf5, 0x38, 0xd9, 0x98, 0x7e, 0x29, 0x59, 0x4a, 0x75, 0x0a, 0xb9, 0x00, 0xb0, 0xea,
	0x2e, 0x44, 0x39, 0x5c, 0x57, 0x80, 0xb6, 0xf0, 0xe1, 0x2f, 0x07, 0x3c, 0x1b, 0x21, 0x03, 0x78,
	0x62, 0x63, 0xdf, 0xeb, 0x69, 0xba, 0xc3, 0xdd, 0xbb, 0xdb, 0xe3, 0xd6, 0xd1, 0x8e, 0xa8, 0x2f,
	0x5b, 0x17, 0x3c, 0x04, 0x8f, 0x89, 0x4c, 0x22, 0x17, 0xaa, 0x0c, 0x5c, 0x4d, 0xff, 0x7a, 0x1d,
	0xfd, 0xa7, 0x1a, 0x4c, 0x9b, 0xb4, 0xf0, 0x1c, 0xf6, 0x56, 0xe2, 0x76, 0x17, 0xce, 0x03, 0xbb,
	0x70, 0x5b, 0xbb, 0xf8, 0xed, 0x00, 0x59, 0x55, 0x87, 0x1c, 0x2d, 0xe9, 0x52, 0x15, 0xe9, 0xb5,
	0x27, 0x27, 0x27, 0xb0, 0xd3, 0xcc, 0x5a, 0xf2, 0x1b, 0x73, 0xfd, 0x3d, 0xda, 0xb3, 0xde, 0x4b,
	0x7e, 0xc3, 0xc8, 0x1b, 0x78, 0xac, 0x3f, 0x6e, 0x8a, 0x33, 0x7d, 0x01, 0xfe, 0x60, 0x3f, 0x32,
	0x7f, 0x3e, 0x5a, 0xfc, 0xf9, 0xe8, 0x83, 0xb8, 0xa6, 0x16, 0x35, 0xd9, 0xd6, 0xd6, 0xdb, 0x7f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x25, 0x9d, 0xdc, 0xe2, 0x51, 0x04, 0x00, 0x00,
}
