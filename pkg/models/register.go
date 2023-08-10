package models

import (
	"fmt"
	"database/sql"
	"lib-manager/pkg/types"
	"golang.org/x/crypto/bcrypt"
  _ "github.com/go-sql-driver/mysql"
	
)



func RegisterUser( db *sql.DB, username , password , reEnterPass , reqAccess string) string{
	var userId int
	var errMsg types.ErrMsg
	// db, err := Connection()
	// 
	// if err != nil {
	// 	errMsg.Msg = "Error in connecting to database"
	// 	return errMsg.Msg
	// }
	// 	defer db.Close()

		// Check for no two Users with same UserName
		rows, err := db.Query("SELECT * FROM users WHERE userName = ?", username)
		if err != nil {}
		if ((rows.Next())) {
		errMsg.Msg = "Username is not unique"
		fmt.Println("Username is not unique")
		return errMsg.Msg
	}
	defer rows.Close()

	hash2, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	
	done := make(chan bool)
	go checkAdminList(db, done)																				// Check If there exist a admin in database or not
	admin := <-done

	if(reqAccess == "adminAccess" ){
		if(admin == false){
			db.Exec("INSERT INTO users (userName, hash, admin) VALUES (?,?,?)", username , hash2 ,1 )		// No admin in database . Make this user Admin
		}else{
			db.Exec("INSERT INTO users (userName, hash, admin) VALUES (?,?,?)", username , hash2 ,0 )		//Add this user to adminReq table and make it client
			_ = db.QueryRow("select id from users where userName = ?", username).Scan(&userId)
			db.Exec("INSERT INTO adminReq ( userId, status) VALUES(?,?)", userId ,0 )
				
		}
	}else{
		db.Exec("INSERT INTO users (userName, hash, admin) VALUES (?,?,?)", username , hash2 ,0 )			//Client user withour adminAccess request
	}

	return "OK"
}


func checkAdminList(db *sql.DB, done chan bool) {
	admin := false

	rows, _ := db.Query("SELECT admin FROM users")
	defer rows.Close()

	for rows.Next() {
		var isAdmin int
		if err := rows.Scan(&isAdmin); err != nil {
			fmt.Println("Error scanning rows:", err)
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


