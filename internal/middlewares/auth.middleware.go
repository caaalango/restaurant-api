package middlewares

import (
	"net/http"

	jwtservice "github.com/calango-productions/api/internal/core/services/jwt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	jwtService *jwtservice.JwtService
}

func NewAuthMiddleware(service *jwtservice.JwtService) *AuthMiddleware {
	return &AuthMiddleware{
		jwtService: service,
	}
}

func (m *AuthMiddleware) Execute() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)

		if session.Get("userData") != nil {
			ctx.Next()
			return
		}

		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			ctx.Abort()
			return
		}

		userData, err := m.jwtService.ValidateToken(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			ctx.Abort()
			return
		}

		session.Set("userData", userData)
		session.Save()

		ctx.Set("ActorToken", userData)
		ctx.Next()
	}
}
