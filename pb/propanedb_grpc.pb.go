// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DatabaseClient is the client API for Database service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DatabaseClient interface {
	CreateDatabase(ctx context.Context, in *PropaneDatabase, opts ...grpc.CallOption) (*PropaneStatus, error)
	Put(ctx context.Context, in *PropanePut, opts ...grpc.CallOption) (*PropaneId, error)
	Get(ctx context.Context, in *PropaneId, opts ...grpc.CallOption) (*PropaneEntity, error)
	Delete(ctx context.Context, in *PropaneId, opts ...grpc.CallOption) (*PropaneStatus, error)
	Search(ctx context.Context, in *PropaneSearch, opts ...grpc.CallOption) (*PropaneEntities, error)
}

type databaseClient struct {
	cc grpc.ClientConnInterface
}

func NewDatabaseClient(cc grpc.ClientConnInterface) DatabaseClient {
	return &databaseClient{cc}
}

func (c *databaseClient) CreateDatabase(ctx context.Context, in *PropaneDatabase, opts ...grpc.CallOption) (*PropaneStatus, error) {
	out := new(PropaneStatus)
	err := c.cc.Invoke(ctx, "/propane.Database/CreateDatabase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) Put(ctx context.Context, in *PropanePut, opts ...grpc.CallOption) (*PropaneId, error) {
	out := new(PropaneId)
	err := c.cc.Invoke(ctx, "/propane.Database/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) Get(ctx context.Context, in *PropaneId, opts ...grpc.CallOption) (*PropaneEntity, error) {
	out := new(PropaneEntity)
	err := c.cc.Invoke(ctx, "/propane.Database/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) Delete(ctx context.Context, in *PropaneId, opts ...grpc.CallOption) (*PropaneStatus, error) {
	out := new(PropaneStatus)
	err := c.cc.Invoke(ctx, "/propane.Database/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) Search(ctx context.Context, in *PropaneSearch, opts ...grpc.CallOption) (*PropaneEntities, error) {
	out := new(PropaneEntities)
	err := c.cc.Invoke(ctx, "/propane.Database/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatabaseServer is the server API for Database service.
// All implementations must embed UnimplementedDatabaseServer
// for forward compatibility
type DatabaseServer interface {
	CreateDatabase(context.Context, *PropaneDatabase) (*PropaneStatus, error)
	Put(context.Context, *PropanePut) (*PropaneId, error)
	Get(context.Context, *PropaneId) (*PropaneEntity, error)
	Delete(context.Context, *PropaneId) (*PropaneStatus, error)
	Search(context.Context, *PropaneSearch) (*PropaneEntities, error)
	mustEmbedUnimplementedDatabaseServer()
}

// UnimplementedDatabaseServer must be embedded to have forward compatible implementations.
type UnimplementedDatabaseServer struct {
}

func (UnimplementedDatabaseServer) CreateDatabase(context.Context, *PropaneDatabase) (*PropaneStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDatabase not implemented")
}
func (UnimplementedDatabaseServer) Put(context.Context, *PropanePut) (*PropaneId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (UnimplementedDatabaseServer) Get(context.Context, *PropaneId) (*PropaneEntity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedDatabaseServer) Delete(context.Context, *PropaneId) (*PropaneStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedDatabaseServer) Search(context.Context, *PropaneSearch) (*PropaneEntities, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedDatabaseServer) mustEmbedUnimplementedDatabaseServer() {}

// UnsafeDatabaseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DatabaseServer will
// result in compilation errors.
type UnsafeDatabaseServer interface {
	mustEmbedUnimplementedDatabaseServer()
}

func RegisterDatabaseServer(s grpc.ServiceRegistrar, srv DatabaseServer) {
	s.RegisterService(&Database_ServiceDesc, srv)
}

func _Database_CreateDatabase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PropaneDatabase)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).CreateDatabase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/propane.Database/CreateDatabase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).CreateDatabase(ctx, req.(*PropaneDatabase))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PropanePut)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/propane.Database/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).Put(ctx, req.(*PropanePut))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PropaneId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/propane.Database/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).Get(ctx, req.(*PropaneId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PropaneId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/propane.Database/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).Delete(ctx, req.(*PropaneId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PropaneSearch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/propane.Database/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).Search(ctx, req.(*PropaneSearch))
	}
	return interceptor(ctx, in, info, handler)
}

// Database_ServiceDesc is the grpc.ServiceDesc for Database service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Database_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "propane.Database",
	HandlerType: (*DatabaseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDatabase",
			Handler:    _Database_CreateDatabase_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _Database_Put_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Database_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Database_Delete_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _Database_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/propanedb.proto",
}