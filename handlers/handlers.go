package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

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
	if r.Method != http.MethodPut {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	long_url := r.PostFormValue("long_url")
	fmt.Println(long_url)

	if long_url == "" {
		http.Error(w, "No long_url header in the request", http.StatusBadRequest)
		return
	}
	if !utils.IsValidURL(long_url) {
		http.Error(w, "Provided link is not a Valid URL", http.StatusBadRequest)
		return
	}

	short_code, err := db.InsertLink(long_url)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(short_code))
}
