package main

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"

	pb "PrimaryNode/proto"

	"google.golang.org/grpc"
)

var claveAES = []byte("QeSU7H3h2XdYhm4Cg4drZouXFV9+pW+mLm8UFTkVKeo=")

type Digimon struct {
	ID          int
	Nombre      string
	Atributo    string
	Estado      string
	NumDataNode int
}

type servidor struct {
	pb.UnimplementedComunicacionServer
	mensajes []Digimon
	mu       sync.Mutex
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

func encriptado(texto string) (string, error) {
	block, err := aes.NewCipher(claveAES)
	if err != nil {
		return "", err
	}

	textoBytes := []byte(texto)
	cipherText := make([]byte, aes.BlockSize+len(textoBytes))

	iv := cipherText[:aes.BlockSize]
	if _, err := rand.Read(iv); err != nil {
		return "", err
	}

	modoCifrado := cipher.NewCFBEncrypter(block, iv)
	modoCifrado.XORKeyStream(cipherText[aes.BlockSize:], textoBytes)

	return hex.EncodeToString(cipherText), nil
}

func splitMensaje(mensaje string) []string {
	return strings.Split(mensaje, "|")
}

func asignarNumDataNode(nombre string) int {
	if nombre[0] >= 'A' && nombre[0] <= 'M' {
		return 1
	}
	return 2
}

func (s *servidor) EnviarEncriptado(ctx context.Context, req *pb.Solicitud) (*pb.Respuesta, error) {
	mensajeDesencriptado, err := desencriptado(req.Mensaje)
	if err != nil {
		return nil, fmt.Errorf("error al desencriptar el mensaje: %v", err)
	}
	s.mu.Lock()
	defer s.mu.Unlock()

	atributos := splitMensaje(mensajeDesencriptado)
	digimon := Digimon{
		ID:          len(s.mensajes) + 1,
		Nombre:      atributos[0],
		Atributo:    atributos[1],
		Estado:      atributos[2],
		NumDataNode: asignarNumDataNode(atributos[0]),
	}
	s.mensajes = append(s.mensajes, digimon)

	archivo, err := os.OpenFile("/app/INFO.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, fmt.Errorf("error al abrir el archivo: %v", err)
	}
	defer archivo.Close()

	linea := fmt.Sprintf("%d,%d,%s,%s\n", digimon.ID, digimon.NumDataNode, digimon.Nombre, digimon.Estado)
	if _, err := archivo.WriteString(linea); err != nil {
		return nil, fmt.Errorf("error al escribir en el archivo: %v", err)
	}
	fmt.Printf("[PRIMARYNODE] Digimon recibido y archivado: %d,%d,%s,%s\n", digimon.ID, digimon.NumDataNode, digimon.Nombre, digimon.Estado)

	mensajeParaDataNode := fmt.Sprintf("%s|%s", digimon.ID, digimon.Atributo)
	mensajeEncriptado, err := encriptado(mensajeParaDataNode)
	if err != nil {
		return nil, fmt.Errorf("error al encriptar el mensaje para el DataNode: %v", err)
	}

	if err := enviarMensajeDataNode(digimon.NumDataNode, mensajeEncriptado); err != nil {
		return nil, fmt.Errorf("error al enviar el mensaje al DataNode: %v", err)
	}

	fmt.Printf("[PRIMARYNODE] Mensaje recibido: %s\n", mensajeDesencriptado)
	return &pb.Respuesta{Resultado: "Mensaje recibido correctamente"}, nil
}

func enviarMensajeDataNode(numDataNode int, mensaje string) error {
	if numDataNode == 1 {
		conn, err := grpc.Dial("datanode1:5002", grpc.WithInsecure())
		if err != nil {
			return fmt.Errorf("error al conectar con el DataNode: %v", err)
		}
		defer conn.Close()

		client := pb.NewComunicacionClient(conn)

		_, err = client.EnviarEncriptado(context.Background(), &pb.Solicitud{Mensaje: mensaje})
		if err != nil {
			return fmt.Errorf("error al enviar el mensaje al DataNode: %v", err)
		}

		fmt.Printf("[PRIMARYNODE] Mensaje enviado al DataNode %d: %s\n", numDataNode, mensaje)
		return nil
	}
	conn, err := grpc.Dial("datanode2:5006", grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("error al conectar con el DataNode: %v", err)
	}
	defer conn.Close()

	client := pb.NewComunicacionClient(conn)

	_, err = client.EnviarEncriptado(context.Background(), &pb.Solicitud{Mensaje: mensaje})
	if err != nil {
		return fmt.Errorf("error al enviar el mensaje al DataNode: %v", err)
	}

	fmt.Printf("[PRIMARYNODE] Mensaje enviado al DataNode %d: %s\n", numDataNode, mensaje)
	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":5007")
	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}

	srv := grpc.NewServer()
	pb.RegisterComunicacionServer(srv, &servidor{})

	fmt.Println("Servidor gRPC escuchando en el puerto :5007")
	if err := srv.Serve(listener); err != nil {
		log.Fatalf("Error en el servidor: %v", err)

	}
}
