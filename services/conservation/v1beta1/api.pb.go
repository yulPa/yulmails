// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package v1beta1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

// Conservation is database representation
type Conservation struct {
	// database unique identifier
	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	// creation date (timestamp without time zone)
	Created string `protobuf:"bytes,2,opt,name=created,proto3" json:"created,omitempty"`
	// number of sent email to keep
	Sent int64 `protobuf:"varint,3,opt,name=sent,proto3" json:"sent,omitempty"`
	// number of unsent email to keep
	Unsent int64 `protobuf:"varint,4,opt,name=unsent,proto3" json:"unsent,omitempty"`
	// keep email content (depend of country law)
	KeepEmailContent     bool     `protobuf:"varint,5,opt,name=keep_email_content,json=keepEmailContent,proto3" json:"keep_email_content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Conservation) Reset()         { *m = Conservation{} }
func (m *Conservation) String() string { return proto.CompactTextString(m) }
func (*Conservation) ProtoMessage()    {}
func (*Conservation) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *Conservation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Conservation.Unmarshal(m, b)
}
func (m *Conservation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Conservation.Marshal(b, m, deterministic)
}
func (m *Conservation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Conservation.Merge(m, src)
}
func (m *Conservation) XXX_Size() int {
	return xxx_messageInfo_Conservation.Size(m)
}
func (m *Conservation) XXX_DiscardUnknown() {
	xxx_messageInfo_Conservation.DiscardUnknown(m)
}

var xxx_messageInfo_Conservation proto.InternalMessageInfo

func (m *Conservation) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Conservation) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *Conservation) GetSent() int64 {
	if m != nil {
		return m.Sent
	}
	return 0
}

func (m *Conservation) GetUnsent() int64 {
	if m != nil {
		return m.Unsent
	}
	return 0
}

func (m *Conservation) GetKeepEmailContent() bool {
	if m != nil {
		return m.KeepEmailContent
	}
	return false
}

func init() {
	proto.RegisterType((*Conservation)(nil), "v1beta1.Conservation")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x50, 0xcf, 0x4a, 0xc3, 0x30,
	0x1c, 0x26, 0xdd, 0xdc, 0x5c, 0x10, 0x29, 0x3f, 0xd9, 0x28, 0x9d, 0x87, 0xd2, 0x53, 0x11, 0x49,
	0xad, 0xde, 0xbc, 0x76, 0x3b, 0x0c, 0x3c, 0xd5, 0xb3, 0x94, 0xb4, 0xfe, 0x1c, 0xc1, 0x2d, 0x29,
	0x6d, 0x56, 0xf0, 0xea, 0x23, 0xe8, 0xa3, 0x89, 0x6f, 0xe0, 0x83, 0x48, 0xd3, 0x0e, 0x3a, 0x74,
	0xb7, 0x7c, 0x7f, 0xf2, 0xe5, 0xfb, 0x42, 0x27, 0xbc, 0x10, 0xac, 0x28, 0x95, 0x56, 0x30, 0xae,
	0xa3, 0x0c, 0x35, 0x8f, 0xdc, 0xcb, 0xb5, 0x52, 0xeb, 0x0d, 0x86, 0xbc, 0x10, 0x21, 0x97, 0x52,
	0x69, 0xae, 0x85, 0x92, 0x55, 0x6b, 0x73, 0xe7, 0x9d, 0x6a, 0x50, 0xb6, 0x7b, 0x09, 0x71, 0x5b,
	0xe8, 0xb7, 0x56, 0xf4, 0x3f, 0x08, 0x3d, 0x8b, 0x95, 0xac, 0xb0, 0xac, 0xcd, 0x25, 0x38, 0xa7,
	0xd6, 0x6a, 0xe1, 0x10, 0x8f, 0x04, 0x83, 0xc4, 0x5a, 0x2d, 0xc0, 0xa1, 0xe3, 0xbc, 0x44, 0xae,
	0xf1, 0xd9, 0xb1, 0x3c, 0x12, 0x4c, 0x92, 0x3d, 0x04, 0xa0, 0xc3, 0x0a, 0xa5, 0x76, 0x06, 0xc6,
	0x6b, 0xce, 0x30, 0xa3, 0xa3, 0x9d, 0x34, 0xec, 0xd0, 0xb0, 0x1d, 0x82, 0x6b, 0x0a, 0xaf, 0x88,
	0x45, 0x8a, 0x5b, 0x2e, 0x36, 0x69, 0xae, 0xa4, 0x6e, 0x3c, 0x27, 0x1e, 0x09, 0x4e, 0x13, 0xbb,
	0x51, 0x96, 0x8d, 0x10, 0xb7, 0xfc, 0xed, 0x37, 0xa1, 0x17, 0xfd, 0x52, 0x8f, 0x58, 0xd6, 0x22,
	0x47, 0x78, 0xa2, 0xf6, 0x83, 0xa8, 0xf4, 0x41, 0xdf, 0x19, 0x6b, 0xe7, 0xb1, 0xfd, 0x3c, 0xb6,
	0x6c, 0xe6, 0xb9, 0x53, 0xd6, 0xfd, 0x0e, 0xeb, 0xdb, 0x7d, 0xe7, 0xfd, 0xeb, 0xe7, 0xd3, 0x02,
	0xb0, 0xc3, 0x3a, 0x0a, 0xf3, 0x9e, 0x72, 0x43, 0x20, 0xa5, 0x10, 0x9b, 0x6d, 0x07, 0x0f, 0xfc,
	0x1f, 0x74, 0x2c, 0x7f, 0x6e, 0xf2, 0xa7, 0xfe, 0x9f, 0xfc, 0x7b, 0x72, 0x95, 0x8d, 0x4c, 0xc7,
	0xbb, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x39, 0xc4, 0x5a, 0x59, 0xc4, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ConservationServiceClient is the client API for ConservationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConservationServiceClient interface {
	// ListConservation returns a stream of List of Conservation
	ListConservation(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (ConservationService_ListConservationClient, error)
	// CreateConservation insert a conservation in dataset
	CreateConservation(ctx context.Context, in *Conservation, opts ...grpc.CallOption) (*Conservation, error)
}

type conservationServiceClient struct {
	cc *grpc.ClientConn
}

func NewConservationServiceClient(cc *grpc.ClientConn) ConservationServiceClient {
	return &conservationServiceClient{cc}
}

func (c *conservationServiceClient) ListConservation(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (ConservationService_ListConservationClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ConservationService_serviceDesc.Streams[0], "/v1beta1.ConservationService/ListConservation", opts...)
	if err != nil {
		return nil, err
	}
	x := &conservationServiceListConservationClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ConservationService_ListConservationClient interface {
	Recv() (*Conservation, error)
	grpc.ClientStream
}

type conservationServiceListConservationClient struct {
	grpc.ClientStream
}

func (x *conservationServiceListConservationClient) Recv() (*Conservation, error) {
	m := new(Conservation)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *conservationServiceClient) CreateConservation(ctx context.Context, in *Conservation, opts ...grpc.CallOption) (*Conservation, error) {
	out := new(Conservation)
	err := c.cc.Invoke(ctx, "/v1beta1.ConservationService/CreateConservation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConservationServiceServer is the server API for ConservationService service.
type ConservationServiceServer interface {
	// ListConservation returns a stream of List of Conservation
	ListConservation(*empty.Empty, ConservationService_ListConservationServer) error
	// CreateConservation insert a conservation in dataset
	CreateConservation(context.Context, *Conservation) (*Conservation, error)
}

func RegisterConservationServiceServer(s *grpc.Server, srv ConservationServiceServer) {
	s.RegisterService(&_ConservationService_serviceDesc, srv)
}

func _ConservationService_ListConservation_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ConservationServiceServer).ListConservation(m, &conservationServiceListConservationServer{stream})
}

type ConservationService_ListConservationServer interface {
	Send(*Conservation) error
	grpc.ServerStream
}

type conservationServiceListConservationServer struct {
	grpc.ServerStream
}

func (x *conservationServiceListConservationServer) Send(m *Conservation) error {
	return x.ServerStream.SendMsg(m)
}

func _ConservationService_CreateConservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Conservation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConservationServiceServer).CreateConservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1beta1.ConservationService/CreateConservation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConservationServiceServer).CreateConservation(ctx, req.(*Conservation))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConservationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1beta1.ConservationService",
	HandlerType: (*ConservationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateConservation",
			Handler:    _ConservationService_CreateConservation_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListConservation",
			Handler:       _ConservationService_ListConservation_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api.proto",
}
