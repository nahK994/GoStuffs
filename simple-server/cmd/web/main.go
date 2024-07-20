package main
import (
	"fmt"
	"net/http"
	"simple-server/pkg/handlers"
)

const portNumber = 8080

func main()  {
	// fmt.Println("Hello World!!!")
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Starting application on port:", portNumber)
	var uri = fmt.Sprintf(":%d", portNumber)
	err := http.ListenAndServe(uri, nil)
	if err != nil {
		fmt.Println(err)
	}
}