package createdishcmd

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/calango-productions/api/internal/core/entities"
	"github.com/calango-productions/api/internal/core/ports"
)

func New(
	dishRepository ports.DishRepository,
) *Command {
	return &Command{
		DishRepository: dishRepository,
	}
}

func (c *Command) Execute(params Params) (*Result, error) {
	log.Printf("create dish initiated: %v", params)

	entities, err := c.getRelatedEntities(params)
	if err != nil {
		log.Printf("error to get related entities: %v", params)
		return &Result{Success: false, Status: http.StatusInternalServerError}, fmt.Errorf("error to get related entities: %v", err)
	}

	if entities.Dish.ClientToken.String() != params.Middle.ClientToken {
		log.Printf("this dish not belogs to client's user: %v", params)
		return &Result{Success: false, Status: http.StatusUnauthorized}, fmt.Errorf("this dish not belogs to client's user")
	}

	if entities.Dish != nil {
		log.Printf("dish already exists: %v", params)
		return &Result{Success: false, Status: http.StatusBadRequest}, fmt.Errorf("dish with same title already exists")
	}

	buildedDish := c.buildDish(params)

	createDish, err := c.persistEntities(buildedDish)
	if err != nil {
		log.Printf("error to creation dish: %v", params)
		return &Result{Success: false, Status: http.StatusInternalServerError}, fmt.Errorf("failed to create: %v", err)
	}

	log.Printf("create dish finished: %v", params)

	return &Result{Success: true, Status: http.StatusOK, Data: ResultData{DishToken: createDish.Token.String()}}, nil
}

func (c *Command) getRelatedEntities(params Params) (*RelatedEntities, error) {
	dish, err := c.DishRepository.GetByKey(ports.GetByKeyConf{Key: "title", Value: params.Core.Title})
	if err != nil {
		return nil, err
	}

	return &RelatedEntities{Dish: dish}, nil
}

func (c *Command) buildDish(params Params) *entities.Dish {
	return &entities.Dish{
		Title:        params.Core.Title,
		Description:  params.Core.Description,
		Restrictions: strings.Join(params.Core.Restrictions, ";"),
		Price:        params.Core.Price,
	}
}

func (c *Command) buildCredential(encryptedPwd string) *entities.Credential {
	return &entities.Credential{
		Password: encryptedPwd,
		Type:     entities.CredentialType(entities.PASSWORD),
	}
}

func (c *Command) persistEntities(dish *entities.Dish) (*entities.Dish, error) {
	dish, err := c.DishRepository.Create(ports.CreateConf[entities.Dish]{Item: *dish})
	if err != nil {
		return nil, fmt.Errorf("failed to create: %v", err)
	}

	return dish, nil
}
