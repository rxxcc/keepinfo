package server

import (
	"fmt"
	"github.com/0xmlx/contacts-app-backend/internal/db"
	"github.com/0xmlx/contacts-app-backend/internal/router"
	"log"
	"net/http"
	"os"
)

var (
	port       = os.Getenv("PORT")
	host       = os.Getenv("HOST")
	portNumber = os.Getenv("DBPORT")
	user       = os.Getenv("USER")
	password   = os.Getenv("PASSWORD")
	dbName     = os.Getenv("DBNAME")
)

// Run starts a new server.
func Run() (*db.DB, error) {
	log.Println("starting a new server at port " + port + "...")

	connInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, portNumber, user, password, dbName)
	dbConn, err := db.ConnectSQL(connInfo)
	if err != nil {
		log.Fatal("cannot connect to the database...")
	}

	log.Println("connected to the database...")

	r := router.NEW()

	if err := http.ListenAndServe(":"+port, r); err != nil {
		return nil, err
	}

	return dbConn, nil
}
