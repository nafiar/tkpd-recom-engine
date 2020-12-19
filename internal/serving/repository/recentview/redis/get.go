package redis

import (
	"log"
	"strconv"

	"github.com/gomodule/redigo/redis"
)

// GetRecentView get recent view activity data from redis
// intentionally remove it's implementation
// reference : slide redis `Get recent view data`
func (r *redisRecentView) GetRecentView(userID int) (result []int, err error) {
	redisConn := r.redisPool.Get()
	defer redisConn.Close()
	redisKey := RECENT_VIEW_KEY + strconv.Itoa(userID)
	tmpResult, err := redis.Strings(redisConn.Do("LRANGE", redisKey, 0, 0))
	if err != nil {
		return
	}
	result = AtoiArray(tmpResult)

	return
}

// AtoiArray Parse array string into array int
// it will skip the entry if the string cannot be converted to int
func AtoiArray(data []string) []int {
	result := make([]int, 0, len(data))
	for _, d := range data {
		if d == "" {
			continue
		}
		tmpInt, err := strconv.Atoi(d)
		if err != nil {
			log.Println("Failed to convert")
		}
		result = append(result, tmpInt)
	}
	return result
}
