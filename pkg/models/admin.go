package models

import (
	"database/sql"
	"lib-manager/pkg/types"
	"net/http"
	"strconv"
    _ "github.com/go-sql-driver/mysql"
)

// Add Books to Database
func AdminAdd(writer http.ResponseWriter, request *http.Request, db *sql.DB, bookname string, Author string, Copies string) string {

	var book types.Books

	rows, _ := db.Query("SELECT * FROM books_record WHERE bookName = ?", bookname)

	if !(rows.Next()) {
		db.Exec("INSERT INTO books_record (bookName, author, copies) VALUES (?, ? ,?)", bookname, Author, Copies)
		return "OK"
	}

	if err := rows.Scan(&book.BookId, &book.BookName, &book.Author, &book.Copies); err != nil {
		panic(err)
	}
	Copy, _ := strconv.Atoi(Copies)
	Copy = book.Copies + Copy
	db.Exec("UPDATE books_record SET copies = ? where bookName = ?", Copy, bookname)

	return "OK"
}

// Remove books from the database
func AdminRemove(writer http.ResponseWriter, request *http.Request, db *sql.DB, bookId string, Copies string) string {

	db.Exec("UPDATE books_record SET copies = ? where bookId = ?", Copies, bookId)
    return "OK"
}

// Fetch Books From Database
func GetBooks(db *sql.DB) (string, []types.Books) {
	var books []types.Books
	var book types.Books

	rows, _ := db.Query("SELECT bookId, bookName, author, copies FROM books_record")
	defer rows.Close()

	for rows.Next() {
		var bookID, bookName, author string
		var copies int
		if err := rows.Scan(&bookID, &bookName, &author, &copies); err != nil {
			panic(err)
		}
		book.BookId = bookID
		book.BookName = bookName
		book.Author = author
		book.Copies = copies
		books = append(books, book)
	}

	// If the datatype is empty
	if len(books) == 0 {
		book.BookId = "empty"
		book.BookName = "empty"
		book.Author = "empty"
		book.Copies = 0
		books = append(books, book)
	}
	return "OK", books
}

// Fetch List of Books Requested for checkout
func GetRequestBooks(db *sql.DB, userId int) (string, []types.RequestBooks) {
	var rows *sql.Rows
	var requestBooks []types.RequestBooks
	var requestBook types.RequestBooks

	rows, _ = db.Query("SELECT * FROM requests WHERE userId=?", userId)
	defer rows.Close()

	for rows.Next() {
		var requestID, bookId, userId, status string

		if err := rows.Scan(&requestID, &bookId, &userId, &status); err != nil {
			panic(err)
		}
		requestBook.RequestId = requestID
		requestBook.BookId = bookId
		requestBook.UserId = userId
		requestBook.Status = status

		requestBooks = append(requestBooks, requestBook)
	}
	// If database is empty
	if len(requestBooks) == 0 {
		requestBook.RequestId = "empty"
		requestBook.BookId = "empty"
		requestBook.UserId = "empty"
		requestBook.Status = "empty"
		requestBooks = append(requestBooks, requestBook)
	}
	return "OK", requestBooks
}

func GetIssuedBooks(db *sql.DB, userId int, admin int) (string, []types.IssuedBook) {
	var rows *sql.Rows
	var issuedBooks []types.IssuedBook
	var issuedBook types.IssuedBook

	if admin == 1 {
		rows, _ = db.Query("SELECT * FROM requests")
	}
	defer rows.Close()

	for rows.Next() {
		var requestID, bookId, userId, status string

		if err := rows.Scan(&requestID, &bookId, &userId, &status); err != nil {
			panic(err)
		}
		issuedBook.RequestId = requestID
		issuedBook.BookId = bookId
		issuedBook.UserId = userId
		issuedBook.Status = status

		issuedBooks = append(issuedBooks, issuedBook)
	}
	// If database is empty
	if len(issuedBooks) == 0 {
		issuedBook.RequestId = "empty"
		issuedBook.BookId = "empty"
		issuedBook.UserId = "empty"
		issuedBook.Status = "empty"
		issuedBooks = append(issuedBooks, issuedBook)
	}
    return "OK", issuedBooks
}

// Fetch list of all clients requesting admin access
func GetAdminRequest(db *sql.DB) (string, []types.AdminRequest) {
	var adminRequests []types.AdminRequest
	var adminRequest types.AdminRequest

	rows, err := db.Query("SELECT * FROM adminReq")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var requestID, userId, status string
		// var copies int
		if err := rows.Scan(&requestID, &userId, &status); err != nil {
			panic(err)
		}
		adminRequest.RequestId = requestID
		adminRequest.UserId = userId
		adminRequest.Status = status

		adminRequests = append(adminRequests, adminRequest)
	}

	if len(adminRequests) == 0 {

		adminRequest.RequestId = "empty"
		adminRequest.UserId = "empty"
		adminRequest.Status = "empty"

		adminRequests = append(adminRequests, adminRequest)
	}
	return "OK", adminRequests
}

// Approve checkin of books requested by the user by the admin
func AdminCheckin(writer http.ResponseWriter, request *http.Request, db *sql.DB, requestId string) string {

	rows, err := db.Query("SELECT bookId FROM requests WHERE reqId = ?", requestId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var bookId string
		if err := rows.Scan(&bookId); err != nil {
			panic(err)
		}
		db.Exec("DELETE FROM requests WHERE reqId = ? ", requestId)

		rows, err := db.Query("SELECT copies FROM books_record WHERE bookId = ?", bookId)
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		for rows.Next() {
			var copies int
			if err := rows.Scan(&copies); err != nil {
				panic(err)
			}
			FinalCopies := copies + 1
			db.Exec("UPDATE books_record SET copies =? where bookId =?", FinalCopies, bookId)
		}
	}
	return "OK"
}

// Approve checkoiut of books requested by the user by the admin
func AdminCheckout(writer http.ResponseWriter, request *http.Request, db *sql.DB, requestId string) string {

	rows, _ := db.Query("SELECT bookId FROM requests WHERE reqId = ?", requestId)
	defer rows.Close()

	for rows.Next() {
		var bookId string
		if err := rows.Scan(&bookId); err != nil {
			panic(err)
		}

		db.Exec("UPDATE requests SET status = 0 WHERE reqId = ? ", requestId)

		rows, _ := db.Query("SELECT copies FROM books_record WHERE bookId = ?", bookId)
		defer rows.Close()

		for rows.Next() {
			var copies int
			if err := rows.Scan(&copies); err != nil {
				panic(err)
			}
			FinalCopies := copies - 1
			db.Exec("UPDATE books_record SET copies =? where bookId =?", FinalCopies, bookId)
		}
	}
	return "OK"
}

// Accept admin request
func AdminAccept(writer http.ResponseWriter, request *http.Request, db *sql.DB, requestId string) string {

	rows, err := db.Query("SELECT userId FROM adminReq WHERE reqId = ?", requestId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var userId string
		if err := rows.Scan(&userId); err != nil {
			panic(err)
		}
		db.Exec("UPDATE users SET admin = 1 WHERE id = ? ", userId)
		db.Exec("DELETE FROM adminReq WHERE reqId = ? ", requestId)

	}
	return "OK"
}

// Deny admin request
func AdminDeny(writer http.ResponseWriter, request *http.Request, db *sql.DB, requestId string) string {
	db.Exec("DELETE FROM adminReq WHERE reqId = ? ", requestId)
	return "OK"
}
