package api

import (
	"errors"
	"net/http"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

func renderError(w http.ResponseWriter, err error, status int) {
	response := Error{
		Status:     status,
		ErrMessage: err.Error(),
	}
	render(w, response)
}

func render(w http.ResponseWriter, response interface{}) {
	jsonResponse, _ := jsoniter.ConfigFastest.Marshal(response)
	w.Write(jsonResponse)
}

// sanitizeBindError remove unused parameter binding error
func sanitizeBindError(errs error) (result error) {
	if errs == nil {
		return
	}

	result = errs

	var errorMessages []string
	messages := strings.Split(errs.Error(), ",")
	for _, m := range messages {
		trim := strings.Trim(m, " ")
		// ignore empty string
		if trim == "" {
			continue
		}
		// ignore parse error
		if strings.Contains(trim, "strconv.Parse") {
			continue
		}
		errorMessages = append(errorMessages, trim)
	}
	if len(errorMessages) == 0 {
		return nil
	}

	return errors.New(strings.Join(errorMessages, ", "))
}
