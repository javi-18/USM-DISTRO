// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// NodoTaiClient is the client API for NodoTai service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodoTaiClient interface {
	SolicitarDatos(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*BattleResponse, error)
	Atacar(ctx context.Context, in *BattleRequest, opts ...grpc.CallOption) (*BattleResponse, error)
}

type nodoTaiClient struct {
	cc grpc.ClientConnInterface
}

func NewNodoTaiClient(cc grpc.ClientConnInterface) NodoTaiClient {
	return &nodoTaiClient{cc}
}

func (c *nodoTaiClient) SolicitarDatos(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*BattleResponse, error) {
	out := new(BattleResponse)
	err := c.cc.Invoke(ctx, "/battle.NodoTai/SolicitarDatos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodoTaiClient) Atacar(ctx context.Context, in *BattleRequest, opts ...grpc.CallOption) (*BattleResponse, error) {
	out := new(BattleResponse)
	err := c.cc.Invoke(ctx, "/battle.NodoTai/Atacar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodoTaiServer is the server API for NodoTai service.
// All implementations must embed UnimplementedNodoTaiServer
// for forward compatibility
type NodoTaiServer interface {
	SolicitarDatos(context.Context, *empty.Empty) (*BattleResponse, error)
	Atacar(context.Context, *BattleRequest) (*BattleResponse, error)
	mustEmbedUnimplementedNodoTaiServer()
}

// UnimplementedNodoTaiServer must be embedded to have forward compatible implementations.
type UnimplementedNodoTaiServer struct {
}

func (UnimplementedNodoTaiServer) SolicitarDatos(context.Context, *empty.Empty) (*BattleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SolicitarDatos not implemented")
}
func (UnimplementedNodoTaiServer) Atacar(context.Context, *BattleRequest) (*BattleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Atacar not implemented")
}
func (UnimplementedNodoTaiServer) mustEmbedUnimplementedNodoTaiServer() {}

// UnsafeNodoTaiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodoTaiServer will
// result in compilation errors.
type UnsafeNodoTaiServer interface {
	mustEmbedUnimplementedNodoTaiServer()
}

func RegisterNodoTaiServer(s grpc.ServiceRegistrar, srv NodoTaiServer) {
	s.RegisterService(&NodoTai_ServiceDesc, srv)
}

func _NodoTai_SolicitarDatos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodoTaiServer).SolicitarDatos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/battle.NodoTai/SolicitarDatos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodoTaiServer).SolicitarDatos(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodoTai_Atacar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BattleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodoTaiServer).Atacar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/battle.NodoTai/Atacar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodoTaiServer).Atacar(ctx, req.(*BattleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NodoTai_ServiceDesc is the grpc.ServiceDesc for NodoTai service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NodoTai_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "battle.NodoTai",
	HandlerType: (*NodoTaiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SolicitarDatos",
			Handler:    _NodoTai_SolicitarDatos_Handler,
		},
		{
			MethodName: "Atacar",
			Handler:    _NodoTai_Atacar_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nodo-tai.proto",
}
