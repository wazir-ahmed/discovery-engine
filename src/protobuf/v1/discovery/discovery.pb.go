// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.15.8
// source: v1/discovery/discovery.proto

package discovery

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PolicyType int32

const (
	// proto3 requires that first value in enum have numeric value of 0
	PolicyType_UNSPECIFIED PolicyType = 0
	PolicyType_KUBEARMOR   PolicyType = 1
	PolicyType_CILIUM      PolicyType = 2
	PolicyType_K8S_NETWORK PolicyType = 3
)

// Enum value maps for PolicyType.
var (
	PolicyType_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "KUBEARMOR",
		2: "CILIUM",
		3: "K8S_NETWORK",
	}
	PolicyType_value = map[string]int32{
		"UNSPECIFIED": 0,
		"KUBEARMOR":   1,
		"CILIUM":      2,
		"K8S_NETWORK": 3,
	}
)

func (x PolicyType) Enum() *PolicyType {
	p := new(PolicyType)
	*p = x
	return p
}

func (x PolicyType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PolicyType) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_discovery_discovery_proto_enumTypes[0].Descriptor()
}

func (PolicyType) Type() protoreflect.EnumType {
	return &file_v1_discovery_discovery_proto_enumTypes[0]
}

func (x PolicyType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PolicyType.Descriptor instead.
func (PolicyType) EnumDescriptor() ([]byte, []int) {
	return file_v1_discovery_discovery_proto_rawDescGZIP(), []int{0}
}

type GetPolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Follow     bool       `protobuf:"varint,1,opt,name=follow,proto3" json:"follow,omitempty"`
	Type       PolicyType `protobuf:"varint,2,opt,name=type,proto3,enum=v1.discovery.PolicyType" json:"type,omitempty"`
	Cluster    string     `protobuf:"bytes,3,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Namespace  string     `protobuf:"bytes,4,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Label      []string   `protobuf:"bytes,5,rep,name=label,proto3" json:"label,omitempty"`
	FromSource string     `protobuf:"bytes,6,opt,name=from_source,json=fromSource,proto3" json:"from_source,omitempty"`
}

func (x *GetPolicyRequest) Reset() {
	*x = GetPolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_discovery_discovery_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPolicyRequest) ProtoMessage() {}

