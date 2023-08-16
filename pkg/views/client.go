package views

import (
	"html/template"
)

func Client(viewMode string) *template.Template {
	mode := map[string]*template.Template{
		"GetClient":       getTemplate("client"),
	}

	temp := mode[viewMode]
	return temp
}
