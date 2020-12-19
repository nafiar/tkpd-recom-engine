package main

import (
	"fmt"
	"log"

	"github.com/nafiar/tkpd-recom-engine/common/config"
	postgreConn "github.com/nafiar/tkpd-recom-engine/common/postgres"
	"github.com/nafiar/tkpd-recom-engine/common/redis"
	apiDelivery "github.com/nafiar/tkpd-recom-engine/internal/serving/delivery/web/api"
	postgreProductRepo "github.com/nafiar/tkpd-recom-engine/internal/serving/repository/product/postgres"
	redisRecentViewRepo "github.com/nafiar/tkpd-recom-engine/internal/serving/repository/recentview/redis"
	httpSimilarProductRepo "github.com/nafiar/tkpd-recom-engine/internal/serving/repository/similarproduct/http"
	productRecommendation "github.com/nafiar/tkpd-recom-engine/internal/serving/usecase/recommendation/product"
)

func main() {
	cfg := config.GetConfig()

	err := config.InitConfig()
	if err != nil {
		err = fmt.Errorf("Could not init config. Err: %+v", err)
		log.Fatal(err)
	}

	cfg = config.GetConfig()

	// Postgre create new connection
	dbPostgre, err := postgreConn.NewConnection(cfg.Postgre["product_recommendation"].ConnString)
	if err != nil {
		log.Fatal(err)
	}
	productRepo := postgreProductRepo.New(dbPostgre)

	redisRecentView := redis.NewConnection(redis.ConnectionConfig{
		Address:   cfg.Redis["recent_view"].Connection,
		MaxActive: cfg.Redis["recent_view"].MaxActive,
		MaxIdle:   cfg.Redis["recent_view"].MaxIdle,
	})

	recentViewRepo := redisRecentViewRepo.New(redisRecentView)
	similarProductRepo := httpSimilarProductRepo.New()

	productRecommendationUC := productRecommendation.New(recentViewRepo, similarProductRepo, productRepo)

	api := apiDelivery.New(productRecommendationUC)
	api.Serve()
}
