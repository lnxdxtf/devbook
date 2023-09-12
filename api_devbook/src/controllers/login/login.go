package controllers_login

import (
	"api/src/auth"
	models_user "api/src/models/user"
	login_repository "api/src/repositories/login"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
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

	repository, err := login_repository.NewRepository()
	if err != nil {
		responses.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	userFound, err := repository.GetUserByEmail(user.Email)

	if err != nil {
		responses.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	authOk := authChecker(userFound, user)
	if !authOk {
		responses.ResponseError(w, http.StatusUnauthorized, err)
		return
	}

	token, err := auth.AuthTokenGen(userFound)
	if err != nil {
		responses.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	data := map[string]interface{}{"token": token}
	tokenEXP, err := auth.AuthTokenExtractExpTime(token)
	if err != nil {
		responses.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	data["exp"] = tokenEXP
	data["id"] = userFound.ID

	responses.ResponseHandler(w, http.StatusOK, responses.Response{Data: data})
}

func authChecker(userFound models_user.User, userToCheck models_user.User) bool {
	if err := security.HashCompare(userFound.Pswrd, userToCheck.Pswrd); err != nil {
		return false
	}
	return true
}
