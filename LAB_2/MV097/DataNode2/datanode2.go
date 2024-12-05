package main

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"

	pb "DataNode2/proto"

	"google.golang.org/grpc"
)

var claveAES = []byte("QeSU7H3h2XdYhm4Cg4drZouXFV9+pW+mLm8UFTkVKeo=")

type Data struct {
	ID       int
	Atributo string
}

type server struct {
	pb.UnimplementedComunicacionServer
	mu sync.Mutex
}

func desencriptado(cipherText string) (string, error) {
	cipherBytes, err := hex.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(claveAES)
	if err != nil {
		return "", err
	}

	if len(cipherBytes) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext muy corto")
	}

	iv := cipherBytes[:aes.BlockSize]
	cipherBytes = cipherBytes[aes.BlockSize:]

	modoDescifrado := cipher.NewCFBDecrypter(block, iv)
	modoDescifrado.XORKeyStream(cipherBytes, cipherBytes)

	return string(cipherBytes), nil
}

func procesado(data string) ([]Data, error) {
	lineas := strings.Split(data, "\n")

	var informacion []Data
	for _, linea := range lineas {
		valores := strings.Split(linea, "|")
		if len(valores) < 2 {
			continue
		}

		id, err := strconv.Atoi(valores[0])
		if err != nil {
			return nil, fmt.Errorf("error al convertir ID: %v", err)
		}

		unico := Data{
			ID:       id,
			Atributo: valores[1],
		}
		informacion = append(informacion, unico)
	}
	return informacion, nil
}

func (s *server) escritura(informacion []Data) error {
	archivo, err := os.OpenFile("/app/INFO_2.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("error al abrir el archivo: %v", err)
	}
	defer archivo.Close()

	for _, data := range informacion {
		linea := fmt.Sprintf("%d,%s\n", data.ID, data.Atributo)
		if _, err := archivo.WriteString(linea); err != nil {
			return fmt.Errorf("error al escribir en el archivo: %v", err)
		}
		fmt.Printf("[DATANODE1] Digimon recibido y archivado ID: %d, Atributo: %s\n", data.ID, data.Atributo)
	}
	return nil
}

func (s *server) Mensaje(ctx context.Context, msg *pb.Solicitud) (*pb.Respuesta, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	decryptedData, err := desencriptado(msg.Mensaje)
	if err != nil {
		return nil, fmt.Errorf("error desencriptando datos: %v", err)
	}
	informacion, err := procesado(decryptedData)
	if err != nil {
		return nil, err
	}

	err = s.escritura(informacion)
	if err != nil {
		return nil, err
	}

	log.Println("Datos escritos en INFO_2.txt correctamente")
	return &pb.Respuesta{Resultado: "Datos recibidos y escritos en INFO_2.txt"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":5006")
	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterComunicacionServer(s, &server{})

	log.Println("Servidor escuchando en el puerto :5006")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error en el servidor: %v", err)
	}
}
