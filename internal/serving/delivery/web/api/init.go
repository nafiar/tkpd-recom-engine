package api

import (
	"github.com/nafiar/tkpd-recom-engine/internal/serving/delivery/web"
	"github.com/nafiar/tkpd-recom-engine/internal/serving/usecase/recommendation"
)

// apiDelivery handler for HTTP api web server
type apiDelivery struct {
	recommendationUseCase recommendation.UseCase
}

// New create new object for HTTP API web server
func New(recomUseCase recommendation.UseCase) web.Delivery {
	api := &apiDelivery{
		recommendationUseCase: recomUseCase,
	}

	return api
}
