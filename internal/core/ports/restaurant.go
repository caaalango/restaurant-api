package ports

import "github.com/calango-productions/api/internal/core/entities"

type RestaurantRepository interface {
	BasePort[entities.Restaurant]
}
