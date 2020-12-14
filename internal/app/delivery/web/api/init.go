package api

import (
	"github.com/nafiar/tkpd-recom-engine/internal/app/delivery/web"
	"github.com/nafiar/tkpd-recom-engine/internal/app/usecase/userdata"
)

// apiDelivery handler for HTTP api web server
type apiDelivery struct {
	userDataUseCase userdata.UseCase
}

// New create new object for HTTP API web server
func New(userData userdata.UseCase) web.Delivery {
	api := &apiDelivery{
		userDataUseCase: userData,
	}

	return api
}
