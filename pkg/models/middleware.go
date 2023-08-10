package models

import (
	"fmt"
	"net/http"
	"database/sql"
	"lib-manager/pkg/types"
  _ "github.com/go-sql-driver/mysql"



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

func Middleware(res http.ResponseWriter, req *http.Request, db *sql.DB ) (string,int , string ,int){
	
	var sessionInfo types.ValidateCookie
	
	// Connect To Database
	// db, err := Connection()
	// if err != nil {
	// 	fmt.Println("Error in connecting to database")
	// }
	// defer db.Close()

	// Validaring cookie
	cookieId := req.Header.Get("Cookie")
	if(len(cookieId) < 10 || cookieId == "SessionID=0000000000000000000000000000000000000000" ){		 //NO Cookies ON User Side
		fmt.Println("Cookie not set")
		return "Cookie Not Set",0,"",0
	}else{
		cookieId = req.Header.Get("Cookie")[10:]
	}

	rows, _ := db.Query("SELECT sessionId, userId FROM cookie")
	for rows.Next() {
		rows.Scan(&sessionInfo.SessionID , &sessionInfo.UserId )
	}

	rows, _ = db.Query("SELECT  admin, userName FROM users WHERE id = ?",sessionInfo.UserId)
	for rows.Next() {
		rows.Scan( &sessionInfo.Admin , &sessionInfo.Username)
	}	
	if(sessionInfo.SessionID == "" || sessionInfo.Username=="" || sessionInfo.UserId == 0){		// Session database table is empty
		return "Nothing in sessionID",0,"",0							
	}

	if cookieId != sessionInfo.SessionID {														// Cookie Id has been tempered
		fmt.Println("Cookie Was Altered On User Side")
		return "Cookie Was Altered On User Side",0,"",0
		
	}
	return "OK", sessionInfo.UserId, sessionInfo.Username , sessionInfo.Admin
}

