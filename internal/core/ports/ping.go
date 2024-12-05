package ports

type PingRepository interface {
	CorePing() error
	RedisPing() error
}
