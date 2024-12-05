package repositories

import (
	"github.com/calango-productions/api/internal/adapters/connections"
	"github.com/calango-productions/api/internal/core/ports"
	"github.com/calango-productions/api/internal/repositories/clientrepo"
	credentialrepo "github.com/calango-productions/api/internal/repositories/credential"
	"github.com/calango-productions/api/internal/repositories/dishrepo"
	"github.com/calango-productions/api/internal/repositories/pingrepo"
	pwdrecorepo "github.com/calango-productions/api/internal/repositories/pwdRecovery"
	"github.com/calango-productions/api/internal/repositories/redisrepo"
	"github.com/calango-productions/api/internal/repositories/userrepo"
)

type RepoProvider interface {
	Credential() ports.CredentialRepository
	PwdRecovery() ports.PasswordRecoveryRepository
	User() ports.UserRepository
	Ping() ports.PingRepository
	Dish() ports.DishRepository
	Client() ports.ClientRepository
	Redis() ports.RedisRepository
}

type Provider struct {
	CredentialRepository  ports.CredentialRepository
	PwdRecoveryRepository ports.PasswordRecoveryRepository
	UserRepository        ports.UserRepository
	PingRepository        ports.PingRepository
	DishRepository        ports.DishRepository
	ClientRepository      ports.ClientRepository
	RedisRepository       ports.RedisRepository
}

func (p Provider) User() ports.UserRepository                    { return p.UserRepository }
func (p Provider) Ping() ports.PingRepository                    { return p.PingRepository }
func (p Provider) Credential() ports.CredentialRepository        { return p.CredentialRepository }
func (p Provider) PwdRecovery() ports.PasswordRecoveryRepository { return p.PwdRecoveryRepository }
func (p Provider) Client() ports.ClientRepository                { return p.ClientRepository }
func (p Provider) Dish() ports.DishRepository                    { return p.DishRepository }
func (p Provider) Redis() ports.RedisRepository                  { return p.RedisRepository }

func New(conn *connections.Connections) *Provider {
	return &Provider{
		UserRepository:        userrepo.New(conn.Databases.Core),
		PingRepository:        pingrepo.New(conn),
		CredentialRepository:  credentialrepo.New(conn.Databases.Core),
		PwdRecoveryRepository: pwdrecorepo.New(conn.Databases.Core),
		DishRepository:        dishrepo.New(conn.Databases.Core),
		ClientRepository:      clientrepo.New(conn.Databases.Core),
		RedisRepository:       redisrepo.New(conn.Databases.Redis),
	}
}
