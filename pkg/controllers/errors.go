package controllers

import (
	"lib-manager/pkg/views"
	"net/http"
)

func PageNotFound(writer http.ResponseWriter, request *http.Request) {
	t := views.Errors("PageNotFound")
	writer.WriteHeader(http.StatusOK)
	t.Execute(writer, nil)
}

func ForbiddenAccess(writer http.ResponseWriter, request *http.Request) {
	t := views.Errors("ForbiddenAccess")
	writer.WriteHeader(http.StatusOK)
	t.Execute(writer, nil)
}

func InternalError(writer http.ResponseWriter, request *http.Request) {
	t := views.Errors("InternalError")
	writer.WriteHeader(http.StatusOK)
	t.Execute(writer, nil)
}
