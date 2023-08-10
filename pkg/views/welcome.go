package views

import (
	"html/template"
)

func StartPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/screens/welcome.html"))
	return temp
}

func LogIn() *template.Template {
	temp := template.Must(template.ParseFiles("templates/screens/login.html"))
	
	return temp
}


func Register() *template.Template {
	temp := template.Must(template.ParseFiles("templates/screens/register.html"))
	return temp
}
