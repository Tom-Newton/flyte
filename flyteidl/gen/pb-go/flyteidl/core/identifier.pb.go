// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/core/identifier.proto

package core

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Indicates a resource type within Flyte.
type ResourceType int32

const (
	ResourceType_UNSPECIFIED ResourceType = 0
	ResourceType_TASK        ResourceType = 1
	ResourceType_WORKFLOW    ResourceType = 2
	ResourceType_LAUNCH_PLAN ResourceType = 3
	// A dataset represents an entity modeled in Flyte DataCatalog. A Dataset is also a versioned entity and can be a compilation of multiple individual objects.
	// Eventually all Catalog objects should be modeled similar to Flyte Objects. The Dataset entities makes it possible for the UI  and CLI to act on the objects
	// in a similar manner to other Flyte objects
	ResourceType_DATASET ResourceType = 4
)

var ResourceType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "TASK",
	2: "WORKFLOW",
	3: "LAUNCH_PLAN",
	4: "DATASET",
}

var ResourceType_value = map[string]int32{
	"UNSPECIFIED": 0,
	"TASK":        1,
	"WORKFLOW":    2,
	"LAUNCH_PLAN": 3,
	"DATASET":     4,
}

func (x ResourceType) String() string {
	return proto.EnumName(ResourceType_name, int32(x))
}

func (ResourceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_adfa846a86e1fa0c, []int{0}
}

// Encapsulation of fields that uniquely identifies a Flyte resource.
type Identifier struct {
	// Identifies the specific type of resource that this identifier corresponds to.
	ResourceType ResourceType `protobuf:"varint,1,opt,name=resource_type,json=resourceType,proto3,enum=flyteidl.core.ResourceType" json:"resource_type,omitempty"`
	// Name of the project the resource belongs to.
	Project string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	// Name of the domain the resource belongs to.
	// A domain can be considered as a subset within a specific project.
	Domain string `protobuf:"bytes,3,opt,name=domain,proto3" json:"domain,omitempty"`
	// User provided value for the resource.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Specific version of the resource.
	Version string `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
	// Optional, org key applied to the resource.
	Org                  string   `protobuf:"bytes,6,opt,name=org,proto3" json:"org,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Identifier) Reset()         { *m = Identifier{} }
func (m *Identifier) String() string { return proto.CompactTextString(m) }
func (*Identifier) ProtoMessage()    {}
func (*Identifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_adfa846a86e1fa0c, []int{0}
}

func (m *Identifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Identifier.Unmarshal(m, b)
}
func (m *Identifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Identifier.Marshal(b, m, deterministic)
}
func (m *Identifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Identifier.Merge(m, src)
}
func (m *Identifier) XXX_Size() int {
	return xxx_messageInfo_Identifier.Size(m)
}
func (m *Identifier) XXX_DiscardUnknown() {
	xxx_messageInfo_Identifier.DiscardUnknown(m)
}

var xxx_messageInfo_Identifier proto.InternalMessageInfo

func (m *Identifier) GetResourceType() ResourceType {
	if m != nil {
		return m.ResourceType
	}
	return ResourceType_UNSPECIFIED
}

func (m *Identifier) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *Identifier) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *Identifier) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Identifier) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Identifier) GetOrg() string {
	if m != nil {
		return m.Org
	}
	return ""
}

