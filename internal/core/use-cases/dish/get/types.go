package getdishcmd

import (
	"github.com/calango-productions/api/internal/core/entities"
	"github.com/calango-productions/api/internal/core/ports"
)

type Command struct {
	DishRepository  ports.DishRepository
	RedisRepository ports.RedisRepository
}

type Params struct {
	Middle MiddleParams
	Core   CoreParams
}

type MiddleParams struct {
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
	Data    entities.Dish
}
