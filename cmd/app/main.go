package main

import (
	"fmt"
	"log"

	"github.com/nafiar/tkpd-recom-engine/common/config"
	postgreConn "github.com/nafiar/tkpd-recom-engine/common/postgres"
	"github.com/nafiar/tkpd-recom-engine/common/redis"
	apiDelivery "github.com/nafiar/tkpd-recom-engine/internal/app/delivery/web/api"
	redisUserInfoRepo "github.com/nafiar/tkpd-recom-engine/internal/app/repository/userinfo/redis"
	postgreProductRepo "github.com/nafiar/tkpd-recom-engine/internal/app/repository/product/postgres"
	userDataInfoUC "github.com/nafiar/tkpd-recom-engine/internal/app/usecase/userdata/info"
	"github.com/nafiar/tkpd-recom-engine/internal/model/product"
)

func main() {
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

	dbPostgre,err := postgreConn.NewConnection(cfg.Postgre["product_recommendation"].ConnString)
	if err != nil {
		log.Fatal(err)
	}
	productRepo := postgreProductRepo.New(dbPostgre)

	productSample, _ := productRepo.GetProduct([]product.Param{})
	log.Println(productSample)

	//productRepo.GetProduct(nil)
	userInfoRepo := redisUserInfoRepo.New(redisUserInfo)
	userDataInfoUC := userDataInfoUC.New(userInfoRepo)

	// redisRecentView := redis.NewConnection(redis.ConnectionConfig{
	// 	Address:   cfg.Redis["recent_view"].Connection,
	// 	MaxActive: cfg.Redis["recent_view"].MaxActive,
	// 	MaxIdle:   cfg.Redis["recent_view"].MaxIdle,
	// })
	// recentViewRepo := redisRecentViewRepo.New(redisRecentView)
	// recentViewUC := recentViewUC.New(recentViewRepo)

	api := apiDelivery.New(userDataInfoUC)
	api.Serve()
}
