package main

import (
	"fmt"
	"log"
)

func main() {
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err.Error())
	}
	server := NewAPIServer("127.0.0.1:8000", store)
	fmt.Println("Starting server on", server.listenAddr)
	server.Run()
}
