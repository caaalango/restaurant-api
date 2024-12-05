package getdishcmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/calango-productions/api/internal/core/ports"
)

func New(
	dishRepository ports.DishRepository,
	redisRepository ports.RedisRepository,
) *Command {
	return &Command{
		DishRepository:  dishRepository,
		RedisRepository: redisRepository,
	}
}

func (c *Command) Execute(params Params) (*Result, error) {
	log.Printf("get dish initiated: %v", params)

	entities, err := c.getRelatedEntities(params)
	if err != nil {
		log.Printf("error to get related entities: %v", params)
		return &Result{Success: false, Status: http.StatusInternalServerError}, fmt.Errorf("error to get related entities: %v", err)
	}

	if entities.Dish.ClientToken.String() != params.Middle.ClientToken {
		log.Printf("this dish not belong to client's user: %v", params)
		return &Result{Success: false, Status: http.StatusForbidden}, fmt.Errorf("this dish not belong to client's user: %v", err)
	}

	if entities.Dish == nil {
		log.Printf("dish not found: %v", params)
		return &Result{Success: false, Status: http.StatusBadRequest}, fmt.Errorf("dish not found")
	}

	log.Printf("get user finished: %v", params)

	return &Result{Success: true, Status: http.StatusOK, Data: *entities.Dish}, nil
}

func (c *Command) getRelatedEntities(params Params) (*RelatedEntities, error) {
	dish, err := c.DishRepository.Get(ports.GetConf{Token: params.Core.DishToken, OnlyActives: true})
	if err != nil {
		return nil, err
	}

	return &RelatedEntities{Dish: dish}, nil
}
