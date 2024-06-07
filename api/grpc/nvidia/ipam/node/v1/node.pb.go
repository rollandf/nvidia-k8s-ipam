// Copyright 2023, NVIDIA CORPORATION & AFFILIATES
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: nvidia/ipam/node/v1/node.proto

package v1

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

// indicates type of the pool which is referred by the name
type PoolType int32

const (
	PoolType_POOL_TYPE_UNSPECIFIED PoolType = 0
	PoolType_POOL_TYPE_IPPOOL      PoolType = 1
	PoolType_POOL_TYPE_CIDRPOOL    PoolType = 2
)

// Enum value maps for PoolType.
var (
	PoolType_name = map[int32]string{
		0: "POOL_TYPE_UNSPECIFIED",
		1: "POOL_TYPE_IPPOOL",
		2: "POOL_TYPE_CIDRPOOL",
	}
	PoolType_value = map[string]int32{
		"POOL_TYPE_UNSPECIFIED": 0,
		"POOL_TYPE_IPPOOL":      1,
		"POOL_TYPE_CIDRPOOL":    2,
	}
)

func (x PoolType) Enum() *PoolType {
	p := new(PoolType)
	*p = x
	return p
}

func (x PoolType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PoolType) Descriptor() protoreflect.EnumDescriptor {
	return file_nvidia_ipam_node_v1_node_proto_enumTypes[0].Descriptor()
}

func (PoolType) Type() protoreflect.EnumType {
	return &file_nvidia_ipam_node_v1_node_proto_enumTypes[0]
}

func (x PoolType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PoolType.Descriptor instead.
func (PoolType) EnumDescriptor() ([]byte, []int) {
	return file_nvidia_ipam_node_v1_node_proto_rawDescGZIP(), []int{0}
}

// AllocateRequest contains parameters for Allocate rpc call
type AllocateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// required, IPAMParameters contains parameters IPAM parameters related to the request
	Parameters *IPAMParameters `protobuf:"bytes,1,opt,name=parameters,proto3" json:"parameters,omitempty"`
}

func (x *AllocateRequest) Reset() {
	*x = AllocateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nvidia_ipam_node_v1_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllocateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllocateRequest) ProtoMessage() {}

