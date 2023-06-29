package routes

import (
	"api/src/controllers"
	"net/http"
)

var UserRoutes = []Route{
	{
		Uri:    "/users",
		Method: http.MethodGet,
		Fn:     controllers.GetAll,
		Auth:   false,
	},
	{
		Uri:    "/users/{id}",
		Method: http.MethodGet,
		Fn:     controllers.GetById,
		Auth:   false,
	},
	{
		Uri:    "/users",
		Method: http.MethodPost,
		Fn:     controllers.Create,
		Auth:   false,
	},
	{
		Uri:    "/users/{id}",
		Method: http.MethodPut,
		Fn:     controllers.UpdateById,
		Auth:   false,
	},
	{
		Uri:    "/users/{id}",
		Method: http.MethodDelete,
		Fn:     controllers.DeleteById,
		Auth:   false,
	},
}
