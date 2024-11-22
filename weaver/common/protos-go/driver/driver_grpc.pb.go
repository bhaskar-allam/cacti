// Copyright IBM Corp. All Rights Reserved.
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: driver/driver.proto

package driver

import (
	context "context"
	common "github.com/hyperledger-cacti/cacti/weaver/common/protos-go/v2/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	DriverCommunication_RequestDriverState_FullMethodName                  = "/driver.driver.DriverCommunication/RequestDriverState"
	DriverCommunication_SubscribeEvent_FullMethodName                      = "/driver.driver.DriverCommunication/SubscribeEvent"
	DriverCommunication_RequestSignedEventSubscriptionQuery_FullMethodName = "/driver.driver.DriverCommunication/RequestSignedEventSubscriptionQuery"
	DriverCommunication_WriteExternalState_FullMethodName                  = "/driver.driver.DriverCommunication/WriteExternalState"
	DriverCommunication_PerformLock_FullMethodName                         = "/driver.driver.DriverCommunication/PerformLock"
	DriverCommunication_CreateAsset_FullMethodName                         = "/driver.driver.DriverCommunication/CreateAsset"
	DriverCommunication_Extinguish_FullMethodName                          = "/driver.driver.DriverCommunication/Extinguish"
	DriverCommunication_AssignAsset_FullMethodName                         = "/driver.driver.DriverCommunication/AssignAsset"
)

// DriverCommunicationClient is the client API for DriverCommunication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DriverCommunicationClient interface {
	// Data Sharing
	// the remote relay sends a RequestDriverState request to its driver with a
	// query defining the data it wants to receive
	RequestDriverState(ctx context.Context, in *common.Query, opts ...grpc.CallOption) (*common.Ack, error)
	// Events Subscription
	// the src-relay uses this endpoint to forward the event subscription request from dest-relay to driver
	SubscribeEvent(ctx context.Context, in *common.EventSubscription, opts ...grpc.CallOption) (*common.Ack, error)
	// Recommended to have TLS mode on for this unsafe endpoint
	// Relay uses this to get Query.requestor_signature and
	// Query.certificate required for event subscription
	RequestSignedEventSubscriptionQuery(ctx context.Context, in *common.EventSubscription, opts ...grpc.CallOption) (*common.Query, error)
	// Events Publication
	// the dest-relay calls the dest-driver on this end point to write the remote network state to the local ledger
	WriteExternalState(ctx context.Context, in *WriteExternalStateMessage, opts ...grpc.CallOption) (*common.Ack, error)
	// As part of SATP, the source reply (sender gateway) sends a PerformLock request to its driver
	// to lock a specific asset
	PerformLock(ctx context.Context, in *PerformLockRequest, opts ...grpc.CallOption) (*common.Ack, error)
	// As part of SATP, the destination reply (receiver gateway) sends a CreateAsset request to its driver
	// to create a specific asset
	CreateAsset(ctx context.Context, in *CreateAssetRequest, opts ...grpc.CallOption) (*common.Ack, error)
	// As part of SATP, the source reply (sender gateway) sends a Extinguish request to its driver
	// to extinguish a specific asset
	Extinguish(ctx context.Context, in *ExtinguishRequest, opts ...grpc.CallOption) (*common.Ack, error)
	// As part of SATP, the destination reply (receiver gateway) sends a AssignAsset request to its driver
	// to assign a specific asset
	AssignAsset(ctx context.Context, in *AssignAssetRequest, opts ...grpc.CallOption) (*common.Ack, error)
}

type driverCommunicationClient struct {
	cc grpc.ClientConnInterface
}

func NewDriverCommunicationClient(cc grpc.ClientConnInterface) DriverCommunicationClient {
	return &driverCommunicationClient{cc}
}

