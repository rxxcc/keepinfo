package controllers

import (
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List of users"))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create new user"))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get a single user"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update User"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete User"))
}
