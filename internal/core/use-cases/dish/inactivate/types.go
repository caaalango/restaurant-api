package inactivatedishcmd

import (
	"github.com/calango-productions/api/internal/core/entities"
	"github.com/calango-productions/api/internal/core/ports"
)

type Command struct {
	DishRepository ports.DishRepository
}

type Params struct {
	Middle MiddleParams
	Core   CoreParams
}

type MiddleParams struct {
	ActorToken  string
	ClientToken string
}

type CoreParams struct {
	DishToken string
}

type RelatedEntities struct {
	Dish *entities.Dish
}

type Result struct {
	Success bool
	Status  int
}
