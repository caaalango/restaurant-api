package listdishescmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/calango-productions/api/internal/core/ports"
)

const MAX_PAGINATION_SIZE = 20

func New(dishRepository ports.DishRepository) *Command {
	return &Command{DishRepository: dishRepository}
}

func (c *Command) Execute(params Params) (*Result, error) {
	log.Printf("list dishes initiated: %v", params)

	entities, err := c.getRelatedEntities(params)
	if err != nil {
		log.Printf("error to get related entities: %v", params)
		return &Result{Success: false, Status: http.StatusInternalServerError}, fmt.Errorf("error to get related entities: %v", err)
	}

	if len(entities.Dishes) == 0 {
		log.Printf("list dished finished: %v", params)
		return &Result{Success: true, Status: http.StatusOK, Data: entities.Dishes}, nil
	}

	if entities.Dishes[0].ClientToken.String() != params.Middle.ClientToken {
		log.Printf("this dish not belongs to this user: %v", params)
		return &Result{Success: false, Status: http.StatusForbidden}, fmt.Errorf("this dish not belongs to this user")
	}

	log.Printf("list dished finished: %v", params)

	return &Result{Success: true, Status: http.StatusOK, Data: entities.Dishes}, nil
}

func (c *Command) getRelatedEntities(params Params) (*RelatedEntities, error) {
	if MAX_PAGINATION_SIZE < params.Core.Size {
		params.Core.Size = MAX_PAGINATION_SIZE
	}

	dishes, err := c.DishRepository.List(ports.ListConf{
		Page:        params.Core.Page,
		Size:        params.Core.Size,
		ClientToken: params.Middle.ClientToken,
	})
	if err != nil {
		return nil, err
	}

	return &RelatedEntities{Dishes: dishes}, nil
}
