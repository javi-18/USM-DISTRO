// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: Clientes/clientes.proto

package proto

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
	ClientesService_EnviarPedido_FullMethodName = "/clientes.ClientesService/EnviarPedido"
)

// ClientesServiceClient is the client API for ClientesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientesServiceClient interface {
	EnviarPedido(ctx context.Context, in *Pedido, opts ...grpc.CallOption) (*Respuesta, error)
}

type clientesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientesServiceClient(cc grpc.ClientConnInterface) ClientesServiceClient {
	return &clientesServiceClient{cc}
}

func (c *clientesServiceClient) EnviarPedido(ctx context.Context, in *Pedido, opts ...grpc.CallOption) (*Respuesta, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Respuesta)
	err := c.cc.Invoke(ctx, ClientesService_EnviarPedido_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientesServiceServer is the server API for ClientesService service.
// All implementations must embed UnimplementedClientesServiceServer
// for forward compatibility.
type ClientesServiceServer interface {
	EnviarPedido(context.Context, *Pedido) (*Respuesta, error)
	mustEmbedUnimplementedClientesServiceServer()
}

// UnimplementedClientesServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedClientesServiceServer struct{}

func (UnimplementedClientesServiceServer) EnviarPedido(context.Context, *Pedido) (*Respuesta, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnviarPedido not implemented")
}
func (UnimplementedClientesServiceServer) mustEmbedUnimplementedClientesServiceServer() {}
func (UnimplementedClientesServiceServer) testEmbeddedByValue()                         {}

// UnsafeClientesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientesServiceServer will
// result in compilation errors.
type UnsafeClientesServiceServer interface {
	mustEmbedUnimplementedClientesServiceServer()
}

func RegisterClientesServiceServer(s grpc.ServiceRegistrar, srv ClientesServiceServer) {
	// If the following call pancis, it indicates UnimplementedClientesServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ClientesService_ServiceDesc, srv)
}

func _ClientesService_EnviarPedido_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Pedido)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientesServiceServer).EnviarPedido(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientesService_EnviarPedido_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientesServiceServer).EnviarPedido(ctx, req.(*Pedido))
	}
	return interceptor(ctx, in, info, handler)
}

// ClientesService_ServiceDesc is the grpc.ServiceDesc for ClientesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "clientes.ClientesService",
	HandlerType: (*ClientesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EnviarPedido",
			Handler:    _ClientesService_EnviarPedido_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Clientes/clientes.proto",
}
