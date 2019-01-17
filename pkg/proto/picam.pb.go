// Code generated by protoc-gen-go. DO NOT EDIT.
// source: picam.proto

package rpi

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

type RequestOpen struct {
	Width                int32    `protobuf:"varint,1,opt,name=width,proto3" json:"width,omitempty"`
	Height               int32    `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	Rotation             int32    `protobuf:"varint,3,opt,name=rotation,proto3" json:"rotation,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestOpen) Reset()         { *m = RequestOpen{} }
func (m *RequestOpen) String() string { return proto.CompactTextString(m) }
func (*RequestOpen) ProtoMessage()    {}
func (*RequestOpen) Descriptor() ([]byte, []int) {
	return fileDescriptor_30fec5ee6d2f922d, []int{0}
}

func (m *RequestOpen) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestOpen.Unmarshal(m, b)
}
func (m *RequestOpen) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestOpen.Marshal(b, m, deterministic)
}
func (m *RequestOpen) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestOpen.Merge(m, src)
}
func (m *RequestOpen) XXX_Size() int {
	return xxx_messageInfo_RequestOpen.Size(m)
}
func (m *RequestOpen) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestOpen.DiscardUnknown(m)
}

var xxx_messageInfo_RequestOpen proto.InternalMessageInfo

func (m *RequestOpen) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *RequestOpen) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *RequestOpen) GetRotation() int32 {
	if m != nil {
		return m.Rotation
	}
	return 0
}

