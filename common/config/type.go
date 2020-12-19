package config

type Config struct {
	Redis   map[string]*RedisCfg
	NSQ     map[string]*NSQCfg
	Postgre map[string]*PostgreCfg
}

type (
	RedisCfg struct {
		Connection string
		MaxActive  int
		MaxIdle    int
	}

	NSQCfg struct {
		ChannelName string
		MaxInFlight int
		NSQLookupd  string
		Topic       string
		Worker      int
	}

	PostgreCfg struct {
		ConnString string
	}
)
