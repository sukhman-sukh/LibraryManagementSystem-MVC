package controllers

import (
	// "encoding/json"
	// "context"
	"fmt"
	"net/http"
	"lib-manager/pkg/views"
	"lib-manager/pkg/models"
)


func GetAdmin(res http.ResponseWriter, req *http.Request) {

	status, _ , _ , admin := models.Middleware(res,req)


	if(status == "OK"){
		fmt.Println("Yes Status is OK ")
		// Admin:= req.Context().Value(admin).(int)
		// fmt.Println(Admin)
		if(admin ==1){	
			fmt.Println("adminnnnnn")

			books := models.GetBooks()

			
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

func AdminCheckinSubmit(res http.ResponseWriter, req *http.Request){
	
	

	http.Redirect(res, req, "/login", http.StatusSeeOther)
}

func AdminAdd(res http.ResponseWriter, req *http.Request){
	t := views.AdminAdd()
	res.WriteHeader(http.StatusOK)
	t.Execute(res, nil)
}

func AdminAddSubmit(res http.ResponseWriter, req *http.Request){

}


func AdminCheckout(res http.ResponseWriter, req *http.Request){
	t := views.AdminCheckout()
	res.WriteHeader(http.StatusOK)
	t.Execute(res, nil)
}

func AdminCheckoutSubmit(res http.ResponseWriter, req *http.Request){

}

