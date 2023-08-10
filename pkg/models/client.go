package models

import (
	"net/http"
    "database/sql"
  _ "github.com/go-sql-driver/mysql"
	
)



func Checkout( res http.ResponseWriter, req *http.Request , db *sql.DB, bookId string , userID int) string {

	// db, err := Connection()
    // var errMsg types.ErrMsg
    // if err!= nil {
    //     errMsg.Msg = "Error in connecting to database"
    // }
    // defer db.Close()

	rows, _ := db.Query("SELECT copies FROM books_record WHERE bookId = ?", bookId)
    defer rows.Close()

    for rows.Next() {
        var copies int
        if err := rows.Scan(&copies); err != nil {
            panic(err)
        }
		db.Exec("UPDATE books_record SET copies = copies-1 WHERE bookId = ?", bookId)
	}
	db.Exec("INSERT INTO requests (bookId, userId , status) VALUES(?, ? ,?)", bookId, userID , -1)
	return "OK"
}


func Checkin(res http.ResponseWriter, req *http.Request , db *sql.DB, reqId string) (string ) {
	// db, err := Connection()
    // var errMsg types.ErrMsg
    // if err!= nil {
    //     errMsg.Msg = "Error in connecting to database"
    // }
    // defer db.Close()

    rows, _ := db.Query("SELECT bookId FROM requests WHERE reqId = ?", reqId)
    defer rows.Close()

    for rows.Next() {
        var bookId string
        if err := rows.Scan(&bookId); err != nil {
            panic(err)
        }
		db.Exec("DELETE FROM requests WHERE reqId = ? ", reqId)

		rows, err := db.Query("SELECT copies FROM books_record WHERE bookId = ?", bookId)
		if err!= nil {
			panic(err)
		}
		defer rows.Close()

		for rows.Next() {
			var copies int
			if err := rows.Scan(&copies); err != nil {
				panic(err)
			}
			FinalCopies := copies+1
			db.Exec("UPDATE books_record SET copies =? where bookId =?", FinalCopies, bookId)
		}
    }
	return "OK"
}