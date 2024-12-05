package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	proto "Clientes/proto"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("dist098:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se logro conectar: %v", err)
	}
	defer conn.Close()

	client := proto.NewClientesServiceClient(conn)

	archivo, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error al abrir el archivo: %v", err)
	}
	defer archivo.Close()

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		line := scanner.Text()
		orden := strings.Split(line, ",")

		pedido := &proto.Pedido{
			Id:          int32(toInt(orden[0])),
			Tipo:        orden[1],
			Nombre:      orden[2],
			Valor:       int32(toInt(orden[3])),
			Escolta:     orden[4],
			Destino:     orden[5],
			Seguimiento: int32(toInt(orden[6])),
		}

		_, err := client.EnviarPedido(context.Background(), pedido)
		if err != nil {
			log.Fatalf("No se logro mandar el mensaje: %v", err)
		}
		fmt.Printf("Mensaje enviado: %v\n", pedido)
		time.Sleep(5 * time.Second)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error al leer el input: %v", err)
	}
}

func toInt(str string) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalf("Error al transformar %v", err)
	}
	return val
}
