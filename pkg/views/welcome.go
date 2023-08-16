package views

import (
	"html/template"
)

func Welcome(viewMode string) *template.Template {
	mode := map[string]*template.Template{
		"StartPage":       getTemplate("welcome"),
		"LogIn":       getTemplate("login"),
		"Register":       getTemplate("register"),
	}

	temp := mode[viewMode]
	return temp
}
