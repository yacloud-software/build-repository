// Code generated by protoc-gen-go.
// source: golang.conradwood.net/apis/heating/heating.proto
// DO NOT EDIT!

/*
Package heating is a generated protocol buffer package.

It is generated from these files:
	golang.conradwood.net/apis/heating/heating.proto

It has these top-level messages:
	SetRequest
	Connection
	ConnectionList
*/
package heating

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "golang.conradwood.net/apis/common"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SetRequest struct {
	NodeID uint64 `protobuf:"varint,1,opt,name=NodeID" json:"NodeID,omitempty"`
	On     bool   `protobuf:"varint,2,opt,name=On" json:"On,omitempty"`
}

func (m *SetRequest) Reset()                    { *m = SetRequest{} }
func (m *SetRequest) String() string            { return proto.CompactTextString(m) }
func (*SetRequest) ProtoMessage()               {}
func (*SetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SetRequest) GetNodeID() uint64 {
	if m != nil {
		return m.NodeID
	}
	return 0
}

func (m *SetRequest) GetOn() bool {
	if m != nil {
		return m.On
	}
	return false
}

type Connection struct {
	NodeID  uint64 `protobuf:"varint,1,opt,name=NodeID" json:"NodeID,omitempty"`
	Mac     string `protobuf:"bytes,2,opt,name=Mac" json:"Mac,omitempty"`
	Name    string `protobuf:"bytes,3,opt,name=Name" json:"Name,omitempty"`
	IP      string `protobuf:"bytes,4,opt,name=IP" json:"IP,omitempty"`
	Relay   bool   `protobuf:"varint,5,opt,name=Relay" json:"Relay,omitempty"`
	Version uint64 `protobuf:"varint,6,opt,name=Version" json:"Version,omitempty"`
}

func (m *Connection) Reset()                    { *m = Connection{} }
func (m *Connection) String() string            { return proto.CompactTextString(m) }
func (*Connection) ProtoMessage()               {}
func (*Connection) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Connection) GetNodeID() uint64 {
	if m != nil {
		return m.NodeID
	}
	return 0
}

func (m *Connection) GetMac() string {
	if m != nil {
		return m.Mac
	}
	return ""
}

func (m *Connection) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Connection) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *Connection) GetRelay() bool {
	if m != nil {
		return m.Relay
	}
	return false
}

func (m *Connection) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

type ConnectionList struct {
	Connections []*Connection `protobuf:"bytes,1,rep,name=Connections" json:"Connections,omitempty"`
}

func (m *ConnectionList) Reset()                    { *m = ConnectionList{} }
func (m *ConnectionList) String() string            { return proto.CompactTextString(m) }
func (*ConnectionList) ProtoMessage()               {}
func (*ConnectionList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ConnectionList) GetConnections() []*Connection {
	if m != nil {
		return m.Connections
	}
	return nil
}

