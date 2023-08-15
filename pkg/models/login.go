package models

import (
	"fmt"
	"time"
	"net/http"
	"crypto/rand"
	"encoding/hex"
	"database/sql"
	"lib-manager/pkg/views"
	"golang.org/x/crypto/bcrypt"
	"lib-manager/pkg/types"
  _ "github.com/go-sql-driver/mysql"
	
)

type ViewData struct {
	Data string
}

func LoginUser(res http.ResponseWriter, req *http.Request , db *sql.DB, userName , password string){

	var user types.UserDetail
	var userId int
	var errMsg types.ErrMsg

	// Check for password authentication
	rows, _ := db.Query("SELECT id ,userName , hash , admin FROM users WHERE userName = ?", userName)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&user.Id ,&user.UserName , &user.Hash , &user.Admin )
		
		if (err != nil){
			errMsg.Msg = "Incorrect ID or Password"
		} else{
			errr := authenticate(res , req ,userName , password , user);
			
			if(errr == "OK"){
				randomBytes := make([]byte, 20)
				rand.Read(randomBytes)
				sessionID := hex.EncodeToString(randomBytes)	// Convert the byte buffer to a hexadecimal string

				// Set the cookie in the response header
				cookie := http.Cookie{
					Name:     "SessionID",
					Value:    sessionID,
					Expires:  time.Now().Add(24 * time.Hour), 
					HttpOnly: true,                           
				}
				http.SetCookie(res, &cookie)
				fmt.Println("Cookie has been set") // Send a response to the client

				_ = db.QueryRow("SELECT userId FROM cookie" ).Scan(&userId)

				if(userId == 0){
					db.Exec("INSERT INTO cookie (sessionId, userId) VALUES (?, ?)", sessionID, user.Id)		//Insert a row in cookie table 
				}else{
					db.Exec("UPDATE cookie SET sessionId = ?, userId = ?", sessionID, user.Id)				// Update a row in cookie table with new sessionId
				}
				if(user.Admin == 1){
					http.Redirect(res, req, "/admin", http.StatusSeeOther)						//Admin redirect

				}else{
					http.Redirect(res, req, "/client", http.StatusSeeOther)						//Client redirect
				}
			

			}else{
				errMsg.Msg = "Incorrect ID or Password"
				t := views.LogIn()
				res.WriteHeader(http.StatusOK)
				t.Execute(res,errMsg )
			
			}
		}
	}


}


func authenticate(res http.ResponseWriter, req *http.Request ,username string , password string , user types.UserDetail) (string){

	err := bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(password))
    
	if(err == nil){
		fmt.Println("Logging into "+ user.UserName)
		return "OK"
	
	}else{
		fmt.Println("Incorrect ID or Password")
		return "Incorrect ID or Password"
	}
}


// Remove Session From Database and redirect to Login page
func Logout(res http.ResponseWriter, req *http.Request , db *sql.DB, userId int ){

	// db, err := Connection()
	// var errMsg types.ErrMsg
	// if err != nil {
	// 	errMsg.Msg = "Error in connecting to database"
	// }
	// defer db.Close()

// Empty cookie and kill the session
	req.Header.Set("Cookie", "" )
	db.Exec("DELETE FROM cookie WHERE userId = ?",userId )

	http.Redirect(res, req, "/login", http.StatusSeeOther)
}