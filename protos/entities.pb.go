// Code generated by protoc-gen-go. DO NOT EDIT.
// source: entities.proto

package protos

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

type CreateUserRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Firstname            string   `protobuf:"bytes,3,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname             string   `protobuf:"bytes,4,opt,name=lastname,proto3" json:"lastname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5f7344b79fcc63d, []int{0}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CreateUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *CreateUserRequest) GetFirstname() string {
	if m != nil {
		return m.Firstname
	}
	return ""
}

func (m *CreateUserRequest) GetLastname() string {
	if m != nil {
		return m.Lastname
	}
	return ""
}

type CreateUserReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserReply) Reset()         { *m = CreateUserReply{} }
func (m *CreateUserReply) String() string { return proto.CompactTextString(m) }
func (*CreateUserReply) ProtoMessage()    {}
func (*CreateUserReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5f7344b79fcc63d, []int{1}
}

func (m *CreateUserReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserReply.Unmarshal(m, b)
}
func (m *CreateUserReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserReply.Marshal(b, m, deterministic)
}
func (m *CreateUserReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserReply.Merge(m, src)
}
func (m *CreateUserReply) XXX_Size() int {
	return xxx_messageInfo_CreateUserReply.Size(m)
}
func (m *CreateUserReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserReply.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserReply proto.InternalMessageInfo

type GetUserRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5f7344b79fcc63d, []int{2}
}

func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (m *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(m, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

func (m *GetUserRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetUserAuthenticateRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserAuthenticateRequest) Reset()         { *m = GetUserAuthenticateRequest{} }
func (m *GetUserAuthenticateRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserAuthenticateRequest) ProtoMessage()    {}
func (*GetUserAuthenticateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5f7344b79fcc63d, []int{3}
}

func (m *GetUserAuthenticateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserAuthenticateRequest.Unmarshal(m, b)
}
func (m *GetUserAuthenticateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserAuthenticateRequest.Marshal(b, m, deterministic)
}
func (m *GetUserAuthenticateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserAuthenticateRequest.Merge(m, src)
}
func (m *GetUserAuthenticateRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserAuthenticateRequest.Size(m)
}
func (m *GetUserAuthenticateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserAuthenticateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserAuthenticateRequest proto.InternalMessageInfo

func (m *GetUserAuthenticateRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *GetUserAuthenticateRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type UserData struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Firstname            string   `protobuf:"bytes,3,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname             string   `protobuf:"bytes,4,opt,name=lastname,proto3" json:"lastname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserData) Reset()         { *m = UserData{} }
func (m *UserData) String() string { return proto.CompactTextString(m) }
func (*UserData) ProtoMessage()    {}
func (*UserData) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5f7344b79fcc63d, []int{4}
}

func (m *UserData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserData.Unmarshal(m, b)
}
func (m *UserData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserData.Marshal(b, m, deterministic)
}
func (m *UserData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserData.Merge(m, src)
}
func (m *UserData) XXX_Size() int {
	return xxx_messageInfo_UserData.Size(m)
}
func (m *UserData) XXX_DiscardUnknown() {
	xxx_messageInfo_UserData.DiscardUnknown(m)
}

var xxx_messageInfo_UserData proto.InternalMessageInfo

func (m *UserData) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserData) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserData) GetFirstname() string {
	if m != nil {
		return m.Firstname
	}
	return ""
}

func (m *UserData) GetLastname() string {
	if m != nil {
		return m.Lastname
	}
	return ""
}

type GetBookRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBookRequest) Reset()         { *m = GetBookRequest{} }
func (m *GetBookRequest) String() string { return proto.CompactTextString(m) }
func (*GetBookRequest) ProtoMessage()    {}
func (*GetBookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5f7344b79fcc63d, []int{5}
}

func (m *GetBookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBookRequest.Unmarshal(m, b)
}
func (m *GetBookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBookRequest.Marshal(b, m, deterministic)
}
func (m *GetBookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBookRequest.Merge(m, src)
}
func (m *GetBookRequest) XXX_Size() int {
	return xxx_messageInfo_GetBookRequest.Size(m)
}
func (m *GetBookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBookRequest proto.InternalMessageInfo

type GetBookReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBookReply) Reset()         { *m = GetBookReply{} }
func (m *GetBookReply) String() string { return proto.CompactTextString(m) }
func (*GetBookReply) ProtoMessage()    {}
func (*GetBookReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5f7344b79fcc63d, []int{6}
}

