package models

import (
	_ "github.com/go-sql-driver/mysql"
	// "golang.org/x/crypto/bcrypt"
	"lib-manager/pkg/types"
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
	

	var sessionInfo types.ValidateCookie
	
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

	rowCheck := rows
	if(rowCheck.Next()){
		fmt.Println("sfsdfdsfsdfsdfsd")
		// return "Nothing in sessionID",0,"",0
	}	
	fmt.Println("sds1")
	for rows.Next() {

		fmt.Println("sds2")
		fmt.Println(rows)
		rows.Scan(&sessionInfo.SessionID , &sessionInfo.UserId )
	

	}

	rows, err = db.Query("SELECT  admin, userName FROM users WHERE  users.id = ?",sessionInfo.UserId)
	if err != nil {
		fmt.Println("Error querying the database:")
	}
	defer rows.Close()
	// if(!(rows.Next())){
	// 	fmt.Println("NOthing in sessionID")
	// 	return "NOthing in sessionID",0,"",0
	// }	
	// fmt.Println("sds1")
	for rows.Next() {
		// fmt.Println("sds2")
		// fmt.Println(rows)
		rows.Scan( &sessionInfo.Admin , &sessionInfo.Username)
		// break
		// if err != nil {
		// 	// log.Fatal("Error scanning rows:", err)
		// }

	}	
	if(sessionInfo.SessionID == "" || sessionInfo.Username=="" || sessionInfo.Admin==0 || sessionInfo.UserId == 0){
		fmt.Println("Nothing in sessionID")
		fmt.Println(sessionInfo.SessionID)
		fmt.Println(sessionInfo.Username)
		fmt.Println(sessionInfo.Admin)
		fmt.Println(sessionInfo.UserId)
		return "Nothing in sessionID",0,"",0
	}
	// if err := rows.Err(); err != nil {
	// 	// log.Fatal(err)
	// }
	fmt.Println(sessionInfo.SessionID , sessionInfo.UserId , sessionInfo.Username , sessionInfo.Admin)

	if err != nil {
		fmt.Println("Internal Server Error")
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return "Cookie Not Set",0,"",0
	}

	fmt.Println("================"+sessionInfo.SessionID)
	if cookieId != sessionInfo.SessionID {
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
		fmt.Println(sessionInfo.UserId)
		fmt.Println(sessionInfo.Admin)
		fmt.Println(sessionInfo.Username)
		fmt.Println("================================")
		

	}
	return "OK", sessionInfo.UserId, sessionInfo.Username , sessionInfo.Admin
}

