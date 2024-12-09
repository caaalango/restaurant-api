package menuctl

import (
	"github.com/calango-productions/api/internal/adapters"
	"github.com/calango-productions/api/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type MenuController struct {
	RestaurantRepository ports.RestaurantRepository
	DishRepository       ports.DishRepository
	CommentRepository    ports.CommentRepository
	RedisRepository      ports.RedisRepository
}

func (mc MenuController) SetUpRoutes(c *gin.Engine) {
	group := c.Group("/menu/:Name")

	group.GET("/", mc.ShowAction)
}

func New(apt *adapters.Adapters) *MenuController {
	return &MenuController{
		RestaurantRepository: apt.Repositories.Restaurant(),
		DishRepository:       apt.Repositories.Dish(),
		CommentRepository:    apt.Repositories.Comment(),
		RedisRepository:      apt.Repositories.Redis(),
	}
}
