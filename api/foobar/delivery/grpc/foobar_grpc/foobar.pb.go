// Code generated by protoc-gen-go. DO NOT EDIT.
// source: foobar.proto

package foobar_grpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Foobar struct {
	ID                   uint64               `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	FoobarContent        string               `protobuf:"bytes,2,opt,name=FoobarContent,proto3" json:"FoobarContent,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Foobar) Reset()         { *m = Foobar{} }
func (m *Foobar) String() string { return proto.CompactTextString(m) }
func (*Foobar) ProtoMessage()    {}
func (*Foobar) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fe8bbf5351ca4f2, []int{0}
}

func (m *Foobar) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Foobar.Unmarshal(m, b)
}
func (m *Foobar) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Foobar.Marshal(b, m, deterministic)
}
func (m *Foobar) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Foobar.Merge(m, src)
}
func (m *Foobar) XXX_Size() int {
	return xxx_messageInfo_Foobar.Size(m)
}
func (m *Foobar) XXX_DiscardUnknown() {
	xxx_messageInfo_Foobar.DiscardUnknown(m)
}

var xxx_messageInfo_Foobar proto.InternalMessageInfo

func (m *Foobar) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Foobar) GetFoobarContent() string {
	if m != nil {
		return m.FoobarContent
	}
	return ""
}

