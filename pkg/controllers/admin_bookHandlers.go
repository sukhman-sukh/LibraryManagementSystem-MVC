package controllers

import (
	"lib-manager/pkg/models"
	"lib-manager/pkg/views"
	"net/http"
)

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
	t := views.Admin("AdminAdd")
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
	t := views.Admin("AdminRemove")
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
