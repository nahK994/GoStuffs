package routes

import (
	"mybank/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/admin/users", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/admin/users/{id}", controllers.UpdateUser).Methods("PUT")
	r.HandleFunc("/admin/users/{id}", controllers.DeleteUser).Methods("DELETE")

	r.HandleFunc("/users/{id}/credit", controllers.CreditBalance).Methods("POST")
	r.HandleFunc("/users/{id}/debit", controllers.DebitBalance).Methods("POST")
	r.HandleFunc("/users/transfer", controllers.TransferMoney).Methods("POST")

	return r
}
