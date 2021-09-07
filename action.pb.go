// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: action.proto

package kvstore

import (
	structpb "github.com/ImSingee/structpb"
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

type Action struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Action:
	//	*Action_Set
	//	*Action_Delete
	//	*Action_Replace
	Action isAction_Action `protobuf_oneof:"action"`
}

func (x *Action) Reset() {
	*x = Action{}
	if protoimpl.UnsafeEnabled {
		mi := &file_action_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Action) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Action) ProtoMessage() {}

func (x *Action) ProtoReflect() protoreflect.Message {
	mi := &file_action_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Action.ProtoReflect.Descriptor instead.
func (*Action) Descriptor() ([]byte, []int) {
	return file_action_proto_rawDescGZIP(), []int{0}
}

func (x *Action) GetAction() isAction_Action {
	if x != nil {
		return x.Action
	}
	return nil
}

func (x *Action) GetSet() *Set {
	if x, ok := x.GetAction().(*Action_Set); ok {
		return x.Set
	}
	return nil
}

func (x *Action) GetDelete() *Delete {
	if x, ok := x.GetAction().(*Action_Delete); ok {
		return x.Delete
	}
	return nil
}

func (x *Action) GetReplace() *Replace {
	if x, ok := x.GetAction().(*Action_Replace); ok {
		return x.Replace
	}
	return nil
}

type isAction_Action interface {
	isAction_Action()
}

type Action_Set struct {
	Set *Set `protobuf:"bytes,1,opt,name=set,proto3,oneof"`
}

type Action_Delete struct {
	Delete *Delete `protobuf:"bytes,2,opt,name=delete,proto3,oneof"`
}

type Action_Replace struct {
	Replace *Replace `protobuf:"bytes,3,opt,name=replace,proto3,oneof"`
}

func (*Action_Set) isAction_Action() {}

func (*Action_Delete) isAction_Action() {}

func (*Action_Replace) isAction_Action() {}

// 设置一个值
type Set struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string          `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value *structpb.Value `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Set) Reset() {
	*x = Set{}
	if protoimpl.UnsafeEnabled {
		mi := &file_action_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Set) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Set) ProtoMessage() {}

func (x *Set) ProtoReflect() protoreflect.Message {
	mi := &file_action_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Set.ProtoReflect.Descriptor instead.
func (*Set) Descriptor() ([]byte, []int) {
	return file_action_proto_rawDescGZIP(), []int{1}
}

func (x *Set) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Set) GetValue() *structpb.Value {
	if x != nil {
		return x.Value
	}
	return nil
}

// 删除一个值
type Delete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *Delete) Reset() {
	*x = Delete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_action_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Delete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Delete) ProtoMessage() {}

func (x *Delete) ProtoReflect() protoreflect.Message {
	mi := &file_action_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Delete.ProtoReflect.Descriptor instead.
func (*Delete) Descriptor() ([]byte, []int) {
	return file_action_proto_rawDescGZIP(), []int{2}
}

func (x *Delete) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

// 完全替换底层数据
type Replace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	New *structpb.Dict `protobuf:"bytes,1,opt,name=new,proto3" json:"new,omitempty"`
}

func (x *Replace) Reset() {
	*x = Replace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_action_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Replace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Replace) ProtoMessage() {}

func (x *Replace) ProtoReflect() protoreflect.Message {
	mi := &file_action_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Replace.ProtoReflect.Descriptor instead.
func (*Replace) Descriptor() ([]byte, []int) {
	return file_action_proto_rawDescGZIP(), []int{3}
}

func (x *Replace) GetNew() *structpb.Dict {
	if x != nil {
		return x.New
	}
	return nil
}

var File_action_proto protoreflect.FileDescriptor

var file_action_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x73, 0x69, 0x6e, 0x67, 0x65, 0x65, 0x2e, 0x6b, 0x76, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x1a, 0x29,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x49, 0x6d, 0x53, 0x69, 0x6e,
	0x67, 0x65, 0x65, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x70, 0x62, 0x2f, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x01, 0x0a, 0x06, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x03, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x73, 0x69, 0x6e, 0x67, 0x65, 0x65, 0x2e, 0x6b, 0x76, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x53, 0x65, 0x74, 0x48, 0x00, 0x52, 0x03, 0x73, 0x65, 0x74, 0x12, 0x30, 0x0a,
	0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x73, 0x69, 0x6e, 0x67, 0x65, 0x65, 0x2e, 0x6b, 0x76, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x00, 0x52, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x33, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x73, 0x69, 0x6e, 0x67, 0x65, 0x65, 0x2e, 0x6b, 0x76, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x48, 0x00, 0x52, 0x07, 0x72, 0x65, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3c,
	0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x1a, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x29, 0x0a, 0x07, 0x52, 0x65, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x03, 0x6e, 0x65, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x52, 0x03,
	0x6e, 0x65, 0x77, 0x42, 0x1d, 0x5a, 0x1b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x49, 0x6d, 0x53, 0x69, 0x6e, 0x67, 0x65, 0x65, 0x2f, 0x6b, 0x76, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_action_proto_rawDescOnce sync.Once
	file_action_proto_rawDescData = file_action_proto_rawDesc
)

func file_action_proto_rawDescGZIP() []byte {
	file_action_proto_rawDescOnce.Do(func() {
		file_action_proto_rawDescData = protoimpl.X.CompressGZIP(file_action_proto_rawDescData)
	})
	return file_action_proto_rawDescData
}

var file_action_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_action_proto_goTypes = []interface{}{
	(*Action)(nil),         // 0: singee.kvstore.Action
	(*Set)(nil),            // 1: singee.kvstore.Set
	(*Delete)(nil),         // 2: singee.kvstore.Delete
	(*Replace)(nil),        // 3: singee.kvstore.Replace
	(*structpb.Value)(nil), // 4: struct.Value
	(*structpb.Dict)(nil),  // 5: struct.Dict
}
var file_action_proto_depIdxs = []int32{
	1, // 0: singee.kvstore.Action.set:type_name -> singee.kvstore.Set
	2, // 1: singee.kvstore.Action.delete:type_name -> singee.kvstore.Delete
	3, // 2: singee.kvstore.Action.replace:type_name -> singee.kvstore.Replace
	4, // 3: singee.kvstore.Set.value:type_name -> struct.Value
	5, // 4: singee.kvstore.Replace.new:type_name -> struct.Dict
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_action_proto_init() }
func file_action_proto_init() {
	if File_action_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_action_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Action); i {
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
		file_action_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Set); i {
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
		file_action_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Delete); i {
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
		file_action_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Replace); i {
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
	file_action_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Action_Set)(nil),
		(*Action_Delete)(nil),
		(*Action_Replace)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_action_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_action_proto_goTypes,
		DependencyIndexes: file_action_proto_depIdxs,
		MessageInfos:      file_action_proto_msgTypes,
	}.Build()
	File_action_proto = out.File
	file_action_proto_rawDesc = nil
	file_action_proto_goTypes = nil
	file_action_proto_depIdxs = nil
}
