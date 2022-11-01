// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: v1/cluster.proto

package aspenv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint32     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Address   string     `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	State     uint32     `protobuf:"varint,3,opt,name=state,proto3" json:"state,omitempty"`
	Heartbeat *Heartbeat `protobuf:"bytes,4,opt,name=heartbeat,proto3" json:"heartbeat,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_cluster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_v1_cluster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_v1_cluster_proto_rawDescGZIP(), []int{0}
}

func (x *Node) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Node) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Node) GetState() uint32 {
	if x != nil {
		return x.State
	}
	return 0
}

func (x *Node) GetHeartbeat() *Heartbeat {
	if x != nil {
		return x.Heartbeat
	}
	return nil
}

type Heartbeat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Generation uint32 `protobuf:"varint,1,opt,name=generation,proto3" json:"generation,omitempty"`
	Version    uint32 `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Heartbeat) Reset() {
	*x = Heartbeat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_cluster_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Heartbeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Heartbeat) ProtoMessage() {}

func (x *Heartbeat) ProtoReflect() protoreflect.Message {
	mi := &file_v1_cluster_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Heartbeat.ProtoReflect.Descriptor instead.
func (*Heartbeat) Descriptor() ([]byte, []int) {
	return file_v1_cluster_proto_rawDescGZIP(), []int{1}
}

func (x *Heartbeat) GetGeneration() uint32 {
	if x != nil {
		return x.Generation
	}
	return 0
}

func (x *Heartbeat) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

type NodeDigest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint32     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Heartbeat *Heartbeat `protobuf:"bytes,2,opt,name=heartbeat,proto3" json:"heartbeat,omitempty"`
}

func (x *NodeDigest) Reset() {
	*x = NodeDigest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_cluster_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeDigest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeDigest) ProtoMessage() {}

func (x *NodeDigest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_cluster_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeDigest.ProtoReflect.Descriptor instead.
func (*NodeDigest) Descriptor() ([]byte, []int) {
	return file_v1_cluster_proto_rawDescGZIP(), []int{2}
}

func (x *NodeDigest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NodeDigest) GetHeartbeat() *Heartbeat {
	if x != nil {
		return x.Heartbeat
	}
	return nil
}

type ClusterGossip struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Digests map[uint32]*NodeDigest `protobuf:"bytes,1,rep,name=digests,proto3" json:"digests,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Nodes   map[uint32]*Node       `protobuf:"bytes,2,rep,name=nodes,proto3" json:"nodes,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ClusterGossip) Reset() {
	*x = ClusterGossip{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_cluster_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterGossip) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterGossip) ProtoMessage() {}

