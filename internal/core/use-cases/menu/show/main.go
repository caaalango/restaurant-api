package showmenucmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/calango-productions/api/internal/core/entities"
	"github.com/calango-productions/api/internal/core/ports"
)

func New(
	restaurantRepository ports.RestaurantRepository,
	dishRepository ports.DishRepository,
	commentRepository ports.CommentRepository,
	redisRepository ports.RedisRepository,
) *Command {
	return &Command{
		RestaurantRepository: restaurantRepository,
		DishRepository:       dishRepository,
		CommentRepository:    commentRepository,
		RedisRepository:      redisRepository,
	}
}

func (c *Command) Execute(params Params) (*Result, error) {
	log.Printf("ShowMenuCommand initiated: %+v", params)

	entities, err := c.getRelatedEntities(params)
	if err != nil {
		log.Printf("ShowMenuCommand failed: %+v, %+v", params, err)
		return &Result{Success: false, Status: http.StatusInternalServerError}, err
	}

	if entities.Restaurant == nil {
		log.Printf("ShowMenuCommand failed: %+v, %+v", params, err)
		return &Result{Success: false, Status: http.StatusNotFound}, err
	}

	result := c.mapDishesByCategory(*entities.Restaurant, entities.Dishes, entities.Comments)

	log.Printf("ShowMenuCommand finished: %+v", params)

	return &Result{Success: true, Status: http.StatusOK, Data: result}, err
}

func (c *Command) getRelatedEntities(params Params) (*RelatedEntities, error) {
	restaurant, err := c.RestaurantRepository.GetByKey(ports.GetByKeyConf{
		Key:         "name",
		Value:       params.Core.Name,
		OnlyActives: true,
	})
	if err != nil {
		return nil, err
	}

	dishes, err := c.DishRepository.List(ports.ListConf{
		FilterField:   "restaurant_token",
		FilterToken:   restaurant.Token,
		HasPagination: false,
		OnlyActives:   true,
	})
	if err != nil {
		return nil, err
	}

	comments, err := c.CommentRepository.List(ports.ListConf{
		FilterField:   "restaurant_token",
		FilterToken:   restaurant.Token,
		HasPagination: false,
		OnlyActives:   true,
	})
	if err != nil {
		return nil, err
	}

	data := &RelatedEntities{
		Restaurant: restaurant,
		Dishes:     dishes,
		Comments:   comments,
	}

	fmt.Printf("%+v", data)

	return data, nil
}

func (c *Command) mapDishesByCategory(restaurant entities.Restaurant, dishes []entities.Dish, comments []entities.Comment) RestaurantResult {
	result := RestaurantResult{}
	result.Name = restaurant.Name
	result.Slogo = restaurant.Slogo

	dishesByCategory := c.selectDishByCategory(dishes)

	for category, dishes := range dishesByCategory {
		categoryResult := CategoryResult{
			Category: category,
		}

		for _, dish := range dishes {
			selectedComments := c.selectCommentsByDish(dish, comments)
			dishResult := DishResult{
				Dish:   dish,
				Notes:  c.calculateLenDishComments(dish, selectedComments),
				Rating: c.calculateAverageDishNote(dish, selectedComments),
			}

			for _, comment := range selectedComments {
				commentResult := CommentResult{
					Comment: comment,
				}
				dishResult.Comments = append(dishResult.Comments, commentResult)
			}

			categoryResult.Dishes = append(categoryResult.Dishes, dishResult)
		}

		result.Categories = append(result.Categories, categoryResult)
	}

	return result
}

func (c *Command) selectDishByCategory(dishes []entities.Dish) map[string][]entities.Dish {
	dishesByCategory := map[string][]entities.Dish{}

	for _, dish := range dishes {
		if _, exists := dishesByCategory[string(dish.Category)]; exists {
			dishesByCategory[string(dish.Category)] = append(dishesByCategory[string(dish.Category)], dish)
		} else {
			dishesByCategory[string(dish.Category)] = []entities.Dish{dish}
		}
	}

	return dishesByCategory
}

func (c *Command) calculateLenDishComments(dish entities.Dish, comments []entities.Comment) int {
	var totalComments int = 0

	for _, comment := range comments {
		if comment.DishToken == dish.Token {
			totalComments++
		}
	}

	return totalComments
}

func (c *Command) calculateAverageDishNote(dish entities.Dish, comments []entities.Comment) int {
	totalNote := 0
	totalComments := 0
	for _, comment := range comments {
		if comment.DishToken == dish.Token {
			totalNote += comment.Note
			totalComments++
		}
	}
	if totalComments == 0 {
		return 0
	}
	totalNote = totalNote / totalComments

	return totalNote
}

func (c *Command) selectCommentsByDish(dish entities.Dish, comments []entities.Comment) []entities.Comment {
	var selected []entities.Comment

	for _, comment := range comments {
		if comment.DishToken == dish.Token {
			selected = append(selected, comment)
		}
	}

	return selected
}
