package routes

import (
	"book-store/pkg/controllers"
	"book-store/pkg/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/books", controllers.GetBooks).Methods(http.MethodGet)
	r.HandleFunc("/books/{id}", handlers.HandleHTTP(controllers.GetBook)).Methods(http.MethodGet)
	r.HandleFunc("/books", handlers.HandleHTTP(controllers.CreateBook)).Methods(http.MethodPost)
	r.HandleFunc("/books/{id}", handlers.HandleHTTP(controllers.UpdateBook)).Methods(http.MethodPut)
	r.HandleFunc("/books/{id}", handlers.HandleHTTP(controllers.DeleteBook)).Methods(http.MethodDelete)

	return r
}
