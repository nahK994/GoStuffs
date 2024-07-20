package handlers

import (
	"net/http"
	"simple-server/pkg/render"
	"fmt"
)


func Home(w http.ResponseWriter, r *http.Request) {
	// _, err := fmt.Fprintf(w, "This is HOME page")
	err := render.RenderTemplate(w, "home.html")
	if err != nil {
		fmt.Println(err)
	}
}

func About(w http.ResponseWriter, r *http.Request) {
	// _, err := fmt.Fprintf(w, "This is ABOUT page")
	err := render.RenderTemplate(w, "about.html")
	if err != nil {
		fmt.Println(err)
	}
}