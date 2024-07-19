package main
import (
	"fmt"
	"errors"
	"net/http"
	"html/template"
)

const portNumber = 8080

func Home(w http.ResponseWriter, r *http.Request) {
	// _, err := fmt.Fprintf(w, "This is HOME page")
	err := parseTemplate(w, "home.html")
	if err != nil {
		fmt.Println(err)
	}
}

func About(w http.ResponseWriter, r *http.Request) {
	// _, err := fmt.Fprintf(w, "This is ABOUT page")
	err := parseTemplate(w, "about.html")
	if err != nil {
		fmt.Println(err)
	}
}

func parseTemplate(w http.ResponseWriter, temaplateName string) error {
	parsedTemplate, _ := template.ParseFiles("./templates/" + temaplateName)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		return errors.New("Cannot be parsed")
	}
	return nil
}

func main()  {
	// fmt.Println("Hello World!!!")
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println("Starting application on port:", portNumber)
	var uri = fmt.Sprintf(":%d", portNumber)
	err := http.ListenAndServe(uri, nil)
	if err != nil {
		fmt.Println(err)
	}
}