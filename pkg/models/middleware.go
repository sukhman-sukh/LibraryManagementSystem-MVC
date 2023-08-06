package models

import (
	_ "github.com/go-sql-driver/mysql"
	// "golang.org/x/crypto/bcrypt"
	"lib-manager/pkg/views"
	"fmt"
	// "database/sql"
	// "context"
	"net/http"

)


type key int 
const (
	admin key = iota
	userID key = iota
	userName string = ""
)


// Validating Cookies
// Returns "Cookie Not Set" when cookie is not there on user side   
// Returns "Cookie Was Altered On User Side" when session id on server and user not matches
// Returns "OK" If cookie is validated

func Middleware(res http.ResponseWriter, req *http.Request ) (string,int , string ,int){
	
	// var sessionId , userName string
	// var userId , admin int
	var sessionID string
	var userId int
	var admin int
	var userName string

	// Connect To Database
	db, err := Connection()
	if err != nil {
		fmt.Println("Error in connecting to database")
		// return "", "", errMsg
	}
	defer db.Close()
	pingErr := db.Ping()
    if pingErr != nil {
        fmt.Println(pingErr)
    }
    fmt.Println("Connected!")

	// Validaring cookie
	cookieId := req.Header.Get("Cookie")


	if(len(cookieId) < 10 || cookieId == "SessionID=0000000000000000000000000000000000000000" ){
		fmt.Println("Cookie not set")
		return "Cookie Not Set",0,"",0
	}else{
		cookieId = req.Header.Get("Cookie")[10:]
	}
	fmt.Println("============================")
	fmt.Println(cookieId)
	fmt.Println("============================")

	// if err 
	// cookie, err := request.Cookie("SessionID") 
	// query := fmt.Sprintf()
	// err = db.Query(query).Scan(&sessionID , &userId ,&admin , &username)
	rows, err := db.Query("SELECT sessionId, userId FROM cookie")
	if err != nil {
		fmt.Println("Error querying the database:")
	}
	defer rows.Close()
	if(IsDbEmpty("cookie" , db)){
		fmt.Println("NOthing in sessionID")
		t := views.LogIn()
		res.WriteHeader(http.StatusOK)
		t.Execute(res, nil)
	}	
	fmt.Println("sds1")
	for rows.Next() {

		fmt.Println("sds2")
		fmt.Println(rows)
		rows.Scan(&sessionID , &userId )
		// break
		// if err != nil {
		// 	// log.Fatal("Error scanning rows:", err)
		// }

	}
	rows, err = db.Query("SELECT  admin, userName FROM users WHERE  users.id = ?",userId)
	if err != nil {
		fmt.Println("Error querying the database:")
	}
	defer rows.Close()
	if(IsDbEmpty("cookie" , db)){
		fmt.Println("NOthing in sessionID")
		t := views.LogIn()
		res.WriteHeader(http.StatusOK)
		t.Execute(res, nil)
	}	
	// fmt.Println("sds1")
	for rows.Next() {
		// fmt.Println("sds2")
		// fmt.Println(rows)
		rows.Scan( &admin , &userName)
		// break
		// if err != nil {
		// 	// log.Fatal("Error scanning rows:", err)
		// }

	}	

	// if err := rows.Err(); err != nil {
	// 	// log.Fatal(err)
	// }
	fmt.Println(sessionID , userId , userName , admin)

	if err != nil {
		fmt.Println("Internal Server Error")
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return "Cookie Not Set",0,"",0
	}

	fmt.Println("================"+sessionID)
	if cookieId != sessionID {
		fmt.Println("Cookie Was Altered On User Side")
		// http.Redirect(res, req , "/", http.StatusSeeOther)
		return "Cookie Was Altered On User Side",0,"",0
		
	} else{
		fmt.Println("Inside ")
		// if(admin == 1){
		// 	ctx := context.WithValue(req.Context(), admin, 1)
		// 	req = req.WithContext(ctx)
		// 	// req.Body.adminAuth = 1
		// }else{
		// 	ctx := context.WithValue(req.Context(), admin, 0)
		// 	req = req.WithContext(ctx)
		// }
		// ctx := context.WithValue(req.Context(), userID, userId)
		// req = req.WithContext(ctx)
		// ctx = context.WithValue(req.Context(), userName, userName)
		// req = req.WithContext(ctx)

		// UserId:= req.Context().Value(userID).(int)
		// UserName:= req.Context().Value(userName).(string)
		// Admin:= req.Context().Value(admin).(int)
		fmt.Println("================================")
		fmt.Println(userId)
		fmt.Println(admin)
		fmt.Println(userName)
		fmt.Println("================================")
		

	}
	return "OK", userId, userName , admin
}

// // Check for authentication of admin
// func IsAdmin(res http.ResponseWriter, req *http.Request )int{
// 	Admin:= req.Context().Value(admin).(int)

// 	if(Admin ==1){
// 		return 1
// 	}else{
// 	return 0
// }
// }
