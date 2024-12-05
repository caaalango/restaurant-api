package dishrepo

import (
	"sync"

	"github.com/calango-productions/api/internal/core/entities"
	"github.com/calango-productions/api/internal/core/ports"
	"github.com/calango-productions/api/internal/repositories/baserepo"
	"github.com/gocraft/dbr/v2"
)

var (
	instance *DishRepository
	once     sync.Once
	table    string
)

type DishRepository struct {
	conn  *dbr.Connection
	table string
	*baserepo.BaseRepository[entities.Dish]
}

func New(conn *dbr.Connection) ports.DishRepository {
	table = "dishes"
	once.Do(func() {
		instance = &DishRepository{
			BaseRepository: baserepo.New[entities.Dish](conn, table),
		}
		instance.conn = conn
		instance.table = table
	})

	return instance
}
