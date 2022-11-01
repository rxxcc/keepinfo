package main

import (
	"github.com/inuoshios/keepinfo/internal/database"
	"github.com/joho/godotenv"
)

func init() {
	// load env files.
	if err := godotenv.Load("app.env"); err != nil {
		app.ErrorLog.Fatalf("error loading .env file: %s", err)
	}
}

// runs the server
func main() {
	db, err := Run()
	if err != nil {
		app.ErrorLog.Fatal(err)
	}

	defer func(db *database.DB) {
		db.SQL.Close()
	}(db)

}
