package controllers

import (
	"api/src/models"
	user_repository "api/src/repositories/user"
	"api/src/responses"
	"encoding/json"
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

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		responses.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare(); err != nil {
		responses.ResponseError(w, http.StatusBadRequest, err)
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
	user.ID = uint64(result)
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
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		responses.ResponseError(w, http.StatusBadRequest, err)
		return
	}
	repository, err := user_repository.NewRepository()
	if err != nil {
		responses.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	result, err := repository.Get(uint(id))
	if err != nil {
		responses.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	responses.ResponseHandler(w, http.StatusOK, responses.Response{Data: result})
}
func Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update User"))
}
func DeleteById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete User"))
}
