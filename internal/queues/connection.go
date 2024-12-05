package queues

import (
	"log"

	"github.com/calango-productions/api/internal/queues/consumers"
	"github.com/calango-productions/api/internal/queues/dispatchers"
	"github.com/streadway/amqp"
)

func ConnectRabbitMQ(rabbitMQConn *amqp.Connection) *dispatchers.Dispatchers {
	ch, err := rabbitMQConn.Channel()
	if err != nil {
		panic("failed to open a channel")
	}
	defer func() {
		if err := ch.Close(); err != nil {
			panic("Failed to close channel")
		}
	}()

	for _, config := range QUEUES_CONFIG {
		args := amqp.Table{}
		for k, v := range config.Args {
			args[k] = v
		}

		_, err := ch.QueueDeclare(
			config.Name,
			config.Durable,
			config.AutoDelete,
			config.Exclusive,
			config.NoWait,
			args,
		)
		if err != nil {
			panic("Failed to declare queue")
		}

		log.Printf("Queue declared: %s", config.Name)
	}

	consumers.New(ch)
	dispatchers := dispatchers.New(ch)

	return dispatchers
}
