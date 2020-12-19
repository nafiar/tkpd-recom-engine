package view

import (
	"encoding/json"
	"time"
)

// messageData store NSQ message json fields
type messageData struct {
	UserID    json.Number `json:"user_id"`
	ProductID json.Number `json:"product_id"`
	Timestamp string      `json:"timestamp"`
}

// RecentViewSignal store recent view normal data
type RecentViewSignal struct {
	UserID    int
	ProductID int
	Timestamp time.Time
}
