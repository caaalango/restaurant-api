package pwdrecorepo

import (
	"sync"

	"github.com/calango-productions/api/internal/core/entities"
	"github.com/calango-productions/api/internal/core/ports"
	"github.com/calango-productions/api/internal/repositories/baserepo"
	"github.com/gocraft/dbr/v2"
)

var (
	instance *PasswordRecoveryRepository
	once     sync.Once
	table    string
)

type PasswordRecoveryRepository struct {
	conn  *dbr.Connection
	table string
	*baserepo.BaseRepository[entities.PasswordRecovery]
}

func New(conn *dbr.Connection) ports.PasswordRecoveryRepository {
	table = "pwd_recoveries"
	once.Do(func() {
		instance = &PasswordRecoveryRepository{
			BaseRepository: baserepo.New[entities.PasswordRecovery](conn, table),
		}
		instance.conn = conn
		instance.table = table
	})

	return instance
}
