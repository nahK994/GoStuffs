package main

import "fmt"

func main() {
	fmt.Println("Hello World!!!")
	server := NewAPIServer("127.0.0.1:8000")
	server.Run()
}
