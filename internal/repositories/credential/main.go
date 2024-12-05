package credentialrepo

import (
	"sync"

	"github.com/calango-productions/api/internal/core/entities"
	"github.com/calango-productions/api/internal/core/ports"
	"github.com/calango-productions/api/internal/repositories/baserepo"
	"github.com/gocraft/dbr/v2"
)

var (
	instance *CredentialRepository
	once     sync.Once
	table    string
)

type CredentialRepository struct {
	conn  *dbr.Connection
	table string
	*baserepo.BaseRepository[entities.Credential]
}

func New(conn *dbr.Connection) ports.CredentialRepository {
	table = "credentials"
	once.Do(func() {
		instance = &CredentialRepository{
			BaseRepository: baserepo.New[entities.Credential](conn, table),
		}
		instance.conn = conn
		instance.table = table
	})

	return instance
}
