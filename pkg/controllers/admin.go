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
			_, requestBook := models.GetRequestBooks(db, userId)
			_, issuedBooks := models.GetIssuedBooks(db, userId, admin)
			_, adminRequest := models.GetAdminRequest(db)
			data := types.Data{
				UserName:     userName,
				Books:        books,
				RequestBook:  requestBook,
				AdminRequest: adminRequest,
				IssuedBooks:  issuedBooks,
			}
			t := views.GetAdmin()
			writer.WriteHeader(http.StatusOK)
			t.Execute(writer, data)
		} else {
			http.Redirect(writer, request, "/error403", http.StatusSeeOther)
		}
	} else {
		http.Redirect(writer, request, "/login", http.StatusSeeOther)
	}

}

func AdminCheckinSubmit(writer http.ResponseWriter, request *http.Request) {

	db, err := models.Connection()

	if err != nil {
		http.Redirect(writer, request, "/error500", http.StatusSeeOther)
	}
	defer db.Close()

	status, _, _, admin := models.Middleware(writer, request, db)

	if status == "OK" {
		if admin == 1 {

			requestId := request.FormValue("reqId")
			_ = models.AdminCheckin(writer, request, db, requestId)

			if status == "OK" {
				http.Redirect(writer, request, "/admin", http.StatusSeeOther)
			}
		} else {
			http.Redirect(writer, request, "/error403", http.StatusSeeOther)
		}
	} else {
		http.Redirect(writer, request, "/error500", http.StatusSeeOther)
	}
}

func AdminAdd(writer http.ResponseWriter, request *http.Request) {
	t := views.AdminAdd()
	writer.WriteHeader(http.StatusOK)
	t.Execute(writer, nil)
}

func AdminAddSubmit(writer http.ResponseWriter, request *http.Request) {

	db, err := models.Connection()
	if err != nil {
		http.Redirect(writer, request, "/error500", http.StatusSeeOther)
	}
	defer db.Close()
	status, _, _, admin := models.Middleware(writer, request, db)

	if status == "OK" {
		if admin == 1 {

			bookname := request.FormValue("bookname")
			Author := request.FormValue("Author")
			Copies := request.FormValue("Copies")

			status := models.AdminAdd(writer, request, db, bookname, Author, Copies)
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

func AdminCheckoutSubmit(writer http.ResponseWriter, request *http.Request) {
	db, err := models.Connection()
	if err != nil {
		http.Redirect(writer, request, "/error500", http.StatusSeeOther)
	}
	defer db.Close()
	status, _, _, admin := models.Middleware(writer, request, db)

	if status == "OK" {
		if admin == 1 {

			requestId := request.FormValue("reqId")
			status := models.AdminCheckout(writer, request, db, requestId)

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

func AdminRemove(writer http.ResponseWriter, request *http.Request) {
	t := views.AdminRemove()
	writer.WriteHeader(http.StatusOK)
	t.Execute(writer, nil)
}

func AdminRemoveSubmit(writer http.ResponseWriter, request *http.Request) {

	db, err := models.Connection()
	if err != nil {
		http.Redirect(writer, request, "/error500", http.StatusSeeOther)
	}
	defer db.Close()
	status, _, _, admin := models.Middleware(writer, request, db)

	if status == "OK" {
		if admin == 1 {

			bookId := request.FormValue("bookId")
			copies := request.FormValue("Copies")

			status := models.AdminRemove(writer, request, db, bookId, copies)

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
