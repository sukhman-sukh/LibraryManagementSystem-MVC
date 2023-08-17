package controllers

import (
	"lib-manager/pkg/models"
	"lib-manager/pkg/types"
	"lib-manager/pkg/middleware"
	"lib-manager/pkg/views"
	"net/http"
)

func GetClient(writer http.ResponseWriter, request *http.Request) {
	db, err := models.Connection()
	if err != nil {
		http.Redirect(writer, request, "/error500", http.StatusSeeOther)
	}
	defer db.Close() 
	
	userID, userName, _, err := middleware.Middleware(writer, request, db)
	if err != nil {
		books, _ := models.GetBooks(db)
		requestBook, _ := models.GetRequestBooks(db, userID, userName)
		data := types.Data{
			UserName:     userName,
			Books:        books,
			RequestBook:  requestBook,
			AdminRequest: nil,
			IssuedBooks:  nil,
		}
		t := views.Client("GetClient")
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

	userID, _, _, _ := middleware.Middleware(writer, request, db)
	bookId := request.FormValue("bookId")

	models.Checkout(writer, request, db, bookId, userID)
	http.Redirect(writer, request, "/", http.StatusSeeOther)
}

func CheckinSubmit(writer http.ResponseWriter, request *http.Request) {
	db, err := models.Connection()
	if err != nil {
		http.Redirect(writer, request, "/error500", http.StatusSeeOther)
	}
	defer db.Close()

	requestId := request.FormValue("reqId")
	models.Checkin(writer, request, db, requestId)
	http.Redirect(writer, request, "/", http.StatusSeeOther)
}
