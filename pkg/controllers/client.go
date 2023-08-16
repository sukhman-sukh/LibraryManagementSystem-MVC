package controllers

import (
	"lib-manager/pkg/models"
	"lib-manager/pkg/types"
	"lib-manager/pkg/views"
	"net/http"
)

// Format

//	bookStatus{
//	    1 : "Requested Check-In"
//	    0: "Check-out"
//	     -1: Requested- check-out
//	}
func GetClient(writer http.ResponseWriter, request *http.Request) {

	db, err := models.Connection()
	if err != nil {
		http.Redirect(writer, request, "/error500", http.StatusSeeOther)
	}
	defer db.Close()

	status, userID, _, _ := models.Middleware(writer, request, db)

	if status == "OK" {
		_, books := models.GetBooks(db)
		_, requestBook := models.GetRequestBooks(db, userID)

		data := types.Data{
			UserName:     userName,
			Books:        books,
			RequestBook:  requestBook,
			AdminRequest: nil,
			IssuedBooks:  nil,
		}

		t := views.GetClient()
		writer.WriteHeader(http.StatusOK)
		t.Execute(writer, data)
	}

}

func CheckoutSubmit(writer http.ResponseWriter, request *http.Request) {

	db, err := models.Connection()
	if err != nil {
		http.Redirect(writer, request, "/error500", http.StatusSeeOther)
	}
	defer db.Close()

	status, userID, _, _ := models.Middleware(writer, request, db)
	bookId := request.FormValue("bookId")

	status = models.Checkout(writer, request, db, bookId, userID)

	if status == "OK" {
		http.Redirect(writer, request, "/", http.StatusSeeOther)
	}

}

func CheckinSubmit(writer http.ResponseWriter, request *http.Request) {
	db, err := models.Connection()

	if err != nil {
		http.Redirect(writer, request, "/error500", http.StatusSeeOther)
	}
	defer db.Close()

	requestId := request.FormValue("reqId")
	status := models.Checkin(writer, request, db, requestId)

	if status == "OK" {
		http.Redirect(writer, request, "/", http.StatusSeeOther)
	}

}
