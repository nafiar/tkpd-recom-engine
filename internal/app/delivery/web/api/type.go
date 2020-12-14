package api

import (
	"net/http"

	"github.com/mholt/binding"
	mUser "github.com/nafiar/tkpd-recom-engine/internal/model/user"
)

type (
	Param struct {
		UserID string
	}
)

func (p *Param) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&p.UserID: "user_id",
	}
}

type (
	Response struct {
		Meta Meta       `json:"meta"`
		Data mUser.Data `json:"data"`
	}

	Meta struct {
		ProcessTime string `json:"process_time"`
	}

	Error struct {
		Status     int    `json:"status"`
		ErrMessage string `json:"error_message"`
	}
)
