package handler

import (
	"log"

	jsoniter "github.com/json-iterator/go"
)

func (r *recentViewUseCase) SetRecentView(messageBody []byte) (err error) {
	message := RecentViewMessage{}
	err = jsoniter.ConfigFastest.Unmarshal(messageBody, &message)
	if err != nil {
		log.Println("Error failed to unmarshall")
	}
	if message.UserID > 0 && message.ProductID > 0 {
		err = r.recentViewRepo.SetRecentView(message.UserID, []int{message.ProductID})
	}
	return
}
