package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/inuoshios/keepinfo/internal/config"
	"github.com/inuoshios/keepinfo/internal/database"
	"github.com/inuoshios/keepinfo/internal/handlers"
	"github.com/inuoshios/keepinfo/internal/router"
)

var app config.Config

// Run starts a new server.
func Run() (*database.DB, error) {
	// adding custom logs
	app.InfoLog = log.New(os.Stdout, "STATUS:\t", log.Ldate|log.Ltime)
	app.ErrorLog = log.New(os.Stdout, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)

	port := os.Getenv("PORT")
	host := os.Getenv("HOST")
	portNumber := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbName := os.Getenv("DBNAME")

	// connecting your application to the database.
	connInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, portNumber, user, password, dbName)
	dbConn, err := database.ConnectSQL(connInfo)
	if err != nil {
		app.InfoLog.Fatalf("error connecting to the database %s", err)
	}

	app.InfoLog.Println("connected to the database successfully...")

	// starting up a server.
	app.InfoLog.Println("starting a new server at port " + port + "...")
	r := router.NEW()

	if err := http.ListenAndServe(":"+port, r); err != nil {
		app.ErrorLog.Fatalf("server error %v", err)
		return nil, err
	}

	repo := handlers.NewRepository(&app, dbConn)
	handlers.NewHandlers(repo)

	return dbConn, nil
}
