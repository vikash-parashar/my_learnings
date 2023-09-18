package render

import (
	"html/template"
	"log"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpName string) error {
	tmpl := template.Must(template.ParseFiles("templates/" + tmpName + ".html"))
	if err := tmpl.Execute(w, nil); err != nil {
		return err
	}
	log.Printf("template %s rendered successfully .", tmpName)
	return nil
}
