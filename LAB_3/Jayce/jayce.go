package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	pb "Jayce/proto"

	"google.golang.org/grpc"
)

type Jayce struct {
	cache   map[string]JayceCache
	mu      sync.Mutex
	broker  string
	logFile string
}

type JayceCache struct {
	Region      string
	Product     string
	VectorClock []*pb.VectorClockEntry
	Quantity    int32
}

func NewJayce(brokerAddress, logFile string) *Jayce {
	return &Jayce{
		cache:   make(map[string]JayceCache),
		broker:  brokerAddress,
		logFile: logFile,
	}
}

func (j *Jayce) Run() {
	fmt.Println("Bienvenido a Jayce. Escribe tus comandos:")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "exit" {
			fmt.Println("Saliendo de Jayce...")
			break
		}
		j.processCommand(input)
	}
}

func (j *Jayce) processCommand(input string) {
	parts := strings.Fields(input)
	if len(parts) < 3 || parts[0] != "ObtenerProducto" {
		fmt.Println("Comando no válido. Uso: ObtenerProducto <nombre_Region> <nombre_Producto>")
		return
	}

	region := parts[1]
	product := parts[2]

	response, err := j.queryBroker(region, product)
	if err != nil {
		log.Printf("Error al consultar el Broker: %v\n", err)
		return
	}

	j.handleResponse(region, product, response)
}

func (j *Jayce) queryBroker(region, product string) (*pb.Response, error) {
	conn, err := grpc.Dial(j.broker, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("no se pudo conectar al Broker: %v", err)
	}
	defer conn.Close()

	client := pb.NewBrokerServiceClient(conn)
	req := &pb.Request{
		CommandType: "ObtenerProducto",
		Region:      region,
		Product:     product,
	}
	return client.RedirectRequest(context.Background(), req)
}

func (j *Jayce) handleResponse(region, product string, response *pb.Response) {
	j.mu.Lock()
	defer j.mu.Unlock()

	cacheKey := fmt.Sprintf("%s:%s", region, product)

	if previous, exists := j.cache[cacheKey]; exists {
		if isVectorClockOlder(response.VectorClock, previous.VectorClock) {
			fmt.Printf("Inconsistencia detectada para %s:%s. Notificando al Broker...\n", region, product)
			j.notifyInconsistency(region, product, response.VectorClock)
			return
		}
	}

	j.cache[cacheKey] = JayceCache{
		Region:      region,
		Product:     product,
		VectorClock: response.VectorClock,
		Quantity:    response.Quantity,
	}

	j.logResponse(region, product, response)

	fmt.Printf("Región: %s, Producto: %s, Cantidad: %d\n", region, product, response.Quantity)
}

func isVectorClockOlder(vc1, vc2 []*pb.VectorClockEntry) bool {
	vc1Map := make(map[string][]int32)
	for _, entry := range vc1 {
		vc1Map[entry.Key] = entry.Values
	}

	for _, entry := range vc2 {
		if vc1Values, exists := vc1Map[entry.Key]; exists {
			for i := 0; i < len(entry.Values); i++ {
				if vc1Values[i] < entry.Values[i] {
					return true
				}
			}
		}
	}
	return false
}

func (j *Jayce) notifyInconsistency(region, product string, vectorClock []*pb.VectorClockEntry) {
	conn, err := grpc.Dial(j.broker, grpc.WithInsecure())
	if err != nil {
		log.Printf("Error al conectar con el Broker: %v\n", err)
		return
	}
	defer conn.Close()

	client := pb.NewBrokerServiceClient(conn)
	req := &pb.InconsistencyRequest{
		Region:      region,
		Product:     product,
		VectorClock: vectorClock,
	}

	_, err = client.HandleInconsistency(context.Background(), req)
	if err != nil {
		log.Printf("Error al notificar inconsistencia al Broker: %v\n", err)
	}
}

func (j *Jayce) logResponse(region, product string, response *pb.Response) {
	logEntry := fmt.Sprintf("Region: %s | Producto: %s | Cantidad: %d | Reloj Vectorial: %v\n", region, product, response.Quantity, response.VectorClock)

	file, err := os.OpenFile(j.logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Error al escribir en el archivo de log: %v\n", err)
		return
	}
	defer file.Close()

	if _, err := file.WriteString(logEntry); err != nil {
		log.Printf("Error al escribir en el archivo de log: %v\n", err)
	}
}

func main() {
	brokerAddress := "dist097:5001"
	jayce := NewJayce(brokerAddress, "jayce_logs.txt")
	jayce.Run()
}
