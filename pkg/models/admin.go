package models

import (
	_ "github.com/go-sql-driver/mysql"
	// "golang.org/x/crypto/bcrypt"
	"lib-manager/pkg/types"
	"fmt"
	"net/http"
	"strconv"
	// "html/template"
	// "lib-manager/pkg/types"
	// "time"
	// "database/sql"
	// "crypto/rand"
	// "encoding/hex"
)



func AdminAdd( res http.ResponseWriter, req *http.Request , bookname string, Author string, Copies string) string {
	// AdminAddSubmit(res, req , bookname, Author, Copies)
	// admin := false

    var book types.Books
	db, err := Connection()
	var errMsg types.ErrMsg
	if err != nil {
		errMsg.Msg = "Error in connecting to database"
        fmt.Println(errMsg.Msg)
		return errMsg.Msg
	}
	defer db.Close()

    fmt.Println("add Books..")

	rows, err := db.Query("select * from books_record where bookName = ?", bookname)
    
	if(!(rows.Next())){
		fmt.Println("1st entry of the book")
		db.Exec("INSERT INTO books_record (bookName, author, copies) VALUES (?, ? ,?)", bookname, Author , Copies)
		return "OK"
	}
    
    if err := rows.Scan(&book.BookId, &book.BookName, &book.Author, &book.Copies); err != nil {
        panic(err)
    }
    
    // fmt.Println("sds1")
    
    Copy, _ := strconv.Atoi(Copies)
    Copy = book.Copies+Copy
    fmt.Println("Updating existting book  " , Copy)
    db.Exec("UPDATE books_record SET copies = ? where bookName = ?", Copy , bookname)	
    
	return "OK"
}

func AdminRemove( res http.ResponseWriter, req *http.Request , bookId string,Copies string) (string ) {
	// AdminAddSubmit(res, req , bookname, Author, Copies)
	// admin := false
	
	db, err := Connection()
	var errMsg types.ErrMsg
	if err != nil {
		errMsg.Msg = "Error in connecting to database"
		return errMsg.Msg 
	}
	defer db.Close()

	db.Exec("UPDATE books_record SET copies = ? where bookId = ?", Copies, bookId)

	return "OK"
}

func GetBooks(res http.ResponseWriter, req *http.Request) (string , []types.Books) {
	db, err := Connection()
	var errMsg types.ErrMsg
	if err != nil {
		errMsg.Msg = "Error in connecting to database"
		// return errMsg.Msg
	}
	defer db.Close()

	rows, err := db.Query("SELECT bookId, bookName, author, copies FROM books_record")
    if err != nil {
        panic(err)
    }
    defer rows.Close()


	// var books []map[string]interface{}
    var books []types.Books
    var book types.Books

    for rows.Next() {
        var bookID , bookName, author string
        var copies int
        if err := rows.Scan(&bookID, &bookName, &author, &copies); err != nil {
            panic(err)
        }
        book.BookId = bookID
        book.BookName = bookName
        book.Author = author
        book.Copies = copies
        books = append(books , book)
    }

    if len(books) == 0 {
        book.BookId = "empty"
        book.BookName = "empty"
        book.Author = "empty"
        book.Copies = 0
    
    }
    books = append(books , book)
    fmt.Println(books)

	return "OK" , books

}


func GetReqBooks(res http.ResponseWriter, req *http.Request) (string , []types.ReqBooks) {
	db, err := Connection()
    var errMsg types.ErrMsg
    if err!= nil {
        errMsg.Msg = "Error in connecting to database"
        // return errMsg.Msg
    }
    defer db.Close()

    rows, err := db.Query("SELECT * FROM requests")
    if err!= nil {
        panic(err)
    }
    defer rows.Close()

    var reqBooks []types.ReqBooks
    var reqBook types.ReqBooks

    for rows.Next() {
        var reqID, date, bookId , userId , status string

        if err := rows.Scan(&reqID, &date, &bookId , &userId , &status); err != nil {
            panic(err)
        }
        reqBook.ReqId = reqID
        reqBook.Date = date
        reqBook.BookId = bookId
        reqBook.UserId = userId
        reqBook.Status = status

        reqBooks = append(reqBooks , reqBook)    
    }
    
    if len(reqBooks) == 0 {
        reqBook.ReqId = "empty"
        reqBook.Date = "empty"
        reqBook.BookId = "empty"
        reqBook.UserId = "empty"
        reqBook.Status = "empty"
        reqBooks = append(reqBooks , reqBook)
    }
    fmt.Println(reqBooks)

	return "OK" , reqBooks

}


func GetAdminReq(res http.ResponseWriter, req *http.Request) (string , []types.AdminReq) {
	db, err := Connection()
    var errMsg types.ErrMsg
    if err!= nil {
        errMsg.Msg = "Error in connecting to database"
        // return errMsg.Msg
    }
    defer db.Close()

    rows, err := db.Query("SELECT * FROM adminReq")
    if err!= nil {
        panic(err)
    }
    defer rows.Close()

    var adminReqs []types.AdminReq
    var adminReq types.AdminReq

    for rows.Next() {
        var reqID, userId , status string
        // var copies int
        if err := rows.Scan(&reqID, &userId , &status); err != nil {
            panic(err)
        }
        adminReq.ReqId = reqID
        adminReq.UserId = userId
        adminReq.Status = status

        adminReqs = append(adminReqs , adminReq)    
    }
    
    if len(adminReqs) == 0 {

        adminReq.ReqId = "empty"
        adminReq.UserId = "empty"
        adminReq.Status = "empty"

        adminReqs = append(adminReqs , adminReq)    
    }
    fmt.Println(adminReqs)

	return "OK" , adminReqs
}


func AdminCheckin(res http.ResponseWriter, req *http.Request , reqId string) (string ) {
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

func AdminCheckout(res http.ResponseWriter, req *http.Request , reqId string) (string ) {
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


func AdminAccept(res http.ResponseWriter, req *http.Request , reqId string) (string ) {
	db, err := Connection()
    var errMsg types.ErrMsg
    if err!= nil {
        errMsg.Msg = "Error in connecting to database"
        // return errMsg.Msg
    }
    defer db.Close()

    rows, err := db.Query("SELECT userId FROM adminReq WHERE reqId = ?", reqId)
    if err!= nil {
        panic(err)
    }
    defer rows.Close()

    for rows.Next() {
        var userId string
        // var copies int
        if err := rows.Scan(&userId); err != nil {
            panic(err)
        }
		db.Exec("UPDATE users SET admin = 1 WHERE id = ? ", userId)
		db.Exec("DELETE FROM adminReq WHERE reqId = ? ", reqId)

    }
	return "OK"
}


func AdminDeny(res http.ResponseWriter, req *http.Request , reqId string) (string ) {
	db, err := Connection()
    var errMsg types.ErrMsg
    if err!= nil {
        errMsg.Msg = "Error in connecting to database"
        // return errMsg.Msg
    }
    defer db.Close()
		db.Exec("DELETE FROM adminReq WHERE reqId = ? ", reqId)
	return "OK"
}