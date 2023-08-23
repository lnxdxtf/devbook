package routes

import (
	controllers_user "api/src/controllers/user"
	"net/http"
)

var UserRoutes = []Route{
	{
		Uri:    "/users",
		Method: http.MethodGet,
		Fn:     controllers_user.GetAll,
		Auth:   true,
	},
	{
		Uri:    "/users/{id}",
		Method: http.MethodGet,
		Fn:     controllers_user.Get,
		Auth:   true,
	},
	{
		Uri:    "/users",
		Method: http.MethodPost,
		Fn:     controllers_user.Create,
		Auth:   false,
	},
	{
		Uri:    "/users/{id}",
		Method: http.MethodPut,
		Fn:     controllers_user.Update,
		Auth:   true,
	},
	{
		Uri:    "/users/{id}",
		Method: http.MethodDelete,
		Fn:     controllers_user.DeleteById,
		Auth:   true,
	},
	{
		Uri:    "/users/{id}/follow",
		Method: http.MethodPost,
		Fn:     controllers_user.Follow,
		Auth:   true,
	},
	{
		Uri:    "/users/{id}/unfollow",
		Method: http.MethodPost,
		Fn:     controllers_user.Unfollow,
		Auth:   true,
	},
	{
		Uri:    "/users/{id}/followers",
		Method: http.MethodGet,
		Fn:     controllers_user.GetUserFollowers,
		Auth:   true,
	},
	{
		Uri:    "/users/{id}/following",
		Method: http.MethodGet,
		Fn:     controllers_user.GetUserFollowing,
		Auth:   true,
	},
	{
		Uri:    "/users/{id}/update-password",
		Method: http.MethodPost,
		Fn:     controllers_user.UpdatePassword,
		Auth:   true,
	},
}
