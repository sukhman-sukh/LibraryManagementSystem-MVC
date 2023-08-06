package controllers

import (
	// "encoding/json"
	// "fmt"
	"net/http"
	"lib-manager/pkg/views"
)


func GetClient(res http.ResponseWriter, req *http.Request) {

	t := views.GetClient()
	res.WriteHeader(http.StatusOK)
	t.Execute(res, nil)

}