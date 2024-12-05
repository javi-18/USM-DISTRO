package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"

	pb "ServidoresHextech/proto"

	"google.golang.org/grpc"
)

type HextechServer struct {
	pb.UnimplementedHextechServiceServer
	regionFiles     map[string]string
	vectorClock     []*pb.VectorClockEntry
	logs            []string
	mutex           sync.Mutex
	serverID        int
	propagationFreq time.Duration
	peers           []string
}

func NewHextechServer(id int, propagationFreq time.Duration, peers []string) *HextechServer {
	return &HextechServer{
		regionFiles:     make(map[string]string),
		vectorClock:     []*pb.VectorClockEntry{},
		logs:            []string{},
		serverID:        id,
		propagationFreq: propagationFreq,
		peers:           peers,
	}
}

func updateVectorClock(vectorClock *[]*pb.VectorClockEntry, key string, values []int32) {
	for _, entry := range *vectorClock {
		if entry.Key == key {
			entry.Values = values
			return
		}
	}
	*vectorClock = append(*vectorClock, &pb.VectorClockEntry{
		Key:    key,
		Values: values,
	})
}

func getValuesFromVectorClock(vectorClock []*pb.VectorClockEntry, key string) ([]int32, bool) {
	for _, entry := range vectorClock {
		if entry.Key == key {
			return entry.Values, true
		}
	}
	return nil, false
}

func (s *HextechServer) AddProduct(ctx context.Context, req *pb.ProductRequest) (*pb.Response, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	region := req.Region
	product := req.Product
	quantity := req.Quantity

	fileName := fmt.Sprintf("data/%s.txt", region)
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		os.MkdirAll("data", os.ModePerm)
		os.Create(fileName)
	}

	file, err := os.OpenFile(fileName, os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	var updated bool
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, fmt.Sprintf("%s ", product)) {
			line = fmt.Sprintf("%s %d", product, quantity)
			updated = true
		}
		lines = append(lines, line)
	}
	if !updated {
		lines = append(lines, fmt.Sprintf("%s %d", product, quantity))
	}

	file.Truncate(0)
	file.Seek(0, 0)
	for _, line := range lines {
		file.WriteString(line + "\n")
	}

	var currentValues []int32
	if values, found := getValuesFromVectorClock(s.vectorClock, region); found {
		currentValues = values
	} else {
		currentValues = make([]int32, len(s.peers)+1)
	}
	currentValues[s.serverID]++
	updateVectorClock(&s.vectorClock, region, currentValues)

	logEntry := fmt.Sprintf("Agregar %s %s %d", region, product, quantity)
	s.logs = append(s.logs, logEntry)

	return &pb.Response{
		Message: "Producto agregado correctamente",
		Success: true,
	}, nil
}

func (s *HextechServer) RenameProduct(ctx context.Context, req *pb.RenameRequest) (*pb.Response, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	region := req.Region
	oldName := req.OldName
	newName := req.NewName

	fileName := fmt.Sprintf("data/%s.txt", region)
	file, err := os.OpenFile(fileName, os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, fmt.Sprintf("%s ", oldName)) {
			line = strings.Replace(line, oldName, newName, 1)
		}
		lines = append(lines, line)
	}

	file.Truncate(0)
	file.Seek(0, 0)
	for _, line := range lines {
		file.WriteString(line + "\n")
	}

	var currentValues []int32
	if values, found := getValuesFromVectorClock(s.vectorClock, region); found {
		currentValues = values
	} else {
		currentValues = make([]int32, len(s.peers)+1)
	}
	currentValues[s.serverID]++
	updateVectorClock(&s.vectorClock, region, currentValues)

	logEntry := fmt.Sprintf("Renombrar %s %s %s", region, oldName, newName)
	s.logs = append(s.logs, logEntry)

	return &pb.Response{
		Message: "Producto renombrado correctamente",
		Success: true,
	}, nil
}

