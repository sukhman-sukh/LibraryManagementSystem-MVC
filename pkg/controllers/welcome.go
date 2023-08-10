package controllers

import (
	"lib-manager/pkg/models"
	"net/http"
	"lib-manager/pkg/views"
	"lib-manager/pkg/types"
)

type key int 
const (
	adminAuthKey key = iota
	userID key = iota
	userName string = ""
)

func Welcome(writer http.ResponseWriter, request *http.Request) {
	
	db, err := models.Connection()
    var errMsg types.ErrMsg
    if err!= nil {
        errMsg.Msg = "Error in connecting to database"
    }
    defer db.Close()

	status, _,_,admin:= models.Middleware(writer,request,db)

	
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
