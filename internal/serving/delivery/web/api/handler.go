package api

import "net/http"

func (a *apiDelivery) Serve() {
	http.HandleFunc("/recommendation/product", a.getRecommendationProduct)
	http.ListenAndServe(":9000", nil)
}