func (m *GetBookReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBookReply.Unmarshal(m, b)
}
func (m *GetBookReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBookReply.Marshal(b, m, deterministic)
}
func (m *GetBookReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBookReply.Merge(m, src)
}
func (m *GetBookReply) XXX_Size() int {
	return xxx_messageInfo_GetBookReply.Size(m)
}
func (m *GetBookReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBookReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetBookReply proto.InternalMessageInfo

type CreateBookRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBookRequest) Reset()         { *m = CreateBookRequest{} }
func (m *CreateBookRequest) String() string { return proto.CompactTextString(m) }
func (*CreateBookRequest) ProtoMessage()    {}
func (*CreateBookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5f7344b79fcc63d, []int{7}
}

func (m *CreateBookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBookRequest.Unmarshal(m, b)
}
func (m *CreateBookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBookRequest.Marshal(b, m, deterministic)
}
func (m *CreateBookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBookRequest.Merge(m, src)
}
func (m *CreateBookRequest) XXX_Size() int {
	return xxx_messageInfo_CreateBookRequest.Size(m)
}
func (m *CreateBookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBookRequest proto.InternalMessageInfo

type CreateBookReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBookReply) Reset()         { *m = CreateBookReply{} }
func (m *CreateBookReply) String() string { return proto.CompactTextString(m) }
func (*CreateBookReply) ProtoMessage()    {}
func (*CreateBookReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5f7344b79fcc63d, []int{8}
}

func (m *CreateBookReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBookReply.Unmarshal(m, b)
}
func (m *CreateBookReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBookReply.Marshal(b, m, deterministic)
}
func (m *CreateBookReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBookReply.Merge(m, src)
}
func (m *CreateBookReply) XXX_Size() int {
	return xxx_messageInfo_CreateBookReply.Size(m)
}
func (m *CreateBookReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBookReply.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBookReply proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreateUserRequest)(nil), "protos.CreateUserRequest")
	proto.RegisterType((*CreateUserReply)(nil), "protos.CreateUserReply")
	proto.RegisterType((*GetUserRequest)(nil), "protos.GetUserRequest")
	proto.RegisterType((*GetUserAuthenticateRequest)(nil), "protos.GetUserAuthenticateRequest")
	proto.RegisterType((*UserData)(nil), "protos.UserData")
	proto.RegisterType((*GetBookRequest)(nil), "protos.GetBookRequest")
	proto.RegisterType((*GetBookReply)(nil), "protos.GetBookReply")
	proto.RegisterType((*CreateBookRequest)(nil), "protos.CreateBookRequest")
	proto.RegisterType((*CreateBookReply)(nil), "protos.CreateBookReply")
}

func init() { proto.RegisterFile("entities.proto", fileDescriptor_a5f7344b79fcc63d) }

