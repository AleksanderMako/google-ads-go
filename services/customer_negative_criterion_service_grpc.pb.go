// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: google/ads/googleads/v11/services/customer_negative_criterion_service.proto

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

// CustomerNegativeCriterionServiceClient is the client API for CustomerNegativeCriterionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomerNegativeCriterionServiceClient interface {
	// Creates or removes criteria. Operation statuses are returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [CriterionError]()
	//   [DatabaseError]()
	//   [FieldError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [MutateError]()
	//   [QuotaError]()
	//   [RequestError]()
	MutateCustomerNegativeCriteria(ctx context.Context, in *MutateCustomerNegativeCriteriaRequest, opts ...grpc.CallOption) (*MutateCustomerNegativeCriteriaResponse, error)
}

type customerNegativeCriterionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerNegativeCriterionServiceClient(cc grpc.ClientConnInterface) CustomerNegativeCriterionServiceClient {
	return &customerNegativeCriterionServiceClient{cc}
}

func (c *customerNegativeCriterionServiceClient) MutateCustomerNegativeCriteria(ctx context.Context, in *MutateCustomerNegativeCriteriaRequest, opts ...grpc.CallOption) (*MutateCustomerNegativeCriteriaResponse, error) {
	out := new(MutateCustomerNegativeCriteriaResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v11.services.CustomerNegativeCriterionService/MutateCustomerNegativeCriteria", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerNegativeCriterionServiceServer is the server API for CustomerNegativeCriterionService service.
// All implementations must embed UnimplementedCustomerNegativeCriterionServiceServer
// for forward compatibility
type CustomerNegativeCriterionServiceServer interface {
	// Creates or removes criteria. Operation statuses are returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [CriterionError]()
	//   [DatabaseError]()
	//   [FieldError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [MutateError]()
	//   [QuotaError]()
	//   [RequestError]()
	MutateCustomerNegativeCriteria(context.Context, *MutateCustomerNegativeCriteriaRequest) (*MutateCustomerNegativeCriteriaResponse, error)
	mustEmbedUnimplementedCustomerNegativeCriterionServiceServer()
}

// UnimplementedCustomerNegativeCriterionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCustomerNegativeCriterionServiceServer struct {
}

func (UnimplementedCustomerNegativeCriterionServiceServer) MutateCustomerNegativeCriteria(context.Context, *MutateCustomerNegativeCriteriaRequest) (*MutateCustomerNegativeCriteriaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateCustomerNegativeCriteria not implemented")
}
func (UnimplementedCustomerNegativeCriterionServiceServer) mustEmbedUnimplementedCustomerNegativeCriterionServiceServer() {
}

// UnsafeCustomerNegativeCriterionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomerNegativeCriterionServiceServer will
// result in compilation errors.
type UnsafeCustomerNegativeCriterionServiceServer interface {
	mustEmbedUnimplementedCustomerNegativeCriterionServiceServer()
}

func RegisterCustomerNegativeCriterionServiceServer(s grpc.ServiceRegistrar, srv CustomerNegativeCriterionServiceServer) {
	s.RegisterService(&CustomerNegativeCriterionService_ServiceDesc, srv)
}

func _CustomerNegativeCriterionService_MutateCustomerNegativeCriteria_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCustomerNegativeCriteriaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerNegativeCriterionServiceServer).MutateCustomerNegativeCriteria(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v11.services.CustomerNegativeCriterionService/MutateCustomerNegativeCriteria",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerNegativeCriterionServiceServer).MutateCustomerNegativeCriteria(ctx, req.(*MutateCustomerNegativeCriteriaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomerNegativeCriterionService_ServiceDesc is the grpc.ServiceDesc for CustomerNegativeCriterionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomerNegativeCriterionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v11.services.CustomerNegativeCriterionService",
	HandlerType: (*CustomerNegativeCriterionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateCustomerNegativeCriteria",
			Handler:    _CustomerNegativeCriterionService_MutateCustomerNegativeCriteria_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v11/services/customer_negative_criterion_service.proto",
}
