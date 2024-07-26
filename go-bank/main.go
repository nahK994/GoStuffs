package main

import "fmt"

func main() {
	server := NewAPIServer("127.0.0.1:8000")
	fmt.Println("Starting server on", server.listenAddr)
	server.Run()
}
