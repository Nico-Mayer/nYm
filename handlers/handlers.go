package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/nico-mayer/nym/db"
	"github.com/nico-mayer/nym/utils"
)

// Get Requests
func GetHomePage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}
	tmpl := template.Must(template.ParseFiles("./static/templates/index.html"))
	tmpl.Execute(w, nil)
}

func GetRedirect(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}
	// Extract the dynamic part of the URL
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	short_code := parts[2]
	link, err := db.GetLink(short_code)

	if err != nil {
		http.Error(w, "No redirect for this short link", http.StatusBadRequest)
	} else {
		http.Redirect(w, r, link.Long_url, http.StatusPermanentRedirect)
	}
}

// Put Requests
func PutCreateLink(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	if r.Method != http.MethodPut {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	long_url := r.PostFormValue("long_url")

	if !utils.IsValidURL(long_url) {
		msg := "Provided url is not valid"
		log.Println(msg)
		type Map = map[string]string
		data := Map{"My_error": msg}
		tmpl := template.Must(template.ParseFiles("./static/templates/index.html"))
		tmpl.ExecuteTemplate(w, "responseContainer", data)
		return
	}

	host := r.Host
	short_code, err := db.InsertLink(long_url)
	short_url := "http://" + host + "/r/" + short_code

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	delta := time.Since(start)
	type Map = map[string]string
	data := Map{"Short_url": short_url, "Delta": fmt.Sprint(delta)}
	tmpl := template.Must(template.ParseFiles("./static/templates/index.html"))
	tmpl.ExecuteTemplate(w, "responseContainer", data)
}
