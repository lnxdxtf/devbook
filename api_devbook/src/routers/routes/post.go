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
		Fn:     controllers_post.GetPosts,
		Auth:   true,
	},
	{
		Uri:    "/posts/{id}",
		Method: http.MethodGet,
		Fn:     controllers_post.GetPostsById,
		Auth:   true,
	},
	{
		Uri:    "/posts/{id}",
		Method: http.MethodDelete,
		Fn:     controllers_post.DeleteById,
		Auth:   true,
	},
}
