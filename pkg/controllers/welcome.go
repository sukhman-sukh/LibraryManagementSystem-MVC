package controllers

import (
	"fmt"
	"lib-manager/pkg/models"
	"net/http"
	// "context"
	"lib-manager/pkg/views"
)

type key int 
const (
	adminAuthKey key = iota
	userID key = iota
	userName string = ""
)

func Welcome(writer http.ResponseWriter, request *http.Request) {
	
	

	status, _,_,admin:= models.Middleware(writer,request)

	
		if(status == "OK"){
		// admin, ok := request.Context().Value(adminAuthKey).(int)

		fmt.Println("============================")
		fmt.Println(admin)
		fmt.Println("============================")
		

		if admin == 1 {
			fmt.Println(writer, "Admin Auth Granted")
			t := views.GetAdmin()
			writer.WriteHeader(http.StatusOK)
			t.Execute(writer, nil)

		} else {
			fmt.Println(writer, "Not an Admin")
			t := views.GetClient()
			writer.WriteHeader(http.StatusOK)
			t.Execute(writer, nil)
		}
	}else{

		fmt.Println("============================")
		fmt.Println("Cookie Not Set")
		fmt.Println("============================")		

		// LogIn(writer , request)
		t := views.StartPage()
		writer.WriteHeader(http.StatusOK)
		t.Execute(writer, nil)
	}
}
