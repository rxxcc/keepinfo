package main

import (
	"log"
)

// runs the server
func main() {
	db, err := Run()
	if err != nil {
		log.Fatal(err)
	}

	defer db.SQL.Close()
}
