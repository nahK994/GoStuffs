package main
import (
	"fmt"
	"errors"
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

func Divide(w http.ResponseWriter, r *http.Request) {
	var value, err = divideValues(2,0)
	if err != nil {
		fmt.Fprintf(w, "Divided by 0 error")
	}
	fmt.Fprintf(w, fmt.Sprintf("%f / %f = %f", 2, 0, value))
}

func divideValues(x, y float32) (float32, error) {
	if y == 0 {
		err := errors.New("Cannot divided by 0")
		return 0, err
	}

	return (x/y), nil
}

func main()  {
	// fmt.Println("Hello World!!!")
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println("Starting application on port:", portNumber)
	var uri = fmt.Sprintf(":%d", portNumber)
	err := http.ListenAndServe(uri, nil)
	if err != nil {
		fmt.Println(err)
	}
}