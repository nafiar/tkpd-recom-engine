package redis

import (
	"log"
	"strconv"
)

// GetRecentView get recent view activity data from redis
// intentionally remove it's implementation
// reference : slide redis `Get recent view data`
func (r *redisRecentView) GetRecentView(userID int) (result []int, err error) {
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
