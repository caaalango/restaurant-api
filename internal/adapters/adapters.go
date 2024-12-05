package adapters

import (
	"github.com/calango-productions/api/internal/adapters/connections"
	"github.com/calango-productions/api/internal/middlewares"
	"github.com/calango-productions/api/internal/queues"
	"github.com/calango-productions/api/internal/queues/consumers"
	"github.com/calango-productions/api/internal/queues/dispatchers"
	"github.com/calango-productions/api/internal/repositories"
)

type Adapters struct {
	Middlewares  *middlewares.Provider
	Repositories *repositories.Provider
	Dispatchers  *dispatchers.Dispatchers
	Consumers    *consumers.Consumers
}

func New(conn *connections.Connections) *Adapters {
	apt := &Adapters{}

	apt.Repositories = repositories.New(conn)

	apt.Middlewares = middlewares.New()

	apt.Dispatchers = queues.ConnectRabbitMQ(conn.RabbitMQ)

	return apt
}
