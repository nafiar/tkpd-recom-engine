package config

type Config struct {
	Redis   map[string]*RedisCfg
	Postgre map[string]*PostgreCfg
}

type (
	RedisCfg struct {
		Connection string
		MaxActive  int
		MaxIdle    int
	}

	PostgreCfg struct {
		ConnString string
	}
)
