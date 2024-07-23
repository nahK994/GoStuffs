package handlers

import (
	"net/http"
	"simple-server/pkg/render"
	"simple-server/pkg/config"
	"simple-server/pkg/models"
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
	err := render.RenderTemplate(w, "home.page.html", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// _, err := fmt.Fprintf(w, "This is ABOUT page")
	data := &models.TemplateData{
		StringMap: map[string]string{
			"name": "Shomi Khan",
			"email": "nkskl6@gmail.com",
			"phone": "+8801676498001",
		},
	}
	err := render.RenderTemplate(w, "about.page.html", data)
	if err != nil {
		fmt.Println(err)
	}
}