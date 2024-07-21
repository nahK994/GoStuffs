package handlers

import (
	"net/http"
	"simple-server/pkg/render"
	"simple-server/pkg/config"
	"fmt"
)

var Repo *Repository
type Repository struct {
	App *config.AppConfig
}

func CreateRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func SetRepo(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// _, err := fmt.Fprintf(w, "This is HOME page")
	err := render.RenderTemplate(w, "home.page.html")
	if err != nil {
		fmt.Println(err)
	}
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// _, err := fmt.Fprintf(w, "This is ABOUT page")
	err := render.RenderTemplate(w, "about.page.html")
	if err != nil {
		fmt.Println(err)
	}
}