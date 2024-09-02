package routes

import (
	"book-store/pkg/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/books", controllers.GetBooks).Methods(http.MethodGet)
	r.HandleFunc("/books/{id}", controllers.GetBook).Methods(http.MethodGet)
	r.HandleFunc("/books", controllers.CreateBook).Methods(http.MethodPost)
	r.HandleFunc("/books/{id}", controllers.UpdateBook).Methods(http.MethodPut)
	r.HandleFunc("/books/{id}", controllers.DeleteBook).Methods(http.MethodDelete)

	return r
}
