// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v4.0.0
// source: proto/SmartContractService.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

var File_proto_SmartContractService_proto protoreflect.FileDescriptor

var file_proto_SmartContractService_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xb7, 0x05, 0x0a,
	0x14, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0e, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x12, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12,
	0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x30, 0x0a, 0x0f, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x38, 0x0a, 0x17, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x43, 0x61, 0x6c,
	0x6c, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x0c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x13, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x47, 0x65, 0x74, 0x42, 0x79, 0x74, 0x65, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x30, 0x0a, 0x0f, 0x67, 0x65, 0x74, 0x42, 0x79, 0x53, 0x6f, 0x6c, 0x69, 0x64, 0x69,
	0x74, 0x79, 0x49, 0x44, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x17, 0x67, 0x65, 0x74, 0x54, 0x78, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x42, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x44, 0x12, 0x0c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x0f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03, 0x88,
	0x02, 0x01, 0x12, 0x40, 0x0a, 0x0e, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0c, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x6e,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x49, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65,
	0x64, 0x65, 0x72, 0x61, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x68, 0x65, 0x64, 0x65, 0x72,
	0x61, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_proto_SmartContractService_proto_goTypes = []interface{}{
	(*Transaction)(nil),         // 0: proto.Transaction
	(*Query)(nil),               // 1: proto.Query
	(*TransactionResponse)(nil), // 2: proto.TransactionResponse
	(*Response)(nil),            // 3: proto.Response
}
var file_proto_SmartContractService_proto_depIdxs = []int32{
	0,  // 0: proto.SmartContractService.createContract:input_type -> proto.Transaction
	0,  // 1: proto.SmartContractService.updateContract:input_type -> proto.Transaction
	0,  // 2: proto.SmartContractService.contractCallMethod:input_type -> proto.Transaction
	1,  // 3: proto.SmartContractService.getContractInfo:input_type -> proto.Query
	1,  // 4: proto.SmartContractService.contractCallLocalMethod:input_type -> proto.Query
	1,  // 5: proto.SmartContractService.ContractGetBytecode:input_type -> proto.Query
	1,  // 6: proto.SmartContractService.getBySolidityID:input_type -> proto.Query
	1,  // 7: proto.SmartContractService.getTxRecordByContractID:input_type -> proto.Query
	0,  // 8: proto.SmartContractService.deleteContract:input_type -> proto.Transaction
	0,  // 9: proto.SmartContractService.systemDelete:input_type -> proto.Transaction
	0,  // 10: proto.SmartContractService.systemUndelete:input_type -> proto.Transaction
	2,  // 11: proto.SmartContractService.createContract:output_type -> proto.TransactionResponse
	2,  // 12: proto.SmartContractService.updateContract:output_type -> proto.TransactionResponse
	2,  // 13: proto.SmartContractService.contractCallMethod:output_type -> proto.TransactionResponse
	3,  // 14: proto.SmartContractService.getContractInfo:output_type -> proto.Response
	3,  // 15: proto.SmartContractService.contractCallLocalMethod:output_type -> proto.Response
	3,  // 16: proto.SmartContractService.ContractGetBytecode:output_type -> proto.Response
	3,  // 17: proto.SmartContractService.getBySolidityID:output_type -> proto.Response
	3,  // 18: proto.SmartContractService.getTxRecordByContractID:output_type -> proto.Response
	2,  // 19: proto.SmartContractService.deleteContract:output_type -> proto.TransactionResponse
	2,  // 20: proto.SmartContractService.systemDelete:output_type -> proto.TransactionResponse
	2,  // 21: proto.SmartContractService.systemUndelete:output_type -> proto.TransactionResponse
	11, // [11:22] is the sub-list for method output_type
	0,  // [0:11] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_proto_SmartContractService_proto_init() }
func file_proto_SmartContractService_proto_init() {
	if File_proto_SmartContractService_proto != nil {
		return
	}
	file_proto_TransactionResponse_proto_init()
	file_proto_Query_proto_init()
	file_proto_Response_proto_init()
	file_proto_Transaction_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_SmartContractService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_SmartContractService_proto_goTypes,
		DependencyIndexes: file_proto_SmartContractService_proto_depIdxs,
	}.Build()
	File_proto_SmartContractService_proto = out.File
	file_proto_SmartContractService_proto_rawDesc = nil
	file_proto_SmartContractService_proto_goTypes = nil
	file_proto_SmartContractService_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SmartContractServiceClient is the client API for SmartContractService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SmartContractServiceClient interface {
	// Creates a contract
	CreateContract(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	// Updates a contract with the content
	UpdateContract(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	// Calls a contract
	ContractCallMethod(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	// Retrieves the contract information
	GetContractInfo(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error)
	// Calls a smart contract to be run on a single node
	ContractCallLocalMethod(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error)
	// Retrieves the byte code of a contract
	ContractGetBytecode(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error)
	// Retrieves a contract by its Solidity address
	GetBySolidityID(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error)
	// Deprecated: Do not use.
	// Always returns an empty record list, as contract accounts are never effective payers for transactions
	GetTxRecordByContractID(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error)
	// Deletes a contract instance and transfers any remaining hbars to a specified receiver
	DeleteContract(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	// Deletes a contract if the submitting account has network admin privileges
	SystemDelete(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	// Undeletes a contract if the submitting account has network admin privileges
	SystemUndelete(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
}

type smartContractServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSmartContractServiceClient(cc grpc.ClientConnInterface) SmartContractServiceClient {
	return &smartContractServiceClient{cc}
}

func (c *smartContractServiceClient) CreateContract(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.SmartContractService/createContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smartContractServiceClient) UpdateContract(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.SmartContractService/updateContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smartContractServiceClient) ContractCallMethod(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.SmartContractService/contractCallMethod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smartContractServiceClient) GetContractInfo(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.SmartContractService/getContractInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smartContractServiceClient) ContractCallLocalMethod(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.SmartContractService/contractCallLocalMethod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smartContractServiceClient) ContractGetBytecode(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.SmartContractService/ContractGetBytecode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smartContractServiceClient) GetBySolidityID(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.SmartContractService/getBySolidityID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *smartContractServiceClient) GetTxRecordByContractID(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.SmartContractService/getTxRecordByContractID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smartContractServiceClient) DeleteContract(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.SmartContractService/deleteContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smartContractServiceClient) SystemDelete(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.SmartContractService/systemDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smartContractServiceClient) SystemUndelete(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.SmartContractService/systemUndelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SmartContractServiceServer is the server API for SmartContractService service.
type SmartContractServiceServer interface {
	// Creates a contract
	CreateContract(context.Context, *Transaction) (*TransactionResponse, error)
	// Updates a contract with the content
	UpdateContract(context.Context, *Transaction) (*TransactionResponse, error)
	// Calls a contract
	ContractCallMethod(context.Context, *Transaction) (*TransactionResponse, error)
	// Retrieves the contract information
	GetContractInfo(context.Context, *Query) (*Response, error)
	// Calls a smart contract to be run on a single node
	ContractCallLocalMethod(context.Context, *Query) (*Response, error)
	// Retrieves the byte code of a contract
	ContractGetBytecode(context.Context, *Query) (*Response, error)
	// Retrieves a contract by its Solidity address
	GetBySolidityID(context.Context, *Query) (*Response, error)
	// Deprecated: Do not use.
	// Always returns an empty record list, as contract accounts are never effective payers for transactions
	GetTxRecordByContractID(context.Context, *Query) (*Response, error)
	// Deletes a contract instance and transfers any remaining hbars to a specified receiver
	DeleteContract(context.Context, *Transaction) (*TransactionResponse, error)
	// Deletes a contract if the submitting account has network admin privileges
	SystemDelete(context.Context, *Transaction) (*TransactionResponse, error)
	// Undeletes a contract if the submitting account has network admin privileges
	SystemUndelete(context.Context, *Transaction) (*TransactionResponse, error)
}

// UnimplementedSmartContractServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSmartContractServiceServer struct {
}

func (*UnimplementedSmartContractServiceServer) CreateContract(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateContract not implemented")
}
func (*UnimplementedSmartContractServiceServer) UpdateContract(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateContract not implemented")
}
func (*UnimplementedSmartContractServiceServer) ContractCallMethod(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContractCallMethod not implemented")
}
func (*UnimplementedSmartContractServiceServer) GetContractInfo(context.Context, *Query) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContractInfo not implemented")
}
func (*UnimplementedSmartContractServiceServer) ContractCallLocalMethod(context.Context, *Query) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContractCallLocalMethod not implemented")
}
func (*UnimplementedSmartContractServiceServer) ContractGetBytecode(context.Context, *Query) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContractGetBytecode not implemented")
}
func (*UnimplementedSmartContractServiceServer) GetBySolidityID(context.Context, *Query) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBySolidityID not implemented")
}
func (*UnimplementedSmartContractServiceServer) GetTxRecordByContractID(context.Context, *Query) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTxRecordByContractID not implemented")
}
func (*UnimplementedSmartContractServiceServer) DeleteContract(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteContract not implemented")
}
func (*UnimplementedSmartContractServiceServer) SystemDelete(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemDelete not implemented")
}
func (*UnimplementedSmartContractServiceServer) SystemUndelete(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemUndelete not implemented")
}

func RegisterSmartContractServiceServer(s *grpc.Server, srv SmartContractServiceServer) {
	s.RegisterService(&_SmartContractService_serviceDesc, srv)
}

func _SmartContractService_CreateContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartContractServiceServer).CreateContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SmartContractService/CreateContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartContractServiceServer).CreateContract(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmartContractService_UpdateContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartContractServiceServer).UpdateContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SmartContractService/UpdateContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartContractServiceServer).UpdateContract(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmartContractService_ContractCallMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartContractServiceServer).ContractCallMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SmartContractService/ContractCallMethod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartContractServiceServer).ContractCallMethod(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmartContractService_GetContractInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartContractServiceServer).GetContractInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SmartContractService/GetContractInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartContractServiceServer).GetContractInfo(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmartContractService_ContractCallLocalMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartContractServiceServer).ContractCallLocalMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SmartContractService/ContractCallLocalMethod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartContractServiceServer).ContractCallLocalMethod(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmartContractService_ContractGetBytecode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartContractServiceServer).ContractGetBytecode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SmartContractService/ContractGetBytecode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartContractServiceServer).ContractGetBytecode(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmartContractService_GetBySolidityID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartContractServiceServer).GetBySolidityID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SmartContractService/GetBySolidityID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartContractServiceServer).GetBySolidityID(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmartContractService_GetTxRecordByContractID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartContractServiceServer).GetTxRecordByContractID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SmartContractService/GetTxRecordByContractID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartContractServiceServer).GetTxRecordByContractID(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmartContractService_DeleteContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartContractServiceServer).DeleteContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SmartContractService/DeleteContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartContractServiceServer).DeleteContract(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmartContractService_SystemDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartContractServiceServer).SystemDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SmartContractService/SystemDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartContractServiceServer).SystemDelete(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmartContractService_SystemUndelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartContractServiceServer).SystemUndelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SmartContractService/SystemUndelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartContractServiceServer).SystemUndelete(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

var _SmartContractService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SmartContractService",
	HandlerType: (*SmartContractServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createContract",
			Handler:    _SmartContractService_CreateContract_Handler,
		},
		{
			MethodName: "updateContract",
			Handler:    _SmartContractService_UpdateContract_Handler,
		},
		{
			MethodName: "contractCallMethod",
			Handler:    _SmartContractService_ContractCallMethod_Handler,
		},
		{
			MethodName: "getContractInfo",
			Handler:    _SmartContractService_GetContractInfo_Handler,
		},
		{
			MethodName: "contractCallLocalMethod",
			Handler:    _SmartContractService_ContractCallLocalMethod_Handler,
		},
		{
			MethodName: "ContractGetBytecode",
			Handler:    _SmartContractService_ContractGetBytecode_Handler,
		},
		{
			MethodName: "getBySolidityID",
			Handler:    _SmartContractService_GetBySolidityID_Handler,
		},
		{
			MethodName: "getTxRecordByContractID",
			Handler:    _SmartContractService_GetTxRecordByContractID_Handler,
		},
		{
			MethodName: "deleteContract",
			Handler:    _SmartContractService_DeleteContract_Handler,
		},
		{
			MethodName: "systemDelete",
			Handler:    _SmartContractService_SystemDelete_Handler,
		},
		{
			MethodName: "systemUndelete",
			Handler:    _SmartContractService_SystemUndelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/SmartContractService.proto",
}
