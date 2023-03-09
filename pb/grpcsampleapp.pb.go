// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: grpcsampleapp.proto

package pb

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcsampleapp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_grpcsampleapp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_grpcsampleapp_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Items struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *Items) Reset() {
	*x = Items{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcsampleapp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Items) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Items) ProtoMessage() {}

func (x *Items) ProtoReflect() protoreflect.Message {
	mi := &file_grpcsampleapp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Items.ProtoReflect.Descriptor instead.
func (*Items) Descriptor() ([]byte, []int) {
	return file_grpcsampleapp_proto_rawDescGZIP(), []int{1}
}

func (x *Items) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcsampleapp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_grpcsampleapp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_grpcsampleapp_proto_rawDescGZIP(), []int{2}
}

func (x *Item) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Item) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UserItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Item *Item `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *UserItem) Reset() {
	*x = UserItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcsampleapp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserItem) ProtoMessage() {}

func (x *UserItem) ProtoReflect() protoreflect.Message {
	mi := &file_grpcsampleapp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserItem.ProtoReflect.Descriptor instead.
func (*UserItem) Descriptor() ([]byte, []int) {
	return file_grpcsampleapp_proto_rawDescGZIP(), []int{3}
}

func (x *UserItem) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UserItem) GetItem() *Item {
	if x != nil {
		return x.Item
	}
	return nil
}

var File_grpcsampleapp_proto protoreflect.FileDescriptor

var file_grpcsampleapp_proto_rawDesc = []byte{
	0x0a, 0x13, 0x67, 0x72, 0x70, 0x63, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x61, 0x70, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x67, 0x72, 0x70, 0x63, 0x73, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x61, 0x70, 0x70, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x2a, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x32, 0x0a,
	0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x29, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x73, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x61, 0x70, 0x70, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x22, 0x2a, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x5c, 0x0a,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x27, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x73, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x61, 0x70, 0x70, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x12, 0x27, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x61, 0x70, 0x70,
	0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x32, 0xbb, 0x02, 0x0a, 0x04,
	0x47, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x61,
	0x70, 0x70, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x73, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x61, 0x70, 0x70, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x12, 0x3c,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x13,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x61, 0x70, 0x70, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x1a, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x61, 0x70, 0x70, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x22, 0x00, 0x30, 0x01, 0x12, 0x40, 0x0a, 0x0b,
	0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x61, 0x70, 0x70, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x74, 0x65, 0x6d, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3c,
	0x0a, 0x08, 0x50, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x09,
	0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x14, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x61, 0x70,
	0x70, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x00, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x68, 0x69, 0x6e, 0x35, 0x6f, 0x6b, 0x2f,
	0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x61, 0x70, 0x70, 0x2d,
	0x77, 0x69, 0x74, 0x68, 0x2d, 0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpcsampleapp_proto_rawDescOnce sync.Once
	file_grpcsampleapp_proto_rawDescData = file_grpcsampleapp_proto_rawDesc
)

func file_grpcsampleapp_proto_rawDescGZIP() []byte {
	file_grpcsampleapp_proto_rawDescOnce.Do(func() {
		file_grpcsampleapp_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpcsampleapp_proto_rawDescData)
	})
	return file_grpcsampleapp_proto_rawDescData
}

var file_grpcsampleapp_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_grpcsampleapp_proto_goTypes = []interface{}{
	(*User)(nil),        // 0: grpcsampleapp.User
	(*Items)(nil),       // 1: grpcsampleapp.Items
	(*Item)(nil),        // 2: grpcsampleapp.Item
	(*UserItem)(nil),    // 3: grpcsampleapp.UserItem
	(*empty.Empty)(nil), // 4: google.protobuf.Empty
}
var file_grpcsampleapp_proto_depIdxs = []int32{
	2, // 0: grpcsampleapp.Items.items:type_name -> grpcsampleapp.Item
	0, // 1: grpcsampleapp.UserItem.user:type_name -> grpcsampleapp.User
	2, // 2: grpcsampleapp.UserItem.item:type_name -> grpcsampleapp.Item
	0, // 3: grpcsampleapp.Game.CreateUser:input_type -> grpcsampleapp.User
	0, // 4: grpcsampleapp.Game.GetUserItems:input_type -> grpcsampleapp.User
	3, // 5: grpcsampleapp.Game.AddItemUser:input_type -> grpcsampleapp.UserItem
	4, // 6: grpcsampleapp.Game.PingPong:input_type -> google.protobuf.Empty
	4, // 7: grpcsampleapp.Game.ListItems:input_type -> google.protobuf.Empty
	0, // 8: grpcsampleapp.Game.CreateUser:output_type -> grpcsampleapp.User
	2, // 9: grpcsampleapp.Game.GetUserItems:output_type -> grpcsampleapp.Item
	4, // 10: grpcsampleapp.Game.AddItemUser:output_type -> google.protobuf.Empty
	4, // 11: grpcsampleapp.Game.PingPong:output_type -> google.protobuf.Empty
	1, // 12: grpcsampleapp.Game.ListItems:output_type -> grpcsampleapp.Items
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_grpcsampleapp_proto_init() }
func file_grpcsampleapp_proto_init() {
	if File_grpcsampleapp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpcsampleapp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_grpcsampleapp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Items); i {
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
		file_grpcsampleapp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
		file_grpcsampleapp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserItem); i {
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
			RawDescriptor: file_grpcsampleapp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpcsampleapp_proto_goTypes,
		DependencyIndexes: file_grpcsampleapp_proto_depIdxs,
		MessageInfos:      file_grpcsampleapp_proto_msgTypes,
	}.Build()
	File_grpcsampleapp_proto = out.File
	file_grpcsampleapp_proto_rawDesc = nil
	file_grpcsampleapp_proto_goTypes = nil
	file_grpcsampleapp_proto_depIdxs = nil
}
