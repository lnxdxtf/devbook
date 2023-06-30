package controllers

import (
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		ResponseError(w, http.StatusBadRequest, err)
		return
	}

	repository, err := repositories.NewRepository()
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	result, err := repository.Insert("INSERT INTO devbook.users (name, nick, email, password) VALUES (?, ?, ?, ?)", user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	user.ID = uint64(result)
	ResponseHandler(w, http.StatusCreated, Response{Data: user})
}

func GetAll(w http.ResponseWriter, r *http.Request) {
}
func GetById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get User By Id"))
}
func UpdateById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update User"))
}
func DeleteById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete User"))
}
