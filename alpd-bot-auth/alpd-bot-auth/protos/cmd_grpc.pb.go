// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CommandServiceClient is the client API for CommandService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommandServiceClient interface {
	ExecuteCommandV1(ctx context.Context, in *CmdRequest, opts ...grpc.CallOption) (*CmdReply, error)
}

type commandServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommandServiceClient(cc grpc.ClientConnInterface) CommandServiceClient {
	return &commandServiceClient{cc}
}

func (c *commandServiceClient) ExecuteCommandV1(ctx context.Context, in *CmdRequest, opts ...grpc.CallOption) (*CmdReply, error) {
	out := new(CmdReply)
	err := c.cc.Invoke(ctx, "/alpd.CommandService/ExecuteCommandV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommandServiceServer is the server API for CommandService service.
// All implementations must embed UnimplementedCommandServiceServer
// for forward compatibility
type CommandServiceServer interface {
	ExecuteCommandV1(context.Context, *CmdRequest) (*CmdReply, error)
	mustEmbedUnimplementedCommandServiceServer()
}

// UnimplementedCommandServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCommandServiceServer struct {
}

func (*UnimplementedCommandServiceServer) ExecuteCommandV1(context.Context, *CmdRequest) (*CmdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteCommandV1 not implemented")
}
func (*UnimplementedCommandServiceServer) mustEmbedUnimplementedCommandServiceServer() {}

func RegisterCommandServiceServer(s *grpc.Server, srv CommandServiceServer) {
	s.RegisterService(&_CommandService_serviceDesc, srv)
}

func _CommandService_ExecuteCommandV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CmdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandServiceServer).ExecuteCommandV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alpd.CommandService/ExecuteCommandV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandServiceServer).ExecuteCommandV1(ctx, req.(*CmdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CommandService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "alpd.CommandService",
	HandlerType: (*CommandServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExecuteCommandV1",
			Handler:    _CommandService_ExecuteCommandV1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cmd.proto",
}