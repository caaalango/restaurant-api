package inactivatedishcmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/calango-productions/api/internal/core/entities"
	"github.com/calango-productions/api/internal/core/ports"
)

func New(dishRepository ports.DishRepository) *Command {
	return &Command{DishRepository: dishRepository}
}

func (c *Command) Execute(params Params) (*Result, error) {
	log.Printf("inactivate dish initiated: %v", params)

	entities, err := c.getRelatedEntities(params)
	if err != nil {
		log.Printf("error to get related entities: %v", params)
		return &Result{Success: false, Status: http.StatusInternalServerError}, fmt.Errorf("error to get related entities: %v", err)
	}

	if entities.Dish == nil {
		log.Printf("dish not found: %v", params)
		return &Result{Success: false, Status: http.StatusBadRequest}, fmt.Errorf("dish not found")
	}

	if params.Middle.ClientToken != entities.Dish.Token.String() {
		log.Printf("dish not belong to client's user: %v", params)
		return &Result{Success: false, Status: http.StatusForbidden}, fmt.Errorf("dish not belong to client's user")
	}

	err = c.persistEntities(entities.Dish)
	if err != nil {
		log.Printf("error to inactivation dish: %v", params)
		return &Result{Success: false, Status: http.StatusInternalServerError}, fmt.Errorf("failed to create: %v", err)
	}

	log.Printf("inactivate dish finished: %v", params)

	return &Result{Success: true, Status: http.StatusOK}, nil
}

func (c *Command) getRelatedEntities(params Params) (*RelatedEntities, error) {
	dish, err := c.DishRepository.Get(ports.GetConf{Token: params.Core.DishToken, OnlyActives: true})
	if err != nil {
		return nil, err
	}

	return &RelatedEntities{Dish: dish}, nil
}

func (c *Command) persistEntities(dish *entities.Dish) error {
	err := c.DishRepository.Inactivate(ports.InactivateConf{Token: dish.Token.String()})
	if err != nil {
		return fmt.Errorf("failed to inactivate: %v", err)
	}

	return nil
}
