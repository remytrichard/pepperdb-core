// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: grpc.proto

/*
Package dappserverpb is a generated protocol buffer package.

It is generated from these files:
	grpc.proto

It has these top-level messages:
	StoreRequest
	StoreResponse
*/
package dappserverpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// The request message of Store
type StoreRequest struct {
	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	File        []byte `protobuf:"bytes,3,opt,name=file,proto3" json:"file,omitempty"`
	Md5         string `protobuf:"bytes,4,opt,name=md5,proto3" json:"md5,omitempty"`
	Type        string `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
}

func (m *StoreRequest) Reset()                    { *m = StoreRequest{} }
func (m *StoreRequest) String() string            { return proto.CompactTextString(m) }
func (*StoreRequest) ProtoMessage()               {}
func (*StoreRequest) Descriptor() ([]byte, []int) { return fileDescriptorGrpc, []int{0} }

func (m *StoreRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StoreRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *StoreRequest) GetFile() []byte {
	if m != nil {
		return m.File
	}
	return nil
}

func (m *StoreRequest) GetMd5() string {
	if m != nil {
		return m.Md5
	}
	return ""
}

func (m *StoreRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

// The reponse message of Store
type StoreResponse struct {
	Ok      bool   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *StoreResponse) Reset()                    { *m = StoreResponse{} }
func (m *StoreResponse) String() string            { return proto.CompactTextString(m) }
func (*StoreResponse) ProtoMessage()               {}
func (*StoreResponse) Descriptor() ([]byte, []int) { return fileDescriptorGrpc, []int{1} }

func (m *StoreResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *StoreResponse) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*StoreRequest)(nil), "dappserverpb.StoreRequest")
	proto.RegisterType((*StoreResponse)(nil), "dappserverpb.StoreResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DAppServer service

type DAppServerClient interface {
	// Store a dapp file
	Store(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*StoreResponse, error)
}

type dAppServerClient struct {
	cc *grpc.ClientConn
}

func NewDAppServerClient(cc *grpc.ClientConn) DAppServerClient {
	return &dAppServerClient{cc}
}

func (c *dAppServerClient) Store(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*StoreResponse, error) {
	out := new(StoreResponse)
	err := grpc.Invoke(ctx, "/dappserverpb.DAppServer/Store", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DAppServer service

type DAppServerServer interface {
	// Store a dapp file
	Store(context.Context, *StoreRequest) (*StoreResponse, error)
}

func RegisterDAppServerServer(s *grpc.Server, srv DAppServerServer) {
	s.RegisterService(&_DAppServer_serviceDesc, srv)
}

func _DAppServer_Store_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DAppServerServer).Store(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dappserverpb.DAppServer/Store",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DAppServerServer).Store(ctx, req.(*StoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DAppServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dappserverpb.DAppServer",
	HandlerType: (*DAppServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Store",
			Handler:    _DAppServer_Store_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc.proto",
}

func init() { proto.RegisterFile("grpc.proto", fileDescriptorGrpc) }

var fileDescriptorGrpc = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcd, 0x4e, 0x85, 0x30,
	0x10, 0x85, 0x85, 0x7b, 0xaf, 0x3f, 0x23, 0x1a, 0x33, 0xab, 0x06, 0x37, 0x84, 0x15, 0x2b, 0x16,
	0x1a, 0x17, 0x2e, 0x35, 0x3e, 0x80, 0x29, 0x4f, 0x00, 0x74, 0x34, 0x04, 0xa1, 0x63, 0x5b, 0x4d,
	0xdc, 0xf8, 0xec, 0xa6, 0x03, 0x26, 0x2c, 0xee, 0xee, 0x6b, 0x67, 0x7a, 0xbe, 0x9c, 0x02, 0xbc,
	0x3b, 0xee, 0x6b, 0x76, 0x36, 0x58, 0xcc, 0x4c, 0xcb, 0xec, 0xc9, 0x7d, 0x93, 0xe3, 0xae, 0xfc,
	0x85, 0xac, 0x09, 0xd6, 0x91, 0xa6, 0xcf, 0x2f, 0xf2, 0x01, 0x11, 0xf6, 0x73, 0x3b, 0x91, 0x4a,
	0x8a, 0xa4, 0xba, 0xd0, 0xc2, 0x58, 0xc0, 0xa5, 0x21, 0xdf, 0xbb, 0x81, 0xc3, 0x60, 0x67, 0x95,
	0xca, 0x68, 0x7b, 0x15, 0x5f, 0xbd, 0x0d, 0x1f, 0xa4, 0x76, 0x45, 0x52, 0x65, 0x5a, 0x18, 0x6f,
	0x60, 0x37, 0x99, 0x07, 0xb5, 0x97, 0xed, 0x88, 0x71, 0x2b, 0xfc, 0x30, 0xa9, 0xc3, 0x92, 0x1d,
	0xb9, 0x7c, 0x84, 0xab, 0xd5, 0xef, 0xd9, 0xce, 0x9e, 0xf0, 0x1a, 0x52, 0x3b, 0x8a, 0xfe, 0x5c,
	0xa7, 0x76, 0x44, 0x05, 0x67, 0xad, 0x31, 0x8e, 0xbc, 0x5f, 0xc5, 0xff, 0xc7, 0xbb, 0x57, 0x80,
	0x97, 0x27, 0xe6, 0x46, 0xaa, 0xe0, 0x33, 0x1c, 0x24, 0x08, 0xf3, 0x7a, 0x5b, 0xb0, 0xde, 0xb6,
	0xcb, 0x6f, 0x8f, 0xce, 0x16, 0x73, 0x79, 0xd2, 0x9d, 0xca, 0x0f, 0xdd, 0xff, 0x05, 0x00, 0x00,
	0xff, 0xff, 0xf7, 0x20, 0x0f, 0x50, 0x2f, 0x01, 0x00, 0x00,
}
