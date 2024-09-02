package handlers

import (
	"book-store/pkg/models"
	"encoding/json"
	"net/http"
)

type controllerFunc func(w http.ResponseWriter, r *http.Request) *models.Response

func HandleHTTP(handleFunc controllerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := handleFunc(w, r)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(resp.StatusCode)
		json.NewEncoder(w).Encode(resp.Body)
	}
}
