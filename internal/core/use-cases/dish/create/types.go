package createdishcmd

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
	ActorToken   string
	Title        string
	Description  string
	Restrictions []string
	Price        int
}

type RelatedEntities struct {
	Dish *entities.Dish
}

type Result struct {
	Success bool
	Status  int
	Data    ResultData
}

type ResultData struct {
	DishToken string
}
