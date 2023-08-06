package controllers

import (
	"fmt"
	"lib-manager/pkg/models"
	"net/http"
	"context"
	"lib-manager/pkg/views"
)

type key int 
const (
	adminAuthKey key = iota
	userID key = iota
	userName string = ""
)

func Welcome(writer http.ResponseWriter, request *http.Request) {
	
	

	status, userId , userName , admin := models.Middleware(writer,request)

	// Setting Values To Req Objects
	if(admin == 1){
		ctx := context.WithValue(request.Context(), admin, 1)
		request = request.WithContext(ctx)
		
	}else{
		ctx := context.WithValue(request.Context(), admin, 0)
		request = request.WithContext(ctx)
	}
	ctx := context.WithValue(request.Context(), userID, userId)
	request = request.WithContext(ctx)
	ctx = context.WithValue(request.Context(), userName, userName)
	request = request.WithContext(ctx)

	UserId:= request.Context().Value(userID).(int)
	UserName:= request.Context().Value(userName).(string)
	Admin:= request.Context().Value(admin).(int)
	fmt.Println("================================")
	fmt.Println(UserId)
	fmt.Println(Admin)
	fmt.Println(UserName)
	fmt.Println("================================")


		if(status == "OK"){
		admin, ok := request.Context().Value(adminAuthKey).(int)

		fmt.Println("============================")
		fmt.Println(admin)
		fmt.Println("============================")
		if !ok {
			// Value not found or not of the expected type
			http.Error(writer, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if admin == 1 {
			fmt.Fprintln(writer, "Admin Auth Granted")
			t := views.GetAdmin()
			writer.WriteHeader(http.StatusOK)
			t.Execute(writer, nil)

		} else {
			fmt.Fprintln(writer, "Not an Admin")
			t := views.GetClient()
			writer.WriteHeader(http.StatusOK)
			t.Execute(writer, nil)
		}
	}else{

		fmt.Println("============================")
		fmt.Println("Cookie Not Set")
		fmt.Println("============================")		

		// LogIn(writer , request)
		t := views.LogIn()
		writer.WriteHeader(http.StatusOK)
		t.Execute(writer, nil)
	}
}
