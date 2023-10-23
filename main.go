package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nico-mayer/nym/config"
	"github.com/nico-mayer/nym/controllers"
	"github.com/nico-mayer/nym/db"
)

func main() {
	err := db.Connect()

	if err != nil {
		fmt.Println("Error on connecting to db")
	} else {
		fmt.Println("Connected to Database")
	}

	initServer()
}

func initServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method)
	})

	http.HandleFunc("/add", controllers.AddLink)

	fmt.Println("Server started on " + config.PORT)
	log.Fatal(http.ListenAndServe(":"+config.PORT, nil))
}
