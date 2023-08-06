package views

import (
	"html/template"
)

func GetClient() *template.Template {
	temp := template.Must(template.ParseFiles("templates/screens/client.html"))
	return temp
}
