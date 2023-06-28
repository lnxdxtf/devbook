package routes

import "net/http"

type Route struct {
	Uri    string
	Method string
	Fn     func(http.ResponseWriter, *http.Request)
	Auth   bool
}
