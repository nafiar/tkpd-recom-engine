package http

import "github.com/nafiar/tkpd-recom-engine/internal/serving/repository/similarproduct"

type httpSimilarProduct struct {
}

func New() similarproduct.Repository {
	return &httpSimilarProduct{}
}
