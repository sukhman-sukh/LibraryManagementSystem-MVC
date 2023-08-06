package views

import (
	"html/template"
)

func GetAdmin() *template.Template {
	temp := template.Must(template.ParseFiles("templates/screens/admin.html"))
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

func AdminRemove() *template.Template {
	temp := template.Must(template.ParseFiles("templates/screens/removeBooks.html"))
    return temp
}

func AdminChoose() *template.Template {
	temp := template.Must(template.ParseFiles("templates/screens/choose.html"))
    return temp
}