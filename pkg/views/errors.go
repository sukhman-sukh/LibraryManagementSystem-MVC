package views

import (
	"html/template"
)

func Errors(viewMode string) *template.Template {
	mode := map[string]*template.Template{
		"ForbiddenAccess":       getTemplate("error403"),
		"PageNotFound":       getTemplate("error404"),
		"InternalError":       getTemplate("error500"),
	}

	temp := mode[viewMode]
	return temp
}
