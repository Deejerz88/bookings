package render

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var tc = make(map[string]*template.Template)

func RenderTemplateSimple(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	//check if we already have template in cache
	_, inMap := tc[t]

	if !inMap {
		// need to create template
		fmt.Println("creating template and adding to cache")
		err = createTemplateCache1(t)
		if err != nil {
			fmt.Println("Error creating template cache :", err)
			return
		}
	} else {
		// we have template
		fmt.Println("getting template from cache")
	}

	tmpl = tc[t]

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}

func createTemplateCache1(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t), "./templates/base.layout.tmpl",
	}

	// parse template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	// add to template cache
	tc[t] = tmpl

	return nil
}
