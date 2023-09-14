package controllers_post

import (
	"api/src/auth"
	models_post "api/src/models/post"
	post_repository "api/src/repositories/post"
	"api/src/responses"
	aws_devbook_s3 "api/src/services/aws"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

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

	result, err := repository.Insert(post)
	if err != nil {
		responses.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	if post.ImageBase64 != "" {
		base64Data := post.ImageBase64
		if strings.Contains(base64Data, ",") {
			base64Data = strings.Split(base64Data, ",")[1]
		}
		decodeBase64Img, err := base64.StdEncoding.DecodeString(base64Data)
		if err != nil {
			responses.ResponseError(w, http.StatusInternalServerError, err)
			return
		}
		s3Service, err := aws_devbook_s3.NewS3Service()
		if err != nil {
			responses.ResponseError(w, http.StatusInternalServerError, err)
			return
		}
		upload := aws_devbook_s3.UploadInput{
			File_name: fmt.Sprintf("user/post/imgs/post_%d.png", result),
			File_type: "image/png",
			File_body: decodeBase64Img,
		}
		if err = s3Service.Upload(upload); err != nil {
			responses.ResponseError(w, http.StatusInternalServerError, err)
			return
		}

	}

	responses.ResponseHandler(w, http.StatusCreated, responses.Response{Data: result})

}

// Get posts from users that the user follows
func GetAll(w http.ResponseWriter, r *http.Request) {
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

func GetRandom(w http.ResponseWriter, r *http.Request) {
	repository, err := post_repository.NewRepository()
	if err != nil {
		responses.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	post, err := repository.GetRandom()
	if err != nil {
		responses.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	
	responses.ResponseHandler(w, http.StatusOK, responses.Response{Data: post})
}

func Get(w http.ResponseWriter, r *http.Request) {
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

func Update(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.AuthTokenExtractDataUser(r)
	if err != nil {
		responses.ResponseError(w, http.StatusUnauthorized, err)
		return
	}

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

	postFound, err := repository.GetById(uint(postID))
	if err != nil {
		responses.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	if postFound.AuthorID != userID {
		responses.ResponseError(w, http.StatusForbidden, errors.New("you can only update your own posts"))
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

	if errs := post.Prepare(); err != nil {
		responses.ResponseError(w, http.StatusBadRequest, errs)
		return
	}

	if err = repository.Update(uint(postID), post); err != nil {
		responses.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	responses.ResponseHandler(w, http.StatusNoContent, nil)

}

func Delete(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.AuthTokenExtractDataUser(r)
	if err != nil {
		responses.ResponseError(w, http.StatusUnauthorized, err)
		return
	}

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

	postFound, err := repository.GetById(uint(postID))
	if err != nil {
		responses.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	if postFound.AuthorID != userID {
		responses.ResponseError(w, http.StatusForbidden, errors.New("you can only delete your own posts"))
		return
	}

	if err = repository.Delete(uint(postID)); err != nil {
		responses.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	responses.ResponseHandler(w, http.StatusNoContent, nil)
}

func UserPosts(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	repository, err := post_repository.NewRepository()
	if err != nil {
		responses.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	posts, err := repository.GetUserPosts(uint(userID))
	if err != nil {
		responses.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	responses.ResponseHandler(w, http.StatusOK, responses.Response{Data: posts})
}

func Like(w http.ResponseWriter, r *http.Request) {
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

	if err = repository.Like(uint(postID)); err != nil {
		responses.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	responses.ResponseHandler(w, http.StatusNoContent, nil)
}

func Unlike(w http.ResponseWriter, r *http.Request) {
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

	if err = repository.Unlike(uint(postID)); err != nil {
		responses.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	responses.ResponseHandler(w, http.StatusNoContent, nil)
}
