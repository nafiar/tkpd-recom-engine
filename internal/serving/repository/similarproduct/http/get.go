package http

import (
	"encoding/json"
	"net/http"
)

// GetSimilarProduct get similar product from model http API
// intentionally remove it's implementation
// reference : Slide HTTP `Create a new request` & `Get the response and how to parse`
func (h *httpSimilarProduct) GetSimilarProduct(id []int) (result []int, err error) {
	// Create http request
	req, err := http.NewRequest("GET", "https://private-6c3ef8-techworkshop.apiary-mock.com/ge", nil)
	if err != nil {
		return
	}

	// Execute http request
	res, err := h.client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	// Parse
	var response Response
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	// Append to result
	for _, v := range response.RelatedProducts {
		result = append(result, int(v.ProductID))
	}
	return
}
