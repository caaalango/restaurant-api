package showmenucmd

import (
	"github.com/calango-productions/api/internal/core/entities"
	"github.com/calango-productions/api/internal/core/ports"
)

type Command struct {
	RestaurantRepository ports.RestaurantRepository
	DishRepository       ports.DishRepository
	CommentRepository    ports.CommentRepository
	RedisRepository      ports.RedisRepository
}

type Params struct {
	Core CoreParams
}

type CoreParams struct {
	Name string
}

type RelatedEntities struct {
	Restaurant *entities.Restaurant
	Dishes     []entities.Dish
	Comments   []entities.Comment
}

type Result struct {
	Success bool
	Status  int
	Data    RestaurantResult
}

type RestaurantResult struct {
	Name       string           `json:"name"`
	Slogo      string           `json:"slogo"`
	Categories []CategoryResult `json:"categories"`
}

type CategoryResult struct {
	Category string       `json:"category"`
	Dishes   []DishResult `json:"dishes"`
}

type DishResult struct {
	Dish     entities.Dish   `json:"dish"`
	Comments []CommentResult `json:"comments"`
	Rating   int             `json:"rating"`
	Notes    int             `json:"notes"`
}

type CommentResult struct {
	Comment entities.Comment `json:"comment"`
}
