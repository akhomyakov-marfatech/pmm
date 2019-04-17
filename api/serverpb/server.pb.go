// Code generated by protoc-gen-go. DO NOT EDIT.
// source: serverpb/server.proto

package serverpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

type VersionRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionRequest) Reset()         { *m = VersionRequest{} }
func (m *VersionRequest) String() string { return proto.CompactTextString(m) }
func (*VersionRequest) ProtoMessage()    {}
func (*VersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9202d8f598083902, []int{0}
}

func (m *VersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionRequest.Unmarshal(m, b)
}
func (m *VersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionRequest.Marshal(b, m, deterministic)
}
func (m *VersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionRequest.Merge(m, src)
}
func (m *VersionRequest) XXX_Size() int {
	return xxx_messageInfo_VersionRequest.Size(m)
}
func (m *VersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VersionRequest proto.InternalMessageInfo

type VersionResponse struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionResponse) Reset()         { *m = VersionResponse{} }
func (m *VersionResponse) String() string { return proto.CompactTextString(m) }
func (*VersionResponse) ProtoMessage()    {}
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9202d8f598083902, []int{1}
}

func (m *VersionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionResponse.Unmarshal(m, b)
}
func (m *VersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionResponse.Marshal(b, m, deterministic)
}
func (m *VersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionResponse.Merge(m, src)
}
func (m *VersionResponse) XXX_Size() int {
	return xxx_messageInfo_VersionResponse.Size(m)
}
func (m *VersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VersionResponse proto.InternalMessageInfo

func (m *VersionResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func init() {
	proto.RegisterType((*VersionRequest)(nil), "server.VersionRequest")
	proto.RegisterType((*VersionResponse)(nil), "server.VersionResponse")
}

func init() { proto.RegisterFile("serverpb/server.proto", fileDescriptor_9202d8f598083902) }

var fileDescriptor_9202d8f598083902 = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x4e, 0x2d, 0x2a,
	0x4b, 0x2d, 0x2a, 0x48, 0xd2, 0x87, 0x30, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xd8, 0x20,
	0x3c, 0x29, 0x99, 0xf4, 0xfc, 0xfc, 0xf4, 0x9c, 0x54, 0xfd, 0xc4, 0x82, 0x4c, 0xfd, 0xc4, 0xbc,
	0xbc, 0xfc, 0x92, 0xc4, 0x92, 0xcc, 0xfc, 0xbc, 0x62, 0x88, 0x2a, 0x29, 0x1d, 0x30, 0x95, 0xac,
	0x9b, 0x9e, 0x9a, 0xa7, 0x5b, 0x5c, 0x9e, 0x98, 0x9e, 0x9e, 0x5a, 0xa4, 0x9f, 0x5f, 0x00, 0x56,
	0x81, 0xa9, 0x5a, 0x49, 0x80, 0x8b, 0x2f, 0x2c, 0xb5, 0xa8, 0x38, 0x33, 0x3f, 0x2f, 0x28, 0xb5,
	0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x49, 0x9b, 0x8b, 0x1f, 0x2e, 0x52, 0x5c, 0x90, 0x9f, 0x57, 0x9c,
	0x2a, 0x24, 0xc1, 0xc5, 0x5e, 0x06, 0x11, 0x92, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x71,
	0x8d, 0x22, 0xb9, 0xd8, 0x82, 0xc1, 0x8e, 0x12, 0xf2, 0xe7, 0x62, 0x87, 0x6a, 0x13, 0x12, 0xd3,
	0x83, 0x3a, 0x1b, 0xd5, 0x64, 0x29, 0x71, 0x0c, 0x71, 0x88, 0xf9, 0x4a, 0xc2, 0x4d, 0x97, 0x9f,
	0x4c, 0x66, 0xe2, 0x15, 0xe2, 0xd6, 0x2f, 0x33, 0xd4, 0x87, 0x1a, 0xed, 0x14, 0x3e, 0xc9, 0xd1,
	0x2b, 0xc8, 0x83, 0x8b, 0x3d, 0x25, 0x35, 0x2d, 0xb1, 0x34, 0xa7, 0x44, 0xc8, 0x96, 0x4b, 0xc8,
	0x31, 0x4f, 0x21, 0xb5, 0xa8, 0x28, 0xbf, 0x48, 0xa1, 0x08, 0xaa, 0x53, 0x4f, 0x48, 0x9d, 0x4b,
	0x55, 0x4a, 0x59, 0x59, 0x3f, 0x25, 0x35, 0x2d, 0x33, 0x2f, 0x13, 0xe2, 0x49, 0x58, 0xe0, 0xb9,
	0x82, 0x94, 0xc2, 0xec, 0x88, 0xe2, 0x80, 0x09, 0x27, 0xb1, 0x81, 0x7d, 0x6e, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x3d, 0x9b, 0x62, 0x60, 0x66, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServerClient is the client API for Server service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServerClient interface {
	// Version returns PMM Server version.
	Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
}

type serverClient struct {
	cc *grpc.ClientConn
}

func NewServerClient(cc *grpc.ClientConn) ServerClient {
	return &serverClient{cc}
}

func (c *serverClient) Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/server.Server/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerServer is the server API for Server service.
type ServerServer interface {
	// Version returns PMM Server version.
	Version(context.Context, *VersionRequest) (*VersionResponse, error)
}

func RegisterServerServer(s *grpc.Server, srv ServerServer) {
	s.RegisterService(&_Server_serviceDesc, srv)
}

func _Server_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Server/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).Version(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Server_serviceDesc = grpc.ServiceDesc{
	ServiceName: "server.Server",
	HandlerType: (*ServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _Server_Version_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "serverpb/server.proto",
}