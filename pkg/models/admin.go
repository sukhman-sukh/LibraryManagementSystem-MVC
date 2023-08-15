package models

import (
	"strconv"
	"net/http"
    "database/sql"
  	"lib-manager/pkg/types"
  _ "github.com/go-sql-driver/mysql"
	
)


// Add Books to Database
func AdminAdd( res http.ResponseWriter, req *http.Request, db *sql.DB , bookname string, Author string, Copies string) string {

    var book types.Books
    
	// db, err := Connection()
	// var errMsg types.ErrMsg
	// if err != nil {
	// 	errMsg.Msg = "Error in connecting to database"
	// 	return errMsg.Msg
	// }
	// defer db.Close()


	rows, _ := db.Query("SELECT * FROM books_record WHERE bookName = ?", bookname)
    
	if(!(rows.Next())){
		db.Exec("INSERT INTO books_record (bookName, author, copies) VALUES (?, ? ,?)", bookname, Author , Copies)
		return "OK"
	}
    
    if err := rows.Scan(&book.BookId, &book.BookName, &book.Author, &book.Copies); err != nil {
        panic(err)
    }
    Copy, _ := strconv.Atoi(Copies)
    Copy = book.Copies+Copy
    db.Exec("UPDATE books_record SET copies = ? where bookName = ?", Copy , bookname)	
    
	return "OK"
}

// Remove books from the database
func AdminRemove( res http.ResponseWriter, req *http.Request , db *sql.DB , bookId string,Copies string) (string ) {

	// db, err := Connection()
	// var errMsg types.ErrMsg
	// if err != nil {
	// 	errMsg.Msg = "Error in connecting to database"
	// 	return errMsg.Msg 
	// }
	// defer db.Close()

	db.Exec("UPDATE books_record SET copies = ? where bookId = ?", Copies, bookId)

	return "OK"
}

// Fetch Books From Database
func GetBooks( db *sql.DB ) (string , []types.Books) {
    var books []types.Books
    var book types.Books

	// db, err := Connection()
	// var errMsg types.ErrMsg
	// if err != nil {
	// 	errMsg.Msg = "Error in connecting to database"
	// }
	// defer db.Close()

	rows, _ := db.Query("SELECT bookId, bookName, author, copies FROM books_record")
    defer rows.Close()
    
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

    // If the datatype is empty
    if len(books) == 0 {
        book.BookId = "empty"
        book.BookName = "empty"
        book.Author = "empty"
        book.Copies = 0
        books = append(books , book)
    }

	return "OK" , books

}

// Fetch List of Books Requested for checkout
func GetReqBooks(db *sql.DB , userId int) (string , []types.ReqBooks) {
    var rows *sql.Rows
    var reqBooks []types.ReqBooks
    var reqBook types.ReqBooks

	// db, err := Connection()
    // var errMsg types.ErrMsg
    // if err!= nil {
    //     errMsg.Msg = "Error in connecting to database"
    // }
    // defer db.Close()


    rows, _ = db.Query("SELECT * FROM requests WHERE userId=?", userId)
    defer rows.Close()

    for rows.Next() {
        var reqID, bookId , userId , status string

        if err := rows.Scan(&reqID, &bookId , &userId , &status); err != nil {
            panic(err)
        }
        reqBook.ReqId = reqID
        reqBook.BookId = bookId
        reqBook.UserId = userId
        reqBook.Status = status

        reqBooks = append(reqBooks , reqBook)    
    }
    // If database is empty
    if len(reqBooks) == 0 {
        reqBook.ReqId = "empty"
        reqBook.BookId = "empty"
        reqBook.UserId = "empty"
        reqBook.Status = "empty"
        reqBooks = append(reqBooks , reqBook)
    }

	return "OK" , reqBooks

}

