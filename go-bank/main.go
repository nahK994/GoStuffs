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
	if err = store.Init(); err != nil {
		log.Fatal(err.Error())
	}

	server := NewAPIServer("127.0.0.1:8000", store, "asdf")
	fmt.Println("Starting server on", server.listenAddr)
	server.Run()
}