func (x *ClusterGossip) ProtoReflect() protoreflect.Message {
	mi := &file_v1_cluster_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterGossip.ProtoReflect.Descriptor instead.
func (*ClusterGossip) Descriptor() ([]byte, []int) {
	return file_v1_cluster_proto_rawDescGZIP(), []int{3}
}

func (x *ClusterGossip) GetDigests() map[uint32]*NodeDigest {
	if x != nil {
		return x.Digests
	}
	return nil
}

func (x *ClusterGossip) GetNodes() map[uint32]*Node {
	if x != nil {
		return x.Nodes
	}
	return nil
}

type ClusterPledge struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterKey string `protobuf:"bytes,1,opt,name=cluster_key,json=clusterKey,proto3" json:"cluster_key,omitempty"`
	NodeId     uint32 `protobuf:"varint,2,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
}

func (x *ClusterPledge) Reset() {
	*x = ClusterPledge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_cluster_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterPledge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterPledge) ProtoMessage() {}

func (x *ClusterPledge) ProtoReflect() protoreflect.Message {
	mi := &file_v1_cluster_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterPledge.ProtoReflect.Descriptor instead.
func (*ClusterPledge) Descriptor() ([]byte, []int) {
	return file_v1_cluster_proto_rawDescGZIP(), []int{4}
}

func (x *ClusterPledge) GetClusterKey() string {
	if x != nil {
		return x.ClusterKey
	}
	return ""
}

func (x *ClusterPledge) GetNodeId() uint32 {
	if x != nil {
		return x.NodeId
	}
	return 0
}

var File_v1_cluster_proto protoreflect.FileDescriptor

var file_v1_cluster_proto_rawDesc = []byte{
	0x0a, 0x10, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x61, 0x73, 0x70, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x22, 0x79, 0x0a, 0x04,
	0x4e, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x31, 0x0a, 0x09, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x73, 0x70, 0x65, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x09, 0x68, 0x65,
	0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x22, 0x45, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74,
	0x62, 0x65, 0x61, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x4f,
	0x0a, 0x0a, 0x4e, 0x6f, 0x64, 0x65, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x31, 0x0a, 0x09,
	0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x61, 0x73, 0x70, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74,
	0x62, 0x65, 0x61, 0x74, 0x52, 0x09, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x22,
	0xa5, 0x02, 0x0a, 0x0d, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x47, 0x6f, 0x73, 0x73, 0x69,
	0x70, 0x12, 0x3e, 0x0a, 0x07, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x61, 0x73, 0x70, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x47, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x2e, 0x44, 0x69, 0x67, 0x65,
	0x73, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74,
	0x73, 0x12, 0x38, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x61, 0x73, 0x70, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x47, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x1a, 0x50, 0x0a, 0x0c, 0x44,
	0x69, 0x67, 0x65, 0x73, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61,
	0x73, 0x70, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x44, 0x69, 0x67, 0x65,
	0x73, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x48, 0x0a,
	0x0a, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x24, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61,
	0x73, 0x70, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x49, 0x0a, 0x0d, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x50, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65,
	0x49, 0x64, 0x32, 0x50, 0x0a, 0x14, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x47, 0x6f, 0x73,
	0x73, 0x69, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x04, 0x45, 0x78,
	0x65, 0x63, 0x12, 0x17, 0x2e, 0x61, 0x73, 0x70, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x47, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x1a, 0x17, 0x2e, 0x61, 0x73,
	0x70, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x47, 0x6f,
	0x73, 0x73, 0x69, 0x70, 0x32, 0x49, 0x0a, 0x0d, 0x50, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x04, 0x45, 0x78, 0x65, 0x63, 0x12, 0x17, 0x2e,
	0x61, 0x73, 0x70, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x50, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x1a, 0x17, 0x2e, 0x61, 0x73, 0x70, 0x65, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x42,
	0x94, 0x01, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x73, 0x70, 0x65, 0x6e, 0x2e, 0x76, 0x31,
	0x42, 0x0c, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x79, 0x6e,
	0x6e, 0x61, 0x78, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x61, 0x73, 0x70, 0x65, 0x6e, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x3b,
	0x61, 0x73, 0x70, 0x65, 0x6e, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x58, 0x58, 0xaa, 0x02, 0x08,
	0x41, 0x73, 0x70, 0x65, 0x6e, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x08, 0x41, 0x73, 0x70, 0x65, 0x6e,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x14, 0x41, 0x73, 0x70, 0x65, 0x6e, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x41, 0x73, 0x70,
	0x65, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_cluster_proto_rawDescOnce sync.Once
	file_v1_cluster_proto_rawDescData = file_v1_cluster_proto_rawDesc
)

func file_v1_cluster_proto_rawDescGZIP() []byte {
	file_v1_cluster_proto_rawDescOnce.Do(func() {
		file_v1_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_cluster_proto_rawDescData)
	})
	return file_v1_cluster_proto_rawDescData
}

var file_v1_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_v1_cluster_proto_goTypes = []interface{}{
	(*Node)(nil),          // 0: aspen.v1.Node
	(*Heartbeat)(nil),     // 1: aspen.v1.Heartbeat
	(*NodeDigest)(nil),    // 2: aspen.v1.NodeDigest
	(*ClusterGossip)(nil), // 3: aspen.v1.ClusterGossip
	(*ClusterPledge)(nil), // 4: aspen.v1.ClusterPledge
	nil,                   // 5: aspen.v1.ClusterGossip.DigestsEntry
	nil,                   // 6: aspen.v1.ClusterGossip.NodesEntry
}
var file_v1_cluster_proto_depIdxs = []int32{
	1, // 0: aspen.v1.Node.heartbeat:type_name -> aspen.v1.Heartbeat
	1, // 1: aspen.v1.NodeDigest.heartbeat:type_name -> aspen.v1.Heartbeat
	5, // 2: aspen.v1.ClusterGossip.digests:type_name -> aspen.v1.ClusterGossip.DigestsEntry
	6, // 3: aspen.v1.ClusterGossip.nodes:type_name -> aspen.v1.ClusterGossip.NodesEntry
	2, // 4: aspen.v1.ClusterGossip.DigestsEntry.value:type_name -> aspen.v1.NodeDigest
	0, // 5: aspen.v1.ClusterGossip.NodesEntry.value:type_name -> aspen.v1.Node
	3, // 6: aspen.v1.ClusterGossipService.Exec:input_type -> aspen.v1.ClusterGossip
	4, // 7: aspen.v1.PledgeService.Exec:input_type -> aspen.v1.ClusterPledge
	3, // 8: aspen.v1.ClusterGossipService.Exec:output_type -> aspen.v1.ClusterGossip
	4, // 9: aspen.v1.PledgeService.Exec:output_type -> aspen.v1.ClusterPledge
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_v1_cluster_proto_init() }
func file_v1_cluster_proto_init() {
	if File_v1_cluster_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_cluster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_cluster_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Heartbeat); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_cluster_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeDigest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_cluster_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterGossip); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_cluster_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterPledge); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_cluster_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_v1_cluster_proto_goTypes,
		DependencyIndexes: file_v1_cluster_proto_depIdxs,
		MessageInfos:      file_v1_cluster_proto_msgTypes,
	}.Build()
	File_v1_cluster_proto = out.File
	file_v1_cluster_proto_rawDesc = nil
	file_v1_cluster_proto_goTypes = nil
	file_v1_cluster_proto_depIdxs = nil
}
