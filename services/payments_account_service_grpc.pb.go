// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: google/ads/googleads/v11/services/payments_account_service.proto

package services

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

// PaymentsAccountServiceClient is the client API for PaymentsAccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaymentsAccountServiceClient interface {
	// Returns all payments accounts associated with all managers
	// between the login customer ID and specified serving customer in the
	// hierarchy, inclusive.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [PaymentsAccountError]()
	//   [QuotaError]()
	//   [RequestError]()
	ListPaymentsAccounts(ctx context.Context, in *ListPaymentsAccountsRequest, opts ...grpc.CallOption) (*ListPaymentsAccountsResponse, error)
}

type paymentsAccountServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentsAccountServiceClient(cc grpc.ClientConnInterface) PaymentsAccountServiceClient {
	return &paymentsAccountServiceClient{cc}
}

func (c *paymentsAccountServiceClient) ListPaymentsAccounts(ctx context.Context, in *ListPaymentsAccountsRequest, opts ...grpc.CallOption) (*ListPaymentsAccountsResponse, error) {
	out := new(ListPaymentsAccountsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v11.services.PaymentsAccountService/ListPaymentsAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentsAccountServiceServer is the server API for PaymentsAccountService service.
// All implementations must embed UnimplementedPaymentsAccountServiceServer
// for forward compatibility
type PaymentsAccountServiceServer interface {
	// Returns all payments accounts associated with all managers
	// between the login customer ID and specified serving customer in the
	// hierarchy, inclusive.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [PaymentsAccountError]()
	//   [QuotaError]()
	//   [RequestError]()
	ListPaymentsAccounts(context.Context, *ListPaymentsAccountsRequest) (*ListPaymentsAccountsResponse, error)
	mustEmbedUnimplementedPaymentsAccountServiceServer()
}

// UnimplementedPaymentsAccountServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPaymentsAccountServiceServer struct {
}

func (UnimplementedPaymentsAccountServiceServer) ListPaymentsAccounts(context.Context, *ListPaymentsAccountsRequest) (*ListPaymentsAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPaymentsAccounts not implemented")
}
func (UnimplementedPaymentsAccountServiceServer) mustEmbedUnimplementedPaymentsAccountServiceServer() {
}

// UnsafePaymentsAccountServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaymentsAccountServiceServer will
// result in compilation errors.
type UnsafePaymentsAccountServiceServer interface {
	mustEmbedUnimplementedPaymentsAccountServiceServer()
}

func RegisterPaymentsAccountServiceServer(s grpc.ServiceRegistrar, srv PaymentsAccountServiceServer) {
	s.RegisterService(&PaymentsAccountService_ServiceDesc, srv)
}

func _PaymentsAccountService_ListPaymentsAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPaymentsAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentsAccountServiceServer).ListPaymentsAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v11.services.PaymentsAccountService/ListPaymentsAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentsAccountServiceServer).ListPaymentsAccounts(ctx, req.(*ListPaymentsAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PaymentsAccountService_ServiceDesc is the grpc.ServiceDesc for PaymentsAccountService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PaymentsAccountService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v11.services.PaymentsAccountService",
	HandlerType: (*PaymentsAccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListPaymentsAccounts",
			Handler:    _PaymentsAccountService_ListPaymentsAccounts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v11/services/payments_account_service.proto",
}