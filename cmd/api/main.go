package main

import (
	"github.com/0xmlx/contacts-app-backend/internal/server"
	"log"
)

// runs the server
func main() {
	db, err := server.Run()
	if err != nil {
		log.Fatal(err)
	}

	defer db.SQL.Close()
}
