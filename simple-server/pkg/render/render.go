package render

import (
	"net/http"
	"html/template"
	"errors"
	"path/filepath"
	"simple-server/pkg/config"
	// "bytes"
)

var app *config.AppConfig

func SetAppConfig(a *config.AppConfig) {
	app = a
}

var functions = template.FuncMap{

}


func RenderTemplate(w http.ResponseWriter, temaplateName string) error {
	var templateCache map[string]*template.Template

	if app.UseCache {
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}

	temp, ok := templateCache[temaplateName]
	if !ok {
		return errors.New("Template not found")
	}

	_ = temp.Execute(w, nil)

	// parsedTemplate, _ := template.ParseFiles("./templates/" + temaplateName)
	// err := parsedTemplate.Execute(w, nil)
	// if err != nil {
	// 	return errors.New("Cannot be parsed")
	// }
	return nil
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}
	
		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return nil, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return nil, err
			}	
		}

		myCache[name] = ts
	}

	return myCache, nil
}