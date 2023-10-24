package controllers

import (
	"net/http"
	"strings"

	"github.com/nico-mayer/nym/models"
	"github.com/nico-mayer/nym/utils"
)

func AddLink(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	longLink := r.Header.Get("long_url")

	if longLink == "" {
		http.Error(w, "No long_url header in the request", http.StatusBadRequest)
		return
	}
	if !utils.IsValidURL(longLink) {
		http.Error(w, "Provided link is not a Valid URL", http.StatusBadRequest)
		return
	}

	short_code, err := models.InsertLink(longLink)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(short_code))
}

func HandleRedirect(w http.ResponseWriter, r *http.Request) {
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
	link, err := models.GetLink(short_code)

	if err != nil {
		http.Error(w, "No redirect for this short link", http.StatusBadRequest)
	} else {
		http.Redirect(w, r, link.Long_url, http.StatusPermanentRedirect)
	}
}
