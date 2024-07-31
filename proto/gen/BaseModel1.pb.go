// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: proto/BaseModel1.proto

// This is the example of the proto file for the GRPC endpoints that will be exposed by this micro service
// The endpoints and messages are defined here, if another service needs to call any of these GRPC endpoints
// it will need to import this proto file
//The logic of theses endpoints will be implemented in the grpcHandler folder, for this example in the BaseModel1GRPCService class

package gen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ExampleCallRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	OtherField string `protobuf:"bytes,2,opt,name=otherField,proto3" json:"otherField,omitempty"`
}

func (x *ExampleCallRequest) Reset() {
	*x = ExampleCallRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_BaseModel1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExampleCallRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleCallRequest) ProtoMessage() {}

func (x *ExampleCallRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_BaseModel1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleCallRequest.ProtoReflect.Descriptor instead.
func (*ExampleCallRequest) Descriptor() ([]byte, []int) {
	return file_proto_BaseModel1_proto_rawDescGZIP(), []int{0}
}

func (x *ExampleCallRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ExampleCallRequest) GetOtherField() string {
	if x != nil {
		return x.OtherField
	}
	return ""
}

type ExampleCallResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionsFound bool  `protobuf:"varint,1,opt,name=transactionsFound,proto3" json:"transactionsFound,omitempty"`
	TransactionsCount int64 `protobuf:"varint,2,opt,name=transactionsCount,proto3" json:"transactionsCount,omitempty"`
}

func (x *ExampleCallResponse) Reset() {
	*x = ExampleCallResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_BaseModel1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExampleCallResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleCallResponse) ProtoMessage() {}

func (x *ExampleCallResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_BaseModel1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleCallResponse.ProtoReflect.Descriptor instead.
func (*ExampleCallResponse) Descriptor() ([]byte, []int) {
	return file_proto_BaseModel1_proto_rawDescGZIP(), []int{1}
}

func (x *ExampleCallResponse) GetTransactionsFound() bool {
	if x != nil {
		return x.TransactionsFound
	}
	return false
}

func (x *ExampleCallResponse) GetTransactionsCount() int64 {
	if x != nil {
		return x.TransactionsCount
	}
	return 0
}

type ExampleCallReturnsEmptyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListOfObjectExample []*ObjectExample `protobuf:"bytes,1,rep,name=listOfObjectExample,proto3" json:"listOfObjectExample,omitempty"`
}

func (x *ExampleCallReturnsEmptyRequest) Reset() {
	*x = ExampleCallReturnsEmptyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_BaseModel1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExampleCallReturnsEmptyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleCallReturnsEmptyRequest) ProtoMessage() {}

func (x *ExampleCallReturnsEmptyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_BaseModel1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleCallReturnsEmptyRequest.ProtoReflect.Descriptor instead.
func (*ExampleCallReturnsEmptyRequest) Descriptor() ([]byte, []int) {
	return file_proto_BaseModel1_proto_rawDescGZIP(), []int{2}
}

func (x *ExampleCallReturnsEmptyRequest) GetListOfObjectExample() []*ObjectExample {
	if x != nil {
		return x.ListOfObjectExample
	}
	return nil
}

type ObjectExample struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId          string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	ReportId        string `protobuf:"bytes,2,opt,name=reportId,proto3" json:"reportId,omitempty"`
	TransactionId   string `protobuf:"bytes,3,opt,name=transactionId,proto3" json:"transactionId,omitempty"`
	JsonData        string `protobuf:"bytes,4,opt,name=jsonData,proto3" json:"jsonData,omitempty"`
	TransactionDate string `protobuf:"bytes,5,opt,name=transactionDate,proto3" json:"transactionDate,omitempty"`
}

func (x *ObjectExample) Reset() {
	*x = ObjectExample{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_BaseModel1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectExample) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectExample) ProtoMessage() {}

func (x *ObjectExample) ProtoReflect() protoreflect.Message {
	mi := &file_proto_BaseModel1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectExample.ProtoReflect.Descriptor instead.
func (*ObjectExample) Descriptor() ([]byte, []int) {
	return file_proto_BaseModel1_proto_rawDescGZIP(), []int{3}
}

func (x *ObjectExample) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ObjectExample) GetReportId() string {
	if x != nil {
		return x.ReportId
	}
	return ""
}

func (x *ObjectExample) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *ObjectExample) GetJsonData() string {
	if x != nil {
		return x.JsonData
	}
	return ""
}

