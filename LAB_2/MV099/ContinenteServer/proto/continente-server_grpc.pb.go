// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// ComunicacionClient is the client API for Comunicacion service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ComunicacionClient interface {
	EnviarEncriptado(ctx context.Context, in *Solicitud, opts ...grpc.CallOption) (*Respuesta, error)
}

type comunicacionClient struct {
	cc grpc.ClientConnInterface
}

func NewComunicacionClient(cc grpc.ClientConnInterface) ComunicacionClient {
	return &comunicacionClient{cc}
}

func (c *comunicacionClient) EnviarEncriptado(ctx context.Context, in *Solicitud, opts ...grpc.CallOption) (*Respuesta, error) {
	out := new(Respuesta)
	err := c.cc.Invoke(ctx, "/server.Comunicacion/EnviarEncriptado", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ComunicacionServer is the server API for Comunicacion service.
// All implementations must embed UnimplementedComunicacionServer
// for forward compatibility
type ComunicacionServer interface {
	EnviarEncriptado(context.Context, *Solicitud) (*Respuesta, error)
	mustEmbedUnimplementedComunicacionServer()
}

// UnimplementedComunicacionServer must be embedded to have forward compatible implementations.
type UnimplementedComunicacionServer struct {
}

func (UnimplementedComunicacionServer) EnviarEncriptado(context.Context, *Solicitud) (*Respuesta, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnviarEncriptado not implemented")
}
func (UnimplementedComunicacionServer) mustEmbedUnimplementedComunicacionServer() {}

// UnsafeComunicacionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ComunicacionServer will
// result in compilation errors.
type UnsafeComunicacionServer interface {
	mustEmbedUnimplementedComunicacionServer()
}

func RegisterComunicacionServer(s grpc.ServiceRegistrar, srv ComunicacionServer) {
	s.RegisterService(&Comunicacion_ServiceDesc, srv)
}

func _Comunicacion_EnviarEncriptado_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Solicitud)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComunicacionServer).EnviarEncriptado(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Comunicacion/EnviarEncriptado",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComunicacionServer).EnviarEncriptado(ctx, req.(*Solicitud))
	}
	return interceptor(ctx, in, info, handler)
}

// Comunicacion_ServiceDesc is the grpc.ServiceDesc for Comunicacion service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Comunicacion_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "server.Comunicacion",
	HandlerType: (*ComunicacionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EnviarEncriptado",
			Handler:    _Comunicacion_EnviarEncriptado_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "continente-server.proto",
}
