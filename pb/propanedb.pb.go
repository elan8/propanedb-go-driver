// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.16.0
// source: api/propanedb.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PropaneDatabase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string                          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DescriptorSet *descriptorpb.FileDescriptorSet `protobuf:"bytes,2,opt,name=descriptor_set,json=descriptorSet,proto3" json:"descriptor_set,omitempty"`
}

func (x *PropaneDatabase) Reset() {
	*x = PropaneDatabase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_propanedb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PropaneDatabase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PropaneDatabase) ProtoMessage() {}

func (x *PropaneDatabase) ProtoReflect() protoreflect.Message {
	mi := &file_api_propanedb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PropaneDatabase.ProtoReflect.Descriptor instead.
func (*PropaneDatabase) Descriptor() ([]byte, []int) {
	return file_api_propanedb_proto_rawDescGZIP(), []int{0}
}

func (x *PropaneDatabase) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PropaneDatabase) GetDescriptorSet() *descriptorpb.FileDescriptorSet {
	if x != nil {
		return x.DescriptorSet
	}
	return nil
}

type PropaneEntities struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entities []*PropaneEntity `protobuf:"bytes,1,rep,name=entities,proto3" json:"entities,omitempty"`
}

func (x *PropaneEntities) Reset() {
	*x = PropaneEntities{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_propanedb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PropaneEntities) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PropaneEntities) ProtoMessage() {}

func (x *PropaneEntities) ProtoReflect() protoreflect.Message {
	mi := &file_api_propanedb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PropaneEntities.ProtoReflect.Descriptor instead.
func (*PropaneEntities) Descriptor() ([]byte, []int) {
	return file_api_propanedb_proto_rawDescGZIP(), []int{1}
}

func (x *PropaneEntities) GetEntities() []*PropaneEntity {
	if x != nil {
		return x.Entities
	}
	return nil
}

type PropaneEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *anypb.Any `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PropaneEntity) Reset() {
	*x = PropaneEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_propanedb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PropaneEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PropaneEntity) ProtoMessage() {}

func (x *PropaneEntity) ProtoReflect() protoreflect.Message {
	mi := &file_api_propanedb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PropaneEntity.ProtoReflect.Descriptor instead.
func (*PropaneEntity) Descriptor() ([]byte, []int) {
	return file_api_propanedb_proto_rawDescGZIP(), []int{2}
}

func (x *PropaneEntity) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

type PropaneStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusMessage string `protobuf:"bytes,1,opt,name=statusMessage,proto3" json:"statusMessage,omitempty"`
}

func (x *PropaneStatus) Reset() {
	*x = PropaneStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_propanedb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PropaneStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PropaneStatus) ProtoMessage() {}

func (x *PropaneStatus) ProtoReflect() protoreflect.Message {
	mi := &file_api_propanedb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PropaneStatus.ProtoReflect.Descriptor instead.
func (*PropaneStatus) Descriptor() ([]byte, []int) {
	return file_api_propanedb_proto_rawDescGZIP(), []int{3}
}

func (x *PropaneStatus) GetStatusMessage() string {
	if x != nil {
		return x.StatusMessage
	}
	return ""
}

type PropaneId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PropaneId) Reset() {
	*x = PropaneId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_propanedb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PropaneId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PropaneId) ProtoMessage() {}

func (x *PropaneId) ProtoReflect() protoreflect.Message {
	mi := &file_api_propanedb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PropaneId.ProtoReflect.Descriptor instead.
func (*PropaneId) Descriptor() ([]byte, []int) {
	return file_api_propanedb_proto_rawDescGZIP(), []int{4}
}

func (x *PropaneId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type PropaneSearch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityType string `protobuf:"bytes,1,opt,name=entityType,proto3" json:"entityType,omitempty"`
	Query      string `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *PropaneSearch) Reset() {
	*x = PropaneSearch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_propanedb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PropaneSearch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PropaneSearch) ProtoMessage() {}

