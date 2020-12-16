package nsq

import (
	"encoding/json"
	"time"

	nsq "github.com/nsqio/go-nsq"

	recentViewUC "github.com/nafiar/tkpd-recom-engine/internal/app/usecase/recentview"
	"github.com/nafiar/tkpd-recom-engine/internal/background/usecase/recentview"
)

type viewHandler struct {
	recentViewUC recentViewUC.UseCase
}

type ViewMessage struct {
	UserID    json.Number `json:"user_id"`
	ProductID json.Number `json:"product_id"`
	Timestamp string      `json:"timestamp"`
}

func (h *viewHandler) HandleMessage(message *nsq.Message) error {
	if h.recentViewUC == nil {
		message.Requeue(0)
		return nil
	}
	h.recentViewUC.SetRecentView(message.Body)
	// parsedMsg, ok := parseAndValidateView(message.Body)
	// if !ok {
	// 	message.Finish()
	// 	return nil
	// }

	return nil
}

// ParseAndValidateView parses NSQ View message in to user product relation
func parseAndValidateView(body []byte) (recentview.RecentViewSignal, bool) {
	var parsed ViewMessage
	var result recentview.RecentViewSignal

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

	result = recentview.RecentViewSignal{
		UserID:    int(userID),
		ProductID: int(productID),
		Timestamp: timestamp,
	}

	return result, true
}
