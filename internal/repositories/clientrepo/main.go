package clientrepo

import (
	"sync"

	"github.com/calango-productions/api/internal/core/entities"
	"github.com/calango-productions/api/internal/core/ports"
	"github.com/calango-productions/api/internal/repositories/baserepo"
	"github.com/gocraft/dbr/v2"
)

var (
	instance *ClientRepository
	once     sync.Once
	table    string
)

type ClientRepository struct {
	conn  *dbr.Connection
	table string
	*baserepo.BaseRepository[entities.Client]
}

func New(conn *dbr.Connection) ports.ClientRepository {
	table = "clients"
	once.Do(func() {
		instance = &ClientRepository{
			BaseRepository: baserepo.New[entities.Client](conn, table),
		}
		instance.conn = conn
		instance.table = table
	})

	return instance
}
