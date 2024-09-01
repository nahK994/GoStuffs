package main

import (
	"go-book-store/pkg/routes"
	"log"
	"net/http"
)

func main() {
	r := routes.RegisterRoutes()

	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
