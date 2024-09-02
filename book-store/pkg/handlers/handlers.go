package handlers

import (
	"book-store/pkg/errors"
	"book-store/pkg/models"
	"encoding/json"
	"net/http"
)

type controllerFunc func(r *http.Request) (*models.Response, error)

func HandleHTTP(handleFunc controllerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		resp, err := handleFunc(r)
		if err != nil {
			switch err.(type) {
			case *errors.NotFound:
				w.WriteHeader(http.StatusNotFound)
			case *errors.BadRequest:
				w.WriteHeader(http.StatusBadRequest)
			}
			json.NewEncoder(w).Encode(err.Error())
		} else {
			w.WriteHeader(resp.StatusCode)
			json.NewEncoder(w).Encode(resp.Body)
		}
	}
}
