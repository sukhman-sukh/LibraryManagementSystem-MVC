package models

import (
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"lib-manager/pkg/types"
	// "lib-manager/pkg/controllers"
	"fmt"
	"net/http"
	// "html/template"
	"lib-manager/pkg/views"
	"time"
	// "database/sql"
	// "context"
	"crypto/rand"
	"encoding/hex"
)

type ViewData struct {
	Data string
}

func LoginUser(res http.ResponseWriter, req *http.Request , userName , password string){

	// fmt.Println("test1", req);
	
	db, err := Connection()
	var errMsg types.ErrMsg
	if err != nil {
		errMsg.Msg = "Error in connecting to database"
		// return "", "", errMsg
	}
	defer db.Close()
	// fmt.Println(userName)

	rows, err := db.Query("SELECT id ,userName , hash , admin FROM users WHERE userName = ?", userName)
	if err != nil {
		errMsg.Msg = "Username Does Not Exist"

		// return "","",errMsg
	}
	defer rows.Close()

	// fmt.Println(rows)
	var user types.UserDetail

	for rows.Next() {
		fmt.Println("INside rows")
		err := rows.Scan(&user.Id ,&user.UserName , &user.Hash , &user.Admin )
		fmt.Println(user)
		
		if (err != nil){
			errMsg.Msg = "Incorrect ID or Password"
			// return "","",errMsg
		} else{
			errr := authenticate(res , req ,userName , password , user);
			
			if(errr == "OK"){
				randomBytes := make([]byte, 20)
				rand.Read(randomBytes)
				
				// fmt.Println("===========================")
				// fmt.Println(randomBytes)
				// fmt.Println("===========================")

				// Convert the byte buffer to a hexadecimal string
				sessionID := hex.EncodeToString(randomBytes)
				fmt.Println(sessionID)
				cookie := http.Cookie{
					Name:     "SessionID",
					Value:    sessionID,
					Expires:  time.Now().Add(24 * time.Hour), // Set cookie expiration time
					HttpOnly: true,                           
				}
			
				// Set the cookie in the response header
				http.SetCookie(res, &cookie)

				// Send a response to the client
				// fmt.Fprintln(res, "Cookie has been set.")
				fmt.Println(req.Header.Get("Cookie"))
				// fmt.Println(req.Cookie("SessionID"))
				fmt.Println("Cookie has been set")

				var userId int
				err = db.QueryRow("SELECT userId FROM cookie" ).Scan(&userId)
				fmt.Println("====================================================================",user.Id)
				if err != nil {}
				if(userId == 0){
					fmt.Println("Empty table of cookie")	
						// No record found, insert a new row
					db.Exec("INSERT INTO cookie (sessionId, userId) VALUES (?, ?)", sessionID, user.Id)
				}else{
					fmt.Println("Table has some values")	
					db.Exec("UPDATE cookie SET sessionId = ?, userId = ?", sessionID, user.Id)
				}
				if(user.Admin == 1){
					fmt.Println("admin123")
					// fmt.Println(req)
					http.Redirect(res, req, "/admin", http.StatusSeeOther)

				}else{
					fmt.Println("not admin")
					http.Redirect(res, req, "/client", http.StatusSeeOther)
				}
			

			}else{
				errMsg.Msg = "Incorrect ID or Password"
				// fmt.Fprintln(errMsg.Msg)
				t := views.LogIn()
				res.WriteHeader(http.StatusOK)
				t.Execute(res,errMsg )
			
			}
		}
	}


}


func authenticate(res http.ResponseWriter, req *http.Request ,username string , password string , user types.UserDetail) (string){

	fmt.Println("Inside authentication block")

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
func Logout(res http.ResponseWriter, req *http.Request , userId int ){

	db, err := Connection()
	var errMsg types.ErrMsg
	if err != nil {
		errMsg.Msg = "Error in connecting to database"
		// return "", "", errMsg
	}
	defer db.Close()

// Empty cookie and kill the session
	req.Header.Set("Cookie", "" )
	fmt.Println( req.Header.Get("Cookie"))
	db.Exec("DELETE FROM cookie WHERE userId = ?",userId )

	http.Redirect(res, req, "/login", http.StatusSeeOther)
}