func (x *GetPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_discovery_discovery_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPolicyRequest.ProtoReflect.Descriptor instead.
func (*GetPolicyRequest) Descriptor() ([]byte, []int) {
	return file_v1_discovery_discovery_proto_rawDescGZIP(), []int{0}
}

func (x *GetPolicyRequest) GetFollow() bool {
	if x != nil {
		return x.Follow
	}
	return false
}

func (x *GetPolicyRequest) GetType() PolicyType {
	if x != nil {
		return x.Type
	}
	return PolicyType_UNSPECIFIED
}

func (x *GetPolicyRequest) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

func (x *GetPolicyRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *GetPolicyRequest) GetLabel() []string {
	if x != nil {
		return x.Label
	}
	return nil
}

func (x *GetPolicyRequest) GetFromSource() string {
	if x != nil {
		return x.FromSource
	}
	return ""
}

type GetPolicyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Policy:
	//	*GetPolicyResponse_Kubearmor
	//	*GetPolicyResponse_Cilium
	//	*GetPolicyResponse_K8SNetwork
	Policy isGetPolicyResponse_Policy `protobuf_oneof:"policy"`
	Error  string                     `protobuf:"bytes,100,opt,name=error,proto3" json:"error,omitempty"`
	Time   *timestamppb.Timestamp     `protobuf:"bytes,101,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *GetPolicyResponse) Reset() {
	*x = GetPolicyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_discovery_discovery_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPolicyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPolicyResponse) ProtoMessage() {}

func (x *GetPolicyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_discovery_discovery_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPolicyResponse.ProtoReflect.Descriptor instead.
func (*GetPolicyResponse) Descriptor() ([]byte, []int) {
	return file_v1_discovery_discovery_proto_rawDescGZIP(), []int{1}
}

func (m *GetPolicyResponse) GetPolicy() isGetPolicyResponse_Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

func (x *GetPolicyResponse) GetKubearmor() *KubeArmorPolicy {
	if x, ok := x.GetPolicy().(*GetPolicyResponse_Kubearmor); ok {
		return x.Kubearmor
	}
	return nil
}

func (x *GetPolicyResponse) GetCilium() *CiliumPolicy {
	if x, ok := x.GetPolicy().(*GetPolicyResponse_Cilium); ok {
		return x.Cilium
	}
	return nil
}

func (x *GetPolicyResponse) GetK8SNetwork() *K8SNetworkPolicy {
	if x, ok := x.GetPolicy().(*GetPolicyResponse_K8SNetwork); ok {
		return x.K8SNetwork
	}
	return nil
}

func (x *GetPolicyResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *GetPolicyResponse) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

type isGetPolicyResponse_Policy interface {
	isGetPolicyResponse_Policy()
}

type GetPolicyResponse_Kubearmor struct {
	Kubearmor *KubeArmorPolicy `protobuf:"bytes,1,opt,name=kubearmor,proto3,oneof"`
}

type GetPolicyResponse_Cilium struct {
	Cilium *CiliumPolicy `protobuf:"bytes,2,opt,name=cilium,proto3,oneof"`
}

type GetPolicyResponse_K8SNetwork struct {
	K8SNetwork *K8SNetworkPolicy `protobuf:"bytes,3,opt,name=k8s_network,json=k8sNetwork,proto3,oneof"`
}

func (*GetPolicyResponse_Kubearmor) isGetPolicyResponse_Policy() {}

func (*GetPolicyResponse_Cilium) isGetPolicyResponse_Policy() {}

func (*GetPolicyResponse_K8SNetwork) isGetPolicyResponse_Policy() {}

type KubeArmorPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Yaml []byte `protobuf:"bytes,1,opt,name=yaml,proto3" json:"yaml,omitempty"`
}

func (x *KubeArmorPolicy) Reset() {
	*x = KubeArmorPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_discovery_discovery_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubeArmorPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubeArmorPolicy) ProtoMessage() {}

func (x *KubeArmorPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_v1_discovery_discovery_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubeArmorPolicy.ProtoReflect.Descriptor instead.
func (*KubeArmorPolicy) Descriptor() ([]byte, []int) {
	return file_v1_discovery_discovery_proto_rawDescGZIP(), []int{2}
}

func (x *KubeArmorPolicy) GetYaml() []byte {
	if x != nil {
		return x.Yaml
	}
	return nil
}

type CiliumPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Yaml []byte `protobuf:"bytes,1,opt,name=yaml,proto3" json:"yaml,omitempty"`
}

func (x *CiliumPolicy) Reset() {
	*x = CiliumPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_discovery_discovery_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CiliumPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CiliumPolicy) ProtoMessage() {}

func (x *CiliumPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_v1_discovery_discovery_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CiliumPolicy.ProtoReflect.Descriptor instead.
func (*CiliumPolicy) Descriptor() ([]byte, []int) {
	return file_v1_discovery_discovery_proto_rawDescGZIP(), []int{3}
}

func (x *CiliumPolicy) GetYaml() []byte {
	if x != nil {
		return x.Yaml
	}
	return nil
}

type K8SNetworkPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Yaml []byte `protobuf:"bytes,1,opt,name=yaml,proto3" json:"yaml,omitempty"`
}

func (x *K8SNetworkPolicy) Reset() {
	*x = K8SNetworkPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_discovery_discovery_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *K8SNetworkPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*K8SNetworkPolicy) ProtoMessage() {}

func (x *K8SNetworkPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_v1_discovery_discovery_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use K8SNetworkPolicy.ProtoReflect.Descriptor instead.
func (*K8SNetworkPolicy) Descriptor() ([]byte, []int) {
	return file_v1_discovery_discovery_proto_rawDescGZIP(), []int{4}
}

func (x *K8SNetworkPolicy) GetYaml() []byte {
	if x != nil {
		return x.Yaml
	}
	return nil
}

var File_v1_discovery_discovery_proto protoreflect.FileDescriptor

var file_v1_discovery_discovery_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x76, 0x31, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc7, 0x01,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x2c, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x72, 0x6f,
	0x6d, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x9b, 0x02, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a,
	0x09, 0x6b, 0x75, 0x62, 0x65, 0x61, 0x72, 0x6d, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e,
	0x4b, 0x75, 0x62, 0x65, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x48,
	0x00, 0x52, 0x09, 0x6b, 0x75, 0x62, 0x65, 0x61, 0x72, 0x6d, 0x6f, 0x72, 0x12, 0x34, 0x0a, 0x06,
	0x63, 0x69, 0x6c, 0x69, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x76,
	0x31, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x43, 0x69, 0x6c, 0x69,
	0x75, 0x6d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x48, 0x00, 0x52, 0x06, 0x63, 0x69, 0x6c, 0x69,
	0x75, 0x6d, 0x12, 0x41, 0x0a, 0x0b, 0x6b, 0x38, 0x73, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x4b, 0x38, 0x73, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x48, 0x00, 0x52, 0x0a, 0x6b, 0x38, 0x73, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x64,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2e, 0x0a, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22, 0x25, 0x0a, 0x0f, 0x4b, 0x75, 0x62, 0x65, 0x41, 0x72, 0x6d,
	0x6f, 0x72, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x61, 0x6d, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x79, 0x61, 0x6d, 0x6c, 0x22, 0x22, 0x0a, 0x0c,
	0x43, 0x69, 0x6c, 0x69, 0x75, 0x6d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x79, 0x61, 0x6d, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x79, 0x61, 0x6d, 0x6c,
	0x22, 0x26, 0x0a, 0x10, 0x4b, 0x38, 0x73, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x61, 0x6d, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x79, 0x61, 0x6d, 0x6c, 0x2a, 0x49, 0x0a, 0x0a, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x4b, 0x55, 0x42, 0x45, 0x41,
	0x52, 0x4d, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x49, 0x4c, 0x49, 0x55, 0x4d,
	0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x4b, 0x38, 0x53, 0x5f, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52,
	0x4b, 0x10, 0x03, 0x32, 0x5d, 0x0a, 0x09, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x12, 0x50, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x1e, 0x2e,
	0x76, 0x31, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x76, 0x31, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x30, 0x01, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x63, 0x63, 0x75, 0x6b, 0x6e, 0x6f, 0x78, 0x2f, 0x6b, 0x6e, 0x6f, 0x78, 0x41, 0x75,
	0x74, 0x6f, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_discovery_discovery_proto_rawDescOnce sync.Once
	file_v1_discovery_discovery_proto_rawDescData = file_v1_discovery_discovery_proto_rawDesc
)

func file_v1_discovery_discovery_proto_rawDescGZIP() []byte {
	file_v1_discovery_discovery_proto_rawDescOnce.Do(func() {
		file_v1_discovery_discovery_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_discovery_discovery_proto_rawDescData)
	})
	return file_v1_discovery_discovery_proto_rawDescData
}

var file_v1_discovery_discovery_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v1_discovery_discovery_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_v1_discovery_discovery_proto_goTypes = []interface{}{
	(PolicyType)(0),               // 0: v1.discovery.PolicyType
	(*GetPolicyRequest)(nil),      // 1: v1.discovery.GetPolicyRequest
	(*GetPolicyResponse)(nil),     // 2: v1.discovery.GetPolicyResponse
	(*KubeArmorPolicy)(nil),       // 3: v1.discovery.KubeArmorPolicy
	(*CiliumPolicy)(nil),          // 4: v1.discovery.CiliumPolicy
	(*K8SNetworkPolicy)(nil),      // 5: v1.discovery.K8sNetworkPolicy
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_v1_discovery_discovery_proto_depIdxs = []int32{
	0, // 0: v1.discovery.GetPolicyRequest.type:type_name -> v1.discovery.PolicyType
	3, // 1: v1.discovery.GetPolicyResponse.kubearmor:type_name -> v1.discovery.KubeArmorPolicy
	4, // 2: v1.discovery.GetPolicyResponse.cilium:type_name -> v1.discovery.CiliumPolicy
	5, // 3: v1.discovery.GetPolicyResponse.k8s_network:type_name -> v1.discovery.K8sNetworkPolicy
	6, // 4: v1.discovery.GetPolicyResponse.time:type_name -> google.protobuf.Timestamp
	1, // 5: v1.discovery.Discovery.GetPolicy:input_type -> v1.discovery.GetPolicyRequest
	2, // 6: v1.discovery.Discovery.GetPolicy:output_type -> v1.discovery.GetPolicyResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_v1_discovery_discovery_proto_init() }
func file_v1_discovery_discovery_proto_init() {
	if File_v1_discovery_discovery_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_discovery_discovery_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPolicyRequest); i {
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
		file_v1_discovery_discovery_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPolicyResponse); i {
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
		file_v1_discovery_discovery_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubeArmorPolicy); i {
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
		file_v1_discovery_discovery_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CiliumPolicy); i {
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
		file_v1_discovery_discovery_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*K8SNetworkPolicy); i {
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
	file_v1_discovery_discovery_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*GetPolicyResponse_Kubearmor)(nil),
		(*GetPolicyResponse_Cilium)(nil),
		(*GetPolicyResponse_K8SNetwork)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_discovery_discovery_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_discovery_discovery_proto_goTypes,
		DependencyIndexes: file_v1_discovery_discovery_proto_depIdxs,
		EnumInfos:         file_v1_discovery_discovery_proto_enumTypes,
		MessageInfos:      file_v1_discovery_discovery_proto_msgTypes,
	}.Build()
	File_v1_discovery_discovery_proto = out.File
	file_v1_discovery_discovery_proto_rawDesc = nil
	file_v1_discovery_discovery_proto_goTypes = nil
	file_v1_discovery_discovery_proto_depIdxs = nil
}
