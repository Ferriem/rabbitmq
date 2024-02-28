package main

import "github.com/Ferriem/rabbitmq/code/RabbitMQ"

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQPubSub("pub-sub")
	rabbitmq.ReceiveSub()
}
