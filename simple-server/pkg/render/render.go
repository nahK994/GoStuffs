package render

import (
	"net/http"
	"html/template"
	"errors"
	"path/filepath"
	"bytes"
)

var functions = template.FuncMap{

}


func RenderTemplate(w http.ResponseWriter, temaplateName string) error {
	templateCache, err := CreateTemplateCache()
	if err != nil {
		return err
	}
	temp, ok := templateCache[temaplateName]
	if !ok {
		return errors.New("Template not found")
	}

	buf := new(bytes.Buffer)
	_ = temp.Execute(buf, nil)

	_, err = buf.WriteTo(w)
	if err != nil {
		return err
	}

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