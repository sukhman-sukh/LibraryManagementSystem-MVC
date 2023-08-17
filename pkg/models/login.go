package models

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"lib-manager/pkg/types"
	"lib-manager/pkg/views"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type ViewData struct {
	Data string
}

func LoginUser(writer http.ResponseWriter, request *http.Request, db *sql.DB, userName, password string) error {
	var user types.UserDetail
	var userId int
	var errorMessage string

	// Check for password authentication
	rows, err := db.Query("SELECT id ,userName , hash , admin FROM users WHERE userName = ?", userName)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.UserName, &user.Hash, &user.Admin)
		if err != nil {
			errorMessage = "Incorrect ID or Password"
		} else {
			response := authenticate(writer, request, userName, password, user)
			if response == "" {
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

				err = db.QueryRow("SELECT userId FROM cookie").Scan(&userId)
				if err != nil {
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
				errorMessage = "Incorrect ID or Password"
				var response types.Response
				response.Message = errorMessage
				t := views.Welcome("LogIn")
				writer.WriteHeader(http.StatusOK)
				t.Execute(writer, response)
			}
		}
	}
	return nil
}

func authenticate(writer http.ResponseWriter, request *http.Request, username string, password string, user types.UserDetail) string {
	err := bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(password))
	if err == nil {
		log.Println("Logging into " + user.UserName)
		return ""
	} else {
		return "Incorrect ID or Password"
	}
}

// Remove Session From Database and redirect to Login page
func Logout(writer http.ResponseWriter, request *http.Request, db *sql.DB, userId int) {
	request.Header.Set("Cookie", "") // Empty cookie and kill the session
	db.Exec("DELETE FROM cookie WHERE userId = ?", userId)
	http.Redirect(writer, request, "/login", http.StatusSeeOther)
}
