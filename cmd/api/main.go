package main

import (
	"github.com/inuoshios/keepinfo/internal/database"
	"github.com/joho/godotenv"
)

// runs the server
func main() {
	// load env files.
	if err := godotenv.Load("app.env"); err != nil {
		app.ErrorLog.Fatalf("error loading .env file: %s", err)
	}

	db, err := Run()
	if err != nil {
		app.ErrorLog.Fatal(err)
	}

	defer func(db *database.DB) {
		db.Close()
	}(db)

}
