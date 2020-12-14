package redis

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

// ConnectionConfig is a configuration for creating a redis connection
// IdleTimeout is in second
type ConnectionConfig struct {
	Address     string
	MaxActive   int
	IdleTimeout int // Seconds
	MaxIdle     int
}

// NewConnection create new redis connection.
func NewConnection(cfg ConnectionConfig) *redis.Pool {
	if cfg.IdleTimeout == 0 {
		cfg.IdleTimeout = 10
	}

	if cfg.MaxIdle == 0 {
		cfg.MaxIdle = 100
	}

	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", cfg.Address)
		},
		IdleTimeout: time.Duration(cfg.IdleTimeout) * time.Second,
		MaxActive:   cfg.MaxActive,
		MaxIdle:     cfg.MaxIdle,
		Wait:        true,
	}
}
