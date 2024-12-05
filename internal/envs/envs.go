package envs

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type key int

const (
	SERVER_PORT key = iota + 1
	DB_USERNAME
	DB_PASSWORD
	DB_HOST
	DB_PORT
	DB_NAME
	RABBITMQ_URL
	OPENSEARCH_URL
	JWT_SECRET
	REDIS_HOST
	REDIS_PORT
	REDIS_PASSWORD
	REDIS_DB
	OPENAI_API_KEY
)

func (k key) String() string {
	switch k {
	case SERVER_PORT:
		return "SERVER_PORT"
	case DB_USERNAME:
		return "DB_USERNAME"
	case DB_PASSWORD:
		return "DB_PASSWORD"
	case DB_HOST:
		return "DB_HOST"
	case DB_PORT:
		return "DB_PORT"
	case DB_NAME:
		return "DB_NAME"
	case RABBITMQ_URL:
		return "RABBITMQ_URL"
	case OPENSEARCH_URL:
		return "OPENSEARCH_URL"
	case JWT_SECRET:
		return "JWT_SECRET"
	case REDIS_HOST:
		return "REDIS_HOST"
	case REDIS_PORT:
		return "REDIS_PORT"
	case REDIS_PASSWORD:
		return "REDIS_PASSWORD"
	case REDIS_DB:
		return "REDIS_DB"
	case OPENAI_API_KEY:
		return "OPENAI_API_KEY"
	default:
		return "Unknown"
	}
}

func Get(key key) (value string) {
	value = os.Getenv(key.String())
	return value
}

func GetInt(key key) (value int) {
	strValue := os.Getenv(key.String())
	value, _ = strconv.Atoi(strValue)
	return
}

func Load() {
	_ = godotenv.Load()
}
