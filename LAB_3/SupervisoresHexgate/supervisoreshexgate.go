package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	pb "SupervisoresHexgate/proto"

	"google.golang.org/grpc"
)

type VectorClockEntry struct {
	Key    string
	Values []int32
}

type SupervisorCache struct {
	Region      string
	Product     string
	VectorClock []VectorClockEntry
	ServerAddr  string
}

type Supervisor struct {
	brokerAddr string
	cache      map[string]SupervisorCache
	mu         sync.Mutex
	logFile    *os.File
}

func NewSupervisor(brokerAddr string) *Supervisor {
	logFile, err := os.OpenFile("supervisor_log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error al abrir archivo de log: %v", err)
	}

	return &Supervisor{
		brokerAddr: brokerAddr,
		cache:      make(map[string]SupervisorCache),
		logFile:    logFile,
	}
}

func (s *Supervisor) Close() {
	s.logFile.Close()
}

func (s *Supervisor) Run() {
	defer s.Close()

	fmt.Println("Bienvenido a Supervisores Hexgate. Escribe tus comandos:")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "exit" {
			fmt.Println("Saliendo de Supervisores Hexgate...")
			break
		}
		s.processCommand(input)
	}
}

func (s *Supervisor) processCommand(input string) {
	parts := strings.Fields(input)
	if len(parts) < 3 {
		fmt.Println("Comando no válido. Formatos soportados:")
		fmt.Println("AgregarProducto <nombre_Region> <nombre_Producto> [valor]")
		fmt.Println("RenombrarProducto <nombre_Region> <nombre_Producto> <nuevo_nombre>")
		fmt.Println("ActualizarValor <nombre_Region> <nombre_Producto> <nuevo_valor>")
		fmt.Println("BorrarProducto <nombre_Region> <nombre_Producto>")
		return
	}

	command := parts[0]
	region := parts[1]
	product := parts[2]

	switch command {
	case "AgregarProducto":
		value := int32(0)
		if len(parts) == 4 {
			fmt.Sscanf(parts[3], "%d", &value)
		}
		s.sendCommandToBroker(&pb.Request{
			CommandType: "AgregarProducto", Region: region, Product: product, Value: value,
		})
	case "RenombrarProducto":
		if len(parts) != 4 {
			fmt.Println("Uso: RenombrarProducto <nombre_Region> <nombre_Producto> <nuevo_nombre>")
			return
		}
		s.sendCommandToBroker(&pb.Request{
			CommandType: "RenombrarProducto", Region: region, Product: product, NewName: parts[3],
		})
	case "ActualizarValor":
		if len(parts) != 4 {
			fmt.Println("Uso: ActualizarValor <nombre_Region> <nombre_Producto> <nuevo_valor>")
			return
		}
		var newValue int32
		fmt.Sscanf(parts[3], "%d", &newValue)
		s.sendCommandToBroker(&pb.Request{
			CommandType: "ActualizarValor", Region: region, Product: product, Value: newValue,
		})
	case "BorrarProducto":
		s.sendCommandToBroker(&pb.Request{
			CommandType: "BorrarProducto", Region: region, Product: product,
		})
	default:
		fmt.Println("Comando no reconocido.")
	}
}

func (s *Supervisor) sendCommandToBroker(req *pb.Request) {
	conn, err := grpc.Dial(s.brokerAddr, grpc.WithInsecure())
	if err != nil {
		log.Printf("Error al conectar con el Broker: %v\n", err)
		return
	}
	defer conn.Close()

	client := pb.NewBrokerServiceClient(conn)
	req.VectorClock = s.convertToProtoVectorClock(s.cache[fmt.Sprintf("%s:%s", req.Region, req.Product)].VectorClock)

	res, err := client.RedirectRequest(context.Background(), req)
	if err != nil {
		log.Printf("Error al enviar comando al Broker: %v\n", err)
		return
	}

	s.handleResponse(req.Region, req.Product, res)
}

func (s *Supervisor) handleResponse(region, product string, response *pb.Response) {
	s.mu.Lock()
	defer s.mu.Unlock()

	cacheKey := fmt.Sprintf("%s:%s", region, product)
	vectorClock := s.convertFromProtoVectorClock(response.VectorClock)

	s.cache[cacheKey] = SupervisorCache{
		Region:      region,
		Product:     product,
		VectorClock: vectorClock,
		ServerAddr:  response.ServerAddress,
	}

	logEntry := fmt.Sprintf("[%s] Región: %s | Producto: %s | Reloj Vectorial: %v | Servidor: %s\n",
		time.Now().Format("2006-01-02 15:04:05"), region, product, vectorClock, response.ServerAddress)
	s.logFile.WriteString(logEntry)

	fmt.Printf("Comando procesado. Reloj vectorial: %v\n", vectorClock)
}

func (s *Supervisor) convertToProtoVectorClock(clock []VectorClockEntry) []*pb.VectorClockEntry {
	protoClock := make([]*pb.VectorClockEntry, len(clock))
	for i, entry := range clock {
		protoClock[i] = &pb.VectorClockEntry{
			Key:    entry.Key,
			Values: entry.Values,
		}
	}
	return protoClock
}

func (s *Supervisor) convertFromProtoVectorClock(protoClock []*pb.VectorClockEntry) []VectorClockEntry {
	clock := make([]VectorClockEntry, len(protoClock))
	for i, entry := range protoClock {
		clock[i] = VectorClockEntry{
			Key:    entry.Key,
			Values: entry.Values,
		}
	}
	return clock
}

func main() {
	supervisor := NewSupervisor("dist097:5001")
	supervisor.Run()
}
