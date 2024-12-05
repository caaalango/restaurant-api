package connections

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/calango-productions/api/internal/adapters/config"
	coreconn "github.com/calango-productions/api/internal/database/core"
	"github.com/go-redis/redis/v8"
	"github.com/gocraft/dbr/v2"
	"github.com/opensearch-project/opensearch-go"
	"github.com/streadway/amqp"
)

type Connections struct {
	Databases  DatabasesConn
	OpenSearch *opensearch.Client
	RabbitMQ   *amqp.Connection
	Closers    []io.Closer
}

type DatabasesConn struct {
	Core  *dbr.Connection
	Redis *redis.Client
}

func New() *Connections {
	return &Connections{}
}

func (c *Connections) ConnectCoreDatabase(conf *config.Config) {
	session, err := coreconn.ConnectCoreDatabase(conf.Databases.Core.DSN)
	fmt.Println(conf.Databases.Core.DSN)
	if err != nil {
		errMessage := fmt.Sprintf("Unable to establish connection with core database: %v", err)
		panic(errMessage)
	}
	c.Databases.Core = session
	c.Closers = append(c.Closers, session)
}

func (c *Connections) ConnectRedis(conf *config.Config) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.Databases.Redis.DSN,
		Password: conf.Databases.Redis.Password,
		DB:       conf.Databases.Redis.DB,
	})

	ctx := context.Background()

	if err := rdb.Ping(ctx).Err(); err != nil {
		errMessage := fmt.Sprintf("Unable to establish connection with Redis: %v", err)
		panic(errMessage)
	}

	c.Databases.Redis = rdb
	c.Closers = append(c.Closers, rdb)
}

func (c *Connections) ConnectRabbitMQ(conf *config.Config) {
	rabbitConn, err := amqp.Dial(conf.RabbitMQURL)
	if err != nil {
		errMessage := fmt.Sprintf("Unable to establish connection with rabbit MQ: %v", err)
		panic(errMessage)
	}
	c.RabbitMQ = rabbitConn
	c.Closers = append(c.Closers, rabbitConn)
}

func (c *Connections) ConnectOpenSearch(conf *config.Config) {
	opensearchClient, err := opensearch.NewClient(opensearch.Config{
		Addresses: []string{conf.OpenSearchURL},
	})
	if err != nil {
		errMessage := fmt.Sprintf("Unable to establish connection with OpenSearch: %v", err)
		panic(errMessage)
	}
	c.OpenSearch = opensearchClient
	c.Closers = append(c.Closers, DefaulCloser)
}

func (c *Connections) Shutdown(ctx context.Context) {
	for _, closer := range c.Closers {
		select {
		case <-ctx.Done():
			log.Println("Shutdown aborted due to context cancellation")
			return
		default:
			if err := closer.Close(); err != nil {
				log.Printf("Error closing component: %v", err)
			}
		}
	}
}