func (x *ObjectExample) GetTransactionDate() string {
	if x != nil {
		return x.TransactionDate
	}
	return ""
}

var File_proto_BaseModel1_proto protoreflect.FileDescriptor

var file_proto_BaseModel1_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x42, 0x61, 0x73, 0x65, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x67, 0x6f, 0x42, 0x61, 0x73, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x12, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x43,
	0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x22, 0x71, 0x0a, 0x13, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x43, 0x61, 0x6c,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x70, 0x0a, 0x1e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4e, 0x0a, 0x13, 0x6c, 0x69, 0x73, 0x74, 0x4f,
	0x66, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x42, 0x61, 0x73, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x45, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x52, 0x13, 0x6c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x22, 0xaf, 0x01, 0x0a, 0x0d, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x12, 0x24, 0x0a,
	0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6a, 0x73, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6a, 0x73, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x28, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61,
	0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x32, 0xd3, 0x01, 0x0a, 0x15, 0x42, 0x61,
	0x73, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x31, 0x47, 0x52, 0x50, 0x43, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x0b, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x43, 0x61,
	0x6c, 0x6c, 0x12, 0x21, 0x2e, 0x67, 0x6f, 0x42, 0x61, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x67, 0x6f, 0x42, 0x61, 0x73, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x43, 0x61, 0x6c,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x62, 0x0a, 0x17, 0x45,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x73, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x2d, 0x2e, 0x67, 0x6f, 0x42, 0x61, 0x73, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x43, 0x61,
	0x6c, 0x6c, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42,
	0x10, 0x5a, 0x0e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x3b, 0x67, 0x65,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_BaseModel1_proto_rawDescOnce sync.Once
	file_proto_BaseModel1_proto_rawDescData = file_proto_BaseModel1_proto_rawDesc
)

func file_proto_BaseModel1_proto_rawDescGZIP() []byte {
	file_proto_BaseModel1_proto_rawDescOnce.Do(func() {
		file_proto_BaseModel1_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_BaseModel1_proto_rawDescData)
	})
	return file_proto_BaseModel1_proto_rawDescData
}

var file_proto_BaseModel1_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_BaseModel1_proto_goTypes = []any{
	(*ExampleCallRequest)(nil),             // 0: goBaseService.ExampleCallRequest
	(*ExampleCallResponse)(nil),            // 1: goBaseService.ExampleCallResponse
	(*ExampleCallReturnsEmptyRequest)(nil), // 2: goBaseService.ExampleCallReturnsEmptyRequest
	(*ObjectExample)(nil),                  // 3: goBaseService.ObjectExample
	(*emptypb.Empty)(nil),                  // 4: google.protobuf.Empty
}
var file_proto_BaseModel1_proto_depIdxs = []int32{
	3, // 0: goBaseService.ExampleCallReturnsEmptyRequest.listOfObjectExample:type_name -> goBaseService.ObjectExample
	0, // 1: goBaseService.BaseModel1GRPCService.ExampleCall:input_type -> goBaseService.ExampleCallRequest
	2, // 2: goBaseService.BaseModel1GRPCService.ExampleCallReturnsEmpty:input_type -> goBaseService.ExampleCallReturnsEmptyRequest
	1, // 3: goBaseService.BaseModel1GRPCService.ExampleCall:output_type -> goBaseService.ExampleCallResponse
	4, // 4: goBaseService.BaseModel1GRPCService.ExampleCallReturnsEmpty:output_type -> google.protobuf.Empty
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_BaseModel1_proto_init() }
func file_proto_BaseModel1_proto_init() {
	if File_proto_BaseModel1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_BaseModel1_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ExampleCallRequest); i {
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
		file_proto_BaseModel1_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ExampleCallResponse); i {
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
		file_proto_BaseModel1_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ExampleCallReturnsEmptyRequest); i {
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
		file_proto_BaseModel1_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ObjectExample); i {
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
			RawDescriptor: file_proto_BaseModel1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_BaseModel1_proto_goTypes,
		DependencyIndexes: file_proto_BaseModel1_proto_depIdxs,
		MessageInfos:      file_proto_BaseModel1_proto_msgTypes,
	}.Build()
	File_proto_BaseModel1_proto = out.File
	file_proto_BaseModel1_proto_rawDesc = nil
	file_proto_BaseModel1_proto_goTypes = nil
	file_proto_BaseModel1_proto_depIdxs = nil
}
