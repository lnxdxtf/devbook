package routes

import (
	controllers_login "api/src/controllers/login"
	"net/http"
)

var LoginRoute = Route{
	Uri:    "/login",
	Method: http.MethodPost,
	Fn:     controllers_login.Login,
	Auth:   false,
}
