package commentrepo

import (
	"sync"

	"github.com/calango-productions/api/internal/core/entities"
	"github.com/calango-productions/api/internal/core/ports"
	"github.com/calango-productions/api/internal/repositories/baserepo"
	"github.com/gocraft/dbr/v2"
)

var (
	instance *CommentRepository
	once     sync.Once
	table    string
)

type CommentRepository struct {
	conn  *dbr.Connection
	table string
	*baserepo.BaseRepository[entities.Comment]
}

func New(conn *dbr.Connection) ports.CommentRepository {
	table = "comments"
	once.Do(func() {
		instance = &CommentRepository{
			BaseRepository: baserepo.New[entities.Comment](conn, table),
		}
		instance.conn = conn
		instance.table = table
	})

	return instance
}
