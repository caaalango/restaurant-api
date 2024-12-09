package userctl

import (
	"github.com/calango-productions/api/internal/adapters"
	"github.com/calango-productions/api/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userRepository       ports.UserRepository
	credentialRepository ports.CredentialRepository
}

func (u UserController) SetUpRoutes(c *gin.Engine) {
	group := c.Group("/users")

	group.POST("/", u.CreateTradionalAction)
}

func New(apt *adapters.Adapters) *UserController {
	return &UserController{
		userRepository:       apt.Repositories.User(),
		credentialRepository: apt.Repositories.Credential(),
	}
}