func (m *Foobar) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Foobar) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type ListFoobar struct {
	Foobar               []*Foobar `protobuf:"bytes,1,rep,name=foobar,proto3" json:"foobar,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListFoobar) Reset()         { *m = ListFoobar{} }
func (m *ListFoobar) String() string { return proto.CompactTextString(m) }
func (*ListFoobar) ProtoMessage()    {}
func (*ListFoobar) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fe8bbf5351ca4f2, []int{1}
}

func (m *ListFoobar) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListFoobar.Unmarshal(m, b)
}
func (m *ListFoobar) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListFoobar.Marshal(b, m, deterministic)
}
func (m *ListFoobar) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListFoobar.Merge(m, src)
}
func (m *ListFoobar) XXX_Size() int {
	return xxx_messageInfo_ListFoobar.Size(m)
}
func (m *ListFoobar) XXX_DiscardUnknown() {
	xxx_messageInfo_ListFoobar.DiscardUnknown(m)
}

var xxx_messageInfo_ListFoobar proto.InternalMessageInfo

func (m *ListFoobar) GetFoobar() []*Foobar {
	if m != nil {
		return m.Foobar
	}
	return nil
}

type DeleteResponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Code                 uint64   `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fe8bbf5351ca4f2, []int{2}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func (m *DeleteResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *DeleteResponse) GetCode() uint64 {
	if m != nil {
		return m.Code
	}
	return 0
}

type FetchRequest struct {
	Num                  uint64   `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchRequest) Reset()         { *m = FetchRequest{} }
func (m *FetchRequest) String() string { return proto.CompactTextString(m) }
func (*FetchRequest) ProtoMessage()    {}
func (*FetchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fe8bbf5351ca4f2, []int{3}
}

func (m *FetchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchRequest.Unmarshal(m, b)
}
func (m *FetchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchRequest.Marshal(b, m, deterministic)
}
func (m *FetchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchRequest.Merge(m, src)
}
func (m *FetchRequest) XXX_Size() int {
	return xxx_messageInfo_FetchRequest.Size(m)
}
func (m *FetchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FetchRequest proto.InternalMessageInfo

func (m *FetchRequest) GetNum() uint64 {
	if m != nil {
		return m.Num
	}
	return 0
}

type SingleRequest struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SingleRequest) Reset()         { *m = SingleRequest{} }
func (m *SingleRequest) String() string { return proto.CompactTextString(m) }
func (*SingleRequest) ProtoMessage()    {}
func (*SingleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fe8bbf5351ca4f2, []int{4}
}

func (m *SingleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SingleRequest.Unmarshal(m, b)
}
func (m *SingleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SingleRequest.Marshal(b, m, deterministic)
}
func (m *SingleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SingleRequest.Merge(m, src)
}
func (m *SingleRequest) XXX_Size() int {
	return xxx_messageInfo_SingleRequest.Size(m)
}
func (m *SingleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SingleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SingleRequest proto.InternalMessageInfo

func (m *SingleRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ErrorMessage struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ErrorMessage) Reset()         { *m = ErrorMessage{} }
func (m *ErrorMessage) String() string { return proto.CompactTextString(m) }
func (*ErrorMessage) ProtoMessage()    {}
func (*ErrorMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fe8bbf5351ca4f2, []int{5}
}

func (m *ErrorMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorMessage.Unmarshal(m, b)
}
func (m *ErrorMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorMessage.Marshal(b, m, deterministic)
}
func (m *ErrorMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorMessage.Merge(m, src)
}
func (m *ErrorMessage) XXX_Size() int {
	return xxx_messageInfo_ErrorMessage.Size(m)
}
func (m *ErrorMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorMessage proto.InternalMessageInfo

func (m *ErrorMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Foobar)(nil), "foobar_grpc.Foobar")
	proto.RegisterType((*ListFoobar)(nil), "foobar_grpc.ListFoobar")
	proto.RegisterType((*DeleteResponse)(nil), "foobar_grpc.DeleteResponse")
	proto.RegisterType((*FetchRequest)(nil), "foobar_grpc.FetchRequest")
	proto.RegisterType((*SingleRequest)(nil), "foobar_grpc.SingleRequest")
	proto.RegisterType((*ErrorMessage)(nil), "foobar_grpc.ErrorMessage")
}

func init() {
	proto.RegisterFile("foobar.proto", fileDescriptor_1fe8bbf5351ca4f2)
}

var fileDescriptor_1fe8bbf5351ca4f2 = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xd1, 0x8a, 0x9b, 0x40,
	0x14, 0x45, 0x63, 0x2c, 0xde, 0x98, 0x50, 0xa6, 0xd0, 0x5a, 0xfb, 0x10, 0x91, 0x3e, 0x08, 0x05,
	0x43, 0xd3, 0x97, 0x14, 0xf2, 0x12, 0x92, 0x26, 0x0d, 0xb4, 0x2f, 0x93, 0xf6, 0xb9, 0x98, 0x78,
	0x63, 0x05, 0x75, 0xdc, 0x99, 0xf1, 0xc7, 0xf6, 0x13, 0xf6, 0xcb, 0x96, 0x38, 0xba, 0xab, 0x4b,
	0xd8, 0xdd, 0xb7, 0x7b, 0x3d, 0xe7, 0x70, 0xe6, 0x9c, 0x2b, 0xd8, 0x67, 0xc6, 0x8e, 0x11, 0x0f,
	0x4b, 0xce, 0x24, 0x23, 0x23, 0xb5, 0xfd, 0x4b, 0x78, 0x79, 0x72, 0xa7, 0x09, 0x63, 0x49, 0x86,
	0xb3, 0x1a, 0x3a, 0x56, 0xe7, 0x99, 0x4c, 0x73, 0x14, 0x32, 0xca, 0x4b, 0xc5, 0xf6, 0x6f, 0x35,
	0x30, 0xb7, 0xb5, 0x80, 0x4c, 0x40, 0xdf, 0x6f, 0x1c, 0xcd, 0xd3, 0x02, 0x83, 0xea, 0xfb, 0x0d,
	0xf9, 0x0c, 0x63, 0x85, 0xac, 0x59, 0x21, 0xb1, 0x90, 0x8e, 0xee, 0x69, 0x81, 0x45, 0xfb, 0x1f,
	0xc9, 0x02, 0xac, 0xbf, 0x65, 0x1c, 0x49, 0x8c, 0x57, 0xd2, 0x31, 0x3c, 0x2d, 0x18, 0xcd, 0xdd,
	0x50, 0xb9, 0x86, 0xad, 0x6b, 0xf8, 0xa7, 0x75, 0xa5, 0x8f, 0xe4, 0x8b, 0x72, 0xcd, 0xb1, 0x51,
	0x0e, 0x5f, 0x56, 0x3e, 0x90, 0xfd, 0xef, 0x00, 0xbf, 0x52, 0x21, 0x9b, 0x77, 0x7f, 0x01, 0x53,
	0x45, 0x76, 0x34, 0x6f, 0x10, 0x8c, 0xe6, 0xef, 0xc2, 0x4e, 0x03, 0xa1, 0x22, 0xd1, 0x86, 0xe2,
	0x2f, 0x61, 0xb2, 0xc1, 0x0c, 0x25, 0x52, 0x14, 0x25, 0x2b, 0x04, 0x92, 0xf7, 0x60, 0x0a, 0x19,
	0xc9, 0x4a, 0xd4, 0xd1, 0x2d, 0xda, 0x6c, 0x84, 0x80, 0x71, 0x62, 0x31, 0xd6, 0xa9, 0x0d, 0x5a,
	0xcf, 0xbe, 0x07, 0xf6, 0x16, 0xe5, 0xe9, 0x3f, 0xc5, 0x9b, 0x0a, 0x85, 0x24, 0x6f, 0x61, 0x50,
	0x54, 0x79, 0xd3, 0xd9, 0x65, 0xf4, 0xa7, 0x30, 0x3e, 0xa4, 0x45, 0x92, 0x61, 0x4b, 0x99, 0x80,
	0x9e, 0xc6, 0x6d, 0xab, 0x69, 0xec, 0x07, 0x60, 0xff, 0xe0, 0x9c, 0xf1, 0xdf, 0x28, 0x44, 0x94,
	0x20, 0x71, 0xe0, 0x4d, 0xae, 0xc6, 0xc6, 0xbf, 0x5d, 0xe7, 0x77, 0x7a, 0x7b, 0x80, 0x9f, 0x51,
	0x11, 0x67, 0xc8, 0xc9, 0x12, 0xac, 0x1d, 0xb6, 0xb1, 0xdd, 0x5e, 0xcc, 0x9e, 0xa9, 0x7b, 0xad,
	0x02, 0xb2, 0x86, 0xf1, 0x0e, 0x65, 0xa7, 0xb8, 0x8f, 0x7d, 0x56, 0x27, 0x98, 0xfb, 0xa1, 0x07,
	0x75, 0x34, 0x0b, 0xb0, 0xd5, 0x05, 0x9b, 0xfd, 0x9a, 0xd3, 0x75, 0xfb, 0x15, 0x98, 0xaa, 0xf9,
	0x67, 0x5f, 0xfe, 0xa9, 0x87, 0x3d, 0x39, 0xd5, 0x57, 0x18, 0x1e, 0x24, 0xe3, 0xf8, 0x7a, 0xd7,
	0xa3, 0x59, 0xff, 0x49, 0xdf, 0xee, 0x03, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x41, 0x61, 0x06, 0x24,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FoobarHandlerClient is the client API for FoobarHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FoobarHandlerClient interface {
	GetFoobar(ctx context.Context, in *SingleRequest, opts ...grpc.CallOption) (*Foobar, error)
	GetListFoobar(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*ListFoobar, error)
	UpdateFoobar(ctx context.Context, in *Foobar, opts ...grpc.CallOption) (*Foobar, error)
	Delete(ctx context.Context, in *SingleRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	Store(ctx context.Context, in *Foobar, opts ...grpc.CallOption) (*Foobar, error)
}

type foobarHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewFoobarHandlerClient(cc grpc.ClientConnInterface) FoobarHandlerClient {
	return &foobarHandlerClient{cc}
}

func (c *foobarHandlerClient) GetFoobar(ctx context.Context, in *SingleRequest, opts ...grpc.CallOption) (*Foobar, error) {
	out := new(Foobar)
	err := c.cc.Invoke(ctx, "/foobar_grpc.FoobarHandler/GetFoobar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foobarHandlerClient) GetListFoobar(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*ListFoobar, error) {
	out := new(ListFoobar)
	err := c.cc.Invoke(ctx, "/foobar_grpc.FoobarHandler/GetListFoobar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foobarHandlerClient) UpdateFoobar(ctx context.Context, in *Foobar, opts ...grpc.CallOption) (*Foobar, error) {
	out := new(Foobar)
	err := c.cc.Invoke(ctx, "/foobar_grpc.FoobarHandler/UpdateFoobar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foobarHandlerClient) Delete(ctx context.Context, in *SingleRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/foobar_grpc.FoobarHandler/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foobarHandlerClient) Store(ctx context.Context, in *Foobar, opts ...grpc.CallOption) (*Foobar, error) {
	out := new(Foobar)
	err := c.cc.Invoke(ctx, "/foobar_grpc.FoobarHandler/Store", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FoobarHandlerServer is the server API for FoobarHandler service.
type FoobarHandlerServer interface {
	GetFoobar(context.Context, *SingleRequest) (*Foobar, error)
	GetListFoobar(context.Context, *FetchRequest) (*ListFoobar, error)
	UpdateFoobar(context.Context, *Foobar) (*Foobar, error)
	Delete(context.Context, *SingleRequest) (*DeleteResponse, error)
	Store(context.Context, *Foobar) (*Foobar, error)
}

// UnimplementedFoobarHandlerServer can be embedded to have forward compatible implementations.
type UnimplementedFoobarHandlerServer struct {
}

func (*UnimplementedFoobarHandlerServer) GetFoobar(ctx context.Context, req *SingleRequest) (*Foobar, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFoobar not implemented")
}
func (*UnimplementedFoobarHandlerServer) GetListFoobar(ctx context.Context, req *FetchRequest) (*ListFoobar, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListFoobar not implemented")
}
func (*UnimplementedFoobarHandlerServer) UpdateFoobar(ctx context.Context, req *Foobar) (*Foobar, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFoobar not implemented")
}
func (*UnimplementedFoobarHandlerServer) Delete(ctx context.Context, req *SingleRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedFoobarHandlerServer) Store(ctx context.Context, req *Foobar) (*Foobar, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Store not implemented")
}

func RegisterFoobarHandlerServer(s *grpc.Server, srv FoobarHandlerServer) {
	s.RegisterService(&_FoobarHandler_serviceDesc, srv)
}

func _FoobarHandler_GetFoobar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoobarHandlerServer).GetFoobar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/foobar_grpc.FoobarHandler/GetFoobar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoobarHandlerServer).GetFoobar(ctx, req.(*SingleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoobarHandler_GetListFoobar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoobarHandlerServer).GetListFoobar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/foobar_grpc.FoobarHandler/GetListFoobar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoobarHandlerServer).GetListFoobar(ctx, req.(*FetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoobarHandler_UpdateFoobar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Foobar)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoobarHandlerServer).UpdateFoobar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/foobar_grpc.FoobarHandler/UpdateFoobar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoobarHandlerServer).UpdateFoobar(ctx, req.(*Foobar))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoobarHandler_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoobarHandlerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/foobar_grpc.FoobarHandler/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoobarHandlerServer).Delete(ctx, req.(*SingleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoobarHandler_Store_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Foobar)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoobarHandlerServer).Store(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/foobar_grpc.FoobarHandler/Store",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoobarHandlerServer).Store(ctx, req.(*Foobar))
	}
	return interceptor(ctx, in, info, handler)
}

var _FoobarHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "foobar_grpc.FoobarHandler",
	HandlerType: (*FoobarHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFoobar",
			Handler:    _FoobarHandler_GetFoobar_Handler,
		},
		{
			MethodName: "GetListFoobar",
			Handler:    _FoobarHandler_GetListFoobar_Handler,
		},
		{
			MethodName: "UpdateFoobar",
			Handler:    _FoobarHandler_UpdateFoobar_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _FoobarHandler_Delete_Handler,
		},
		{
			MethodName: "Store",
			Handler:    _FoobarHandler_Store_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "foobar.proto",
}