func (x *PropaneSearch) ProtoReflect() protoreflect.Message {
	mi := &file_api_propanedb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PropaneSearch.ProtoReflect.Descriptor instead.
func (*PropaneSearch) Descriptor() ([]byte, []int) {
	return file_api_propanedb_proto_rawDescGZIP(), []int{5}
}

func (x *PropaneSearch) GetEntityType() string {
	if x != nil {
		return x.EntityType
	}
	return ""
}

func (x *PropaneSearch) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

var File_api_propanedb_proto protoreflect.FileDescriptor

var file_api_propanedb_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x70, 0x61, 0x6e, 0x65, 0x64, 0x62, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x72, 0x6f, 0x70, 0x61, 0x6e, 0x65, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x70, 0x0a, 0x0f, 0x50,
	0x72, 0x6f, 0x70, 0x61, 0x6e, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x49, 0x0a, 0x0e, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x5f, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x52, 0x0d,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x22, 0x45, 0x0a,
	0x0f, 0x50, 0x72, 0x6f, 0x70, 0x61, 0x6e, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x12, 0x32, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x61, 0x6e, 0x65, 0x2e, 0x50, 0x72, 0x6f,
	0x70, 0x61, 0x6e, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x22, 0x39, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x70, 0x61, 0x6e, 0x65, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x35, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x70, 0x61, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x24, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x1b, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x70, 0x61, 0x6e,
	0x65, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x45, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x70, 0x61, 0x6e, 0x65, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x32, 0xb0, 0x02, 0x0a, 0x08, 0x44,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x70,
	0x61, 0x6e, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x61, 0x6e, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x61, 0x6e, 0x65, 0x2e, 0x50, 0x72,
	0x6f, 0x70, 0x61, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x33, 0x0a,
	0x03, 0x50, 0x75, 0x74, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x61, 0x6e, 0x65, 0x2e, 0x50,
	0x72, 0x6f, 0x70, 0x61, 0x6e, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x70,
	0x72, 0x6f, 0x70, 0x61, 0x6e, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x61, 0x6e, 0x65, 0x49, 0x64,
	0x22, 0x00, 0x12, 0x33, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x70,
	0x61, 0x6e, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x61, 0x6e, 0x65, 0x49, 0x64, 0x1a, 0x16, 0x2e,
	0x70, 0x72, 0x6f, 0x70, 0x61, 0x6e, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x61, 0x6e, 0x65, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x61, 0x6e, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x70,
	0x61, 0x6e, 0x65, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x61, 0x6e, 0x65, 0x2e,
	0x50, 0x72, 0x6f, 0x70, 0x61, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12,
	0x3c, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x70,
	0x61, 0x6e, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x61, 0x6e, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x61, 0x6e, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x70,
	0x61, 0x6e, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22, 0x00, 0x42, 0x25, 0x0a,
	0x07, 0x70, 0x72, 0x6f, 0x70, 0x61, 0x6e, 0x65, 0x42, 0x0c, 0x50, 0x72, 0x6f, 0x70, 0x61, 0x6e,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x03, 0x70, 0x62, 0x2f, 0xa2, 0x02, 0x04,
	0x50, 0x72, 0x6f, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_propanedb_proto_rawDescOnce sync.Once
	file_api_propanedb_proto_rawDescData = file_api_propanedb_proto_rawDesc
)

func file_api_propanedb_proto_rawDescGZIP() []byte {
	file_api_propanedb_proto_rawDescOnce.Do(func() {
		file_api_propanedb_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_propanedb_proto_rawDescData)
	})
	return file_api_propanedb_proto_rawDescData
}

var file_api_propanedb_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_propanedb_proto_goTypes = []interface{}{
	(*PropaneDatabase)(nil),                // 0: propane.PropaneDatabase
	(*PropaneEntities)(nil),                // 1: propane.PropaneEntities
	(*PropaneEntity)(nil),                  // 2: propane.PropaneEntity
	(*PropaneStatus)(nil),                  // 3: propane.PropaneStatus
	(*PropaneId)(nil),                      // 4: propane.PropaneId
	(*PropaneSearch)(nil),                  // 5: propane.PropaneSearch
	(*descriptorpb.FileDescriptorSet)(nil), // 6: google.protobuf.FileDescriptorSet
	(*anypb.Any)(nil),                      // 7: google.protobuf.Any
}
var file_api_propanedb_proto_depIdxs = []int32{
	6, // 0: propane.PropaneDatabase.descriptor_set:type_name -> google.protobuf.FileDescriptorSet
	2, // 1: propane.PropaneEntities.entities:type_name -> propane.PropaneEntity
	7, // 2: propane.PropaneEntity.data:type_name -> google.protobuf.Any
	0, // 3: propane.Database.CreateDatabase:input_type -> propane.PropaneDatabase
	2, // 4: propane.Database.Put:input_type -> propane.PropaneEntity
	4, // 5: propane.Database.Get:input_type -> propane.PropaneId
	4, // 6: propane.Database.Delete:input_type -> propane.PropaneId
	5, // 7: propane.Database.Search:input_type -> propane.PropaneSearch
	3, // 8: propane.Database.CreateDatabase:output_type -> propane.PropaneStatus
	4, // 9: propane.Database.Put:output_type -> propane.PropaneId
	2, // 10: propane.Database.Get:output_type -> propane.PropaneEntity
	3, // 11: propane.Database.Delete:output_type -> propane.PropaneStatus
	1, // 12: propane.Database.Search:output_type -> propane.PropaneEntities
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_propanedb_proto_init() }
func file_api_propanedb_proto_init() {
	if File_api_propanedb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_propanedb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PropaneDatabase); i {
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
		file_api_propanedb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PropaneEntities); i {
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
		file_api_propanedb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PropaneEntity); i {
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
		file_api_propanedb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PropaneStatus); i {
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
		file_api_propanedb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PropaneId); i {
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
		file_api_propanedb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PropaneSearch); i {
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
			RawDescriptor: file_api_propanedb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_propanedb_proto_goTypes,
		DependencyIndexes: file_api_propanedb_proto_depIdxs,
		MessageInfos:      file_api_propanedb_proto_msgTypes,
	}.Build()
	File_api_propanedb_proto = out.File
	file_api_propanedb_proto_rawDesc = nil
	file_api_propanedb_proto_goTypes = nil
	file_api_propanedb_proto_depIdxs = nil
}