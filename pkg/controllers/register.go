package controllers

import (
	"lib-manager/pkg/models"
	"lib-manager/pkg/views"
	"lib-manager/pkg/types"
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
	if status != "" {
		var response types.Response
		response.Message = status
		t := views.Welcome("Register")
		writer.WriteHeader(http.StatusOK)
		t.Execute(writer, response)
	} else {
		t := views.Welcome("StartPage")
		writer.WriteHeader(http.StatusOK)
		t.Execute(writer, nil)
	}
}

// To open register page
func RegisterPage(writer http.ResponseWriter, request *http.Request) {
	t := views.Welcome("Register")
	writer.WriteHeader(http.StatusOK)
	t.Execute(writer, nil)
}
