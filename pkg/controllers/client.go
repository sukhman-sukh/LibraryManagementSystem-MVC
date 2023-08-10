package controllers

import (
	"net/http"
	"lib-manager/pkg/models"
	"lib-manager/pkg/views"
	"lib-manager/pkg/types"
)

// Format

// bookStatus{
//     1 : "Requested Check-In"
//     0: "Check-out"
//      -1: Requested- check-out
// }
func GetClient(res http.ResponseWriter, req *http.Request) {

	db, err := models.Connection()
    var errMsg types.ErrMsg
    if err!= nil {
        errMsg.Msg = "Error in connecting to database"
    }
    defer db.Close()

	status, userID , _ , admin := models.Middleware(res,req ,db)

	if(status == "OK"){
	_ ,books := models.GetBooks(db)
	_, reqBook := models.GetReqBooks(db, userID , admin)

	data := types.Data{
		UserName: userName,
		Books:     books,
		ReqBook:  reqBook,
		AdminReq:nil,
	}

	
	t := views.GetClient()
	res.WriteHeader(http.StatusOK)
	t.Execute(res, data)}

}

func Checkout(res http.ResponseWriter, req *http.Request){
	t := views.Checkout()
    res.WriteHeader(http.StatusOK)
    t.Execute(res, nil)
}


func CheckoutSubmit(res http.ResponseWriter, req *http.Request){

	db, err := models.Connection()
    var errMsg types.ErrMsg
    if err!= nil {
        errMsg.Msg = "Error in connecting to database"
    }
    defer db.Close()

	status, userID , _ , _ := models.Middleware(res,req,db)
	bookId := req.FormValue("bookId");

	status = models.Checkout(res, req ,db, bookId , userID)

	if(status == "OK"){
		http.Redirect(res, req, "/", http.StatusSeeOther)
	}

}


func Checkin(res http.ResponseWriter, req *http.Request){
	t := views.Checkin()
    res.WriteHeader(http.StatusOK)
    t.Execute(res, nil)
}


func CheckinSubmit(res http.ResponseWriter, req *http.Request){
	db, err := models.Connection()
    var errMsg types.ErrMsg
    if err!= nil {
        errMsg.Msg = "Error in connecting to database"
    }
    defer db.Close()

	reqId := req.FormValue("reqId");
	status := models.Checkin(res, req ,db, reqId)

	if(status == "OK"){
		http.Redirect(res, req, "/", http.StatusSeeOther)
	}

}

