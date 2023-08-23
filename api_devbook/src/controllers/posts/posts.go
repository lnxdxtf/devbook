package controllers_post

import (
	"api/src/auth"
	models_post "api/src/models/post"
	post_repository "api/src/repositories/post"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Create(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.AuthTokenExtractDataUser(r)
	if err != nil {
		responses.ResponseError(w, http.StatusUnauthorized, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	var post models_post.Post
	if err = json.Unmarshal(body, &post); err != nil {
		responses.ResponseError(w, http.StatusBadRequest, err)
		return
	}
	post.AuthorID = userID
	if errs := post.Prepare(); errs != nil {
		responses.ResponseError(w, http.StatusBadRequest, errs)
		return
	}

	repository, err := post_repository.NewRepository()
	if err != nil {
		responses.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	postID, err := repository.Insert(post)
	if err != nil {
		responses.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	responses.ResponseHandler(w, http.StatusCreated, responses.Response{Data: postID})

}

// Get posts from users that the user follows
func GetPosts(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.AuthTokenExtractDataUser(r)
	if err != nil {
		responses.ResponseError(w, http.StatusUnauthorized, err)
		return
	}
	repository, err := post_repository.NewRepository()
	if err != nil {
		responses.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	posts, err := repository.GetAll(userID)
	if err != nil {
		responses.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	responses.ResponseHandler(w, http.StatusOK, responses.Response{Data: posts})
}

func GetPostsById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	repository, err := post_repository.NewRepository()
	if err != nil {
		responses.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	post, err := repository.GetById(uint(postID))
	if err != nil {
		responses.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	responses.ResponseHandler(w, http.StatusOK, responses.Response{Data: post})
}

func DeleteById(w http.ResponseWriter, r *http.Request) {

}
