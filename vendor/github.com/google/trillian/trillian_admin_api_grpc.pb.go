// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package trillian

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

// TrillianAdminClient is the client API for TrillianAdmin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TrillianAdminClient interface {
	// Lists all trees the requester has access to.
	ListTrees(ctx context.Context, in *ListTreesRequest, opts ...grpc.CallOption) (*ListTreesResponse, error)
	// Retrieves a tree by ID.
	GetTree(ctx context.Context, in *GetTreeRequest, opts ...grpc.CallOption) (*Tree, error)
	// Creates a new tree.
	// System-generated fields are not required and will be ignored if present,
	// e.g.: tree_id, create_time and update_time.
	// Returns the created tree, with all system-generated fields assigned.
	CreateTree(ctx context.Context, in *CreateTreeRequest, opts ...grpc.CallOption) (*Tree, error)
	// Updates a tree.
	// See Tree for details. Readonly fields cannot be updated.
	UpdateTree(ctx context.Context, in *UpdateTreeRequest, opts ...grpc.CallOption) (*Tree, error)
	// Soft-deletes a tree.
	// A soft-deleted tree may be undeleted for a certain period, after which
	// it'll be permanently deleted.
	DeleteTree(ctx context.Context, in *DeleteTreeRequest, opts ...grpc.CallOption) (*Tree, error)
	// Undeletes a soft-deleted a tree.
	// A soft-deleted tree may be undeleted for a certain period, after which
	// it'll be permanently deleted.
	UndeleteTree(ctx context.Context, in *UndeleteTreeRequest, opts ...grpc.CallOption) (*Tree, error)
}

type trillianAdminClient struct {
	cc grpc.ClientConnInterface
}

func NewTrillianAdminClient(cc grpc.ClientConnInterface) TrillianAdminClient {
	return &trillianAdminClient{cc}
}

