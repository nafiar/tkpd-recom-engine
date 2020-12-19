package main

import (
	"context"
	"fmt"
	"log"

	"github.com/nafiar/tkpd-recom-engine/common/config"
	"github.com/nafiar/tkpd-recom-engine/common/redis"
	recentview_delivery_nsq "github.com/nafiar/tkpd-recom-engine/internal/background/delivery/view/nsq"
	recentview_repository_redis "github.com/nafiar/tkpd-recom-engine/internal/background/repository/recentView/redis"
	view_usecase_useractivity "github.com/nafiar/tkpd-recom-engine/internal/background/usecase/useractivity/view"
)

func main() {
	log.Println("init background")

	cfg := config.GetConfig()

	err := config.InitConfig()
	if err != nil {
		err = fmt.Errorf("Could not init config. Err: %+v", err)
		log.Fatal(err)
	}

	cfg = config.GetConfig()

	log.Println("finish init process")

	redisUserInfo := redis.NewConnection(redis.ConnectionConfig{
		Address:   cfg.Redis["user_data"].Connection,
		MaxActive: cfg.Redis["user_data"].MaxActive,
		MaxIdle:   cfg.Redis["user_data"].MaxIdle,
	})

	recentviewRepository := recentview_repository_redis.New(redisUserInfo)

	recentviewHandler := view_usecase_useractivity.New(recentviewRepository)

	delivery, err := recentview_delivery_nsq.New(recentviewHandler, recentview_delivery_nsq.ListenerConfig{
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
