package controllers_user

import (
	"api/src/auth"
	models_user "api/src/models/user"
	user_repository "api/src/repositories/user"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models_user.User
	if err = json.Unmarshal(body, &user); err != nil {
		responses.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	if errs := user.Prepare(models_user.Signup); errs != nil {
		responses.ResponseError(w, http.StatusBadRequest, errs)
		return
	}

	repository, err := user_repository.NewRepository()
	if err != nil {
		responses.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	result, err := repository.Insert(user)
	if err != nil {
		responses.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	user.ID = result
	responses.ResponseHandler(w, http.StatusCreated, responses.Response{Data: user})
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	nameOrNick := "%" + strings.ToLower(r.URL.Query().Get("nick")) + "%"
	repository, err := user_repository.NewRepository()
	if err != nil {
		responses.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	result, err := repository.GetAll(nameOrNick)
	if err != nil {
		responses.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	responses.ResponseHandler(w, http.StatusOK, responses.Response{Data: result})
}

func Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.ResponseError(w, http.StatusBadRequest, err)
		return
	}
	repository, err := user_repository.NewRepository()
	if err != nil {
		responses.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	result, err := repository.GetById(uint(id))
	if err != nil {
		responses.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	responses.ResponseHandler(w, http.StatusOK, responses.Response{Data: result})
}

func Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.ResponseError(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	userIDToken, err := auth.AuthTokenExtractDataUser(r)
	if err != nil {
		responses.ResponseError(w, http.StatusUnauthorized, err)
		return
	}

	if userIDToken != uint(id) {
		responses.ResponseError(w, http.StatusForbidden, errors.New("cannot update another user"))
		return
	}

	var user models_user.User
	if err = json.Unmarshal(body, &user); err != nil {
		responses.ResponseError(w, http.StatusBadRequest, err)
		return
	}
	if errs := user.Prepare(models_user.Update); errs != nil {
		responses.ResponseError(w, http.StatusBadRequest, errs)
		return
	}
	repository, err := user_repository.NewRepository()
	if err != nil {
		responses.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	if err = repository.Update(uint(id), user); err != nil {
		responses.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	responses.ResponseHandler(w, http.StatusNoContent, nil)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	userIDToken, err := auth.AuthTokenExtractDataUser(r)
	if err != nil {
		responses.ResponseError(w, http.StatusUnauthorized, err)
		return
	}

	if userIDToken != uint(id) {
		responses.ResponseError(w, http.StatusForbidden, errors.New("cannot delete another user"))
		return
	}

	repository, err := user_repository.NewRepository()
	if err != nil {
		responses.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	if err = repository.Delete(uint(id)); err != nil {
		responses.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

}

func Follow(w http.ResponseWriter, r *http.Request) {
	// User that will follow - Follower
	followerID, err := auth.AuthTokenExtractDataUser(r)
	if err != nil {
		responses.ResponseError(w, http.StatusUnauthorized, err)
		return
	}
	params := mux.Vars(r)
	// User that will be followed - User
	userID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.ResponseError(w, http.StatusBadRequest, err)
		return
	}
	if followerID == uint(userID) {
		responses.ResponseError(w, http.StatusForbidden, errors.New("cannot follow yourself"))
		return
	}
	repository, err := user_repository.NewRepository()
	if err != nil {
		responses.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	if err = repository.FollowUser(uint(userID), uint(followerID)); err != nil {
		responses.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	responses.ResponseHandler(w, http.StatusNoContent, nil)
}

func Unfollow(w http.ResponseWriter, r *http.Request) {
	// User that will follow - Follower
	followerID, err := auth.AuthTokenExtractDataUser(r)
	if err != nil {
		responses.ResponseError(w, http.StatusUnauthorized, err)
		return
	}
	params := mux.Vars(r)
	// User that will be followed - User
	userID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.ResponseError(w, http.StatusBadRequest, err)
		return
	}
	if followerID == uint(userID) {
		responses.ResponseError(w, http.StatusForbidden, errors.New("cannot follow yourself"))
		return
	}
	repository, err := user_repository.NewRepository()
	if err != nil {
		responses.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	if err = repository.UnfollowUser(uint(userID), uint(followerID)); err != nil {
		responses.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	responses.ResponseHandler(w, http.StatusNoContent, nil)

}

func GetUserFollowers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	repository, err := user_repository.NewRepository()
	if err != nil {
		responses.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	result, err := repository.GetUserFollowers(uint(userID))
	if err != nil {
		responses.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	responses.ResponseHandler(w, http.StatusOK, responses.Response{Data: result})
}

func GetUserFollowing(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	repository, err := user_repository.NewRepository()
	if err != nil {
		responses.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	result, err := repository.GetUserFollowing(uint(userID))
	if err != nil {
		responses.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	responses.ResponseHandler(w, http.StatusOK, responses.Response{Data: result})
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.ResponseError(w, http.StatusBadRequest, err)
		return
	}
	userIDToken, err := auth.AuthTokenExtractDataUser(r)
	if err != nil {
		responses.ResponseError(w, http.StatusUnauthorized, err)
		return
	}

	if userIDToken != uint(userID) {
		responses.ResponseError(w, http.StatusForbidden, errors.New("cannot update password from another user"))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	var pswrdUpdtData models_user.PswrdUpdate
	if err = json.Unmarshal(body, &pswrdUpdtData); err != nil {
		responses.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	repository, err := user_repository.NewRepository()
	if err != nil {
		responses.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	pswrdUserSaved, err := repository.GetPswrd(uint(userID))
	if err != nil {
		responses.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.HashCompare(pswrdUserSaved, pswrdUpdtData.Current); err != nil {
		responses.ResponseError(w, http.StatusUnauthorized, errors.New("invalid password"))
		return
	}

	newHashPswrd, err := security.HashGen(pswrdUpdtData.New)
	if err != nil {
		responses.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	if err = repository.UpdatePswrd(uint(userID), string(newHashPswrd)); err != nil {
		responses.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	responses.ResponseHandler(w, http.StatusNoContent, nil)

}
