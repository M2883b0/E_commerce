// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.26.1
// source: ai_agent/ai_agent.proto

package ai_agent

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AiAgent_UserRequest_FullMethodName = "/api.ai_agent.AiAgent/UserRequest"
)

// AiAgentClient is the client API for AiAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AiAgentClient interface {
	UserRequest(ctx context.Context, in *UserRequestReq, opts ...grpc.CallOption) (*UserRequestResp, error)
}

type aiAgentClient struct {
	cc grpc.ClientConnInterface
}

func NewAiAgentClient(cc grpc.ClientConnInterface) AiAgentClient {
	return &aiAgentClient{cc}
}

func (c *aiAgentClient) UserRequest(ctx context.Context, in *UserRequestReq, opts ...grpc.CallOption) (*UserRequestResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserRequestResp)
	err := c.cc.Invoke(ctx, AiAgent_UserRequest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AiAgentServer is the server API for AiAgent service.
// All implementations must embed UnimplementedAiAgentServer
// for forward compatibility.
type AiAgentServer interface {
	UserRequest(context.Context, *UserRequestReq) (*UserRequestResp, error)
	mustEmbedUnimplementedAiAgentServer()
}

// UnimplementedAiAgentServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAiAgentServer struct{}

func (UnimplementedAiAgentServer) UserRequest(context.Context, *UserRequestReq) (*UserRequestResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRequest not implemented")
}
func (UnimplementedAiAgentServer) mustEmbedUnimplementedAiAgentServer() {}
func (UnimplementedAiAgentServer) testEmbeddedByValue()                 {}

// UnsafeAiAgentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AiAgentServer will
// result in compilation errors.
type UnsafeAiAgentServer interface {
	mustEmbedUnimplementedAiAgentServer()
}

func RegisterAiAgentServer(s grpc.ServiceRegistrar, srv AiAgentServer) {
	// If the following call pancis, it indicates UnimplementedAiAgentServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AiAgent_ServiceDesc, srv)
}

func _AiAgent_UserRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequestReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AiAgentServer).UserRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AiAgent_UserRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AiAgentServer).UserRequest(ctx, req.(*UserRequestReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AiAgent_ServiceDesc is the grpc.ServiceDesc for AiAgent service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AiAgent_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.ai_agent.AiAgent",
	HandlerType: (*AiAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserRequest",
			Handler:    _AiAgent_UserRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ai_agent/ai_agent.proto",
}
