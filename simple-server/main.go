package main
import (
	"fmt"
	"net/http"
)

const portNumber = 8080

func Home(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "This is HOME page")
	if err != nil {
		fmt.Println(err)
	}
}

func About(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "This is ABOUT page")
	if err != nil {
		fmt.Println(err)
	}
}

func main()  {
	fmt.Println("Hello World!!!")

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println("Starting application on port:", portNumber)
	var uri = fmt.Sprintf(":%d", portNumber)
	err := http.ListenAndServe(uri, nil)
	if err != nil {
		fmt.Println(err)
	}
}