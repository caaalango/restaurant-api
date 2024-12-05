package pingrepo

import (
	"context"
	"sync"

	"github.com/calango-productions/api/internal/adapters/connections"
)

var (
	instance PingRepository
	once     sync.Once
)

type PingRepository struct {
	connections *connections.Connections
}

func New(conn *connections.Connections) PingRepository {
	once.Do(func() {
		instance = PingRepository{connections: conn}
	})
	return instance
}

func (p PingRepository) CorePing() error {
	return p.connections.Databases.Core.Ping()
}

func (p PingRepository) RedisPing() error {
	ctx := context.Background()
	return p.connections.Databases.Redis.Ping(ctx).Err()
}
