package main

import (
	"fmt"
	"log"
	"net/http"
	"simple-server/pkg/config"
	"simple-server/pkg/handlers"
	"simple-server/pkg/render"
)

const portNumber = 8080

func main() {
	// fmt.Println("Hello World!!!")
	fmt.Println("Starting application on port:", portNumber)

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	var app config.AppConfig
	app.TemplateCache = templateCache
	app.UseCache = false
	render.SetAppConfig(&app)

	repo := handlers.CreateRepo(&app)
	handlers.SetRepo(repo)

	var uri = fmt.Sprintf("127.0.0.1:%d", portNumber)
	srv := &http.Server{
		Handler: routes(),
		Addr:    uri,
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
