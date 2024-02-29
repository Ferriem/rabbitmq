package main

import (
	"os"

	"github.com/Ferriem/rabbitmq/code/RabbitMQ"
)

func main() {
	for _, s := range os.Args[1:] {
		rabbitmq := RabbitMQ.NewRabbitMQTopic("Topic", s)
		rabbitmq.ReceiveTopic()
	}
}
