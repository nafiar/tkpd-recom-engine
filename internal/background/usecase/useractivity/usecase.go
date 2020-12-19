package useractivity

import (
	"github.com/nsqio/go-nsq"
)

// UseCase represent recent view business logic
type UseCase interface {
	HandleMessage(msg *nsq.Message) error
}
