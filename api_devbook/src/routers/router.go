package routers

import (
	"api/src/middlewares"
	"api/src/routers/routes"

	"github.com/gorilla/mux"
)

func Create() *mux.Router {
	return config(mux.NewRouter())
}

func config(r *mux.Router) *mux.Router {
	routesPckg := routes.UserRoutes
	routesPckg = append(routesPckg, routes.LoginRoute)
	routesPckg = append(routesPckg, routes.PostRoutes...) // ... is a spread operator | this will add one by one in the slice and not add the slice itself

	for _, route := range routesPckg {
		if route.Auth {
			r.HandleFunc(route.Uri, middlewares.LoggerMiddleware(middlewares.AuthMiddleware(route.Fn))).Methods(route.Method)
			continue
		} else {
			r.HandleFunc(route.Uri, middlewares.LoggerMiddleware(route.Fn)).Methods(route.Method)
		}
	}
	return r
}
