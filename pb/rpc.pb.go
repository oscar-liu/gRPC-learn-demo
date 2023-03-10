//proto3标准

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.5
// source: rpc.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// 请求数据体
type RequestMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RequestMessage) Reset() {
	*x = RequestMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestMessage) ProtoMessage() {}

func (x *RequestMessage) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestMessage.ProtoReflect.Descriptor instead.
func (*RequestMessage) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *RequestMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RequestMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// 返回数据体
type ResponseMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ResponseMessage) Reset() {
	*x = ResponseMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseMessage) ProtoMessage() {}

func (x *ResponseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseMessage.ProtoReflect.Descriptor instead.
func (*ResponseMessage) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResponseMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_rpc_proto protoreflect.FileDescriptor

var file_rpc_proto_rawDesc = []byte{
	0x0a, 0x09, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3e, 0x0a, 0x0e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3f, 0x0a, 0x0f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x84, 0x02, 0x0a,
	0x06, 0x47, 0x72, 0x65, 0x65, 0x74, 0x73, 0x12, 0x34, 0x0a, 0x0d, 0x73, 0x61, 0x79, 0x55, 0x6e,
	0x61, 0x72, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x0f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x10, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a,
	0x14, 0x73, 0x61, 0x79, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x0f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x10, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x3d, 0x0a, 0x14,
	0x53, 0x61, 0x79, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x0f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x10, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x46, 0x0a, 0x1b, 0x53,
	0x61, 0x79, 0x42, 0x69, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x0f, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x10, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x28,
	0x01, 0x30, 0x01, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_proto_rawDescOnce sync.Once
	file_rpc_proto_rawDescData = file_rpc_proto_rawDesc
)

func file_rpc_proto_rawDescGZIP() []byte {
	file_rpc_proto_rawDescOnce.Do(func() {
		file_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_proto_rawDescData)
	})
	return file_rpc_proto_rawDescData
}

