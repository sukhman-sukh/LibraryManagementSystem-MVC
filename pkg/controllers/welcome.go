package controllers

import (
	"lib-manager/pkg/models"
	"net/http"
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
			if admin == 1 {
				http.Redirect(writer, request, "/admin", http.StatusSeeOther)

			} else {
				http.Redirect(writer, request, "/client", http.StatusSeeOther)
			}
		}else{
			t := views.StartPage()
			writer.WriteHeader(http.StatusOK)
			t.Execute(writer, nil)
		}
}
