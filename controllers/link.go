package controllers

import (
	"net/http"

	"github.com/nico-mayer/nym/models"
	"github.com/nico-mayer/nym/utils"
)

func AddLink(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	longLink := r.Header.Get("longLink")

	if longLink == "" {
		http.Error(w, "No longLink header in the request", http.StatusBadRequest)
		return
	}
	if !utils.IsValidURL(longLink) {
		http.Error(w, "Provided link is not a Valid URL", http.StatusBadRequest)
		return
	}

	err := models.InsertLink(longLink)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
