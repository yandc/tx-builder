// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.26.1
// source: builder/v1/tx.proto

package v1

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
	Tx_BuildTx_FullMethodName              = "/api.builder.v1.Tx/BuildTx"
	Tx_SignTx_FullMethodName               = "/api.builder.v1.Tx/SignTx"
	Tx_SendRawTx_FullMethodName            = "/api.builder.v1.Tx/SendRawTx"
	Tx_SendTx_FullMethodName               = "/api.builder.v1.Tx/SendTx"
	Tx_GetBalance_FullMethodName           = "/api.builder.v1.Tx/GetBalance"
	Tx_GetTransactionByHash_FullMethodName = "/api.builder.v1.Tx/GetTransactionByHash"
	Tx_GetTransactionList_FullMethodName   = "/api.builder.v1.Tx/GetTransactionList"
	Tx_GetAssetList_FullMethodName         = "/api.builder.v1.Tx/GetAssetList"
)

// TxClient is the client API for Tx service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TxClient interface {
	BuildTx(ctx context.Context, in *TxInfoRequest, opts ...grpc.CallOption) (*BuildTxReply, error)
	SignTx(ctx context.Context, in *SignTxRequest, opts ...grpc.CallOption) (*SignTxReply, error)
	SendRawTx(ctx context.Context, in *SendRawTxRequest, opts ...grpc.CallOption) (*SendRawTxReply, error)
	SendTx(ctx context.Context, in *TxInfoRequest, opts ...grpc.CallOption) (*SendRawTxReply, error)
	GetBalance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceReply, error)
	GetTransactionByHash(ctx context.Context, in *GetTransactionByHashRequest, opts ...grpc.CallOption) (*TransactionRecord, error)
	GetTransactionList(ctx context.Context, in *PageListRequest, opts ...grpc.CallOption) (*PageListResponse, error)
	GetAssetList(ctx context.Context, in *PageListAssetRequest, opts ...grpc.CallOption) (*PageListAssetResponse, error)
}

type txClient struct {
	cc grpc.ClientConnInterface
}

func NewTxClient(cc grpc.ClientConnInterface) TxClient {
	return &txClient{cc}
}

