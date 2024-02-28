package main

import (
	"fmt"
	"time"

	"github.com/Ferriem/rabbitmq/code/RabbitMQ"
)

func main() {
	rabbitmq1 := RabbitMQ.NewRabbitMQRouting("routing", "routing_key1")
	rabbitmq2 := RabbitMQ.NewRabbitMQRouting("routing", "routing_key2")
	for i := 0; i < 10; i++ {
		rabbitmq1.PublishRouting("Routing1" + fmt.Sprintf("%d", i))
		rabbitmq2.PublishRouting("Routing2" + fmt.Sprintf("%d", i))
		fmt.Printf("Publish: %d\n", i)
		time.Sleep(1 * time.Second)
	}
}
