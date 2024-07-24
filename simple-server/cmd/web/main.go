package main
import (
	"fmt"
	"net/http"
	"simple-server/pkg/handlers"
	"simple-server/pkg/config"
	"simple-server/pkg/render"
	"log"
)

const portNumber = 8080

func main()  {
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

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	// var uri = fmt.Sprintf(":%d", portNumber)
	// err1 := http.ListenAndServe(uri, nil)
	// if err1 != nil {
	// 	fmt.Println(err1)
	// }

	srv := &http.Server{
        Handler:      routes(&app)
        Addr:         "127.0.0.1:8000",
	}

	err := srv.ListenAndServe()
	log.Fatal(err)
}