var file_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_proto_goTypes = []interface{}{
	(*RequestMessage)(nil),  // 0: RequestMessage
	(*ResponseMessage)(nil), // 1: ResponseMessage
}
var file_rpc_proto_depIdxs = []int32{
	0, // 0: Greets.sayUnaryHello:input_type -> RequestMessage
	0, // 1: Greets.sayServerStreamHello:input_type -> RequestMessage
	0, // 2: Greets.SayClientStreamHello:input_type -> RequestMessage
	0, // 3: Greets.SayBiDirectionalStreamHello:input_type -> RequestMessage
	1, // 4: Greets.sayUnaryHello:output_type -> ResponseMessage
	1, // 5: Greets.sayServerStreamHello:output_type -> ResponseMessage
	1, // 6: Greets.SayClientStreamHello:output_type -> ResponseMessage
	1, // 7: Greets.SayBiDirectionalStreamHello:output_type -> ResponseMessage
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpc_proto_init() }
func file_rpc_proto_init() {
	if File_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestMessage); i {
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
		file_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseMessage); i {
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
			RawDescriptor: file_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_proto_goTypes,
		DependencyIndexes: file_rpc_proto_depIdxs,
		MessageInfos:      file_rpc_proto_msgTypes,
	}.Build()
	File_rpc_proto = out.File
	file_rpc_proto_rawDesc = nil
	file_rpc_proto_goTypes = nil
	file_rpc_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GreetsClient is the client API for Greets service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreetsClient interface {
	// 简单 RPC（Unary RPC）
	SayUnaryHello(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*ResponseMessage, error)
	// 服务端流式 RPC（Server streaming RPC）
	SayServerStreamHello(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (Greets_SayServerStreamHelloClient, error)
	// 客户端流式 RPC（Client streaming RPC）
	SayClientStreamHello(ctx context.Context, opts ...grpc.CallOption) (Greets_SayClientStreamHelloClient, error)
	// 双向流式 RPC（Bi-directional streaming RPC）
	SayBiDirectionalStreamHello(ctx context.Context, opts ...grpc.CallOption) (Greets_SayBiDirectionalStreamHelloClient, error)
}

type greetsClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetsClient(cc grpc.ClientConnInterface) GreetsClient {
	return &greetsClient{cc}
}

func (c *greetsClient) SayUnaryHello(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*ResponseMessage, error) {
	out := new(ResponseMessage)
	err := c.cc.Invoke(ctx, "/Greets/sayUnaryHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greetsClient) SayServerStreamHello(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (Greets_SayServerStreamHelloClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Greets_serviceDesc.Streams[0], "/Greets/sayServerStreamHello", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetsSayServerStreamHelloClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Greets_SayServerStreamHelloClient interface {
	Recv() (*ResponseMessage, error)
	grpc.ClientStream
}

type greetsSayServerStreamHelloClient struct {
	grpc.ClientStream
}

func (x *greetsSayServerStreamHelloClient) Recv() (*ResponseMessage, error) {
	m := new(ResponseMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetsClient) SayClientStreamHello(ctx context.Context, opts ...grpc.CallOption) (Greets_SayClientStreamHelloClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Greets_serviceDesc.Streams[1], "/Greets/SayClientStreamHello", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetsSayClientStreamHelloClient{stream}
	return x, nil
}

type Greets_SayClientStreamHelloClient interface {
	Send(*RequestMessage) error
	CloseAndRecv() (*ResponseMessage, error)
	grpc.ClientStream
}

type greetsSayClientStreamHelloClient struct {
	grpc.ClientStream
}

func (x *greetsSayClientStreamHelloClient) Send(m *RequestMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetsSayClientStreamHelloClient) CloseAndRecv() (*ResponseMessage, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ResponseMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetsClient) SayBiDirectionalStreamHello(ctx context.Context, opts ...grpc.CallOption) (Greets_SayBiDirectionalStreamHelloClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Greets_serviceDesc.Streams[2], "/Greets/SayBiDirectionalStreamHello", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetsSayBiDirectionalStreamHelloClient{stream}
	return x, nil
}

type Greets_SayBiDirectionalStreamHelloClient interface {
	Send(*RequestMessage) error
	Recv() (*ResponseMessage, error)
	grpc.ClientStream
}

type greetsSayBiDirectionalStreamHelloClient struct {
	grpc.ClientStream
}

func (x *greetsSayBiDirectionalStreamHelloClient) Send(m *RequestMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetsSayBiDirectionalStreamHelloClient) Recv() (*ResponseMessage, error) {
	m := new(ResponseMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetsServer is the server API for Greets service.
type GreetsServer interface {
	// 简单 RPC（Unary RPC）
	SayUnaryHello(context.Context, *RequestMessage) (*ResponseMessage, error)
	// 服务端流式 RPC（Server streaming RPC）
	SayServerStreamHello(*RequestMessage, Greets_SayServerStreamHelloServer) error
	// 客户端流式 RPC（Client streaming RPC）
	SayClientStreamHello(Greets_SayClientStreamHelloServer) error
	// 双向流式 RPC（Bi-directional streaming RPC）
	SayBiDirectionalStreamHello(Greets_SayBiDirectionalStreamHelloServer) error
}

// UnimplementedGreetsServer can be embedded to have forward compatible implementations.
type UnimplementedGreetsServer struct {
}

func (*UnimplementedGreetsServer) SayUnaryHello(context.Context, *RequestMessage) (*ResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayUnaryHello not implemented")
}
func (*UnimplementedGreetsServer) SayServerStreamHello(*RequestMessage, Greets_SayServerStreamHelloServer) error {
	return status.Errorf(codes.Unimplemented, "method SayServerStreamHello not implemented")
}
func (*UnimplementedGreetsServer) SayClientStreamHello(Greets_SayClientStreamHelloServer) error {
	return status.Errorf(codes.Unimplemented, "method SayClientStreamHello not implemented")
}
func (*UnimplementedGreetsServer) SayBiDirectionalStreamHello(Greets_SayBiDirectionalStreamHelloServer) error {
	return status.Errorf(codes.Unimplemented, "method SayBiDirectionalStreamHello not implemented")
}

func RegisterGreetsServer(s *grpc.Server, srv GreetsServer) {
	s.RegisterService(&_Greets_serviceDesc, srv)
}

func _Greets_SayUnaryHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetsServer).SayUnaryHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Greets/SayUnaryHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetsServer).SayUnaryHello(ctx, req.(*RequestMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greets_SayServerStreamHello_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RequestMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreetsServer).SayServerStreamHello(m, &greetsSayServerStreamHelloServer{stream})
}

type Greets_SayServerStreamHelloServer interface {
	Send(*ResponseMessage) error
	grpc.ServerStream
}

type greetsSayServerStreamHelloServer struct {
	grpc.ServerStream
}

func (x *greetsSayServerStreamHelloServer) Send(m *ResponseMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _Greets_SayClientStreamHello_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetsServer).SayClientStreamHello(&greetsSayClientStreamHelloServer{stream})
}

type Greets_SayClientStreamHelloServer interface {
	SendAndClose(*ResponseMessage) error
	Recv() (*RequestMessage, error)
	grpc.ServerStream
}

type greetsSayClientStreamHelloServer struct {
	grpc.ServerStream
}

func (x *greetsSayClientStreamHelloServer) SendAndClose(m *ResponseMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetsSayClientStreamHelloServer) Recv() (*RequestMessage, error) {
	m := new(RequestMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Greets_SayBiDirectionalStreamHello_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetsServer).SayBiDirectionalStreamHello(&greetsSayBiDirectionalStreamHelloServer{stream})
}

type Greets_SayBiDirectionalStreamHelloServer interface {
	Send(*ResponseMessage) error
	Recv() (*RequestMessage, error)
	grpc.ServerStream
}

type greetsSayBiDirectionalStreamHelloServer struct {
	grpc.ServerStream
}

func (x *greetsSayBiDirectionalStreamHelloServer) Send(m *ResponseMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetsSayBiDirectionalStreamHelloServer) Recv() (*RequestMessage, error) {
	m := new(RequestMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Greets_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Greets",
	HandlerType: (*GreetsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "sayUnaryHello",
			Handler:    _Greets_SayUnaryHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "sayServerStreamHello",
			Handler:       _Greets_SayServerStreamHello_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SayClientStreamHello",
			Handler:       _Greets_SayClientStreamHello_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "SayBiDirectionalStreamHello",
			Handler:       _Greets_SayBiDirectionalStreamHello_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "rpc.proto",
}
