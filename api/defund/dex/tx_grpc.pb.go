// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: defund/dex/tx.proto

package dex

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

const (
	Msg_CreatePair_FullMethodName       = "/defund.dex.Msg/CreatePair"
	Msg_CreatePool_FullMethodName       = "/defund.dex.Msg/CreatePool"
	Msg_CreateRangedPool_FullMethodName = "/defund.dex.Msg/CreateRangedPool"
	Msg_Deposit_FullMethodName          = "/defund.dex.Msg/Deposit"
	Msg_Withdraw_FullMethodName         = "/defund.dex.Msg/Withdraw"
	Msg_LimitOrder_FullMethodName       = "/defund.dex.Msg/LimitOrder"
	Msg_MarketOrder_FullMethodName      = "/defund.dex.Msg/MarketOrder"
	Msg_MMOrder_FullMethodName          = "/defund.dex.Msg/MMOrder"
	Msg_CancelOrder_FullMethodName      = "/defund.dex.Msg/CancelOrder"
	Msg_CancelAllOrders_FullMethodName  = "/defund.dex.Msg/CancelAllOrders"
	Msg_UpdateParams_FullMethodName     = "/defund.dex.Msg/UpdateParams"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// CreatePair defines a method for creating a pair
	CreatePair(ctx context.Context, in *MsgCreatePair, opts ...grpc.CallOption) (*MsgCreatePairResponse, error)
	// CreatePool defines a method for creating a pool
	CreatePool(ctx context.Context, in *MsgCreatePool, opts ...grpc.CallOption) (*MsgCreatePoolResponse, error)
	// CreateRangePool defines a method for creating a ranged pool
	CreateRangedPool(ctx context.Context, in *MsgCreateRangedPool, opts ...grpc.CallOption) (*MsgCreateRangedPoolResponse, error)
	// Deposit defines a method for depositing coins to the pool
	Deposit(ctx context.Context, in *MsgDeposit, opts ...grpc.CallOption) (*MsgDepositResponse, error)
	// Withdraw defines a method for withdrawing pool coin from the pool
	Withdraw(ctx context.Context, in *MsgWithdraw, opts ...grpc.CallOption) (*MsgWithdrawResponse, error)
	// LimitOrder defines a method for making a limit order
	LimitOrder(ctx context.Context, in *MsgLimitOrder, opts ...grpc.CallOption) (*MsgLimitOrderResponse, error)
	// MarketOrder defines a method for making a market order
	MarketOrder(ctx context.Context, in *MsgMarketOrder, opts ...grpc.CallOption) (*MsgMarketOrderResponse, error)
	// MsgMMOrder defines a method for making a MM(market making) order
	MMOrder(ctx context.Context, in *MsgMMOrder, opts ...grpc.CallOption) (*MsgMMOrderResponse, error)
	// CancelOrder defines a method for cancelling an order
	CancelOrder(ctx context.Context, in *MsgCancelOrder, opts ...grpc.CallOption) (*MsgCancelOrderResponse, error)
	// CancelAllOrders defines a method for cancelling all orders
	CancelAllOrders(ctx context.Context, in *MsgCancelAllOrders, opts ...grpc.CallOption) (*MsgCancelAllOrdersResponse, error)
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreatePair(ctx context.Context, in *MsgCreatePair, opts ...grpc.CallOption) (*MsgCreatePairResponse, error) {
	out := new(MsgCreatePairResponse)
	err := c.cc.Invoke(ctx, Msg_CreatePair_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreatePool(ctx context.Context, in *MsgCreatePool, opts ...grpc.CallOption) (*MsgCreatePoolResponse, error) {
	out := new(MsgCreatePoolResponse)
	err := c.cc.Invoke(ctx, Msg_CreatePool_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateRangedPool(ctx context.Context, in *MsgCreateRangedPool, opts ...grpc.CallOption) (*MsgCreateRangedPoolResponse, error) {
	out := new(MsgCreateRangedPoolResponse)
	err := c.cc.Invoke(ctx, Msg_CreateRangedPool_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Deposit(ctx context.Context, in *MsgDeposit, opts ...grpc.CallOption) (*MsgDepositResponse, error) {
	out := new(MsgDepositResponse)
	err := c.cc.Invoke(ctx, Msg_Deposit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Withdraw(ctx context.Context, in *MsgWithdraw, opts ...grpc.CallOption) (*MsgWithdrawResponse, error) {
	out := new(MsgWithdrawResponse)
	err := c.cc.Invoke(ctx, Msg_Withdraw_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) LimitOrder(ctx context.Context, in *MsgLimitOrder, opts ...grpc.CallOption) (*MsgLimitOrderResponse, error) {
	out := new(MsgLimitOrderResponse)
	err := c.cc.Invoke(ctx, Msg_LimitOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) MarketOrder(ctx context.Context, in *MsgMarketOrder, opts ...grpc.CallOption) (*MsgMarketOrderResponse, error) {
	out := new(MsgMarketOrderResponse)
	err := c.cc.Invoke(ctx, Msg_MarketOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) MMOrder(ctx context.Context, in *MsgMMOrder, opts ...grpc.CallOption) (*MsgMMOrderResponse, error) {
	out := new(MsgMMOrderResponse)
	err := c.cc.Invoke(ctx, Msg_MMOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CancelOrder(ctx context.Context, in *MsgCancelOrder, opts ...grpc.CallOption) (*MsgCancelOrderResponse, error) {
	out := new(MsgCancelOrderResponse)
	err := c.cc.Invoke(ctx, Msg_CancelOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CancelAllOrders(ctx context.Context, in *MsgCancelAllOrders, opts ...grpc.CallOption) (*MsgCancelAllOrdersResponse, error) {
	out := new(MsgCancelAllOrdersResponse)
	err := c.cc.Invoke(ctx, Msg_CancelAllOrders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// CreatePair defines a method for creating a pair
	CreatePair(context.Context, *MsgCreatePair) (*MsgCreatePairResponse, error)
	// CreatePool defines a method for creating a pool
	CreatePool(context.Context, *MsgCreatePool) (*MsgCreatePoolResponse, error)
	// CreateRangePool defines a method for creating a ranged pool
	CreateRangedPool(context.Context, *MsgCreateRangedPool) (*MsgCreateRangedPoolResponse, error)
	// Deposit defines a method for depositing coins to the pool
	Deposit(context.Context, *MsgDeposit) (*MsgDepositResponse, error)
	// Withdraw defines a method for withdrawing pool coin from the pool
	Withdraw(context.Context, *MsgWithdraw) (*MsgWithdrawResponse, error)
	// LimitOrder defines a method for making a limit order
	LimitOrder(context.Context, *MsgLimitOrder) (*MsgLimitOrderResponse, error)
	// MarketOrder defines a method for making a market order
	MarketOrder(context.Context, *MsgMarketOrder) (*MsgMarketOrderResponse, error)
	// MsgMMOrder defines a method for making a MM(market making) order
	MMOrder(context.Context, *MsgMMOrder) (*MsgMMOrderResponse, error)
	// CancelOrder defines a method for cancelling an order
	CancelOrder(context.Context, *MsgCancelOrder) (*MsgCancelOrderResponse, error)
	// CancelAllOrders defines a method for cancelling all orders
	CancelAllOrders(context.Context, *MsgCancelAllOrders) (*MsgCancelAllOrdersResponse, error)
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) CreatePair(context.Context, *MsgCreatePair) (*MsgCreatePairResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePair not implemented")
}
func (UnimplementedMsgServer) CreatePool(context.Context, *MsgCreatePool) (*MsgCreatePoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePool not implemented")
}
func (UnimplementedMsgServer) CreateRangedPool(context.Context, *MsgCreateRangedPool) (*MsgCreateRangedPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRangedPool not implemented")
}
func (UnimplementedMsgServer) Deposit(context.Context, *MsgDeposit) (*MsgDepositResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deposit not implemented")
}
func (UnimplementedMsgServer) Withdraw(context.Context, *MsgWithdraw) (*MsgWithdrawResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Withdraw not implemented")
}
func (UnimplementedMsgServer) LimitOrder(context.Context, *MsgLimitOrder) (*MsgLimitOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LimitOrder not implemented")
}
func (UnimplementedMsgServer) MarketOrder(context.Context, *MsgMarketOrder) (*MsgMarketOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarketOrder not implemented")
}
func (UnimplementedMsgServer) MMOrder(context.Context, *MsgMMOrder) (*MsgMMOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MMOrder not implemented")
}
func (UnimplementedMsgServer) CancelOrder(context.Context, *MsgCancelOrder) (*MsgCancelOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelOrder not implemented")
}
func (UnimplementedMsgServer) CancelAllOrders(context.Context, *MsgCancelAllOrders) (*MsgCancelAllOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelAllOrders not implemented")
}
func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_CreatePair_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreatePair)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreatePair(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreatePair_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreatePair(ctx, req.(*MsgCreatePair))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreatePool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreatePool)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreatePool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreatePool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreatePool(ctx, req.(*MsgCreatePool))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateRangedPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateRangedPool)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateRangedPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreateRangedPool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateRangedPool(ctx, req.(*MsgCreateRangedPool))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Deposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeposit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Deposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_Deposit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Deposit(ctx, req.(*MsgDeposit))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Withdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgWithdraw)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Withdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_Withdraw_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Withdraw(ctx, req.(*MsgWithdraw))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_LimitOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgLimitOrder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).LimitOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_LimitOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).LimitOrder(ctx, req.(*MsgLimitOrder))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_MarketOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgMarketOrder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).MarketOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_MarketOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).MarketOrder(ctx, req.(*MsgMarketOrder))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_MMOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgMMOrder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).MMOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_MMOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).MMOrder(ctx, req.(*MsgMMOrder))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCancelOrder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CancelOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CancelOrder(ctx, req.(*MsgCancelOrder))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CancelAllOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCancelAllOrders)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CancelAllOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CancelAllOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CancelAllOrders(ctx, req.(*MsgCancelAllOrders))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "defund.dex.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePair",
			Handler:    _Msg_CreatePair_Handler,
		},
		{
			MethodName: "CreatePool",
			Handler:    _Msg_CreatePool_Handler,
		},
		{
			MethodName: "CreateRangedPool",
			Handler:    _Msg_CreateRangedPool_Handler,
		},
		{
			MethodName: "Deposit",
			Handler:    _Msg_Deposit_Handler,
		},
		{
			MethodName: "Withdraw",
			Handler:    _Msg_Withdraw_Handler,
		},
		{
			MethodName: "LimitOrder",
			Handler:    _Msg_LimitOrder_Handler,
		},
		{
			MethodName: "MarketOrder",
			Handler:    _Msg_MarketOrder_Handler,
		},
		{
			MethodName: "MMOrder",
			Handler:    _Msg_MMOrder_Handler,
		},
		{
			MethodName: "CancelOrder",
			Handler:    _Msg_CancelOrder_Handler,
		},
		{
			MethodName: "CancelAllOrders",
			Handler:    _Msg_CancelAllOrders_Handler,
		},
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "defund/dex/tx.proto",
}
