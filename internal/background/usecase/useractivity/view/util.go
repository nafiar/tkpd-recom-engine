package view

import (
	"encoding/json"
	"time"
	// "github.com/nafiar/tkpd-recom-engine/internal/background/usecase/recentview"
)

// ParseAndValidateView parses NSQ View message in to user product relation
func parseAndValidateView(body []byte) (RecentViewSignal, bool) {
	var parsed messageData
	var result RecentViewSignal

	err := json.Unmarshal(body, &parsed)
	if err != nil {
		return result, false
	}

	userID, err := parsed.UserID.Int64()
	if err != nil || userID == 0 {
		return result, false
	}

	productID, err := parsed.ProductID.Int64()
	if err != nil || productID == 0 {
		return result, false
	}

	timestamp, err := time.Parse(time.RFC3339, parsed.Timestamp)
	if err != nil {
		return result, false
	}

	result = RecentViewSignal{
		UserID:    int(userID),
		ProductID: int(productID),
		Timestamp: timestamp,
	}

	return result, true
}
