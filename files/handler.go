package main

import (
	"fmt"

	nsq "github.com/nsqio/go-nsq"
)

type simpleHandler struct{}

func (h *simpleHandler) HandleMessage(message *nsq.Message) error {
	fmt.Println("From NSQ:", string(message.Body))

	// we can do many things here:
	// we can parse and validate the message
	// and store the data to redis.
	return nil
}




