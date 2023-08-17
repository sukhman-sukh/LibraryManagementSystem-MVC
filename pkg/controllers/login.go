package controllers

import (
	"lib-manager/pkg/models"
	"lib-manager/pkg/views"
	"net/http"
	"lib-manager/pkg/middleware"
)

// To open login page
func LogIn(writer http.ResponseWriter, request *http.Request) {
	t := views.Welcome("LogIn")
	writer.WriteHeader(http.StatusOK)
	t.Execute(writer, nil)
}

// Check credentials for login
func ChecklogIn(writer http.ResponseWriter, request *http.Request) {
	username := request.FormValue("username")
	password := request.FormValue("password")

	db, err := models.Connection()
	if err != nil {
		http.Redirect(writer, request, "/error500", http.StatusSeeOther)
	}
	defer db.Close()

	models.LoginUser(writer, request, db, username, password)
}

// Logout And End Session
func Logout(writer http.ResponseWriter, request *http.Request) {
	db, err := models.Connection()
	if err != nil {
		http.Redirect(writer, request, "/error500", http.StatusSeeOther)
	}
	defer db.Close()

	userID, _, _, err := middleware.Middleware(writer, request, db)
	if err != nil {
		http.Redirect(writer, request, "/login", http.StatusSeeOther)
	} else {
		models.Logout(writer, request, db, userID)
	}
}
