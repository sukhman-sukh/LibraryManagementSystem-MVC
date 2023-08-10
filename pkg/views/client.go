package views

import (
	"html/template"
)

func GetClient() *template.Template {
	temp := template.Must(template.ParseFiles("templates/screens/client.html"))
	return temp
}

func Checkin() *template.Template {
	temp := template.Must(template.ParseFiles("templates/screens/reqCheckin.html"))
	return temp
}
func Checkout() *template.Template {
	temp := template.Must(template.ParseFiles("templates/screens/reqCheckout.html"))
	return temp
}