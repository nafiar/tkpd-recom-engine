package config

type Config struct {
	Redis map[string]*RedisCfg
}

type (
	RedisCfg struct {
		Connection string
		MaxActive  int
		MaxIdle    int
	}
)
