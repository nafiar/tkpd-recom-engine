package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mholt/binding"
)

func (a *apiDelivery) getUserInfo(w http.ResponseWriter, req *http.Request) {
	var param Param
	var response Response
	var err error

	start := time.Now()

	err = binding.Bind(req, &param)
	if err = sanitizeBindError(err); err != nil {
		renderError(w, err, http.StatusBadRequest)
		return
	}

	data, err := a.userDataUseCase.GetData(param.UserID)
	if err != nil {
		renderError(w, err, http.StatusInternalServerError)
		return
	}

	response.Data = data
	response.Meta.ProcessTime = fmt.Sprintf("%.2fms", time.Since(start).Seconds()*1000)

	render(w, response)
}
