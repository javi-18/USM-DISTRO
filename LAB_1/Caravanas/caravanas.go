package main

import (
	"context"
	"fmt"
	"log"
	"net"

	proto "Caravanas/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50052"
)

type server struct {
	proto.UnimplementedCaravanasServiceServer
}

func (s *server) ProcesarPedido(ctx context.Context, req *proto.Pedido) (*proto.Estado, error) {
	fmt.Printf("Procesando mensaje: %v\n", req)

	return &proto.Estado{Seguimiento: 1}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("error al escuchar: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterCaravanasServiceServer(s, &server{})
	reflection.Register(s)

	fmt.Println("Servicio Caravanas corriendo en el puerto", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("error al servir: %v", err)
	}
}
