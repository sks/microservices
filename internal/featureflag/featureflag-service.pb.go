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

type CreateFlagRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	TenantId             string   `protobuf:"bytes,3,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateFlagRequest) Reset()         { *m = CreateFlagRequest{} }
func (m *CreateFlagRequest) String() string { return proto.CompactTextString(m) }
func (*CreateFlagRequest) ProtoMessage()    {}
func (*CreateFlagRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fff38ecf66d7d015, []int{2}
}

func (m *CreateFlagRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateFlagRequest.Unmarshal(m, b)
}
func (m *CreateFlagRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateFlagRequest.Marshal(b, m, deterministic)
}
func (m *CreateFlagRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateFlagRequest.Merge(m, src)
}
func (m *CreateFlagRequest) XXX_Size() int {
	return xxx_messageInfo_CreateFlagRequest.Size(m)
}
func (m *CreateFlagRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateFlagRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateFlagRequest proto.InternalMessageInfo

func (m *CreateFlagRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *CreateFlagRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *CreateFlagRequest) GetTenantId() string {
	if m != nil {
		return m.TenantId
	}
	return ""
}

func (m *CreateFlagRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type CreateFlagResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateFlagResponse) Reset()         { *m = CreateFlagResponse{} }
func (m *CreateFlagResponse) String() string { return proto.CompactTextString(m) }
func (*CreateFlagResponse) ProtoMessage()    {}
func (*CreateFlagResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fff38ecf66d7d015, []int{3}
}

func (m *CreateFlagResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateFlagResponse.Unmarshal(m, b)
}
func (m *CreateFlagResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateFlagResponse.Marshal(b, m, deterministic)
}
func (m *CreateFlagResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateFlagResponse.Merge(m, src)
}
func (m *CreateFlagResponse) XXX_Size() int {
	return xxx_messageInfo_CreateFlagResponse.Size(m)
}
func (m *CreateFlagResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateFlagResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateFlagResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GetFlagRequest)(nil), "featureflag.GetFlagRequest")
	proto.RegisterType((*GetFlagResponse)(nil), "featureflag.GetFlagResponse")
	proto.RegisterType((*CreateFlagRequest)(nil), "featureflag.CreateFlagRequest")
	proto.RegisterType((*CreateFlagResponse)(nil), "featureflag.CreateFlagResponse")
}

func init() {
	proto.RegisterFile("protos/featureflag-service.proto", fileDescriptor_fff38ecf66d7d015)
}

var fileDescriptor_fff38ecf66d7d015 = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xcd, 0x4e, 0xf3, 0x30,
	0x10, 0xfc, 0xd2, 0x7e, 0x50, 0xba, 0x45, 0xfc, 0x58, 0x3d, 0x54, 0x2d, 0x82, 0x28, 0x17, 0x7a,
	0xc1, 0x48, 0xf0, 0x00, 0x48, 0x45, 0x0a, 0xe2, 0x44, 0x55, 0x6e, 0x5c, 0x90, 0x9b, 0x6c, 0x83,
	0xd5, 0x60, 0x07, 0xaf, 0x53, 0x09, 0x1e, 0x87, 0x27, 0x45, 0x8d, 0x43, 0x71, 0xc4, 0xcf, 0x2d,
	0xbb, 0x93, 0x9d, 0x99, 0x1d, 0x2f, 0x84, 0x85, 0xd1, 0x56, 0xd3, 0xf9, 0x02, 0x85, 0x2d, 0x0d,
	0x2e, 0x72, 0x91, 0x9d, 0x11, 0x9a, 0x95, 0x4c, 0x90, 0x57, 0x10, 0xeb, 0x79, 0x50, 0x74, 0x05,
	0x7b, 0x37, 0x68, 0xe3, 0x5c, 0x64, 0x33, 0x7c, 0x29, 0x91, 0x2c, 0x3b, 0x80, 0xf6, 0x12, 0x5f,
	0x07, 0x41, 0x18, 0x8c, 0xbb, 0xb3, 0xf5, 0x27, 0x1b, 0x41, 0xd7, 0xa2, 0x12, 0xca, 0x3e, 0xca,
	0x74, 0xd0, 0xaa, 0xfa, 0x3b, 0xae, 0x71, 0x9b, 0x46, 0xa7, 0xb0, 0xbf, 0x21, 0xa0, 0x42, 0x2b,
	0x42, 0xd6, 0x87, 0xad, 0x95, 0xc8, 0x4b, 0xac, 0x39, 0x5c, 0x11, 0xbd, 0xc1, 0xe1, 0xb5, 0x41,
	0x61, 0xf1, 0x6f, 0xb1, 0xcd, 0x70, 0xcb, 0x1b, 0x6e, 0x5a, 0x68, 0x37, 0x2d, 0xb0, 0x10, 0x7a,
	0x29, 0x52, 0x62, 0x64, 0x61, 0xa5, 0x56, 0x83, 0xff, 0x15, 0xec, 0xb7, 0xa2, 0x3e, 0x30, 0x5f,
	0xdb, 0xf9, 0xbc, 0x78, 0x0f, 0xa0, 0x73, 0xef, 0xa2, 0x61, 0x31, 0x74, 0xea, 0x35, 0xd8, 0x88,
	0x7b, 0x01, 0xf1, 0x66, 0x3a, 0xc3, 0xa3, 0x9f, 0x41, 0xc7, 0x18, 0xfd, 0x63, 0x77, 0x00, 0x5f,
	0x4a, 0xec, 0xb8, 0xf1, 0xf7, 0xb7, 0xf5, 0x87, 0x27, 0xbf, 0xe2, 0x9f, 0x84, 0x93, 0x29, 0x8c,
	0xa5, 0xe6, 0x99, 0xb4, 0x4f, 0xe5, 0x9c, 0xd3, 0x92, 0xf8, 0xb3, 0x4c, 0x8c, 0xae, 0x5f, 0x94,
	0xfc, 0x79, 0x9a, 0xec, 0xc6, 0xae, 0x5a, 0x53, 0xd0, 0x34, 0x78, 0xe8, 0x4b, 0x65, 0xd1, 0x28,
	0x91, 0xfb, 0xd7, 0x30, 0xdf, 0xae, 0xce, 0xe0, 0xf2, 0x23, 0x00, 0x00, 0xff, 0xff, 0x29, 0x40,
	0x6c, 0xfb, 0x2a, 0x02, 0x00, 0x00,
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
	CreateFlag(ctx context.Context, in *CreateFlagRequest, opts ...grpc.CallOption) (*CreateFlagResponse, error)
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

func (c *serviceClient) CreateFlag(ctx context.Context, in *CreateFlagRequest, opts ...grpc.CallOption) (*CreateFlagResponse, error) {
	out := new(CreateFlagResponse)
	err := c.cc.Invoke(ctx, "/featureflag.Service/CreateFlag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
type ServiceServer interface {
	// get a flag value
	GetFlag(context.Context, *GetFlagRequest) (*GetFlagResponse, error)
	CreateFlag(context.Context, *CreateFlagRequest) (*CreateFlagResponse, error)
}

// UnimplementedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (*UnimplementedServiceServer) GetFlag(ctx context.Context, req *GetFlagRequest) (*GetFlagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFlag not implemented")
}
func (*UnimplementedServiceServer) CreateFlag(ctx context.Context, req *CreateFlagRequest) (*CreateFlagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFlag not implemented")
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

func _Service_CreateFlag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFlagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).CreateFlag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/featureflag.Service/CreateFlag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).CreateFlag(ctx, req.(*CreateFlagRequest))
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
		{
			MethodName: "CreateFlag",
			Handler:    _Service_CreateFlag_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/featureflag-service.proto",
}
