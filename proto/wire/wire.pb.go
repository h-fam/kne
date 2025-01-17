// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: proto/wire.proto

package wire

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

type Pod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	SrcIp  string  `protobuf:"bytes,2,opt,name=src_ip,json=srcIp,proto3" json:"src_ip,omitempty"`
	NetNs  string  `protobuf:"bytes,3,opt,name=net_ns,json=netNs,proto3" json:"net_ns,omitempty"`
	KubeNs string  `protobuf:"bytes,4,opt,name=kube_ns,json=kubeNs,proto3" json:"kube_ns,omitempty"`
	Links  []*Link `protobuf:"bytes,5,rep,name=links,proto3" json:"links,omitempty"`
}

func (x *Pod) Reset() {
	*x = Pod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_wire_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pod) ProtoMessage() {}

func (x *Pod) ProtoReflect() protoreflect.Message {
	mi := &file_proto_wire_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pod.ProtoReflect.Descriptor instead.
func (*Pod) Descriptor() ([]byte, []int) {
	return file_proto_wire_proto_rawDescGZIP(), []int{0}
}

func (x *Pod) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Pod) GetSrcIp() string {
	if x != nil {
		return x.SrcIp
	}
	return ""
}

func (x *Pod) GetNetNs() string {
	if x != nil {
		return x.NetNs
	}
	return ""
}

func (x *Pod) GetKubeNs() string {
	if x != nil {
		return x.KubeNs
	}
	return ""
}

func (x *Pod) GetLinks() []*Link {
	if x != nil {
		return x.Links
	}
	return nil
}

type Link struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PeerPod   string `protobuf:"bytes,1,opt,name=peer_pod,json=peerPod,proto3" json:"peer_pod,omitempty"`
	LocalIntf string `protobuf:"bytes,2,opt,name=local_intf,json=localIntf,proto3" json:"local_intf,omitempty"`
	PeerIntf  string `protobuf:"bytes,3,opt,name=peer_intf,json=peerIntf,proto3" json:"peer_intf,omitempty"`
	LocalIp   string `protobuf:"bytes,4,opt,name=local_ip,json=localIp,proto3" json:"local_ip,omitempty"`
	PeerIp    string `protobuf:"bytes,5,opt,name=peer_ip,json=peerIp,proto3" json:"peer_ip,omitempty"`
	Uid       int64  `protobuf:"varint,6,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *Link) Reset() {
	*x = Link{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_wire_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Link) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Link) ProtoMessage() {}

func (x *Link) ProtoReflect() protoreflect.Message {
	mi := &file_proto_wire_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Link.ProtoReflect.Descriptor instead.
func (*Link) Descriptor() ([]byte, []int) {
	return file_proto_wire_proto_rawDescGZIP(), []int{1}
}

func (x *Link) GetPeerPod() string {
	if x != nil {
		return x.PeerPod
	}
	return ""
}

func (x *Link) GetLocalIntf() string {
	if x != nil {
		return x.LocalIntf
	}
	return ""
}

func (x *Link) GetPeerIntf() string {
	if x != nil {
		return x.PeerIntf
	}
	return ""
}

func (x *Link) GetLocalIp() string {
	if x != nil {
		return x.LocalIp
	}
	return ""
}

func (x *Link) GetPeerIp() string {
	if x != nil {
		return x.PeerIp
	}
	return ""
}

func (x *Link) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type PodQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	KubeNs string `protobuf:"bytes,2,opt,name=kube_ns,json=kubeNs,proto3" json:"kube_ns,omitempty"`
}

func (x *PodQuery) Reset() {
	*x = PodQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_wire_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PodQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PodQuery) ProtoMessage() {}

func (x *PodQuery) ProtoReflect() protoreflect.Message {
	mi := &file_proto_wire_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PodQuery.ProtoReflect.Descriptor instead.
func (*PodQuery) Descriptor() ([]byte, []int) {
	return file_proto_wire_proto_rawDescGZIP(), []int{2}
}

func (x *PodQuery) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PodQuery) GetKubeNs() string {
	if x != nil {
		return x.KubeNs
	}
	return ""
}

type SkipQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pod    string `protobuf:"bytes,1,opt,name=pod,proto3" json:"pod,omitempty"`
	Peer   string `protobuf:"bytes,2,opt,name=peer,proto3" json:"peer,omitempty"`
	KubeNs string `protobuf:"bytes,3,opt,name=kube_ns,json=kubeNs,proto3" json:"kube_ns,omitempty"`
}

func (x *SkipQuery) Reset() {
	*x = SkipQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_wire_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SkipQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SkipQuery) ProtoMessage() {}

func (x *SkipQuery) ProtoReflect() protoreflect.Message {
	mi := &file_proto_wire_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SkipQuery.ProtoReflect.Descriptor instead.
func (*SkipQuery) Descriptor() ([]byte, []int) {
	return file_proto_wire_proto_rawDescGZIP(), []int{3}
}

func (x *SkipQuery) GetPod() string {
	if x != nil {
		return x.Pod
	}
	return ""
}

func (x *SkipQuery) GetPeer() string {
	if x != nil {
		return x.Peer
	}
	return ""
}

func (x *SkipQuery) GetKubeNs() string {
	if x != nil {
		return x.KubeNs
	}
	return ""
}

type BoolResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response bool `protobuf:"varint,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *BoolResponse) Reset() {
	*x = BoolResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_wire_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoolResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoolResponse) ProtoMessage() {}

