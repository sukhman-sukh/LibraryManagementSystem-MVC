package controllers

import (
	// "encoding/json"
	"fmt"
	"net/http"
	"lib-manager/pkg/views"
	"lib-manager/pkg/models"
	"lib-manager/pkg/types"
)

// To submit registration form
func Register(res http.ResponseWriter, req *http.Request) {
	var errMsg types.ErrMsg
	username := req.FormValue("username")
	password := req.FormValue("password")
	reEnterPass := req.FormValue("reEnterPass") 
	adminAccess := req.FormValue("adminAccess")

	fmt.Println(username , password , reEnterPass , adminAccess)
	status := models.RegisterUser(username , password , reEnterPass , adminAccess)
	errMsg.Msg = status
	if(status == "OK"){
		fmt.Println("+++++++++++++++++++++",status)
		t := views.StartPage()
		res.WriteHeader(http.StatusOK)
		t.Execute(res,status)
	}else{
		fmt.Println("+++++++++++++++++++++",status)

		t := views.Register()
		res.WriteHeader(http.StatusOK)
		t.Execute(res,errMsg)
	}


}

// To open register page
func RegisterPage(res http.ResponseWriter, req *http.Request){
	t := views.Register()
	res.WriteHeader(http.StatusOK)
	t.Execute(res,nil )
}