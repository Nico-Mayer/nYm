package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/nico-mayer/nym/config"
	"github.com/nico-mayer/nym/db"
	"github.com/nico-mayer/nym/handlers"
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
		if r.Method != http.MethodGet {
			return
		}
		tmpl := template.Must(template.ParseFiles("./static/templates/index.html"))
		tmpl.Execute(w, nil)
	})
	http.HandleFunc("/r/", handlers.HandleRedirect)
	http.HandleFunc("/add", handlers.AddLink)

	fmt.Println("Server started on " + config.PORT)
	log.Fatal(http.ListenAndServe(":"+config.PORT, nil))
}
