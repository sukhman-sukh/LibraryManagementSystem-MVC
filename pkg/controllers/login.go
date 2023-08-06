package controllers

import (
	// "encoding/json"
	"fmt"
	"context"
	"net/http"
	"lib-manager/pkg/models"
	"lib-manager/pkg/views"
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

	// models.loginUser(res ,req, username , password)
	models.LoginUser(res ,req, username , password)
}

// Logout And End Session
func Logout(res http.ResponseWriter, req *http.Request){
	status, userId , userName , admin := models.Middleware(res,req)

	// Setting Values To Req Objects
	if(admin == 1){
		ctx := context.WithValue(req.Context(), admin, 1)
		req = req.WithContext(ctx)
		
	}else{
		ctx := context.WithValue(req.Context(), admin, 0)
		req = req.WithContext(ctx)
	}
	ctx := context.WithValue(req.Context(), userID, userId)
	req = req.WithContext(ctx)
	ctx = context.WithValue(req.Context(), userName, userName)
	req = req.WithContext(ctx)

	UserId:= req.Context().Value(userID).(int)
	UserName:= req.Context().Value(userName).(string)
	Admin:= req.Context().Value(admin).(int)
	fmt.Println("================================")
	fmt.Println(UserId)
	fmt.Println(Admin)
	fmt.Println(UserName)
	fmt.Println("================================")


	if(status == "OK"){
		models.Logout(res , req)	
	}else{
		http.Redirect(res, req, "/login", http.StatusSeeOther)
	}


}