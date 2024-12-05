package updatedishcmd

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/calango-productions/api/internal/core/ports"
)

func New(dishRepository ports.DishRepository) *Command {
	return &Command{DishRepository: dishRepository}
}

func (c *Command) Execute(params Params) (*Result, error) {
	log.Printf("update dish initiated: %v", params)

	entities, err := c.getRelatedEntities(params)
	if err != nil {
		log.Printf("error to get related entities: %v", params)
		return &Result{Success: false, Status: http.StatusInternalServerError}, fmt.Errorf("error to get related entities: %v", err)
	}

	if entities.Actor.ClientToken != entities.Dish.ClientToken {
		log.Printf("user not belong to this client: %v", params)
		return &Result{Success: false, Status: http.StatusMethodNotAllowed}, fmt.Errorf("user not belong to this exists")
	}

	if entities.Dish == nil {
		log.Printf("dish already exists: %v", params)
		return &Result{Success: false, Status: http.StatusMethodNotAllowed}, fmt.Errorf("dish already exists")
	}

	buildedDish := c.buildUpdateMap(params)

	err = c.persistEntities(params, buildedDish)
	if err != nil {
		log.Printf("error to creation dish: %v", params)
		return &Result{Success: false, Status: http.StatusInternalServerError}, fmt.Errorf("failed to create: %v", err)
	}

	log.Printf("create traditional user finished: %v", params)

	return &Result{Success: true, Status: http.StatusOK}, nil
}

func (c *Command) getRelatedEntities(params Params) (*RelatedEntities, error) {
	dish, err := c.DishRepository.Get(ports.GetConf{Token: params.Core.DishToken})
	if err != nil {
		return nil, err
	}

	return &RelatedEntities{Dish: dish}, nil
}

func (c *Command) buildUpdateMap(params Params) map[string]interface{} {
	return map[string]interface{}{
		"Description": params.Core.Description,
		"Rescritions": strings.Join(params.Core.Restrictions, ";"),
		"Price":       params.Core.Price,
	}
}

func (c *Command) persistEntities(params Params, updates map[string]interface{}) error {
	err := c.DishRepository.Update(ports.UpdateConf{Token: params.Core.DishToken, Updates: updates})
	if err != nil {
		return fmt.Errorf("failed to create: %v", err)
	}

	return nil
}
