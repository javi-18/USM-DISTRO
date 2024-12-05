package main

import (
	pb "NodoTai/proto"
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	"google.golang.org/protobuf/types/known/emptypb"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedNodoTaiServer
	vida           int
	datosRecibidos int
	CD             int
}

func (s *server) Atacar(ctx context.Context, req *pb.BattleRequest) (*pb.BattleResponse, error) {
	log.Printf("Recibiendo ataque de %d puntos de daño", req.Damage)
	s.vida -= int(req.Damage)

	if s.vida <= 0 {
		log.Println("Greymon/Garurumon han sido derrotados.")
		return &pb.BattleResponse{Result: "Nodo Tai ha sido derrotado."}, nil
	}

	return &pb.BattleResponse{Result: fmt.Sprintf("Ataque recibido. Vida restante: %d", s.vida)}, nil
}

func (s *server) SolicitarDatos(ctx context.Context, req *emptypb.Empty) (*pb.BattleResponse, error) {
	log.Printf("Solicitando datos al Primary Node...")
	s.datosRecibidos += 10

	log.Printf("Datos recibidos: %d", s.datosRecibidos)
	return &pb.BattleResponse{Result: fmt.Sprintf("Datos recibidos: %d", s.datosRecibidos)}, nil
}

func (s *server) AtacarDiaboromon(ctx context.Context, req *pb.BattleRequest) (*pb.BattleResponse, error) {
	if s.datosRecibidos >= s.CD {
		log.Println("Ataque exitoso. Diaboromon derrotado.")
		return &pb.BattleResponse{Result: "Diaboromon derrotado. ¡Victoria!"}, nil
	} else {
		log.Println("Ataque fallido. Diaboromon repele el ataque.")
		s.vida -= 10
		return &pb.BattleResponse{Result: fmt.Sprintf("Ataque fallido. Vida restante: %d", s.vida)}, nil
	}
}

func main() {
	file, err := os.Open("/app/INPUT.txt")
	if err != nil {
		log.Fatalf("Error al abrir INPUT.txt: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	valores := []int{}
	for scanner.Scan() {
		linea := scanner.Text()
		partes := stringToIntSlice(linea)
		valores = append(valores, partes...)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error leyendo INPUT.txt: %v", err)
	}

	nodoTai := &server{
		vida:           valores[0],
		datosRecibidos: 0,
		CD:             valores[2],
	}

	connDiaboromon, err := grpc.Dial("diaboromon:5004", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error de conexión con Diaboromon: %v", err)
	}
	defer connDiaboromon.Close()
	log.Println("Conexión establecida con Diaboromon")

	connPrimary, err := grpc.Dial("primarynode:5007", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error de conexión con Primary Node: %v", err)
	}
	defer connPrimary.Close()
	log.Println("Conexión establecida con Primary Node")

	listener, err := net.Listen("tcp", ":5008")
	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterNodoTaiServer(grpcServer, nodoTai)

	log.Println("Nodo Tai escuchando en el puerto 5008...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Error al iniciar el servidor gRPC: %v", err)
	}
}

func stringToIntSlice(input string) []int {
	parts := []int{}
	for _, part := range strings.Split(input, ",") {
		num, _ := strconv.Atoi(part)
		parts = append(parts, num)
	}
	return parts
}
