package controllers

import (
	"lib-manager/pkg/models"
	"lib-manager/pkg/views"
	"net/http"
	"lib-manager/pkg/middleware"
)

type key int

const (
	adminAuthKey key    = iota
	userID       key    = iota
	userName     string = ""
)

func Welcome(writer http.ResponseWriter, request *http.Request) {
	db, err := models.Connection()
	if err != nil {
		http.Redirect(writer, request, "/error500", http.StatusSeeOther)
	}
	defer db.Close()

	_, _, admin, err := middleware.Middleware(writer, request, db)
	if err != nil {
		t := views.Welcome("StartPage")
		writer.WriteHeader(http.StatusOK)
		t.Execute(writer, nil)
	} else {
		if admin == 1 {
			http.Redirect(writer, request, "/admin", http.StatusSeeOther)
		} else {
			http.Redirect(writer, request, "/client", http.StatusSeeOther)
		}
	}
}
