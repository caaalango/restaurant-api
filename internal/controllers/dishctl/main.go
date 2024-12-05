package dishctl

import (
	"github.com/calango-productions/api/internal/adapters"
	"github.com/calango-productions/api/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type DishController struct {
	DishRepository  ports.DishRepository
	RedisRepository ports.RedisRepository
}

func (u DishController) SetUpRoutes(c *gin.Engine) {
	c.Group("/dishes")

	c.POST("/", u.CreateAction)
	c.GET("/", u.ListAction)
	c.GET("/:DishToken", u.GetAction)
	c.PUT("/:DishToken", u.UpdateAction)
	c.DELETE("/:DishToken", u.InactivateAction)
}

func New(apt *adapters.Adapters) *DishController {
	return &DishController{
		RedisRepository: apt.Repositories.Redis(),
		DishRepository:  apt.Repositories.Dish(),
	}
}
