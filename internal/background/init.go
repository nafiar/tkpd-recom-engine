package background

import (
	"context"
	"fmt"
	"log"

	"github.com/nafiar/tkpd-recom-engine/common/config"
	"github.com/nafiar/tkpd-recom-engine/common/redis"
	recentview_delivery_nsq "github.com/nafiar/tkpd-recom-engine/internal/background/delivery/view/nsq"
	recentview_repository_redis "github.com/nafiar/tkpd-recom-engine/internal/background/repository/recentView/redis"
	recentview_usecase_handler "github.com/nafiar/tkpd-recom-engine/internal/background/usecase/recentview/handler"
)

func Run() {
	log.Println("init background")

	cfg := config.GetConfig()

	err := config.InitConfig()
	if err != nil {
		err = fmt.Errorf("Could not init config. Err: %+v", err)
		log.Fatal(err)
	}

	cfg = config.GetConfig()

	redisUserInfo := redis.NewConnection(redis.ConnectionConfig{
		Address:   cfg.Redis["user_data"].Connection,
		MaxActive: cfg.Redis["user_data"].MaxActive,
		MaxIdle:   cfg.Redis["user_data"].MaxIdle,
	})

	recentviewRepository := recentview_repository_redis.New(redisUserInfo)

	recentviewHandler := recentview_usecase_handler.New(recentviewRepository)

	delivery, err := recentview_delivery_nsq.NewViewListener(recentviewHandler, recentview_delivery_nsq.ListenerConfig{
		ChannelName: cfg.NSQ["activity_view_stream"].ChannelName,
		MaxInFlight: cfg.NSQ["activity_view_stream"].MaxInFlight,
		NSQLookupd:  cfg.NSQ["activity_view_stream"].NSQLookupd,
		Topic:       cfg.NSQ["activity_view_stream"].Topic,
		Worker:      cfg.NSQ["activity_view_stream"].Worker,
	})
	if err != nil {
		log.Println("error initializing NSQ listener")
		return
	}

	delivery.Run(context.Background())

	done := make(chan struct{})
	<-done
}
