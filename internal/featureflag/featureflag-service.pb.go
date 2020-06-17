// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/featureflag-service.proto

package featureflag

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

// The request message containing the user's name.
type GetFlagRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	TenantId             string   `protobuf:"bytes,2,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFlagRequest) Reset()         { *m = GetFlagRequest{} }
func (m *GetFlagRequest) String() string { return proto.CompactTextString(m) }
func (*GetFlagRequest) ProtoMessage()    {}
func (*GetFlagRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fff38ecf66d7d015, []int{0}
}

func (m *GetFlagRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFlagRequest.Unmarshal(m, b)
}
func (m *GetFlagRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFlagRequest.Marshal(b, m, deterministic)
}
func (m *GetFlagRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFlagRequest.Merge(m, src)
}
func (m *GetFlagRequest) XXX_Size() int {
	return xxx_messageInfo_GetFlagRequest.Size(m)
}
func (m *GetFlagRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFlagRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFlagRequest proto.InternalMessageInfo

func (m *GetFlagRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *GetFlagRequest) GetTenantId() string {
	if m != nil {
		return m.TenantId
	}
	return ""
}

// The response message containing the greetings
type GetFlagResponse struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFlagResponse) Reset()         { *m = GetFlagResponse{} }
func (m *GetFlagResponse) String() string { return proto.CompactTextString(m) }
func (*GetFlagResponse) ProtoMessage()    {}
func (*GetFlagResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fff38ecf66d7d015, []int{1}
}

func (m *GetFlagResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFlagResponse.Unmarshal(m, b)
}
func (m *GetFlagResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFlagResponse.Marshal(b, m, deterministic)
}
func (m *GetFlagResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFlagResponse.Merge(m, src)
}
func (m *GetFlagResponse) XXX_Size() int {
	return xxx_messageInfo_GetFlagResponse.Size(m)
}
func (m *GetFlagResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFlagResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetFlagResponse proto.InternalMessageInfo

func (m *GetFlagResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*GetFlagRequest)(nil), "featureflag.GetFlagRequest")
	proto.RegisterType((*GetFlagResponse)(nil), "featureflag.GetFlagResponse")
}

func init() {
	proto.RegisterFile("protos/featureflag-service.proto", fileDescriptor_fff38ecf66d7d015)
}

var fileDescriptor_fff38ecf66d7d015 = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0x4f, 0x4b, 0x4d, 0x2c, 0x29, 0x2d, 0x4a, 0x4d, 0xcb, 0x49, 0x4c, 0xd7, 0x2d,
	0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x03, 0x4b, 0x09, 0x71, 0x23, 0x49, 0x29, 0xd9, 0x73,
	0xf1, 0xb9, 0xa7, 0x96, 0xb8, 0xe5, 0x24, 0xa6, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08,
	0x09, 0x70, 0x31, 0x67, 0xa7, 0x56, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0x98, 0x42,
	0xd2, 0x5c, 0x9c, 0x25, 0xa9, 0x79, 0x89, 0x79, 0x25, 0xf1, 0x99, 0x29, 0x12, 0x4c, 0x60, 0x71,
	0x0e, 0x88, 0x80, 0x67, 0x8a, 0x92, 0x3a, 0x17, 0x3f, 0xdc, 0x80, 0xe2, 0x82, 0xfc, 0xbc, 0xe2,
	0x54, 0x21, 0x11, 0x2e, 0xd6, 0xb2, 0xc4, 0x9c, 0xd2, 0x54, 0xa8, 0x19, 0x10, 0x8e, 0x51, 0x20,
	0x17, 0x7b, 0x30, 0xc4, 0x1d, 0x42, 0x6e, 0x5c, 0xec, 0x50, 0x3d, 0x42, 0xd2, 0x7a, 0x48, 0xae,
	0xd1, 0x43, 0x75, 0x8a, 0x94, 0x0c, 0x76, 0x49, 0x88, 0x35, 0x4a, 0x0c, 0x4e, 0x01, 0x5c, 0x1a,
	0x99, 0xf9, 0x7a, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc5, 0xd9, 0xc5, 0x7a, 0xb9, 0x99,
	0xc9, 0x45, 0xf9, 0x50, 0xdf, 0x16, 0x23, 0x6b, 0x2e, 0x76, 0xe2, 0x71, 0x83, 0xf0, 0x40, 0x46,
	0x14, 0x07, 0x30, 0x46, 0x89, 0x64, 0xe6, 0x95, 0xa4, 0x16, 0xe5, 0x25, 0xe6, 0x20, 0x87, 0x54,
	0x12, 0x1b, 0x38, 0x88, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd1, 0xfe, 0xc9, 0x2e, 0x46,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceClient interface {
	// get a flag value
	GetFlag(ctx context.Context, in *GetFlagRequest, opts ...grpc.CallOption) (*GetFlagResponse, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) GetFlag(ctx context.Context, in *GetFlagRequest, opts ...grpc.CallOption) (*GetFlagResponse, error) {
	out := new(GetFlagResponse)
	err := c.cc.Invoke(ctx, "/featureflag.Service/GetFlag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
type ServiceServer interface {
	// get a flag value
	GetFlag(context.Context, *GetFlagRequest) (*GetFlagResponse, error)
}

// UnimplementedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (*UnimplementedServiceServer) GetFlag(ctx context.Context, req *GetFlagRequest) (*GetFlagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFlag not implemented")
}

func RegisterServiceServer(s *grpc.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_GetFlag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFlagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetFlag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/featureflag.Service/GetFlag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetFlag(ctx, req.(*GetFlagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "featureflag.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFlag",
			Handler:    _Service_GetFlag_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/featureflag-service.proto",
}
