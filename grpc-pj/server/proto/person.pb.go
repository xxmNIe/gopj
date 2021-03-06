// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.15.8
// source: proto/person.proto

package proto

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

type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age  int64  `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_person_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_proto_person_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_proto_person_proto_rawDescGZIP(), []int{0}
}

func (x *Person) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Person) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Person) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

type SavePersonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Person *Person `protobuf:"bytes,1,opt,name=person,proto3" json:"person,omitempty"`
}

func (x *SavePersonRequest) Reset() {
	*x = SavePersonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_person_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SavePersonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SavePersonRequest) ProtoMessage() {}

func (x *SavePersonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_person_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SavePersonRequest.ProtoReflect.Descriptor instead.
func (*SavePersonRequest) Descriptor() ([]byte, []int) {
	return file_proto_person_proto_rawDescGZIP(), []int{1}
}

func (x *SavePersonRequest) GetPerson() *Person {
	if x != nil {
		return x.Person
	}
	return nil
}

type SavePersonReplay struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SavePersonReplay) Reset() {
	*x = SavePersonReplay{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_person_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SavePersonReplay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SavePersonReplay) ProtoMessage() {}

func (x *SavePersonReplay) ProtoReflect() protoreflect.Message {
	mi := &file_proto_person_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SavePersonReplay.ProtoReflect.Descriptor instead.
func (*SavePersonReplay) Descriptor() ([]byte, []int) {
	return file_proto_person_proto_rawDescGZIP(), []int{2}
}

func (x *SavePersonReplay) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type SavePersonListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Peoples []*Person `protobuf:"bytes,1,rep,name=peoples,proto3" json:"peoples,omitempty"`
}

func (x *SavePersonListRequest) Reset() {
	*x = SavePersonListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_person_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SavePersonListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SavePersonListRequest) ProtoMessage() {}

func (x *SavePersonListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_person_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SavePersonListRequest.ProtoReflect.Descriptor instead.
func (*SavePersonListRequest) Descriptor() ([]byte, []int) {
	return file_proto_person_proto_rawDescGZIP(), []int{3}
}

func (x *SavePersonListRequest) GetPeoples() []*Person {
	if x != nil {
		return x.Peoples
	}
	return nil
}

type SavePersonListReplay struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *SavePersonListReplay) Reset() {
	*x = SavePersonListReplay{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_person_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SavePersonListReplay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SavePersonListReplay) ProtoMessage() {}

func (x *SavePersonListReplay) ProtoReflect() protoreflect.Message {
	mi := &file_proto_person_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SavePersonListReplay.ProtoReflect.Descriptor instead.
func (*SavePersonListReplay) Descriptor() ([]byte, []int) {
	return file_proto_person_proto_rawDescGZIP(), []int{4}
}

func (x *SavePersonListReplay) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_proto_person_proto protoreflect.FileDescriptor

var file_proto_person_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3e, 0x0a, 0x06, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x61, 0x67, 0x65, 0x22, 0x3a, 0x0a, 0x11, 0x53,
	0x61, 0x76, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x25, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52,
	0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x22, 0x2c, 0x0a, 0x10, 0x53, 0x61, 0x76, 0x65, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x40, 0x0a, 0x15, 0x53, 0x61, 0x76, 0x65, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27,
	0x0a, 0x07, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x07,
	0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x73, 0x22, 0x2c, 0x0a, 0x14, 0x53, 0x61, 0x76, 0x65, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x97, 0x01, 0x0a, 0x06, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65,
	0x12, 0x41, 0x0a, 0x0a, 0x53, 0x61, 0x76, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x18,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x61, 0x76, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x61,
	0x79, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0b, 0x53, 0x61, 0x76, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x73, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x22, 0x00, 0x42,
	0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_person_proto_rawDescOnce sync.Once
	file_proto_person_proto_rawDescData = file_proto_person_proto_rawDesc
)

func file_proto_person_proto_rawDescGZIP() []byte {
	file_proto_person_proto_rawDescOnce.Do(func() {
		file_proto_person_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_person_proto_rawDescData)
	})
	return file_proto_person_proto_rawDescData
}

var file_proto_person_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_person_proto_goTypes = []interface{}{
	(*Person)(nil),                // 0: proto.Person
	(*SavePersonRequest)(nil),     // 1: proto.SavePersonRequest
	(*SavePersonReplay)(nil),      // 2: proto.SavePersonReplay
	(*SavePersonListRequest)(nil), // 3: proto.SavePersonListRequest
	(*SavePersonListReplay)(nil),  // 4: proto.SavePersonListReplay
}
var file_proto_person_proto_depIdxs = []int32{
	0, // 0: proto.SavePersonRequest.person:type_name -> proto.Person
	0, // 1: proto.SavePersonListRequest.peoples:type_name -> proto.Person
	1, // 2: proto.People.SavePerson:input_type -> proto.SavePersonRequest
	3, // 3: proto.People.SavePersons:input_type -> proto.SavePersonListRequest
	2, // 4: proto.People.SavePerson:output_type -> proto.SavePersonReplay
	4, // 5: proto.People.SavePersons:output_type -> proto.SavePersonListReplay
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_person_proto_init() }
func file_proto_person_proto_init() {
	if File_proto_person_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_person_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person); i {
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
		file_proto_person_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SavePersonRequest); i {
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
		file_proto_person_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SavePersonReplay); i {
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
		file_proto_person_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SavePersonListRequest); i {
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
		file_proto_person_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SavePersonListReplay); i {
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
			RawDescriptor: file_proto_person_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_person_proto_goTypes,
		DependencyIndexes: file_proto_person_proto_depIdxs,
		MessageInfos:      file_proto_person_proto_msgTypes,
	}.Build()
	File_proto_person_proto = out.File
	file_proto_person_proto_rawDesc = nil
	file_proto_person_proto_goTypes = nil
	file_proto_person_proto_depIdxs = nil
}
