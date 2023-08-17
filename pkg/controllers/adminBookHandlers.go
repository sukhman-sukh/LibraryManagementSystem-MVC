package controllers

import (
	"lib-manager/pkg/middleware"
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

	_, _, admin, err := middleware.Middleware(writer, request, db)
	if err != nil {
		http.Redirect(writer, request, "/login", http.StatusSeeOther)
	} else {
		if admin == 1 {
			requestId := request.FormValue("reqId")
			models.AdminCheckin(writer, request, db, requestId)
			http.Redirect(writer, request, "/admin", http.StatusSeeOther)
		} else {
			http.Redirect(writer, request, "/error403", http.StatusSeeOther)
		}
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

	_, _, admin, err := middleware.Middleware(writer, request, db)
	if err != nil {
		http.Redirect(writer, request, "/login", http.StatusSeeOther)
	} else {
		if admin == 1 {
			bookname := request.FormValue("bookname")
			Author := request.FormValue("Author")
			Copies := request.FormValue("Copies")
			models.AdminAdd(writer, request, db, bookname, Author, Copies)
			http.Redirect(writer, request, "/admin", http.StatusSeeOther)
		} else {
			http.Redirect(writer, request, "/error403", http.StatusSeeOther)
		}
	}
}

func AdminCheckoutSubmit(writer http.ResponseWriter, request *http.Request) {
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
			models.AdminCheckout(writer, request, db, requestId)
			http.Redirect(writer, request, "/admin", http.StatusSeeOther)
		} else {
			http.Redirect(writer, request, "/error403", http.StatusSeeOther)
		}
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

	_, _, admin, err := middleware.Middleware(writer, request, db)
	if err != nil {
		http.Redirect(writer, request, "/login", http.StatusSeeOther)
	} else {
		if admin == 1 {
			bookId := request.FormValue("bookId")
			copies := request.FormValue("Copies")
			models.AdminRemove(writer, request, db, bookId, copies)
			http.Redirect(writer, request, "/admin", http.StatusSeeOther)
		} else {
			http.Redirect(writer, request, "/error403", http.StatusSeeOther)
		}
	}
}
