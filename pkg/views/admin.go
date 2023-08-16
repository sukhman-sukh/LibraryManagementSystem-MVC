package views

import (
	"html/template"
)

func Admin(viewMode string) *template.Template {
	mode := map[string]*template.Template{
		"GetAdmin":       getTemplate("admin"),
		"AdminAdd":    getTemplate("addBooks"),
		"AdminRemove": getTemplate("removeBooks"),
	}

	temp := mode[viewMode]
	return temp
}

func getTemplate(viewMode string) *template.Template {
	return template.Must(template.ParseFiles("templates/screens/" + viewMode+".html"))
}
