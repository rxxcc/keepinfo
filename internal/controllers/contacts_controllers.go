package controllers

import "net/http"

func GetContacts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List of contacts"))
}

func CreateContact(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create new contact"))
}

func GetContact(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get a single contact"))
}

func UpdateContact(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update contact"))
}

func DeleteContact(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete contact"))
}
