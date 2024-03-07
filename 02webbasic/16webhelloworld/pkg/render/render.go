package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func RenderTemplates(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template: ", err)
		return
	}
}

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// Kiểm tra xem template đã có trong cache chưa
	_, inMap := tc[t]
	if !inMap {
		// Tạo template
		log.Println("creating template and add to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		log.Println("using cache template")
	}

	tmpl = tc[t]
	err = tmpl.Execute(w, nil)

	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		"./templates/" + t, "./templates/base.layout.tmpl",
	}

	// Parse templates
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	tc[t] = tmpl
	return nil
}
