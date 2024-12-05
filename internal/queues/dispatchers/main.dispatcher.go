package dispatchers

import (
	"fmt"

	"github.com/streadway/amqp"
)

func (d *Dispatchers) MainDispatcher(message string) error {
	exchange := ""
	routingKey := "main"
	mandatory := false
	immediate := false

	publishing := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
	}

	err := d.channel.Publish(
		exchange,
		routingKey,
		mandatory,
		immediate,
		publishing,
	)
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	return nil
}
