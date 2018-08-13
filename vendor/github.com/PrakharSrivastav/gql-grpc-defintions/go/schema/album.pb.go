// Code generated by protoc-gen-go. DO NOT EDIT.
// source: schema/album.proto

package no_sysco_middleware_grpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"

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

type Album struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ArtistId             string   `protobuf:"bytes,3,opt,name=artistId,proto3" json:"artistId,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Album) Reset()         { *m = Album{} }
func (m *Album) String() string { return proto.CompactTextString(m) }
func (*Album) ProtoMessage()    {}
func (*Album) Descriptor() ([]byte, []int) {
	return fileDescriptor_album_123f201eee8cf64d, []int{0}
}
func (m *Album) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Album.Unmarshal(m, b)
}
func (m *Album) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Album.Marshal(b, m, deterministic)
}
func (dst *Album) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Album.Merge(dst, src)
}
func (m *Album) XXX_Size() int {
	return xxx_messageInfo_Album.Size(m)
}
func (m *Album) XXX_DiscardUnknown() {
	xxx_messageInfo_Album.DiscardUnknown(m)
}

var xxx_messageInfo_Album proto.InternalMessageInfo

func (m *Album) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Album) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Album) GetArtistId() string {
	if m != nil {
		return m.ArtistId
	}
	return ""
}

func (m *Album) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type SimpleAlbumRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleAlbumRequest) Reset()         { *m = SimpleAlbumRequest{} }
func (m *SimpleAlbumRequest) String() string { return proto.CompactTextString(m) }
func (*SimpleAlbumRequest) ProtoMessage()    {}
func (*SimpleAlbumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_album_123f201eee8cf64d, []int{1}
}
func (m *SimpleAlbumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleAlbumRequest.Unmarshal(m, b)
}
func (m *SimpleAlbumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleAlbumRequest.Marshal(b, m, deterministic)
}
func (dst *SimpleAlbumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleAlbumRequest.Merge(dst, src)
}
func (m *SimpleAlbumRequest) XXX_Size() int {
	return xxx_messageInfo_SimpleAlbumRequest.Size(m)
}
func (m *SimpleAlbumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleAlbumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleAlbumRequest proto.InternalMessageInfo

func (m *SimpleAlbumRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Album)(nil), "no.sysco.middleware.grpc.Album")
	proto.RegisterType((*SimpleAlbumRequest)(nil), "no.sysco.middleware.grpc.SimpleAlbumRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AlbumServiceClient is the client API for AlbumService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AlbumServiceClient interface {
	GetAll(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (AlbumService_GetAllClient, error)
	GetAlbumByArtist(ctx context.Context, in *SimpleAlbumRequest, opts ...grpc.CallOption) (AlbumService_GetAlbumByArtistClient, error)
	GetAlbumByTrack(ctx context.Context, in *SimpleAlbumRequest, opts ...grpc.CallOption) (*Album, error)
}

type albumServiceClient struct {
	cc *grpc.ClientConn
}

func NewAlbumServiceClient(cc *grpc.ClientConn) AlbumServiceClient {
	return &albumServiceClient{cc}
}

func (c *albumServiceClient) GetAll(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (AlbumService_GetAllClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AlbumService_serviceDesc.Streams[0], "/no.sysco.middleware.grpc.AlbumService/GetAll", opts...)
	if err != nil {
		return nil, err
	}
	x := &albumServiceGetAllClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AlbumService_GetAllClient interface {
	Recv() (*Album, error)
	grpc.ClientStream
}

type albumServiceGetAllClient struct {
	grpc.ClientStream
}

func (x *albumServiceGetAllClient) Recv() (*Album, error) {
	m := new(Album)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *albumServiceClient) GetAlbumByArtist(ctx context.Context, in *SimpleAlbumRequest, opts ...grpc.CallOption) (AlbumService_GetAlbumByArtistClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AlbumService_serviceDesc.Streams[1], "/no.sysco.middleware.grpc.AlbumService/GetAlbumByArtist", opts...)
	if err != nil {
		return nil, err
	}
	x := &albumServiceGetAlbumByArtistClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AlbumService_GetAlbumByArtistClient interface {
	Recv() (*Album, error)
	grpc.ClientStream
}

type albumServiceGetAlbumByArtistClient struct {
	grpc.ClientStream
}

func (x *albumServiceGetAlbumByArtistClient) Recv() (*Album, error) {
	m := new(Album)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *albumServiceClient) GetAlbumByTrack(ctx context.Context, in *SimpleAlbumRequest, opts ...grpc.CallOption) (*Album, error) {
	out := new(Album)
	err := c.cc.Invoke(ctx, "/no.sysco.middleware.grpc.AlbumService/GetAlbumByTrack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlbumServiceServer is the server API for AlbumService service.
type AlbumServiceServer interface {
	GetAll(*empty.Empty, AlbumService_GetAllServer) error
	GetAlbumByArtist(*SimpleAlbumRequest, AlbumService_GetAlbumByArtistServer) error
	GetAlbumByTrack(context.Context, *SimpleAlbumRequest) (*Album, error)
}

func RegisterAlbumServiceServer(s *grpc.Server, srv AlbumServiceServer) {
	s.RegisterService(&_AlbumService_serviceDesc, srv)
}

func _AlbumService_GetAll_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AlbumServiceServer).GetAll(m, &albumServiceGetAllServer{stream})
}

type AlbumService_GetAllServer interface {
	Send(*Album) error
	grpc.ServerStream
}

type albumServiceGetAllServer struct {
	grpc.ServerStream
}

func (x *albumServiceGetAllServer) Send(m *Album) error {
	return x.ServerStream.SendMsg(m)
}

func _AlbumService_GetAlbumByArtist_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SimpleAlbumRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AlbumServiceServer).GetAlbumByArtist(m, &albumServiceGetAlbumByArtistServer{stream})
}

type AlbumService_GetAlbumByArtistServer interface {
	Send(*Album) error
	grpc.ServerStream
}

type albumServiceGetAlbumByArtistServer struct {
	grpc.ServerStream
}

func (x *albumServiceGetAlbumByArtistServer) Send(m *Album) error {
	return x.ServerStream.SendMsg(m)
}

func _AlbumService_GetAlbumByTrack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleAlbumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlbumServiceServer).GetAlbumByTrack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/no.sysco.middleware.grpc.AlbumService/GetAlbumByTrack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlbumServiceServer).GetAlbumByTrack(ctx, req.(*SimpleAlbumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AlbumService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "no.sysco.middleware.grpc.AlbumService",
	HandlerType: (*AlbumServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAlbumByTrack",
			Handler:    _AlbumService_GetAlbumByTrack_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAll",
			Handler:       _AlbumService_GetAll_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetAlbumByArtist",
			Handler:       _AlbumService_GetAlbumByArtist_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "schema/album.proto",
}

func init() { proto.RegisterFile("schema/album.proto", fileDescriptor_album_123f201eee8cf64d) }

var fileDescriptor_album_123f201eee8cf64d = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0xdf, 0x4a, 0xf4, 0x30,
	0x10, 0xc5, 0xbf, 0xed, 0xb7, 0x2e, 0x3a, 0x8a, 0xca, 0x08, 0x52, 0xd6, 0x0b, 0x97, 0xc5, 0x0b,
	0x2f, 0x24, 0x15, 0x7d, 0x82, 0x2d, 0x2e, 0xe2, 0xed, 0xae, 0x2f, 0x90, 0x26, 0xb3, 0x35, 0x98,
	0x34, 0x35, 0x49, 0x95, 0xbe, 0x8d, 0x8f, 0x2a, 0x4d, 0xf1, 0x0f, 0x4a, 0xc1, 0x0b, 0xef, 0x92,
	0x93, 0x99, 0xf3, 0x9b, 0x33, 0x01, 0xf4, 0xe2, 0x81, 0x0c, 0xcf, 0xb8, 0x2e, 0x1a, 0xc3, 0x6a,
	0x67, 0x83, 0xc5, 0xb4, 0xb2, 0xcc, 0xb7, 0x5e, 0x58, 0x66, 0x94, 0x94, 0x9a, 0x5e, 0xb8, 0x23,
	0x56, 0xba, 0x5a, 0x4c, 0x4f, 0x4a, 0x6b, 0x4b, 0x4d, 0x59, 0xac, 0x2b, 0x9a, 0x4d, 0x46, 0xa6,
	0x0e, 0x6d, 0xdf, 0x36, 0x57, 0xb0, 0xb5, 0xe8, 0x5c, 0x70, 0x1f, 0x12, 0x25, 0xd3, 0xd1, 0x6c,
	0x74, 0xbe, 0xb3, 0x4a, 0x94, 0x44, 0x84, 0x71, 0xc5, 0x0d, 0xa5, 0x49, 0x54, 0xe2, 0x19, 0xa7,
	0xb0, 0xcd, 0x5d, 0x50, 0x3e, 0xdc, 0xc9, 0xf4, 0x7f, 0xd4, 0x3f, 0xee, 0x38, 0x83, 0x5d, 0x49,
	0x5e, 0x38, 0x55, 0x07, 0x65, 0xab, 0x74, 0x1c, 0x9f, 0xbf, 0x4a, 0xf3, 0x33, 0xc0, 0xb5, 0x32,
	0xb5, 0xa6, 0x08, 0x5c, 0xd1, 0x53, 0x43, 0x3e, 0x7c, 0xe7, 0x5e, 0xbd, 0x26, 0xb0, 0x17, 0x0b,
	0xd6, 0xe4, 0x9e, 0x95, 0x20, 0x5c, 0xc2, 0xe4, 0x96, 0xc2, 0x42, 0x6b, 0x3c, 0x66, 0x7d, 0x12,
	0xf6, 0x9e, 0x84, 0x2d, 0xbb, 0x24, 0xd3, 0x53, 0x36, 0x94, 0x9d, 0x45, 0xa7, 0xf9, 0xbf, 0xcb,
	0x11, 0x12, 0x1c, 0x46, 0x9b, 0xa2, 0x31, 0x79, 0xbb, 0x88, 0x53, 0xe3, 0xc5, 0x70, 0xe3, 0xcf,
	0x49, 0x7f, 0x87, 0x29, 0xe0, 0xe0, 0x13, 0x73, 0xef, 0xb8, 0x78, 0xfc, 0x73, 0x4a, 0x9e, 0xc1,
	0xe0, 0x67, 0xe7, 0x47, 0x3d, 0x9a, 0x7b, 0xba, 0xa1, 0x8d, 0xaa, 0x54, 0xb7, 0xf9, 0x62, 0x12,
	0x17, 0x76, 0xfd, 0x16, 0x00, 0x00, 0xff, 0xff, 0x3d, 0x3c, 0x24, 0xb5, 0x38, 0x02, 0x00, 0x00,
}