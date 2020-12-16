package nsq

import (
	"context"
	"log"
	"strings"

	nsq "github.com/nsqio/go-nsq"

	view_delivery "github.com/nafiar/tkpd-recom-engine/internal/background/delivery/view"
	"github.com/nafiar/tkpd-recom-engine/internal/background/usecase/recentview"
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

func NewViewListener(usecase recentview.Usecase, listenerConfig ListenerConfig) (view_delivery.Delivery, error) {
	cfg := nsq.NewConfig()
	cfg.MaxInFlight = listenerConfig.MaxInFlight

	consumer, err := nsq.NewConsumer(listenerConfig.Topic, listenerConfig.ChannelName, cfg)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	handler := &viewHandler{
		usecase: usecase,
	}

	consumer.AddConcurrentHandlers(handler, listenerConfig.Worker)
	consumer.SetLoggerLevel(nsq.LogLevelError)

	return &viewListener{
		lookupdAddress: listenerConfig.NSQLookupd,
		consumer:       consumer,
	}, nil
}

func (v *viewListener) Run(ctx context.Context) {
	lookupds := strings.Split(v.lookupdAddress, ",")
	err := v.consumer.ConnectToNSQLookupds(lookupds)
	if err != nil {
		log.Println(err)
	}
}
