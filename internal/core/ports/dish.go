package ports

import "github.com/calango-productions/api/internal/core/entities"

type DishRepository interface {
	BasePort[entities.Dish]
}
