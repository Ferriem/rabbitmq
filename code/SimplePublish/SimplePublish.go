package main

import (
	"fmt"

	"github.com/Ferriem/rabbitmq/code/RabbitMQ"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("Simple")
	// rabbitmq.PublishSimple("I am Simple!")
	for i := 0; i < 100; i++ {
		rabbitmq.PublishSimple("I am Simple! + " + fmt.Sprintf("%d", i))
	}
	fmt.Printf("Success")
}
