package controllers

import (
	// "encoding/json"
	"context"
	"fmt"
	"net/http"
	"lib-manager/pkg/views"
	"lib-manager/pkg/models"
)


func GetAdmin(res http.ResponseWriter, req *http.Request) {

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
	

	// var admin int
	if(status == "OK"){
		fmt.Println("Yes Status is OK ")
		// Admin:= req.Context().Value(admin).(int)
		// fmt.Println(Admin)
		if(models.IsAdmin(res , req) ==1){	
			fmt.Println("adminnnnnn")		
			t := views.GetAdmin()
			res.WriteHeader(http.StatusOK)
			t.Execute(res, nil)
		}else{
			http.Redirect(res, req, "/login", http.StatusSeeOther)
		}
	}else{
		http.Redirect(res, req, "/login", http.StatusSeeOther)
	}

}

func AdminCheckin(res http.ResponseWriter, req *http.Request){
	t := views.Checkin()
	res.WriteHeader(http.StatusOK)
	t.Execute(res, nil)
}