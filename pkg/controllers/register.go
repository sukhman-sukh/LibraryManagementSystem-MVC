package controllers

import (
	// "encoding/json"
	"fmt"
	"net/http"
	"lib-manager/pkg/views"
	"lib-manager/pkg/models"
)

// To submit registration form
func Register(res http.ResponseWriter, req *http.Request) {
	username := req.FormValue("username")
	password := req.FormValue("password")
	reEnterPass := req.FormValue("reEnterPass") 
	adminAccess := req.FormValue("adminAccess")

	fmt.Println(username , password , reEnterPass , adminAccess)
	status := models.RegisterUser(username , password , reEnterPass , adminAccess)
	if(status == "OK"){
		t := views.StartPage()
		res.WriteHeader(http.StatusOK)
		t.Execute(res,status)
	}else{
		t := views.StartPage()
		res.WriteHeader(http.StatusOK)
		t.Execute(res,status)
	}


}

// To open register page
func RegisterPage(res http.ResponseWriter, req *http.Request){
	t := views.Register()
	res.WriteHeader(http.StatusOK)
	t.Execute(res,nil )
}