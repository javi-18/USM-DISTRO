package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

const (
	rabbitMQURL = "amqp://guest:guest@dist098:5672/"
	queueName   = "raquis_queue"
)

func main() {
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		log.Fatalf("error al conectar a RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("error al abrir el canal: %v", err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("error al registrar el consumidor: %v", err)
	}

	for msg := range msgs {
		fmt.Printf("Mensaje recibido de RabbitMQ: %s\n", msg.Body)
	}
}
