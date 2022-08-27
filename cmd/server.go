package main

import (
	"github.com/0xmlx/contacts-app-backend/api/router"
	"log"
	"net/http"
	"os"
)

// Run starts a new server.
func run() error {
	var port = os.Getenv("PORT")
	log.Println("starting a new server at port " + port + "...")

	r := router.NEW()

	if err := http.ListenAndServe(":"+port, r); err != nil {
		return err
	}

	return nil
}
