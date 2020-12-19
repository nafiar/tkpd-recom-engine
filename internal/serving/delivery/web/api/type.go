package api

import (
	"net/http"

	"github.com/mholt/binding"
	mProduct "github.com/nafiar/tkpd-recom-engine/internal/model/product"
)

type (
	Param struct {
		UserID    int
		ProductID int
	}
)

func (p *Param) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&p.UserID:    "user_id",
		&p.ProductID: "product_id",
	}
}

type (
	Response struct {
		Meta Meta            `json:"meta"`
		Data []mProduct.Data `json:"data"`
	}

	Meta struct {
		ProcessTime string `json:"process_time"`
	}

	Error struct {
		Status     int    `json:"status"`
		ErrMessage string `json:"error_message"`
	}
)
