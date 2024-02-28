package main

import (
	"fmt"
	"time"

	"github.com/Ferriem/rabbitmq/code/RabbitMQ"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQPubSub("pub-sub")
	for i := 0; i < 100; i++ {
		rabbitmq.PublishPub("pub-sub" + fmt.Sprintf("%d", i))
		fmt.Printf("Publish: %d\n", i)
		time.Sleep(1 * time.Second)
	}

}
