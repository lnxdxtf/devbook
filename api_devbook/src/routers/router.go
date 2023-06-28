package routers

import (
	"api/src/routers/routes"

	"github.com/gorilla/mux"
)

func Create() *mux.Router {
	return config(mux.NewRouter())
}

func config(r *mux.Router) *mux.Router {
	for _, route := range routes.UserRoutes {
		r.HandleFunc(route.Uri, route.Fn).Methods(route.Method)
	}
	return r
} 
