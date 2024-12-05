package listdishescmd

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
	ClientToken string
}

type CoreParams struct {
	Page int
	Size int
}

type RelatedEntities struct {
	Dishes []entities.Dish
}

type Result struct {
	Success bool
	Status  int
	Data    []entities.Dish
}
