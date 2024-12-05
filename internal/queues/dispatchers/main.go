package dispatchers

import (
	"github.com/streadway/amqp"
)

type Dispatchers struct {
	channel *amqp.Channel
}

func New(ch *amqp.Channel) *Dispatchers {
	return &Dispatchers{channel: ch}
}
