package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/inuoshios/keepinfo/internal/config"
	"github.com/inuoshios/keepinfo/internal/db"
	"github.com/inuoshios/keepinfo/internal/router"
)

var (
	port           = os.Getenv("PORT")
	host           = os.Getenv("HOST")
	portNumber     = os.Getenv("DBPORT")
	user           = os.Getenv("USER")
	password       = os.Getenv("PASSWORD")
	dbName         = os.Getenv("DBNAME")
	app            config.Config
	sessionManager *scs.SessionManager
)

// Run starts a new server.
func Run() (*db.DB, error) {
	// This initializes a new session manager.
	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	sessionManager.Cookie.Secure = false

	app.Session = sessionManager

	// connecting your application to the database.
	connInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, portNumber, user, password, dbName)
	dbConn, err := db.ConnectSQL(connInfo)
	if err != nil {
		log.Fatal("cannot connect to the database...")
	}

	log.Println("connected to the database...")

	// starting up a server.
	log.Println("starting a new server at port " + port + "...")
	r := router.NEW()

	if err := http.ListenAndServe(":"+port, r); err != nil {
		return nil, err
	}

	return dbConn, nil
}
