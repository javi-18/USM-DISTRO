package main

import (
	"context"
	"fmt"
	"log"
	"net"

	proto "Logistica/proto"

	pb "Caravanas/proto"

	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port        = ":50051"
	rabbitMQURL = "amqp://guest:guest@rabbitmq:5672/"
)

type server struct {
	proto.UnimplementedLogisticaServiceServer
}

func (s *server) EnviarPedido(ctx context.Context, req *pb.Pedido) (*proto.Respuesta, error) {
	fmt.Printf("Mensaje recibido: %v\n", req)

	conn, err := grpc.Dial("dist099:50052", grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("no se pudo conectar con caravanas: %v", err)
	}
	defer conn.Close()

	caravanasClient := pb.NewCaravanasServiceClient(conn)
	_, err = caravanasClient.ProcesarPedido(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("no se pudo enviar el mensaje a caravanas: %v", err)
	}

	rabbitConn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		return nil, fmt.Errorf("no se pudo conectar a RabbitMQ: %v", err)
	}
	defer rabbitConn.Close()

	ch, err := rabbitConn.Channel()
	if err != nil {
		return nil, fmt.Errorf("no se pudo abrir el canal: %v", err)
	}
	defer ch.Close()

	err = ch.Publish(
		"",
		"raquis_queue",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(fmt.Sprintf("%v", req)),
		})
	if err != nil {
		return nil, fmt.Errorf("no se pudo publicar el mensaje: %v", err)
	}

	return &proto.Respuesta{Seguimiento: req.Seguimiento}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("error al escuchar: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterLogisticaServiceServer(s, &server{})
	reflection.Register(s)

	fmt.Println("Servicio Logística corriendo en el puerto", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("error al servir: %v", err)
	}
}
