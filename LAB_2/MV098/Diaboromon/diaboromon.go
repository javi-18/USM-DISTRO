package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	pb "Diaboromon/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedDiaboromonServer
}

func (s *server) AtacarNodoTai(ctx context.Context, req *pb.BattleRequest) (*pb.BattleResponse, error) {
	log.Printf("Atacando Nodo Tai con %d puntos de daño y tipo de ataque %s", req.Damage, req.AttackType)

	response := &pb.BattleResponse{
		Result: "Ataque exitoso en Nodo Tai",
	}

	return response, nil
}

func LeerInputTXT(ruta string) (int, error) {
	file, err := os.Open(ruta)
	if err != nil {
		return 0, fmt.Errorf("Error al abrir input.txt: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		linea := scanner.Text()
		elementos := strings.Split(linea, ",")

		if len(elementos) < 3 {
			return 0, fmt.Errorf("Formato incorrecto en input.txt")
		}

		tiempo, err := strconv.Atoi(elementos[2])
		if err != nil {
			return 0, fmt.Errorf("Error al convertir el tercer elemento en tiempo: %v", err)
		}
		return tiempo, nil
	}
	return 0, fmt.Errorf("Archivo input.txt vacío")
}

func generarAtaque() *pb.BattleRequest {

	return &pb.BattleRequest{
		Damage:     10,
		AttackType: "Fuego",
	}
}

func main() {
	conn, err := grpc.Dial("nodotai:5008", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error de conexión con Nodo Tai: %v", err)
	}
	defer conn.Close()

	client := pb.NewNodoTaiClient(conn)

	listener, err := net.Listen("tcp", ":5004")
	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterDiaboromonServer(grpcServer, &server{})

	log.Println("Diaboromon escuchando en el puerto 5004")

	tiempoAtaque, err := LeerInputTXT("/app/INPUT.txt")
	if err != nil {
		log.Fatalf("Error al leer tiempo de ataque desde input.txt: %v", err)
	}
	log.Printf("Tiempo entre ataques leído desde input.txt: %d segundos", tiempoAtaque)

	go func() {
		for {
			ataque := generarAtaque()
			log.Println("Enviando ataque al Nodo Tai...")
			response, err := client.Atacar(context.Background(), ataque)
			if err != nil {
				log.Printf("Error al atacar a Nodo Tai: %v", err)
				return
			}
			log.Printf("Respuesta del Nodo Tai: %s", response.Result)

			time.Sleep(time.Duration(tiempoAtaque) * time.Second)
		}
	}()

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Error al iniciar el servidor gRPC: %v", err)
	}
}
