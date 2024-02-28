package main

import "github.com/Ferriem/rabbitmq/code/RabbitMQ"

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQRouting("routing ", "routing_key1")
	rabbitmq.ReceiveRouting()
}