var fileDescriptor_a5f7344b79fcc63d = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xcb, 0x4e, 0x32, 0x31,
	0x14, 0x0e, 0x03, 0xff, 0x0f, 0x9c, 0x98, 0x11, 0x0a, 0x11, 0x9c, 0xb8, 0x20, 0x5d, 0xb9, 0x62,
	0x01, 0x0b, 0xb7, 0x22, 0x26, 0xc6, 0x98, 0x18, 0x43, 0xe2, 0x03, 0x54, 0x39, 0xc6, 0xea, 0x40,
	0xc7, 0xf6, 0x10, 0xc3, 0xca, 0xd7, 0xf3, 0xb1, 0x4c, 0x3b, 0x17, 0xca, 0x30, 0x2b, 0x56, 0x93,
	0x73, 0xfb, 0xe6, 0xbb, 0xa4, 0x10, 0xe2, 0x9a, 0x24, 0x49, 0x34, 0xe3, 0x44, 0x2b, 0x52, 0xec,
	0xbf, 0xfb, 0x18, 0xfe, 0x03, 0xdd, 0xb9, 0x46, 0x41, 0xf8, 0x6c, 0x50, 0x2f, 0xf0, 0x6b, 0x83,
	0x86, 0x58, 0x1f, 0xfe, 0xe1, 0x4a, 0xc8, 0x78, 0x58, 0x1b, 0xd5, 0x2e, 0xdb, 0x8b, 0xb4, 0x60,
	0x11, 0xb4, 0x12, 0x61, 0xcc, 0xb7, 0xd2, 0xcb, 0x61, 0xe0, 0x06, 0x45, 0xcd, 0x2e, 0xa0, 0xfd,
	0x26, 0xb5, 0xa1, 0xb5, 0x58, 0xe1, 0xb0, 0xee, 0x86, 0xbb, 0x86, 0xbd, 0x8c, 0x45, 0x36, 0x6c,
	0xa4, 0x97, 0x79, 0xcd, 0xbb, 0x70, 0xea, 0x13, 0x48, 0xe2, 0x2d, 0x1f, 0x41, 0x78, 0x87, 0xe4,
	0x13, 0x0a, 0x21, 0x90, 0x4b, 0xc7, 0xa6, 0xbe, 0x08, 0xe4, 0x92, 0x3f, 0x42, 0x94, 0x6d, 0xcc,
	0x36, 0xf4, 0x6e, 0xa5, 0xbd, 0x0a, 0xc2, 0xa3, 0xe9, 0xf3, 0x0f, 0x68, 0x59, 0xb0, 0x5b, 0x41,
	0xa2, 0xfc, 0xaf, 0x1d, 0x5a, 0xe0, 0xa3, 0x1d, 0x2f, 0xb8, 0xe3, 0xd4, 0xdd, 0x28, 0xf5, 0x99,
	0xf1, 0xe5, 0x21, 0x9c, 0x14, 0x1d, 0xab, 0xbf, 0x97, 0x67, 0xe2, 0x2f, 0x15, 0x3e, 0x15, 0x7b,
	0x93, 0xdf, 0x00, 0x9a, 0x4f, 0xa8, 0x8d, 0x34, 0xc4, 0xae, 0x01, 0x76, 0x36, 0xb2, 0xf3, 0x34,
	0x65, 0x33, 0x3e, 0xc8, 0x36, 0x1a, 0x54, 0x8d, 0x92, 0x78, 0xcb, 0xa6, 0xd0, 0xcc, 0x3c, 0x65,
	0x67, 0xf9, 0xce, 0x7e, 0x0c, 0x51, 0x27, 0xef, 0x17, 0x66, 0x3d, 0x40, 0xaf, 0x22, 0x08, 0xc6,
	0x4b, 0x00, 0x15, 0x29, 0x55, 0x80, 0x5d, 0x39, 0x06, 0x56, 0xdf, 0x1e, 0x03, 0xcf, 0x85, 0xa8,
	0x7f, 0xd0, 0xb7, 0xd4, 0x0b, 0xf1, 0xee, 0xb6, 0x24, 0xde, 0x3f, 0x1f, 0x54, 0x8d, 0xac, 0x95,
	0x6d, 0x68, 0xce, 0xd5, 0x9a, 0xb4, 0x8a, 0x27, 0x0d, 0x08, 0x66, 0xf7, 0x2f, 0xe9, 0xfb, 0x98,
	0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0x80, 0x52, 0xeb, 0x27, 0x38, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PersistClient is the client API for Persist service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PersistClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserReply, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*UserData, error)
	GetUserAuthenticate(ctx context.Context, in *GetUserAuthenticateRequest, opts ...grpc.CallOption) (*UserData, error)
	GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*GetBookReply, error)
	CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*CreateBookReply, error)
}

type persistClient struct {
	cc *grpc.ClientConn
}

func NewPersistClient(cc *grpc.ClientConn) PersistClient {
	return &persistClient{cc}
}

