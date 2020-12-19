package redis

import (
	"log"
	"strconv"

	"github.com/gomodule/redigo/redis"
)

func (r *redisRecentView) SetRecentView(userID int, productIDs []int) (err error) {
	redisConn := r.redisPool.Get()
	defer redisConn.Close()
	redisKey := RECENT_VIEW_KEY + strconv.Itoa(userID)
	var redisArgs []interface{}
	redisArgs = append(redisArgs, redisKey)
	for _, pid := range productIDs {
		redisArgs = append(redisArgs, pid)
	}
	_, err = redis.Int64(redisConn.Do("LPUSH", redisArgs...))
	if err != nil {
		log.Println("error cannot do lpush")
	}

	_, err = redis.Int64(redisConn.Do("EXPIRE", redisKey, 300))
	if err != nil {
		log.Println("error cannot do expire")
	}
	return
}
