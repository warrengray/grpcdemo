// Code generated by protoc-gen-go. DO NOT EDIT.
// source: demo.proto

package rpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type HelloWorldRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	NickName             string   `protobuf:"bytes,2,opt,name=NickName,proto3" json:"NickName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloWorldRequest) Reset()         { *m = HelloWorldRequest{} }
func (m *HelloWorldRequest) String() string { return proto.CompactTextString(m) }
func (*HelloWorldRequest) ProtoMessage()    {}
func (*HelloWorldRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{0}
}

func (m *HelloWorldRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloWorldRequest.Unmarshal(m, b)
}
func (m *HelloWorldRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloWorldRequest.Marshal(b, m, deterministic)
}
func (m *HelloWorldRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloWorldRequest.Merge(m, src)
}
func (m *HelloWorldRequest) XXX_Size() int {
	return xxx_messageInfo_HelloWorldRequest.Size(m)
}
func (m *HelloWorldRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloWorldRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloWorldRequest proto.InternalMessageInfo

func (m *HelloWorldRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HelloWorldRequest) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

type HelloWorldResponse struct {
	Greeting             string   `protobuf:"bytes,1,opt,name=Greeting,proto3" json:"Greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloWorldResponse) Reset()         { *m = HelloWorldResponse{} }
func (m *HelloWorldResponse) String() string { return proto.CompactTextString(m) }
func (*HelloWorldResponse) ProtoMessage()    {}
func (*HelloWorldResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{1}
}

