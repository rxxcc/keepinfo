package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/inuoshios/keepinfo/internal/config"
	"github.com/inuoshios/keepinfo/internal/database"
)

var (
	app config.Config
)

// Run starts a new server.
func Run() (*database.DB, error) {
	// adding custom logs
	app.InfoLog = log.New(os.Stdout, "[STATUS]"+" ", log.Ldate|log.Ltime)
	app.ErrorLog = log.New(os.Stdout, "[ERROR]"+" ", log.Ldate|log.Ltime|log.Lshortfile)

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

	app.InfoLog.Println("connected to the database successfully")

	repo := NewRepository(&app, dbConn)
	NewHandlers(repo)

	// starting up a server.

	// if err := http.ListenAndServe(":"+port, r); err != nil {
	// 	app.ErrorLog.Fatalf("server error - %v", err)
	// 	return nil, err
	// }
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: NEW(),
	}

	shutDownErr := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)

		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		s := <-quit

		app.InfoLog.Printf("service interrupt received %s", s.String())

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer func() {
			cancel()
		}()

		shutDownErr <- srv.Shutdown(ctx)

	}()

	app.InfoLog.Println("starting a new server at port " + port)

	err = srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return nil, fmt.Errorf("http server shutdown error %w", err)
	}

	if err = <-shutDownErr; err != nil {
		return nil, err
	}

	return dbConn, nil
}