func (s *HextechServer) DeleteProduct(ctx context.Context, req *pb.ProductRequest) (*pb.Response, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	region := req.Region
	product := req.Product

	fileName := fmt.Sprintf("data/%s.txt", region)
	file, err := os.OpenFile(fileName, os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	deleted := false
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, fmt.Sprintf("%s ", product)) {
			deleted = true
			continue
		}
		lines = append(lines, line)
	}

	if !deleted {
		return &pb.Response{
			Message: "Producto no encontrado",
			Success: false,
		}, nil
	}

	file.Truncate(0)
	file.Seek(0, 0)
	for _, line := range lines {
		file.WriteString(line + "\n")
	}

	var currentValues []int32
	if values, found := getValuesFromVectorClock(s.vectorClock, region); found {
		currentValues = values
	} else {
		currentValues = make([]int32, len(s.peers)+1)
	}
	currentValues[s.serverID]++
	updateVectorClock(&s.vectorClock, region, currentValues)

	logEntry := fmt.Sprintf("Eliminar %s %s", region, product)
	s.logs = append(s.logs, logEntry)

	return &pb.Response{
		Message: "Producto eliminado correctamente",
		Success: true,
	}, nil
}

func (s *HextechServer) PropagateChanges(ctx context.Context, req *pb.PropagateRequest) (*pb.Empty, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, entry := range req.VectorClock {
		currentValues, found := getValuesFromVectorClock(s.vectorClock, entry.Key)
		if found {
			for i := range currentValues {
				if i < len(entry.Values) && entry.Values[i] > currentValues[i] {
					currentValues[i] = entry.Values[i]
				}
			}
		} else {
			s.vectorClock = append(s.vectorClock, entry)
		}
	}

	s.logs = append(s.logs, req.Logs...)

	log.Println("Cambios propagados recibidos y aplicados.")
	return &pb.Empty{}, nil
}

func (s *HextechServer) SyncState(ctx context.Context, req *pb.SyncRequest) (*pb.SyncResponse, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	log.Printf("Sincronización solicitada desde nodo %s", req.SourceNode)

	return &pb.SyncResponse{
		UpdatedClock: s.vectorClock,
		UpdatedLogs:  s.logs,
	}, nil
}

func (s *HextechServer) propagateChanges() {
	for {
		time.Sleep(s.propagationFreq)
		s.mutex.Lock()
		defer s.mutex.Unlock()

		for _, peer := range s.peers {
			conn, err := grpc.Dial(peer, grpc.WithInsecure())
			if err != nil {
				log.Printf("Error al conectar con el servidor %s: %v", peer, err)
				continue
			}
			defer conn.Close()

			client := pb.NewHextechServiceClient(conn)
			req := &pb.PropagateRequest{
				VectorClock: s.vectorClock,
				Logs:        s.logs,
			}
			_, err = client.PropagateChanges(context.Background(), req)
			if err != nil {
				log.Printf("Error al propagar cambios a %s: %v", peer, err)
			} else {
				log.Printf("Cambios propagados exitosamente a %s", peer)
			}
		}
		s.logs = []string{}
	}
}

func main() {
	ports := []int{5002, 5003, 5004}
	peersList := [][]string{
		{"localhost:5003", "localhost:5004"},
		{"localhost:5002", "localhost:5004"},
		{"localhost:5002", "localhost:5003"},
	}

	for i, port := range ports {
		server := NewHextechServer(i, 30*time.Second, peersList[i])

		listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err != nil {
			log.Fatalf("Error al iniciar el servidor en el puerto %d: %v", port, err)
		}
		log.Printf("Servidor iniciado en el puerto %d", port)

		grpcServer := grpc.NewServer()
		pb.RegisterHextechServiceServer(grpcServer, server)

		go func() {
			server.propagateChanges()
		}()

		go func() {
			if err := grpcServer.Serve(listener); err != nil {
				log.Fatalf("Error al iniciar el servidor gRPC: %v", err)
			}
		}()
	}

	select {}
}