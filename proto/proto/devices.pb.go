// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.3
// source: devices.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Device struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name              string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Owner             string                 `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	PublicKey         string                 `protobuf:"bytes,3,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	Address           string                 `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	CreatedAt         *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Connected         bool                   `protobuf:"varint,6,opt,name=connected,proto3" json:"connected,omitempty"`
	LastHandshakeTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=last_handshake_time,json=lastHandshakeTime,proto3" json:"last_handshake_time,omitempty"`
	ReceiveBytes      int64                  `protobuf:"varint,8,opt,name=receive_bytes,json=receiveBytes,proto3" json:"receive_bytes,omitempty"`
	TransmitBytes     int64                  `protobuf:"varint,9,opt,name=transmit_bytes,json=transmitBytes,proto3" json:"transmit_bytes,omitempty"`
	Endpoint          string                 `protobuf:"bytes,10,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	OwnerName         string                 `protobuf:"bytes,11,opt,name=owner_name,json=ownerName,proto3" json:"owner_name,omitempty"`
	OwnerEmail        string                 `protobuf:"bytes,12,opt,name=owner_email,json=ownerEmail,proto3" json:"owner_email,omitempty"`
}

func (x *Device) Reset() {
	*x = Device{}
	if protoimpl.UnsafeEnabled {
		mi := &file_devices_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Device) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Device) ProtoMessage() {}

func (x *Device) ProtoReflect() protoreflect.Message {
	mi := &file_devices_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Device.ProtoReflect.Descriptor instead.
func (*Device) Descriptor() ([]byte, []int) {
	return file_devices_proto_rawDescGZIP(), []int{0}
}

func (x *Device) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Device) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *Device) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *Device) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Device) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Device) GetConnected() bool {
	if x != nil {
		return x.Connected
	}
	return false
}

func (x *Device) GetLastHandshakeTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastHandshakeTime
	}
	return nil
}

func (x *Device) GetReceiveBytes() int64 {
	if x != nil {
		return x.ReceiveBytes
	}
	return 0
}

func (x *Device) GetTransmitBytes() int64 {
	if x != nil {
		return x.TransmitBytes
	}
	return 0
}

func (x *Device) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *Device) GetOwnerName() string {
	if x != nil {
		return x.OwnerName
	}
	return ""
}

func (x *Device) GetOwnerEmail() string {
	if x != nil {
		return x.OwnerEmail
	}
	return ""
}

type AddDeviceReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	PublicKey string `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (x *AddDeviceReq) Reset() {
	*x = AddDeviceReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_devices_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddDeviceReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddDeviceReq) ProtoMessage() {}

func (x *AddDeviceReq) ProtoReflect() protoreflect.Message {
	mi := &file_devices_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddDeviceReq.ProtoReflect.Descriptor instead.
func (*AddDeviceReq) Descriptor() ([]byte, []int) {
	return file_devices_proto_rawDescGZIP(), []int{1}
}

func (x *AddDeviceReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddDeviceReq) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

type ListDevicesReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListDevicesReq) Reset() {
	*x = ListDevicesReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_devices_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDevicesReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDevicesReq) ProtoMessage() {}

func (x *ListDevicesReq) ProtoReflect() protoreflect.Message {
	mi := &file_devices_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDevicesReq.ProtoReflect.Descriptor instead.
func (*ListDevicesReq) Descriptor() ([]byte, []int) {
	return file_devices_proto_rawDescGZIP(), []int{2}
}

type ListDevicesRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Device `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListDevicesRes) Reset() {
	*x = ListDevicesRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_devices_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDevicesRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDevicesRes) ProtoMessage() {}

func (x *ListDevicesRes) ProtoReflect() protoreflect.Message {
	mi := &file_devices_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDevicesRes.ProtoReflect.Descriptor instead.
func (*ListDevicesRes) Descriptor() ([]byte, []int) {
	return file_devices_proto_rawDescGZIP(), []int{3}
}

func (x *ListDevicesRes) GetItems() []*Device {
	if x != nil {
		return x.Items
	}
	return nil
}

type DeleteDeviceReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// admin's may delete a device owned
	// by someone other than the current user
	// if empty, defaults to the current user
	Owner *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *DeleteDeviceReq) Reset() {
	*x = DeleteDeviceReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_devices_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDeviceReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDeviceReq) ProtoMessage() {}

func (x *DeleteDeviceReq) ProtoReflect() protoreflect.Message {
	mi := &file_devices_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDeviceReq.ProtoReflect.Descriptor instead.
func (*DeleteDeviceReq) Descriptor() ([]byte, []int) {
	return file_devices_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteDeviceReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeleteDeviceReq) GetOwner() *wrapperspb.StringValue {
	if x != nil {
		return x.Owner
	}
	return nil
}

type ListAllDevicesReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListAllDevicesReq) Reset() {
	*x = ListAllDevicesReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_devices_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAllDevicesReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAllDevicesReq) ProtoMessage() {}

func (x *ListAllDevicesReq) ProtoReflect() protoreflect.Message {
	mi := &file_devices_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAllDevicesReq.ProtoReflect.Descriptor instead.
func (*ListAllDevicesReq) Descriptor() ([]byte, []int) {
	return file_devices_proto_rawDescGZIP(), []int{5}
}

type ListAllDevicesRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Device `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListAllDevicesRes) Reset() {
	*x = ListAllDevicesRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_devices_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAllDevicesRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAllDevicesRes) ProtoMessage() {}

