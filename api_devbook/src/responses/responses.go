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

func ResponseError(w http.ResponseWriter, code int, err error) {
	ResponseHandler(w, code, ErrorResponse{Error: err.Error()})
}