func (m *HelloWorldResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloWorldResponse.Unmarshal(m, b)
}
func (m *HelloWorldResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloWorldResponse.Marshal(b, m, deterministic)
}
func (m *HelloWorldResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloWorldResponse.Merge(m, src)
}
func (m *HelloWorldResponse) XXX_Size() int {
	return xxx_messageInfo_HelloWorldResponse.Size(m)
}
func (m *HelloWorldResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloWorldResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloWorldResponse proto.InternalMessageInfo

func (m *HelloWorldResponse) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

type LetterResponse struct {
	Letter               string   `protobuf:"bytes,1,opt,name=Letter,proto3" json:"Letter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LetterResponse) Reset()         { *m = LetterResponse{} }
func (m *LetterResponse) String() string { return proto.CompactTextString(m) }
func (*LetterResponse) ProtoMessage()    {}
func (*LetterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{2}
}

func (m *LetterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LetterResponse.Unmarshal(m, b)
}
func (m *LetterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LetterResponse.Marshal(b, m, deterministic)
}
func (m *LetterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LetterResponse.Merge(m, src)
}
func (m *LetterResponse) XXX_Size() int {
	return xxx_messageInfo_LetterResponse.Size(m)
}
func (m *LetterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LetterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LetterResponse proto.InternalMessageInfo

func (m *LetterResponse) GetLetter() string {
	if m != nil {
		return m.Letter
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloWorldRequest)(nil), "rpc.HelloWorldRequest")
	proto.RegisterType((*HelloWorldResponse)(nil), "rpc.HelloWorldResponse")
	proto.RegisterType((*LetterResponse)(nil), "rpc.LetterResponse")
}

func init() { proto.RegisterFile("demo.proto", fileDescriptor_ca53982754088a9d) }

var fileDescriptor_ca53982754088a9d = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x49, 0xcd, 0xcd,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0x2a, 0x48, 0x56, 0x72, 0xe6, 0x12, 0xf4,
	0x48, 0xcd, 0xc9, 0xc9, 0x0f, 0xcf, 0x2f, 0xca, 0x49, 0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e,
	0x11, 0x12, 0xe2, 0x62, 0xf1, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02,
	0xb3, 0x85, 0xa4, 0xb8, 0x38, 0xfc, 0x32, 0x93, 0xb3, 0xc1, 0xe2, 0x4c, 0x60, 0x71, 0x38, 0x5f,
	0xc9, 0x80, 0x4b, 0x08, 0xd9, 0x90, 0xe2, 0x82, 0xfc, 0xbc, 0x62, 0xb0, 0x0e, 0xf7, 0xa2, 0xd4,
	0xd4, 0x92, 0xcc, 0xbc, 0x74, 0xa8, 0x49, 0x70, 0xbe, 0x92, 0x06, 0x17, 0x9f, 0x4f, 0x6a, 0x49,
	0x49, 0x6a, 0x11, 0x5c, 0xb5, 0x18, 0x17, 0x1b, 0x44, 0x04, 0xaa, 0x16, 0xca, 0x33, 0xea, 0x62,
	0xe4, 0xe2, 0x76, 0x49, 0xcd, 0xcd, 0x0f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x15, 0xb2, 0xe5,
	0xe2, 0x42, 0xd8, 0x25, 0x24, 0xa6, 0x57, 0x54, 0x90, 0xac, 0x87, 0xe1, 0x03, 0x29, 0x71, 0x0c,
	0x71, 0xa8, 0x35, 0x36, 0x5c, 0xdc, 0xc1, 0x05, 0xa9, 0x39, 0x39, 0xbe, 0x95, 0x60, 0x5f, 0xe1,
	0xd2, 0x2f, 0x0c, 0x16, 0x47, 0x75, 0xa2, 0x01, 0x63, 0x12, 0x1b, 0x38, 0xe4, 0x8c, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x5c, 0xcd, 0x23, 0x64, 0x47, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DemoServiceClient is the client API for DemoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DemoServiceClient interface {
	HelloWorld(ctx context.Context, in *HelloWorldRequest, opts ...grpc.CallOption) (*HelloWorldResponse, error)
	SpellMyName(ctx context.Context, in *HelloWorldRequest, opts ...grpc.CallOption) (DemoService_SpellMyNameClient, error)
}

type demoServiceClient struct {
	cc *grpc.ClientConn
}

func NewDemoServiceClient(cc *grpc.ClientConn) DemoServiceClient {
	return &demoServiceClient{cc}
}

func (c *demoServiceClient) HelloWorld(ctx context.Context, in *HelloWorldRequest, opts ...grpc.CallOption) (*HelloWorldResponse, error) {
	out := new(HelloWorldResponse)
	err := c.cc.Invoke(ctx, "/rpc.DemoService/HelloWorld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demoServiceClient) SpellMyName(ctx context.Context, in *HelloWorldRequest, opts ...grpc.CallOption) (DemoService_SpellMyNameClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DemoService_serviceDesc.Streams[0], "/rpc.DemoService/SpellMyName", opts...)
	if err != nil {
		return nil, err
	}
	x := &demoServiceSpellMyNameClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DemoService_SpellMyNameClient interface {
	Recv() (*LetterResponse, error)
	grpc.ClientStream
}

type demoServiceSpellMyNameClient struct {
	grpc.ClientStream
}

func (x *demoServiceSpellMyNameClient) Recv() (*LetterResponse, error) {
	m := new(LetterResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DemoServiceServer is the server API for DemoService service.
type DemoServiceServer interface {
	HelloWorld(context.Context, *HelloWorldRequest) (*HelloWorldResponse, error)
	SpellMyName(*HelloWorldRequest, DemoService_SpellMyNameServer) error
}

// UnimplementedDemoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDemoServiceServer struct {
}

func (*UnimplementedDemoServiceServer) HelloWorld(ctx context.Context, req *HelloWorldRequest) (*HelloWorldResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HelloWorld not implemented")
}
func (*UnimplementedDemoServiceServer) SpellMyName(req *HelloWorldRequest, srv DemoService_SpellMyNameServer) error {
	return status.Errorf(codes.Unimplemented, "method SpellMyName not implemented")
}

func RegisterDemoServiceServer(s *grpc.Server, srv DemoServiceServer) {
	s.RegisterService(&_DemoService_serviceDesc, srv)
}

func _DemoService_HelloWorld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloWorldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoServiceServer).HelloWorld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.DemoService/HelloWorld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoServiceServer).HelloWorld(ctx, req.(*HelloWorldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DemoService_SpellMyName_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloWorldRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DemoServiceServer).SpellMyName(m, &demoServiceSpellMyNameServer{stream})
}

type DemoService_SpellMyNameServer interface {
	Send(*LetterResponse) error
	grpc.ServerStream
}

type demoServiceSpellMyNameServer struct {
	grpc.ServerStream
}

func (x *demoServiceSpellMyNameServer) Send(m *LetterResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _DemoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.DemoService",
	HandlerType: (*DemoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HelloWorld",
			Handler:    _DemoService_HelloWorld_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SpellMyName",
			Handler:       _DemoService_SpellMyName_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "demo.proto",
}
