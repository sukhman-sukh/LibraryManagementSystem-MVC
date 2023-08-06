package models

import (
	_ "github.com/go-sql-driver/mysql"
	// "golang.org/x/crypto/bcrypt"
	"lib-manager/pkg/types"
	"fmt"
	"net/http"
	// "strconv"
	// "html/template"
	// "lib-manager/pkg/views"
	// "time"
	// "database/sql"
	// "crypto/rand"
	// "encoding/hex"
)



func Checkout( res http.ResponseWriter, req *http.Request , reqId string) string {
	// AdminAddSubmit(res, req , bookname, Author, Copies)
	// admin := false
	
	db, err := Connection()
    var errMsg types.ErrMsg
    if err!= nil {
        errMsg.Msg = "Error in connecting to database"
        // return errMsg.Msg
    }
    defer db.Close()

    rows, err := db.Query("SELECT bookId FROM requests WHERE reqId = ?", reqId)
    if err!= nil {
        panic(err)
    }
    defer rows.Close()

    for rows.Next() {
        var bookId string
       
        if err := rows.Scan(&bookId); err != nil {
            panic(err)
        }
		db.Exec("UPDATE requests SET status = 0 WHERE reqId = ? ", reqId)

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
			// FinalCopies:= strconv.ParseInt(copies, 10, 64)
			fmt.Printf("copies")
			FinalCopies := copies-1
			db.Exec("UPDATE books_record SET copies =? where bookId =?", FinalCopies, bookId)
		}
    }
	return "OK"
}


func Checkin(res http.ResponseWriter, req *http.Request , reqId string) (string ) {
	db, err := Connection()
    var errMsg types.ErrMsg
    if err!= nil {
        errMsg.Msg = "Error in connecting to database"
        // return errMsg.Msg
    }
    defer db.Close()

    rows, err := db.Query("SELECT bookId FROM requests WHERE reqId = ?", reqId)
    if err!= nil {
        panic(err)
    }
    defer rows.Close()

    for rows.Next() {
        var bookId string
        // var copies int
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
			// FinalCopies:= strconv.ParseInt(copies, 10, 64)
			fmt.Printf("copies")
			FinalCopies := copies+1
			db.Exec("UPDATE books_record SET copies =? where bookId =?", FinalCopies, bookId)
		}
    }
	return "OK"
}