func (x *AllocateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nvidia_ipam_node_v1_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllocateRequest.ProtoReflect.Descriptor instead.
func (*AllocateRequest) Descriptor() ([]byte, []int) {
	return file_nvidia_ipam_node_v1_node_proto_rawDescGZIP(), []int{0}
}

func (x *AllocateRequest) GetParameters() *IPAMParameters {
	if x != nil {
		return x.Parameters
	}
	return nil
}

// IPAMParameters common message which contains information used in all rpc calls
type IPAMParameters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// required, list of pools in which IP addresses should be allocated,
	// must contain one or two elements (dual-stack, IPv4 + IPv6 use-case)
	Pools []string `protobuf:"bytes,1,rep,name=pools,proto3" json:"pools,omitempty"`
	// required, a unique plaintext identifier for a container, allocated by the runtime
	CniContainerid string `protobuf:"bytes,2,opt,name=cni_containerid,json=cniContainerid,proto3" json:"cni_containerid,omitempty"`
	// required, name of the interface inside the container
	CniIfname string `protobuf:"bytes,3,opt,name=cni_ifname,json=cniIfname,proto3" json:"cni_ifname,omitempty"`
	// required, additional metadata to identify IP allocation
	Metadata *IPAMMetadata `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// type of the pool which is refered by the name in the pools field
	PoolType PoolType `protobuf:"varint,5,opt,name=pool_type,json=poolType,proto3,enum=nvidia.ipam.node.v1.PoolType" json:"pool_type,omitempty"`
	// conatins IP that were statically requested
	RequestedIps []string `protobuf:"bytes,6,rep,name=requested_ips,json=requestedIps,proto3" json:"requested_ips,omitempty"`
}

func (x *IPAMParameters) Reset() {
	*x = IPAMParameters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nvidia_ipam_node_v1_node_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPAMParameters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPAMParameters) ProtoMessage() {}

func (x *IPAMParameters) ProtoReflect() protoreflect.Message {
	mi := &file_nvidia_ipam_node_v1_node_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPAMParameters.ProtoReflect.Descriptor instead.
func (*IPAMParameters) Descriptor() ([]byte, []int) {
	return file_nvidia_ipam_node_v1_node_proto_rawDescGZIP(), []int{1}
}

func (x *IPAMParameters) GetPools() []string {
	if x != nil {
		return x.Pools
	}
	return nil
}

func (x *IPAMParameters) GetCniContainerid() string {
	if x != nil {
		return x.CniContainerid
	}
	return ""
}

func (x *IPAMParameters) GetCniIfname() string {
	if x != nil {
		return x.CniIfname
	}
	return ""
}

func (x *IPAMParameters) GetMetadata() *IPAMMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *IPAMParameters) GetPoolType() PoolType {
	if x != nil {
		return x.PoolType
	}
	return PoolType_POOL_TYPE_UNSPECIFIED
}

func (x *IPAMParameters) GetRequestedIps() []string {
	if x != nil {
		return x.RequestedIps
	}
	return nil
}

// IPAMMetadata contains metadata for IPAM calls
type IPAMMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// required, name of the k8s pod
	K8SPodName string `protobuf:"bytes,1,opt,name=k8s_pod_name,json=k8sPodName,proto3" json:"k8s_pod_name,omitempty"`
	// required, namespace of the k8s pod
	K8SPodNamespace string `protobuf:"bytes,2,opt,name=k8s_pod_namespace,json=k8sPodNamespace,proto3" json:"k8s_pod_namespace,omitempty"`
	// optional, UID of the k8s pod, k8s_pod_uid exist in containerd >= 1.6 cr-io >= 0.3
	K8SPodUid string `protobuf:"bytes,3,opt,name=k8s_pod_uid,json=k8sPodUid,proto3" json:"k8s_pod_uid,omitempty"`
	// optional, PCI device ID related to the allocation
	DeviceId string `protobuf:"bytes,4,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
}

func (x *IPAMMetadata) Reset() {
	*x = IPAMMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nvidia_ipam_node_v1_node_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPAMMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPAMMetadata) ProtoMessage() {}

func (x *IPAMMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_nvidia_ipam_node_v1_node_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPAMMetadata.ProtoReflect.Descriptor instead.
func (*IPAMMetadata) Descriptor() ([]byte, []int) {
	return file_nvidia_ipam_node_v1_node_proto_rawDescGZIP(), []int{2}
}

func (x *IPAMMetadata) GetK8SPodName() string {
	if x != nil {
		return x.K8SPodName
	}
	return ""
}

func (x *IPAMMetadata) GetK8SPodNamespace() string {
	if x != nil {
		return x.K8SPodNamespace
	}
	return ""
}

func (x *IPAMMetadata) GetK8SPodUid() string {
	if x != nil {
		return x.K8SPodUid
	}
	return ""
}

func (x *IPAMMetadata) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

// IsAllocatedRequest contains parameters for IsAllocated rpc call
type IsAllocatedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// required, IPAMParameters contains parameters IPAM parameters related to the request
	Parameters *IPAMParameters `protobuf:"bytes,1,opt,name=parameters,proto3" json:"parameters,omitempty"`
}

func (x *IsAllocatedRequest) Reset() {
	*x = IsAllocatedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nvidia_ipam_node_v1_node_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsAllocatedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsAllocatedRequest) ProtoMessage() {}

func (x *IsAllocatedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nvidia_ipam_node_v1_node_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsAllocatedRequest.ProtoReflect.Descriptor instead.
func (*IsAllocatedRequest) Descriptor() ([]byte, []int) {
	return file_nvidia_ipam_node_v1_node_proto_rawDescGZIP(), []int{3}
}

func (x *IsAllocatedRequest) GetParameters() *IPAMParameters {
	if x != nil {
		return x.Parameters
	}
	return nil
}

// DeallocateRequest contains parameters for Deallocate rpc call
type DeallocateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// required, IPAMParameters contains parameters IPAM parameters related to the request
	Parameters *IPAMParameters `protobuf:"bytes,1,opt,name=parameters,proto3" json:"parameters,omitempty"`
}

