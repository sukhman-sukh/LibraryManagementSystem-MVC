package views

import (
	"html/template"
)

func ForbiddenAccess() *template.Template {
	temp := template.Must(template.ParseFiles("templates/screens/error403.html"))
	return temp
}

func PageNotFound() *template.Template {
	temp := template.Must(template.ParseFiles("templates/screens/error404.html"))
	return temp
}

func InternalError() *template.Template {
	temp := template.Must(template.ParseFiles("templates/screens/error500.html"))
	return temp
}
