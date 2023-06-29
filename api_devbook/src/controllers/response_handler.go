package controllers

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Status any         `json:"status"`
	Error  interface{} `json:"error"`
}

type Response struct {
	Status any         `json:"status"`
	Data   interface{} `json:"data"`
}

func ResponseJsonWriter(w http.ResponseWriter, code int, output interface{}) {
	response, _ := json.Marshal(output)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
