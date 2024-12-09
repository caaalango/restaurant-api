package docsctl

import (
	"github.com/calango-productions/api/internal/adapters"
	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

type DocsController struct {
	adapters *adapters.Adapters
}

func New(adapters *adapters.Adapters) DocsController {
	return DocsController{adapters: adapters}
}

func (u DocsController) SetUpRoutes(c *gin.Engine) {
	group := c.Group("/docs")

	docs.SwaggerInfo.BasePath = "/api"

	group.GET("/*any", ginSwagger.WrapHandler(files.Handler))
}
