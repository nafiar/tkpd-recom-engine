package view

import (
	"log"

	"github.com/nsqio/go-nsq"
)

// HandleMessage will handle listened message
// intentionally removed handler logic
// reference : Slide 15 `Create NSQ Handler Func`
func (v *viewUserActivity) HandleMessage(message *nsq.Message) error {
	// safeguard when repository is nil
	if v.recentViewRepo == nil {
		message.Requeue(0)
		return nil
	}

	// this is the process to parse NSQ message to ViewMessage struct
	parsedMsg, ok := parseAndValidateView(message.Body)
	if !ok {
		log.Println("error when parsing message ")
		message.Finish()
		return nil
	}

	err := v.recentViewRepo.SetRecentView(parsedMsg.UserID, []int{parsedMsg.ProductID})
	if err != nil {
		return err
	}

	return nil
}
