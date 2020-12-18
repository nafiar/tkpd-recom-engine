package redis

import (
	"github.com/gomodule/redigo/redis"

	"github.com/nafiar/tkpd-recom-engine/internal/app/repository/recentView"
)

// redisRecentView store connection pool for user info
type redisRecentView struct {
	redisPool *redis.Pool
}

// New create new object for recent view data source
func New(pool *redis.Pool) recentView.Repository {
	return &redisRecentView{
		redisPool: pool,
	}
}
