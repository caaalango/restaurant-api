package jwtservice

import (
	"errors"
	"time"

	"github.com/calango-productions/api/internal/envs"
	"github.com/calango-productions/api/internal/types"
	"github.com/golang-jwt/jwt/v4"
)

var tokenDuration = time.Minute * 30

type JwtService struct {
	SecretKey     string
	TokenDuration time.Duration
}

func New() *JwtService {
	return &JwtService{
		SecretKey:     envs.Get(envs.JWT_SECRET),
		TokenDuration: tokenDuration,
	}
}

type Claims struct {
	UserToken   string `json:"userToken"`
	ClientToken string `json:"clientToken"`
	jwt.RegisteredClaims
}

func (tm *JwtService) CreateToken(userData types.UserData, role string) (string, error) {
	claims := Claims{
		UserToken:   userData.UserToken,
		ClientToken: userData.ClientToken,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tm.TokenDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(tm.SecretKey))
}

func (tm *JwtService) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(tm.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