func GetIssuedBooks(db *sql.DB , userId int ,  admin int) (string , []types.IssuedBook ) {
    var rows *sql.Rows
    var issuedBooks []types.IssuedBook
    var issuedBook types.IssuedBook

	// db, err := Connection()
    // var errMsg types.ErrMsg
    // if err!= nil {
    //     errMsg.Msg = "Error in connecting to database"
    // }
    // defer db.Close()

    if(admin == 1){
        rows, _ = db.Query("SELECT * FROM requests")
    }
    defer rows.Close()

    for rows.Next() {
        var reqID, bookId , userId , status string

        if err := rows.Scan(&reqID, &bookId , &userId , &status); err != nil {
            panic(err)
        }
        issuedBook.ReqId = reqID
        issuedBook.BookId = bookId
        issuedBook.UserId = userId
        issuedBook.Status = status

        issuedBooks = append(issuedBooks , issuedBook)    
    }
    // If database is empty
    if len(issuedBooks) == 0 {
        issuedBook.ReqId = "empty"
        issuedBook.BookId = "empty"
        issuedBook.UserId = "empty"
        issuedBook.Status = "empty"
        issuedBooks = append(issuedBooks , issuedBook)
    }

	return "OK" , issuedBooks

}


// Fetch list of all clients requesting admin access
func GetAdminReq( db *sql.DB) (string , []types.AdminReq) {
    var adminReqs []types.AdminReq
    var adminReq types.AdminReq

	// db, err := Connection()
    // var errMsg types.ErrMsg
    // if err!= nil {
    //     errMsg.Msg = "Error in connecting to database"
    // }
    // defer db.Close()

    rows, err := db.Query("SELECT * FROM adminReq")
    if err!= nil {
        panic(err)
    }
    defer rows.Close()


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
	return "OK" , adminReqs
}

// Approve checkin of books requested by the user by the admin
func AdminCheckin(res http.ResponseWriter, req *http.Request , db *sql.DB, reqId string) (string ) {
	// db, err := Connection()
    // var errMsg types.ErrMsg
    // if err!= nil {
    //     errMsg.Msg = "Error in connecting to database"
    // }
    // defer db.Close()

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

// Approve checkoiut of books requested by the user by the admin
func AdminCheckout(res http.ResponseWriter, req *http.Request , db *sql.DB , reqId string) (string ) {

    rows, _ := db.Query("SELECT bookId FROM requests WHERE reqId = ?", reqId)
    defer rows.Close()

    for rows.Next() {
        var bookId string
        if err := rows.Scan(&bookId); err != nil {
            panic(err)
        }
        
		db.Exec("UPDATE requests SET status = 0 WHERE reqId = ? ", reqId)

		rows, _ := db.Query("SELECT copies FROM books_record WHERE bookId = ?", bookId)
		defer rows.Close()

		for rows.Next() {
			var copies int
			if err := rows.Scan(&copies); err != nil {
				panic(err)
			}
			FinalCopies := copies-1
			db.Exec("UPDATE books_record SET copies =? where bookId =?", FinalCopies, bookId)
		}
    }
	return "OK"
}

// Accept admin request
func AdminAccept(res http.ResponseWriter, req *http.Request , db *sql.DB, reqId string) (string ) {
	// db, err := Connection()
    // var errMsg types.ErrMsg
    // if err!= nil {
    //     errMsg.Msg = "Error in connecting to database"
    // }
    // defer db.Close()

    rows, err := db.Query("SELECT userId FROM adminReq WHERE reqId = ?", reqId)
    if err!= nil {
        panic(err)
    }
    defer rows.Close()

    for rows.Next() {
        var userId string
        if err := rows.Scan(&userId); err != nil {
            panic(err)
        }
		db.Exec("UPDATE users SET admin = 1 WHERE id = ? ", userId)
		db.Exec("DELETE FROM adminReq WHERE reqId = ? ", reqId)

    }
	return "OK"
}

// Deny admin request
func AdminDeny(res http.ResponseWriter, req *http.Request , db *sql.DB , reqId string) (string ) {
	// db, err := Connection()
    // var errMsg types.ErrMsg
    // if err!= nil {
    //     errMsg.Msg = "Error in connecting to database"
    // }
    // defer db.Close()
		db.Exec("DELETE FROM adminReq WHERE reqId = ? ", reqId)
	return "OK"
}