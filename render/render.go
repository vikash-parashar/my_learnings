package render

import (
	"html/template"
	"log"

	"github.com/gin-gonic/gin"
)

func RenderTemplate(c *gin.Context, tmpName string) error {
	tmpl := template.Must(template.ParseFiles("templates/" + tmpName + ".html"))
	if err := tmpl.Execute(c.Writer, nil); err != nil {
		return err
	}
	log.Printf("Template %s rendered successfully.", tmpName)
	return nil
}
