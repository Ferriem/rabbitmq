package main

import (
	"fmt"
	"time"

	"github.com/Ferriem/rabbitmq/code/RabbitMQ"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQTopic("Topic", "topic.one")
	for i := 0; i <= 10; i++ {
		rabbitmq.PublishTopic("I am topic one!" + fmt.Sprintf("%d", i))
		fmt.Printf("Publish: %d\n", i)
		time.Sleep(1 * time.Second)
	}
}