func (x *ListAllDevicesRes) ProtoReflect() protoreflect.Message {
	mi := &file_devices_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAllDevicesRes.ProtoReflect.Descriptor instead.
func (*ListAllDevicesRes) Descriptor() ([]byte, []int) {
	return file_devices_proto_rawDescGZIP(), []int{6}
}

func (x *ListAllDevicesRes) GetItems() []*Device {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_devices_proto protoreflect.FileDescriptor

var file_devices_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb8, 0x03, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x4a, 0x0a, 0x13, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x11, 0x6c, 0x61, 0x73, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68,
	0x61, 0x6b, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x0e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x42, 0x79,
	0x74, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x22,
	0x41, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b,
	0x65, 0x79, 0x22, 0x10, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x22, 0x35, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x59, 0x0a, 0x0f, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x22, 0x13, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c,
	0x6c, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x22, 0x38, 0x0a, 0x11, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x12, 0x23, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0x93, 0x02, 0x0a, 0x07, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x12, 0x31, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x13,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x70, 0x65, 0x63,
	0x69, 0x66, 0x69, 0x63, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x22,
	0x00, 0x12, 0x40, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x6c, 0x6c, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x1a,
	0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x26, 0x5a, 0x24, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x61, 0x61, 0x73, 0x2d, 0x61,
	0x70, 0x70, 0x2f, 0x57, 0x61, 0x61, 0x53, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_devices_proto_rawDescOnce sync.Once
	file_devices_proto_rawDescData = file_devices_proto_rawDesc
)

func file_devices_proto_rawDescGZIP() []byte {
	file_devices_proto_rawDescOnce.Do(func() {
		file_devices_proto_rawDescData = protoimpl.X.CompressGZIP(file_devices_proto_rawDescData)
	})
	return file_devices_proto_rawDescData
}

var file_devices_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_devices_proto_goTypes = []interface{}{
	(*Device)(nil),                 // 0: proto.Device
	(*AddDeviceReq)(nil),           // 1: proto.AddDeviceReq
	(*ListDevicesReq)(nil),         // 2: proto.ListDevicesReq
	(*ListDevicesRes)(nil),         // 3: proto.ListDevicesRes
	(*DeleteDeviceReq)(nil),        // 4: proto.DeleteDeviceReq
	(*ListAllDevicesReq)(nil),      // 5: proto.ListAllDevicesReq
	(*ListAllDevicesRes)(nil),      // 6: proto.ListAllDevicesRes
	(*timestamppb.Timestamp)(nil),  // 7: google.protobuf.Timestamp
	(*wrapperspb.StringValue)(nil), // 8: google.protobuf.StringValue
	(*emptypb.Empty)(nil),          // 9: google.protobuf.Empty
}
var file_devices_proto_depIdxs = []int32{
	7, // 0: proto.Device.created_at:type_name -> google.protobuf.Timestamp
	7, // 1: proto.Device.last_handshake_time:type_name -> google.protobuf.Timestamp
	0, // 2: proto.ListDevicesRes.items:type_name -> proto.Device
	8, // 3: proto.DeleteDeviceReq.owner:type_name -> google.protobuf.StringValue
	0, // 4: proto.ListAllDevicesRes.items:type_name -> proto.Device
	1, // 5: proto.Devices.AddDevice:input_type -> proto.AddDeviceReq
	2, // 6: proto.Devices.ListSpecificDeviceForUser:input_type -> proto.ListDevicesReq
	4, // 7: proto.Devices.DeleteDevice:input_type -> proto.DeleteDeviceReq
	5, // 8: proto.Devices.ListAllDevices:input_type -> proto.ListAllDevicesReq
	0, // 9: proto.Devices.AddDevice:output_type -> proto.Device
	3, // 10: proto.Devices.ListSpecificDeviceForUser:output_type -> proto.ListDevicesRes
	9, // 11: proto.Devices.DeleteDevice:output_type -> google.protobuf.Empty
	6, // 12: proto.Devices.ListAllDevices:output_type -> proto.ListAllDevicesRes
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_devices_proto_init() }
func file_devices_proto_init() {
	if File_devices_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_devices_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Device); i {
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
		file_devices_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddDeviceReq); i {
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
		file_devices_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDevicesReq); i {
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
		file_devices_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDevicesRes); i {
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
		file_devices_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDeviceReq); i {
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
		file_devices_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAllDevicesReq); i {
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
		file_devices_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAllDevicesRes); i {
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
			RawDescriptor: file_devices_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_devices_proto_goTypes,
		DependencyIndexes: file_devices_proto_depIdxs,
		MessageInfos:      file_devices_proto_msgTypes,
	}.Build()
	File_devices_proto = out.File
	file_devices_proto_rawDesc = nil
	file_devices_proto_goTypes = nil
	file_devices_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DevicesClient is the client API for Devices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DevicesClient interface {
	AddDevice(ctx context.Context, in *AddDeviceReq, opts ...grpc.CallOption) (*Device, error)
	ListSpecificDeviceForUser(ctx context.Context, in *ListDevicesReq, opts ...grpc.CallOption) (*ListDevicesRes, error)
	DeleteDevice(ctx context.Context, in *DeleteDeviceReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListAllDevices(ctx context.Context, in *ListAllDevicesReq, opts ...grpc.CallOption) (*ListAllDevicesRes, error)
}

type devicesClient struct {
	cc grpc.ClientConnInterface
}

func NewDevicesClient(cc grpc.ClientConnInterface) DevicesClient {
	return &devicesClient{cc}
}

func (c *devicesClient) AddDevice(ctx context.Context, in *AddDeviceReq, opts ...grpc.CallOption) (*Device, error) {
	out := new(Device)
	err := c.cc.Invoke(ctx, "/proto.Devices/AddDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesClient) ListSpecificDeviceForUser(ctx context.Context, in *ListDevicesReq, opts ...grpc.CallOption) (*ListDevicesRes, error) {
	out := new(ListDevicesRes)
	err := c.cc.Invoke(ctx, "/proto.Devices/ListSpecificDeviceForUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesClient) DeleteDevice(ctx context.Context, in *DeleteDeviceReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.Devices/DeleteDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesClient) ListAllDevices(ctx context.Context, in *ListAllDevicesReq, opts ...grpc.CallOption) (*ListAllDevicesRes, error) {
	out := new(ListAllDevicesRes)
	err := c.cc.Invoke(ctx, "/proto.Devices/ListAllDevices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DevicesServer is the server API for Devices service.
type DevicesServer interface {
	AddDevice(context.Context, *AddDeviceReq) (*Device, error)
	ListSpecificDeviceForUser(context.Context, *ListDevicesReq) (*ListDevicesRes, error)
	DeleteDevice(context.Context, *DeleteDeviceReq) (*emptypb.Empty, error)
	ListAllDevices(context.Context, *ListAllDevicesReq) (*ListAllDevicesRes, error)
}

// UnimplementedDevicesServer can be embedded to have forward compatible implementations.
type UnimplementedDevicesServer struct {
}

func (*UnimplementedDevicesServer) AddDevice(context.Context, *AddDeviceReq) (*Device, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDevice not implemented")
}
func (*UnimplementedDevicesServer) ListSpecificDeviceForUser(context.Context, *ListDevicesReq) (*ListDevicesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSpecificDeviceForUser not implemented")
}
func (*UnimplementedDevicesServer) DeleteDevice(context.Context, *DeleteDeviceReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDevice not implemented")
}
func (*UnimplementedDevicesServer) ListAllDevices(context.Context, *ListAllDevicesReq) (*ListAllDevicesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAllDevices not implemented")
}

func RegisterDevicesServer(s *grpc.Server, srv DevicesServer) {
	s.RegisterService(&_Devices_serviceDesc, srv)
}

func _Devices_AddDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDeviceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).AddDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Devices/AddDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).AddDevice(ctx, req.(*AddDeviceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devices_ListSpecificDeviceForUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDevicesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).ListSpecificDeviceForUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Devices/ListSpecificDeviceForUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).ListSpecificDeviceForUser(ctx, req.(*ListDevicesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devices_DeleteDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDeviceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).DeleteDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Devices/DeleteDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).DeleteDevice(ctx, req.(*DeleteDeviceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devices_ListAllDevices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAllDevicesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).ListAllDevices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Devices/ListAllDevices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).ListAllDevices(ctx, req.(*ListAllDevicesReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Devices_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Devices",
	HandlerType: (*DevicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddDevice",
			Handler:    _Devices_AddDevice_Handler,
		},
		{
			MethodName: "ListSpecificDeviceForUser",
			Handler:    _Devices_ListSpecificDeviceForUser_Handler,
		},
		{
			MethodName: "DeleteDevice",
			Handler:    _Devices_DeleteDevice_Handler,
		},
		{
			MethodName: "ListAllDevices",
			Handler:    _Devices_ListAllDevices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "devices.proto",
}