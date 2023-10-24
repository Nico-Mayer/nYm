package main

import (
	"log"
	"net/http"

	"github.com/nico-mayer/nym/config"
	"github.com/nico-mayer/nym/db"
	"github.com/nico-mayer/nym/handlers"
)

func main() {
	err := db.Connect()

	if err != nil {
		log.Println("Error on connecting to db")
	} else {
		log.Println("Connected to Database")
	}

	initServer()
}

func initServer() {
	// Get Requests
	http.HandleFunc("/", handlers.GetHomePage)
	http.HandleFunc("/r/", handlers.GetRedirect)
	// Put Requests
	http.HandleFunc("/add", handlers.PutCreateLink)

	log.Println("Server started on " + config.PORT)
	log.Fatal(http.ListenAndServe(":"+config.PORT, nil))
}
