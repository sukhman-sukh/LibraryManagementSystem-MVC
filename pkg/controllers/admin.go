package controllers

import (
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

	db, err := models.Connection()
    var errMsg types.ErrMsg
    if err!= nil {
        errMsg.Msg = "Error in connecting to database"
    }
    defer db.Close()

	status, userId , userName , admin := models.Middleware(res,req , db)

	if(status == "OK"){
		fmt.Println("Yes Status is OK ")
		if(admin ==1){	

			
			_ ,books := models.GetBooks(db)
			_, reqBook := models.GetReqBooks(db, userId )
			_, issuedBooks := models.GetIssuedBooks(db, userId , admin)
			_ , adminReq := models.GetAdminReq(db)
			data := types.Data{
				UserName: userName,
				Books:     books,
				ReqBook:  reqBook,
				AdminReq:adminReq,
				IssuedBooks: issuedBooks,
			}
		
			fmt.Println(data)
			
			t := views.GetAdmin()
			res.WriteHeader(http.StatusOK)
			t.Execute(res, data)
		}else{
			http.Redirect(res, req, "/error403", http.StatusSeeOther)	
		}
	}else{
		http.Redirect(res, req, "/login", http.StatusSeeOther)
	}

}

func AdminCheckinSubmit(res http.ResponseWriter, req *http.Request){

	db, err := models.Connection()
    var errMsg types.ErrMsg
    if err!= nil {
        errMsg.Msg = "Error in connecting to database"
    }
    defer db.Close()

	status, _ , _ , admin := models.Middleware(res,req , db)


	if(status == "OK"){
		if(admin ==1){	

			reqId := req.FormValue("reqId");
			_ = models.AdminCheckin(res, req ,db , reqId)

			if(status == "OK"){
				http.Redirect(res, req, "/admin", http.StatusSeeOther)
			}
		}else{
			http.Redirect(res, req, "/error403", http.StatusSeeOther)	
		}
	}else{
		http.Redirect(res, req, "/", http.StatusSeeOther)
	}
}

func AdminAdd(res http.ResponseWriter, req *http.Request){
	t := views.AdminAdd()
	res.WriteHeader(http.StatusOK)
	t.Execute(res, nil)
}

func AdminAddSubmit(res http.ResponseWriter, req *http.Request){

	db, err := models.Connection()
    var errMsg types.ErrMsg
    if err!= nil {
        errMsg.Msg = "Error in connecting to database"
    }
    defer db.Close()
	status, _ , _ , admin := models.Middleware(res,req, db)


	if(status == "OK"){
		if(admin ==1){	

			bookname := req.FormValue("bookname");
			Author := req.FormValue("Author")
			Copies := req.FormValue("Copies");

			status := models.AdminAdd(res, req ,db , bookname, Author, Copies)
			if(status == "OK"){
				http.Redirect(res, req, "/admin", http.StatusSeeOther)
			}
		}else{
			http.Redirect(res, req, "/error403", http.StatusSeeOther)	
		}
	}else{
		http.Redirect(res, req, "/login", http.StatusSeeOther)
	}
}


func AdminCheckoutSubmit(res http.ResponseWriter, req *http.Request){
	db, err := models.Connection()
    var errMsg types.ErrMsg
    if err!= nil {
        errMsg.Msg = "Error in connecting to database"
    }
    defer db.Close()
	status, _ , _ , admin := models.Middleware(res,req,db)

	if(status == "OK"){
		if(admin ==1){	

			reqId := req.FormValue("reqId");
			status := models.AdminCheckout(res, req ,db, reqId)

			if(status == "OK"){
				http.Redirect(res, req, "/admin", http.StatusSeeOther)
			}
		}else{
			http.Redirect(res, req, "/error403", http.StatusSeeOther)	
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

	db, err := models.Connection()
    var errMsg types.ErrMsg
    if err!= nil {
        errMsg.Msg = "Error in connecting to database"
    }
    defer db.Close()
	status, _ , _ , admin := models.Middleware(res,req,db)


	if(status == "OK"){
		if(admin ==1){	

			bookId := req.FormValue("bookId");
			copies := req.FormValue("Copies");

			status := models.AdminRemove(res, req ,db, bookId, copies)

			if(status == "OK"){
				http.Redirect(res, req, "/admin", http.StatusSeeOther)
			}
		}else{
			http.Redirect(res, req, "/error403", http.StatusSeeOther)	
		}
	}else{
		http.Redirect(res, req, "/login", http.StatusSeeOther)
	}
}


func AdminAccept(res http.ResponseWriter, req *http.Request){

	db, err := models.Connection()
    var errMsg types.ErrMsg
    if err!= nil {
        errMsg.Msg = "Error in connecting to database"
    }
    defer db.Close()

	status, _ , _ , admin := models.Middleware(res,req,db)

	if(status == "OK"){
		if(admin ==1){	

			reqId := req.FormValue("reqId");
			status := models.AdminAccept(res, req ,db , reqId)

			if(status == "OK"){
				http.Redirect(res, req, "/admin", http.StatusSeeOther)
			}
		}else{
			http.Redirect(res, req, "/error403", http.StatusSeeOther)	
		}
	}else{
		http.Redirect(res, req, "/login", http.StatusSeeOther)
	}
}

func AdminDeny(res http.ResponseWriter, req *http.Request){
	db, err := models.Connection()
    var errMsg types.ErrMsg
    if err!= nil {
        errMsg.Msg = "Error in connecting to database"
    }
    defer db.Close()

	status, _ , _ , admin := models.Middleware(res,req,db)


	if(status == "OK"){
		if(admin ==1){	

			reqId := req.FormValue("reqId");

			status := models.AdminAccept(res, req ,db , reqId)

			if(status == "OK"){
				http.Redirect(res, req, "/admin", http.StatusSeeOther)
			}
		}else{
			http.Redirect(res, req, "/error403", http.StatusSeeOther)	
		}
	}else{
		http.Redirect(res, req, "/login", http.StatusSeeOther)
	}
}


