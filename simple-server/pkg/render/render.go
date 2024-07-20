package render

import (
	"net/http"
	"html/template"
	"errors"
)


func RenderTemplate(w http.ResponseWriter, temaplateName string) error {
	parsedTemplate, _ := template.ParseFiles("./templates/" + temaplateName)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		return errors.New("Cannot be parsed")
	}
	return nil
}