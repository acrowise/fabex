// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fabex.proto

package proto

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

type RequestRange struct {
	Startblock           int64    `protobuf:"varint,1,opt,name=startblock,proto3" json:"startblock,omitempty"`
	Endblock             int64    `protobuf:"varint,2,opt,name=endblock,proto3" json:"endblock,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestRange) Reset()         { *m = RequestRange{} }
func (m *RequestRange) String() string { return proto.CompactTextString(m) }
func (*RequestRange) ProtoMessage()    {}
func (*RequestRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7d7206373264ff4, []int{0}
}

func (m *RequestRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestRange.Unmarshal(m, b)
}
func (m *RequestRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestRange.Marshal(b, m, deterministic)
}
func (m *RequestRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestRange.Merge(m, src)
}
func (m *RequestRange) XXX_Size() int {
	return xxx_messageInfo_RequestRange.Size(m)
}
func (m *RequestRange) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestRange.DiscardUnknown(m)
}

var xxx_messageInfo_RequestRange proto.InternalMessageInfo

func (m *RequestRange) GetStartblock() int64 {
	if m != nil {
		return m.Startblock
	}
	return 0
}

func (m *RequestRange) GetEndblock() int64 {
	if m != nil {
		return m.Endblock
	}
	return 0
}

type Entry struct {
	Channelid            string   `protobuf:"bytes,1,opt,name=channelid,proto3" json:"channelid,omitempty"`
	Txid                 string   `protobuf:"bytes,2,opt,name=txid,proto3" json:"txid,omitempty"`
	Hash                 string   `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	Previoushash         string   `protobuf:"bytes,4,opt,name=previoushash,proto3" json:"previoushash,omitempty"`
	Blocknum             uint64   `protobuf:"varint,5,opt,name=blocknum,proto3" json:"blocknum,omitempty"`
	Payload              string   `protobuf:"bytes,6,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Entry) Reset()         { *m = Entry{} }
func (m *Entry) String() string { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()    {}
func (*Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7d7206373264ff4, []int{1}
}

func (m *Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entry.Unmarshal(m, b)
}
func (m *Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entry.Marshal(b, m, deterministic)
}
func (m *Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entry.Merge(m, src)
}
func (m *Entry) XXX_Size() int {
	return xxx_messageInfo_Entry.Size(m)
}
func (m *Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_Entry proto.InternalMessageInfo

func (m *Entry) GetChannelid() string {
	if m != nil {
		return m.Channelid
	}
	return ""
}

func (m *Entry) GetTxid() string {
	if m != nil {
		return m.Txid
	}
	return ""
}

func (m *Entry) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Entry) GetPrevioushash() string {
	if m != nil {
		return m.Previoushash
	}
	return ""
}

func (m *Entry) GetBlocknum() uint64 {
	if m != nil {
		return m.Blocknum
	}
	return 0
}

func (m *Entry) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func init() {
	proto.RegisterType((*RequestRange)(nil), "fabex.RequestRange")
	proto.RegisterType((*Entry)(nil), "fabex.Entry")
}

func init() { proto.RegisterFile("fabex.proto", fileDescriptor_d7d7206373264ff4) }

var fileDescriptor_d7d7206373264ff4 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xcf, 0x4a, 0xf3, 0x40,
	0x10, 0x27, 0x5f, 0x9b, 0xf6, 0xcb, 0x58, 0x2f, 0x23, 0x42, 0x28, 0x22, 0xa5, 0x88, 0xd4, 0x83,
	0xa9, 0x28, 0xbe, 0x40, 0xa0, 0x4a, 0x3d, 0x49, 0xf0, 0xe4, 0x6d, 0xd3, 0xdd, 0x36, 0x4b, 0xd3,
	0xdd, 0x98, 0x6c, 0x42, 0xf2, 0x44, 0xbe, 0x8a, 0x8f, 0x25, 0x3b, 0xb1, 0xb5, 0x45, 0xe8, 0x69,
	0xe7, 0xf7, 0x67, 0x76, 0x7f, 0x33, 0x0b, 0x27, 0x4b, 0x16, 0x8b, 0x3a, 0xc8, 0x72, 0x6d, 0x34,
	0xba, 0x04, 0xc6, 0x2f, 0x30, 0x88, 0xc4, 0x47, 0x29, 0x0a, 0x13, 0x31, 0xb5, 0x12, 0x78, 0x09,
	0x50, 0x18, 0x96, 0x9b, 0x38, 0xd5, 0x8b, 0xb5, 0xef, 0x8c, 0x9c, 0x49, 0x27, 0xda, 0x63, 0x70,
	0x08, 0xff, 0x85, 0xe2, 0xad, 0xfa, 0x8f, 0xd4, 0x1d, 0x1e, 0x7f, 0x3a, 0xe0, 0xce, 0x94, 0xc9,
	0x1b, 0xbc, 0x00, 0x6f, 0x91, 0x30, 0xa5, 0x44, 0x2a, 0x39, 0x5d, 0xe2, 0x45, 0xbf, 0x04, 0x22,
	0x74, 0x4d, 0x2d, 0x39, 0xf5, 0x7b, 0x11, 0xd5, 0x96, 0x4b, 0x58, 0x91, 0xf8, 0x9d, 0x96, 0xb3,
	0x35, 0x8e, 0x61, 0x90, 0xe5, 0xa2, 0x92, 0xba, 0x2c, 0x48, 0xeb, 0x92, 0x76, 0xc0, 0xd9, 0x3c,
	0xf4, 0xb8, 0x2a, 0x37, 0xbe, 0x3b, 0x72, 0x26, 0xdd, 0x68, 0x87, 0xd1, 0x87, 0x7e, 0xc6, 0x9a,
	0x54, 0x33, 0xee, 0xf7, 0xa8, 0x75, 0x0b, 0xef, 0xbf, 0x1c, 0x70, 0x9f, 0xec, 0xfc, 0x18, 0x40,
	0x7f, 0x56, 0x67, 0xa9, 0xce, 0x05, 0x9e, 0x05, 0xed, 0x7e, 0xf6, 0xf7, 0x31, 0x1c, 0xfc, 0x90,
	0x34, 0xd7, 0x9d, 0x83, 0x37, 0xe0, 0x3d, 0x0b, 0x13, 0x36, 0x6f, 0xf5, 0x9c, 0xe3, 0x81, 0xf8,
	0xc7, 0x7a, 0x0b, 0xa7, 0x64, 0x0d, 0xb7, 0x79, 0x8e, 0xdb, 0x1f, 0xe1, 0xdc, 0xda, 0xad, 0x79,
	0xae, 0x96, 0x3a, 0x6c, 0x5e, 0xdb, 0xb0, 0xc7, 0xdb, 0xc2, 0xeb, 0xf7, 0xab, 0x95, 0x34, 0x49,
	0x19, 0x07, 0x0b, 0xbd, 0x99, 0x56, 0x8c, 0xcb, 0x8d, 0x54, 0x45, 0xc2, 0xd6, 0xba, 0x9a, 0x92,
	0x73, 0x4a, 0xff, 0x1d, 0xf7, 0xe8, 0x78, 0xf8, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xab, 0x97, 0x1a,
	0x68, 0x05, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FabexClient is the client API for Fabex service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FabexClient interface {
	Explore(ctx context.Context, in *RequestRange, opts ...grpc.CallOption) (Fabex_ExploreClient, error)
	GetByTxId(ctx context.Context, in *Entry, opts ...grpc.CallOption) (Fabex_GetByTxIdClient, error)
	GetByBlocknum(ctx context.Context, in *Entry, opts ...grpc.CallOption) (Fabex_GetByBlocknumClient, error)
	GetBlockInfoByPayload(ctx context.Context, in *Entry, opts ...grpc.CallOption) (Fabex_GetBlockInfoByPayloadClient, error)
}

type fabexClient struct {
	cc *grpc.ClientConn
}

func NewFabexClient(cc *grpc.ClientConn) FabexClient {
	return &fabexClient{cc}
}

func (c *fabexClient) Explore(ctx context.Context, in *RequestRange, opts ...grpc.CallOption) (Fabex_ExploreClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Fabex_serviceDesc.Streams[0], "/fabex.Fabex/Explore", opts...)
	if err != nil {
		return nil, err
	}
	x := &fabexExploreClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Fabex_ExploreClient interface {
	Recv() (*Entry, error)
	grpc.ClientStream
}

type fabexExploreClient struct {
	grpc.ClientStream
}

func (x *fabexExploreClient) Recv() (*Entry, error) {
	m := new(Entry)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fabexClient) GetByTxId(ctx context.Context, in *Entry, opts ...grpc.CallOption) (Fabex_GetByTxIdClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Fabex_serviceDesc.Streams[1], "/fabex.Fabex/GetByTxId", opts...)
	if err != nil {
		return nil, err
	}
	x := &fabexGetByTxIdClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Fabex_GetByTxIdClient interface {
	Recv() (*Entry, error)
	grpc.ClientStream
}

type fabexGetByTxIdClient struct {
	grpc.ClientStream
}

func (x *fabexGetByTxIdClient) Recv() (*Entry, error) {
	m := new(Entry)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fabexClient) GetByBlocknum(ctx context.Context, in *Entry, opts ...grpc.CallOption) (Fabex_GetByBlocknumClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Fabex_serviceDesc.Streams[2], "/fabex.Fabex/GetByBlocknum", opts...)
	if err != nil {
		return nil, err
	}
	x := &fabexGetByBlocknumClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Fabex_GetByBlocknumClient interface {
	Recv() (*Entry, error)
	grpc.ClientStream
}

type fabexGetByBlocknumClient struct {
	grpc.ClientStream
}

func (x *fabexGetByBlocknumClient) Recv() (*Entry, error) {
	m := new(Entry)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fabexClient) GetBlockInfoByPayload(ctx context.Context, in *Entry, opts ...grpc.CallOption) (Fabex_GetBlockInfoByPayloadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Fabex_serviceDesc.Streams[3], "/fabex.Fabex/GetBlockInfoByPayload", opts...)
	if err != nil {
		return nil, err
	}
	x := &fabexGetBlockInfoByPayloadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Fabex_GetBlockInfoByPayloadClient interface {
	Recv() (*Entry, error)
	grpc.ClientStream
}

type fabexGetBlockInfoByPayloadClient struct {
	grpc.ClientStream
}

func (x *fabexGetBlockInfoByPayloadClient) Recv() (*Entry, error) {
	m := new(Entry)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FabexServer is the server API for Fabex service.
type FabexServer interface {
	Explore(*RequestRange, Fabex_ExploreServer) error
	GetByTxId(*Entry, Fabex_GetByTxIdServer) error
	GetByBlocknum(*Entry, Fabex_GetByBlocknumServer) error
	GetBlockInfoByPayload(*Entry, Fabex_GetBlockInfoByPayloadServer) error
}

// UnimplementedFabexServer can be embedded to have forward compatible implementations.
type UnimplementedFabexServer struct {
}

func (*UnimplementedFabexServer) Explore(req *RequestRange, srv Fabex_ExploreServer) error {
	return status.Errorf(codes.Unimplemented, "method Explore not implemented")
}
func (*UnimplementedFabexServer) GetByTxId(req *Entry, srv Fabex_GetByTxIdServer) error {
	return status.Errorf(codes.Unimplemented, "method GetByTxId not implemented")
}
func (*UnimplementedFabexServer) GetByBlocknum(req *Entry, srv Fabex_GetByBlocknumServer) error {
	return status.Errorf(codes.Unimplemented, "method GetByBlocknum not implemented")
}
func (*UnimplementedFabexServer) GetBlockInfoByPayload(req *Entry, srv Fabex_GetBlockInfoByPayloadServer) error {
	return status.Errorf(codes.Unimplemented, "method GetBlockInfoByPayload not implemented")
}

func RegisterFabexServer(s *grpc.Server, srv FabexServer) {
	s.RegisterService(&_Fabex_serviceDesc, srv)
}

func _Fabex_Explore_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RequestRange)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FabexServer).Explore(m, &fabexExploreServer{stream})
}

type Fabex_ExploreServer interface {
	Send(*Entry) error
	grpc.ServerStream
}

type fabexExploreServer struct {
	grpc.ServerStream
}

func (x *fabexExploreServer) Send(m *Entry) error {
	return x.ServerStream.SendMsg(m)
}

func _Fabex_GetByTxId_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Entry)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FabexServer).GetByTxId(m, &fabexGetByTxIdServer{stream})
}

type Fabex_GetByTxIdServer interface {
	Send(*Entry) error
	grpc.ServerStream
}

type fabexGetByTxIdServer struct {
	grpc.ServerStream
}

func (x *fabexGetByTxIdServer) Send(m *Entry) error {
	return x.ServerStream.SendMsg(m)
}

func _Fabex_GetByBlocknum_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Entry)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FabexServer).GetByBlocknum(m, &fabexGetByBlocknumServer{stream})
}

type Fabex_GetByBlocknumServer interface {
	Send(*Entry) error
	grpc.ServerStream
}

type fabexGetByBlocknumServer struct {
	grpc.ServerStream
}

func (x *fabexGetByBlocknumServer) Send(m *Entry) error {
	return x.ServerStream.SendMsg(m)
}

func _Fabex_GetBlockInfoByPayload_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Entry)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FabexServer).GetBlockInfoByPayload(m, &fabexGetBlockInfoByPayloadServer{stream})
}

type Fabex_GetBlockInfoByPayloadServer interface {
	Send(*Entry) error
	grpc.ServerStream
}

type fabexGetBlockInfoByPayloadServer struct {
	grpc.ServerStream
}

func (x *fabexGetBlockInfoByPayloadServer) Send(m *Entry) error {
	return x.ServerStream.SendMsg(m)
}

var _Fabex_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fabex.Fabex",
	HandlerType: (*FabexServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Explore",
			Handler:       _Fabex_Explore_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetByTxId",
			Handler:       _Fabex_GetByTxId_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetByBlocknum",
			Handler:       _Fabex_GetByBlocknum_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetBlockInfoByPayload",
			Handler:       _Fabex_GetBlockInfoByPayload_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "fabex.proto",
}
