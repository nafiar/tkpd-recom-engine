package redis

import (
	"log"
)

// SetRecentView will execute LPUSH command to push user recent view data
// Intentionally removed
// reference slide : redis set function
func (r *redisRecentView) SetRecentView(userID int, productIDs []int) (err error) {
	redisConn := r.redisPool.Get()
	defer redisConn.Close()

	log.Println("inserting redis userID : %v, productIDs : %v", userID, productIDs)
	return
}