func (c *persistClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserReply, error) {
	out := new(CreateUserReply)
	err := c.cc.Invoke(ctx, "/protos.Persist/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *persistClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*UserData, error) {
	out := new(UserData)
	err := c.cc.Invoke(ctx, "/protos.Persist/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *persistClient) GetUserAuthenticate(ctx context.Context, in *GetUserAuthenticateRequest, opts ...grpc.CallOption) (*UserData, error) {
	out := new(UserData)
	err := c.cc.Invoke(ctx, "/protos.Persist/GetUserAuthenticate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *persistClient) GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*GetBookReply, error) {
	out := new(GetBookReply)
	err := c.cc.Invoke(ctx, "/protos.Persist/GetBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *persistClient) CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*CreateBookReply, error) {
	out := new(CreateBookReply)
	err := c.cc.Invoke(ctx, "/protos.Persist/CreateBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PersistServer is the server API for Persist service.
type PersistServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserReply, error)
	GetUser(context.Context, *GetUserRequest) (*UserData, error)
	GetUserAuthenticate(context.Context, *GetUserAuthenticateRequest) (*UserData, error)
	GetBook(context.Context, *GetBookRequest) (*GetBookReply, error)
	CreateBook(context.Context, *CreateBookRequest) (*CreateBookReply, error)
}

// UnimplementedPersistServer can be embedded to have forward compatible implementations.
type UnimplementedPersistServer struct {
}

func (*UnimplementedPersistServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedPersistServer) GetUser(ctx context.Context, req *GetUserRequest) (*UserData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (*UnimplementedPersistServer) GetUserAuthenticate(ctx context.Context, req *GetUserAuthenticateRequest) (*UserData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserAuthenticate not implemented")
}
func (*UnimplementedPersistServer) GetBook(ctx context.Context, req *GetBookRequest) (*GetBookReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBook not implemented")
}
func (*UnimplementedPersistServer) CreateBook(ctx context.Context, req *CreateBookRequest) (*CreateBookReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBook not implemented")
}

func RegisterPersistServer(s *grpc.Server, srv PersistServer) {
	s.RegisterService(&_Persist_serviceDesc, srv)
}

func _Persist_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersistServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Persist/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersistServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Persist_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersistServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Persist/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersistServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Persist_GetUserAuthenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserAuthenticateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersistServer).GetUserAuthenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Persist/GetUserAuthenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersistServer).GetUserAuthenticate(ctx, req.(*GetUserAuthenticateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Persist_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersistServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Persist/GetBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersistServer).GetBook(ctx, req.(*GetBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Persist_CreateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersistServer).CreateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Persist/CreateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersistServer).CreateBook(ctx, req.(*CreateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Persist_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Persist",
	HandlerType: (*PersistServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _Persist_CreateUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _Persist_GetUser_Handler,
		},
		{
			MethodName: "GetUserAuthenticate",
			Handler:    _Persist_GetUserAuthenticate_Handler,
		},
		{
			MethodName: "GetBook",
			Handler:    _Persist_GetBook_Handler,
		},
		{
			MethodName: "CreateBook",
			Handler:    _Persist_CreateBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "entities.proto",
}

// ControlClient is the client API for Control service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ControlClient interface {
}

type controlClient struct {
	cc *grpc.ClientConn
}

func NewControlClient(cc *grpc.ClientConn) ControlClient {
	return &controlClient{cc}
}

// ControlServer is the server API for Control service.
type ControlServer interface {
}

// UnimplementedControlServer can be embedded to have forward compatible implementations.
type UnimplementedControlServer struct {
}

func RegisterControlServer(s *grpc.Server, srv ControlServer) {
	s.RegisterService(&_Control_serviceDesc, srv)
}

var _Control_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Control",
	HandlerType: (*ControlServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "entities.proto",
}

// AIClient is the client API for AI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AIClient interface {
}

type aIClient struct {
	cc *grpc.ClientConn
}

func NewAIClient(cc *grpc.ClientConn) AIClient {
	return &aIClient{cc}
}

// AIServer is the server API for AI service.
type AIServer interface {
}

// UnimplementedAIServer can be embedded to have forward compatible implementations.
type UnimplementedAIServer struct {
}

func RegisterAIServer(s *grpc.Server, srv AIServer) {
	s.RegisterService(&_AI_serviceDesc, srv)
}

var _AI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.AI",
	HandlerType: (*AIServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "entities.proto",
}