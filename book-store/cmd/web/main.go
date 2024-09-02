package main

import (
	"book-store/pkg/config"
	"book-store/pkg/routes"
	"log"
	"net/http"
)

func main() {
	config.InitDB()
	r := routes.RegisterRoutes()

	log.Println("Starting server on :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
