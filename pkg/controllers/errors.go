package controllers

import (
	"net/http"
	"lib-manager/pkg/views"
)

func PageNotFound(res http.ResponseWriter, req *http.Request){
	t := views.PageNotFound()
	res.WriteHeader(http.StatusOK)
	t.Execute(res,nil )
}

func ForbiddenAccess(res http.ResponseWriter, req *http.Request){
	t := views.ForbiddenAccess()
	res.WriteHeader(http.StatusOK)
	t.Execute(res,nil )
}


