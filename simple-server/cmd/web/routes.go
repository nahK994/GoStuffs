package main

import (
	"net/http"
	"simple-server/pkg/handlers"

	"github.com/gorilla/mux"
)

func routes() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.Repo.Home)
	r.HandleFunc("/about", handlers.Repo.About)
	return r
}