func (x *BoolResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_wire_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoolResponse.ProtoReflect.Descriptor instead.
func (*BoolResponse) Descriptor() ([]byte, []int) {
	return file_proto_wire_proto_rawDescGZIP(), []int{4}
}

func (x *BoolResponse) GetResponse() bool {
	if x != nil {
		return x.Response
	}
	return false
}

type RemotePod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NetNs    string `protobuf:"bytes,1,opt,name=net_ns,json=netNs,proto3" json:"net_ns,omitempty"`
	IntfName string `protobuf:"bytes,2,opt,name=intf_name,json=intfName,proto3" json:"intf_name,omitempty"`
	IntfIp   string `protobuf:"bytes,3,opt,name=intf_ip,json=intfIp,proto3" json:"intf_ip,omitempty"`
	PeerVtep string `protobuf:"bytes,4,opt,name=peer_vtep,json=peerVtep,proto3" json:"peer_vtep,omitempty"`
	KubeNs   string `protobuf:"bytes,5,opt,name=kube_ns,json=kubeNs,proto3" json:"kube_ns,omitempty"`
	Vni      int64  `protobuf:"varint,6,opt,name=vni,proto3" json:"vni,omitempty"`
}

func (x *RemotePod) Reset() {
	*x = RemotePod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_wire_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemotePod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemotePod) ProtoMessage() {}

