package routes

import (
	controllers_post "api/src/controllers/posts"
	"net/http"
)

var PostRoutes = []Route{
	{
		Uri:    "/posts",
		Method: http.MethodPost,
		Fn:     controllers_post.Create,
		Auth:   true,
	},
	{
		Uri:    "/posts",
		Method: http.MethodGet,
		Fn:     controllers_post.GetAll,
		Auth:   true,
	},
	{
		Uri:    "/posts/{id}",
		Method: http.MethodGet,
		Fn:     controllers_post.Get,
		Auth:   true,
	},
	{
		Uri:    "/posts/{id}",
		Method: http.MethodPut,
		Fn:     controllers_post.Update,
		Auth:   true,
	},
	{
		Uri:    "/posts/{id}",
		Method: http.MethodDelete,
		Fn:     controllers_post.Delete,
		Auth:   true,
	},
	{
		Uri:    "/users/{id}/posts",
		Method: http.MethodGet,
		Fn:     controllers_post.UserPosts,
		Auth:   true,
	},
	{
		Uri:    "/posts/{id}/like",
		Method: http.MethodPost,
		Fn:     controllers_post.Like,
		Auth:   true,
	},
	{
		Uri:    "/posts/{id}/unlike",
		Method: http.MethodPost,
		Fn:     controllers_post.Unlike,
		Auth:   true,
	},
}
