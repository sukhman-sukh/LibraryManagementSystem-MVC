package controllers

import (
	"net/http"
	"lib-manager/pkg/models"
	"lib-manager/pkg/views"
	"lib-manager/pkg/types"
)

// To open login page
func LogIn(res http.ResponseWriter, req *http.Request){
	t := views.LogIn()
	res.WriteHeader(http.StatusOK)
	t.Execute(res,nil )
}

//  Check credentials for login
func ChecklogIn(res http.ResponseWriter, req *http.Request) {
	username := req.FormValue("username")
	password := req.FormValue("password")

	db, err := models.Connection()
    var errMsg types.ErrMsg
    if err!= nil {
        errMsg.Msg = "Error in connecting to database"
    }
    defer db.Close()

	models.LoginUser(res ,req,db, username , password)

}

// Logout And End Session
func Logout(res http.ResponseWriter, req *http.Request){
	db, err := models.Connection()
    var errMsg types.ErrMsg
    if err!= nil {
        errMsg.Msg = "Error in connecting to database"
    }
    defer db.Close()

	status,userID,_,_ := models.Middleware(res,req,db)

	if(status == "OK"){
		
		models.Logout(res , req ,db, userID)	
	}else{
		http.Redirect(res, req, "/login", http.StatusSeeOther)
	}


}