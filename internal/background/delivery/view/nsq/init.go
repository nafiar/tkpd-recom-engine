package nsq

import (
	"log"

	nsq "github.com/nsqio/go-nsq"

	view_delivery "github.com/nafiar/tkpd-recom-engine/internal/background/delivery/view"
	"github.com/nafiar/tkpd-recom-engine/internal/background/usecase/useractivity"
)

type viewListener struct {
	lookupdAddress string
	consumer       *nsq.Consumer
}

type ListenerConfig struct {
	MaxInFlight int
	ChannelName string
	NSQLookupd  string
	Topic       string
	Worker      int
}

// New create new object for nsq view listener
// intentionally removed consumer initialization
// reference : NSQ Slide 14 `Creating a new listener`
func New(usecase useractivity.UseCase, listenerConfig ListenerConfig) (view_delivery.Delivery, error) {
	cfg := nsq.NewConfig()
	cfg.MaxInFlight = listenerConfig.MaxInFlight

	consumer, err := nsq.NewConsumer(listenerConfig.Topic, listenerConfig.ChannelName, cfg)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	consumer.AddConcurrentHandlers(usecase, listenerConfig.Worker)
	consumer.SetLoggerLevel(nsq.LogLevelError)

	return &viewListener{
		lookupdAddress: listenerConfig.NSQLookupd,
		consumer:       consumer,
	}, nil
}
