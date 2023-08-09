package models

import (
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"lib-manager/pkg/types"
	"fmt"
	// "net/http"
	// "html/template"
	// "lib-manager/pkg/views"
	// "time"
	"database/sql"
	// "crypto/rand"
	// "encoding/hex"
)



func RegisterUser( username , password , reEnterPass , reqAccess string) string{
	
	// admin := false
	
	db, err := Connection()
	var errMsg types.ErrMsg
	if err != nil {
		errMsg.Msg = "Error in connecting to database"
		return errMsg.Msg
	}
		defer db.Close()

		rows, err := db.Query("select * from users where userName = ?", username)
		if err != nil {}
		if ((rows.Next())) {
		errMsg.Msg = "Username is not unique"
		fmt.Println("Username is not unique")
		return errMsg.Msg
	}
	defer rows.Close()

	fmt.Println("Above hash2 in register")
	hash2, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	
	done := make(chan bool)
	go checkAdminList(db, done)
	admin := <-done
	var userId int
	fmt.Println(hash2)
	fmt.Println(reqAccess)
	if(reqAccess == "adminAccess" ){
	if(admin == false){
		db.Exec("INSERT INTO users (userName, hash, admin) VALUES (?,?,?)", username , hash2 ,1 )
	}else{
		db.Exec("INSERT INTO users (userName, hash, admin) VALUES (?,?,?)", username , hash2 ,0 )
		
			err = db.QueryRow("select id from users where userName = ?", username).Scan(&userId)
			if err != nil {}else{
				db.Exec("INSERT INTO adminReq ( userId, status) VALUES(?,?)", userId ,0 )
			}
	}
	}else{
		db.Exec("INSERT INTO users (userName, hash, admin) VALUES (?,?,?)", username , hash2 ,0 )
	}

	return "OK"
}


func checkAdminList(db *sql.DB, done chan bool) {
	admin := false

	rows, err := db.Query("SELECT admin FROM users")
	if err != nil {
		fmt.Println("Error querying the database:", err)
		done <- false
		return
	}
	defer rows.Close()

	for rows.Next() {
		var isAdmin int
		if err := rows.Scan(&isAdmin); err != nil {
			fmt.Println("Error scanning rows:", err)
			done <- false
			return
		}
		fmt.Println(isAdmin)
		if isAdmin == 1 {
			admin = true
		}
	}

	done <- admin
	return
}


