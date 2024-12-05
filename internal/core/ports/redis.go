package ports

import "time"

type RedisStoreConf struct {
	Key        string
	Expiration time.Duration
	Data       any
}

type RedisRescueConf struct {
	Key  string
	Data any
}

type RedisRepository interface {
	Store(RedisStoreConf) error
	Rescue(RedisRescueConf) (any, error)
}
