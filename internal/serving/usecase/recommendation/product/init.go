package product

import (
	"github.com/nafiar/tkpd-recom-engine/internal/serving/repository/product"
	"github.com/nafiar/tkpd-recom-engine/internal/serving/repository/recentview"
	"github.com/nafiar/tkpd-recom-engine/internal/serving/repository/similarproduct"
	"github.com/nafiar/tkpd-recom-engine/internal/serving/usecase/recommendation"
)

// productRecommendation handler for HTTP api web server
type productRecommendation struct {
	recentViewRepo     recentview.Repository
	similarProductRepo similarproduct.Repository
	productRepo        product.Repository
}

func New(recentViewRepo recentview.Repository, similarProductRepo similarproduct.Repository, productRepo product.Repository) recommendation.UseCase {
	return &productRecommendation{
		recentViewRepo:     recentViewRepo,
		similarProductRepo: similarProductRepo,
		productRepo:        productRepo,
	}
}
