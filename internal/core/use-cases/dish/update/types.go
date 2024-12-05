package updatedishcmd

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
	DishToken    string
	Description  string
	Restrictions []string
	Price        int
}

type RelatedEntities struct {
	Actor *entities.User
	Dish  *entities.Dish
}

type Result struct {
	Success bool
	Status  int
}
