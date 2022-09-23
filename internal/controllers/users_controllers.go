package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/inuoshios/keepinfo/internal/models"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.User

	json.NewDecoder(r.Body).Decode(&user)
	log.Println(user)
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login"))
}

func SignOut(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Logout"))
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
