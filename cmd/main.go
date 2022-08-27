package main

import "log"

// runs the server
func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}

}
