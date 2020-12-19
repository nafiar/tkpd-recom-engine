package http

import (
	"net/http"

	"github.com/nafiar/tkpd-recom-engine/internal/serving/repository/similarproduct"
)

type httpSimilarProduct struct {
	client http.Client
}

func New() similarproduct.Repository {
	return &httpSimilarProduct{
		client: http.Client{},
	}
}