// Encapsulation of fields that uniquely identifies a Flyte workflow execution
type WorkflowExecutionIdentifier struct {
	// Name of the project the resource belongs to.
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	// Name of the domain the resource belongs to.
	// A domain can be considered as a subset within a specific project.
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	// User or system provided value for the resource.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Optional, org key applied to the resource.
	Org                  string   `protobuf:"bytes,5,opt,name=org,proto3" json:"org,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkflowExecutionIdentifier) Reset()         { *m = WorkflowExecutionIdentifier{} }
func (m *WorkflowExecutionIdentifier) String() string { return proto.CompactTextString(m) }
func (*WorkflowExecutionIdentifier) ProtoMessage()    {}
func (*WorkflowExecutionIdentifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_adfa846a86e1fa0c, []int{1}
}

func (m *WorkflowExecutionIdentifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowExecutionIdentifier.Unmarshal(m, b)
}
func (m *WorkflowExecutionIdentifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowExecutionIdentifier.Marshal(b, m, deterministic)
}
func (m *WorkflowExecutionIdentifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowExecutionIdentifier.Merge(m, src)
}
func (m *WorkflowExecutionIdentifier) XXX_Size() int {
	return xxx_messageInfo_WorkflowExecutionIdentifier.Size(m)
}
func (m *WorkflowExecutionIdentifier) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowExecutionIdentifier.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowExecutionIdentifier proto.InternalMessageInfo

func (m *WorkflowExecutionIdentifier) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *WorkflowExecutionIdentifier) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *WorkflowExecutionIdentifier) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WorkflowExecutionIdentifier) GetOrg() string {
	if m != nil {
		return m.Org
	}
	return ""
}

// Encapsulation of fields that identify a Flyte node execution entity.
type NodeExecutionIdentifier struct {
	NodeId               string                       `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	ExecutionId          *WorkflowExecutionIdentifier `protobuf:"bytes,2,opt,name=execution_id,json=executionId,proto3" json:"execution_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *NodeExecutionIdentifier) Reset()         { *m = NodeExecutionIdentifier{} }
func (m *NodeExecutionIdentifier) String() string { return proto.CompactTextString(m) }
func (*NodeExecutionIdentifier) ProtoMessage()    {}
func (*NodeExecutionIdentifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_adfa846a86e1fa0c, []int{2}
}

func (m *NodeExecutionIdentifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeExecutionIdentifier.Unmarshal(m, b)
}
func (m *NodeExecutionIdentifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeExecutionIdentifier.Marshal(b, m, deterministic)
}
func (m *NodeExecutionIdentifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeExecutionIdentifier.Merge(m, src)
}
func (m *NodeExecutionIdentifier) XXX_Size() int {
	return xxx_messageInfo_NodeExecutionIdentifier.Size(m)
}
func (m *NodeExecutionIdentifier) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeExecutionIdentifier.DiscardUnknown(m)
}

var xxx_messageInfo_NodeExecutionIdentifier proto.InternalMessageInfo

func (m *NodeExecutionIdentifier) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *NodeExecutionIdentifier) GetExecutionId() *WorkflowExecutionIdentifier {
	if m != nil {
		return m.ExecutionId
	}
	return nil
}

// Encapsulation of fields that identify a Flyte task execution entity.
type TaskExecutionIdentifier struct {
	TaskId               *Identifier              `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	NodeExecutionId      *NodeExecutionIdentifier `protobuf:"bytes,2,opt,name=node_execution_id,json=nodeExecutionId,proto3" json:"node_execution_id,omitempty"`
	RetryAttempt         uint32                   `protobuf:"varint,3,opt,name=retry_attempt,json=retryAttempt,proto3" json:"retry_attempt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *TaskExecutionIdentifier) Reset()         { *m = TaskExecutionIdentifier{} }
func (m *TaskExecutionIdentifier) String() string { return proto.CompactTextString(m) }
func (*TaskExecutionIdentifier) ProtoMessage()    {}
func (*TaskExecutionIdentifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_adfa846a86e1fa0c, []int{3}
}

func (m *TaskExecutionIdentifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskExecutionIdentifier.Unmarshal(m, b)
}
func (m *TaskExecutionIdentifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskExecutionIdentifier.Marshal(b, m, deterministic)
}
func (m *TaskExecutionIdentifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskExecutionIdentifier.Merge(m, src)
}
func (m *TaskExecutionIdentifier) XXX_Size() int {
	return xxx_messageInfo_TaskExecutionIdentifier.Size(m)
}
func (m *TaskExecutionIdentifier) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskExecutionIdentifier.DiscardUnknown(m)
}

var xxx_messageInfo_TaskExecutionIdentifier proto.InternalMessageInfo

func (m *TaskExecutionIdentifier) GetTaskId() *Identifier {
	if m != nil {
		return m.TaskId
	}
	return nil
}

func (m *TaskExecutionIdentifier) GetNodeExecutionId() *NodeExecutionIdentifier {
	if m != nil {
		return m.NodeExecutionId
	}
	return nil
}

func (m *TaskExecutionIdentifier) GetRetryAttempt() uint32 {
	if m != nil {
		return m.RetryAttempt
	}
	return 0
}

// Encapsulation of fields the uniquely identify a signal.
type SignalIdentifier struct {
	// Unique identifier for a signal.
	SignalId string `protobuf:"bytes,1,opt,name=signal_id,json=signalId,proto3" json:"signal_id,omitempty"`
	// Identifies the Flyte workflow execution this signal belongs to.
	ExecutionId          *WorkflowExecutionIdentifier `protobuf:"bytes,2,opt,name=execution_id,json=executionId,proto3" json:"execution_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *SignalIdentifier) Reset()         { *m = SignalIdentifier{} }
func (m *SignalIdentifier) String() string { return proto.CompactTextString(m) }
func (*SignalIdentifier) ProtoMessage()    {}
func (*SignalIdentifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_adfa846a86e1fa0c, []int{4}
}

func (m *SignalIdentifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignalIdentifier.Unmarshal(m, b)
}
func (m *SignalIdentifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignalIdentifier.Marshal(b, m, deterministic)
}
func (m *SignalIdentifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignalIdentifier.Merge(m, src)
}
func (m *SignalIdentifier) XXX_Size() int {
	return xxx_messageInfo_SignalIdentifier.Size(m)
}
func (m *SignalIdentifier) XXX_DiscardUnknown() {
	xxx_messageInfo_SignalIdentifier.DiscardUnknown(m)
}

var xxx_messageInfo_SignalIdentifier proto.InternalMessageInfo

func (m *SignalIdentifier) GetSignalId() string {
	if m != nil {
		return m.SignalId
	}
	return ""
}

func (m *SignalIdentifier) GetExecutionId() *WorkflowExecutionIdentifier {
	if m != nil {
		return m.ExecutionId
	}
	return nil
}

func init() {
	proto.RegisterEnum("flyteidl.core.ResourceType", ResourceType_name, ResourceType_value)
	proto.RegisterType((*Identifier)(nil), "flyteidl.core.Identifier")
	proto.RegisterType((*WorkflowExecutionIdentifier)(nil), "flyteidl.core.WorkflowExecutionIdentifier")
	proto.RegisterType((*NodeExecutionIdentifier)(nil), "flyteidl.core.NodeExecutionIdentifier")
	proto.RegisterType((*TaskExecutionIdentifier)(nil), "flyteidl.core.TaskExecutionIdentifier")
	proto.RegisterType((*SignalIdentifier)(nil), "flyteidl.core.SignalIdentifier")
}

func init() { proto.RegisterFile("flyteidl/core/identifier.proto", fileDescriptor_adfa846a86e1fa0c) }

var fileDescriptor_adfa846a86e1fa0c = []byte{
	// 481 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xcf, 0x8f, 0xd2, 0x40,
	0x14, 0xb6, 0x0b, 0x5b, 0xd8, 0x57, 0x70, 0xeb, 0x1c, 0xa4, 0x86, 0xc4, 0x6c, 0x30, 0x31, 0x9b,
	0x4d, 0x6c, 0x13, 0xbc, 0x19, 0x0f, 0xd6, 0x5d, 0x36, 0x36, 0x8b, 0xec, 0xa6, 0x94, 0x90, 0x78,
	0x21, 0xa5, 0x7d, 0xd4, 0x11, 0xe8, 0x34, 0xd3, 0x41, 0xed, 0xc5, 0xc4, 0xff, 0xcc, 0x93, 0x7f,
	0x97, 0x99, 0x52, 0x96, 0x42, 0x58, 0xbd, 0x78, 0x7b, 0x3f, 0xbe, 0x79, 0xdf, 0xf7, 0xde, 0xbc,
	0x07, 0xcf, 0x67, 0x8b, 0x4c, 0x20, 0x0d, 0x17, 0x56, 0xc0, 0x38, 0x5a, 0x34, 0xc4, 0x58, 0xd0,
	0x19, 0x45, 0x6e, 0x26, 0x9c, 0x09, 0x46, 0x9a, 0x9b, 0xbc, 0x29, 0xf3, 0x9d, 0x5f, 0x0a, 0x80,
	0x73, 0x8f, 0x21, 0xef, 0xa0, 0xc9, 0x31, 0x65, 0x2b, 0x1e, 0xe0, 0x44, 0x64, 0x09, 0x1a, 0xca,
	0x99, 0x72, 0xfe, 0xb8, 0xdb, 0x36, 0x77, 0x5e, 0x99, 0x6e, 0x81, 0xf1, 0xb2, 0x04, 0xdd, 0x06,
	0x2f, 0x79, 0xc4, 0x80, 0x5a, 0xc2, 0xd9, 0x17, 0x0c, 0x84, 0x71, 0x74, 0xa6, 0x9c, 0x9f, 0xb8,
	0x1b, 0x97, 0x3c, 0x05, 0x35, 0x64, 0x4b, 0x9f, 0xc6, 0x46, 0x25, 0x4f, 0x14, 0x1e, 0x21, 0x50,
	0x8d, 0xfd, 0x25, 0x1a, 0xd5, 0x3c, 0x9a, 0xdb, 0xb2, 0xca, 0x57, 0xe4, 0x29, 0x65, 0xb1, 0x71,
	0xbc, 0xae, 0x52, 0xb8, 0x44, 0x87, 0x0a, 0xe3, 0x91, 0xa1, 0xe6, 0x51, 0x69, 0x76, 0x56, 0xd0,
	0x1e, 0x33, 0x3e, 0x9f, 0x2d, 0xd8, 0xb7, 0xde, 0x77, 0x0c, 0x56, 0x82, 0xb2, 0xb8, 0xd4, 0x52,
	0x49, 0x90, 0xf2, 0x90, 0xa0, 0xa3, 0x7f, 0x0a, 0x2a, 0x68, 0x8f, 0xb7, 0xb4, 0x3f, 0x15, 0x68,
	0x0d, 0x58, 0x88, 0x87, 0x38, 0x5b, 0x50, 0x8b, 0x59, 0x88, 0x13, 0x1a, 0x16, 0x9c, 0xaa, 0x74,
	0x9d, 0x90, 0x7c, 0x84, 0x06, 0x6e, 0xf0, 0x32, 0x2b, 0x89, 0xb5, 0xee, 0xc5, 0xde, 0x78, 0xff,
	0xd2, 0x8e, 0xab, 0xe1, 0x36, 0xd8, 0xf9, 0xad, 0x40, 0xcb, 0xf3, 0xd3, 0xf9, 0x21, 0x0d, 0x5d,
	0xa8, 0x09, 0x3f, 0x9d, 0x6f, 0x34, 0x68, 0xdd, 0x67, 0x7b, 0x2c, 0xa5, 0xa2, 0xaa, 0x44, 0x3a,
	0x21, 0x71, 0xe1, 0x49, 0xae, 0xfb, 0x80, 0xc6, 0x97, 0x7b, 0xaf, 0x1f, 0x68, 0xdd, 0x3d, 0x8d,
	0x77, 0x13, 0xe4, 0x85, 0x5c, 0x29, 0xc1, 0xb3, 0x89, 0x2f, 0x04, 0x2e, 0x13, 0x91, 0xff, 0x7e,
	0x53, 0x6e, 0x8d, 0xe0, 0x99, 0xbd, 0x8e, 0x75, 0x7e, 0x80, 0x3e, 0xa4, 0x51, 0xec, 0x2f, 0x4a,
	0x0d, 0xb4, 0xe1, 0x24, 0xcd, 0x63, 0xdb, 0x31, 0xd6, 0xd3, 0x02, 0xf4, 0x9f, 0x07, 0x79, 0x31,
	0x82, 0x46, 0x79, 0xa7, 0xc9, 0x29, 0x68, 0xa3, 0xc1, 0xf0, 0xae, 0x77, 0xe9, 0x5c, 0x3b, 0xbd,
	0x2b, 0xfd, 0x11, 0xa9, 0x43, 0xd5, 0xb3, 0x87, 0x37, 0xba, 0x42, 0x1a, 0x50, 0x1f, 0xdf, 0xba,
	0x37, 0xd7, 0xfd, 0xdb, 0xb1, 0x7e, 0x24, 0x81, 0x7d, 0x7b, 0x34, 0xb8, 0xfc, 0x30, 0xb9, 0xeb,
	0xdb, 0x03, 0xbd, 0x42, 0x34, 0xa8, 0x5d, 0xd9, 0x9e, 0x3d, 0xec, 0x79, 0x7a, 0xf5, 0xfd, 0xdb,
	0x4f, 0x6f, 0x22, 0x2a, 0x3e, 0xaf, 0xa6, 0x66, 0xc0, 0x96, 0x56, 0xae, 0x8d, 0xf1, 0x68, 0x6d,
	0x58, 0xf7, 0x87, 0x1a, 0x61, 0x6c, 0x25, 0xd3, 0x57, 0x11, 0xb3, 0x76, 0x6e, 0x77, 0xaa, 0xe6,
	0x17, 0xfb, 0xfa, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x03, 0xea, 0x9d, 0xb8, 0xd3, 0x03, 0x00,
	0x00,
}
