package main

import "database/sql"

// runs the server
func main() {
	db, err := Run()
	if err != nil {
		app.ErrorLog.Fatal(err)
	}

	defer func(SQL *sql.DB) {
		err := SQL.Close()
		if err != nil {
			app.ErrorLog.Fatal(err)
		}
	}(db.SQL)
}
