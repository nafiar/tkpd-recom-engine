package main

import (
	"log"
	nsq "github.com/nsqio/go-nsq"
)

func main() {
	cfg := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("test", "dari_golang", cfg)
	if err != nil {
		log.Fatalln(err)
	}
	handler := &simpleHandler{}
	consumer.AddConcurrentHandlers(handler, 1)
	err = consumer.ConnectToNSQLookupds([]string{"127.0.0.1:4161"})
	if err != nil {
		log.Fatalln(err)
	}
	done := make(chan int)
	<-done
}



