package main

import(
	"github.com/gorilla/mux"
	"simple-server/pkg/config"
	"simple-server/pkg/handlers"
	"net/http"
)


func routes(app *config.AppConfig) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.Repo.Home)
	r.HandleFunc("/about", handlers.Repo.About)
	return r
}