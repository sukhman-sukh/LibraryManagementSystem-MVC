package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(db *sql.DB, username, password, reEnterPass, requestAccess string) string {
	var userId int

	// Check for no two Users with same UserName
	rows, err := db.Query("SELECT * FROM users WHERE userName = ?", username)
	if err != nil {
		return "Row not found"
	}
	if rows.Next() {
		return "Username is not unique"
	}
	defer rows.Close()

	hash2, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	done := make(chan bool)
	go checkAdminList(db, done) // Check If there exist a admin in database or not
	admin := <-done

	if requestAccess == "adminAccess" {
		if admin == false {
			db.Exec("INSERT INTO users (userName, hash, admin) VALUES (?,?,?)", username, hash2, 1) // No admin in database . Make this user Admin
		} else {
			db.Exec("INSERT INTO users (userName, hash, admin) VALUES (?,?,?)", username, hash2, 0) //Add this user to adminReq table and make it client
			err = db.QueryRow("select id from users where userName = ?", username).Scan(&userId)
			if err != nil {
				return ""
			}
			db.Exec("INSERT INTO adminReq ( userId, status) VALUES(?,?)", userId, 0)
		}
	} else {
		db.Exec("INSERT INTO users (userName, hash, admin) VALUES (?,?,?)", username, hash2, 0) //Client user withour adminAccess request
	}
	return ""
}

func checkAdminList(db *sql.DB, done chan bool) {
	admin := false

	rows, _ := db.Query("SELECT admin FROM users")
	defer rows.Close()

	for rows.Next() {
		var isAdmin int
		if err := rows.Scan(&isAdmin); err != nil {
			done <- false
			return
		}
		if isAdmin == 1 {
			admin = true
		}
	}
	done <- admin
	return
}
