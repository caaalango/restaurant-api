package restaurantrepo

import (
	"sync"

	"github.com/calango-productions/api/internal/core/entities"
	"github.com/calango-productions/api/internal/core/ports"
	"github.com/calango-productions/api/internal/repositories/baserepo"
	"github.com/gocraft/dbr/v2"
)

var (
	instance *RestaurantRepository
	once     sync.Once
	table    string
)

type RestaurantRepository struct {
	conn  *dbr.Connection
	table string
	*baserepo.BaseRepository[entities.Restaurant]
}

func New(conn *dbr.Connection) ports.RestaurantRepository {
	table = "restaurants"
	once.Do(func() {
		instance = &RestaurantRepository{
			BaseRepository: baserepo.New[entities.Restaurant](conn, table),
		}
		instance.conn = conn
		instance.table = table
	})

	return instance
}
