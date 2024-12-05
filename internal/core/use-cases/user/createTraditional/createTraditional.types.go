package createtradusercmd

import (
	"github.com/calango-productions/api/internal/core/entities"
	"github.com/calango-productions/api/internal/core/ports"
)

type Command struct {
	userRepository ports.UserRepository
}

type Params struct {
	ActorToken string
	Email      string
	Password   string
}

type RelatedEntities struct {
	User *entities.User
}

type Result struct {
	Success bool
	Status  int
	Data    EntityIdentifier
}

type EntityIdentifier struct {
	UserToken string
}