func (x *RemotePod) ProtoReflect() protoreflect.Message {
	mi := &file_proto_wire_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemotePod.ProtoReflect.Descriptor instead.
func (*RemotePod) Descriptor() ([]byte, []int) {
	return file_proto_wire_proto_rawDescGZIP(), []int{5}
}

func (x *RemotePod) GetNetNs() string {
	if x != nil {
		return x.NetNs
	}
	return ""
}

func (x *RemotePod) GetIntfName() string {
	if x != nil {
		return x.IntfName
	}
	return ""
}

func (x *RemotePod) GetIntfIp() string {
	if x != nil {
		return x.IntfIp
	}
	return ""
}

func (x *RemotePod) GetPeerVtep() string {
	if x != nil {
		return x.PeerVtep
	}
	return ""
}

func (x *RemotePod) GetKubeNs() string {
	if x != nil {
		return x.KubeNs
	}
	return ""
}

func (x *RemotePod) GetVni() int64 {
	if x != nil {
		return x.Vni
	}
	return 0
}

var File_proto_wire_proto protoreflect.FileDescriptor

var file_proto_wire_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x69, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x6b, 0x6e, 0x65, 0x2e, 0x77, 0x69, 0x72, 0x65, 0x22, 0x86, 0x01, 0x0a,
	0x03, 0x50, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x72, 0x63, 0x5f,
	0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x72, 0x63, 0x49, 0x70, 0x12,
	0x15, 0x0a, 0x06, 0x6e, 0x65, 0x74, 0x5f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6e, 0x65, 0x74, 0x4e, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x6b, 0x75, 0x62, 0x65, 0x5f, 0x6e,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6b, 0x75, 0x62, 0x65, 0x4e, 0x73, 0x12,
	0x24, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x6b, 0x6e, 0x65, 0x2e, 0x77, 0x69, 0x72, 0x65, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x05,
	0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x22, 0xa3, 0x01, 0x0a, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x19,
	0x0a, 0x08, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x70, 0x65, 0x65, 0x72, 0x50, 0x6f, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x74, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x49, 0x6e, 0x74, 0x66, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x65, 0x65, 0x72,
	0x5f, 0x69, 0x6e, 0x74, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x65, 0x65,
	0x72, 0x49, 0x6e, 0x74, 0x66, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x69,
	0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x49, 0x70,
	0x12, 0x17, 0x0a, 0x07, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x69, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x37, 0x0a, 0x08, 0x50,
	0x6f, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6b,
	0x75, 0x62, 0x65, 0x5f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6b, 0x75,
	0x62, 0x65, 0x4e, 0x73, 0x22, 0x4a, 0x0a, 0x09, 0x53, 0x6b, 0x69, 0x70, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x70, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x65, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x65, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x6b, 0x75, 0x62, 0x65, 0x5f,
	0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6b, 0x75, 0x62, 0x65, 0x4e, 0x73,
	0x22, 0x2a, 0x0a, 0x0c, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xa0, 0x01, 0x0a,
	0x09, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x6f, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x6e, 0x65,
	0x74, 0x5f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x65, 0x74, 0x4e,
	0x73, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x66, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x66, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x69, 0x6e, 0x74, 0x66, 0x5f, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x69, 0x6e, 0x74, 0x66, 0x49, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x65, 0x65, 0x72, 0x5f,
	0x76, 0x74, 0x65, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x65, 0x65, 0x72,
	0x56, 0x74, 0x65, 0x70, 0x12, 0x17, 0x0a, 0x07, 0x6b, 0x75, 0x62, 0x65, 0x5f, 0x6e, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6b, 0x75, 0x62, 0x65, 0x4e, 0x73, 0x12, 0x10, 0x0a,
	0x03, 0x76, 0x6e, 0x69, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x76, 0x6e, 0x69, 0x32,
	0x8f, 0x02, 0x0a, 0x05, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x12, 0x28, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x12, 0x2e, 0x6b, 0x6e, 0x65, 0x2e, 0x77, 0x69, 0x72, 0x65, 0x2e, 0x50, 0x6f, 0x64, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x1a, 0x0d, 0x2e, 0x6b, 0x6e, 0x65, 0x2e, 0x77, 0x69, 0x72, 0x65, 0x2e,
	0x50, 0x6f, 0x64, 0x12, 0x31, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x12,
	0x0d, 0x2e, 0x6b, 0x6e, 0x65, 0x2e, 0x77, 0x69, 0x72, 0x65, 0x2e, 0x50, 0x6f, 0x64, 0x1a, 0x16,
	0x2e, 0x6b, 0x6e, 0x65, 0x2e, 0x77, 0x69, 0x72, 0x65, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0b, 0x53, 0x6b, 0x69, 0x70, 0x52, 0x65,
	0x76, 0x65, 0x72, 0x73, 0x65, 0x12, 0x13, 0x2e, 0x6b, 0x6e, 0x65, 0x2e, 0x77, 0x69, 0x72, 0x65,
	0x2e, 0x53, 0x6b, 0x69, 0x70, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x16, 0x2e, 0x6b, 0x6e, 0x65,
	0x2e, 0x77, 0x69, 0x72, 0x65, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x53, 0x6b, 0x69, 0x70, 0x12, 0x13, 0x2e, 0x6b, 0x6e, 0x65,
	0x2e, 0x77, 0x69, 0x72, 0x65, 0x2e, 0x53, 0x6b, 0x69, 0x70, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a,
	0x16, 0x2e, 0x6b, 0x6e, 0x65, 0x2e, 0x77, 0x69, 0x72, 0x65, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x49, 0x73, 0x53, 0x6b, 0x69,
	0x70, 0x70, 0x65, 0x64, 0x12, 0x13, 0x2e, 0x6b, 0x6e, 0x65, 0x2e, 0x77, 0x69, 0x72, 0x65, 0x2e,
	0x53, 0x6b, 0x69, 0x70, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x16, 0x2e, 0x6b, 0x6e, 0x65, 0x2e,
	0x77, 0x69, 0x72, 0x65, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x32, 0x3f, 0x0a, 0x06, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x6b, 0x6e, 0x65, 0x2e, 0x77, 0x69, 0x72, 0x65,
	0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x6f, 0x64, 0x1a, 0x16, 0x2e, 0x6b, 0x6e, 0x65,
	0x2e, 0x77, 0x69, 0x72, 0x65, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6b, 0x6e, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x77, 0x69, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_wire_proto_rawDescOnce sync.Once
	file_proto_wire_proto_rawDescData = file_proto_wire_proto_rawDesc
)

