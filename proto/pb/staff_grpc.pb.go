// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: staff.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// StaffServiceClient is the client API for StaffService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StaffServiceClient interface {
	ListStaff(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListStaffResponse, error)
	CreateStaff(ctx context.Context, in *CreateStaffRequest, opts ...grpc.CallOption) (*CreateStaffResponse, error)
	UpdateStaff(ctx context.Context, in *UpdateStaffRequest, opts ...grpc.CallOption) (*UpdateStaffResponse, error)
	FindStaffById(ctx context.Context, in *StaffId, opts ...grpc.CallOption) (*Staff, error)
	DeleteStaff(ctx context.Context, in *StaffId, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type staffServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStaffServiceClient(cc grpc.ClientConnInterface) StaffServiceClient {
	return &staffServiceClient{cc}
}

func (c *staffServiceClient) ListStaff(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListStaffResponse, error) {
	out := new(ListStaffResponse)
	err := c.cc.Invoke(ctx, "/StaffService/listStaff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffServiceClient) CreateStaff(ctx context.Context, in *CreateStaffRequest, opts ...grpc.CallOption) (*CreateStaffResponse, error) {
	out := new(CreateStaffResponse)
	err := c.cc.Invoke(ctx, "/StaffService/createStaff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffServiceClient) UpdateStaff(ctx context.Context, in *UpdateStaffRequest, opts ...grpc.CallOption) (*UpdateStaffResponse, error) {
	out := new(UpdateStaffResponse)
	err := c.cc.Invoke(ctx, "/StaffService/updateStaff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffServiceClient) FindStaffById(ctx context.Context, in *StaffId, opts ...grpc.CallOption) (*Staff, error) {
	out := new(Staff)
	err := c.cc.Invoke(ctx, "/StaffService/findStaffById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffServiceClient) DeleteStaff(ctx context.Context, in *StaffId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/StaffService/deleteStaff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StaffServiceServer is the server API for StaffService service.
// All implementations must embed UnimplementedStaffServiceServer
// for forward compatibility
type StaffServiceServer interface {
	ListStaff(context.Context, *emptypb.Empty) (*ListStaffResponse, error)
	CreateStaff(context.Context, *CreateStaffRequest) (*CreateStaffResponse, error)
	UpdateStaff(context.Context, *UpdateStaffRequest) (*UpdateStaffResponse, error)
	FindStaffById(context.Context, *StaffId) (*Staff, error)
	DeleteStaff(context.Context, *StaffId) (*emptypb.Empty, error)
	mustEmbedUnimplementedStaffServiceServer()
}

// UnimplementedStaffServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStaffServiceServer struct {
}

func (UnimplementedStaffServiceServer) ListStaff(context.Context, *emptypb.Empty) (*ListStaffResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStaff not implemented")
}
func (UnimplementedStaffServiceServer) CreateStaff(context.Context, *CreateStaffRequest) (*CreateStaffResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStaff not implemented")
}
func (UnimplementedStaffServiceServer) UpdateStaff(context.Context, *UpdateStaffRequest) (*UpdateStaffResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStaff not implemented")
}
func (UnimplementedStaffServiceServer) FindStaffById(context.Context, *StaffId) (*Staff, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindStaffById not implemented")
}
func (UnimplementedStaffServiceServer) DeleteStaff(context.Context, *StaffId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStaff not implemented")
}
func (UnimplementedStaffServiceServer) mustEmbedUnimplementedStaffServiceServer() {}

// UnsafeStaffServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StaffServiceServer will
// result in compilation errors.
type UnsafeStaffServiceServer interface {
	mustEmbedUnimplementedStaffServiceServer()
}

func RegisterStaffServiceServer(s grpc.ServiceRegistrar, srv StaffServiceServer) {
	s.RegisterService(&StaffService_ServiceDesc, srv)
}

func _StaffService_ListStaff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).ListStaff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/StaffService/listStaff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).ListStaff(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffService_CreateStaff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStaffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).CreateStaff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/StaffService/createStaff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).CreateStaff(ctx, req.(*CreateStaffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffService_UpdateStaff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStaffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).UpdateStaff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/StaffService/updateStaff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).UpdateStaff(ctx, req.(*UpdateStaffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffService_FindStaffById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StaffId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).FindStaffById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/StaffService/findStaffById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).FindStaffById(ctx, req.(*StaffId))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffService_DeleteStaff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StaffId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).DeleteStaff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/StaffService/deleteStaff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).DeleteStaff(ctx, req.(*StaffId))
	}
	return interceptor(ctx, in, info, handler)
}

// StaffService_ServiceDesc is the grpc.ServiceDesc for StaffService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StaffService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "StaffService",
	HandlerType: (*StaffServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "listStaff",
			Handler:    _StaffService_ListStaff_Handler,
		},
		{
			MethodName: "createStaff",
			Handler:    _StaffService_CreateStaff_Handler,
		},
		{
			MethodName: "updateStaff",
			Handler:    _StaffService_UpdateStaff_Handler,
		},
		{
			MethodName: "findStaffById",
			Handler:    _StaffService_FindStaffById_Handler,
		},
		{
			MethodName: "deleteStaff",
			Handler:    _StaffService_DeleteStaff_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "staff.proto",
}