func (x *DeallocateRequest) Reset() {
	*x = DeallocateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nvidia_ipam_node_v1_node_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeallocateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeallocateRequest) ProtoMessage() {}

func (x *DeallocateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nvidia_ipam_node_v1_node_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeallocateRequest.ProtoReflect.Descriptor instead.
func (*DeallocateRequest) Descriptor() ([]byte, []int) {
	return file_nvidia_ipam_node_v1_node_proto_rawDescGZIP(), []int{4}
}

func (x *DeallocateRequest) GetParameters() *IPAMParameters {
	if x != nil {
		return x.Parameters
	}
	return nil
}

// AllocateResponse contains reply for Allocate rpc call
type AllocateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// list of allocated IPs
	Allocations []*AllocationInfo `protobuf:"bytes,1,rep,name=allocations,proto3" json:"allocations,omitempty"`
}

func (x *AllocateResponse) Reset() {
	*x = AllocateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nvidia_ipam_node_v1_node_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllocateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllocateResponse) ProtoMessage() {}

func (x *AllocateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nvidia_ipam_node_v1_node_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllocateResponse.ProtoReflect.Descriptor instead.
func (*AllocateResponse) Descriptor() ([]byte, []int) {
	return file_nvidia_ipam_node_v1_node_proto_rawDescGZIP(), []int{5}
}

func (x *AllocateResponse) GetAllocations() []*AllocationInfo {
	if x != nil {
		return x.Allocations
	}
	return nil
}

// AllocationInfo contains information about the allocation
type AllocationInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of the pool in which this IP was allocated
	Pool string `protobuf:"bytes,1,opt,name=pool,proto3" json:"pool,omitempty"`
	// allocated IP together with prefix length, e.g. 192.168.10.33/24
	Ip string `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	// gateway for allocated IP
	Gateway string `protobuf:"bytes,3,opt,name=gateway,proto3" json:"gateway,omitempty"`
	// type of the pool which is refered by the name in the pools field
	PoolType PoolType `protobuf:"varint,4,opt,name=pool_type,json=poolType,proto3,enum=nvidia.ipam.node.v1.PoolType" json:"pool_type,omitempty"`
}

func (x *AllocationInfo) Reset() {
	*x = AllocationInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nvidia_ipam_node_v1_node_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllocationInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllocationInfo) ProtoMessage() {}

func (x *AllocationInfo) ProtoReflect() protoreflect.Message {
	mi := &file_nvidia_ipam_node_v1_node_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllocationInfo.ProtoReflect.Descriptor instead.
func (*AllocationInfo) Descriptor() ([]byte, []int) {
	return file_nvidia_ipam_node_v1_node_proto_rawDescGZIP(), []int{6}
}

func (x *AllocationInfo) GetPool() string {
	if x != nil {
		return x.Pool
	}
	return ""
}

func (x *AllocationInfo) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *AllocationInfo) GetGateway() string {
	if x != nil {
		return x.Gateway
	}
	return ""
}

func (x *AllocationInfo) GetPoolType() PoolType {
	if x != nil {
		return x.PoolType
	}
	return PoolType_POOL_TYPE_UNSPECIFIED
}

// IsAllocatedReply contains reply for IsAllocated rpc call
type IsAllocatedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *IsAllocatedResponse) Reset() {
	*x = IsAllocatedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nvidia_ipam_node_v1_node_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsAllocatedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsAllocatedResponse) ProtoMessage() {}

func (x *IsAllocatedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nvidia_ipam_node_v1_node_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsAllocatedResponse.ProtoReflect.Descriptor instead.
func (*IsAllocatedResponse) Descriptor() ([]byte, []int) {
	return file_nvidia_ipam_node_v1_node_proto_rawDescGZIP(), []int{7}
}

// DeallocateReply contains reply for Deallocate rpc call
type DeallocateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeallocateResponse) Reset() {
	*x = DeallocateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nvidia_ipam_node_v1_node_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeallocateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeallocateResponse) ProtoMessage() {}

func (x *DeallocateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nvidia_ipam_node_v1_node_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeallocateResponse.ProtoReflect.Descriptor instead.
func (*DeallocateResponse) Descriptor() ([]byte, []int) {
	return file_nvidia_ipam_node_v1_node_proto_rawDescGZIP(), []int{8}
}

var File_nvidia_ipam_node_v1_node_proto protoreflect.FileDescriptor

var file_nvidia_ipam_node_v1_node_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6e, 0x76, 0x69, 0x64, 0x69, 0x61, 0x2f, 0x69, 0x70, 0x61, 0x6d, 0x2f, 0x6e, 0x6f,
	0x64, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x13, 0x6e, 0x76, 0x69, 0x64, 0x69, 0x61, 0x2e, 0x69, 0x70, 0x61, 0x6d, 0x2e, 0x6e, 0x6f,
	0x64, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x56, 0x0a, 0x0f, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6e,
	0x76, 0x69, 0x64, 0x69, 0x61, 0x2e, 0x69, 0x70, 0x61, 0x6d, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x49, 0x50, 0x41, 0x4d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x73, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x22, 0x8e, 0x02,
	0x0a, 0x0e, 0x49, 0x50, 0x41, 0x4d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x05, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6e, 0x69, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x63, 0x6e, 0x69, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x69, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x6e, 0x69, 0x5f, 0x69, 0x66, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6e, 0x69, 0x49, 0x66, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3d,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x6e, 0x76, 0x69, 0x64, 0x69, 0x61, 0x2e, 0x69, 0x70, 0x61, 0x6d, 0x2e, 0x6e,
	0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x50, 0x41, 0x4d, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3a, 0x0a,
	0x09, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1d, 0x2e, 0x6e, 0x76, 0x69, 0x64, 0x69, 0x61, 0x2e, 0x69, 0x70, 0x61, 0x6d, 0x2e, 0x6e,
	0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6f, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x08, 0x70, 0x6f, 0x6f, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x69, 0x70, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x49, 0x70, 0x73, 0x22, 0x99,
	0x01, 0x0a, 0x0c, 0x49, 0x50, 0x41, 0x4d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x20, 0x0a, 0x0c, 0x6b, 0x38, 0x73, 0x5f, 0x70, 0x6f, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6b, 0x38, 0x73, 0x50, 0x6f, 0x64, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x2a, 0x0a, 0x11, 0x6b, 0x38, 0x73, 0x5f, 0x70, 0x6f, 0x64, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6b, 0x38,
	0x73, 0x50, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1e, 0x0a,
	0x0b, 0x6b, 0x38, 0x73, 0x5f, 0x70, 0x6f, 0x64, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6b, 0x38, 0x73, 0x50, 0x6f, 0x64, 0x55, 0x69, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x22, 0x59, 0x0a, 0x12, 0x49, 0x73,
	0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x43, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6e, 0x76, 0x69, 0x64, 0x69, 0x61, 0x2e, 0x69, 0x70,
	0x61, 0x6d, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x50, 0x41, 0x4d, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x73, 0x22, 0x58, 0x0a, 0x11, 0x44, 0x65, 0x61, 0x6c, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x0a, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x6e, 0x76, 0x69, 0x64, 0x69, 0x61, 0x2e, 0x69, 0x70, 0x61, 0x6d, 0x2e, 0x6e, 0x6f, 0x64,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x50, 0x41, 0x4d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x73, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x22,
	0x59, 0x0a, 0x10, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0b, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6e, 0x76, 0x69, 0x64, 0x69,
	0x61, 0x2e, 0x69, 0x70, 0x61, 0x6d, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x61,
	0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x8a, 0x01, 0x0a, 0x0e, 0x41,
	0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x6f, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x6f,
	0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x70, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x3a, 0x0a, 0x09, 0x70,
	0x6f, 0x6f, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d,
	0x2e, 0x6e, 0x76, 0x69, 0x64, 0x69, 0x61, 0x2e, 0x69, 0x70, 0x61, 0x6d, 0x2e, 0x6e, 0x6f, 0x64,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6f, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x70,
	0x6f, 0x6f, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x49, 0x73, 0x41, 0x6c, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14,
	0x0a, 0x12, 0x44, 0x65, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x53, 0x0a, 0x08, 0x50, 0x6f, 0x6f, 0x6c, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x19, 0x0a, 0x15, 0x50, 0x4f, 0x4f, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x50,
	0x4f, 0x4f, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x50, 0x50, 0x4f, 0x4f, 0x4c, 0x10,
	0x01, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x4f, 0x4f, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43,
	0x49, 0x44, 0x52, 0x50, 0x4f, 0x4f, 0x4c, 0x10, 0x02, 0x32, 0xad, 0x02, 0x0a, 0x0b, 0x49, 0x50,
	0x41, 0x4d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x08, 0x41, 0x6c, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x65, 0x12, 0x24, 0x2e, 0x6e, 0x76, 0x69, 0x64, 0x69, 0x61, 0x2e, 0x69,
	0x70, 0x61, 0x6d, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6e, 0x76,
	0x69, 0x64, 0x69, 0x61, 0x2e, 0x69, 0x70, 0x61, 0x6d, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x62, 0x0a, 0x0b, 0x49, 0x73, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x27, 0x2e, 0x6e, 0x76, 0x69, 0x64, 0x69, 0x61, 0x2e, 0x69, 0x70, 0x61,
	0x6d, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x73, 0x41, 0x6c, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x6e,
	0x76, 0x69, 0x64, 0x69, 0x61, 0x2e, 0x69, 0x70, 0x61, 0x6d, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x49, 0x73, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x0a, 0x44, 0x65, 0x61, 0x6c,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x12, 0x26, 0x2e, 0x6e, 0x76, 0x69, 0x64, 0x69, 0x61, 0x2e,
	0x69, 0x70, 0x61, 0x6d, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x61,
	0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27,
	0x2e, 0x6e, 0x76, 0x69, 0x64, 0x69, 0x61, 0x2e, 0x69, 0x70, 0x61, 0x6d, 0x2e, 0x6e, 0x6f, 0x64,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_nvidia_ipam_node_v1_node_proto_rawDescOnce sync.Once
	file_nvidia_ipam_node_v1_node_proto_rawDescData = file_nvidia_ipam_node_v1_node_proto_rawDesc
)

func file_nvidia_ipam_node_v1_node_proto_rawDescGZIP() []byte {
	file_nvidia_ipam_node_v1_node_proto_rawDescOnce.Do(func() {
		file_nvidia_ipam_node_v1_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_nvidia_ipam_node_v1_node_proto_rawDescData)
	})
	return file_nvidia_ipam_node_v1_node_proto_rawDescData
}

var file_nvidia_ipam_node_v1_node_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_nvidia_ipam_node_v1_node_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_nvidia_ipam_node_v1_node_proto_goTypes = []interface{}{
	(PoolType)(0),               // 0: nvidia.ipam.node.v1.PoolType
	(*AllocateRequest)(nil),     // 1: nvidia.ipam.node.v1.AllocateRequest
	(*IPAMParameters)(nil),      // 2: nvidia.ipam.node.v1.IPAMParameters
	(*IPAMMetadata)(nil),        // 3: nvidia.ipam.node.v1.IPAMMetadata
	(*IsAllocatedRequest)(nil),  // 4: nvidia.ipam.node.v1.IsAllocatedRequest
	(*DeallocateRequest)(nil),   // 5: nvidia.ipam.node.v1.DeallocateRequest
	(*AllocateResponse)(nil),    // 6: nvidia.ipam.node.v1.AllocateResponse
	(*AllocationInfo)(nil),      // 7: nvidia.ipam.node.v1.AllocationInfo
	(*IsAllocatedResponse)(nil), // 8: nvidia.ipam.node.v1.IsAllocatedResponse
	(*DeallocateResponse)(nil),  // 9: nvidia.ipam.node.v1.DeallocateResponse
}
var file_nvidia_ipam_node_v1_node_proto_depIdxs = []int32{
	2,  // 0: nvidia.ipam.node.v1.AllocateRequest.parameters:type_name -> nvidia.ipam.node.v1.IPAMParameters
	3,  // 1: nvidia.ipam.node.v1.IPAMParameters.metadata:type_name -> nvidia.ipam.node.v1.IPAMMetadata
	0,  // 2: nvidia.ipam.node.v1.IPAMParameters.pool_type:type_name -> nvidia.ipam.node.v1.PoolType
	2,  // 3: nvidia.ipam.node.v1.IsAllocatedRequest.parameters:type_name -> nvidia.ipam.node.v1.IPAMParameters
	2,  // 4: nvidia.ipam.node.v1.DeallocateRequest.parameters:type_name -> nvidia.ipam.node.v1.IPAMParameters
	7,  // 5: nvidia.ipam.node.v1.AllocateResponse.allocations:type_name -> nvidia.ipam.node.v1.AllocationInfo
	0,  // 6: nvidia.ipam.node.v1.AllocationInfo.pool_type:type_name -> nvidia.ipam.node.v1.PoolType
	1,  // 7: nvidia.ipam.node.v1.IPAMService.Allocate:input_type -> nvidia.ipam.node.v1.AllocateRequest
	4,  // 8: nvidia.ipam.node.v1.IPAMService.IsAllocated:input_type -> nvidia.ipam.node.v1.IsAllocatedRequest
	5,  // 9: nvidia.ipam.node.v1.IPAMService.Deallocate:input_type -> nvidia.ipam.node.v1.DeallocateRequest
	6,  // 10: nvidia.ipam.node.v1.IPAMService.Allocate:output_type -> nvidia.ipam.node.v1.AllocateResponse
	8,  // 11: nvidia.ipam.node.v1.IPAMService.IsAllocated:output_type -> nvidia.ipam.node.v1.IsAllocatedResponse
	9,  // 12: nvidia.ipam.node.v1.IPAMService.Deallocate:output_type -> nvidia.ipam.node.v1.DeallocateResponse
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_nvidia_ipam_node_v1_node_proto_init() }
func file_nvidia_ipam_node_v1_node_proto_init() {
	if File_nvidia_ipam_node_v1_node_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nvidia_ipam_node_v1_node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllocateRequest); i {
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
		file_nvidia_ipam_node_v1_node_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPAMParameters); i {
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
		file_nvidia_ipam_node_v1_node_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPAMMetadata); i {
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
		file_nvidia_ipam_node_v1_node_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsAllocatedRequest); i {
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
		file_nvidia_ipam_node_v1_node_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeallocateRequest); i {
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
		file_nvidia_ipam_node_v1_node_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllocateResponse); i {
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
		file_nvidia_ipam_node_v1_node_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllocationInfo); i {
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
		file_nvidia_ipam_node_v1_node_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsAllocatedResponse); i {
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
		file_nvidia_ipam_node_v1_node_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeallocateResponse); i {
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
			RawDescriptor: file_nvidia_ipam_node_v1_node_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nvidia_ipam_node_v1_node_proto_goTypes,
		DependencyIndexes: file_nvidia_ipam_node_v1_node_proto_depIdxs,
		EnumInfos:         file_nvidia_ipam_node_v1_node_proto_enumTypes,
		MessageInfos:      file_nvidia_ipam_node_v1_node_proto_msgTypes,
	}.Build()
	File_nvidia_ipam_node_v1_node_proto = out.File
	file_nvidia_ipam_node_v1_node_proto_rawDesc = nil
	file_nvidia_ipam_node_v1_node_proto_goTypes = nil
	file_nvidia_ipam_node_v1_node_proto_depIdxs = nil
}
