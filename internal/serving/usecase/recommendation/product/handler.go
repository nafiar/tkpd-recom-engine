package product

import (
	"fmt"

	mProduct "github.com/nafiar/tkpd-recom-engine/internal/model/product"
	mRecom "github.com/nafiar/tkpd-recom-engine/internal/serving/usecase/recommendation"
)

// GetRecommendationData generate recommendation products based on param
// first it will retrieve user activity view data
// then, it will enrich it with product similarity model
// eventually it will get additional product detail data before returned as response
func (p *productRecommendation) GetRecommendationData(param mRecom.Param) (data []mProduct.Data, err error) {
	// paranoid check for upstream data
	if err = p.safeguard(); err != nil {
		err = fmt.Errorf("failed to get data, " + err.Error())
		return
	}

	// get list of user's recent view products
	recView, err := p.recentViewRepo.GetRecentView(param.UserID)
	if err != nil {
		return
	}

	// get similar products based on recent view products
	similarProducts, err := p.similarProductRepo.GetSimilarProduct(recView)
	if err != nil {
		return
	}

	// get additional product detail before return it to response
	data, err = p.productRepo.GetProductByIDs(similarProducts)
	if err != nil {
		return
	}

	return
}
