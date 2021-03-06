// Code generated by protoc-gen-go. DO NOT EDIT.
// source: allstream.proto

//声明 包名

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

//stream 请求结构
type AllStreamReq struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AllStreamReq) Reset()         { *m = AllStreamReq{} }
func (m *AllStreamReq) String() string { return proto.CompactTextString(m) }
func (*AllStreamReq) ProtoMessage()    {}
func (*AllStreamReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_025094d0c4b09dae, []int{0}
}

func (m *AllStreamReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllStreamReq.Unmarshal(m, b)
}
func (m *AllStreamReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllStreamReq.Marshal(b, m, deterministic)
}
func (m *AllStreamReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllStreamReq.Merge(m, src)
}
func (m *AllStreamReq) XXX_Size() int {
	return xxx_messageInfo_AllStreamReq.Size(m)
}
func (m *AllStreamReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AllStreamReq.DiscardUnknown(m)
}

var xxx_messageInfo_AllStreamReq proto.InternalMessageInfo

func (m *AllStreamReq) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

//stream 返回结构
type AllStreamResp struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AllStreamResp) Reset()         { *m = AllStreamResp{} }
func (m *AllStreamResp) String() string { return proto.CompactTextString(m) }
func (*AllStreamResp) ProtoMessage()    {}
func (*AllStreamResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_025094d0c4b09dae, []int{1}
}

func (m *AllStreamResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllStreamResp.Unmarshal(m, b)
}
func (m *AllStreamResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllStreamResp.Marshal(b, m, deterministic)
}
func (m *AllStreamResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllStreamResp.Merge(m, src)
}
func (m *AllStreamResp) XXX_Size() int {
	return xxx_messageInfo_AllStreamResp.Size(m)
}
func (m *AllStreamResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AllStreamResp.DiscardUnknown(m)
}

var xxx_messageInfo_AllStreamResp proto.InternalMessageInfo

func (m *AllStreamResp) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*AllStreamReq)(nil), "helloworld.AllStreamReq")
	proto.RegisterType((*AllStreamResp)(nil), "helloworld.AllStreamResp")
}

func init() { proto.RegisterFile("allstream.proto", fileDescriptor_025094d0c4b09dae) }

var fileDescriptor_025094d0c4b09dae = []byte{
	// 137 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0xcc, 0xc9, 0x29,
	0x2e, 0x29, 0x4a, 0x4d, 0xcc, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xca, 0x48, 0xcd,
	0xc9, 0xc9, 0x2f, 0xcf, 0x2f, 0xca, 0x49, 0x51, 0x52, 0xe2, 0xe2, 0x71, 0xcc, 0xc9, 0x09, 0x06,
	0x4b, 0x07, 0xa5, 0x16, 0x0a, 0x09, 0x71, 0xb1, 0xa4, 0x24, 0x96, 0x24, 0x4a, 0x30, 0x2a, 0x30,
	0x6a, 0x70, 0x06, 0x81, 0xd9, 0x4a, 0xca, 0x5c, 0xbc, 0x48, 0x6a, 0x8a, 0x0b, 0xb0, 0x29, 0x32,
	0x8a, 0xe4, 0xe2, 0x87, 0x2b, 0x0a, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0x12, 0x72, 0xe3, 0xe2, 0x84,
	0x0b, 0x09, 0x49, 0xe8, 0x21, 0x6c, 0xd5, 0x43, 0xb6, 0x52, 0x4a, 0x12, 0x87, 0x4c, 0x71, 0x81,
	0x12, 0x83, 0x06, 0xa3, 0x01, 0x63, 0x12, 0x1b, 0xd8, 0xd9, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x91, 0xb3, 0x8c, 0x71, 0xc9, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AllStreamServerClient is the client API for AllStreamServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AllStreamServerClient interface {
	// 双向推送流
	AllStream(ctx context.Context, opts ...grpc.CallOption) (AllStreamServer_AllStreamClient, error)
}

type allStreamServerClient struct {
	cc *grpc.ClientConn
}

func NewAllStreamServerClient(cc *grpc.ClientConn) AllStreamServerClient {
	return &allStreamServerClient{cc}
}

func (c *allStreamServerClient) AllStream(ctx context.Context, opts ...grpc.CallOption) (AllStreamServer_AllStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AllStreamServer_serviceDesc.Streams[0], "/helloworld.AllStreamServer/AllStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &allStreamServerAllStreamClient{stream}
	return x, nil
}

type AllStreamServer_AllStreamClient interface {
	Send(*AllStreamReq) error
	Recv() (*AllStreamResp, error)
	grpc.ClientStream
}

type allStreamServerAllStreamClient struct {
	grpc.ClientStream
}

func (x *allStreamServerAllStreamClient) Send(m *AllStreamReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *allStreamServerAllStreamClient) Recv() (*AllStreamResp, error) {
	m := new(AllStreamResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AllStreamServerServer is the server API for AllStreamServer service.
type AllStreamServerServer interface {
	// 双向推送流
	AllStream(AllStreamServer_AllStreamServer) error
}

func RegisterAllStreamServerServer(s *grpc.Server, srv AllStreamServerServer) {
	s.RegisterService(&_AllStreamServer_serviceDesc, srv)
}

func _AllStreamServer_AllStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AllStreamServerServer).AllStream(&allStreamServerAllStreamServer{stream})
}

type AllStreamServer_AllStreamServer interface {
	Send(*AllStreamResp) error
	Recv() (*AllStreamReq, error)
	grpc.ServerStream
}

type allStreamServerAllStreamServer struct {
	grpc.ServerStream
}

func (x *allStreamServerAllStreamServer) Send(m *AllStreamResp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *allStreamServerAllStreamServer) Recv() (*AllStreamReq, error) {
	m := new(AllStreamReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _AllStreamServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.AllStreamServer",
	HandlerType: (*AllStreamServerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AllStream",
			Handler:       _AllStreamServer_AllStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "allstream.proto",
}
