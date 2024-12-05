package consumers

import (
	"github.com/streadway/amqp"
)

type Consumers struct {
	channel *amqp.Channel
}

func New(channel *amqp.Channel) {
	InitMainConsumer(channel)
}
