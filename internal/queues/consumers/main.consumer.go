package consumers

import (
	"fmt"

	"github.com/streadway/amqp"
)

func InitMainConsumer(channel *amqp.Channel) error {
	queueName := "main"
	consumerName := "main"
	autoAck := true
	exclusive := false
	noLocal := false
	noWait := false
	args := amqp.Table{}

	_, err := channel.Consume(
		queueName,
		consumerName,
		autoAck,
		exclusive,
		noLocal,
		noWait,
		args,
	)
	if err != nil {
		return fmt.Errorf("failed to register a consumer: %w", err)
	}

	done := make(chan bool)

	go func() {
		defer close(done)

		fmt.Println("Hello, World")

		done <- true
	}()

	select {}
}
