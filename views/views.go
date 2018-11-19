package views

import (
	"html/template"
	"log"
)

var templates *template.Template

func Init() error {
	var err error
	templates, err = template.ParseGlob("files/var/www/templates/*")
	if err != nil {
		return err
	}

	log.Println("Parse template successfully")
	return nil
}

func Get() *template.Template {
	return templates
}
