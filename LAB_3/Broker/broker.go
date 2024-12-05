package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"

	pb "Broker/proto"

	"google.golang.org/grpc"
)

type BrokerServer struct {
	pb.UnimplementedBrokerServiceServer
	servers []string
	mu      sync.Mutex
}

func NewBrokerServer() *BrokerServer {
	return &BrokerServer{
		servers: []string{"dist100:5004", "dist100:5005", "dist100:5006"},
	}
}

func (s *BrokerServer) RedirectRequest(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	log.Printf("Recibiendo comando del tipo: %s", req.CommandType)

	switch req.CommandType {
	case "Supervisor":
		server := s.selectServerForSupervisors(req.VectorClock)
		return s.forwardRequest(ctx, server, req)
	case "Jayce":
		server := s.selectServerForJayce(req.VectorClock)
		return s.forwardRequest(ctx, server, req)
	default:
		server := s.selectRandomServer()
		return s.forwardRequest(ctx, server, req)
	}
}

func (s *BrokerServer) HandleInconsistency(ctx context.Context, req *pb.InconsistencyRequest) (*pb.Response, error) {
	log.Println("Manejando inconsistencia...")
	server := s.determineServer(req.VectorClock)

	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		log.Printf("Error al conectar con el servidor %s: %v", server, err)
		return nil, err
	}
	defer conn.Close()

	client := pb.NewHextechServiceClient(conn)
	res, err := client.ResolveInconsistency(ctx, req)
	if err != nil {
		log.Printf("Error al resolver inconsistencia: %v", err)
		return nil, err
	}

	log.Printf("Inconsistencia resuelta por el servidor: %s", server)
	return res, nil
}

func (s *BrokerServer) selectRandomServer() string {
	rand.Seed(time.Now().UnixNano())
	return s.servers[rand.Intn(len(s.servers))]
}

func (s *BrokerServer) selectServerForSupervisors(vectorClock []*pb.VectorClockEntry) string {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, server := range s.servers {
		if s.isCompatibleWithReadYourWrites(server, vectorClock) {
			log.Printf("Servidor seleccionado (Read Your Writes): %s", server)
			return server
		}
	}
	log.Println("No se encontró un servidor compatible. Usando el servidor por defecto.")
	return s.servers[0]
}

func (s *BrokerServer) selectServerForJayce(vectorClock []*pb.VectorClockEntry) string {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, server := range s.servers {
		if s.isCompatibleWithMonotonicReads(server, vectorClock) {
			log.Printf("Servidor seleccionado (Monotonic Reads): %s", server)
			return server
		}
	}
	log.Println("No se encontró un servidor compatible. Usando el servidor por defecto.")
	return s.servers[0]
}

func (s *BrokerServer) determineServer(vectorClock []*pb.VectorClockEntry) string {
	for _, server := range s.servers {
		if s.isMostUpdatedServer(server, vectorClock) {
			log.Printf("Servidor determinado para manejar inconsistencia: %s", server)
			return server
		}
	}
	log.Println("No se encontró un servidor actualizado. Usando el servidor por defecto.")
	return s.servers[0]
}

func (s *BrokerServer) forwardRequest(ctx context.Context, server string, req *pb.Request) (*pb.Response, error) {
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		log.Printf("Error al conectar con el servidor %s: %v", server, err)
		return nil, err
	}
	defer conn.Close()

	client := pb.NewHextechServiceClient(conn)
	res, err := client.HandleRequest(ctx, req)
	if err != nil {
		log.Printf("Error al manejar la solicitud en el servidor %s: %v", server, err)
		return nil, err
	}

	log.Printf("Solicitud manejada exitosamente por el servidor %s", server)
	return res, nil
}

func (s *BrokerServer) getServerVectorClock(server string) ([]*pb.VectorClockEntry, error) {
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("error al conectar con el servidor %s: %v", server, err)
	}
	defer conn.Close()

	client := pb.NewHextechServiceClient(conn)
	res, err := client.GetVectorClock(context.Background(), &pb.Empty{})
	if err != nil {
		return nil, fmt.Errorf("error al obtener reloj vectorial del servidor %s: %v", server, err)
	}

	return res.VectorClock, nil
}

func (s *BrokerServer) isCompatibleWithReadYourWrites(server string, clientVectorClock []*pb.VectorClockEntry) bool {
	serverVectorClock, err := s.getServerVectorClock(server)
	if err != nil {
		log.Printf("Error al obtener reloj vectorial del servidor %s: %v", server, err)
		return false
	}

	for i := 0; i < len(clientVectorClock); i++ {
		for j := 0; j < len(serverVectorClock); j++ {
			if clientVectorClock[i].Key == serverVectorClock[j].Key && serverVectorClock[j].Values[0] < clientVectorClock[i].Values[0] {
				return false
			}
		}
	}
	return true
}

func (s *BrokerServer) isCompatibleWithMonotonicReads(server string, clientVectorClock []*pb.VectorClockEntry) bool {
	serverVectorClock, err := s.getServerVectorClock(server)
	if err != nil {
		log.Printf("Error al obtener reloj vectorial del servidor %s: %v", server, err)
		return false
	}

	for i := 0; i < len(clientVectorClock); i++ {
		for j := 0; j < len(serverVectorClock); j++ {
			if clientVectorClock[i].Key == serverVectorClock[j].Key && serverVectorClock[j].Values[0] < clientVectorClock[i].Values[0] {
				return false
			}
		}
	}
	return true
}

func (s *BrokerServer) isMostUpdatedServer(server string, clientVectorClock []*pb.VectorClockEntry) bool {
	serverVectorClock, err := s.getServerVectorClock(server)
	if err != nil {
		log.Printf("Error al obtener reloj vectorial del servidor %s: %v", server, err)
		return false
	}

	for i := 0; i < len(clientVectorClock); i++ {
		for j := 0; j < len(serverVectorClock); j++ {
			if clientVectorClock[i].Key == serverVectorClock[j].Key && serverVectorClock[j].Values[0] < clientVectorClock[i].Values[0] {
				return false
			}
		}
	}
	return true
}

func main() {
	listener, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("Error al iniciar el listener: %v", err)
	}
	fmt.Println("Broker corriendo en el puerto 5001")
	grpcServer := grpc.NewServer()
	broker := NewBrokerServer()
	pb.RegisterBrokerServiceServer(grpcServer, broker)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Error al iniciar el servidor gRPC: %v", err)
	}
}
