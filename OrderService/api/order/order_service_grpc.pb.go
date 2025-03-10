// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.26.1
// source: order/order_service.proto

package order

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
	OrderService_PlaceOrder_FullMethodName      = "/api.order.OrderService/PlaceOrder"
	OrderService_ListOrder_FullMethodName       = "/api.order.OrderService/ListOrder"
	OrderService_GetOrderById_FullMethodName    = "/api.order.OrderService/GetOrderById"
	OrderService_DelOrderById_FullMethodName    = "/api.order.OrderService/DelOrderById"
	OrderService_MarkOrderPaid_FullMethodName   = "/api.order.OrderService/MarkOrderPaid"
	OrderService_MarkOrderCancel_FullMethodName = "/api.order.OrderService/MarkOrderCancel"
)

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderServiceClient interface {
	PlaceOrder(ctx context.Context, in *PlaceOrderReq, opts ...grpc.CallOption) (*PlaceOrderResp, error)
	ListOrder(ctx context.Context, in *ListOrderReq, opts ...grpc.CallOption) (*ListOrderResp, error)
	GetOrderById(ctx context.Context, in *GetOrderByIdReq, opts ...grpc.CallOption) (*GetOrderByIdResp, error)
	DelOrderById(ctx context.Context, in *DelOrderByIdReq, opts ...grpc.CallOption) (*DelOrderByIdResp, error)
	MarkOrderPaid(ctx context.Context, in *MarkOrderPaidReq, opts ...grpc.CallOption) (*MarkOrderPaidResp, error)
	MarkOrderCancel(ctx context.Context, in *MarkOrderCancelReq, opts ...grpc.CallOption) (*MarkOrderCancelResp, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) PlaceOrder(ctx context.Context, in *PlaceOrderReq, opts ...grpc.CallOption) (*PlaceOrderResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PlaceOrderResp)
	err := c.cc.Invoke(ctx, OrderService_PlaceOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) ListOrder(ctx context.Context, in *ListOrderReq, opts ...grpc.CallOption) (*ListOrderResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListOrderResp)
	err := c.cc.Invoke(ctx, OrderService_ListOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetOrderById(ctx context.Context, in *GetOrderByIdReq, opts ...grpc.CallOption) (*GetOrderByIdResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetOrderByIdResp)
	err := c.cc.Invoke(ctx, OrderService_GetOrderById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) DelOrderById(ctx context.Context, in *DelOrderByIdReq, opts ...grpc.CallOption) (*DelOrderByIdResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DelOrderByIdResp)
	err := c.cc.Invoke(ctx, OrderService_DelOrderById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) MarkOrderPaid(ctx context.Context, in *MarkOrderPaidReq, opts ...grpc.CallOption) (*MarkOrderPaidResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MarkOrderPaidResp)
	err := c.cc.Invoke(ctx, OrderService_MarkOrderPaid_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) MarkOrderCancel(ctx context.Context, in *MarkOrderCancelReq, opts ...grpc.CallOption) (*MarkOrderCancelResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MarkOrderCancelResp)
	err := c.cc.Invoke(ctx, OrderService_MarkOrderCancel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
// All implementations must embed UnimplementedOrderServiceServer
// for forward compatibility.
type OrderServiceServer interface {
	PlaceOrder(context.Context, *PlaceOrderReq) (*PlaceOrderResp, error)
	ListOrder(context.Context, *ListOrderReq) (*ListOrderResp, error)
	GetOrderById(context.Context, *GetOrderByIdReq) (*GetOrderByIdResp, error)
	DelOrderById(context.Context, *DelOrderByIdReq) (*DelOrderByIdResp, error)
	MarkOrderPaid(context.Context, *MarkOrderPaidReq) (*MarkOrderPaidResp, error)
	MarkOrderCancel(context.Context, *MarkOrderCancelReq) (*MarkOrderCancelResp, error)
	mustEmbedUnimplementedOrderServiceServer()
}

// UnimplementedOrderServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedOrderServiceServer struct{}

func (UnimplementedOrderServiceServer) PlaceOrder(context.Context, *PlaceOrderReq) (*PlaceOrderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaceOrder not implemented")
}
func (UnimplementedOrderServiceServer) ListOrder(context.Context, *ListOrderReq) (*ListOrderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrder not implemented")
}
func (UnimplementedOrderServiceServer) GetOrderById(context.Context, *GetOrderByIdReq) (*GetOrderByIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderById not implemented")
}
func (UnimplementedOrderServiceServer) DelOrderById(context.Context, *DelOrderByIdReq) (*DelOrderByIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelOrderById not implemented")
}
func (UnimplementedOrderServiceServer) MarkOrderPaid(context.Context, *MarkOrderPaidReq) (*MarkOrderPaidResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkOrderPaid not implemented")
}
func (UnimplementedOrderServiceServer) MarkOrderCancel(context.Context, *MarkOrderCancelReq) (*MarkOrderCancelResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkOrderCancel not implemented")
}
func (UnimplementedOrderServiceServer) mustEmbedUnimplementedOrderServiceServer() {}
func (UnimplementedOrderServiceServer) testEmbeddedByValue()                      {}

// UnsafeOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServiceServer will
// result in compilation errors.
type UnsafeOrderServiceServer interface {
	mustEmbedUnimplementedOrderServiceServer()
}

func RegisterOrderServiceServer(s grpc.ServiceRegistrar, srv OrderServiceServer) {
	// If the following call pancis, it indicates UnimplementedOrderServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&OrderService_ServiceDesc, srv)
}

func _OrderService_PlaceOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaceOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).PlaceOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_PlaceOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).PlaceOrder(ctx, req.(*PlaceOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_ListOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).ListOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_ListOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).ListOrder(ctx, req.(*ListOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetOrderById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetOrderById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_GetOrderById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetOrderById(ctx, req.(*GetOrderByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_DelOrderById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelOrderByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).DelOrderById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_DelOrderById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).DelOrderById(ctx, req.(*DelOrderByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_MarkOrderPaid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkOrderPaidReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).MarkOrderPaid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_MarkOrderPaid_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).MarkOrderPaid(ctx, req.(*MarkOrderPaidReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_MarkOrderCancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkOrderCancelReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).MarkOrderCancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_MarkOrderCancel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).MarkOrderCancel(ctx, req.(*MarkOrderCancelReq))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderService_ServiceDesc is the grpc.ServiceDesc for OrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.order.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PlaceOrder",
			Handler:    _OrderService_PlaceOrder_Handler,
		},
		{
			MethodName: "ListOrder",
			Handler:    _OrderService_ListOrder_Handler,
		},
		{
			MethodName: "GetOrderById",
			Handler:    _OrderService_GetOrderById_Handler,
		},
		{
			MethodName: "DelOrderById",
			Handler:    _OrderService_DelOrderById_Handler,
		},
		{
			MethodName: "MarkOrderPaid",
			Handler:    _OrderService_MarkOrderPaid_Handler,
		},
		{
			MethodName: "MarkOrderCancel",
			Handler:    _OrderService_MarkOrderCancel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order/order_service.proto",
}
