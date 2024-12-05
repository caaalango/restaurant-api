package middlewares

import jwtservice "github.com/calango-productions/api/internal/core/services/jwt"

type Middlewares interface {
	AuthMiddleware() AuthMiddleware
	CorsMiddleware() CorsMiddleware
}

type Provider struct {
	AuthMiddleware AuthMiddleware
	CorsMiddleware CorsMiddleware
}

func (p Provider) Auth() AuthMiddleware { return p.AuthMiddleware }
func (p Provider) Cors() CorsMiddleware { return p.CorsMiddleware }

func New() *Provider {
	jwtService := jwtservice.New()

	return &Provider{
		AuthMiddleware: *NewAuthMiddleware(jwtService),
		CorsMiddleware: *NewCorsMiddleware(),
	}
}
