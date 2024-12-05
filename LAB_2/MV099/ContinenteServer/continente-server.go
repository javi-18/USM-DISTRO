package main

import (
	"bufio"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	pb "ContinenteServer/proto"

	"google.golang.org/grpc"
)

type Digimon struct {
	Nombre   string
	Atributo string
	Estado   string
}

var claveAES = []byte("QeSU7H3h2XdYhm4Cg4drZouXFV9+pW+mLm8UFTkVKeo=")

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

func enviarElementos(client pb.ComunicacionClient, elementos []Digimon, TE int) {
	for i := 0; i < 6 && i < len(elementos); i++ {
		enviarElemento(client, elementos[i])
	}

	for i := 6; i < len(elementos); i++ {
		enviarElemento(client, elementos[i])
		time.Sleep(time.Duration(TE) * time.Second)
	}
}

func enviarElemento(client pb.ComunicacionClient, elemento Digimon) {
	mensaje := fmt.Sprintf("%s|%s|%s", elemento.Nombre, elemento.Atributo, elemento.Estado)

	mensajeCifrado, err := encriptado(mensaje)
	if err != nil {
		log.Fatalf("Error al cifrar el mensaje: %v", err)
	}

	req := &pb.Solicitud{
		Mensaje: mensajeCifrado,
	}

	_, err = client.EnviarEncriptado(context.Background(), req)
	if err != nil {
		log.Fatalf("Error al enviar el elemento cifrado: %v", err)
	}
	fmt.Printf("[CONTINENTE-SERVER] Estado enviado: Nombre %s, Estado: %s", elemento.Nombre, elemento.Estado)
}

func main() {
	//LEER ARCHIVO INPUT Y GUARDAR VALORES DE INICIALIZACION--------------------------
	archivo, err := os.Open("/app/INPUT.txt")
	if err != nil {
		log.Fatalf("Error al abrir el archivo INPUT.txt: %v", err)
	}
	defer archivo.Close()

	scanner := bufio.NewScanner(archivo)
	scanner.Scan()
	linea := scanner.Text()
	valores := strings.Split(linea, ",")
	PS, _ := strconv.ParseFloat(valores[0], 64)
	TE, _ := strconv.Atoi(valores[1])

	//LEER ARCHIVO DIGIMONS Y GUARDAR VALORES DE INICIALIZACION-----------------------

	archivo2, err := os.Open("/app/DIGIMONS.txt")
	if err != nil {
		log.Fatalf("Error al abrir el archivo DIGIMONS.txt: %v", err)
	}
	defer archivo2.Close()

	var digimons []Digimon
	scanner2 := bufio.NewScanner(archivo2)
	for scanner2.Scan() {
		linea := scanner.Text()
		valores := strings.Split(linea, ",")
		nombre := strings.TrimSpace(valores[0])
		atributo := strings.TrimSpace(valores[1])

		estado := "Sacrificado"
		if rand.Float64() > PS {
			estado = "No-Sacrificado"
		}

		digimons = append(digimons, Digimon{
			Nombre:   nombre,
			Atributo: atributo,
			Estado:   estado,
		})
	}

	//CONEXION CON PRIMARYNODE Y ENVIO ENCRIPTADO DE DATOS----------------------------
	enviar := digimons[(len(digimons) * 2 / 3):]

	conn, err := grpc.Dial("primary-node:5007", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error de conexión con PrimaryNode: %v", err)
	}
	defer conn.Close()

	client := pb.NewComunicacionClient(conn)
	enviarElementos(client, enviar, TE)
}
