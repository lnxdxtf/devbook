package controllers

import "net/http"

func Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create User"))
}
func GetAll(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get all Users"))
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
