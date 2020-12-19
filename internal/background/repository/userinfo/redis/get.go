package redis

import (
	"strconv"
	"strings"

	"github.com/gomodule/redigo/redis"
	m "github.com/nafiar/tkpd-recom-engine/internal/model/user"
)

// GetData get user data from redis the parse it's response into user model data
func (r *redisUserInfo) GetData(userID string) (data m.Data, err error) {
	redisConn := r.redisPool.Get()
	defer redisConn.Close()

	redisKey := UserInfoKey + userID

	resp, err := redis.String(redisConn.Do("GET", redisKey))
	if err != nil {
		return
	}

	data = parseToData(resp)
	data.ID = userID

	return
}

// parseToData will split response using | delimiter,
// then put parsed value into result
func parseToData(resp string) (result m.Data) {
	parsedValue := make([]string, ExpectedValueLength)
	infos := strings.Split(resp, "|")

	for i, info := range infos {
		if i >= ExpectedValueLength {
			break
		}

		parsedValue[i] = info
	}

	age, _ := strconv.Atoi(parsedValue[1])

	result.Name = parsedValue[0]
	result.Age = age

	return
}
