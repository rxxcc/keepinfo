package main

// runs the server
func main() {
	db, err := Run()
	if err != nil {
		app.ErrorLog.Fatal(err)
	}

	defer db.SQL.Close()
}
