// Code generated by protoc-gen-go. DO NOT EDIT.
// source: anraft/proto/peer_proto/peer.proto

package peer_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type ReqHeader struct {
	Version              int32    `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Seqid                int64    `protobuf:"varint,3,opt,name=seqid,proto3" json:"seqid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqHeader) Reset()         { *m = ReqHeader{} }
func (m *ReqHeader) String() string { return proto.CompactTextString(m) }
func (*ReqHeader) ProtoMessage()    {}
func (*ReqHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_peer_f592c43529ae4103, []int{0}
}
func (m *ReqHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqHeader.Unmarshal(m, b)
}
func (m *ReqHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqHeader.Marshal(b, m, deterministic)
}
func (dst *ReqHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqHeader.Merge(dst, src)
}
func (m *ReqHeader) XXX_Size() int {
	return xxx_messageInfo_ReqHeader.Size(m)
}
func (m *ReqHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqHeader.DiscardUnknown(m)
}

var xxx_messageInfo_ReqHeader proto.InternalMessageInfo

func (m *ReqHeader) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *ReqHeader) GetSeqid() int64 {
	if m != nil {
		return m.Seqid
	}
	return 0
}

type ResHeader struct {
	Seqid                int64    `protobuf:"varint,1,opt,name=seqid,proto3" json:"seqid,omitempty"`
	RetCode              int32    `protobuf:"varint,2,opt,name=ret_code,json=retCode,proto3" json:"ret_code,omitempty"`
	RetMsg               string   `protobuf:"bytes,4,opt,name=ret_msg,json=retMsg,proto3" json:"ret_msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResHeader) Reset()         { *m = ResHeader{} }
func (m *ResHeader) String() string { return proto.CompactTextString(m) }
func (*ResHeader) ProtoMessage()    {}
func (*ResHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_peer_f592c43529ae4103, []int{1}
}
func (m *ResHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResHeader.Unmarshal(m, b)
}
func (m *ResHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResHeader.Marshal(b, m, deterministic)
}
func (dst *ResHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResHeader.Merge(dst, src)
}
func (m *ResHeader) XXX_Size() int {
	return xxx_messageInfo_ResHeader.Size(m)
}
func (m *ResHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_ResHeader.DiscardUnknown(m)
}

var xxx_messageInfo_ResHeader proto.InternalMessageInfo

func (m *ResHeader) GetSeqid() int64 {
	if m != nil {
		return m.Seqid
	}
	return 0
}

func (m *ResHeader) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *ResHeader) GetRetMsg() string {
	if m != nil {
		return m.RetMsg
	}
	return ""
}

type HeartBeatReq struct {
	Header               *ReqHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *HeartBeatReq) Reset()         { *m = HeartBeatReq{} }
func (m *HeartBeatReq) String() string { return proto.CompactTextString(m) }
func (*HeartBeatReq) ProtoMessage()    {}
func (*HeartBeatReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_peer_f592c43529ae4103, []int{2}
}
func (m *HeartBeatReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartBeatReq.Unmarshal(m, b)
}
func (m *HeartBeatReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartBeatReq.Marshal(b, m, deterministic)
}
func (dst *HeartBeatReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartBeatReq.Merge(dst, src)
}
func (m *HeartBeatReq) XXX_Size() int {
	return xxx_messageInfo_HeartBeatReq.Size(m)
}
func (m *HeartBeatReq) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartBeatReq.DiscardUnknown(m)
}

var xxx_messageInfo_HeartBeatReq proto.InternalMessageInfo

func (m *HeartBeatReq) GetHeader() *ReqHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type HeartBeatRes struct {
	Header               *ResHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *HeartBeatRes) Reset()         { *m = HeartBeatRes{} }
func (m *HeartBeatRes) String() string { return proto.CompactTextString(m) }
func (*HeartBeatRes) ProtoMessage()    {}
func (*HeartBeatRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_peer_f592c43529ae4103, []int{3}
}
func (m *HeartBeatRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartBeatRes.Unmarshal(m, b)
}
func (m *HeartBeatRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartBeatRes.Marshal(b, m, deterministic)
}
func (dst *HeartBeatRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartBeatRes.Merge(dst, src)
}
func (m *HeartBeatRes) XXX_Size() int {
	return xxx_messageInfo_HeartBeatRes.Size(m)
}
func (m *HeartBeatRes) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartBeatRes.DiscardUnknown(m)
}

var xxx_messageInfo_HeartBeatRes proto.InternalMessageInfo

func (m *HeartBeatRes) GetHeader() *ResHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func init() {
	proto.RegisterType((*ReqHeader)(nil), "peer_proto.ReqHeader")
	proto.RegisterType((*ResHeader)(nil), "peer_proto.ResHeader")
	proto.RegisterType((*HeartBeatReq)(nil), "peer_proto.HeartBeatReq")
	proto.RegisterType((*HeartBeatRes)(nil), "peer_proto.HeartBeatRes")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PeerClient is the client API for Peer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PeerClient interface {
	HeartBeat(ctx context.Context, in *HeartBeatReq, opts ...grpc.CallOption) (*HeartBeatRes, error)
}

type peerClient struct {
	cc *grpc.ClientConn
}

func NewPeerClient(cc *grpc.ClientConn) PeerClient {
	return &peerClient{cc}
}

func (c *peerClient) HeartBeat(ctx context.Context, in *HeartBeatReq, opts ...grpc.CallOption) (*HeartBeatRes, error) {
	out := new(HeartBeatRes)
	err := c.cc.Invoke(ctx, "/peer_proto.Peer/HeartBeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PeerServer is the server API for Peer service.
type PeerServer interface {
	HeartBeat(context.Context, *HeartBeatReq) (*HeartBeatRes, error)
}

func RegisterPeerServer(s *grpc.Server, srv PeerServer) {
	s.RegisterService(&_Peer_serviceDesc, srv)
}

func _Peer_HeartBeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartBeatReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerServer).HeartBeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/peer_proto.Peer/HeartBeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerServer).HeartBeat(ctx, req.(*HeartBeatReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Peer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "peer_proto.Peer",
	HandlerType: (*PeerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HeartBeat",
			Handler:    _Peer_HeartBeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "anraft/proto/peer_proto/peer.proto",
}

func init() {
	proto.RegisterFile("anraft/proto/peer_proto/peer.proto", fileDescriptor_peer_f592c43529ae4103)
}

var fileDescriptor_peer_f592c43529ae4103 = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8f, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x59, 0xdb, 0x46, 0x33, 0x7a, 0x1a, 0x14, 0x57, 0x4f, 0x21, 0xa7, 0x5c, 0x4c, 0xa1,
	0x1e, 0x45, 0x04, 0x3d, 0xd8, 0x8b, 0x20, 0x0b, 0x9e, 0xc3, 0x6a, 0x9e, 0xb5, 0x07, 0xbb, 0xcd,
	0xcc, 0xe2, 0xef, 0x97, 0xa6, 0x35, 0x89, 0x82, 0xf4, 0xf6, 0x66, 0xf6, 0x7d, 0x6f, 0xdf, 0x50,
	0xee, 0x57, 0xe2, 0xdf, 0xe3, 0x74, 0x2d, 0x21, 0x86, 0xe9, 0x1a, 0x90, 0xaa, 0x97, 0x65, 0x2b,
	0x99, 0xfa, 0x75, 0x7e, 0x43, 0xa9, 0x43, 0x33, 0x87, 0xaf, 0x21, 0x6c, 0xe9, 0xf0, 0x0b, 0xa2,
	0xcb, 0xb0, 0xb2, 0x26, 0x33, 0xc5, 0xc4, 0xfd, 0x8c, 0x7c, 0x4a, 0x13, 0x45, 0xb3, 0xac, 0xed,
	0x28, 0x33, 0xc5, 0xc8, 0x6d, 0x87, 0xfc, 0x65, 0x03, 0xeb, 0x0e, 0xee, 0x2c, 0x66, 0x60, 0xe1,
	0x0b, 0x3a, 0x12, 0xc4, 0xea, 0x2d, 0xd4, 0xb0, 0x07, 0xdb, 0x4c, 0x41, 0x7c, 0x08, 0x35, 0xf8,
	0x9c, 0x36, 0xb2, 0xfa, 0xd4, 0x85, 0x1d, 0x67, 0xa6, 0x48, 0x5d, 0x22, 0x88, 0x4f, 0xba, 0xc8,
	0x6f, 0xe9, 0x64, 0x0e, 0x2f, 0xf1, 0x1e, 0x3e, 0x3a, 0x34, 0x7c, 0x45, 0xc9, 0x47, 0xfb, 0x47,
	0x1b, 0x7d, 0x3c, 0x3b, 0x2b, 0xfb, 0x03, 0xca, 0xae, 0xbd, 0xdb, 0x99, 0xfe, 0xe0, 0xba, 0x0f,
	0xd7, 0xdf, 0xf8, 0xec, 0x91, 0xc6, 0xcf, 0x80, 0xf0, 0x1d, 0xa5, 0x5d, 0x0c, 0xdb, 0x21, 0x33,
	0x2c, 0x77, 0xf9, 0xdf, 0x8b, 0xbe, 0x26, 0xed, 0xee, 0xfa, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x61,
	0xad, 0x49, 0x3a, 0x93, 0x01, 0x00, 0x00,
}
