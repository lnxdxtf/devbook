package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ResponseJsonWriter(w, http.StatusBadRequest, ErrorResponse{Status: http.StatusBadRequest, Error: "Cannot process the request body :("})
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		ResponseJsonWriter(w, http.StatusBadRequest, ErrorResponse{Status: http.StatusBadRequest, Error: "Error converting user data :("})
	}
	mysql_db := database.MySQLDB{}
	db, err := mysql_db.Connect()
	if err != nil {
		ResponseJsonWriter(w, http.StatusInternalServerError, ErrorResponse{Status: http.StatusInternalServerError, Error: "Error connecting to database :("})
	}

	repository := repositories.NewRepository(db)
	query := "INSERT INTO users (name, nick, email, password) VALUES (?, ?, ?, ?)"
	result, err := repository.Insert(query, user.Name, user.Nick, user.Email, user.Password, user.CreatedAt)
	if err != nil {
		ResponseJsonWriter(w, http.StatusInternalServerError, ErrorResponse{Status: http.StatusInternalServerError, Error: "Error inserting user :("})
	}
	type LastInsertId struct {
		Id int64 `json:"id"`
	}
	id, err := result.LastInsertId()
	if err != nil {
		ResponseJsonWriter(w, http.StatusInternalServerError, ErrorResponse{Status: http.StatusInternalServerError, Error: "Error getting last inserted id :(. But the user was inserted successfully"})
	}

	data := LastInsertId{Id: id}
	ResponseJsonWriter(w, http.StatusCreated, data)
}
func GetAll(w http.ResponseWriter, r *http.Request) {
	ResponseJsonWriter(w, http.StatusBadRequest, ErrorResponse{Status: http.StatusBadRequest, Error: "Cannot process the request body :("})
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
