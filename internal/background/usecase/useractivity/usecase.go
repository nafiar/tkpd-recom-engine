package useractivity

import (
	"github.com/nsqio/go-nsq"
)

// Usecase represent recent view business logic
type Usecase interface {
	HandleMessage(msg *nsq.Message) error
}
