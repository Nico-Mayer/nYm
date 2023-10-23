package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nico-mayer/nym/config"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method)
	})

	fmt.Println("Server started on " + config.PORT)
	log.Fatal(http.ListenAndServe(":"+config.PORT, nil))
}
