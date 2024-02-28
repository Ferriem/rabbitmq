package main

import "github.com/Ferriem/rabbitmq/code/RabbitMQ"

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQRouting("routing ", "routing_key2")
	rabbitmq.ReceiveRouting()
}
