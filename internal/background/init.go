package background

import (
	"context"
	"log"

	recentview_delivery_nsq "github.com/nafiar/tkpd-recom-engine/internal/background/delivery/view/nsq"
	recentview_usecase_handler "github.com/nafiar/tkpd-recom-engine/internal/background/usecase/recentview/handler"
)

func Run() {
	log.Println("init background")

	recentviewHandler := recentview_usecase_handler.New()

	delivery, err := recentview_delivery_nsq.NewViewListener(recentviewHandler, recentview_delivery_nsq.ListenerConfig{
		ChannelName: "recommendation_generator",
		MaxInFlight: 100,
		NSQLookupd:  "localhost:4161",
		Topic:       "recent_view_stream",
		Worker:      1,
	})
	if err != nil {
		log.Println("error initializing NSQ listener")
		return
	}

	delivery.Run(context.Background())

	done := make(chan struct{})
	<-done
}
