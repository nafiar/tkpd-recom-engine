package redis

import (
	"github.com/gomodule/redigo/redis"
	"github.com/nafiar/tkpd-recom-engine/internal/serving/repository/recentview"
)

// redisRecentView store connection pool for user info
type redisRecentView struct {
	redisPool *redis.Pool
}

// New create new object for recent view data source
func New(pool *redis.Pool) recentview.Repository {
	return &redisRecentView{
		redisPool: pool,
	}
}
