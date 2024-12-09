package repositories

import (
	"github.com/calango-productions/api/internal/adapters/connections"
	"github.com/calango-productions/api/internal/core/ports"
	"github.com/calango-productions/api/internal/repositories/commentrepo"
	credentialrepo "github.com/calango-productions/api/internal/repositories/credential"
	"github.com/calango-productions/api/internal/repositories/dishrepo"
	"github.com/calango-productions/api/internal/repositories/pingrepo"
	"github.com/calango-productions/api/internal/repositories/redisrepo"
	"github.com/calango-productions/api/internal/repositories/restaurantrepo"
	"github.com/calango-productions/api/internal/repositories/userrepo"
)

type RepoProvider interface {
	Credential() ports.CredentialRepository
	User() ports.UserRepository
	Ping() ports.PingRepository
	Dish() ports.DishRepository
	Redis() ports.RedisRepository
	Comment() ports.CommentRepository
	Restaurant() ports.RestaurantRepository
}

type Provider struct {
	CredentialRepository ports.CredentialRepository
	UserRepository       ports.UserRepository
	PingRepository       ports.PingRepository
	DishRepository       ports.DishRepository
	RedisRepository      ports.RedisRepository
	CommentRepository    ports.CommentRepository
	RestaurantRepository ports.RestaurantRepository
}

func (p Provider) User() ports.UserRepository             { return p.UserRepository }
func (p Provider) Ping() ports.PingRepository             { return p.PingRepository }
func (p Provider) Credential() ports.CredentialRepository { return p.CredentialRepository }
func (p Provider) Dish() ports.DishRepository             { return p.DishRepository }
func (p Provider) Redis() ports.RedisRepository           { return p.RedisRepository }
func (p Provider) Comment() ports.CommentRepository       { return p.CommentRepository }
func (p Provider) Restaurant() ports.RestaurantRepository { return p.RestaurantRepository }

func New(conn *connections.Connections) *Provider {
	return &Provider{
		UserRepository:       userrepo.New(conn.Databases.Core),
		PingRepository:       pingrepo.New(conn),
		CredentialRepository: credentialrepo.New(conn.Databases.Core),
		DishRepository:       dishrepo.New(conn.Databases.Core),
		RedisRepository:      redisrepo.New(conn.Databases.Redis),
		CommentRepository:    commentrepo.New(conn.Databases.Core),
		RestaurantRepository: restaurantrepo.New(conn.Databases.Core),
	}
}
