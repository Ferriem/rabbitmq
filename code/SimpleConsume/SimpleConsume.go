package main

import "github.com/Ferriem/rabbitmq/code/RabbitMQ"

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("Simple")
	rabbitmq.ConsumeSimple()

}
