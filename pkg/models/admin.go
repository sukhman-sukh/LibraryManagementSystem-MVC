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



func AdminAdd( res http.ResponseWriter, req *http.Request , bookname string, Author string, Copies string) string {
	// AdminAddSubmit(res, req , bookname, Author, Copies)
	// admin := false
	
	db, err := Connection()
	var errMsg types.ErrMsg
	if err != nil {
		errMsg.Msg = "Error in connecting to database"
		return errMsg.Msg
	}
	defer db.Close()

	rows, err := db.Query("select * from books_record where bookName = ?", bookname)
	if(IsDbEmpty("cookie" , db)){
		fmt.Println("1st entry of the book")
		db.Exec("INSERT INTO books_record (bookName, author, copies) VALUES (?, ? ,?)", bookname, Author , Copies)
		return "OK"
	}	
	fmt.Println("sds1")
	for rows.Next() {
		fmt.Println("Updating existting book")
		db.Exec("UPDATE books_record SET copies = ? where bookName = ?", Copies, bookname)
		
	}

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

func GetBooks(res http.ResponseWriter, req *http.Request) (string , []map[string]interface{}) {
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


	var books []map[string]interface{}

    for rows.Next() {
        var bookID, bookName, author string
        var copies int
        if err := rows.Scan(&bookID, &bookName, &author, &copies); err != nil {
            panic(err)
        }
        books = append(books, map[string]interface{}{
            "bookId":   bookID,
            "bookName": bookName,
            "author":   author,
            "copies":   copies,
        })
    }

    if len(books) == 0 {
        books = append(books, map[string]interface{}{
            "bookId":   "empty",
            "bookName": "empty",
            "author":   "empty",
            "copies":   "empty",
        })
    }

    fmt.Println(books)

	return "OK" , books

}


func GetReqBooks(res http.ResponseWriter, req *http.Request) (string , []map[string]interface{}) {
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

	var reqBook []map[string]interface{}

    for rows.Next() {
        var reqID, date, bookId , userId , status string
        // var copies int
        if err := rows.Scan(&reqID, &date, &bookId , &userId , &status); err != nil {
            panic(err)
        }
        reqBook = append(reqBook, map[string]interface{}{

			"reqID": reqID,
			"date": date, 
			"bookId" : bookId, 
			"userId":userId , 
			"status" :status,
		})
    }

    if len(reqBook) == 0 {
        reqBook = append(reqBook, map[string]interface{}{
			"reqID": "empty",
			"date": "empty", 
			"bookId" :"empty", 
			"userId":"empty", 
			"status" :"empty",
        })
    }

    fmt.Println(reqBook)

	return "OK" , reqBook
}


func GetAdminReq(res http.ResponseWriter, req *http.Request) (string , []map[string]interface{}) {
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

	var adminReq []map[string]interface{}

    for rows.Next() {
        var reqID, userId , status string
        // var copies int
        if err := rows.Scan(&reqID, &userId , &status); err != nil {
            panic(err)
        }
        adminReq = append(adminReq, map[string]interface{}{

			"reqID": reqID,
			"userId":userId , 
			"status" :status,
		})
    }

    if len(adminReq) == 0 {
        adminReq = append(adminReq, map[string]interface{}{
			"reqID": "empty", 
			"userId":"empty", 
			"status" :"empty",
        })
    }

    fmt.Println(adminReq)

	return "OK" , adminReq
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