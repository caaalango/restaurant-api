package docsctl

import (
	"github.com/calango-productions/api/internal/adapters"
	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

type DocsController struct{}

func (u DocsController) SetUpRoutes(c *gin.Engine) {
	c.Group("/docs")

	docs.SwaggerInfo.BasePath = "/api"
	c.GET("/*any", ginSwagger.WrapHandler(files.Handler))
}

func New(apt *adapters.Adapters) *DocsController {
	return &DocsController{}
}
