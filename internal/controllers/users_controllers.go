package controllers

import (
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create new user"))
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login"))
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