func (c *trillianAdminClient) ListTrees(ctx context.Context, in *ListTreesRequest, opts ...grpc.CallOption) (*ListTreesResponse, error) {
	out := new(ListTreesResponse)
	err := c.cc.Invoke(ctx, "/trillian.TrillianAdmin/ListTrees", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trillianAdminClient) GetTree(ctx context.Context, in *GetTreeRequest, opts ...grpc.CallOption) (*Tree, error) {
	out := new(Tree)
	err := c.cc.Invoke(ctx, "/trillian.TrillianAdmin/GetTree", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trillianAdminClient) CreateTree(ctx context.Context, in *CreateTreeRequest, opts ...grpc.CallOption) (*Tree, error) {
	out := new(Tree)
	err := c.cc.Invoke(ctx, "/trillian.TrillianAdmin/CreateTree", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trillianAdminClient) UpdateTree(ctx context.Context, in *UpdateTreeRequest, opts ...grpc.CallOption) (*Tree, error) {
	out := new(Tree)
	err := c.cc.Invoke(ctx, "/trillian.TrillianAdmin/UpdateTree", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trillianAdminClient) DeleteTree(ctx context.Context, in *DeleteTreeRequest, opts ...grpc.CallOption) (*Tree, error) {
	out := new(Tree)
	err := c.cc.Invoke(ctx, "/trillian.TrillianAdmin/DeleteTree", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trillianAdminClient) UndeleteTree(ctx context.Context, in *UndeleteTreeRequest, opts ...grpc.CallOption) (*Tree, error) {
	out := new(Tree)
	err := c.cc.Invoke(ctx, "/trillian.TrillianAdmin/UndeleteTree", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrillianAdminServer is the server API for TrillianAdmin service.
// All implementations should embed UnimplementedTrillianAdminServer
// for forward compatibility
type TrillianAdminServer interface {
	// Lists all trees the requester has access to.
	ListTrees(context.Context, *ListTreesRequest) (*ListTreesResponse, error)
	// Retrieves a tree by ID.
	GetTree(context.Context, *GetTreeRequest) (*Tree, error)
	// Creates a new tree.
	// System-generated fields are not required and will be ignored if present,
	// e.g.: tree_id, create_time and update_time.
	// Returns the created tree, with all system-generated fields assigned.
	CreateTree(context.Context, *CreateTreeRequest) (*Tree, error)
	// Updates a tree.
	// See Tree for details. Readonly fields cannot be updated.
	UpdateTree(context.Context, *UpdateTreeRequest) (*Tree, error)
	// Soft-deletes a tree.
	// A soft-deleted tree may be undeleted for a certain period, after which
	// it'll be permanently deleted.
	DeleteTree(context.Context, *DeleteTreeRequest) (*Tree, error)
	// Undeletes a soft-deleted a tree.
	// A soft-deleted tree may be undeleted for a certain period, after which
	// it'll be permanently deleted.
	UndeleteTree(context.Context, *UndeleteTreeRequest) (*Tree, error)
}

// UnimplementedTrillianAdminServer should be embedded to have forward compatible implementations.
type UnimplementedTrillianAdminServer struct {
}

func (UnimplementedTrillianAdminServer) ListTrees(context.Context, *ListTreesRequest) (*ListTreesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTrees not implemented")
}
func (UnimplementedTrillianAdminServer) GetTree(context.Context, *GetTreeRequest) (*Tree, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTree not implemented")
}
func (UnimplementedTrillianAdminServer) CreateTree(context.Context, *CreateTreeRequest) (*Tree, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTree not implemented")
}
func (UnimplementedTrillianAdminServer) UpdateTree(context.Context, *UpdateTreeRequest) (*Tree, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTree not implemented")
}
func (UnimplementedTrillianAdminServer) DeleteTree(context.Context, *DeleteTreeRequest) (*Tree, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTree not implemented")
}
func (UnimplementedTrillianAdminServer) UndeleteTree(context.Context, *UndeleteTreeRequest) (*Tree, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UndeleteTree not implemented")
}

// UnsafeTrillianAdminServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TrillianAdminServer will
// result in compilation errors.
type UnsafeTrillianAdminServer interface {
	mustEmbedUnimplementedTrillianAdminServer()
}

func RegisterTrillianAdminServer(s grpc.ServiceRegistrar, srv TrillianAdminServer) {
	s.RegisterService(&TrillianAdmin_ServiceDesc, srv)
}

func _TrillianAdmin_ListTrees_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTreesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrillianAdminServer).ListTrees(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trillian.TrillianAdmin/ListTrees",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrillianAdminServer).ListTrees(ctx, req.(*ListTreesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrillianAdmin_GetTree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTreeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrillianAdminServer).GetTree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trillian.TrillianAdmin/GetTree",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrillianAdminServer).GetTree(ctx, req.(*GetTreeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrillianAdmin_CreateTree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTreeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrillianAdminServer).CreateTree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trillian.TrillianAdmin/CreateTree",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrillianAdminServer).CreateTree(ctx, req.(*CreateTreeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrillianAdmin_UpdateTree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTreeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrillianAdminServer).UpdateTree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trillian.TrillianAdmin/UpdateTree",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrillianAdminServer).UpdateTree(ctx, req.(*UpdateTreeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrillianAdmin_DeleteTree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTreeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrillianAdminServer).DeleteTree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trillian.TrillianAdmin/DeleteTree",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrillianAdminServer).DeleteTree(ctx, req.(*DeleteTreeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrillianAdmin_UndeleteTree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UndeleteTreeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrillianAdminServer).UndeleteTree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trillian.TrillianAdmin/UndeleteTree",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrillianAdminServer).UndeleteTree(ctx, req.(*UndeleteTreeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TrillianAdmin_ServiceDesc is the grpc.ServiceDesc for TrillianAdmin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TrillianAdmin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "trillian.TrillianAdmin",
	HandlerType: (*TrillianAdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListTrees",
			Handler:    _TrillianAdmin_ListTrees_Handler,
		},
		{
			MethodName: "GetTree",
			Handler:    _TrillianAdmin_GetTree_Handler,
		},
		{
			MethodName: "CreateTree",
			Handler:    _TrillianAdmin_CreateTree_Handler,
		},
		{
			MethodName: "UpdateTree",
			Handler:    _TrillianAdmin_UpdateTree_Handler,
		},
		{
			MethodName: "DeleteTree",
			Handler:    _TrillianAdmin_DeleteTree_Handler,
		},
		{
			MethodName: "UndeleteTree",
			Handler:    _TrillianAdmin_UndeleteTree_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "trillian_admin_api.proto",
}