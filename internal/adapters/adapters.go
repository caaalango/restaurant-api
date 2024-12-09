package adapters

import (
	"github.com/calango-productions/api/internal/adapters/connections"
	"github.com/calango-productions/api/internal/middlewares"
	"github.com/calango-productions/api/internal/repositories"
)

type Adapters struct {
	Middlewares  *middlewares.Provider
	Repositories *repositories.Provider
}

func New(conn *connections.Connections) *Adapters {
	apt := &Adapters{}

	apt.Repositories = repositories.New(conn)

	apt.Middlewares = middlewares.New()

	return apt
}