func (c *driverCommunicationClient) RequestDriverState(ctx context.Context, in *common.Query, opts ...grpc.CallOption) (*common.Ack, error) {
	out := new(common.Ack)
	err := c.cc.Invoke(ctx, DriverCommunication_RequestDriverState_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverCommunicationClient) SubscribeEvent(ctx context.Context, in *common.EventSubscription, opts ...grpc.CallOption) (*common.Ack, error) {
	out := new(common.Ack)
	err := c.cc.Invoke(ctx, DriverCommunication_SubscribeEvent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverCommunicationClient) RequestSignedEventSubscriptionQuery(ctx context.Context, in *common.EventSubscription, opts ...grpc.CallOption) (*common.Query, error) {
	out := new(common.Query)
	err := c.cc.Invoke(ctx, DriverCommunication_RequestSignedEventSubscriptionQuery_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverCommunicationClient) WriteExternalState(ctx context.Context, in *WriteExternalStateMessage, opts ...grpc.CallOption) (*common.Ack, error) {
	out := new(common.Ack)
	err := c.cc.Invoke(ctx, DriverCommunication_WriteExternalState_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverCommunicationClient) PerformLock(ctx context.Context, in *PerformLockRequest, opts ...grpc.CallOption) (*common.Ack, error) {
	out := new(common.Ack)
	err := c.cc.Invoke(ctx, DriverCommunication_PerformLock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverCommunicationClient) CreateAsset(ctx context.Context, in *CreateAssetRequest, opts ...grpc.CallOption) (*common.Ack, error) {
	out := new(common.Ack)
	err := c.cc.Invoke(ctx, DriverCommunication_CreateAsset_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverCommunicationClient) Extinguish(ctx context.Context, in *ExtinguishRequest, opts ...grpc.CallOption) (*common.Ack, error) {
	out := new(common.Ack)
	err := c.cc.Invoke(ctx, DriverCommunication_Extinguish_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverCommunicationClient) AssignAsset(ctx context.Context, in *AssignAssetRequest, opts ...grpc.CallOption) (*common.Ack, error) {
	out := new(common.Ack)
	err := c.cc.Invoke(ctx, DriverCommunication_AssignAsset_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DriverCommunicationServer is the server API for DriverCommunication service.
// All implementations must embed UnimplementedDriverCommunicationServer
// for forward compatibility
type DriverCommunicationServer interface {
	// Data Sharing
	// the remote relay sends a RequestDriverState request to its driver with a
	// query defining the data it wants to receive
	RequestDriverState(context.Context, *common.Query) (*common.Ack, error)
	// Events Subscription
	// the src-relay uses this endpoint to forward the event subscription request from dest-relay to driver
	SubscribeEvent(context.Context, *common.EventSubscription) (*common.Ack, error)
	// Recommended to have TLS mode on for this unsafe endpoint
	// Relay uses this to get Query.requestor_signature and
	// Query.certificate required for event subscription
	RequestSignedEventSubscriptionQuery(context.Context, *common.EventSubscription) (*common.Query, error)
	// Events Publication
	// the dest-relay calls the dest-driver on this end point to write the remote network state to the local ledger
	WriteExternalState(context.Context, *WriteExternalStateMessage) (*common.Ack, error)
	// As part of SATP, the source reply (sender gateway) sends a PerformLock request to its driver
	// to lock a specific asset
	PerformLock(context.Context, *PerformLockRequest) (*common.Ack, error)
	// As part of SATP, the destination reply (receiver gateway) sends a CreateAsset request to its driver
	// to create a specific asset
	CreateAsset(context.Context, *CreateAssetRequest) (*common.Ack, error)
	// As part of SATP, the source reply (sender gateway) sends a Extinguish request to its driver
	// to extinguish a specific asset
	Extinguish(context.Context, *ExtinguishRequest) (*common.Ack, error)
	// As part of SATP, the destination reply (receiver gateway) sends a AssignAsset request to its driver
	// to assign a specific asset
	AssignAsset(context.Context, *AssignAssetRequest) (*common.Ack, error)
	mustEmbedUnimplementedDriverCommunicationServer()
}

// UnimplementedDriverCommunicationServer must be embedded to have forward compatible implementations.
type UnimplementedDriverCommunicationServer struct {
}

func (UnimplementedDriverCommunicationServer) RequestDriverState(context.Context, *common.Query) (*common.Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestDriverState not implemented")
}
func (UnimplementedDriverCommunicationServer) SubscribeEvent(context.Context, *common.EventSubscription) (*common.Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubscribeEvent not implemented")
}
func (UnimplementedDriverCommunicationServer) RequestSignedEventSubscriptionQuery(context.Context, *common.EventSubscription) (*common.Query, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestSignedEventSubscriptionQuery not implemented")
}
func (UnimplementedDriverCommunicationServer) WriteExternalState(context.Context, *WriteExternalStateMessage) (*common.Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteExternalState not implemented")
}
func (UnimplementedDriverCommunicationServer) PerformLock(context.Context, *PerformLockRequest) (*common.Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PerformLock not implemented")
}
func (UnimplementedDriverCommunicationServer) CreateAsset(context.Context, *CreateAssetRequest) (*common.Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAsset not implemented")
}
func (UnimplementedDriverCommunicationServer) Extinguish(context.Context, *ExtinguishRequest) (*common.Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Extinguish not implemented")
}
func (UnimplementedDriverCommunicationServer) AssignAsset(context.Context, *AssignAssetRequest) (*common.Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignAsset not implemented")
}
func (UnimplementedDriverCommunicationServer) mustEmbedUnimplementedDriverCommunicationServer() {}

// UnsafeDriverCommunicationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DriverCommunicationServer will
// result in compilation errors.
type UnsafeDriverCommunicationServer interface {
	mustEmbedUnimplementedDriverCommunicationServer()
}

func RegisterDriverCommunicationServer(s grpc.ServiceRegistrar, srv DriverCommunicationServer) {
	s.RegisterService(&DriverCommunication_ServiceDesc, srv)
}

func _DriverCommunication_RequestDriverState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverCommunicationServer).RequestDriverState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DriverCommunication_RequestDriverState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverCommunicationServer).RequestDriverState(ctx, req.(*common.Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverCommunication_SubscribeEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.EventSubscription)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverCommunicationServer).SubscribeEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DriverCommunication_SubscribeEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverCommunicationServer).SubscribeEvent(ctx, req.(*common.EventSubscription))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverCommunication_RequestSignedEventSubscriptionQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.EventSubscription)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverCommunicationServer).RequestSignedEventSubscriptionQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DriverCommunication_RequestSignedEventSubscriptionQuery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverCommunicationServer).RequestSignedEventSubscriptionQuery(ctx, req.(*common.EventSubscription))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverCommunication_WriteExternalState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteExternalStateMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverCommunicationServer).WriteExternalState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DriverCommunication_WriteExternalState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverCommunicationServer).WriteExternalState(ctx, req.(*WriteExternalStateMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverCommunication_PerformLock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PerformLockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverCommunicationServer).PerformLock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DriverCommunication_PerformLock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverCommunicationServer).PerformLock(ctx, req.(*PerformLockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverCommunication_CreateAsset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAssetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverCommunicationServer).CreateAsset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DriverCommunication_CreateAsset_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverCommunicationServer).CreateAsset(ctx, req.(*CreateAssetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverCommunication_Extinguish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExtinguishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverCommunicationServer).Extinguish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DriverCommunication_Extinguish_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverCommunicationServer).Extinguish(ctx, req.(*ExtinguishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverCommunication_AssignAsset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignAssetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverCommunicationServer).AssignAsset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DriverCommunication_AssignAsset_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverCommunicationServer).AssignAsset(ctx, req.(*AssignAssetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DriverCommunication_ServiceDesc is the grpc.ServiceDesc for DriverCommunication service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DriverCommunication_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "driver.driver.DriverCommunication",
	HandlerType: (*DriverCommunicationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestDriverState",
			Handler:    _DriverCommunication_RequestDriverState_Handler,
		},
		{
			MethodName: "SubscribeEvent",
			Handler:    _DriverCommunication_SubscribeEvent_Handler,
		},
		{
			MethodName: "RequestSignedEventSubscriptionQuery",
			Handler:    _DriverCommunication_RequestSignedEventSubscriptionQuery_Handler,
		},
		{
			MethodName: "WriteExternalState",
			Handler:    _DriverCommunication_WriteExternalState_Handler,
		},
		{
			MethodName: "PerformLock",
			Handler:    _DriverCommunication_PerformLock_Handler,
		},
		{
			MethodName: "CreateAsset",
			Handler:    _DriverCommunication_CreateAsset_Handler,
		},
		{
			MethodName: "Extinguish",
			Handler:    _DriverCommunication_Extinguish_Handler,
		},
		{
			MethodName: "AssignAsset",
			Handler:    _DriverCommunication_AssignAsset_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "driver/driver.proto",
}
