package controllers

import (
	"lib-manager/pkg/middleware"
	"lib-manager/pkg/models"
	"lib-manager/pkg/types"
	"lib-manager/pkg/views"
	"net/http"
)

func GetAdmin(writer http.ResponseWriter, request *http.Request) {
	db, err := models.Connection()
	if err != nil {
		http.Redirect(writer, request, "/error500", http.StatusSeeOther)
	}
	defer db.Close()

	userId, userName, admin, err := middleware.Middleware(writer, request, db)
	if err != nil {
		http.Redirect(writer, request, "/login", http.StatusSeeOther)
	} else {
		if admin == 1 {
			books, _ := models.GetBooks(db)
			requestBook, _ := models.GetRequestBooks(db, userId, userName)
			issuedBooks, _ := models.GetIssuedBooks(db, userId, admin, userName)
			adminRequest, _ := models.GetAdminRequest(db)
			data := types.Data{
				UserName:     userName,
				Books:        books,
				RequestBook:  requestBook,
				AdminRequest: adminRequest,
				IssuedBooks:  issuedBooks,
			}
			t := views.Admin("GetAdmin")
			writer.WriteHeader(http.StatusOK)
			t.Execute(writer, data)
		} else {
			http.Redirect(writer, request, "/error403", http.StatusSeeOther)
		}
	}
}

func AdminAccept(writer http.ResponseWriter, request *http.Request) {
	db, err := models.Connection()
	if err != nil {
		http.Redirect(writer, request, "/error500", http.StatusSeeOther)
	}
	defer db.Close()

	_, _, admin, err := middleware.Middleware(writer, request, db)
	if err != nil {
		http.Redirect(writer, request, "/login", http.StatusSeeOther)
	} else {
		if admin == 1 {
			requestId := request.FormValue("reqId")
			models.AdminAccept(writer, request, db, requestId)
			http.Redirect(writer, request, "/admin", http.StatusSeeOther)
		} else {
			http.Redirect(writer, request, "/error403", http.StatusSeeOther)
		}
	}
}

func AdminDeny(writer http.ResponseWriter, request *http.Request) {
	db, err := models.Connection()
	if err != nil {
		http.Redirect(writer, request, "/error500", http.StatusSeeOther)
	}
	defer db.Close()

	_, _, admin, err := middleware.Middleware(writer, request, db)
	if err != nil {
		http.Redirect(writer, request, "/login", http.StatusSeeOther)
	} else {
		if admin == 1 {
			requestId := request.FormValue("reqId")
			models.AdminAccept(writer, request, db, requestId)
			http.Redirect(writer, request, "/admin", http.StatusSeeOther)
		} else {
			http.Redirect(writer, request, "/error403", http.StatusSeeOther)
		}
	}
}