func file_proto_wire_proto_rawDescGZIP() []byte {
	file_proto_wire_proto_rawDescOnce.Do(func() {
		file_proto_wire_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_wire_proto_rawDescData)
	})
	return file_proto_wire_proto_rawDescData
}

var file_proto_wire_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_wire_proto_goTypes = []interface{}{
	(*Pod)(nil),          // 0: kne.wire.Pod
	(*Link)(nil),         // 1: kne.wire.Link
	(*PodQuery)(nil),     // 2: kne.wire.PodQuery
	(*SkipQuery)(nil),    // 3: kne.wire.SkipQuery
	(*BoolResponse)(nil), // 4: kne.wire.BoolResponse
	(*RemotePod)(nil),    // 5: kne.wire.RemotePod
}
var file_proto_wire_proto_depIdxs = []int32{
	1, // 0: kne.wire.Pod.links:type_name -> kne.wire.Link
	2, // 1: kne.wire.Local.Get:input_type -> kne.wire.PodQuery
	0, // 2: kne.wire.Local.SetAlive:input_type -> kne.wire.Pod
	3, // 3: kne.wire.Local.SkipReverse:input_type -> kne.wire.SkipQuery
	3, // 4: kne.wire.Local.Skip:input_type -> kne.wire.SkipQuery
	3, // 5: kne.wire.Local.IsSkipped:input_type -> kne.wire.SkipQuery
	5, // 6: kne.wire.Remote.Update:input_type -> kne.wire.RemotePod
	0, // 7: kne.wire.Local.Get:output_type -> kne.wire.Pod
	4, // 8: kne.wire.Local.SetAlive:output_type -> kne.wire.BoolResponse
	4, // 9: kne.wire.Local.SkipReverse:output_type -> kne.wire.BoolResponse
	4, // 10: kne.wire.Local.Skip:output_type -> kne.wire.BoolResponse
	4, // 11: kne.wire.Local.IsSkipped:output_type -> kne.wire.BoolResponse
	4, // 12: kne.wire.Remote.Update:output_type -> kne.wire.BoolResponse
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_wire_proto_init() }
func file_proto_wire_proto_init() {
	if File_proto_wire_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_wire_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pod); i {
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
		file_proto_wire_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Link); i {
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
		file_proto_wire_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PodQuery); i {
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
		file_proto_wire_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SkipQuery); i {
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
		file_proto_wire_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoolResponse); i {
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
		file_proto_wire_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemotePod); i {
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
			RawDescriptor: file_proto_wire_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_proto_wire_proto_goTypes,
		DependencyIndexes: file_proto_wire_proto_depIdxs,
		MessageInfos:      file_proto_wire_proto_msgTypes,
	}.Build()
	File_proto_wire_proto = out.File
	file_proto_wire_proto_rawDesc = nil
	file_proto_wire_proto_goTypes = nil
	file_proto_wire_proto_depIdxs = nil
}
