package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mholt/binding"
	mRecom "github.com/nafiar/tkpd-recom-engine/internal/serving/usecase/recommendation"
)

func (a *apiDelivery) getRecommendationProduct(w http.ResponseWriter, req *http.Request) {
	var param Param
	var response Response
	var err error

	start := time.Now()

	err = binding.Bind(req, &param)
	if err = sanitizeBindError(err); err != nil {
		renderError(w, err, http.StatusBadRequest)
		return
	}

	data, err := a.recommendationUseCase.GetRecommendationData(mRecom.Param{UserID: param.UserID, ProductID: param.ProductID})
	if err != nil {
		renderError(w, err, http.StatusInternalServerError)
		return
	}

	response.Data = data
	response.Meta.ProcessTime = fmt.Sprintf("%.2fms", time.Since(start).Seconds()*1000)

	render(w, response)
}
