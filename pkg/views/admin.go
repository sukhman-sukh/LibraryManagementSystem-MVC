package views

import (
	"html/template"
)

func GetAdmin() *template.Template {
	temp := template.Must(template.ParseFiles("templates/screens/admin.html"))
	return temp
}

func Checkin() *template.Template {
	temp := template.Must(template.ParseFiles("templates/screens/checkin.html"))
	return temp
}

func AdminAdd() *template.Template {
	temp := template.Must(template.ParseFiles("templates/screens/addBooks.html"))
	return temp
}


func AdminCheckout() *template.Template {
	temp := template.Must(template.ParseFiles("templates/screens/checkout.html"))
	return temp
}