func (c *txClient) BuildTx(ctx context.Context, in *TxInfoRequest, opts ...grpc.CallOption) (*BuildTxReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BuildTxReply)
	err := c.cc.Invoke(ctx, Tx_BuildTx_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *txClient) SignTx(ctx context.Context, in *SignTxRequest, opts ...grpc.CallOption) (*SignTxReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SignTxReply)
	err := c.cc.Invoke(ctx, Tx_SignTx_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *txClient) SendRawTx(ctx context.Context, in *SendRawTxRequest, opts ...grpc.CallOption) (*SendRawTxReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendRawTxReply)
	err := c.cc.Invoke(ctx, Tx_SendRawTx_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *txClient) SendTx(ctx context.Context, in *TxInfoRequest, opts ...grpc.CallOption) (*SendRawTxReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendRawTxReply)
	err := c.cc.Invoke(ctx, Tx_SendTx_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *txClient) GetBalance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BalanceReply)
	err := c.cc.Invoke(ctx, Tx_GetBalance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *txClient) GetTransactionByHash(ctx context.Context, in *GetTransactionByHashRequest, opts ...grpc.CallOption) (*TransactionRecord, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TransactionRecord)
	err := c.cc.Invoke(ctx, Tx_GetTransactionByHash_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *txClient) GetTransactionList(ctx context.Context, in *PageListRequest, opts ...grpc.CallOption) (*PageListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PageListResponse)
	err := c.cc.Invoke(ctx, Tx_GetTransactionList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *txClient) GetAssetList(ctx context.Context, in *PageListAssetRequest, opts ...grpc.CallOption) (*PageListAssetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PageListAssetResponse)
	err := c.cc.Invoke(ctx, Tx_GetAssetList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TxServer is the server API for Tx service.
// All implementations must embed UnimplementedTxServer
// for forward compatibility.
type TxServer interface {
	BuildTx(context.Context, *TxInfoRequest) (*BuildTxReply, error)
	SignTx(context.Context, *SignTxRequest) (*SignTxReply, error)
	SendRawTx(context.Context, *SendRawTxRequest) (*SendRawTxReply, error)
	SendTx(context.Context, *TxInfoRequest) (*SendRawTxReply, error)
	GetBalance(context.Context, *BalanceRequest) (*BalanceReply, error)
	GetTransactionByHash(context.Context, *GetTransactionByHashRequest) (*TransactionRecord, error)
	GetTransactionList(context.Context, *PageListRequest) (*PageListResponse, error)
	GetAssetList(context.Context, *PageListAssetRequest) (*PageListAssetResponse, error)
	mustEmbedUnimplementedTxServer()
}

// UnimplementedTxServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTxServer struct{}

func (UnimplementedTxServer) BuildTx(context.Context, *TxInfoRequest) (*BuildTxReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuildTx not implemented")
}
func (UnimplementedTxServer) SignTx(context.Context, *SignTxRequest) (*SignTxReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignTx not implemented")
}
func (UnimplementedTxServer) SendRawTx(context.Context, *SendRawTxRequest) (*SendRawTxReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendRawTx not implemented")
}
func (UnimplementedTxServer) SendTx(context.Context, *TxInfoRequest) (*SendRawTxReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTx not implemented")
}
func (UnimplementedTxServer) GetBalance(context.Context, *BalanceRequest) (*BalanceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}
func (UnimplementedTxServer) GetTransactionByHash(context.Context, *GetTransactionByHashRequest) (*TransactionRecord, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionByHash not implemented")
}
func (UnimplementedTxServer) GetTransactionList(context.Context, *PageListRequest) (*PageListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionList not implemented")
}
func (UnimplementedTxServer) GetAssetList(context.Context, *PageListAssetRequest) (*PageListAssetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAssetList not implemented")
}
func (UnimplementedTxServer) mustEmbedUnimplementedTxServer() {}
func (UnimplementedTxServer) testEmbeddedByValue()            {}

// UnsafeTxServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TxServer will
// result in compilation errors.
type UnsafeTxServer interface {
	mustEmbedUnimplementedTxServer()
}

func RegisterTxServer(s grpc.ServiceRegistrar, srv TxServer) {
	// If the following call pancis, it indicates UnimplementedTxServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Tx_ServiceDesc, srv)
}

func _Tx_BuildTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TxServer).BuildTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tx_BuildTx_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TxServer).BuildTx(ctx, req.(*TxInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tx_SignTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TxServer).SignTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tx_SignTx_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TxServer).SignTx(ctx, req.(*SignTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tx_SendRawTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRawTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TxServer).SendRawTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tx_SendRawTx_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TxServer).SendRawTx(ctx, req.(*SendRawTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tx_SendTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TxServer).SendTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tx_SendTx_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TxServer).SendTx(ctx, req.(*TxInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tx_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TxServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tx_GetBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TxServer).GetBalance(ctx, req.(*BalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tx_GetTransactionByHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionByHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TxServer).GetTransactionByHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tx_GetTransactionByHash_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TxServer).GetTransactionByHash(ctx, req.(*GetTransactionByHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tx_GetTransactionList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TxServer).GetTransactionList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tx_GetTransactionList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TxServer).GetTransactionList(ctx, req.(*PageListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tx_GetAssetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageListAssetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TxServer).GetAssetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tx_GetAssetList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TxServer).GetAssetList(ctx, req.(*PageListAssetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Tx_ServiceDesc is the grpc.ServiceDesc for Tx service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Tx_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.builder.v1.Tx",
	HandlerType: (*TxServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BuildTx",
			Handler:    _Tx_BuildTx_Handler,
		},
		{
			MethodName: "SignTx",
			Handler:    _Tx_SignTx_Handler,
		},
		{
			MethodName: "SendRawTx",
			Handler:    _Tx_SendRawTx_Handler,
		},
		{
			MethodName: "SendTx",
			Handler:    _Tx_SendTx_Handler,
		},
		{
			MethodName: "GetBalance",
			Handler:    _Tx_GetBalance_Handler,
		},
		{
			MethodName: "GetTransactionByHash",
			Handler:    _Tx_GetTransactionByHash_Handler,
		},
		{
			MethodName: "GetTransactionList",
			Handler:    _Tx_GetTransactionList_Handler,
		},
		{
			MethodName: "GetAssetList",
			Handler:    _Tx_GetAssetList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "builder/v1/tx.proto",
}
