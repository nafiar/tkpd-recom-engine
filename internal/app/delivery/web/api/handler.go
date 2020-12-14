package api

import (
	"net/http"
)

func (a *apiDelivery) Serve() {
	http.HandleFunc("/userinfo", a.getUserInfo)
	http.ListenAndServe(":9000", nil)
}
