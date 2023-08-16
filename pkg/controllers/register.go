package controllers

import (
	"lib-manager/pkg/models"
	"lib-manager/pkg/views"
	"net/http"
)

// To submit registration form
func Register(writer http.ResponseWriter, request *http.Request) {

	username := request.FormValue("username")
	password := request.FormValue("password")
	reEnterPass := request.FormValue("reEnterPass")
	adminAccess := request.FormValue("adminAccess")

	db, err := models.Connection()
	if err != nil {
		http.Redirect(writer, request, "/error500", http.StatusSeeOther)
	}
	defer db.Close()

	status := models.RegisterUser(db, username, password, reEnterPass, adminAccess)

	if status == "OK" {
		t := views.StartPage()
		writer.WriteHeader(http.StatusOK)
		t.Execute(writer, status)
	} else {

		t := views.Register()
		writer.WriteHeader(http.StatusOK)
		t.Execute(writer, status)
	}

}

// To open register page
func RegisterPage(writer http.ResponseWriter, request *http.Request) {
	t := views.Register()
	writer.WriteHeader(http.StatusOK)
	t.Execute(writer, nil)
}
