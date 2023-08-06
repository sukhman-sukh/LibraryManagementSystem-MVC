package controllers

import (
	// "encoding/json"
	// "context"
	"fmt"
	"net/http"
	"lib-manager/pkg/views"
	"lib-manager/pkg/models"
	"lib-manager/pkg/types"
)



// Format

// bookStatus{
//     1 : "Requested Check-In"
//     0: "Check-out"
//      -1: Requested- check-out
// }

func GetAdmin(res http.ResponseWriter, req *http.Request) {

	status, _ , userName , admin := models.Middleware(res,req)


	if(status == "OK"){
		fmt.Println("Yes Status is OK ")
		// Admin:= req.Context().Value(admin).(int)
		// fmt.Println(Admin)
		if(admin ==1){	
			fmt.Println("adminnnnnn")

			_ ,books := models.GetBooks(res ,req)
			_, reqBook := models.GetReqBooks(res ,req)
			_ , adminReq := models.GetAdminReq(res ,req)
			// username: userName, data: books, reqdata: reqBook, adminReq: adminReq
			// var data types.Data 
			data := types.Data{
				UserName: userName,
				Books:     books,
				ReqBook:  reqBook,
				AdminReq: adminReq,
			}
			
			// data.UserName = userName
			// data.Books=books
            // data.ReqBook=reqBook
        	// data.AdminReq = adminReq
            
			fmt.Println(data)
			
			t := views.GetAdmin()
			res.WriteHeader(http.StatusOK)
			t.Execute(res, data)
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
	
	status, _ , _ , admin := models.Middleware(res,req)


	if(status == "OK"){
		if(admin ==1){	
			fmt.Println("adminnnnnn")
			// username := req.FormValue("username")

			reqId := req.FormValue("reqId");
			// copies := req.FormValue("Copies");

			_ = models.AdminCheckin(res, req , reqId)

			if(status == "OK"){
				http.Redirect(res, req, "/admin", http.StatusSeeOther)
			}
		}else{
			http.Redirect(res, req, "/login", http.StatusSeeOther)
		}
	}else{
		http.Redirect(res, req, "/", http.StatusSeeOther)
	}

	// http.Redirect(res, req, "/login", http.StatusSeeOther)
}

func AdminAdd(res http.ResponseWriter, req *http.Request){
	t := views.AdminAdd()
	res.WriteHeader(http.StatusOK)
	t.Execute(res, nil)
}

func AdminAddSubmit(res http.ResponseWriter, req *http.Request){
	status, _ , _ , admin := models.Middleware(res,req)


	if(status == "OK"){
		if(admin ==1){	
			fmt.Println("adminnnnnn")
			// username := req.FormValue("username")

			bookname := req.FormValue("bookname");
			Author := req.FormValue("Author")
			Copies := req.FormValue("Copies");

			status := models.AdminAdd(res, req , bookname, Author, Copies)
			if(status == "OK"){
				http.Redirect(res, req, "/admin", http.StatusSeeOther)
			}
		}else{
			http.Redirect(res, req, "/login", http.StatusSeeOther)
		}
	}else{
		http.Redirect(res, req, "/login", http.StatusSeeOther)
	}
}


func AdminCheckout(res http.ResponseWriter, req *http.Request){
	t := views.AdminCheckout()
	res.WriteHeader(http.StatusOK)
	t.Execute(res, nil)
}

func AdminCheckoutSubmit(res http.ResponseWriter, req *http.Request){
	status, _ , _ , admin := models.Middleware(res,req)

	if(status == "OK"){
		if(admin ==1){	
			fmt.Println("adminnnnnn")
			// username := req.FormValue("username")

			reqId := req.FormValue("reqId");

			status := models.AdminCheckout(res, req , reqId)

			if(status == "OK"){
				http.Redirect(res, req, "/admin", http.StatusSeeOther)
			}
		}else{
			http.Redirect(res, req, "/login", http.StatusSeeOther)
		}
	}else{
		http.Redirect(res, req, "/login", http.StatusSeeOther)
	}
}

func AdminRemove(res http.ResponseWriter, req *http.Request){
	t := views.AdminRemove()
    res.WriteHeader(http.StatusOK)
    t.Execute(res, nil)
}


func AdminRemoveSubmit(res http.ResponseWriter, req *http.Request){

	status, _ , _ , admin := models.Middleware(res,req)


	if(status == "OK"){
		if(admin ==1){	
			fmt.Println("adminnnnnn")
			// username := req.FormValue("username")

			bookId := req.FormValue("bookId");
			copies := req.FormValue("Copies");

			status := models.AdminRemove(res, req , bookId, copies)

			if(status == "OK"){
				http.Redirect(res, req, "/admin", http.StatusSeeOther)
			}
		}else{
			http.Redirect(res, req, "/login", http.StatusSeeOther)
		}
	}else{
		http.Redirect(res, req, "/login", http.StatusSeeOther)
	}
}

func AdminChoose(res http.ResponseWriter, req *http.Request){
	t := views.AdminChoose()
    res.WriteHeader(http.StatusOK)
    t.Execute(res, nil)
}

func AdminAccept(res http.ResponseWriter, req *http.Request){

	status, _ , _ , admin := models.Middleware(res,req)

	if(status == "OK"){
		if(admin ==1){	
			fmt.Println("adminnnnnn")
			// username := req.FormValue("username")

			reqId := req.FormValue("reqId");

			status := models.AdminAccept(res, req , reqId)

			if(status == "OK"){
				http.Redirect(res, req, "/admin", http.StatusSeeOther)
			}
		}else{
			http.Redirect(res, req, "/login", http.StatusSeeOther)
		}
	}else{
		http.Redirect(res, req, "/login", http.StatusSeeOther)
	}
}

func AdminDeny(res http.ResponseWriter, req *http.Request){
	status, _ , _ , admin := models.Middleware(res,req)


	if(status == "OK"){
		if(admin ==1){	
			fmt.Println("adminnnnnn")

			reqId := req.FormValue("reqId");

			status := models.AdminAccept(res, req , reqId)

			if(status == "OK"){
				http.Redirect(res, req, "/admin", http.StatusSeeOther)
			}
		}else{
			http.Redirect(res, req, "/login", http.StatusSeeOther)
		}
	}else{
		http.Redirect(res, req, "/login", http.StatusSeeOther)
	}
}


