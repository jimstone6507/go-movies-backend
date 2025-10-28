package main

import (
	"fmt"
	"log"
	"net/http"
)

// http port
const port = 8080


// application configuration
type application struct {
	Domain string
}

func main() {
	log.Println("starting app....")

	var app application

	app.Domain = "example.com"

	// set application config

	// read from command line

	// connect to the database


	log.Println("Starting the application on port ", port)
	// start a web server
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}	
}
