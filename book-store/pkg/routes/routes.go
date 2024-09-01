package routes

import (
	"book-store/pkg/config"
	"book-store/pkg/controllers"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	config.InitDB()
	r := mux.NewRouter()

	r.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", controllers.GetBook).Methods("GET")
	r.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")

	return r
}