func init() {
	proto.RegisterType((*SetRequest)(nil), "heating.SetRequest")
	proto.RegisterType((*Connection)(nil), "heating.Connection")
	proto.RegisterType((*ConnectionList)(nil), "heating.ConnectionList")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for HeatingService service

type HeatingServiceClient interface {
	// turn a relay on/off
	SetRelay(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*common.Void, error)
	// what's connected:
	Connected(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*ConnectionList, error)
}

type heatingServiceClient struct {
	cc *grpc.ClientConn
}

func NewHeatingServiceClient(cc *grpc.ClientConn) HeatingServiceClient {
	return &heatingServiceClient{cc}
}

func (c *heatingServiceClient) SetRelay(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*common.Void, error) {
	out := new(common.Void)
	err := grpc.Invoke(ctx, "/heating.HeatingService/SetRelay", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *heatingServiceClient) Connected(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*ConnectionList, error) {
	out := new(ConnectionList)
	err := grpc.Invoke(ctx, "/heating.HeatingService/Connected", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for HeatingService service

type HeatingServiceServer interface {
	// turn a relay on/off
	SetRelay(context.Context, *SetRequest) (*common.Void, error)
	// what's connected:
	Connected(context.Context, *common.Void) (*ConnectionList, error)
}

func RegisterHeatingServiceServer(s *grpc.Server, srv HeatingServiceServer) {
	s.RegisterService(&_HeatingService_serviceDesc, srv)
}

func _HeatingService_SetRelay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeatingServiceServer).SetRelay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heating.HeatingService/SetRelay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeatingServiceServer).SetRelay(ctx, req.(*SetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HeatingService_Connected_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeatingServiceServer).Connected(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heating.HeatingService/Connected",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeatingServiceServer).Connected(ctx, req.(*common.Void))
	}
	return interceptor(ctx, in, info, handler)
}

var _HeatingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "heating.HeatingService",
	HandlerType: (*HeatingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetRelay",
			Handler:    _HeatingService_SetRelay_Handler,
		},
		{
			MethodName: "Connected",
			Handler:    _HeatingService_Connected_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "golang.conradwood.net/apis/heating/heating.proto",
}

func init() { proto.RegisterFile("golang.conradwood.net/apis/heating/heating.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x51, 0x4d, 0x4b, 0x03, 0x41,
	0x0c, 0x65, 0xb7, 0xdf, 0xa9, 0x14, 0x19, 0x45, 0x87, 0x9e, 0xca, 0xe2, 0xa1, 0x17, 0xb7, 0x52,
	0xf5, 0x0f, 0xa8, 0xa0, 0x05, 0x6d, 0xcb, 0x16, 0x7a, 0x1f, 0x77, 0x43, 0x1d, 0x68, 0x93, 0xba,
	0x33, 0x2a, 0xfe, 0x05, 0x7f, 0xb5, 0xec, 0x74, 0x96, 0xad, 0xa0, 0x9e, 0x92, 0x97, 0xe4, 0x25,
	0xe1, 0x3d, 0xb8, 0x58, 0xf1, 0x5a, 0xd1, 0x2a, 0x4e, 0x99, 0x72, 0x95, 0x7d, 0x30, 0x67, 0x31,
	0xa1, 0x1d, 0xa9, 0xad, 0x36, 0xa3, 0x17, 0x54, 0x56, 0xd3, 0xaa, 0x8c, 0xf1, 0x36, 0x67, 0xcb,
	0xa2, 0xe5, 0x61, 0x3f, 0xfe, 0x87, 0x9a, 0xf2, 0x66, 0xc3, 0xe4, 0xc3, 0x8e, 0x18, 0x5d, 0x01,
	0x2c, 0xd0, 0x26, 0xf8, 0xfa, 0x86, 0xc6, 0x8a, 0x13, 0x68, 0x4e, 0x39, 0xc3, 0xc9, 0x9d, 0x0c,
	0x06, 0xc1, 0xb0, 0x9e, 0x78, 0x24, 0x7a, 0x10, 0xce, 0x48, 0x86, 0x83, 0x60, 0xd8, 0x4e, 0xc2,
	0x19, 0x45, 0x5f, 0x01, 0xc0, 0x2d, 0x13, 0x61, 0x6a, 0x35, 0xd3, 0x9f, 0xb4, 0x43, 0xa8, 0x3d,
	0xa9, 0xd4, 0xf1, 0x3a, 0x49, 0x91, 0x0a, 0x01, 0xf5, 0xa9, 0xda, 0xa0, 0xac, 0xb9, 0x92, 0xcb,
	0x8b, 0xe5, 0x93, 0xb9, 0xac, 0xbb, 0x4a, 0x38, 0x99, 0x8b, 0x63, 0x68, 0x24, 0xb8, 0x56, 0x9f,
	0xb2, 0xe1, 0xee, 0xed, 0x80, 0x90, 0xd0, 0x5a, 0x62, 0x6e, 0x34, 0x93, 0x6c, 0xba, 0x23, 0x25,
	0x8c, 0xee, 0xa1, 0x57, 0xfd, 0xf2, 0xa8, 0x8d, 0x15, 0xd7, 0xd0, 0xad, 0x2a, 0x46, 0x06, 0x83,
	0xda, 0xb0, 0x3b, 0x3e, 0x8a, 0x4b, 0xc9, 0xaa, 0x5e, 0xb2, 0x3f, 0x37, 0x36, 0xd0, 0x7b, 0xd8,
	0x8d, 0x2c, 0x30, 0x7f, 0xd7, 0x29, 0x8a, 0x73, 0x68, 0x3b, 0x75, 0x8a, 0x07, 0x2a, 0x7e, 0x25,
	0x58, 0xff, 0x20, 0xf6, 0x6a, 0x2e, 0x59, 0x67, 0x62, 0x0c, 0x1d, 0xbf, 0x0f, 0x33, 0xf1, 0xa3,
	0xd5, 0x3f, 0xfd, 0xe5, 0x7a, 0xf1, 0xeb, 0xcd, 0x19, 0x44, 0x84, 0x76, 0xdf, 0x2f, 0xef, 0x60,
	0x61, 0x59, 0x49, 0x7a, 0x6e, 0x3a, 0xb7, 0x2e, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x00, 0xb3,
	0x40, 0x78, 0x1a, 0x02, 0x00, 0x00,
}
