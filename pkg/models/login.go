package models

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"fmt"
	"lib-manager/pkg/types"
	"lib-manager/pkg/views"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type ViewData struct {
	Data string
}

func LoginUser(writer http.ResponseWriter, request *http.Request, db *sql.DB, userName, password string) {

	var user types.UserDetail
	var userId int
	var errorMessage types.ErrorMessage

	// Check for password authentication
	rows, _ := db.Query("SELECT id ,userName , hash , admin FROM users WHERE userName = ?", userName)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.UserName, &user.Hash, &user.Admin)

		if err != nil {
			errorMessage.Message = "Incorrect ID or Password"
		} else {
			errr := authenticate(writer, request, userName, password, user)

			if errr == "OK" {
				randomBytes := make([]byte, 20)
				rand.Read(randomBytes)
				sessionID := hex.EncodeToString(randomBytes) // Convert the byte buffer to a hexadecimal string

				// Set the cookie in the response header
				cookie := http.Cookie{
					Name:     "SessionID",
					Value:    sessionID,
					Expires:  time.Now().Add(24 * time.Hour),
					HttpOnly: true,
				}
				http.SetCookie(writer, &cookie)
				fmt.Println("Cookie has been set") // Send a response to the client

				_ = db.QueryRow("SELECT userId FROM cookie").Scan(&userId)

				if userId == 0 {
					db.Exec("INSERT INTO cookie (sessionId, userId) VALUES (?, ?)", sessionID, user.Id) //Insert a row in cookie table
				} else {
					db.Exec("UPDATE cookie SET sessionId = ?, userId = ?", sessionID, user.Id) // Update a row in cookie table with new sessionId
				}
				if user.Admin == 1 {
					http.Redirect(writer, request, "/admin", http.StatusSeeOther) //Admin redirect

				} else {
					http.Redirect(writer, request, "/client", http.StatusSeeOther) //Client redirect
				}

			} else {
				errorMessage.Message = "Incorrect ID or Password"
				t := views.LogIn()
				writer.WriteHeader(http.StatusOK)
				t.Execute(writer, errorMessage)

			}
		}
	}

}

func authenticate(writer http.ResponseWriter, request *http.Request, username string, password string, user types.UserDetail) string {

	err := bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(password))

	if err == nil {
		fmt.Println("Logging into " + user.UserName)
		return "OK"

	} else {
		fmt.Println("Incorrect ID or Password")
		return "Incorrect ID or Password"
	}
}

// Remove Session From Database and redirect to Login page
func Logout(writer http.ResponseWriter, request *http.Request, db *sql.DB, userId int) {

	// Empty cookie and kill the session
	request.Header.Set("Cookie", "")
	db.Exec("DELETE FROM cookie WHERE userId = ?", userId)

	http.Redirect(writer, request, "/login", http.StatusSeeOther)
}
