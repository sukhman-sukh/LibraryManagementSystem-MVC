package controllers

import (
	// "encoding/json"
	// "fmt"
	// "context"
	"net/http"
	"lib-manager/pkg/models"
	"lib-manager/pkg/views"
	"fmt"
)

// To open login page
func LogIn(res http.ResponseWriter, req *http.Request){
	t := views.LogIn()
	res.WriteHeader(http.StatusOK)
	t.Execute(res,nil )
}


// To open login page with errors
func LogInError(res http.ResponseWriter, req *http.Request , error string){
	t := views.LogIn()
	res.WriteHeader(http.StatusOK)
	t.Execute(res,error )
}

//  Check credentials for login
func ChecklogIn(res http.ResponseWriter, req *http.Request) {
	username := req.FormValue("username")
	password := req.FormValue("password")
	fmt.Println("testing", req)

	// models.loginUser(res ,req, username , password)
	models.LoginUser(res ,req, username , password)
}

// Logout And End Session
func Logout(res http.ResponseWriter, req *http.Request){
	status,userID,_,_ := models.Middleware(res,req)

	if(status == "OK"){
		models.Logout(res , req , userID)	
	}else{
		http.Redirect(res, req, "/login", http.StatusSeeOther)
	}


}