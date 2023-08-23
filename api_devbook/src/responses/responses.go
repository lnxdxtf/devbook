package responses

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error interface{} `json:"error"`
}

type Response struct {
	Data interface{} `json:"data"`
}

func ResponseHandler(w http.ResponseWriter, code int, output interface{}) {
	response, _ := json.Marshal(output)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func ResponseError(w http.ResponseWriter, code int, err interface{}) {
	switch e := err.(type) {
	case error:
		ResponseHandler(w, code, ErrorResponse{Error: e.Error()})
	case []error:
		var errMessages []string
		for _, err := range e {
			errMessages = append(errMessages, err.Error())
		}
		ResponseHandler(w, code, ErrorResponse{Error: errMessages})
	default:
		ResponseHandler(w, code, ErrorResponse{Error: "an unexpected error occurred"})
	}
}
