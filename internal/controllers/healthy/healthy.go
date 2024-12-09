package healthy

import (
	"log"
	"net/http"

	"github.com/calango-productions/api/internal/adapters"
	"github.com/gin-gonic/gin"
)

type HealthController struct {
	adapters *adapters.Adapters
}

func New(adapters *adapters.Adapters) HealthController {
	return HealthController{adapters: adapters}
}

func (h HealthController) SetUpRoutes(c *gin.Engine) {
	c.GET("/health", h.Health)
	c.GET("/checks/coredb", h.Health)
	c.GET("/checks/redis", h.Health)
}

func (h HealthController) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "system running",
	})
}

func (h HealthController) CheckCoreDatabase(c *gin.Context) {
	if err := h.adapters.Repositories.PingRepository.CorePing(); err != nil {
		log.Printf("failed to ping core database: err: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "core database system unavailable",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "system running",
	})
}

func (h HealthController) CheckRedis(c *gin.Context) {
	if err := h.adapters.Repositories.PingRepository.RedisPing(); err != nil {
		log.Printf("failed to ping redis: err: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "redis system unavailable",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "database running",
	})
}
