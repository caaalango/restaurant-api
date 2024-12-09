package ports

import "github.com/calango-productions/api/internal/core/entities"

type ListCompleteConf struct {
	RestaurantToken string
}

type CommentRepository interface {
	BasePort[entities.Comment]
}
