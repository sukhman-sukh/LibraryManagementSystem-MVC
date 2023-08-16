package controllers

import (
	"lib-manager/pkg/models"
	"lib-manager/pkg/types"
	"lib-manager/pkg/views"
	"net/http"
)

// Mapping Code to status
//
//	bookStatus{
//	    1 : "Requested Check-In"
//	    0: "Check-out"
//	     -1: Requested- check-out
//	}
func GetAdmin(writer http.ResponseWriter, request *http.Request) {

	db, err := models.Connection()
	if err != nil {
		http.Redirect(writer, request, "/error500", http.StatusSeeOther)
	}
	defer db.Close()

	status, userId, userName, admin := models.Middleware(writer, request, db)

	if status == "OK" {
		if admin == 1 {

			_, books := models.GetBooks(db)
			_, requestBook := models.GetRequestBooks(db, userId ,userName)
			_, issuedBooks := models.GetIssuedBooks(db, userId, admin ,userName)
			_, adminRequest := models.GetAdminRequest(db)
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
	} else {
		http.Redirect(writer, request, "/login", http.StatusSeeOther)
	}

}



func AdminAccept(writer http.ResponseWriter, request *http.Request) {

	db, err := models.Connection()
	if err != nil {
		http.Redirect(writer, request, "/error500", http.StatusSeeOther)
	}
	defer db.Close()

	status, _, _, admin := models.Middleware(writer, request, db)

	if status == "OK" {
		if admin == 1 {

			requestId := request.FormValue("reqId")
			status := models.AdminAccept(writer, request, db, requestId)

			if status == "OK" {
				http.Redirect(writer, request, "/admin", http.StatusSeeOther)
			}
		} else {
			http.Redirect(writer, request, "/error403", http.StatusSeeOther)
		}
	} else {
		http.Redirect(writer, request, "/login", http.StatusSeeOther)
	}
}

func AdminDeny(writer http.ResponseWriter, request *http.Request) {
	db, err := models.Connection()
	if err != nil {
		http.Redirect(writer, request, "/error500", http.StatusSeeOther)
	}
	defer db.Close()

	status, _, _, admin := models.Middleware(writer, request, db)

	if status == "OK" {
		if admin == 1 {

			requestId := request.FormValue("reqId")

			status := models.AdminAccept(writer, request, db, requestId)

			if status == "OK" {
				http.Redirect(writer, request, "/admin", http.StatusSeeOther)
			}
		} else {
			http.Redirect(writer, request, "/error403", http.StatusSeeOther)
		}
	} else {
		http.Redirect(writer, request, "/login", http.StatusSeeOther)
	}
}
