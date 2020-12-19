package recommendation

import (
	mProduct "github.com/nafiar/tkpd-recom-engine/internal/model/product"
)

// UseCase interface for recommendation product
type UseCase interface {
	// GetRecommendationData get similar product based on param
	GetRecommendationData(param Param) (data []mProduct.Data, err error)
}
