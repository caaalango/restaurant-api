package config

import (
	"fmt"

	"github.com/calango-productions/api/internal/envs"
)

type Config struct {
	ServerPort    string
	Databases     Databases
	RabbitMQURL   string
	OpenSearchURL string
}

type Databases struct {
	Core  Core
	Redis Redis
}

type Core struct {
	DSN string
}

type Redis struct {
	DSN      string
	Password string
	DB       int
}

func New() *Config {
	envs.Load()

	return &Config{
		ServerPort: envs.Get(envs.SERVER_PORT),
		Databases: Databases{
			Core{DSN: buildCoreDatabaseStringConnection()},
			Redis{DSN: buildCoreRedisStringConnection(),
				Password: envs.Get(envs.REDIS_PASSWORD),
				DB:       envs.GetInt(envs.REDIS_DB)},
		},
		RabbitMQURL:   envs.Get(envs.RABBITMQ_URL),
		OpenSearchURL: envs.Get(envs.OPENSEARCH_URL),
	}
}

func buildCoreDatabaseStringConnection() string {
	dbUsername := envs.Get(envs.DB_USERNAME)
	dbPassword := envs.Get(envs.DB_PASSWORD)
	dbHost := envs.Get(envs.DB_HOST)
	dbPort := envs.Get(envs.DB_PORT)
	dbName := envs.Get(envs.DB_NAME)

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		dbHost,
		dbUsername,
		dbPassword,
		dbName,
		dbPort,
	)

	return dsn
}

func buildCoreRedisStringConnection() string {
	redisHost := envs.Get(envs.REDIS_HOST)
	redisPort := envs.Get(envs.REDIS_PORT)

	dsn := fmt.Sprintf(
		"%s:%s",
		redisHost,
		redisPort,
	)

	return dsn
}
