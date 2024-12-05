package queues

type QueueConfig struct {
	Name       string
	Type       string
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	NoWait     bool
	Args       map[string]interface{}
}

var QUEUES_CONFIG = []QueueConfig{
	{
		Name:       "main",
		Type:       "fifo",
		Durable:    true,
		AutoDelete: false,
		Exclusive:  false,
		NoWait:     false,
		Args: map[string]interface{}{
			"x-queue-type":  "quorum",
			"x-message-ttl": 60000,
			"x-max-length":  500,
		},
	},
}
