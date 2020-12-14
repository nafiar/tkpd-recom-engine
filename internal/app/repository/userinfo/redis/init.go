package redis

import (
	"github.com/gomodule/redigo/redis"
	"github.com/nafiar/tkpd-recom-engine/internal/app/repository/userinfo"
)

// redisUserInfo strore connection pool for user info
type redisUserInfo struct {
	redisPool *redis.Pool
}

// New create new object for user info data source
func New(pool *redis.Pool) userinfo.Repository {
	return &redisUserInfo{
		redisPool: pool,
	}
}
