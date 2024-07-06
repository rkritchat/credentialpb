// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: credential.proto

package credentialpb

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

type CreateCredentialReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *CreateCredentialReq) Reset() {
	*x = CreateCredentialReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_credential_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCredentialReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCredentialReq) ProtoMessage() {}

func (x *CreateCredentialReq) ProtoReflect() protoreflect.Message {
	mi := &file_credential_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCredentialReq.ProtoReflect.Descriptor instead.
func (*CreateCredentialReq) Descriptor() ([]byte, []int) {
	return file_credential_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCredentialReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateCredentialReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type CommonResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSucces bool   `protobuf:"varint,1,opt,name=isSucces,proto3" json:"isSucces,omitempty"`
	Message  string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CommonResp) Reset() {
	*x = CommonResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_credential_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonResp) ProtoMessage() {}

func (x *CommonResp) ProtoReflect() protoreflect.Message {
	mi := &file_credential_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonResp.ProtoReflect.Descriptor instead.
func (*CommonResp) Descriptor() ([]byte, []int) {
	return file_credential_proto_rawDescGZIP(), []int{1}
}

func (x *CommonResp) GetIsSucces() bool {
	if x != nil {
		return x.IsSucces
	}
	return false
}

func (x *CommonResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetCredentialReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *GetCredentialReq) Reset() {
	*x = GetCredentialReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_credential_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCredentialReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCredentialReq) ProtoMessage() {}

func (x *GetCredentialReq) ProtoReflect() protoreflect.Message {
	mi := &file_credential_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCredentialReq.ProtoReflect.Descriptor instead.
func (*GetCredentialReq) Descriptor() ([]byte, []int) {
	return file_credential_proto_rawDescGZIP(), []int{2}
}

func (x *GetCredentialReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type GetCredentialResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *GetCredentialResp) Reset() {
	*x = GetCredentialResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_credential_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCredentialResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCredentialResp) ProtoMessage() {}

func (x *GetCredentialResp) ProtoReflect() protoreflect.Message {
	mi := &file_credential_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCredentialResp.ProtoReflect.Descriptor instead.
func (*GetCredentialResp) Descriptor() ([]byte, []int) {
	return file_credential_proto_rawDescGZIP(), []int{3}
}

func (x *GetCredentialResp) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetCredentialResp) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

var File_credential_proto protoreflect.FileDescriptor

var file_credential_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x47, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x42, 0x0a, 0x0a, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x28, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x45, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x32, 0x8d, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x14, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x71,
	0x1a, 0x0b, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12,
	0x3f, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x11, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00,
	0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72,
	0x6b, 0x72, 0x69, 0x74, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_credential_proto_rawDescOnce sync.Once
	file_credential_proto_rawDescData = file_credential_proto_rawDesc
)

func file_credential_proto_rawDescGZIP() []byte {
	file_credential_proto_rawDescOnce.Do(func() {
		file_credential_proto_rawDescData = protoimpl.X.CompressGZIP(file_credential_proto_rawDescData)
	})
	return file_credential_proto_rawDescData
}

var file_credential_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_credential_proto_goTypes = []any{
	(*CreateCredentialReq)(nil), // 0: CreateCredentialReq
	(*CommonResp)(nil),          // 1: CommonResp
	(*GetCredentialReq)(nil),    // 2: GetCredentialReq
	(*GetCredentialResp)(nil),   // 3: GetCredentialResp
}
var file_credential_proto_depIdxs = []int32{
	0, // 0: CredentialService.CreateCredential:input_type -> CreateCredentialReq
	2, // 1: CredentialService.GetCredentialByEmail:input_type -> GetCredentialReq
	1, // 2: CredentialService.CreateCredential:output_type -> CommonResp
	3, // 3: CredentialService.GetCredentialByEmail:output_type -> GetCredentialResp
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_credential_proto_init() }
func file_credential_proto_init() {
	if File_credential_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_credential_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateCredentialReq); i {
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
		file_credential_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CommonResp); i {
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
		file_credential_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetCredentialReq); i {
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
		file_credential_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetCredentialResp); i {
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
			RawDescriptor: file_credential_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_credential_proto_goTypes,
		DependencyIndexes: file_credential_proto_depIdxs,
		MessageInfos:      file_credential_proto_msgTypes,
	}.Build()
	File_credential_proto = out.File
	file_credential_proto_rawDesc = nil
	file_credential_proto_goTypes = nil
	file_credential_proto_depIdxs = nil
}