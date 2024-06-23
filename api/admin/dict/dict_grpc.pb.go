// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: admin/dict/dict.proto

package dict

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Dict_CreateDict_FullMethodName            = "/api.admin.dict.Dict/CreateDict"
	Dict_UpdateDict_FullMethodName            = "/api.admin.dict.Dict/UpdateDict"
	Dict_ListDict_FullMethodName              = "/api.admin.dict.Dict/ListDict"
	Dict_BatchUpdateDictStatus_FullMethodName = "/api.admin.dict.Dict/BatchUpdateDictStatus"
	Dict_DeleteDict_FullMethodName            = "/api.admin.dict.Dict/DeleteDict"
	Dict_GetDict_FullMethodName               = "/api.admin.dict.Dict/GetDict"
)

// DictClient is the client API for Dict service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 字典服务
type DictClient interface {
	CreateDict(ctx context.Context, in *CreateDictRequest, opts ...grpc.CallOption) (*CreateDictReply, error)
	// 更新用户
	UpdateDict(ctx context.Context, in *UpdateDictRequest, opts ...grpc.CallOption) (*UpdateDictReply, error)
	// 字典列表
	ListDict(ctx context.Context, in *GetDictSelectListRequest, opts ...grpc.CallOption) (*ListDictReply, error)
	// 批量修改字典状态
	BatchUpdateDictStatus(ctx context.Context, in *BatchUpdateDictStatusRequest, opts ...grpc.CallOption) (*BatchUpdateDictStatusReply, error)
	DeleteDict(ctx context.Context, in *DeleteDictRequest, opts ...grpc.CallOption) (*DeleteDictReply, error)
	GetDict(ctx context.Context, in *GetDictRequest, opts ...grpc.CallOption) (*GetDictReply, error)
}

type dictClient struct {
	cc grpc.ClientConnInterface
}

func NewDictClient(cc grpc.ClientConnInterface) DictClient {
	return &dictClient{cc}
}

func (c *dictClient) CreateDict(ctx context.Context, in *CreateDictRequest, opts ...grpc.CallOption) (*CreateDictReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateDictReply)
	err := c.cc.Invoke(ctx, Dict_CreateDict_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictClient) UpdateDict(ctx context.Context, in *UpdateDictRequest, opts ...grpc.CallOption) (*UpdateDictReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateDictReply)
	err := c.cc.Invoke(ctx, Dict_UpdateDict_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictClient) ListDict(ctx context.Context, in *GetDictSelectListRequest, opts ...grpc.CallOption) (*ListDictReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListDictReply)
	err := c.cc.Invoke(ctx, Dict_ListDict_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictClient) BatchUpdateDictStatus(ctx context.Context, in *BatchUpdateDictStatusRequest, opts ...grpc.CallOption) (*BatchUpdateDictStatusReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BatchUpdateDictStatusReply)
	err := c.cc.Invoke(ctx, Dict_BatchUpdateDictStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictClient) DeleteDict(ctx context.Context, in *DeleteDictRequest, opts ...grpc.CallOption) (*DeleteDictReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteDictReply)
	err := c.cc.Invoke(ctx, Dict_DeleteDict_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictClient) GetDict(ctx context.Context, in *GetDictRequest, opts ...grpc.CallOption) (*GetDictReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDictReply)
	err := c.cc.Invoke(ctx, Dict_GetDict_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DictServer is the server API for Dict service.
// All implementations must embed UnimplementedDictServer
// for forward compatibility
//
// 字典服务
type DictServer interface {
	CreateDict(context.Context, *CreateDictRequest) (*CreateDictReply, error)
	// 更新用户
	UpdateDict(context.Context, *UpdateDictRequest) (*UpdateDictReply, error)
	// 字典列表
	ListDict(context.Context, *GetDictSelectListRequest) (*ListDictReply, error)
	// 批量修改字典状态
	BatchUpdateDictStatus(context.Context, *BatchUpdateDictStatusRequest) (*BatchUpdateDictStatusReply, error)
	DeleteDict(context.Context, *DeleteDictRequest) (*DeleteDictReply, error)
	GetDict(context.Context, *GetDictRequest) (*GetDictReply, error)
	mustEmbedUnimplementedDictServer()
}

// UnimplementedDictServer must be embedded to have forward compatible implementations.
type UnimplementedDictServer struct {
}

func (UnimplementedDictServer) CreateDict(context.Context, *CreateDictRequest) (*CreateDictReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDict not implemented")
}
func (UnimplementedDictServer) UpdateDict(context.Context, *UpdateDictRequest) (*UpdateDictReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDict not implemented")
}
func (UnimplementedDictServer) ListDict(context.Context, *GetDictSelectListRequest) (*ListDictReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDict not implemented")
}
func (UnimplementedDictServer) BatchUpdateDictStatus(context.Context, *BatchUpdateDictStatusRequest) (*BatchUpdateDictStatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchUpdateDictStatus not implemented")
}
func (UnimplementedDictServer) DeleteDict(context.Context, *DeleteDictRequest) (*DeleteDictReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDict not implemented")
}
func (UnimplementedDictServer) GetDict(context.Context, *GetDictRequest) (*GetDictReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDict not implemented")
}
func (UnimplementedDictServer) mustEmbedUnimplementedDictServer() {}

// UnsafeDictServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DictServer will
// result in compilation errors.
type UnsafeDictServer interface {
	mustEmbedUnimplementedDictServer()
}

func RegisterDictServer(s grpc.ServiceRegistrar, srv DictServer) {
	s.RegisterService(&Dict_ServiceDesc, srv)
}

func _Dict_CreateDict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictServer).CreateDict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dict_CreateDict_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictServer).CreateDict(ctx, req.(*CreateDictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dict_UpdateDict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictServer).UpdateDict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dict_UpdateDict_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictServer).UpdateDict(ctx, req.(*UpdateDictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dict_ListDict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDictSelectListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictServer).ListDict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dict_ListDict_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictServer).ListDict(ctx, req.(*GetDictSelectListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dict_BatchUpdateDictStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchUpdateDictStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictServer).BatchUpdateDictStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dict_BatchUpdateDictStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictServer).BatchUpdateDictStatus(ctx, req.(*BatchUpdateDictStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dict_DeleteDict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictServer).DeleteDict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dict_DeleteDict_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictServer).DeleteDict(ctx, req.(*DeleteDictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dict_GetDict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictServer).GetDict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dict_GetDict_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictServer).GetDict(ctx, req.(*GetDictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Dict_ServiceDesc is the grpc.ServiceDesc for Dict service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Dict_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.admin.dict.Dict",
	HandlerType: (*DictServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDict",
			Handler:    _Dict_CreateDict_Handler,
		},
		{
			MethodName: "UpdateDict",
			Handler:    _Dict_UpdateDict_Handler,
		},
		{
			MethodName: "ListDict",
			Handler:    _Dict_ListDict_Handler,
		},
		{
			MethodName: "BatchUpdateDictStatus",
			Handler:    _Dict_BatchUpdateDictStatus_Handler,
		},
		{
			MethodName: "DeleteDict",
			Handler:    _Dict_DeleteDict_Handler,
		},
		{
			MethodName: "GetDict",
			Handler:    _Dict_GetDict_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin/dict/dict.proto",
}