type ResponseImage struct {
	ImageBytes           []byte   `protobuf:"bytes,1,opt,name=imageBytes,proto3" json:"imageBytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseImage) Reset()         { *m = ResponseImage{} }
func (m *ResponseImage) String() string { return proto.CompactTextString(m) }
func (*ResponseImage) ProtoMessage()    {}
func (*ResponseImage) Descriptor() ([]byte, []int) {
	return fileDescriptor_30fec5ee6d2f922d, []int{1}
}

func (m *ResponseImage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseImage.Unmarshal(m, b)
}
func (m *ResponseImage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseImage.Marshal(b, m, deterministic)
}
func (m *ResponseImage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseImage.Merge(m, src)
}
func (m *ResponseImage) XXX_Size() int {
	return xxx_messageInfo_ResponseImage.Size(m)
}
func (m *ResponseImage) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseImage.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseImage proto.InternalMessageInfo

func (m *ResponseImage) GetImageBytes() []byte {
	if m != nil {
		return m.ImageBytes
	}
	return nil
}

func init() {
	proto.RegisterType((*RequestOpen)(nil), "rpi.RequestOpen")
	proto.RegisterType((*ResponseImage)(nil), "rpi.ResponseImage")
}

func init() { proto.RegisterFile("picam.proto", fileDescriptor_30fec5ee6d2f922d) }

var fileDescriptor_30fec5ee6d2f922d = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0xbb, 0xd6, 0x2d, 0x76, 0x5a, 0x41, 0x06, 0x91, 0x65, 0x0f, 0x2a, 0x01, 0xc1, 0x53,
	0x14, 0x7d, 0x03, 0x0b, 0x8a, 0x27, 0x25, 0x82, 0x9e, 0xd3, 0x76, 0xe8, 0x0e, 0x98, 0x4d, 0x4c,
	0x46, 0xc5, 0x17, 0xf0, 0xb9, 0xc5, 0xb4, 0x68, 0x7a, 0xcb, 0xff, 0xe5, 0xe7, 0xe7, 0x63, 0x60,
	0x12, 0x78, 0x61, 0x9d, 0x0e, 0xd1, 0x8b, 0xc7, 0x61, 0x0c, 0xdc, 0xc2, 0xdc, 0x26, 0x5a, 0x03,
	0xf5, 0x02, 0x13, 0x43, 0x6f, 0xef, 0x94, 0xe4, 0x21, 0x50, 0x8f, 0x87, 0x50, 0x7f, 0xf2, 0x52,
	0xba, 0xa6, 0x3a, 0xad, 0xce, 0x6b, 0xb3, 0x0e, 0x78, 0x04, 0xa3, 0x8e, 0x78, 0xd5, 0x49, 0xb3,
	0x93, 0xf1, 0x26, 0x61, 0x0b, 0x7b, 0xd1, 0x8b, 0x15, 0xf6, 0x7d, 0x33, 0xcc, 0x3f, 0x7f, 0x59,
	0x5d, 0xc0, 0xbe, 0xa1, 0x14, 0x7c, 0x9f, 0xe8, 0xde, 0xd9, 0x15, 0xe1, 0x31, 0x00, 0xff, 0x3e,
	0x6e, 0xbe, 0x84, 0x52, 0xde, 0x9f, 0x9a, 0x82, 0x5c, 0x7d, 0x57, 0x30, 0x7d, 0xe4, 0x99, 0x75,
	0x4f, 0x14, 0x3f, 0x78, 0x41, 0x78, 0x06, 0xbb, 0xd9, 0xe9, 0x40, 0xc7, 0xc0, 0xba, 0xb0, 0x6c,
	0xc7, 0x99, 0x3c, 0x7b, 0x5e, 0xaa, 0x01, 0x9e, 0x40, 0x3d, 0x7b, 0xf5, 0x89, 0xf0, 0x9f, 0x6e,
	0x17, 0x34, 0x8c, 0xef, 0x48, 0x6e, 0xa3, 0x75, 0x94, 0xca, 0x12, 0x6e, 0x76, 0x0b, 0x49, 0x35,
	0xb8, 0xac, 0xe6, 0xa3, 0x7c, 0x99, 0xeb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x4c, 0x74,
	0xae, 0x39, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PiCamServiceClient is the client API for PiCamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PiCamServiceClient interface {
	Open(ctx context.Context, in *RequestOpen, opts ...grpc.CallOption) (*Void, error)
	Close(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Void, error)
	GetFrames(ctx context.Context, in *Void, opts ...grpc.CallOption) (PiCamService_GetFramesClient, error)
}

type piCamServiceClient struct {
	cc *grpc.ClientConn
}

func NewPiCamServiceClient(cc *grpc.ClientConn) PiCamServiceClient {
	return &piCamServiceClient{cc}
}

func (c *piCamServiceClient) Open(ctx context.Context, in *RequestOpen, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/rpi.PiCamService/Open", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *piCamServiceClient) Close(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/rpi.PiCamService/Close", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *piCamServiceClient) GetFrames(ctx context.Context, in *Void, opts ...grpc.CallOption) (PiCamService_GetFramesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PiCamService_serviceDesc.Streams[0], "/rpi.PiCamService/GetFrames", opts...)
	if err != nil {
		return nil, err
	}
	x := &piCamServiceGetFramesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PiCamService_GetFramesClient interface {
	Recv() (*ResponseImage, error)
	grpc.ClientStream
}

type piCamServiceGetFramesClient struct {
	grpc.ClientStream
}

func (x *piCamServiceGetFramesClient) Recv() (*ResponseImage, error) {
	m := new(ResponseImage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PiCamServiceServer is the server API for PiCamService service.
type PiCamServiceServer interface {
	Open(context.Context, *RequestOpen) (*Void, error)
	Close(context.Context, *Void) (*Void, error)
	GetFrames(*Void, PiCamService_GetFramesServer) error
}

func RegisterPiCamServiceServer(s *grpc.Server, srv PiCamServiceServer) {
	s.RegisterService(&_PiCamService_serviceDesc, srv)
}

func _PiCamService_Open_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestOpen)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PiCamServiceServer).Open(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpi.PiCamService/Open",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PiCamServiceServer).Open(ctx, req.(*RequestOpen))
	}
	return interceptor(ctx, in, info, handler)
}

func _PiCamService_Close_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PiCamServiceServer).Close(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpi.PiCamService/Close",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PiCamServiceServer).Close(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _PiCamService_GetFrames_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Void)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PiCamServiceServer).GetFrames(m, &piCamServiceGetFramesServer{stream})
}

type PiCamService_GetFramesServer interface {
	Send(*ResponseImage) error
	grpc.ServerStream
}

type piCamServiceGetFramesServer struct {
	grpc.ServerStream
}

func (x *piCamServiceGetFramesServer) Send(m *ResponseImage) error {
	return x.ServerStream.SendMsg(m)
}

var _PiCamService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpi.PiCamService",
	HandlerType: (*PiCamServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Open",
			Handler:    _PiCamService_Open_Handler,
		},
		{
			MethodName: "Close",
			Handler:    _PiCamService_Close_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetFrames",
			Handler:       _PiCamService_GetFrames_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "picam.proto",
}
